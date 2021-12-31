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

// makeOperationServerPostNmsTrafficBlacklistCmd returns a cmd to handle operation postNmsTrafficBlacklist
func makeOperationServerPostNmsTrafficBlacklistCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "postNmsTrafficBlacklist",
		Short: ``,
		RunE:  runOperationServerPostNmsTrafficBlacklist,
	}

	if err := registerOperationServerPostNmsTrafficBlacklistParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServerPostNmsTrafficBlacklist uses cmd flags to call endpoint api
func runOperationServerPostNmsTrafficBlacklist(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := serverops.NewPostNmsTrafficBlacklistParams()
	if err, _ := retrieveOperationServerPostNmsTrafficBlacklistBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationServerPostNmsTrafficBlacklistResult(appCli.Server.PostNmsTrafficBlacklist(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationServerPostNmsTrafficBlacklistParamFlags registers all flags needed to fill params
func registerOperationServerPostNmsTrafficBlacklistParamFlags(cmd *cobra.Command) error {
	if err := registerOperationServerPostNmsTrafficBlacklistBodyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationServerPostNmsTrafficBlacklistBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	// warning: array body models.BlacklistSchema is not supported by go-swagger cli yet

	return nil
}

func retrieveOperationServerPostNmsTrafficBlacklistBodyFlag(m *serverops.PostNmsTrafficBlacklistParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {

		// warning: array body models.BlacklistSchema is not supported by go-swagger cli yet

	}
	return nil, retAdded
}

// parseOperationServerPostNmsTrafficBlacklistResult parses request result and return the string content
func parseOperationServerPostNmsTrafficBlacklistResult(resp0 *server.PostNmsTrafficBlacklistOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*server.PostNmsTrafficBlacklistOK)
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
		resp1, ok := iResp1.(*server.PostNmsTrafficBlacklistUnauthorized)
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
		resp2, ok := iResp2.(*server.PostNmsTrafficBlacklistForbidden)
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
		resp3, ok := iResp3.(*server.PostNmsTrafficBlacklistInternalServerError)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
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