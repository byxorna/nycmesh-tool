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

// makeOperationDevicesGetDevicesEroutersIDNetflowCmd returns a cmd to handle operation getDevicesEroutersIdNetflow
func makeOperationDevicesGetDevicesEroutersIDNetflowCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getDevicesEroutersIdNetflow",
		Short: ``,
		RunE:  runOperationDevicesGetDevicesEroutersIDNetflow,
	}

	if err := registerOperationDevicesGetDevicesEroutersIDNetflowParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDevicesGetDevicesEroutersIDNetflow uses cmd flags to call endpoint api
func runOperationDevicesGetDevicesEroutersIDNetflow(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := devices.NewGetDevicesEroutersIDNetflowParams()
	if err, _ := retrieveOperationDevicesGetDevicesEroutersIDNetflowIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesGetDevicesEroutersIDNetflowInterfaceIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDevicesGetDevicesEroutersIDNetflowResult(appCli.Devices.GetDevicesEroutersIDNetflow(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDevicesGetDevicesEroutersIDNetflowParamFlags registers all flags needed to fill params
func registerOperationDevicesGetDevicesEroutersIDNetflowParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDevicesGetDevicesEroutersIDNetflowIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesGetDevicesEroutersIDNetflowInterfaceIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDevicesGetDevicesEroutersIDNetflowIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationDevicesGetDevicesEroutersIDNetflowInterfaceIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	interfaceIdDescription := ``

	var interfaceIdFlagName string
	if cmdPrefix == "" {
		interfaceIdFlagName = "interfaceId"
	} else {
		interfaceIdFlagName = fmt.Sprintf("%v.interfaceId", cmdPrefix)
	}

	var interfaceIdFlagDefault string

	_ = cmd.PersistentFlags().String(interfaceIdFlagName, interfaceIdFlagDefault, interfaceIdDescription)

	return nil
}

func retrieveOperationDevicesGetDevicesEroutersIDNetflowIDFlag(m *devices.GetDevicesEroutersIDNetflowParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDevicesGetDevicesEroutersIDNetflowInterfaceIDFlag(m *devices.GetDevicesEroutersIDNetflowParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("interfaceId") {

		var interfaceIdFlagName string
		if cmdPrefix == "" {
			interfaceIdFlagName = "interfaceId"
		} else {
			interfaceIdFlagName = fmt.Sprintf("%v.interfaceId", cmdPrefix)
		}

		interfaceIdFlagValue, err := cmd.Flags().GetString(interfaceIdFlagName)
		if err != nil {
			return err, false
		}
		m.InterfaceID = &interfaceIdFlagValue

	}
	return nil, retAdded
}

// parseOperationDevicesGetDevicesEroutersIDNetflowResult parses request result and return the string content
func parseOperationDevicesGetDevicesEroutersIDNetflowResult(resp0 *devices.GetDevicesEroutersIDNetflowOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*devices.GetDevicesEroutersIDNetflowOK)
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
		resp1, ok := iResp1.(*devices.GetDevicesEroutersIDNetflowBadRequest)
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
		resp2, ok := iResp2.(*devices.GetDevicesEroutersIDNetflowUnauthorized)
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
		resp3, ok := iResp3.(*devices.GetDevicesEroutersIDNetflowNotFound)
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
		resp4, ok := iResp4.(*devices.GetDevicesEroutersIDNetflowInternalServerError)
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