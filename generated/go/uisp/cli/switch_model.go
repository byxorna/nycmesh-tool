// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for Switch

// register flags to command
func registerModelSwitchFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSwitchPorts(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSwitchVlanCapable(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSwitchVlanEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSwitchPorts(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: ports Ports1 array type is not supported by go-swagger cli yet

	return nil
}

func registerSwitchVlanCapable(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	vlanCapableDescription := `Required. `

	var vlanCapableFlagName string
	if cmdPrefix == "" {
		vlanCapableFlagName = "vlanCapable"
	} else {
		vlanCapableFlagName = fmt.Sprintf("%v.vlanCapable", cmdPrefix)
	}

	var vlanCapableFlagDefault bool

	_ = cmd.PersistentFlags().Bool(vlanCapableFlagName, vlanCapableFlagDefault, vlanCapableDescription)

	return nil
}

func registerSwitchVlanEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	vlanEnabledDescription := `Required. `

	var vlanEnabledFlagName string
	if cmdPrefix == "" {
		vlanEnabledFlagName = "vlanEnabled"
	} else {
		vlanEnabledFlagName = fmt.Sprintf("%v.vlanEnabled", cmdPrefix)
	}

	var vlanEnabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(vlanEnabledFlagName, vlanEnabledFlagDefault, vlanEnabledDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSwitchFlags(depth int, m *models.Switch, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, portsAdded := retrieveSwitchPortsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || portsAdded

	err, vlanCapableAdded := retrieveSwitchVlanCapableFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || vlanCapableAdded

	err, vlanEnabledAdded := retrieveSwitchVlanEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || vlanEnabledAdded

	return nil, retAdded
}

func retrieveSwitchPortsFlags(depth int, m *models.Switch, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	portsFlagName := fmt.Sprintf("%v.ports", cmdPrefix)
	if cmd.Flags().Changed(portsFlagName) {
		// warning: ports array type Ports1 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveSwitchVlanCapableFlags(depth int, m *models.Switch, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	vlanCapableFlagName := fmt.Sprintf("%v.vlanCapable", cmdPrefix)
	if cmd.Flags().Changed(vlanCapableFlagName) {

		var vlanCapableFlagName string
		if cmdPrefix == "" {
			vlanCapableFlagName = "vlanCapable"
		} else {
			vlanCapableFlagName = fmt.Sprintf("%v.vlanCapable", cmdPrefix)
		}

		vlanCapableFlagValue, err := cmd.Flags().GetBool(vlanCapableFlagName)
		if err != nil {
			return err, false
		}
		m.VlanCapable = &vlanCapableFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSwitchVlanEnabledFlags(depth int, m *models.Switch, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	vlanEnabledFlagName := fmt.Sprintf("%v.vlanEnabled", cmdPrefix)
	if cmd.Flags().Changed(vlanEnabledFlagName) {

		var vlanEnabledFlagName string
		if cmdPrefix == "" {
			vlanEnabledFlagName = "vlanEnabled"
		} else {
			vlanEnabledFlagName = fmt.Sprintf("%v.vlanEnabled", cmdPrefix)
		}

		vlanEnabledFlagValue, err := cmd.Flags().GetBool(vlanEnabledFlagName)
		if err != nil {
			return err, false
		}
		m.VlanEnabled = &vlanEnabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}
