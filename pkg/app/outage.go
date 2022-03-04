package app

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/outages"
	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

type OutageUISP models.Outage

// Outage is a single UISP Outage, which impacts a single Device/Site pair
type Outage struct {
	NN int
	OutageUISP
	Start time.Time
}

var (
	outageMapPollInterval         = 60 * time.Second
	outageMapErrorBackoffInterval = 30 * time.Second
	ignoreOutagesLongerThan       = 60 * 24 * time.Hour // 60 days
)

type HealthStatus string

const (
	HealthStatusOnline      HealthStatus = "online"
	HealthStatusUnreachable HealthStatus = "unreachable"
	HealthStatusDegraded    HealthStatus = "degraded"
	HealthStatusOffline     HealthStatus = "offline"
	HealthStatusUnknown     HealthStatus = "unknown"
)

// OutageMap maps a NN to a list of device outages. Each OutageMap is a
// snapshot of UISP outage data at a point in time
type OutageMap struct {
	data map[int][]Outage
}

// outageNNGroups takes a current and previous outage, and returns the list of NNs that:
// newOutageNNs: NNs that are newly in the Outage state (at least 1 device in current map is unavailable)
// clearedNNs: NNs that were in outage, but now are all good
// prevImpactedNNs: all NNs impacted in the previous OutageMap
// currentImpactedNNs: all NNs impacted in the current OutageMap
func outageNNGroups(currOutageMap, prevOutageMap OutageMap) (newOutageNNs, clearedNNs, prevImpactedNNs, currentImpactedNNs []int) {
	prevImpactedNNs = prevOutageMap.ImpactedNNs()
	currentImpactedNNs = currOutageMap.ImpactedNNs()

	for _, nn := range currentImpactedNNs {
		if !contains(prevImpactedNNs, nn) {
			newOutageNNs = append(newOutageNNs, nn)
		}
	}

	for _, nn := range prevImpactedNNs {
		if !contains(currentImpactedNNs, nn) {
			clearedNNs = append(clearedNNs, nn)
		}
	}

	sort.Ints(newOutageNNs)
	sort.Ints(clearedNNs)
	sort.Ints(prevImpactedNNs)
	sort.Ints(currentImpactedNNs)

	return newOutageNNs, clearedNNs, prevImpactedNNs, currentImpactedNNs
}

func (om *OutageMap) ImpactedNNs() []int {
	impacted := []int{}
	for k := range om.data {
		impacted = append(impacted, k)
	}

	sort.Slice(impacted, func(i, j int) bool {
		return i < j
	})

	return impacted
}

func (om *OutageMap) GetOutagesDeviceNamesByOutageType(nn int, outageType string) []string {
	devOutages, err := om.NodeDeviceOutages(nn)
	if err != nil {
		return []string{}
	}
	// maps outageType to list of device names in that outage state
	m := map[string][]string{}

	// collect all devices by outage type
	for _, o := range devOutages {
		if o.Device == nil || o.Device.Name == nil {
			// not a device outage, not sure what to do here
			continue
		}
		if devs, ok := m[o.Type]; ok {
			m[o.Type] = append(devs, *o.Device.Name)
		} else {
			m[o.Type] = []string{*o.Device.Name}
		}
	}

	res := m[outageType]
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return res
}

func (om *OutageMap) OutageStartTime(nn int) (outageStart *time.Time, err error) {
	outages, ok := om.data[nn]
	if !ok {
		return nil, fmt.Errorf("no outage for nn:%d found", nn)
	}

	for _, o := range outages {
		o := o
		if outageStart == nil || outageStart.After(o.Start) {
			outageStart = &o.Start
		}
	}
	return
}

func (om *OutageMap) NodeDeviceOutages(nn int) ([]Outage, error) {
	if os, ok := om.data[nn]; ok {
		return os, nil
	}
	// TODO: if node not found is this error, or just no outages?
	return []Outage{}, nil
}

