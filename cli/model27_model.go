// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"
"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for Model27

// register flags to command
func registerModelModel27Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel27Interface(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel27Interface(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var interfaceFlagName string
	if cmdPrefix == "" {
		interfaceFlagName = "interface"
	} else {
		interfaceFlagName = fmt.Sprintf("%v.interface", cmdPrefix)
	}

	if err := registerModelVlansInterfaceSchemaFlags(depth+1, interfaceFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModel27Flags(depth int, m *models.Model27, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, interfaceAdded := retrieveModel27InterfaceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || interfaceAdded

	return nil, retAdded
}

func retrieveModel27InterfaceFlags(depth int, m *models.Model27, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	interfaceFlagName := fmt.Sprintf("%v.interface", cmdPrefix)
	if cmd.Flags().Changed(interfaceFlagName) {
		// info: complex object interface VlansInterfaceSchema is retrieved outside this Changed() block
	}
	interfaceFlagValue := m.Interface
	if swag.IsZero(interfaceFlagValue) {
		interfaceFlagValue = &models.VlansInterfaceSchema{}
	}

	err, interfaceAdded := retrieveModelVlansInterfaceSchemaFlags(depth+1, interfaceFlagValue, interfaceFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || interfaceAdded
	if interfaceAdded {
		m.Interface = interfaceFlagValue
	}

	return nil, retAdded
}
