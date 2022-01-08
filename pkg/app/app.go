package app

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/cli"
	"github.com/byxorna/nycmesh-tool/client"
	"github.com/byxorna/nycmesh-tool/pkg/cache"
	"github.com/byxorna/nycmesh-tool/pkg/nycmesh"
	"github.com/spf13/cobra"
)

type App struct {
	*client.UISPAPI
	*nycmesh.Client

	// TODO: wire this up as a general caching layer, instead of if being in nycmesh.Client?
	diskCache cache.DiskCache
}

func New(cmd *cobra.Command, args []string) (*App, error) {
	a := App{}

	diskCache, err := cache.NewDiskVCache()
	if err != nil {
		return nil, fmt.Errorf("unable to create new disk cache: %w", err)
	}
	a.diskCache = diskCache

	c, err := cli.NewClient(cmd, args)
	if err != nil {
		return nil, fmt.Errorf("unable to create cli: %w", err)
	}
	a.UISPAPI = c

	nycmeshClient, err := nycmesh.New()
	if err != nil {
		return nil, fmt.Errorf("unable to create nycmesh client: %w", err)
	}
	a.Client = nycmeshClient
	return &a, nil
}
