// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// Schema cli for Uisps

// register flags to command
func registerModelUispsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUispsVlans(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUispsVlans(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var vlansFlagName string
	if cmdPrefix == "" {
		vlansFlagName = "vlans"
	} else {
		vlansFlagName = fmt.Sprintf("%v.vlans", cmdPrefix)
	}

	if err := registerModelVlansFlags(depth+1, vlansFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUispsFlags(depth int, m *models.Uisps, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, vlansAdded := retrieveUispsVlansFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || vlansAdded

	return nil, retAdded
}

func retrieveUispsVlansFlags(depth int, m *models.Uisps, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	vlansFlagName := fmt.Sprintf("%v.vlans", cmdPrefix)
	if cmd.Flags().Changed(vlansFlagName) {
		// info: complex object vlans Vlans is retrieved outside this Changed() block
	}
	vlansFlagValue := m.Vlans
	if swag.IsZero(vlansFlagValue) {
		vlansFlagValue = &models.Vlans{}
	}

	err, vlansAdded := retrieveModelVlansFlags(depth+1, vlansFlagValue, vlansFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || vlansAdded
	if vlansAdded {
		m.Vlans = vlansFlagValue
	}

	return nil, retAdded
}
