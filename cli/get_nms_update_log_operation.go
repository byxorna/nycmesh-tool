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

// makeOperationServerGetNmsUpdateLogCmd returns a cmd to handle operation getNmsUpdateLog
func makeOperationServerGetNmsUpdateLogCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getNmsUpdateLog",
		Short: ``,
		RunE:  runOperationServerGetNmsUpdateLog,
	}

	if err := registerOperationServerGetNmsUpdateLogParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServerGetNmsUpdateLog uses cmd flags to call endpoint api
func runOperationServerGetNmsUpdateLog(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := serverops.NewGetNmsUpdateLogParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationServerGetNmsUpdateLogResult(appCli.Server.GetNmsUpdateLog(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationServerGetNmsUpdateLogParamFlags registers all flags needed to fill params
func registerOperationServerGetNmsUpdateLogParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationServerGetNmsUpdateLogResult parses request result and return the string content
func parseOperationServerGetNmsUpdateLogResult(resp0 *server.GetNmsUpdateLogOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*server.GetNmsUpdateLogOK)
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
		resp1, ok := iResp1.(*server.GetNmsUpdateLogUnauthorized)
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
		resp2, ok := iResp2.(*server.GetNmsUpdateLogForbidden)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp3 interface{} = respErr
		resp3, ok := iResp3.(*server.GetNmsUpdateLogInternalServerError)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp4 interface{} = respErr
		resp4, ok := iResp4.(*server.GetNmsUpdateLogServiceUnavailable)
		if ok {
			if !swag.IsZero(resp4) && !swag.IsZero(resp4.Payload) {
				msgStr, err := json.Marshal(resp4.Payload)
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
