// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/devices"
	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationDevicesPutDevicesBlackboxesIDConfigCmd returns a cmd to handle operation putDevicesBlackboxesIdConfig
func makeOperationDevicesPutDevicesBlackboxesIDConfigCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "putDevicesBlackboxesIdConfig",
		Short: ``,
		RunE:  runOperationDevicesPutDevicesBlackboxesIDConfig,
	}

	if err := registerOperationDevicesPutDevicesBlackboxesIDConfigParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDevicesPutDevicesBlackboxesIDConfig uses cmd flags to call endpoint api
func runOperationDevicesPutDevicesBlackboxesIDConfig(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := devices.NewPutDevicesBlackboxesIDConfigParams()
	if err, _ := retrieveOperationDevicesPutDevicesBlackboxesIDConfigBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesPutDevicesBlackboxesIDConfigIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDevicesPutDevicesBlackboxesIDConfigResult(appCli.Devices.PutDevicesBlackboxesIDConfig(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDevicesPutDevicesBlackboxesIDConfigParamFlags registers all flags needed to fill params
func registerOperationDevicesPutDevicesBlackboxesIDConfigParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDevicesPutDevicesBlackboxesIDConfigBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesPutDevicesBlackboxesIDConfigIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDevicesPutDevicesBlackboxesIDConfigBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelDeviceBlackBoxConfigFlags(0, "deviceBlackBoxConfig", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationDevicesPutDevicesBlackboxesIDConfigIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationDevicesPutDevicesBlackboxesIDConfigBodyFlag(m *devices.PutDevicesBlackboxesIDConfigParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.DeviceBlackBoxConfig{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.DeviceBlackBoxConfig: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.DeviceBlackBoxConfig{}
	}
	err, added := retrieveModelDeviceBlackBoxConfigFlags(0, bodyValueModel, "deviceBlackBoxConfig", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Body = bodyValueModel
	}
	if dryRun && debug {

		bodyValueDebugBytes, err := json.Marshal(m.Body)
		if err != nil {
			return err, false
		}
		logDebugf("Body dry-run payload: %v", string(bodyValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationDevicesPutDevicesBlackboxesIDConfigIDFlag(m *devices.PutDevicesBlackboxesIDConfigParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationDevicesPutDevicesBlackboxesIDConfigResult parses request result and return the string content
func parseOperationDevicesPutDevicesBlackboxesIDConfigResult(resp0 *devices.PutDevicesBlackboxesIDConfigOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*devices.PutDevicesBlackboxesIDConfigOK)
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
		resp1, ok := iResp1.(*devices.PutDevicesBlackboxesIDConfigBadRequest)
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
		resp2, ok := iResp2.(*devices.PutDevicesBlackboxesIDConfigUnauthorized)
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
		resp3, ok := iResp3.(*devices.PutDevicesBlackboxesIDConfigForbidden)
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
		resp4, ok := iResp4.(*devices.PutDevicesBlackboxesIDConfigNotFound)
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
		resp5, ok := iResp5.(*devices.PutDevicesBlackboxesIDConfigConflict)
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
		resp6, ok := iResp6.(*devices.PutDevicesBlackboxesIDConfigInternalServerError)
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