package cmd

import (
	"fmt"
	"time"

	"github.com/byxorna/nycmesh-tool/pkg/app"
	"github.com/spf13/cobra"
)

var (
	watchCmd = &cobra.Command{
		Use:   "watch",
		Short: "keep an eye on events happening in the mesh",
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("....")
		},
	}

	watchLogsCmd = &cobra.Command{
		Use:   "logs",
		Short: "watch logs from UISP",
		RunE: func(cmd *cobra.Command, args []string) error {
			a, err := app.New(cmd, args)
			if err != nil {
				return err
			}

			logCh := make(chan app.LogEvent, 10)

			go func() {
				// as log events are produced by the watcher, print em out in a nice format
				for le := range logCh {
					fmt.Printf("%s\t%s\t%s\n", le.Time.Local().Format(time.RFC3339), *le.Level, *le.Message)
				}
			}()

			if err := a.WatchLogs(cmd.Context(), logCh); err != nil {
				return fmt.Errorf("log watch failed: %w", err)
			}
			return nil
		},
	}

	watchOutagesCmd = &cobra.Command{
		Use:   "outages",
		Short: "watch outages from UISP",
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("implement watch outages")
		},
	}
)

func init() {
	watchCmd.AddCommand(watchLogsCmd)
	watchCmd.AddCommand(watchOutagesCmd)
	rootCmd.AddCommand(watchCmd)
}
