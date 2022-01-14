package cmd

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	"github.com/byxorna/nycmesh-tool/pkg/app"
	"github.com/byxorna/nycmesh-tool/pkg/meshapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var filterNodeID int

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
		deviceID, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("%s is not a valid device identifier: %w", args[0], err)
		}

		devs, err := a.MeshAPIDevices(args...)
		if err != nil {
			return fmt.Errorf("error getting device %d: %w", deviceID, err)
		}

		d, ok := devs[deviceID]
		if !ok {
			return fmt.Errorf("device %d not found", deviceID)
		}

		pp, err := json.MarshalIndent(d, "", "\t")
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

		// fetch all mesh-api devices
		meshDevs, err := a.MeshAPIDevices(args...)
		if err != nil {
			return fmt.Errorf("error fetching mesh-api devices: %w", err)
		}

		headers := []string{"NN", "Name", "ID", "SSID", "Mfg", "Model", "Geo", "Status", "Notes"}
		data := make([][]string, len(meshDevs))

		allDevList := []*meshapi.Device{}
		for _, d := range meshDevs {
			locald := d
			allDevList = append(allDevList, locald)
		}
		sort.Slice(allDevList, func(i, j int) bool {
			return allDevList[i].Name < allDevList[j].Name
		})

		switch viper.GetString("core.format") {
		case "json":
			pp, err := json.MarshalIndent(allDevList, "", "\t")
			if err != nil {
				return err
			}
			fmt.Printf("%s\n", pp)
		case "table":
			for _, d := range allDevList {
				// populate table columns
				data = append(data, []string{fmt.Sprintf("%d", d.NodeID), d.Name, fmt.Sprintf("%d", d.ID), d.SSID, d.Type.Manufacturer, d.Type.Name, meshapi.GeoURI(d.Latitude, d.Longitude, d.AltitudeMeters), d.Status, d.Notes})
			}
			a.Tableify(headers, data)
		default:
			return fmt.Errorf("unsupported format: %s", viper.GetString("core.format"))
		}

		return nil
	},
}

func init() {
	deviceCmd.AddCommand(deviceListCmd)
	deviceCmd.AddCommand(deviceGetCmd)
	meshapiCmd.AddCommand(deviceCmd)

	deviceCmd.PersistentFlags().StringVar(&status, "status", StatusActive, "filter for status")
	deviceListCmd.Flags().IntVar(&filterNodeID, "node", 0, "only show devices for the given node ID")
}
