// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/client/server"
	serverops "github.com/byxorna/nycmesh-tool/client/server"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationServerGetNmsServerconfigCmd returns a cmd to handle operation getNmsServerconfig
func makeOperationServerGetNmsServerconfigCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getNmsServerconfig",
		Short: ``,
		RunE:  runOperationServerGetNmsServerconfig,
	}

	if err := registerOperationServerGetNmsServerconfigParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServerGetNmsServerconfig uses cmd flags to call endpoint api
func runOperationServerGetNmsServerconfig(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := serverops.NewGetNmsServerconfigParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationServerGetNmsServerconfigResult(appCli.Server.GetNmsServerconfig(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationServerGetNmsServerconfigParamFlags registers all flags needed to fill params
func registerOperationServerGetNmsServerconfigParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationServerGetNmsServerconfigResult parses request result and return the string content
func parseOperationServerGetNmsServerconfigResult(resp0 *server.GetNmsServerconfigOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*server.GetNmsServerconfigOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*server.GetNmsServerconfigInternalServerError)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
