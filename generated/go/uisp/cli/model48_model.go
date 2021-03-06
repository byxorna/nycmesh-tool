// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for Model48

// register flags to command
func registerModelModel48Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel48Interface(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel48Mode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel48Interface(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var interfaceFlagName string
	if cmdPrefix == "" {
		interfaceFlagName = "interface"
	} else {
		interfaceFlagName = fmt.Sprintf("%v.interface", cmdPrefix)
	}

	if err := registerModelInterface2Flags(depth+1, interfaceFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel48Mode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	modeDescription := `Enum: ["tagged","untagged","quiq"]. Required. VLAN mode`

	var modeFlagName string
	if cmdPrefix == "" {
		modeFlagName = "mode"
	} else {
		modeFlagName = fmt.Sprintf("%v.mode", cmdPrefix)
	}

	var modeFlagDefault string

	_ = cmd.PersistentFlags().String(modeFlagName, modeFlagDefault, modeDescription)

	if err := cmd.RegisterFlagCompletionFunc(modeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["tagged","untagged","quiq"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModel48Flags(depth int, m *models.Model48, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, interfaceAdded := retrieveModel48InterfaceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || interfaceAdded

	err, modeAdded := retrieveModel48ModeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || modeAdded

	return nil, retAdded
}

func retrieveModel48InterfaceFlags(depth int, m *models.Model48, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	interfaceFlagName := fmt.Sprintf("%v.interface", cmdPrefix)
	if cmd.Flags().Changed(interfaceFlagName) {
		// info: complex object interface Interface2 is retrieved outside this Changed() block
	}
	interfaceFlagValue := m.Interface
	if swag.IsZero(interfaceFlagValue) {
		interfaceFlagValue = &models.Interface2{}
	}

	err, interfaceAdded := retrieveModelInterface2Flags(depth+1, interfaceFlagValue, interfaceFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || interfaceAdded
	if interfaceAdded {
		m.Interface = interfaceFlagValue
	}

	return nil, retAdded
}

func retrieveModel48ModeFlags(depth int, m *models.Model48, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	modeFlagName := fmt.Sprintf("%v.mode", cmdPrefix)
	if cmd.Flags().Changed(modeFlagName) {

		var modeFlagName string
		if cmdPrefix == "" {
			modeFlagName = "mode"
		} else {
			modeFlagName = fmt.Sprintf("%v.mode", cmdPrefix)
		}

		modeFlagValue, err := cmd.Flags().GetString(modeFlagName)
		if err != nil {
			return err, false
		}
		m.Mode = &modeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
