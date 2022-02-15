package app

import (
	"context"
	"log"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/outages"
	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

type OutageUISP models.Outage

type Outage struct {
	NN int
	OutageUISP
	HealthStatus
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

type OutageMap struct {
	data map[int][]Outage
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

func (om *OutageMap) NodeAggregatedHealth(nn int) (outageStart *time.Time, numUnreachable int, numOutage int) {
	outages, ok := om.data[nn]
	if !ok {
		return
	}

	for _, o := range outages {
		o := o
		if outageStart == nil || outageStart.After(o.Start) {
			outageStart = &o.Start
		}

		switch o.Type {
		case "outage":
			numOutage = numOutage + 1
		case "unreachable":
			numUnreachable = numUnreachable + 1
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

func (a *App) GetFullOutageMap(ctx context.Context) (OutageMap, error) {
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

func (a *App) outageConsumer(ctx context.Context, wg *sync.WaitGroup, fountain <-chan OutageMap) {

	// read from outages that are valid to us, and update our status map
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				break
			case outageMap := <-fountain:
				for _, nn := range outageMap.ImpactedNNs() {
					outageStart, nunreachable, noutage := outageMap.NodeAggregatedHealth(nn)
					if outageStart == nil {
						continue
					}

					outageDevices := outageMap.GetOutagesDeviceNamesByOutageType(nn, "outage")
					unreachableDevices := outageMap.GetOutagesDeviceNamesByOutageType(nn, "unreachable")

					dur := time.Since(*outageStart)

					if dur > ignoreOutagesLongerThan {
						//log.Printf("nn:%d has outage > %s, ignoring", nn, ignoreOutagesLongerThan)
						continue
					}
					switch {
					case noutage > 0 && nunreachable == 0:
						log.Printf("nn:%d: degraded for %s (%d devices in outage: %s)", nn, dur.String(), len(outageDevices), strings.Join(outageDevices, ","))
					case noutage > 0 && nunreachable > 0:
						log.Printf("nn:%d: degraded for %s (%d devices in outage: %s, %d devices unreachable: %s)", nn, dur.String(), len(outageDevices), strings.Join(outageDevices, ","), len(unreachableDevices), strings.Join(unreachableDevices, ","))
					case noutage == 0 && nunreachable > 0:
						log.Printf("nn:%d: degraded for %s (%d devices unreachable: %s)", nn, dur.String(), len(unreachableDevices), strings.Join(unreachableDevices, ","))
					default:
						log.Printf("nn:%d: online", nn)
					}
				}

				if !a.config.Daemon.EnableSlack {
					log.Printf("slack support disabled via --enable-slack=false - not updating slack hub topics with status")
					continue
				}

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
			log.Printf("fetching outage map from UISP")
			om, err := a.GetFullOutageMap(ctx)
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
