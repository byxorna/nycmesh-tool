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

// makeOperationDevicesPutDevicesDeviceidInterfacesInterfacenameCmd returns a cmd to handle operation putDevicesDeviceidInterfacesInterfacename
func makeOperationDevicesPutDevicesDeviceidInterfacesInterfacenameCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "putDevicesDeviceidInterfacesInterfacename",
		Short: ``,
		RunE:  runOperationDevicesPutDevicesDeviceidInterfacesInterfacename,
	}

	if err := registerOperationDevicesPutDevicesDeviceidInterfacesInterfacenameParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDevicesPutDevicesDeviceidInterfacesInterfacename uses cmd flags to call endpoint api
func runOperationDevicesPutDevicesDeviceidInterfacesInterfacename(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := devices.NewPutDevicesDeviceidInterfacesInterfacenameParams()
	if err, _ := retrieveOperationDevicesPutDevicesDeviceidInterfacesInterfacenameBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesPutDevicesDeviceidInterfacesInterfacenameDeviceIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesPutDevicesDeviceidInterfacesInterfacenameInterfaceNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesPutDevicesDeviceidInterfacesInterfacenameIsForceApplyFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDevicesPutDevicesDeviceidInterfacesInterfacenameResult(appCli.Devices.PutDevicesDeviceidInterfacesInterfacename(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDevicesPutDevicesDeviceidInterfacesInterfacenameParamFlags registers all flags needed to fill params
func registerOperationDevicesPutDevicesDeviceidInterfacesInterfacenameParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDevicesPutDevicesDeviceidInterfacesInterfacenameBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesPutDevicesDeviceidInterfacesInterfacenameDeviceIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesPutDevicesDeviceidInterfacesInterfacenameInterfaceNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesPutDevicesDeviceidInterfacesInterfacenameIsForceApplyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDevicesPutDevicesDeviceidInterfacesInterfacenameBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelDeviceInterfaceSchemaFlags(0, "deviceInterfaceSchema", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationDevicesPutDevicesDeviceidInterfacesInterfacenameDeviceIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationDevicesPutDevicesDeviceidInterfacesInterfacenameInterfaceNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	interfaceNameDescription := `Required. `

	var interfaceNameFlagName string
	if cmdPrefix == "" {
		interfaceNameFlagName = "interfaceName"
	} else {
		interfaceNameFlagName = fmt.Sprintf("%v.interfaceName", cmdPrefix)
	}

	var interfaceNameFlagDefault string

	_ = cmd.PersistentFlags().String(interfaceNameFlagName, interfaceNameFlagDefault, interfaceNameDescription)

	return nil
}
func registerOperationDevicesPutDevicesDeviceidInterfacesInterfacenameIsForceApplyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	isForceApplyDescription := ``

	var isForceApplyFlagName string
	if cmdPrefix == "" {
		isForceApplyFlagName = "isForceApply"
	} else {
		isForceApplyFlagName = fmt.Sprintf("%v.isForceApply", cmdPrefix)
	}

	var isForceApplyFlagDefault string

	_ = cmd.PersistentFlags().String(isForceApplyFlagName, isForceApplyFlagDefault, isForceApplyDescription)

	return nil
}

func retrieveOperationDevicesPutDevicesDeviceidInterfacesInterfacenameBodyFlag(m *devices.PutDevicesDeviceidInterfacesInterfacenameParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.DeviceInterfaceSchema{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.DeviceInterfaceSchema: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.DeviceInterfaceSchema{}
	}
	err, added := retrieveModelDeviceInterfaceSchemaFlags(0, bodyValueModel, "deviceInterfaceSchema", cmd)
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
func retrieveOperationDevicesPutDevicesDeviceidInterfacesInterfacenameDeviceIDFlag(m *devices.PutDevicesDeviceidInterfacesInterfacenameParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDevicesPutDevicesDeviceidInterfacesInterfacenameInterfaceNameFlag(m *devices.PutDevicesDeviceidInterfacesInterfacenameParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("interfaceName") {

		var interfaceNameFlagName string
		if cmdPrefix == "" {
			interfaceNameFlagName = "interfaceName"
		} else {
			interfaceNameFlagName = fmt.Sprintf("%v.interfaceName", cmdPrefix)
		}

		interfaceNameFlagValue, err := cmd.Flags().GetString(interfaceNameFlagName)
		if err != nil {
			return err, false
		}
		m.InterfaceName = interfaceNameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationDevicesPutDevicesDeviceidInterfacesInterfacenameIsForceApplyFlag(m *devices.PutDevicesDeviceidInterfacesInterfacenameParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("isForceApply") {

		var isForceApplyFlagName string
		if cmdPrefix == "" {
			isForceApplyFlagName = "isForceApply"
		} else {
			isForceApplyFlagName = fmt.Sprintf("%v.isForceApply", cmdPrefix)
		}

		isForceApplyFlagValue, err := cmd.Flags().GetString(isForceApplyFlagName)
		if err != nil {
			return err, false
		}
		m.IsForceApply = &isForceApplyFlagValue

	}
	return nil, retAdded
}

// parseOperationDevicesPutDevicesDeviceidInterfacesInterfacenameResult parses request result and return the string content
func parseOperationDevicesPutDevicesDeviceidInterfacesInterfacenameResult(resp0 *devices.PutDevicesDeviceidInterfacesInterfacenameOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*devices.PutDevicesDeviceidInterfacesInterfacenameOK)
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
		resp1, ok := iResp1.(*devices.PutDevicesDeviceidInterfacesInterfacenameBadRequest)
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
		resp2, ok := iResp2.(*devices.PutDevicesDeviceidInterfacesInterfacenameUnauthorized)
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
		resp3, ok := iResp3.(*devices.PutDevicesDeviceidInterfacesInterfacenameForbidden)
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
		resp4, ok := iResp4.(*devices.PutDevicesDeviceidInterfacesInterfacenameNotFound)
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
		resp5, ok := iResp5.(*devices.PutDevicesDeviceidInterfacesInterfacenameInternalServerError)
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
