// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for DataVlan

// register flags to command
func registerModelDataVlanFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDataVlanEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDataVlanID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDataVlanEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	enabledDescription := `Required. `

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

func registerDataVlanID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `Required. `

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault int64

	_ = cmd.PersistentFlags().Int64(idFlagName, idFlagDefault, idDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDataVlanFlags(depth int, m *models.DataVlan, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, enabledAdded := retrieveDataVlanEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || enabledAdded

	err, idAdded := retrieveDataVlanIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	return nil, retAdded
}

func retrieveDataVlanEnabledFlags(depth int, m *models.DataVlan, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Enabled = &enabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDataVlanIDFlags(depth int, m *models.DataVlan, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetInt64(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = &idFlagValue

		retAdded = true
	}

	return nil, retAdded
}