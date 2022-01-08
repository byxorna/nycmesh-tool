package cmd

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	"github.com/byxorna/nycmesh-tool/pkg/app"
	"github.com/byxorna/nycmesh-tool/pkg/nycmesh"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	status         string
	StatusAny      = "any"
	StatusActive   = "active"
	StatusInactive = "inactive"
)

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

		idArgs := []int{}
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				return fmt.Errorf("%s is not a valid node id: %w", arg, err)
			}
			idArgs = append(idArgs, id)
		}

		nodes, err := a.Nodes(idArgs...)
		if err != nil {
			return err
		}

		headers := []string{"ID", "Geo", "Status", "Devices", "Building", "Notes"}
		data := make([][]string, len(nodes))

		nodeNumbers := make([]int, 0, len(nodes))
		for _, n := range nodes {
			nodeNumbers = append(nodeNumbers, int(n.ID))
		}
		sort.Ints(nodeNumbers)

		orderedNodes := []*nycmesh.Node{}
		for _, nn := range nodeNumbers {
			n := nodes[nn]

			if status == StatusAny || status == n.Status {
				orderedNodes = append(orderedNodes, &n)
				data = append(data, []string{
					fmt.Sprintf("%d", n.ID),
					n.GeoURI(),
					fmt.Sprintf("%s", n.Status),
					fmt.Sprintf("%d", len(n.Devices)),
					fmt.Sprintf("%s", n.Building),
					fmt.Sprintf("%s", n.Notes),
				})
			}
		}

		switch viper.GetString("core.format") {
		case "json":
			pp, err := json.MarshalIndent(orderedNodes, "", "\t")
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

func init() {
	nodeCmd.PersistentFlags().StringVar(&status, "status", StatusActive, "filter for status")
	nodeCmd.AddCommand(nodeListCmd)
	nodeCmd.AddCommand(nodeGetCmd)
	rootCmd.AddCommand(nodeCmd)
}
