// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for DevicesBackup

// register flags to command
func registerModelDevicesBackupFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDevicesBackupDeviceIds(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDevicesBackupDeviceIds(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: deviceIds DeviceIds1 array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDevicesBackupFlags(depth int, m *models.DevicesBackup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, deviceIdsAdded := retrieveDevicesBackupDeviceIdsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceIdsAdded

	return nil, retAdded
}

func retrieveDevicesBackupDeviceIdsFlags(depth int, m *models.DevicesBackup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deviceIdsFlagName := fmt.Sprintf("%v.deviceIds", cmdPrefix)
	if cmd.Flags().Changed(deviceIdsFlagName) {
		// warning: deviceIds array type DeviceIds1 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
