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

// makeOperationDevicesPutDevicesDeviceidBackupsCmd returns a cmd to handle operation putDevicesDeviceidBackups
func makeOperationDevicesPutDevicesDeviceidBackupsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "putDevicesDeviceidBackups",
		Short: ``,
		RunE:  runOperationDevicesPutDevicesDeviceidBackups,
	}

	if err := registerOperationDevicesPutDevicesDeviceidBackupsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDevicesPutDevicesDeviceidBackups uses cmd flags to call endpoint api
func runOperationDevicesPutDevicesDeviceidBackups(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := devices.NewPutDevicesDeviceidBackupsParams()
	if err, _ := retrieveOperationDevicesPutDevicesDeviceidBackupsBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesPutDevicesDeviceidBackupsDeviceIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDevicesPutDevicesDeviceidBackupsResult(appCli.Devices.PutDevicesDeviceidBackups(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDevicesPutDevicesDeviceidBackupsParamFlags registers all flags needed to fill params
func registerOperationDevicesPutDevicesDeviceidBackupsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDevicesPutDevicesDeviceidBackupsBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesPutDevicesDeviceidBackupsDeviceIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDevicesPutDevicesDeviceidBackupsBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelModel114Flags(0, "model114", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationDevicesPutDevicesDeviceidBackupsDeviceIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	deviceIdDescription := `Required. `

	var deviceIdFlagName string
	if cmdPrefix == "" {
		deviceIdFlagName = "deviceId"
	} else {
		deviceIdFlagName = fmt.Sprintf("%v.deviceId", cmdPrefix)
	}

	var deviceIdFlagDefault string

	_ = cmd.PersistentFlags().String(deviceIdFlagName, deviceIdFlagDefault, deviceIdDescription)

	return nil
}

func retrieveOperationDevicesPutDevicesDeviceidBackupsBodyFlag(m *devices.PutDevicesDeviceidBackupsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.Model114{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.Model114: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.Model114{}
	}
	err, added := retrieveModelModel114Flags(0, bodyValueModel, "model114", cmd)
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
func retrieveOperationDevicesPutDevicesDeviceidBackupsDeviceIDFlag(m *devices.PutDevicesDeviceidBackupsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("deviceId") {

		var deviceIdFlagName string
		if cmdPrefix == "" {
			deviceIdFlagName = "deviceId"
		} else {
			deviceIdFlagName = fmt.Sprintf("%v.deviceId", cmdPrefix)
		}

		deviceIdFlagValue, err := cmd.Flags().GetString(deviceIdFlagName)
		if err != nil {
			return err, false
		}
		m.DeviceID = deviceIdFlagValue

	}
	return nil, retAdded
}

// parseOperationDevicesPutDevicesDeviceidBackupsResult parses request result and return the string content
func parseOperationDevicesPutDevicesDeviceidBackupsResult(resp0 *devices.PutDevicesDeviceidBackupsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*devices.PutDevicesDeviceidBackupsOK)
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
		resp1, ok := iResp1.(*devices.PutDevicesDeviceidBackupsBadRequest)
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
		resp2, ok := iResp2.(*devices.PutDevicesDeviceidBackupsUnauthorized)
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
		resp3, ok := iResp3.(*devices.PutDevicesDeviceidBackupsForbidden)
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
		resp4, ok := iResp4.(*devices.PutDevicesDeviceidBackupsNotFound)
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
		resp5, ok := iResp5.(*devices.PutDevicesDeviceidBackupsInternalServerError)
		if ok {
			if !swag.IsZero(resp5) && !swag.IsZero(resp5.Payload) {
				msgStr, err := json.Marshal(resp5.Payload)
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
