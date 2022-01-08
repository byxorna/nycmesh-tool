package cmd

import (
	"github.com/byxorna/nycmesh-tool/generated/go/uisp/cli"
	"github.com/go-openapi/strfmt"
)

func init() {
	uispCmd, _ := cli.MakeRootCmd()
	uispCmd.Use = "uisp"
	uispCmd.Short = "UISP API operations"
	rootCmd.AddCommand(uispCmd)

}
