package app

import (
	"fmt"
	"log"
	"strconv"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/devices"
	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/sites"
	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/byxorna/nycmesh-tool/pkg/meshapi"
)

type FusedDevice struct {
	NodeNumber    int                          `json:"nn"`
	MeshAPINode   *meshapi.Node                `json:"meshapi_node"`
	MeshAPIDevice *meshapi.Device              `json:"meshapi_device"`
	UISP          *models.DeviceStatusOverview `json:"uisp_device"`
}

func (a *App) UISPDevices() ([]*models.DeviceStatusOverview, error) {
	ok, err := a.UISPAPI.Devices.GetDevices(devices.NewGetDevicesParams(), nil)
	if err != nil {
		return nil, err
	}

	return ok.Payload, nil
}

func (a *App) MeshAPIDevices(ids ...string) (map[int]*meshapi.Device, error) {
	idFilter := map[int]interface{}{}
	for _, idarg := range ids {
		n, err := strconv.Atoi(idarg)
		if err != nil {
			return nil, fmt.Errorf("%s is not a valid device identifier: %w", idarg, err)
		}
		idFilter[n] = true
	}

	nodes, err := a.MeshAPINodes()
	if err != nil {
		return nil, err
	}

	devs := map[int]*meshapi.Device{}
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

	return devs, nil
}

// joinWithUISPData takes an unaugmented map of devices, and populates UISP fields
func (a *App) joinWithUISPData(devs map[int]*meshapi.Device) (map[int]*meshapi.Device, error) {
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
