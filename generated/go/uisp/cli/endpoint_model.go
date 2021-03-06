// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for Endpoint

// register flags to command
func registerModelEndpointFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerEndpointAll(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointReadOnly(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointReadWrite(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerEndpointAll(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	allDescription := `Required. Number of client sites with read access.`

	var allFlagName string
	if cmdPrefix == "" {
		allFlagName = "all"
	} else {
		allFlagName = fmt.Sprintf("%v.all", cmdPrefix)
	}

	var allFlagDefault float64

	_ = cmd.PersistentFlags().Float64(allFlagName, allFlagDefault, allDescription)

	return nil
}

func registerEndpointReadOnly(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	readOnlyDescription := `Required. Number of client sites with read-only access.`

	var readOnlyFlagName string
	if cmdPrefix == "" {
		readOnlyFlagName = "readOnly"
	} else {
		readOnlyFlagName = fmt.Sprintf("%v.readOnly", cmdPrefix)
	}

	var readOnlyFlagDefault float64

	_ = cmd.PersistentFlags().Float64(readOnlyFlagName, readOnlyFlagDefault, readOnlyDescription)

	return nil
}

func registerEndpointReadWrite(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	readWriteDescription := `Required. Number of client sites with read-write access.`

	var readWriteFlagName string
	if cmdPrefix == "" {
		readWriteFlagName = "readWrite"
	} else {
		readWriteFlagName = fmt.Sprintf("%v.readWrite", cmdPrefix)
	}

	var readWriteFlagDefault float64

	_ = cmd.PersistentFlags().Float64(readWriteFlagName, readWriteFlagDefault, readWriteDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelEndpointFlags(depth int, m *models.Endpoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, allAdded := retrieveEndpointAllFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || allAdded

	err, readOnlyAdded := retrieveEndpointReadOnlyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || readOnlyAdded

	err, readWriteAdded := retrieveEndpointReadWriteFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || readWriteAdded

	return nil, retAdded
}

func retrieveEndpointAllFlags(depth int, m *models.Endpoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	allFlagName := fmt.Sprintf("%v.all", cmdPrefix)
	if cmd.Flags().Changed(allFlagName) {

		var allFlagName string
		if cmdPrefix == "" {
			allFlagName = "all"
		} else {
			allFlagName = fmt.Sprintf("%v.all", cmdPrefix)
		}

		allFlagValue, err := cmd.Flags().GetFloat64(allFlagName)
		if err != nil {
			return err, false
		}
		m.All = &allFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveEndpointReadOnlyFlags(depth int, m *models.Endpoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	readOnlyFlagName := fmt.Sprintf("%v.readOnly", cmdPrefix)
	if cmd.Flags().Changed(readOnlyFlagName) {

		var readOnlyFlagName string
		if cmdPrefix == "" {
			readOnlyFlagName = "readOnly"
		} else {
			readOnlyFlagName = fmt.Sprintf("%v.readOnly", cmdPrefix)
		}

		readOnlyFlagValue, err := cmd.Flags().GetFloat64(readOnlyFlagName)
		if err != nil {
			return err, false
		}
		m.ReadOnly = &readOnlyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveEndpointReadWriteFlags(depth int, m *models.Endpoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	readWriteFlagName := fmt.Sprintf("%v.readWrite", cmdPrefix)
	if cmd.Flags().Changed(readWriteFlagName) {

		var readWriteFlagName string
		if cmdPrefix == "" {
			readWriteFlagName = "readWrite"
		} else {
			readWriteFlagName = fmt.Sprintf("%v.readWrite", cmdPrefix)
		}

		readWriteFlagValue, err := cmd.Flags().GetFloat64(readWriteFlagName)
		if err != nil {
			return err, false
		}
		m.ReadWrite = &readWriteFlagValue

		retAdded = true
	}

	return nil, retAdded
}
