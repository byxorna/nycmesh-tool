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

var expDeviceCmd = &cobra.Command{
	Use:     "device",
	Short:   "interact with devices",
	Aliases: []string{"d", "devices"},
}

var expDeviceListCmd = &cobra.Command{
	Use:     "list",
	Short:   "list devices",
	Aliases: []string{"find", "l", "f"},
	RunE: func(cmd *cobra.Command, args []string) error {
		a, err := app.New(cmd, args)
		if err != nil {
			return err
		}

		// --- load necessary data ---

		// fetch all mesh-api devices
		log.Printf("loading mesh-api devices...")
		meshDevs, err := a.MeshAPIDevices()
		if err != nil {
			return fmt.Errorf("error fetching mesh-api devices: %w", err)
		}

		// fetch all mesh-api nodes
		log.Printf("loading mesh-api nodes...")
		meshNodes, err := a.MeshAPINodes()
		if err != nil {
			return fmt.Errorf("error fetching mesh-api nodes: %w", err)
		}

		// fetch all UISP devices
		log.Printf("loading UISP devices...")
		uispDevs, err := a.UISPDevices()
		if err != nil {
			return fmt.Errorf("error fetching UISP devices: %w", err)
		}

		// --- field extraction for display ---

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
					// ignore node number extraction errors, continue on
				} else {
					dev.NodeNumber = nn
					nodeNumberStr = fmt.Sprintf("%d", nn)
					if node, ok := meshNodes[nn]; ok {
						dev.MeshAPINode = &node

						// now, we try to _best effort_ identify which device in mesh-api matches
						// which UISP device. This is super haxx right now, because the best key
						// we have to join the structs together is whether or not the SSID field
						// in both structs is populated.
						// NOTE: previously I was thinking about using mac address to perform this
						// join, but discovered that SNMP/black box devices in UISP do not have
						// mac addresses
						// If we dont know what device in mesh-api a UISP device matches with, we
						// leave output field `"meshapi_device": null` in the FusedDevice struct
						for _, d := range node.Devices {
							locald := d
							if dev.UISP.Attributes != nil && dev.UISP.Attributes.Ssid != "" && d.SSID == dev.UISP.Attributes.Ssid {
								log.Printf("matched UISP device %s to mesh-api device %d by ssid=%s", dev.UISP.Identification.Name, locald.ID, locald.SSID)
								dev.MeshAPIDevice = &locald
								break
							}
						}

					}
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

		// --- render output ---

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
	expDeviceCmd.AddCommand(expDeviceListCmd)
	experimentCmd.AddCommand(expDeviceCmd)

	expDeviceCmd.PersistentFlags().StringVar(&status, "status", StatusActive, "filter for status")
	expDeviceListCmd.Flags().IntVar(&filterNodeID, "node", 0, "only show devices for the given node ID")
}
