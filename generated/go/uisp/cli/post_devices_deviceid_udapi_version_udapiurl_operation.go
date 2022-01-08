// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/devices"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationDevicesPostDevicesDeviceidUdapiVersionUdapiurlCmd returns a cmd to handle operation postDevicesDeviceidUdapiVersionUdapiurl
func makeOperationDevicesPostDevicesDeviceidUdapiVersionUdapiurlCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "postDevicesDeviceidUdapiVersionUdapiurl",
		Short: ``,
		RunE:  runOperationDevicesPostDevicesDeviceidUdapiVersionUdapiurl,
	}

	if err := registerOperationDevicesPostDevicesDeviceidUdapiVersionUdapiurlParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDevicesPostDevicesDeviceidUdapiVersionUdapiurl uses cmd flags to call endpoint api
func runOperationDevicesPostDevicesDeviceidUdapiVersionUdapiurl(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := devices.NewPostDevicesDeviceidUdapiVersionUdapiurlParams()
	if err, _ := retrieveOperationDevicesPostDevicesDeviceidUdapiVersionUdapiurlDeviceIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesPostDevicesDeviceidUdapiVersionUdapiurlUdapiURLFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesPostDevicesDeviceidUdapiVersionUdapiurlVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDevicesPostDevicesDeviceidUdapiVersionUdapiurlResult(appCli.Devices.PostDevicesDeviceidUdapiVersionUdapiurl(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDevicesPostDevicesDeviceidUdapiVersionUdapiurlParamFlags registers all flags needed to fill params
func registerOperationDevicesPostDevicesDeviceidUdapiVersionUdapiurlParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDevicesPostDevicesDeviceidUdapiVersionUdapiurlDeviceIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesPostDevicesDeviceidUdapiVersionUdapiurlUdapiURLParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesPostDevicesDeviceidUdapiVersionUdapiurlVersionParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDevicesPostDevicesDeviceidUdapiVersionUdapiurlDeviceIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationDevicesPostDevicesDeviceidUdapiVersionUdapiurlUdapiURLParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	udapiUrlDescription := `Required. `

	var udapiUrlFlagName string
	if cmdPrefix == "" {
		udapiUrlFlagName = "udapiUrl"
	} else {
		udapiUrlFlagName = fmt.Sprintf("%v.udapiUrl", cmdPrefix)
	}

	var udapiUrlFlagDefault string

	_ = cmd.PersistentFlags().String(udapiUrlFlagName, udapiUrlFlagDefault, udapiUrlDescription)

	return nil
}
func registerOperationDevicesPostDevicesDeviceidUdapiVersionUdapiurlVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	versionDescription := `Required. `

	var versionFlagName string
	if cmdPrefix == "" {
		versionFlagName = "version"
	} else {
		versionFlagName = fmt.Sprintf("%v.version", cmdPrefix)
	}

	var versionFlagDefault string

	_ = cmd.PersistentFlags().String(versionFlagName, versionFlagDefault, versionDescription)

	return nil
}

func retrieveOperationDevicesPostDevicesDeviceidUdapiVersionUdapiurlDeviceIDFlag(m *devices.PostDevicesDeviceidUdapiVersionUdapiurlParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDevicesPostDevicesDeviceidUdapiVersionUdapiurlUdapiURLFlag(m *devices.PostDevicesDeviceidUdapiVersionUdapiurlParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("udapiUrl") {

		var udapiUrlFlagName string
		if cmdPrefix == "" {
			udapiUrlFlagName = "udapiUrl"
		} else {
			udapiUrlFlagName = fmt.Sprintf("%v.udapiUrl", cmdPrefix)
		}

		udapiUrlFlagValue, err := cmd.Flags().GetString(udapiUrlFlagName)
		if err != nil {
			return err, false
		}
		m.UdapiURL = udapiUrlFlagValue

	}
	return nil, retAdded
}
func retrieveOperationDevicesPostDevicesDeviceidUdapiVersionUdapiurlVersionFlag(m *devices.PostDevicesDeviceidUdapiVersionUdapiurlParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("version") {

		var versionFlagName string
		if cmdPrefix == "" {
			versionFlagName = "version"
		} else {
			versionFlagName = fmt.Sprintf("%v.version", cmdPrefix)
		}

		versionFlagValue, err := cmd.Flags().GetString(versionFlagName)
		if err != nil {
			return err, false
		}
		m.Version = versionFlagValue

	}
	return nil, retAdded
}

// parseOperationDevicesPostDevicesDeviceidUdapiVersionUdapiurlResult parses request result and return the string content
func parseOperationDevicesPostDevicesDeviceidUdapiVersionUdapiurlResult(resp0 *devices.PostDevicesDeviceidUdapiVersionUdapiurlOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*devices.PostDevicesDeviceidUdapiVersionUdapiurlOK)
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
		resp1, ok := iResp1.(*devices.PostDevicesDeviceidUdapiVersionUdapiurlBadRequest)
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
		resp2, ok := iResp2.(*devices.PostDevicesDeviceidUdapiVersionUdapiurlUnauthorized)
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
		resp3, ok := iResp3.(*devices.PostDevicesDeviceidUdapiVersionUdapiurlForbidden)
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
		resp4, ok := iResp4.(*devices.PostDevicesDeviceidUdapiVersionUdapiurlNotFound)
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
		resp5, ok := iResp5.(*devices.PostDevicesDeviceidUdapiVersionUdapiurlInternalServerError)
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
		resp6, ok := iResp6.(*devices.PostDevicesDeviceidUdapiVersionUdapiurlNotImplemented)
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
		msgStr := fmt.Sprintf("%v", resp0.Payload)
		return string(msgStr), nil
	}

	return "", nil
}