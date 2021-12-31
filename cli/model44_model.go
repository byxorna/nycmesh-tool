// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for Model44

// register flags to command
func registerModelModel44Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel44IncludeVlans(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel44NativeVlan(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel44Port(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel44IncludeVlans(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: includeVlans IncludeVlans array type is not supported by go-swagger cli yet

	return nil
}

func registerModel44NativeVlan(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nativeVlanDescription := `Required. Native VLAN identification`

	var nativeVlanFlagName string
	if cmdPrefix == "" {
		nativeVlanFlagName = "nativeVlan"
	} else {
		nativeVlanFlagName = fmt.Sprintf("%v.nativeVlan", cmdPrefix)
	}

	var nativeVlanFlagDefault int64

	_ = cmd.PersistentFlags().Int64(nativeVlanFlagName, nativeVlanFlagDefault, nativeVlanDescription)

	return nil
}

func registerModel44Port(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	portDescription := `Required. Identification of port`

	var portFlagName string
	if cmdPrefix == "" {
		portFlagName = "port"
	} else {
		portFlagName = fmt.Sprintf("%v.port", cmdPrefix)
	}

	var portFlagDefault string

	_ = cmd.PersistentFlags().String(portFlagName, portFlagDefault, portDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModel44Flags(depth int, m *models.Model44, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, includeVlansAdded := retrieveModel44IncludeVlansFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || includeVlansAdded

	err, nativeVlanAdded := retrieveModel44NativeVlanFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nativeVlanAdded

	err, portAdded := retrieveModel44PortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || portAdded

	return nil, retAdded
}

func retrieveModel44IncludeVlansFlags(depth int, m *models.Model44, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	includeVlansFlagName := fmt.Sprintf("%v.includeVlans", cmdPrefix)
	if cmd.Flags().Changed(includeVlansFlagName) {
		// warning: includeVlans array type IncludeVlans is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveModel44NativeVlanFlags(depth int, m *models.Model44, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nativeVlanFlagName := fmt.Sprintf("%v.nativeVlan", cmdPrefix)
	if cmd.Flags().Changed(nativeVlanFlagName) {

		var nativeVlanFlagName string
		if cmdPrefix == "" {
			nativeVlanFlagName = "nativeVlan"
		} else {
			nativeVlanFlagName = fmt.Sprintf("%v.nativeVlan", cmdPrefix)
		}

		nativeVlanFlagValue, err := cmd.Flags().GetInt64(nativeVlanFlagName)
		if err != nil {
			return err, false
		}
		m.NativeVlan = &nativeVlanFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel44PortFlags(depth int, m *models.Model44, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	portFlagName := fmt.Sprintf("%v.port", cmdPrefix)
	if cmd.Flags().Changed(portFlagName) {

		var portFlagName string
		if cmdPrefix == "" {
			portFlagName = "port"
		} else {
			portFlagName = fmt.Sprintf("%v.port", cmdPrefix)
		}

		portFlagValue, err := cmd.Flags().GetString(portFlagName)
		if err != nil {
			return err, false
		}
		m.Port = &portFlagValue

		retAdded = true
	}

	return nil, retAdded
}
