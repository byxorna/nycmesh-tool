// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for Model107

// register flags to command
func registerModelModel107Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel107UnmsKey(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel107UnmsKey(depth int, cmdPrefix string, cmd *cobra.Command) error {
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
func retrieveModelModel107Flags(depth int, m *models.Model107, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, unmsKeyAdded := retrieveModel107UnmsKeyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || unmsKeyAdded

	return nil, retAdded
}

func retrieveModel107UnmsKeyFlags(depth int, m *models.Model107, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
