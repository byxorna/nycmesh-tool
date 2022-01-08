// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for Model93

// register flags to command
func registerModelModel93Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel93Cidr(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel93Type(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel93Cidr(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	cidrDescription := `Required. `

	var cidrFlagName string
	if cmdPrefix == "" {
		cidrFlagName = "cidr"
	} else {
		cidrFlagName = fmt.Sprintf("%v.cidr", cmdPrefix)
	}

	var cidrFlagDefault string

	_ = cmd.PersistentFlags().String(cidrFlagName, cidrFlagDefault, cidrDescription)

	return nil
}

func registerModel93Type(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["static","dhcp","dhcpv6"]. Required. `

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	if err := cmd.RegisterFlagCompletionFunc(typeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["static","dhcp","dhcpv6"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModel93Flags(depth int, m *models.Model93, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, cidrAdded := retrieveModel93CidrFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || cidrAdded

	err, typeAdded := retrieveModel93TypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveModel93CidrFlags(depth int, m *models.Model93, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	cidrFlagName := fmt.Sprintf("%v.cidr", cmdPrefix)
	if cmd.Flags().Changed(cidrFlagName) {

		var cidrFlagName string
		if cmdPrefix == "" {
			cidrFlagName = "cidr"
		} else {
			cidrFlagName = fmt.Sprintf("%v.cidr", cmdPrefix)
		}

		cidrFlagValue, err := cmd.Flags().GetString(cidrFlagName)
		if err != nil {
			return err, false
		}
		m.Cidr = &cidrFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel93TypeFlags(depth int, m *models.Model93, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	typeFlagName := fmt.Sprintf("%v.type", cmdPrefix)
	if cmd.Flags().Changed(typeFlagName) {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = &typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}