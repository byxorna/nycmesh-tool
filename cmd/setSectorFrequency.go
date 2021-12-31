package cmd

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/pkg/app"
	"github.com/spf13/cobra"
)

type setSectorFrequencyOpts struct {
	apply     bool
	frequency string
	devices   []string
}

var (
	ssfo = setSectorFrequencyOpts{}
)

// setSectorFrequencyCmd represents the setSectorFrequency command
var setSectorFrequencyCmd = &cobra.Command{
	Use:   "setSectorFrequency",
	Short: "Update sector radio frequency",
	Long: `Allow operators to update the frequency of a sector,
either with a manually specified frequency, or an automated selection.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if ssfo.apply {
			return app.SetSectorFrequency(ssfo.frequency, ssfo.devices)
		} else {
			return fmt.Errorf("You must pass --apply to apply these changes")
		}
	},
}

func init() {
	rootCmd.AddCommand(setSectorFrequencyCmd)
	setSectorFrequencyCmd.PersistentFlags().BoolVar(&ssfo.apply, "apply", false, "Apply the change")
	setSectorFrequencyCmd.PersistentFlags().StringVar(&ssfo.frequency, "frequency", "auto", "The frequency to set")
	setSectorFrequencyCmd.PersistentFlags().StringArrayVar(&ssfo.devices, "devices", []string{}, "device identifiers to apply change to")
}
