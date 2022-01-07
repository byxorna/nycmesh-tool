package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/pkg/app"
	"github.com/spf13/cobra"
)

var deviceCmd = &cobra.Command{
	Use:     "device",
	Short:   "interact with devices",
	Aliases: []string{"d", "devices"},
}

var deviceGetCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"get", "s", "g"},
	Short:   "show a device",
	RunE: func(cmd *cobra.Command, args []string) error {
		a, err := app.New(cmd, args)
		if err != nil {
			return err
		}

		if len(args) != 1 || args[0] == "" {
			return fmt.Errorf("one device identifier argument is required")
		}

		deviceID := args[0]

		devices, err := a.Devices(deviceID)
		if err != nil {
			return err
		}

		if len(devices) == 0 {
			return fmt.Errorf("device %s not found", deviceID)
		}

    pp, err := json.MarshalIndent(devices[0], "", "\t")
    if err != nil {
      return err
    }
    fmt.Printf("%s\n", pp)

		return nil
	},
}

var deviceListCmd = &cobra.Command{
	Use:     "list",
	Short:   "list devices",
	Aliases: []string{"find", "l", "f"},
	RunE: func(cmd *cobra.Command, args []string) error {
		a, err := app.New(cmd, args)
		if err != nil {
			return err
		}

		devices, err := a.Devices(args...)
		if err != nil {
			return err
		}

		headers := []string{"ID", "Name", "Status" }
		data := make([][]string, len(devices))

		for _, dev := range devices {
			data = append(data, []string{
				fmt.Sprintf("%d", dev.ID),
				fmt.Sprintf("%s", dev.Name),
				fmt.Sprintf("%s", dev.Status),
			})
		}

		a.Tableify(headers, data)
		return nil
	},
}

func init() {
	deviceCmd.AddCommand(deviceListCmd)
	deviceCmd.AddCommand(deviceGetCmd)
	rootCmd.AddCommand(deviceCmd)
}
