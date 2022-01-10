package cmd

import (
	"github.com/spf13/cobra"
)

var experimentCmd = &cobra.Command{
	Use:     "experiment",
	Short:   "interact with experimental commands",
	Aliases: []string{"exp", "experimental", "ex"},
}

func init() {
	rootCmd.AddCommand(experimentCmd)
}
