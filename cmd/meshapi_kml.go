package cmd

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/pkg/app"
	"github.com/spf13/cobra"
)

var kmlCmd = &cobra.Command{
	Use:   "kml",
	Short: "Fetch the mesh-api's KML of the network",
	RunE: func(cmd *cobra.Command, args []string) error {
		a, err := app.New(cmd, args)
		if err != nil {
			return err
		}

		kml, err := a.MeshAPIKML()
		if err != nil {
			return fmt.Errorf("unable to fetch KML: %w", err)
		}

		fmt.Printf("%s\n", kml)

		return nil
	},
}

func init() {
	meshapiCmd.AddCommand(kmlCmd)
}
