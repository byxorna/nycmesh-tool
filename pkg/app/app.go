package app

import (
	"fmt"
	"strconv"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/cli"
	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client"
	"github.com/byxorna/nycmesh-tool/pkg/cache"
	"github.com/byxorna/nycmesh-tool/pkg/meshapi"
	"github.com/byxorna/nycmesh-tool/pkg/version"
	"github.com/spf13/cobra"
)

type App struct {
	*client.UISPAPI
	MeshAPIClient *meshapi.Client

	// TODO: wire this up as a general caching layer, instead of if being in meshapi.Client?
	diskCache cache.DiskCache
}

func New(cmd *cobra.Command, args []string) (*App, error) {
	a := App{}

	diskCache, err := cache.NewDiskVCache()
	if err != nil {
		return nil, fmt.Errorf("unable to create new disk cache: %w", err)
	}
	a.diskCache = diskCache

	c, err := cli.NewClientCustom(cmd, args)
	if err != nil {
		return nil, fmt.Errorf("unable to create cli: %w", err)
	}
	a.UISPAPI = c

	nycmeshClient, err := meshapi.New()
	if err != nil {
		return nil, fmt.Errorf("unable to create nycmesh client: %w", err)
	}
	a.MeshAPIClient = nycmeshClient
	return &a, nil
}

func VersionString() string {
	return fmt.Sprintf("%s (commit:%s branch:%s built:%s)", version.Release, version.GitCommit, version.GitBranch, version.BuildDate)
}

func (a *App) MeshAPINodes(ids ...string) (map[int]meshapi.Node, error) {
	var idInts []int
	if len(ids) > 0 {
		idFilter, err := getIDFilter(ids)
		if err != nil {
			return nil, err
		}
		for i := range idFilter {
			idInts = append(idInts, i)
		}
	}
	return a.MeshAPIClient.Nodes(idInts...)
}

func getIDFilter(args []string) (map[int]interface{}, error) {
	idFilter := map[int]interface{}{}
	for _, idarg := range args {
		n, err := strconv.Atoi(idarg)
		if err != nil {
			return nil, fmt.Errorf("%s is not a valid mesh-api ID: %w", idarg, err)
		}
		idFilter[n] = true
	}
	return idFilter, nil
}
