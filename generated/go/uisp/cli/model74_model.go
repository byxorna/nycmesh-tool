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

// Schema cli for Model74

// register flags to command
func registerModelModel74Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel74Addresses(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel74Index(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel74Mac(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel74Name(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel74Type(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel74Addresses(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: addresses Addresses3 array type is not supported by go-swagger cli yet

	return nil
}

func registerModel74Index(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	indexDescription := `Required. `

	var indexFlagName string
	if cmdPrefix == "" {
		indexFlagName = "index"
	} else {
		indexFlagName = fmt.Sprintf("%v.index", cmdPrefix)
	}

	var indexFlagDefault int64

	_ = cmd.PersistentFlags().Int64(indexFlagName, indexFlagDefault, indexDescription)

	return nil
}

func registerModel74Mac(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModel74Name(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Required. `

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

func registerModel74Type(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["eth","sfp+","wlan"]. `

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string = "eth"

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	if err := cmd.RegisterFlagCompletionFunc(typeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["eth","sfp+","wlan"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModel74Flags(depth int, m *models.Model74, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, addressesAdded := retrieveModel74AddressesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || addressesAdded

	err, indexAdded := retrieveModel74IndexFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || indexAdded

	err, macAdded := retrieveModel74MacFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || macAdded

	err, nameAdded := retrieveModel74NameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, typeAdded := retrieveModel74TypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveModel74AddressesFlags(depth int, m *models.Model74, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	addressesFlagName := fmt.Sprintf("%v.addresses", cmdPrefix)
	if cmd.Flags().Changed(addressesFlagName) {
		// warning: addresses array type Addresses3 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveModel74IndexFlags(depth int, m *models.Model74, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	indexFlagName := fmt.Sprintf("%v.index", cmdPrefix)
	if cmd.Flags().Changed(indexFlagName) {

		var indexFlagName string
		if cmdPrefix == "" {
			indexFlagName = "index"
		} else {
			indexFlagName = fmt.Sprintf("%v.index", cmdPrefix)
		}

		indexFlagValue, err := cmd.Flags().GetInt64(indexFlagName)
		if err != nil {
			return err, false
		}
		m.Index = &indexFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel74MacFlags(depth int, m *models.Model74, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel74NameFlags(depth int, m *models.Model74, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Name = &nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel74TypeFlags(depth int, m *models.Model74, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
