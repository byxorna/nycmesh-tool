package cmd

import (
	"log"

	"github.com/byxorna/nycmesh-tool/pkg/cache"
	"github.com/spf13/cobra"
)

// cacheCmd represents the cache command
var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "manage cached files for this tool",
}

var cacheClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clear cache",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("clearing cache")

		diskCache, err := cache.NewDiskVCache()
		if err != nil {
			return err
		}
		err = diskCache.EraseAll()
		if err != nil {
			return err
		}
		return nil

	},
}

func init() {
	cacheCmd.AddCommand(cacheClearCmd)
	rootCmd.AddCommand(cacheCmd)
}
