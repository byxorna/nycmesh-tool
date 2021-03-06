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

// makeOperationDevicesGetDevicesDeviceidBackupsBackupidCmd returns a cmd to handle operation getDevicesDeviceidBackupsBackupid
func makeOperationDevicesGetDevicesDeviceidBackupsBackupidCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getDevicesDeviceidBackupsBackupid",
		Short: ``,
		RunE:  runOperationDevicesGetDevicesDeviceidBackupsBackupid,
	}

	if err := registerOperationDevicesGetDevicesDeviceidBackupsBackupidParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDevicesGetDevicesDeviceidBackupsBackupid uses cmd flags to call endpoint api
func runOperationDevicesGetDevicesDeviceidBackupsBackupid(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := devices.NewGetDevicesDeviceidBackupsBackupidParams()
	if err, _ := retrieveOperationDevicesGetDevicesDeviceidBackupsBackupidBackupIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesGetDevicesDeviceidBackupsBackupidDeviceIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesGetDevicesDeviceidBackupsBackupidReplaceUnmsKeyFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDevicesGetDevicesDeviceidBackupsBackupidResult(appCli.Devices.GetDevicesDeviceidBackupsBackupid(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDevicesGetDevicesDeviceidBackupsBackupidParamFlags registers all flags needed to fill params
func registerOperationDevicesGetDevicesDeviceidBackupsBackupidParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDevicesGetDevicesDeviceidBackupsBackupidBackupIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesGetDevicesDeviceidBackupsBackupidDeviceIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesGetDevicesDeviceidBackupsBackupidReplaceUnmsKeyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDevicesGetDevicesDeviceidBackupsBackupidBackupIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	backupIdDescription := `Required. `

	var backupIdFlagName string
	if cmdPrefix == "" {
		backupIdFlagName = "backupId"
	} else {
		backupIdFlagName = fmt.Sprintf("%v.backupId", cmdPrefix)
	}

	var backupIdFlagDefault string

	_ = cmd.PersistentFlags().String(backupIdFlagName, backupIdFlagDefault, backupIdDescription)

	return nil
}
func registerOperationDevicesGetDevicesDeviceidBackupsBackupidDeviceIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationDevicesGetDevicesDeviceidBackupsBackupidReplaceUnmsKeyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	replaceUnmsKeyDescription := ``

	var replaceUnmsKeyFlagName string
	if cmdPrefix == "" {
		replaceUnmsKeyFlagName = "replaceUnmsKey"
	} else {
		replaceUnmsKeyFlagName = fmt.Sprintf("%v.replaceUnmsKey", cmdPrefix)
	}

	var replaceUnmsKeyFlagDefault bool

	_ = cmd.PersistentFlags().Bool(replaceUnmsKeyFlagName, replaceUnmsKeyFlagDefault, replaceUnmsKeyDescription)

	return nil
}

func retrieveOperationDevicesGetDevicesDeviceidBackupsBackupidBackupIDFlag(m *devices.GetDevicesDeviceidBackupsBackupidParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("backupId") {

		var backupIdFlagName string
		if cmdPrefix == "" {
			backupIdFlagName = "backupId"
		} else {
			backupIdFlagName = fmt.Sprintf("%v.backupId", cmdPrefix)
		}

		backupIdFlagValue, err := cmd.Flags().GetString(backupIdFlagName)
		if err != nil {
			return err, false
		}
		m.BackupID = backupIdFlagValue

	}
	return nil, retAdded
}
func retrieveOperationDevicesGetDevicesDeviceidBackupsBackupidDeviceIDFlag(m *devices.GetDevicesDeviceidBackupsBackupidParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDevicesGetDevicesDeviceidBackupsBackupidReplaceUnmsKeyFlag(m *devices.GetDevicesDeviceidBackupsBackupidParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("replaceUnmsKey") {

		var replaceUnmsKeyFlagName string
		if cmdPrefix == "" {
			replaceUnmsKeyFlagName = "replaceUnmsKey"
		} else {
			replaceUnmsKeyFlagName = fmt.Sprintf("%v.replaceUnmsKey", cmdPrefix)
		}

		replaceUnmsKeyFlagValue, err := cmd.Flags().GetBool(replaceUnmsKeyFlagName)
		if err != nil {
			return err, false
		}
		m.ReplaceUnmsKey = &replaceUnmsKeyFlagValue

	}
	return nil, retAdded
}

// parseOperationDevicesGetDevicesDeviceidBackupsBackupidResult parses request result and return the string content
func parseOperationDevicesGetDevicesDeviceidBackupsBackupidResult(respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*devices.GetDevicesDeviceidBackupsBackupidBadRequest)
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
		resp1, ok := iResp1.(*devices.GetDevicesDeviceidBackupsBackupidUnauthorized)
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
		resp2, ok := iResp2.(*devices.GetDevicesDeviceidBackupsBackupidForbidden)
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
		resp3, ok := iResp3.(*devices.GetDevicesDeviceidBackupsBackupidNotFound)
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
		resp4, ok := iResp4.(*devices.GetDevicesDeviceidBackupsBackupidInternalServerError)
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
	return "", nil
}
