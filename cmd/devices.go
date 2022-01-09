package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/byxorna/nycmesh-tool/pkg/app"
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

		// fetch all UISP devices
		uispDevs, err := a.UISPDevices()
		if err != nil {
			return fmt.Errorf("error fetching UISP devices: %w", err)
		}
		headers := []string{"Name", "NN", "MAC", "Model", "Type", "Site", "IP", "Firmware", "Frequency", "ChWidth", "LinkScore%"}
		data := make([][]string, len(meshDevs))

		orderedDevs := []*app.FusedDevice{}

		{
			for _, dso := range uispDevs {
				localdso := dso
				dev := app.FusedDevice{
					UISP: localdso,
				}

				// fields we explicitly extract for table view
				var name, nodeNumberStr, devType, mac, model, site, ip, fw, freq, chwidth, linkScore string

				nn, err := getNNFromUISPDevice(localdso)
				if err != nil {
					//return fmt.Errorf("err nn: %s: %w", localdso.Identification.Name, err)
					// ignore node number extraction errors, continue on
				} else {
					dev.NodeNumber = nn
					nodeNumberStr = fmt.Sprintf("%d", nn)
				}

				if meshDev, ok := meshDevs[dev.NodeNumber]; ok {
					log.Printf("found meshapi dev nn:%d from uisp device name %s", dev.NodeNumber, localdso.Identification.Name)
					dev.MeshAPI = meshDev
				}

				orderedDevs = append(orderedDevs, &dev)

				if dso.Identification != nil {
					name = dso.Identification.Name
					devType = dso.Identification.Type
					mac = dso.Identification.Mac
					model = dso.Identification.Model
					fw = dso.Identification.FirmwareVersion
					if dso.Identification.Site != nil {
						site = dso.Identification.Site.Name
					}
					if dso.IPAddress != nil {
						ip = *dso.IPAddress
					}

					if dso.Overview != nil {
						freq = fmt.Sprintf("%.0f", dso.Overview.Frequency)
						chwidth = fmt.Sprintf("%.0f", dso.Overview.ChannelWidth)
						if dso.Overview.LinkScore != nil && dso.Overview.LinkScore.Score != nil {
							linkScore = fmt.Sprintf("%.0f", float64(100.0)*(*dso.Overview.LinkScore.Score))
						}
					}
				}

				data = append(data, []string{name, nodeNumberStr, mac, model, devType, site, ip, fw, freq, chwidth, linkScore})
			}
		}

		switch viper.GetString("core.format") {
		case "json":
			pp, err := json.MarshalIndent(orderedDevs, "", "\t")
			if err != nil {
				return err
			}
			fmt.Printf("%s\n", pp)
		case "table":
			a.Tableify(headers, data)
		default:
			return fmt.Errorf("unsupported format: %s", viper.GetString("core.format"))
		}

		return nil
	},
}

func getNNFromUISPDevice(d *models.DeviceStatusOverview) (int, error) {
	extracter := regexp.MustCompile(`nycmesh-(?:[^\d]+-)?(\d{3,6})\b`)
	var n string
	if d.Identification != nil {
		n = d.Identification.Name
	}
	if extracter.MatchString(n) {
		res := extracter.FindStringSubmatch(n)
		var match string
		switch len(res) {
		case 2:
			match = res[1]
		case 3:
			match = res[2]
		}
		nn, err := strconv.Atoi(match)
		if err != nil {
			return 0, fmt.Errorf("unable to parse %s to nn %s to int: %w", n, match, err)
		}
		return nn, nil
	}
	return 0, fmt.Errorf("unable to derive nn from %s", n)
}

func init() {
	deviceCmd.AddCommand(deviceListCmd)
	deviceCmd.AddCommand(deviceGetCmd)
	rootCmd.AddCommand(deviceCmd)

	deviceCmd.PersistentFlags().StringVar(&status, "status", StatusActive, "filter for status")
	deviceListCmd.Flags().IntVar(&filterNodeID, "node", 0, "only show devices for the given node ID")
}
