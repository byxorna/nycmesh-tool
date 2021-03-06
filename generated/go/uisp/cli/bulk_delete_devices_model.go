// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for BulkDeleteDevices

// register flags to command
func registerModelBulkDeleteDevicesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerBulkDeleteDevicesDeletedIds(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBulkDeleteDevicesDiscovered(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBulkDeleteDevicesMessage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBulkDeleteDevicesUndeletedIds(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerBulkDeleteDevicesDeletedIds(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: deletedIds DeletedIds array type is not supported by go-swagger cli yet

	return nil
}

func registerBulkDeleteDevicesDiscovered(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	discoveredDescription := ``

	var discoveredFlagName string
	if cmdPrefix == "" {
		discoveredFlagName = "discovered"
	} else {
		discoveredFlagName = fmt.Sprintf("%v.discovered", cmdPrefix)
	}

	var discoveredFlagDefault bool

	_ = cmd.PersistentFlags().Bool(discoveredFlagName, discoveredFlagDefault, discoveredDescription)

	return nil
}

func registerBulkDeleteDevicesMessage(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	messageDescription := `Required. `

	var messageFlagName string
	if cmdPrefix == "" {
		messageFlagName = "message"
	} else {
		messageFlagName = fmt.Sprintf("%v.message", cmdPrefix)
	}

	var messageFlagDefault string

	_ = cmd.PersistentFlags().String(messageFlagName, messageFlagDefault, messageDescription)

	return nil
}

func registerBulkDeleteDevicesUndeletedIds(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: undeletedIds UndeletedIds array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelBulkDeleteDevicesFlags(depth int, m *models.BulkDeleteDevices, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, deletedIdsAdded := retrieveBulkDeleteDevicesDeletedIdsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deletedIdsAdded

	err, discoveredAdded := retrieveBulkDeleteDevicesDiscoveredFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || discoveredAdded

	err, messageAdded := retrieveBulkDeleteDevicesMessageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || messageAdded

	err, undeletedIdsAdded := retrieveBulkDeleteDevicesUndeletedIdsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || undeletedIdsAdded

	return nil, retAdded
}

func retrieveBulkDeleteDevicesDeletedIdsFlags(depth int, m *models.BulkDeleteDevices, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deletedIdsFlagName := fmt.Sprintf("%v.deletedIds", cmdPrefix)
	if cmd.Flags().Changed(deletedIdsFlagName) {
		// warning: deletedIds array type DeletedIds is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveBulkDeleteDevicesDiscoveredFlags(depth int, m *models.BulkDeleteDevices, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	discoveredFlagName := fmt.Sprintf("%v.discovered", cmdPrefix)
	if cmd.Flags().Changed(discoveredFlagName) {

		var discoveredFlagName string
		if cmdPrefix == "" {
			discoveredFlagName = "discovered"
		} else {
			discoveredFlagName = fmt.Sprintf("%v.discovered", cmdPrefix)
		}

		discoveredFlagValue, err := cmd.Flags().GetBool(discoveredFlagName)
		if err != nil {
			return err, false
		}
		m.Discovered = discoveredFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveBulkDeleteDevicesMessageFlags(depth int, m *models.BulkDeleteDevices, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	messageFlagName := fmt.Sprintf("%v.message", cmdPrefix)
	if cmd.Flags().Changed(messageFlagName) {

		var messageFlagName string
		if cmdPrefix == "" {
			messageFlagName = "message"
		} else {
			messageFlagName = fmt.Sprintf("%v.message", cmdPrefix)
		}

		messageFlagValue, err := cmd.Flags().GetString(messageFlagName)
		if err != nil {
			return err, false
		}
		m.Message = &messageFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveBulkDeleteDevicesUndeletedIdsFlags(depth int, m *models.BulkDeleteDevices, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	undeletedIdsFlagName := fmt.Sprintf("%v.undeletedIds", cmdPrefix)
	if cmd.Flags().Changed(undeletedIdsFlagName) {
		// warning: undeletedIds array type UndeletedIds is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
