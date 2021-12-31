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

// Schema cli for OutInterface1

// register flags to command
func registerModelOutInterface1Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerOutInterface1ID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOutInterface1Mac(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOutInterface1MacOverride(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOutInterface1Name(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOutInterface1Type(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerOutInterface1ID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := ``

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerOutInterface1Mac(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	macDescription := ``

	var macFlagName string
	if cmdPrefix == "" {
		macFlagName = "mac"
	} else {
		macFlagName = fmt.Sprintf("%v.mac", cmdPrefix)
	}

	var macFlagDefault string

	_ = cmd.PersistentFlags().String(macFlagName, macFlagDefault, macDescription)

	return nil
}

func registerOutInterface1MacOverride(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	macOverrideDescription := ``

	var macOverrideFlagName string
	if cmdPrefix == "" {
		macOverrideFlagName = "macOverride"
	} else {
		macOverrideFlagName = fmt.Sprintf("%v.macOverride", cmdPrefix)
	}

	var macOverrideFlagDefault string

	_ = cmd.PersistentFlags().String(macOverrideFlagName, macOverrideFlagDefault, macOverrideDescription)

	return nil
}

func registerOutInterface1Name(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := ``

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerOutInterface1Type(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["eth","switch","port","pon","br","pppoe","vlan","sfp+","wlan","ath","lag","loopback","tunnel"]. `

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
			if err := json.Unmarshal([]byte(`["eth","switch","port","pon","br","pppoe","vlan","sfp+","wlan","ath","lag","loopback","tunnel"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelOutInterface1Flags(depth int, m *models.OutInterface1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, idAdded := retrieveOutInterface1IDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, macAdded := retrieveOutInterface1MacFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || macAdded

	err, macOverrideAdded := retrieveOutInterface1MacOverrideFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || macOverrideAdded

	err, nameAdded := retrieveOutInterface1NameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, typeAdded := retrieveOutInterface1TypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveOutInterface1IDFlags(depth int, m *models.OutInterface1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOutInterface1MacFlags(depth int, m *models.OutInterface1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	macFlagName := fmt.Sprintf("%v.mac", cmdPrefix)
	if cmd.Flags().Changed(macFlagName) {

		var macFlagName string
		if cmdPrefix == "" {
			macFlagName = "mac"
		} else {
			macFlagName = fmt.Sprintf("%v.mac", cmdPrefix)
		}

		macFlagValue, err := cmd.Flags().GetString(macFlagName)
		if err != nil {
			return err, false
		}
		m.Mac = macFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOutInterface1MacOverrideFlags(depth int, m *models.OutInterface1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	macOverrideFlagName := fmt.Sprintf("%v.macOverride", cmdPrefix)
	if cmd.Flags().Changed(macOverrideFlagName) {

		var macOverrideFlagName string
		if cmdPrefix == "" {
			macOverrideFlagName = "macOverride"
		} else {
			macOverrideFlagName = fmt.Sprintf("%v.macOverride", cmdPrefix)
		}

		macOverrideFlagValue, err := cmd.Flags().GetString(macOverrideFlagName)
		if err != nil {
			return err, false
		}
		m.MacOverride = macOverrideFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOutInterface1NameFlags(depth int, m *models.OutInterface1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOutInterface1TypeFlags(depth int, m *models.OutInterface1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Type = typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