func (a *App) FetchOutageMap(ctx context.Context) (OutageMap, error) {
	inProgress := true
	params := outages.NewGetOutagesParams().
		WithDefaults().
		WithContext(ctx).
		WithCount(1000).
		WithPage(float64(1)).
		WithInProgress(&inProgress)

	data := []Outage{}
	for {
		res, err := a.UISPAPI.Outages.GetOutages(params, nil)
		if err != nil {
			log.Printf("fetching outages failed: %s", err.Error())
			time.Sleep(logFetchErrorBackoff)
			continue
		}

		for _, outage := range res.Payload.Items {
			if outage.Ongoing != nil && !*outage.Ongoing {
				// skip outages that are concluded, if any leaked through our query
				continue
			}

			outageStart, err := time.Parse(time.RFC3339, *outage.StartTimestamp)
			if err != nil {
				log.Printf("unable to parse outage start timestamp %s on outage %s: %s", *outage.StartTimestamp, *outage.ID, err.Error())
				continue
			}

			dat := Outage{
				Start:      outageStart,
				OutageUISP: OutageUISP(*outage),
			}

			{ // TODO: this code should be extracted into a helper function
				if outage.Device != nil {
					// try to find NN from device name
					if nn, err := GetNNFromDeviceName(*outage.Device.Name); err == nil && dat.NN == 0 {
						dat.NN = nn
					}
				}

				if outage.Site != nil {
					// try to find NN from device name
					if nn, err := GetNNFromSiteName(*outage.Site.Name); err == nil && dat.NN == 0 {
						dat.NN = nn
					}
					// TODO: recurse to aggregate all parent site identifiers as well
				}
			}

			data = append(data, dat)
		}

		if res.Payload.Pagination.Pages != nil && float64(*res.Payload.Pagination.Pages) == params.Page {
			// all done with pagination, lets wrap up this fetch
			om := OutageMap{data: map[int][]Outage{}}
			for _, d := range data {
				if existing, ok := om.data[d.NN]; ok {
					om.data[d.NN] = append(existing, d)

				} else {
					om.data[d.NN] = []Outage{d}
				}
			}
			return om, nil
		}

		// set the params for the next page
		params.SetPage(float64(*res.Payload.Pagination.Page + 1))
	}
}

func contains(elems []int, v int) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func allNNs(outageMap OutageMap, prevOutageMap OutageMap) []int {
	nns := map[int]interface{}{}
	for _, nn := range outageMap.ImpactedNNs() {
		nns[nn] = struct{}{}
	}
	for _, nn := range prevOutageMap.ImpactedNNs() {
		nns[nn] = struct{}{}
	}

	l := []int{}
	for nn, _ := range nns {
		l = append(l, nn)
	}

	sort.Ints(l)
	return l

}

func (a *App) processNewOutageMap(outageMap OutageMap, prevOutageMap OutageMap) error {
	// determine new outages only
	newOutageNNs, clearedNNs, _, currentImpactedNNs := outageNNGroups(outageMap, prevOutageMap)

	for _, nn := range clearedNNs {
		if outageBegin, err := prevOutageMap.OutageStartTime(nn); err != nil {
			log.Printf("nn:%d outage resolved (unable to verify outage duration: %s)", nn, err.Error())
		} else {
			log.Printf("nn:%d outage resolved after %s", nn, time.Since(*outageBegin).String())
		}
	}

	for _, nn := range newOutageNNs {
		outageStart, err := outageMap.OutageStartTime(nn)
		if err != nil {
			log.Printf("unable to determine outage start time: %s", err.Error())
			continue
		}

		outageDevices := outageMap.GetOutagesDeviceNamesByOutageType(nn, "outage")
		unreachableDevices := outageMap.GetOutagesDeviceNamesByOutageType(nn, "unreachable")

		dur := time.Since(*outageStart)

		durString := fmt.Sprintf("%s", dur.String())
		// prettify how we show long outages, so we dont see tons of microseconds for 1w long outages
		switch {
		case dur.Hours() > 7*24:
			durString = fmt.Sprintf("%.1f weeks", dur.Round(time.Hour).Hours()/(24*7))
		case dur.Hours() > 24:
			durString = fmt.Sprintf("%.1f days", dur.Round(time.Hour).Hours()/24)
		case dur.Minutes() > 10:
			durString = fmt.Sprintf("%s", dur.Round(time.Minute))
		default:
			durString = fmt.Sprintf("%s", dur.Round(time.Second))
		}

		if dur > ignoreOutagesLongerThan {
			//log.Printf("nn:%d has outage > %s, ignoring", nn, ignoreOutagesLongerThan)
			continue
		}

		msgs := []string{}
		if len(outageDevices) > 0 {
			msgs = append(msgs, fmt.Sprintf("%d in outage: %s", len(outageDevices), strings.Join(outageDevices, ",")))
		}
		if len(unreachableDevices) > 0 {
			msgs = append(msgs, fmt.Sprintf("%d unreachable: %s", len(unreachableDevices), strings.Join(unreachableDevices, ",")))
		}
		log.Printf("nn:%d in outage for %s (%s)", nn, durString, strings.Join(msgs, ", "))
	}

	for _, nn := range currentImpactedNNs {
		if !contains(newOutageNNs, nn) {
			currOutages, _ := outageMap.NodeDeviceOutages(nn)
			prevOutages, _ := prevOutageMap.NodeDeviceOutages(nn)

			outageDevices := outageMap.GetOutagesDeviceNamesByOutageType(nn, "outage")
			unreachableDevices := outageMap.GetOutagesDeviceNamesByOutageType(nn, "unreachable")
			msgs := []string{}
			if len(outageDevices) > 0 {
				msgs = append(msgs, fmt.Sprintf("%d in outage: %s", len(outageDevices), strings.Join(outageDevices, ",")))
			}
			if len(unreachableDevices) > 0 {
				msgs = append(msgs, fmt.Sprintf("%d unreachable: %s", len(unreachableDevices), strings.Join(unreachableDevices, ",")))
			}

			delta := len(prevOutages) - len(currOutages)
			switch {
			case delta > 0:
				log.Printf("nn:%d outage continues but got better: %s", nn, strings.Join(msgs, ", "))
			case delta < 0:
				log.Printf("nn:%d outage intensifies: %s", nn, strings.Join(msgs, ", "))
			default:
				// no change in device outage states, ignore
			}
		}
	}

	return nil
}

