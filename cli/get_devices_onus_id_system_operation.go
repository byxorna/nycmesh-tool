// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/client/devices"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationDevicesGetDevicesOnusIDSystemCmd returns a cmd to handle operation getDevicesOnusIdSystem
func makeOperationDevicesGetDevicesOnusIDSystemCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getDevicesOnusIdSystem",
		Short: ``,
		RunE:  runOperationDevicesGetDevicesOnusIDSystem,
	}

	if err := registerOperationDevicesGetDevicesOnusIDSystemParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDevicesGetDevicesOnusIDSystem uses cmd flags to call endpoint api
func runOperationDevicesGetDevicesOnusIDSystem(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := devices.NewGetDevicesOnusIDSystemParams()
	if err, _ := retrieveOperationDevicesGetDevicesOnusIDSystemIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDevicesGetDevicesOnusIDSystemResult(appCli.Devices.GetDevicesOnusIDSystem(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDevicesGetDevicesOnusIDSystemParamFlags registers all flags needed to fill params
func registerOperationDevicesGetDevicesOnusIDSystemParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDevicesGetDevicesOnusIDSystemIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDevicesGetDevicesOnusIDSystemIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. `

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func retrieveOperationDevicesGetDevicesOnusIDSystemIDFlag(m *devices.GetDevicesOnusIDSystemParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("id") {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

	}
	return nil, retAdded
}

// parseOperationDevicesGetDevicesOnusIDSystemResult parses request result and return the string content
func parseOperationDevicesGetDevicesOnusIDSystemResult(resp0 *devices.GetDevicesOnusIDSystemOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*devices.GetDevicesOnusIDSystemOK)
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
		resp1, ok := iResp1.(*devices.GetDevicesOnusIDSystemBadRequest)
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
		resp2, ok := iResp2.(*devices.GetDevicesOnusIDSystemUnauthorized)
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
		resp3, ok := iResp3.(*devices.GetDevicesOnusIDSystemForbidden)
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
		resp4, ok := iResp4.(*devices.GetDevicesOnusIDSystemNotFound)
		if ok {
			if !swag.IsZero(resp4) && !swag.IsZero(resp4.Payload) {
				msgStr, err := json.Marshal(resp4.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp5 interface{} = respErr
		resp5, ok := iResp5.(*devices.GetDevicesOnusIDSystemInternalServerError)
		if ok {
			if !swag.IsZero(resp5) && !swag.IsZero(resp5.Payload) {
				msgStr, err := json.Marshal(resp5.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp6 interface{} = respErr
		resp6, ok := iResp6.(*devices.GetDevicesOnusIDSystemNotImplemented)
		if ok {
			if !swag.IsZero(resp6) && !swag.IsZero(resp6.Payload) {
				msgStr, err := json.Marshal(resp6.Payload)
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