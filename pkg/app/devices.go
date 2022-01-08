package app

import (
	"fmt"
	"log"
	"strconv"

	"github.com/byxorna/nycmesh-tool/client/devices"
	"github.com/byxorna/nycmesh-tool/client/sites"
	"github.com/byxorna/nycmesh-tool/pkg/nycmesh"
)

func (a *App) Devices(ids ...string) (map[int]*nycmesh.Device, error) {
	nodes, err := a.Nodes()
	if err != nil {
		return nil, fmt.Errorf("unable to fetch devices: %w", err)
	}

	idFilter := map[int]interface{}{}
	for _, idarg := range ids {
		n, err := strconv.Atoi(idarg)
		if err != nil {
			return nil, fmt.Errorf("%s is not a valid device identifier: %w", idarg, err)
		}
		idFilter[n] = true
	}

	devs := map[int]*nycmesh.Device{}
	for _, n := range nodes {
		for _, d := range n.Devices {
			locald := d
			locald.NodeID = n.ID
			_, matchesFilter := idFilter[d.ID]
			if len(idFilter) > 0 && !matchesFilter {
				// if there is a filter provided, and the device's ID is not in the filter
				// (i.e. return is len(idFilter)), then do not add to map
				continue
			}

			devs[d.ID] = &locald
		}
	}

	/*
				modifiedDevs, err := a.joinWithUISPData(devs)
				if err != nil {
					return nil, fmt.Errorf("error joining uisp data with devices: %w", err)
				}
		    return modifiedDevs, nil
	*/

	return devs, nil
}

// joinWithUISPData takes an unaugmented map of devices, and populates UISP fields
func (a *App) joinWithUISPData(devs map[int]*nycmesh.Device) (map[int]*nycmesh.Device, error) {
	interval := "day"
	sitesok, err := a.UISPAPI.Sites.GetSites(sites.NewGetSitesParams().WithDefaults(), nil)
	if err != nil {
		return nil, err
	}
	for _, v := range sitesok.GetPayload() {
		log.Printf("fetching site %s %s", *v.ID, v.Name)
		p := sites.NewGetSitesSiteidTrafficSummaryParams().
			WithSiteID(*v.ID).
			WithInterval(interval)
		siteTrafficSummary, err := a.UISPAPI.Sites.GetSitesSiteidTrafficSummary(p, nil)
		if err != nil {
			return nil, fmt.Errorf("error getting traffic summary for %s: %w", v.Name, err)
		}
		log.Printf("[traffic last %s] up: %s | down: %s", interval, ByteCountSI(int64(*siteTrafficSummary.Payload.Upload)), ByteCountSI(int64(*siteTrafficSummary.Payload.Download)))
	}

	ok, err := a.UISPAPI.Devices.GetDevices(devices.NewGetDevicesParams(), nil)
	if err != nil {
		log.Printf("%+v", ok)
		return nil, fmt.Errorf("unable to get devices: %w", err)
	}
	return nil, fmt.Errorf("implement joinWithUISPData")
}

func ByteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func ByteCountIEC(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}

// TODO(gabe): implement me
func (a *App) ListSectorsByFrequency() error {
	params := devices.NewGetDevicesParams()
	devs, err := a.UISPAPI.Devices.GetDevices(params, nil)
	if err != nil {
		return fmt.Errorf("unable to get devices: %w", err)
	}
	fmt.Printf("res: %+v\n", devs)
	for _, d := range devs.Payload {
		fmt.Printf("%+v\n", d)
	}
	return nil
}

// TODO(gabe): implement me
func (a *App) SetSectorFrequency(frequency string, devids ...[]string) error {
	log.Printf("setSectorFrequency called: freq:%s devs:%v\n", frequency, devids)

	params := devices.NewGetDevicesParams()
	devs, err := a.UISPAPI.Devices.GetDevices(params, nil)
	if err != nil {
		return err
	}
	fmt.Printf("res: %+v\n", devs)
	for _, d := range devs.Payload {
		fmt.Printf("%+v\n", d)
	}
	return fmt.Errorf("TODO implement me")
}
