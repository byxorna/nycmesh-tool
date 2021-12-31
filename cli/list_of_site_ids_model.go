// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for ListOfSiteIds

// register flags to command
func registerModelListOfSiteIdsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerListOfSiteIdsIds(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerListOfSiteIdsIds(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: ids Ids array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelListOfSiteIdsFlags(depth int, m *models.ListOfSiteIds, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, idsAdded := retrieveListOfSiteIdsIdsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idsAdded

	return nil, retAdded
}

func retrieveListOfSiteIdsIdsFlags(depth int, m *models.ListOfSiteIds, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idsFlagName := fmt.Sprintf("%v.ids", cmdPrefix)
	if cmd.Flags().Changed(idsFlagName) {
		// warning: ids array type Ids is not supported by go-swagger cli yet
	}

	return nil, retAdded
}