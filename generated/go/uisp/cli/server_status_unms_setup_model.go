// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for ServerStatusUnmsSetup

// register flags to command
func registerModelServerStatusUnmsSetupFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerServerStatusUnmsSetupIsConfigured(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerStatusUnmsSetupIsFirstDeviceConfigured(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerStatusUnmsSetupIsSubnetConfigured(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerServerStatusUnmsSetupIsConfigured(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	isConfiguredDescription := ``

	var isConfiguredFlagName string
	if cmdPrefix == "" {
		isConfiguredFlagName = "isConfigured"
	} else {
		isConfiguredFlagName = fmt.Sprintf("%v.isConfigured", cmdPrefix)
	}

	var isConfiguredFlagDefault bool

	_ = cmd.PersistentFlags().Bool(isConfiguredFlagName, isConfiguredFlagDefault, isConfiguredDescription)

	return nil
}

func registerServerStatusUnmsSetupIsFirstDeviceConfigured(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	isFirstDeviceConfiguredDescription := ``

	var isFirstDeviceConfiguredFlagName string
	if cmdPrefix == "" {
		isFirstDeviceConfiguredFlagName = "isFirstDeviceConfigured"
	} else {
		isFirstDeviceConfiguredFlagName = fmt.Sprintf("%v.isFirstDeviceConfigured", cmdPrefix)
	}

	var isFirstDeviceConfiguredFlagDefault bool

	_ = cmd.PersistentFlags().Bool(isFirstDeviceConfiguredFlagName, isFirstDeviceConfiguredFlagDefault, isFirstDeviceConfiguredDescription)

	return nil
}

func registerServerStatusUnmsSetupIsSubnetConfigured(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	isSubnetConfiguredDescription := ``

	var isSubnetConfiguredFlagName string
	if cmdPrefix == "" {
		isSubnetConfiguredFlagName = "isSubnetConfigured"
	} else {
		isSubnetConfiguredFlagName = fmt.Sprintf("%v.isSubnetConfigured", cmdPrefix)
	}

	var isSubnetConfiguredFlagDefault bool

	_ = cmd.PersistentFlags().Bool(isSubnetConfiguredFlagName, isSubnetConfiguredFlagDefault, isSubnetConfiguredDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServerStatusUnmsSetupFlags(depth int, m *models.ServerStatusUnmsSetup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, isConfiguredAdded := retrieveServerStatusUnmsSetupIsConfiguredFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || isConfiguredAdded

	err, isFirstDeviceConfiguredAdded := retrieveServerStatusUnmsSetupIsFirstDeviceConfiguredFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || isFirstDeviceConfiguredAdded

	err, isSubnetConfiguredAdded := retrieveServerStatusUnmsSetupIsSubnetConfiguredFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || isSubnetConfiguredAdded

	return nil, retAdded
}

func retrieveServerStatusUnmsSetupIsConfiguredFlags(depth int, m *models.ServerStatusUnmsSetup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	isConfiguredFlagName := fmt.Sprintf("%v.isConfigured", cmdPrefix)
	if cmd.Flags().Changed(isConfiguredFlagName) {

		var isConfiguredFlagName string
		if cmdPrefix == "" {
			isConfiguredFlagName = "isConfigured"
		} else {
			isConfiguredFlagName = fmt.Sprintf("%v.isConfigured", cmdPrefix)
		}

		isConfiguredFlagValue, err := cmd.Flags().GetBool(isConfiguredFlagName)
		if err != nil {
			return err, false
		}
		m.IsConfigured = isConfiguredFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerStatusUnmsSetupIsFirstDeviceConfiguredFlags(depth int, m *models.ServerStatusUnmsSetup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	isFirstDeviceConfiguredFlagName := fmt.Sprintf("%v.isFirstDeviceConfigured", cmdPrefix)
	if cmd.Flags().Changed(isFirstDeviceConfiguredFlagName) {

		var isFirstDeviceConfiguredFlagName string
		if cmdPrefix == "" {
			isFirstDeviceConfiguredFlagName = "isFirstDeviceConfigured"
		} else {
			isFirstDeviceConfiguredFlagName = fmt.Sprintf("%v.isFirstDeviceConfigured", cmdPrefix)
		}

		isFirstDeviceConfiguredFlagValue, err := cmd.Flags().GetBool(isFirstDeviceConfiguredFlagName)
		if err != nil {
			return err, false
		}
		m.IsFirstDeviceConfigured = isFirstDeviceConfiguredFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerStatusUnmsSetupIsSubnetConfiguredFlags(depth int, m *models.ServerStatusUnmsSetup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	isSubnetConfiguredFlagName := fmt.Sprintf("%v.isSubnetConfigured", cmdPrefix)
	if cmd.Flags().Changed(isSubnetConfiguredFlagName) {

		var isSubnetConfiguredFlagName string
		if cmdPrefix == "" {
			isSubnetConfiguredFlagName = "isSubnetConfigured"
		} else {
			isSubnetConfiguredFlagName = fmt.Sprintf("%v.isSubnetConfigured", cmdPrefix)
		}

		isSubnetConfiguredFlagValue, err := cmd.Flags().GetBool(isSubnetConfiguredFlagName)
		if err != nil {
			return err, false
		}
		m.IsSubnetConfigured = isSubnetConfiguredFlagValue

		retAdded = true
	}

	return nil, retAdded
}
