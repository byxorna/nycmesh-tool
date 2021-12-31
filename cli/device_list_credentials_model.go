// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for DeviceListCredentials

// register flags to command
func registerModelDeviceListCredentialsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDeviceListCredentialsCredentials(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceListCredentialsIsPassphraseMissing(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceListCredentialsCredentials(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: credentials Credentials array type is not supported by go-swagger cli yet

	return nil
}

func registerDeviceListCredentialsIsPassphraseMissing(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	isPassphraseMissingDescription := ``

	var isPassphraseMissingFlagName string
	if cmdPrefix == "" {
		isPassphraseMissingFlagName = "isPassphraseMissing"
	} else {
		isPassphraseMissingFlagName = fmt.Sprintf("%v.isPassphraseMissing", cmdPrefix)
	}

	var isPassphraseMissingFlagDefault bool

	_ = cmd.PersistentFlags().Bool(isPassphraseMissingFlagName, isPassphraseMissingFlagDefault, isPassphraseMissingDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDeviceListCredentialsFlags(depth int, m *models.DeviceListCredentials, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, credentialsAdded := retrieveDeviceListCredentialsCredentialsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || credentialsAdded

	err, isPassphraseMissingAdded := retrieveDeviceListCredentialsIsPassphraseMissingFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || isPassphraseMissingAdded

	return nil, retAdded
}

func retrieveDeviceListCredentialsCredentialsFlags(depth int, m *models.DeviceListCredentials, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	credentialsFlagName := fmt.Sprintf("%v.credentials", cmdPrefix)
	if cmd.Flags().Changed(credentialsFlagName) {
		// warning: credentials array type Credentials is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDeviceListCredentialsIsPassphraseMissingFlags(depth int, m *models.DeviceListCredentials, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	isPassphraseMissingFlagName := fmt.Sprintf("%v.isPassphraseMissing", cmdPrefix)
	if cmd.Flags().Changed(isPassphraseMissingFlagName) {

		var isPassphraseMissingFlagName string
		if cmdPrefix == "" {
			isPassphraseMissingFlagName = "isPassphraseMissing"
		} else {
			isPassphraseMissingFlagName = fmt.Sprintf("%v.isPassphraseMissing", cmdPrefix)
		}

		isPassphraseMissingFlagValue, err := cmd.Flags().GetBool(isPassphraseMissingFlagName)
		if err != nil {
			return err, false
		}
		m.IsPassphraseMissing = isPassphraseMissingFlagValue

		retAdded = true
	}

	return nil, retAdded
}
