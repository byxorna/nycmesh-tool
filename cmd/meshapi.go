package cmd

import (
	"github.com/spf13/cobra"
)

var meshapiCmd = &cobra.Command{
	Use:     "meshapi",
	Short:   "interact with the mesh-api",
	Aliases: []string{"mesh", "mesh-api"},
}

func init() {
	rootCmd.AddCommand(meshapiCmd)
}
