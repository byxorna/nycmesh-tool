package cmd

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	"github.com/byxorna/nycmesh-tool/pkg/app"
	"github.com/spf13/cobra"
)

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:     "node",
	Aliases: []string{"n", "nodes"},
	Short:   "nodes, hubs, and supernodes",
}

var nodeGetCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"get", "s", "g"},
	Short:   "show a node",
	RunE: func(cmd *cobra.Command, args []string) error {
		a, err := app.New(cmd, args)
		if err != nil {
			return err
		}

		if len(args) != 1 {
			return fmt.Errorf("one device identifier argument is required")
		}

		nodes, err := a.Nodes()
		if err != nil {
			return err
		}

		for _, a := range args {
			n, err := strconv.Atoi(a)
			if err != nil {
				return fmt.Errorf("%s is not a valid node identifier: %w", a, err)
			}

			node, ok := nodes[n]
			if !ok {
				return fmt.Errorf("node %d not found", n)
			}
			pp, err := json.MarshalIndent(node, "", "\t")
			if err != nil {
				return err
			}
			fmt.Printf("%s\n", pp)
		}

		return nil
	},
}

var nodeListCmd = &cobra.Command{
	Use:     "list",
	Short:   "list nodes",
	Aliases: []string{"find", "l", "f"},
	RunE: func(cmd *cobra.Command, args []string) error {
		a, err := app.New(cmd, args)
		if err != nil {
			return err
		}

		nodes, err := a.Nodes()
		if err != nil {
			return err
		}

		headers := []string{"ID", "Geo", "Status", "Devices", "Notes"}
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
				n.GeoURI(),
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