func (a *App) outageConsumer(ctx context.Context, wg *sync.WaitGroup, fountain <-chan OutageMap) {

	// read from outages that are valid to us, and update our status map
	wg.Add(1)
	go func() {
		defer wg.Done()
		var prevOutageMap OutageMap

		for {
			select {
			case <-ctx.Done():
				break
			case outageMap := <-fountain:
				if err := a.processNewOutageMap(outageMap, prevOutageMap); err != nil {
					log.Printf("error processing outage map: %s", err.Error())
				}

				if a.config.Daemon.EnableSlack {
					log.Printf("TODO @gabe enable when ready to turn on outage notifications into slack threads")
					//log.Printf("syncing outages to slack threads")
					//baseURL := fmt.Sprintf("%s://%s/nms", a.config.UISP.Scheme, a.config.UISP.Hostname)
					//if err := a.updateSlackOutageThreadsFromOutageMaps(baseURL, outageMap, prevOutageMap); err != nil {
					//	log.Printf("error syncing outage map to slack: %s", err.Error())
					//}
				}

				// record our previous outage map, so we can detect new vs existing outages
				prevOutageMap = outageMap

			}
		}
	}()
}

func (a *App) runOutageInformer(ctx context.Context) error {
	wg := sync.WaitGroup{}
	fountain := make(chan OutageMap, 1)

	consumers := []func(context.Context, *sync.WaitGroup, <-chan OutageMap){}
	if a.config.Daemon.OutageDetection {
		consumers = append(consumers, a.outageConsumer)
	}

	a.startOutageMapProducer(ctx, &wg, fountain)
	outputChannels := a.startOutageConsumerMultiplexer(ctx, &wg, fountain, len(consumers))
	for i, startCoroutine := range consumers {
		startCoroutine(ctx, &wg, outputChannels[i])
	}

	wg.Wait()
	return nil
}

func (a *App) startOutageMapProducer(ctx context.Context, wg *sync.WaitGroup, ch chan<- OutageMap) {
	wg.Add(1)
	ctxDone := ctx.Done()

	go func() {
		defer func() {
			close(ch)
			wg.Done()
		}()

		doFetch := func() {
			om, err := a.FetchOutageMap(ctx)
			if err != nil {
				log.Printf("error getting outage map: %s", err.Error())
				time.Sleep(outageMapErrorBackoffInterval)
			}
			ch <- om
		}

		ticker := time.Tick(outageMapPollInterval)
		doFetch()
		for {
			select {
			case <-ctxDone:
				break
			case <-ticker:
				doFetch()
			}
		}
	}()
}

// TODO: good candidate for genericization with startLogConsumerMultiplexer
func (a *App) startOutageConsumerMultiplexer(ctx context.Context, wg *sync.WaitGroup, srcCh <-chan OutageMap, numChannelsToMultiplexTo int) []chan OutageMap {
	wg.Add(1)
	ctxDone := ctx.Done()

	downstreamChannels := make([]chan OutageMap, numChannelsToMultiplexTo)
	for i := range downstreamChannels {
		downstreamChannels[i] = make(chan OutageMap)
	}

	go func() {
		defer func() {
			for _, ch := range downstreamChannels {
				close(ch)
			}
			wg.Done()
		}()

		for {
			select {
			case <-ctxDone:
				break
			case evt := <-srcCh:
				// multiplex this log to all our downstreamChannels
				for _, consumerCh := range downstreamChannels {
					consumerCh <- evt
				}
			}
		}
	}()

	return downstreamChannels
}
