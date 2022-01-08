package cmd

import (
	"github.com/byxorna/nycmesh-tool/generated/go/uisp/cli"
)

func init() {
	uispCmd, err := cli.MakeRootCmdCustom("uisp", "UISP API operations")
	if err != nil {
		panic(err)
	}
	rootCmd.AddCommand(uispCmd)
}
