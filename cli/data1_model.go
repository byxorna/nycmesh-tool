// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for Data1

// register flags to command
func registerModelData1Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerData1DeviceID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerData1DeviceID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	deviceIdDescription := ``

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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelData1Flags(depth int, m *models.Data1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, deviceIdAdded := retrieveData1DeviceIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceIdAdded

	return nil, retAdded
}

func retrieveData1DeviceIDFlags(depth int, m *models.Data1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deviceIdFlagName := fmt.Sprintf("%v.deviceId", cmdPrefix)
	if cmd.Flags().Changed(deviceIdFlagName) {

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

		retAdded = true
	}

	return nil, retAdded
}
