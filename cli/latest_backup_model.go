// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for LatestBackup

// register flags to command
func registerModelLatestBackupFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerLatestBackupID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerLatestBackupTimestamp(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerLatestBackupID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `Required. Latest backup ID.`

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerLatestBackupTimestamp(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: primitive timestamp strfmt.Date is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelLatestBackupFlags(depth int, m *models.LatestBackup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, idAdded := retrieveLatestBackupIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, timestampAdded := retrieveLatestBackupTimestampFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || timestampAdded

	return nil, retAdded
}

func retrieveLatestBackupIDFlags(depth int, m *models.LatestBackup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = &idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveLatestBackupTimestampFlags(depth int, m *models.LatestBackup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	timestampFlagName := fmt.Sprintf("%v.timestamp", cmdPrefix)
	if cmd.Flags().Changed(timestampFlagName) {

		// warning: primitive timestamp strfmt.Date is not supported by go-swagger cli yet

		retAdded = true
	}

	return nil, retAdded
}
