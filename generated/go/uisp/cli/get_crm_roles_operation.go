// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/c_r_m"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationCrmGetCrmRolesCmd returns a cmd to handle operation getCrmRoles
func makeOperationCrmGetCrmRolesCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getCrmRoles",
		Short: ``,
		RunE:  runOperationCrmGetCrmRoles,
	}

	if err := registerOperationCrmGetCrmRolesParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationCrmGetCrmRoles uses cmd flags to call endpoint api
func runOperationCrmGetCrmRoles(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := c_r_m.NewGetCrmRolesParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationCrmGetCrmRolesResult(appCli.Crm.GetCrmRoles(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationCrmGetCrmRolesParamFlags registers all flags needed to fill params
func registerOperationCrmGetCrmRolesParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationCrmGetCrmRolesResult parses request result and return the string content
func parseOperationCrmGetCrmRolesResult(resp0 *c_r_m.GetCrmRolesOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*c_r_m.GetCrmRolesOK)
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
		resp1, ok := iResp1.(*c_r_m.GetCrmRolesUnauthorized)
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
		resp2, ok := iResp2.(*c_r_m.GetCrmRolesForbidden)
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
		resp3, ok := iResp3.(*c_r_m.GetCrmRolesInternalServerError)
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
		resp4, ok := iResp4.(*c_r_m.GetCrmRolesServiceUnavailable)
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
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
