package app

import (
	"fmt"
	"log"
	"strconv"

	"github.com/byxorna/nycmesh-tool/client/devices"
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

	return devs, nil
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
