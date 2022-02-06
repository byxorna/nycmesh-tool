package cmd

import (
	"log"

	"github.com/byxorna/nycmesh-tool/pkg/app"
	"github.com/spf13/cobra"
)

var (
	daemonCmd = &cobra.Command{
		Use:   "daemon",
		Short: "launch a daemon",
		RunE: func(cmd *cobra.Command, args []string) error {
			a, err := app.New(cmd, args)
			if err != nil {
				return err
			}

			log.Print("launching daemon...")
			errs := a.RunDaemon()
			if len(errs) > 0 {
				log.Printf("daemon reported %d errors", len(errs))

				// for now, just return the first error to kill the program
				// as the errors shouild have already been logged as they were intercepted
				return errs[0]

			}
			return nil
		},
	}
)

func init() {
	// TODO: add pidfile, etc type flags?
	rootCmd.AddCommand(daemonCmd)
}
