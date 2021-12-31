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

// makeOperationDevicesPostDevicesDeviceidInterfacesInterfacenameBlockCmd returns a cmd to handle operation postDevicesDeviceidInterfacesInterfacenameBlock
func makeOperationDevicesPostDevicesDeviceidInterfacesInterfacenameBlockCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "postDevicesDeviceidInterfacesInterfacenameBlock",
		Short: ``,
		RunE:  runOperationDevicesPostDevicesDeviceidInterfacesInterfacenameBlock,
	}

	if err := registerOperationDevicesPostDevicesDeviceidInterfacesInterfacenameBlockParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDevicesPostDevicesDeviceidInterfacesInterfacenameBlock uses cmd flags to call endpoint api
func runOperationDevicesPostDevicesDeviceidInterfacesInterfacenameBlock(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := devices.NewPostDevicesDeviceidInterfacesInterfacenameBlockParams()
	if err, _ := retrieveOperationDevicesPostDevicesDeviceidInterfacesInterfacenameBlockDeviceIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesPostDevicesDeviceidInterfacesInterfacenameBlockInterfaceNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDevicesPostDevicesDeviceidInterfacesInterfacenameBlockResult(appCli.Devices.PostDevicesDeviceidInterfacesInterfacenameBlock(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDevicesPostDevicesDeviceidInterfacesInterfacenameBlockParamFlags registers all flags needed to fill params
func registerOperationDevicesPostDevicesDeviceidInterfacesInterfacenameBlockParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDevicesPostDevicesDeviceidInterfacesInterfacenameBlockDeviceIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesPostDevicesDeviceidInterfacesInterfacenameBlockInterfaceNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDevicesPostDevicesDeviceidInterfacesInterfacenameBlockDeviceIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationDevicesPostDevicesDeviceidInterfacesInterfacenameBlockInterfaceNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationDevicesPostDevicesDeviceidInterfacesInterfacenameBlockDeviceIDFlag(m *devices.PostDevicesDeviceidInterfacesInterfacenameBlockParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDevicesPostDevicesDeviceidInterfacesInterfacenameBlockInterfaceNameFlag(m *devices.PostDevicesDeviceidInterfacesInterfacenameBlockParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationDevicesPostDevicesDeviceidInterfacesInterfacenameBlockResult parses request result and return the string content
func parseOperationDevicesPostDevicesDeviceidInterfacesInterfacenameBlockResult(resp0 *devices.PostDevicesDeviceidInterfacesInterfacenameBlockOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*devices.PostDevicesDeviceidInterfacesInterfacenameBlockOK)
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
		resp1, ok := iResp1.(*devices.PostDevicesDeviceidInterfacesInterfacenameBlockBadRequest)
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
		resp2, ok := iResp2.(*devices.PostDevicesDeviceidInterfacesInterfacenameBlockUnauthorized)
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
		resp3, ok := iResp3.(*devices.PostDevicesDeviceidInterfacesInterfacenameBlockForbidden)
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
		resp4, ok := iResp4.(*devices.PostDevicesDeviceidInterfacesInterfacenameBlockNotFound)
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
		resp5, ok := iResp5.(*devices.PostDevicesDeviceidInterfacesInterfacenameBlockInternalServerError)
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
