// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/client/authorization"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationAuthorizationGetUserTotpauthCmd returns a cmd to handle operation getUserTotpauth
func makeOperationAuthorizationGetUserTotpauthCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getUserTotpauth",
		Short: ``,
		RunE:  runOperationAuthorizationGetUserTotpauth,
	}

	if err := registerOperationAuthorizationGetUserTotpauthParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationAuthorizationGetUserTotpauth uses cmd flags to call endpoint api
func runOperationAuthorizationGetUserTotpauth(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := authorization.NewGetUserTotpauthParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationAuthorizationGetUserTotpauthResult(appCli.Authorization.GetUserTotpauth(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationAuthorizationGetUserTotpauthParamFlags registers all flags needed to fill params
func registerOperationAuthorizationGetUserTotpauthParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationAuthorizationGetUserTotpauthResult parses request result and return the string content
func parseOperationAuthorizationGetUserTotpauthResult(resp0 *authorization.GetUserTotpauthOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*authorization.GetUserTotpauthOK)
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
		resp1, ok := iResp1.(*authorization.GetUserTotpauthUnauthorized)
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
		resp2, ok := iResp2.(*authorization.GetUserTotpauthForbidden)
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
		resp3, ok := iResp3.(*authorization.GetUserTotpauthInternalServerError)
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
