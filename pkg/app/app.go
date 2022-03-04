package app

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/cli"
	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client"
	"github.com/byxorna/nycmesh-tool/pkg/cache"
	"github.com/byxorna/nycmesh-tool/pkg/meshapi"
	"github.com/byxorna/nycmesh-tool/pkg/version"
	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
)

type App struct {
	*client.UISPAPI

	Slack *slack.Client

	MeshAPIClient *meshapi.Client

	// TODO: wire this up as a general caching layer, instead of if being in meshapi.Client?
	diskCache cache.DiskCache

	daemonBootupTimestamp time.Time
	config                *Config
}

func New(cmd *cobra.Command, args []string) (*App, error) {
	a := App{}

	cfg, err := NewConfig()
	if err != nil {
		return nil, err
	}
	a.config = cfg

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

	if a.config.Daemon.EnableSlack {
		a.Slack = slack.New(a.config.Slack.token)
	}
	return &a, nil
}

func VersionString() string {
	return fmt.Sprintf("%s (commit:%s branch:%s built:%s)", version.Release, version.GitCommit, version.GitBranch, version.BuildDate)
}

func (a *App) MeshAPIKML() ([]byte, error) {
	return a.MeshAPIClient.KML()
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

	///

	var body []byte
	var err error
	if !a.diskCache.Has("nodes") {
		// get nodes from API
		nodes, err := a.MeshAPIClient.Nodes()
		// write ndoes into cache if its happy
		if err != nil {
			return nil, fmt.Errorf("error getting mesh-api nodes: %w", err)
		}

		// TODO: this could be improved with a proper read-thru cache, but I don't
		// want to implement it :P
		// Regardless, this cache deserialization business could definitely be
		// DRYed up somewhat...

		if x, err := json.Marshal(&nodes); err != nil {
			return nil, fmt.Errorf("unable to marshal nodes for caching: %w", err)
		} else {
			err = a.diskCache.Write("nodes", x)
			if err != nil {
				return nil, err
			}
		}
	}

	body, err = a.diskCache.Read("nodes")
	if err != nil {
		return nil, fmt.Errorf("cache error: %w", err)
	}

	// TODO: here too, we could DRY this unmarshalling. having to deserialize structs outside
	// the meshapi package is pretty rough...
	var nodes []meshapi.Node
	if err = json.Unmarshal(body, &nodes); err != nil {
		return nil, fmt.Errorf("unable to decode nodes: %w\n%s", err, body)
	}

	idFilter := map[int]bool{}
	for _, id := range idInts {
		idFilter[id] = true
	}

	m := map[int]meshapi.Node{}
	for _, node := range nodes {
		if _, matchesFilter := idFilter[node.ID]; !matchesFilter && len(idFilter) > 0 {
			// only return nodes that match provided filter if necessary
			continue
		}
		m[node.ID] = node
	}

	return m, nil
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
