// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for OnuPolicies1

// register flags to command
func registerModelOnuPolicies1Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerOnuPolicies1DefaultState(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuPolicies1DhcpOption82(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerOnuPolicies1DefaultState(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	defaultStateDescription := `Enum: ["enabled","disabled"]. Required. `

	var defaultStateFlagName string
	if cmdPrefix == "" {
		defaultStateFlagName = "defaultState"
	} else {
		defaultStateFlagName = fmt.Sprintf("%v.defaultState", cmdPrefix)
	}

	var defaultStateFlagDefault string

	_ = cmd.PersistentFlags().String(defaultStateFlagName, defaultStateFlagDefault, defaultStateDescription)

	if err := cmd.RegisterFlagCompletionFunc(defaultStateFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerOnuPolicies1DhcpOption82(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	dhcpOption82Description := ``

	var dhcpOption82FlagName string
	if cmdPrefix == "" {
		dhcpOption82FlagName = "dhcpOption82"
	} else {
		dhcpOption82FlagName = fmt.Sprintf("%v.dhcpOption82", cmdPrefix)
	}

	var dhcpOption82FlagDefault bool

	_ = cmd.PersistentFlags().Bool(dhcpOption82FlagName, dhcpOption82FlagDefault, dhcpOption82Description)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelOnuPolicies1Flags(depth int, m *models.OnuPolicies1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, defaultStateAdded := retrieveOnuPolicies1DefaultStateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || defaultStateAdded

	err, dhcpOption82Added := retrieveOnuPolicies1DhcpOption82Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dhcpOption82Added

	return nil, retAdded
}

func retrieveOnuPolicies1DefaultStateFlags(depth int, m *models.OnuPolicies1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	defaultStateFlagName := fmt.Sprintf("%v.defaultState", cmdPrefix)
	if cmd.Flags().Changed(defaultStateFlagName) {

		var defaultStateFlagName string
		if cmdPrefix == "" {
			defaultStateFlagName = "defaultState"
		} else {
			defaultStateFlagName = fmt.Sprintf("%v.defaultState", cmdPrefix)
		}

		defaultStateFlagValue, err := cmd.Flags().GetString(defaultStateFlagName)
		if err != nil {
			return err, false
		}
		m.DefaultState = &defaultStateFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuPolicies1DhcpOption82Flags(depth int, m *models.OnuPolicies1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dhcpOption82FlagName := fmt.Sprintf("%v.dhcpOption82", cmdPrefix)
	if cmd.Flags().Changed(dhcpOption82FlagName) {

		var dhcpOption82FlagName string
		if cmdPrefix == "" {
			dhcpOption82FlagName = "dhcpOption82"
		} else {
			dhcpOption82FlagName = fmt.Sprintf("%v.dhcpOption82", cmdPrefix)
		}

		dhcpOption82FlagValue, err := cmd.Flags().GetBool(dhcpOption82FlagName)
		if err != nil {
			return err, false
		}
		m.DhcpOption82 = dhcpOption82FlagValue

		retAdded = true
	}

	return nil, retAdded
}