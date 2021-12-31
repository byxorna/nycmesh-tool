// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for BulkDeleteSites

// register flags to command
func registerModelBulkDeleteSitesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerBulkDeleteSitesDeletedIds(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBulkDeleteSitesMessage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBulkDeleteSitesUndeletedIds(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerBulkDeleteSitesDeletedIds(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: deletedIds DeletedIds array type is not supported by go-swagger cli yet

	return nil
}

func registerBulkDeleteSitesMessage(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerBulkDeleteSitesUndeletedIds(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: undeletedIds UndeletedIds array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelBulkDeleteSitesFlags(depth int, m *models.BulkDeleteSites, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, deletedIdsAdded := retrieveBulkDeleteSitesDeletedIdsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deletedIdsAdded

	err, messageAdded := retrieveBulkDeleteSitesMessageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || messageAdded

	err, undeletedIdsAdded := retrieveBulkDeleteSitesUndeletedIdsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || undeletedIdsAdded

	return nil, retAdded
}

func retrieveBulkDeleteSitesDeletedIdsFlags(depth int, m *models.BulkDeleteSites, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveBulkDeleteSitesMessageFlags(depth int, m *models.BulkDeleteSites, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveBulkDeleteSitesUndeletedIdsFlags(depth int, m *models.BulkDeleteSites, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
