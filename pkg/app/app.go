package app

import (
	"fmt"
	"github.com/byxorna/nycmesh-tool/cli"
	"github.com/byxorna/nycmesh-tool/client"
	"github.com/byxorna/nycmesh-tool/client/devices"
	"github.com/spf13/cobra"
	"log"
)

type App struct {
	*client.UISPAPI
}

func New(cmd *cobra.Command, args []string) (*App, error) {
	a := App{}
	c, err := cli.NewClient(cmd, args)
	if err != nil {
		return nil, err
	}
	a.UISPAPI = c
	return &a, nil
}

func (a *App) ListSectorsByFrequency() error {

	params := devices.NewGetDevicesParams()
	devs, err := a.Devices.GetDevices(params, nil)
	if err != nil {
		return err
	}
	fmt.Printf("res: %+v\n", devs)
	for _, d := range devs.Payload {
		fmt.Printf("%+v\n", d)
	}
	return nil
}
func (a *App) SetSectorFrequency(frequency string, devids ...[]string) error {
	log.Printf("setSectorFrequency called: freq:%s devs:%v\n", frequency, devids)

	params := devices.NewGetDevicesParams()
	devs, err := a.Devices.GetDevices(params, nil)
	if err != nil {
		return err
	}
	fmt.Printf("res: %+v\n", devs)
	for _, d := range devs.Payload {
		fmt.Printf("%+v\n", d)
	}
	return fmt.Errorf("TODO implement me")
}
