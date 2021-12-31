// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for PayloadSystemUnmsKey

// register flags to command
func registerModelPayloadSystemUnmsKeyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPayloadSystemUnmsKeyDeviceIds(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPayloadSystemUnmsKeyUnmsKey(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPayloadSystemUnmsKeyDeviceIds(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: deviceIds DeviceIds array type is not supported by go-swagger cli yet

	return nil
}

func registerPayloadSystemUnmsKeyUnmsKey(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	unmsKeyDescription := `Required. `

	var unmsKeyFlagName string
	if cmdPrefix == "" {
		unmsKeyFlagName = "unmsKey"
	} else {
		unmsKeyFlagName = fmt.Sprintf("%v.unmsKey", cmdPrefix)
	}

	var unmsKeyFlagDefault string

	_ = cmd.PersistentFlags().String(unmsKeyFlagName, unmsKeyFlagDefault, unmsKeyDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPayloadSystemUnmsKeyFlags(depth int, m *models.PayloadSystemUnmsKey, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, deviceIdsAdded := retrievePayloadSystemUnmsKeyDeviceIdsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceIdsAdded

	err, unmsKeyAdded := retrievePayloadSystemUnmsKeyUnmsKeyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || unmsKeyAdded

	return nil, retAdded
}

func retrievePayloadSystemUnmsKeyDeviceIdsFlags(depth int, m *models.PayloadSystemUnmsKey, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deviceIdsFlagName := fmt.Sprintf("%v.deviceIds", cmdPrefix)
	if cmd.Flags().Changed(deviceIdsFlagName) {
		// warning: deviceIds array type DeviceIds is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrievePayloadSystemUnmsKeyUnmsKeyFlags(depth int, m *models.PayloadSystemUnmsKey, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	unmsKeyFlagName := fmt.Sprintf("%v.unmsKey", cmdPrefix)
	if cmd.Flags().Changed(unmsKeyFlagName) {

		var unmsKeyFlagName string
		if cmdPrefix == "" {
			unmsKeyFlagName = "unmsKey"
		} else {
			unmsKeyFlagName = fmt.Sprintf("%v.unmsKey", cmdPrefix)
		}

		unmsKeyFlagValue, err := cmd.Flags().GetString(unmsKeyFlagName)
		if err != nil {
			return err, false
		}
		m.UnmsKey = &unmsKeyFlagValue

		retAdded = true
	}

	return nil, retAdded
}
