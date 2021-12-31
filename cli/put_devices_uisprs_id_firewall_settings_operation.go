// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/client/devices"
	"github.com/byxorna/nycmesh-tool/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationDevicesPutDevicesUisprsIDFirewallSettingsCmd returns a cmd to handle operation putDevicesUisprsIdFirewallSettings
func makeOperationDevicesPutDevicesUisprsIDFirewallSettingsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "putDevicesUisprsIdFirewallSettings",
		Short: ``,
		RunE:  runOperationDevicesPutDevicesUisprsIDFirewallSettings,
	}

	if err := registerOperationDevicesPutDevicesUisprsIDFirewallSettingsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDevicesPutDevicesUisprsIDFirewallSettings uses cmd flags to call endpoint api
func runOperationDevicesPutDevicesUisprsIDFirewallSettings(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := devices.NewPutDevicesUisprsIDFirewallSettingsParams()
	if err, _ := retrieveOperationDevicesPutDevicesUisprsIDFirewallSettingsBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesPutDevicesUisprsIDFirewallSettingsIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDevicesPutDevicesUisprsIDFirewallSettingsResult(appCli.Devices.PutDevicesUisprsIDFirewallSettings(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDevicesPutDevicesUisprsIDFirewallSettingsParamFlags registers all flags needed to fill params
func registerOperationDevicesPutDevicesUisprsIDFirewallSettingsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDevicesPutDevicesUisprsIDFirewallSettingsBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesPutDevicesUisprsIDFirewallSettingsIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDevicesPutDevicesUisprsIDFirewallSettingsBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelModel56Flags(0, "model56", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationDevicesPutDevicesUisprsIDFirewallSettingsIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationDevicesPutDevicesUisprsIDFirewallSettingsBodyFlag(m *devices.PutDevicesUisprsIDFirewallSettingsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.Model56{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.Model56: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.Model56{}
	}
	err, added := retrieveModelModel56Flags(0, bodyValueModel, "model56", cmd)
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
func retrieveOperationDevicesPutDevicesUisprsIDFirewallSettingsIDFlag(m *devices.PutDevicesUisprsIDFirewallSettingsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationDevicesPutDevicesUisprsIDFirewallSettingsResult parses request result and return the string content
func parseOperationDevicesPutDevicesUisprsIDFirewallSettingsResult(resp0 *devices.PutDevicesUisprsIDFirewallSettingsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*devices.PutDevicesUisprsIDFirewallSettingsOK)
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
		resp1, ok := iResp1.(*devices.PutDevicesUisprsIDFirewallSettingsBadRequest)
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
		resp2, ok := iResp2.(*devices.PutDevicesUisprsIDFirewallSettingsUnauthorized)
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
		resp3, ok := iResp3.(*devices.PutDevicesUisprsIDFirewallSettingsNotFound)
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
		resp4, ok := iResp4.(*devices.PutDevicesUisprsIDFirewallSettingsInternalServerError)
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
