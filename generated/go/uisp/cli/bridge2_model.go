// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for Bridge2

// register flags to command
func registerModelBridge2Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerBridge2IncludeVlans(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBridge2NativeVlan(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerBridge2IncludeVlans(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: includeVlans IncludeVlans2 map type is not supported by go-swagger cli yet

	return nil
}

func registerBridge2NativeVlan(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: nativeVlan NativeVlan map type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelBridge2Flags(depth int, m *models.Bridge2, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, includeVlansAdded := retrieveBridge2IncludeVlansFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || includeVlansAdded

	err, nativeVlanAdded := retrieveBridge2NativeVlanFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nativeVlanAdded

	return nil, retAdded
}

func retrieveBridge2IncludeVlansFlags(depth int, m *models.Bridge2, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	includeVlansFlagName := fmt.Sprintf("%v.includeVlans", cmdPrefix)
	if cmd.Flags().Changed(includeVlansFlagName) {
		// warning: includeVlans map type IncludeVlans2 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveBridge2NativeVlanFlags(depth int, m *models.Bridge2, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nativeVlanFlagName := fmt.Sprintf("%v.nativeVlan", cmdPrefix)
	if cmd.Flags().Changed(nativeVlanFlagName) {
		// warning: nativeVlan map type NativeVlan is not supported by go-swagger cli yet
	}

	return nil, retAdded
}