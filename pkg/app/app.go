package app

import (
	"github.com/byxorna/nycmesh-tool/cli"
	"github.com/byxorna/nycmesh-tool/client"
	"github.com/byxorna/nycmesh-tool/pkg/nycmesh"
	"github.com/spf13/cobra"
)

type App struct {
	*client.UISPAPI
	*nycmesh.Client
}

func New(cmd *cobra.Command, args []string) (*App, error) {
	a := App{}
	c, err := cli.NewClient(cmd, args)
	if err != nil {
		return nil, err
	}
	a.UISPAPI = c

	nycmeshClient, err := nycmesh.New()
	if err != nil {
		return nil, err
	}
	a.Client = nycmeshClient
	return &a, nil
}
