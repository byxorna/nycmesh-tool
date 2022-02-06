package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/byxorna/nycmesh-tool/pkg/app"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	watchSince   time.Duration
	watchOnlyNNs []int

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
				errColor := color.New(color.FgHiRed).SprintFunc()
				infoColor := color.New(color.FgHiGreen).SprintFunc()
				warnColor := color.New(color.FgHiYellow).SprintFunc()
				hiWhiteColor := color.New(color.FgHiWhite).SprintFunc()
				levelColor := func(level string, in string) string {
					switch level {
					case "info":
						return fmt.Sprintf("%s", infoColor(in))
					case "warning":
						return fmt.Sprintf("%s", warnColor(in))
					case "error":
						return fmt.Sprintf("%s", errColor(in))
					default:
						return in
					}
				}
				// as log events are produced by the watcher, print em out in a nice format
				for le := range logCh {
					var nnstr string
					if le.NN > 0 {
						nnstr = fmt.Sprintf("nn:%s", hiWhiteColor(fmt.Sprintf("%d", le.NN)))
					}

					filterMatch := true
					if len(watchOnlyNNs) > 0 {
						filterMatch = false
						for _, nn := range watchOnlyNNs {
							filterMatch = (filterMatch || nn == le.NN)
						}
					}

					if filterMatch {
						fmt.Printf("%s\t%s\t%s\t%s\n", le.Time.Local().Format(time.RFC3339), levelColor(*le.Level, *le.Level), nnstr, *le.Message)
					}
				}
			}()

			log.Printf("watching logs since %s ago", watchSince)
			if err := a.WatchLogs(cmd.Context(), time.Now().Add(-(watchSince)), logCh); err != nil {
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
	watchLogsCmd.Flags().DurationVar(&watchSince, "since", time.Minute*15, "Begin watching for events this far in the past")
	watchLogsCmd.Flags().IntSliceVar(&watchOnlyNNs, "nn", []int{}, "filter logs for only those concerning devices/sites for these NNs (WARNING: this filter is performed clientside)")

	watchCmd.AddCommand(watchLogsCmd)
	watchCmd.AddCommand(watchOutagesCmd)

	rootCmd.AddCommand(watchCmd)
}
