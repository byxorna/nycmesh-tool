// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for JumboMtu

// register flags to command
func registerModelJumboMtuFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerJumboMtuEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerJumboMtuIsCurrentMtuSynced(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerJumboMtuMaxMtu(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerJumboMtuEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	enabledDescription := `Set to true if jumbo frames are enabled.`

	var enabledFlagName string
	if cmdPrefix == "" {
		enabledFlagName = "enabled"
	} else {
		enabledFlagName = fmt.Sprintf("%v.enabled", cmdPrefix)
	}

	var enabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(enabledFlagName, enabledFlagDefault, enabledDescription)

	return nil
}

func registerJumboMtuIsCurrentMtuSynced(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	isCurrentMtuSyncedDescription := `Set to true if all interfaces have the same mtu.`

	var isCurrentMtuSyncedFlagName string
	if cmdPrefix == "" {
		isCurrentMtuSyncedFlagName = "isCurrentMtuSynced"
	} else {
		isCurrentMtuSyncedFlagName = fmt.Sprintf("%v.isCurrentMtuSynced", cmdPrefix)
	}

	var isCurrentMtuSyncedFlagDefault bool

	_ = cmd.PersistentFlags().Bool(isCurrentMtuSyncedFlagName, isCurrentMtuSyncedFlagDefault, isCurrentMtuSyncedDescription)

	return nil
}

func registerJumboMtuMaxMtu(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	maxMtuDescription := `Maximum configurable MTU value in bytes.`

	var maxMtuFlagName string
	if cmdPrefix == "" {
		maxMtuFlagName = "maxMtu"
	} else {
		maxMtuFlagName = fmt.Sprintf("%v.maxMtu", cmdPrefix)
	}

	var maxMtuFlagDefault int64

	_ = cmd.PersistentFlags().Int64(maxMtuFlagName, maxMtuFlagDefault, maxMtuDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelJumboMtuFlags(depth int, m *models.JumboMtu, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, enabledAdded := retrieveJumboMtuEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || enabledAdded

	err, isCurrentMtuSyncedAdded := retrieveJumboMtuIsCurrentMtuSyncedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || isCurrentMtuSyncedAdded

	err, maxMtuAdded := retrieveJumboMtuMaxMtuFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || maxMtuAdded

	return nil, retAdded
}

func retrieveJumboMtuEnabledFlags(depth int, m *models.JumboMtu, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	enabledFlagName := fmt.Sprintf("%v.enabled", cmdPrefix)
	if cmd.Flags().Changed(enabledFlagName) {

		var enabledFlagName string
		if cmdPrefix == "" {
			enabledFlagName = "enabled"
		} else {
			enabledFlagName = fmt.Sprintf("%v.enabled", cmdPrefix)
		}

		enabledFlagValue, err := cmd.Flags().GetBool(enabledFlagName)
		if err != nil {
			return err, false
		}
		m.Enabled = enabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveJumboMtuIsCurrentMtuSyncedFlags(depth int, m *models.JumboMtu, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	isCurrentMtuSyncedFlagName := fmt.Sprintf("%v.isCurrentMtuSynced", cmdPrefix)
	if cmd.Flags().Changed(isCurrentMtuSyncedFlagName) {

		var isCurrentMtuSyncedFlagName string
		if cmdPrefix == "" {
			isCurrentMtuSyncedFlagName = "isCurrentMtuSynced"
		} else {
			isCurrentMtuSyncedFlagName = fmt.Sprintf("%v.isCurrentMtuSynced", cmdPrefix)
		}

		isCurrentMtuSyncedFlagValue, err := cmd.Flags().GetBool(isCurrentMtuSyncedFlagName)
		if err != nil {
			return err, false
		}
		m.IsCurrentMtuSynced = isCurrentMtuSyncedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveJumboMtuMaxMtuFlags(depth int, m *models.JumboMtu, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	maxMtuFlagName := fmt.Sprintf("%v.maxMtu", cmdPrefix)
	if cmd.Flags().Changed(maxMtuFlagName) {

		var maxMtuFlagName string
		if cmdPrefix == "" {
			maxMtuFlagName = "maxMtu"
		} else {
			maxMtuFlagName = fmt.Sprintf("%v.maxMtu", cmdPrefix)
		}

		maxMtuFlagValue, err := cmd.Flags().GetInt64(maxMtuFlagName)
		if err != nil {
			return err, false
		}
		m.MaxMtu = &maxMtuFlagValue

		retAdded = true
	}

	return nil, retAdded
}