package cmd

import (
	"fmt"
	"github.com/byxorna/nycmesh-tool/pkg/app"
	"github.com/byxorna/nycmesh-tool/pkg/nycmesh"
	"github.com/spf13/cobra"
	"sort"
	"strconv"
)

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "nodes, hubs, and supernodes",
}

var nodeGetCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"get", "s"},
	Short:   "show a node",
	RunE: func(cmd *cobra.Command, args []string) error {
		a, err := app.New(cmd, args)
		if err != nil {
			return err
		}

		nodes, err := a.Nodes()
		if err != nil {
			return err
		}

		nodesToShow := []nycmesh.Node{}

		for _, a := range args {
			n, err := strconv.Atoi(a)
			if err != nil {
				return fmt.Errorf("%s is not a valid node identifier: %w", a, err)
			}

			node, ok := nodes[n]
			if !ok {
				return fmt.Errorf("node %d not found", n)
			}
			nodesToShow = append(nodesToShow, node)
		}

		headers := []string{"ID", "Lat", "Lon", "Altitude", "Status", "Devices", "Notes"}
		data := make([][]string, len(nodes))

		for i, n := range nodesToShow {
			data[i] = []string{
				fmt.Sprintf("%d", n.ID),
				fmt.Sprintf("%f", n.Latitude),
				fmt.Sprintf("%f", n.Longitude),
				fmt.Sprintf("%f", n.AltitudeMeters),
				fmt.Sprintf("%s", n.Status),
				fmt.Sprintf("%d", len(n.Devices)),
				fmt.Sprintf("%s", n.Notes),
			}
		}

		a.Tableify(headers, data)

		return nil
	},
}

var nodeListCmd = &cobra.Command{
	Use:   "list",
	Short: "list nodes",
	RunE: func(cmd *cobra.Command, args []string) error {
		a, err := app.New(cmd, args)
		if err != nil {
			return err
		}

		nodes, err := a.Nodes()
		if err != nil {
			return err
		}

		headers := []string{"ID", "Lat", "Lon", "Altitude", "Status", "Devices", "Notes"}
		data := make([][]string, len(nodes))

		nodeNumbers := make([]int, 0, len(nodes))
		for _, n := range nodes {
			nodeNumbers = append(nodeNumbers, int(n.ID))
		}
		sort.Ints(nodeNumbers)

		for _, nn := range nodeNumbers {
			n := nodes[nn]
			data = append(data, []string{
				fmt.Sprintf("%d", n.ID),
				fmt.Sprintf("%f", n.Latitude),
				fmt.Sprintf("%f", n.Longitude),
				fmt.Sprintf("%f", n.AltitudeMeters),
				fmt.Sprintf("%s", n.Status),
				fmt.Sprintf("%d", len(n.Devices)),
				fmt.Sprintf("%s", n.Notes),
			})
		}

		a.Tableify(headers, data)
		return nil
	},
}

func init() {
	nodeCmd.AddCommand(nodeListCmd)
	nodeCmd.AddCommand(nodeGetCmd)
	rootCmd.AddCommand(nodeCmd)
}
