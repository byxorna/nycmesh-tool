package app


import (
	"fmt"
  "github.com/byxorna/nycmesh-tool/pkg/nycmesh"
	"github.com/byxorna/nycmesh-tool/client/devices"
)

func (a *App) Devices(ids ...string) ([]*nycmesh.Device, error) {
  return nil, fmt.Errorf("implement me Devices()")
}

// TODO(gabe): implement me
func (a *App) ListSectorsByFrequency() error {
	params := devices.NewGetDevicesParams()
	devs, err := a.UISPAPI.Devices.GetDevices(params, nil)
	if err != nil {
		return err
	}
	fmt.Printf("res: %+v\n", devs)
	for _, d := range devs.Payload {
		fmt.Printf("%+v\n", d)
	}
	return nil
}


