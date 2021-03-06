// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/server"
	serverops "github.com/byxorna/nycmesh-tool/generated/go/uisp/client/server"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationServerGetNmsVersionLatestCmd returns a cmd to handle operation getNmsVersionLatest
func makeOperationServerGetNmsVersionLatestCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getNmsVersionLatest",
		Short: ``,
		RunE:  runOperationServerGetNmsVersionLatest,
	}

	if err := registerOperationServerGetNmsVersionLatestParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServerGetNmsVersionLatest uses cmd flags to call endpoint api
func runOperationServerGetNmsVersionLatest(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := serverops.NewGetNmsVersionLatestParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationServerGetNmsVersionLatestResult(appCli.Server.GetNmsVersionLatest(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationServerGetNmsVersionLatestParamFlags registers all flags needed to fill params
func registerOperationServerGetNmsVersionLatestParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationServerGetNmsVersionLatestResult parses request result and return the string content
func parseOperationServerGetNmsVersionLatestResult(resp0 *server.GetNmsVersionLatestOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*server.GetNmsVersionLatestOK)
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
		resp1, ok := iResp1.(*server.GetNmsVersionLatestForbidden)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*server.GetNmsVersionLatestInternalServerError)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr := fmt.Sprintf("%v", resp0.Payload)
		return string(msgStr), nil
	}

	return "", nil
}
