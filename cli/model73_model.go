// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for Model73

// register flags to command
func registerModelModel73Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel73Ids(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel73Ids(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: ids Ids array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModel73Flags(depth int, m *models.Model73, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, idsAdded := retrieveModel73IdsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idsAdded

	return nil, retAdded
}

func retrieveModel73IdsFlags(depth int, m *models.Model73, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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