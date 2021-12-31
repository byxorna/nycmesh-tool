// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for Model95

// register flags to command
func registerModelModel95Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel95Addresses(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel95Bridge(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel95BridgeID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel95Description(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel95Name(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel95ProxyARP(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel95Addresses(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: addresses Addresses4 array type is not supported by go-swagger cli yet

	return nil
}

func registerModel95Bridge(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var bridgeFlagName string
	if cmdPrefix == "" {
		bridgeFlagName = "bridge"
	} else {
		bridgeFlagName = fmt.Sprintf("%v.bridge", cmdPrefix)
	}

	if err := registerModelBridge1Flags(depth+1, bridgeFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel95BridgeID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	bridgeIdDescription := `Required. `

	var bridgeIdFlagName string
	if cmdPrefix == "" {
		bridgeIdFlagName = "bridgeId"
	} else {
		bridgeIdFlagName = fmt.Sprintf("%v.bridgeId", cmdPrefix)
	}

	var bridgeIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(bridgeIdFlagName, bridgeIdFlagDefault, bridgeIdDescription)

	return nil
}

func registerModel95Description(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := ``

	var descriptionFlagName string
	if cmdPrefix == "" {
		descriptionFlagName = "description"
	} else {
		descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
	}

	var descriptionFlagDefault string

	_ = cmd.PersistentFlags().String(descriptionFlagName, descriptionFlagDefault, descriptionDescription)

	return nil
}

func registerModel95Name(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModel95ProxyARP(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	proxyARPDescription := ``

	var proxyARPFlagName string
	if cmdPrefix == "" {
		proxyARPFlagName = "proxyARP"
	} else {
		proxyARPFlagName = fmt.Sprintf("%v.proxyARP", cmdPrefix)
	}

	var proxyARPFlagDefault bool

	_ = cmd.PersistentFlags().Bool(proxyARPFlagName, proxyARPFlagDefault, proxyARPDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModel95Flags(depth int, m *models.Model95, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, addressesAdded := retrieveModel95AddressesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || addressesAdded

	err, bridgeAdded := retrieveModel95BridgeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || bridgeAdded

	err, bridgeIdAdded := retrieveModel95BridgeIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || bridgeIdAdded

	err, descriptionAdded := retrieveModel95DescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, nameAdded := retrieveModel95NameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, proxyARPAdded := retrieveModel95ProxyARPFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || proxyARPAdded

	return nil, retAdded
}

func retrieveModel95AddressesFlags(depth int, m *models.Model95, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	addressesFlagName := fmt.Sprintf("%v.addresses", cmdPrefix)
	if cmd.Flags().Changed(addressesFlagName) {
		// warning: addresses array type Addresses4 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveModel95BridgeFlags(depth int, m *models.Model95, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	bridgeFlagName := fmt.Sprintf("%v.bridge", cmdPrefix)
	if cmd.Flags().Changed(bridgeFlagName) {
		// info: complex object bridge Bridge1 is retrieved outside this Changed() block
	}
	bridgeFlagValue := m.Bridge
	if swag.IsZero(bridgeFlagValue) {
		bridgeFlagValue = &models.Bridge1{}
	}

	err, bridgeAdded := retrieveModelBridge1Flags(depth+1, bridgeFlagValue, bridgeFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || bridgeAdded
	if bridgeAdded {
		m.Bridge = bridgeFlagValue
	}

	return nil, retAdded
}

func retrieveModel95BridgeIDFlags(depth int, m *models.Model95, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	bridgeIdFlagName := fmt.Sprintf("%v.bridgeId", cmdPrefix)
	if cmd.Flags().Changed(bridgeIdFlagName) {

		var bridgeIdFlagName string
		if cmdPrefix == "" {
			bridgeIdFlagName = "bridgeId"
		} else {
			bridgeIdFlagName = fmt.Sprintf("%v.bridgeId", cmdPrefix)
		}

		bridgeIdFlagValue, err := cmd.Flags().GetInt64(bridgeIdFlagName)
		if err != nil {
			return err, false
		}
		m.BridgeID = &bridgeIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel95DescriptionFlags(depth int, m *models.Model95, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	descriptionFlagName := fmt.Sprintf("%v.description", cmdPrefix)
	if cmd.Flags().Changed(descriptionFlagName) {

		var descriptionFlagName string
		if cmdPrefix == "" {
			descriptionFlagName = "description"
		} else {
			descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
		}

		descriptionFlagValue, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			return err, false
		}
		m.Description = descriptionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel95NameFlags(depth int, m *models.Model95, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel95ProxyARPFlags(depth int, m *models.Model95, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	proxyARPFlagName := fmt.Sprintf("%v.proxyARP", cmdPrefix)
	if cmd.Flags().Changed(proxyARPFlagName) {

		var proxyARPFlagName string
		if cmdPrefix == "" {
			proxyARPFlagName = "proxyARP"
		} else {
			proxyARPFlagName = fmt.Sprintf("%v.proxyARP", cmdPrefix)
		}

		proxyARPFlagValue, err := cmd.Flags().GetBool(proxyARPFlagName)
		if err != nil {
			return err, false
		}
		m.ProxyARP = proxyARPFlagValue

		retAdded = true
	}

	return nil, retAdded
}
