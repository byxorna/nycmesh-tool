package cmd

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/byxorna/nycmesh-tool/pkg/selfupdate"
	"github.com/spf13/cobra"
)

var (
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Self-update commands",
	}

	updateCheckCmd = &cobra.Command{
		Use:   "check",
		Short: "Check for tool updates by querying github releases",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Print("checking for updates")
			_ = checkForUpdate()
			return nil
		},
	}
)

func init() {
	updateCmd.AddCommand(updateCheckCmd)
	rootCmd.AddCommand(updateCmd)
}

func checkForUpdate() error {
	c, err := selfupdate.NewUpdaterGithub()
	if err != nil {
		return fmt.Errorf("unable to create selfupdater: %w", err)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	releaseAssetName := fmt.Sprintf("%s-%s-%s", rootCmd.Use, runtime.GOOS, runtime.GOARCH)
	if runtime.GOOS == "windows" {
		releaseAssetName = fmt.Sprintf("%s.exe", releaseAssetName)
	}
	r, err := c.HasNewerRelease(ctx, releaseAssetName)

	switch err {
	case selfupdate.ErrConnectionError:
		log.Printf("update check failed: %s", err.Error())
	case selfupdate.ErrNoNewRelease:
		log.Printf("up to date! (%s)", r.TagName)
	case nil:
		if r.DownloadURL != "" {
			binName := fmt.Sprintf(`$HOME/bin/%s`, releaseAssetName)
			log.Printf(`🎉 %s is available! Download it with 'curl -o %s "%s" && chmod +x %s'`, binName, r.TagName, r.DownloadURL, binName)
		} else {
			// probably, we dont have builds available for this releaseAssetName
			log.Printf("🥳 %s is available, but I didnt find a release asset named '%s'. See %s for available builds", releaseAssetName, r.TagName, r.URL)
		}
	default:
		log.Printf("update error: %s", err.Error())
	}
	return err
}
