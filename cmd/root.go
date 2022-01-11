package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/byxorna/nycmesh-tool/pkg/app"
	"github.com/byxorna/nycmesh-tool/pkg/selfupdate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var (
	rootCmd = &cobra.Command{
		Use:   "nycmesh-tool",
		Short: "NYC Mesh tool provides CLI access to UISP, mesh-api, and more",
		Long: `NYC Mesh tool provides an intuitive CLI that wraps the UISP API,
    mesh-api (https://github.com/meshcenter/mesh-api), as well as various
    business logic to operate mesh networks at scale`,
		Version: app.VersionString(),
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	// TODO: dont exhaust github rate limit by checking each invocation
	_ = checkForUpdate()

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func checkForUpdate() error {
	if c, err := selfupdate.NewUpdaterGithub(); err == nil {
		log.Printf("checking for update...")
		ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancelFunc()

		releaseAssetName := fmt.Sprintf("%s-%s-%s", rootCmd.Use, runtime.GOOS, runtime.GOARCH)
		if runtime.GOOS == "windows" {
			releaseAssetName = fmt.Sprintf("%s.exe", releaseAssetName)
		}
		r, err := c.HasNewerRelease(ctx, releaseAssetName)

		switch err {
		case selfupdate.ErrConnectionError:
			log.Print("connection error, skipping version check")
		case selfupdate.ErrNoNewRelease:
			log.Print("no new release available")
		case nil:
			if r.DownloadURL != "" {
				binName := fmt.Sprintf(`$HOME/bin/%s`, releaseAssetName)
				log.Printf(`🎉 %s is available! Download it with 'curl -o %s "%s" && chmod +x %s'`, binName, r.TagName, r.DownloadURL, binName)
			} else {
				// probably, we dont have builds available for this releaseAssetName
				log.Printf("🥳 %s is available, but I didnt find a release asset named '%s'. See %s for available builds", releaseAssetName, r.TagName, r.URL)
			}
		default:
			log.Print(err.Error())
		}
	}
	return nil
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("format", "f", "json", "output format (table or json)")
	viper.BindPFlag("core.format", rootCmd.PersistentFlags().Lookup("format"))
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nycmesh-tool.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".nycmesh-tool")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func GetIntSlice(i *[]string) ([]int, error) {
	var arr = *i
	ret := []int{}
	for _, str := range arr {
		one_int, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("%s is not a valid ID format: %w", str, err)
		}
		ret = append(ret, one_int)
	}
	return ret, nil
}
