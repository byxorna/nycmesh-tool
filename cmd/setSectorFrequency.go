package cmd

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/pkg/app"
	"github.com/spf13/cobra"
)

// TODO: remove this in favor of viper
type setSectorFrequencyOpts struct {
	apply     bool
	frequency string
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
		app, err := app.New(cmd, args)
		if err != nil {
			return err
		}
		if ssfo.apply {
			return app.SetSectorFrequency(ssfo.frequency, args)
		} else {
			app.ListSectorsByFrequency()
			return fmt.Errorf("You must pass --apply to apply these changes")
		}
	},
}

func init() {
	rootCmd.AddCommand(setSectorFrequencyCmd)
	setSectorFrequencyCmd.PersistentFlags().BoolVar(&ssfo.apply, "apply", false, "Apply the change")
	setSectorFrequencyCmd.PersistentFlags().StringVar(&ssfo.frequency, "frequency", "auto", "The frequency to set")

}
