package app

import "github.com/spf13/viper"

type Config struct {
	slackToken   string
	outputFormat string
	DaemonConfig
}

type DaemonConfig struct {
	DFSEventDetection bool
	EnableSlack       bool
}

func NewConfig() *Config {
	return &Config{
		outputFormat: viper.GetString("core.format"),
		slackToken:   viper.GetString("slack.token"),
		DaemonConfig: DaemonConfig{
			DFSEventDetection: viper.GetBool("daemon.dfs-event-detection"),
			EnableSlack:       viper.GetBool("daemon.enable-slack"),
		},
	}
}
