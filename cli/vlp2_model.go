// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for Vlp2

// register flags to command
func registerModelVlp2Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerVlp2Enabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVlp2Ports(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVlp2Tagged(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVlp2VlanID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerVlp2Enabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	enabledDescription := `Required. `

	var enabledFlagName string
	if cmdPrefix == "" {
		enabledFlagName = "enabled"
	} else {
		enabledFlagName = fmt.Sprintf("%v.enabled", cmdPrefix)
	}

	var enabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(enabledFlagName, enabledFlagDefault, enabledDescription)

	return nil
}

func registerVlp2Ports(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: ports Ports3 array type is not supported by go-swagger cli yet

	return nil
}

func registerVlp2Tagged(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	taggedDescription := `Required. `

	var taggedFlagName string
	if cmdPrefix == "" {
		taggedFlagName = "tagged"
	} else {
		taggedFlagName = fmt.Sprintf("%v.tagged", cmdPrefix)
	}

	var taggedFlagDefault bool

	_ = cmd.PersistentFlags().Bool(taggedFlagName, taggedFlagDefault, taggedDescription)

	return nil
}

func registerVlp2VlanID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	vlanIdDescription := `Required. `

	var vlanIdFlagName string
	if cmdPrefix == "" {
		vlanIdFlagName = "vlanId"
	} else {
		vlanIdFlagName = fmt.Sprintf("%v.vlanId", cmdPrefix)
	}

	var vlanIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(vlanIdFlagName, vlanIdFlagDefault, vlanIdDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelVlp2Flags(depth int, m *models.Vlp2, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, enabledAdded := retrieveVlp2EnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || enabledAdded

	err, portsAdded := retrieveVlp2PortsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || portsAdded

	err, taggedAdded := retrieveVlp2TaggedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || taggedAdded

	err, vlanIdAdded := retrieveVlp2VlanIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || vlanIdAdded

	return nil, retAdded
}

func retrieveVlp2EnabledFlags(depth int, m *models.Vlp2, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	enabledFlagName := fmt.Sprintf("%v.enabled", cmdPrefix)
	if cmd.Flags().Changed(enabledFlagName) {

		var enabledFlagName string
		if cmdPrefix == "" {
			enabledFlagName = "enabled"
		} else {
			enabledFlagName = fmt.Sprintf("%v.enabled", cmdPrefix)
		}

		enabledFlagValue, err := cmd.Flags().GetBool(enabledFlagName)
		if err != nil {
			return err, false
		}
		m.Enabled = &enabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVlp2PortsFlags(depth int, m *models.Vlp2, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	portsFlagName := fmt.Sprintf("%v.ports", cmdPrefix)
	if cmd.Flags().Changed(portsFlagName) {
		// warning: ports array type Ports3 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveVlp2TaggedFlags(depth int, m *models.Vlp2, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	taggedFlagName := fmt.Sprintf("%v.tagged", cmdPrefix)
	if cmd.Flags().Changed(taggedFlagName) {

		var taggedFlagName string
		if cmdPrefix == "" {
			taggedFlagName = "tagged"
		} else {
			taggedFlagName = fmt.Sprintf("%v.tagged", cmdPrefix)
		}

		taggedFlagValue, err := cmd.Flags().GetBool(taggedFlagName)
		if err != nil {
			return err, false
		}
		m.Tagged = &taggedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVlp2VlanIDFlags(depth int, m *models.Vlp2, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	vlanIdFlagName := fmt.Sprintf("%v.vlanId", cmdPrefix)
	if cmd.Flags().Changed(vlanIdFlagName) {

		var vlanIdFlagName string
		if cmdPrefix == "" {
			vlanIdFlagName = "vlanId"
		} else {
			vlanIdFlagName = fmt.Sprintf("%v.vlanId", cmdPrefix)
		}

		vlanIdFlagValue, err := cmd.Flags().GetInt64(vlanIdFlagName)
		if err != nil {
			return err, false
		}
		m.VlanID = &vlanIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}
