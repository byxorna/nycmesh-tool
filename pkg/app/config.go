package app

import (
	"github.com/spf13/viper"
)

type Config struct {
	Core   CoreConfig   `json:"core"`
	Slack  SlackConfig  `json:"slack"`
	Daemon DaemonConfig `json:"daemon"`
}

type SlackConfig struct {
	token string `json:"token"`
}

type CoreConfig struct {
	OutputFormat string `json:"format"`
}

type DaemonConfig struct {
	DFSEventDetection bool `json:"dfs-event-detection"`
	EnableSlack       bool `json:"enable-slack"`
}

func NewConfig() (*Config, error) {
	// TODO: there is surely a more clever way to get this struct populated via
	// viper.Unmarshal, but I cannot be troubled to figure out the proper struct tags
	// or why yaml bugs like https://github.com/spf13/viper/issues/338 matter
	cfg := &Config{
		Core: CoreConfig{
			OutputFormat: viper.GetString("core.format"),
		},
		Slack: SlackConfig{
			token: viper.GetString("slack.token"),
		},
		Daemon: DaemonConfig{
			DFSEventDetection: viper.GetBool("daemon.dfs-event-detection"),
			EnableSlack:       viper.GetBool("daemon.enable-slack"),
		},
	}

	return cfg, nil
}
