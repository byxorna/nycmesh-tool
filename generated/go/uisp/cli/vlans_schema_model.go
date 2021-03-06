// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for VlansSchema

// register flags to command
func registerModelVlansSchemaFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerVlansSchemaVlans(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerVlansSchemaVlans(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: vlans Vlans4 array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelVlansSchemaFlags(depth int, m *models.VlansSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, vlansAdded := retrieveVlansSchemaVlansFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || vlansAdded

	return nil, retAdded
}

func retrieveVlansSchemaVlansFlags(depth int, m *models.VlansSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	vlansFlagName := fmt.Sprintf("%v.vlans", cmdPrefix)
	if cmd.Flags().Changed(vlansFlagName) {
		// warning: vlans array type Vlans4 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
