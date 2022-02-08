package cmd

import (
	"context"
	"log"

	"github.com/byxorna/nycmesh-tool/pkg/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
			errs := a.RunDaemon(context.Background())
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
	daemonCmd.Flags().Bool("dfs-event-detection", true, "enable DFS event detection feature")
	viper.BindPFlag("daemon.dfs-event-detection", daemonCmd.Flags().Lookup("dfs-event-detection"))

	daemonCmd.Flags().Bool("enable-slack", false, "enable Slack integration for the daemon")
	viper.BindPFlag("daemon.enable-slack", daemonCmd.Flags().Lookup("enable-slack"))

	daemonCmd.Flags().String("slack-token", "slack.token", "specify the slack.token to connect with slack")
	viper.BindPFlag("slack.token", daemonCmd.Flags().Lookup("slack-token"))
	rootCmd.AddCommand(daemonCmd)
}
