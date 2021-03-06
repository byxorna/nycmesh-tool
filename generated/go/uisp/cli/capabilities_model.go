// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for Capabilities

// register flags to command
func registerModelCapabilitiesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCapabilitiesLoadBalanceValues(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCapabilitiesSupportAutoEdge(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCapabilitiesSupportCableTest(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCapabilitiesSupportReset(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCapabilitiesLoadBalanceValues(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: loadBalanceValues LoadBalanceValues array type is not supported by go-swagger cli yet

	return nil
}

func registerCapabilitiesSupportAutoEdge(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	supportAutoEdgeDescription := ``

	var supportAutoEdgeFlagName string
	if cmdPrefix == "" {
		supportAutoEdgeFlagName = "supportAutoEdge"
	} else {
		supportAutoEdgeFlagName = fmt.Sprintf("%v.supportAutoEdge", cmdPrefix)
	}

	var supportAutoEdgeFlagDefault bool

	_ = cmd.PersistentFlags().Bool(supportAutoEdgeFlagName, supportAutoEdgeFlagDefault, supportAutoEdgeDescription)

	return nil
}

func registerCapabilitiesSupportCableTest(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	supportCableTestDescription := ``

	var supportCableTestFlagName string
	if cmdPrefix == "" {
		supportCableTestFlagName = "supportCableTest"
	} else {
		supportCableTestFlagName = fmt.Sprintf("%v.supportCableTest", cmdPrefix)
	}

	var supportCableTestFlagDefault bool

	_ = cmd.PersistentFlags().Bool(supportCableTestFlagName, supportCableTestFlagDefault, supportCableTestDescription)

	return nil
}

func registerCapabilitiesSupportReset(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	supportResetDescription := ``

	var supportResetFlagName string
	if cmdPrefix == "" {
		supportResetFlagName = "supportReset"
	} else {
		supportResetFlagName = fmt.Sprintf("%v.supportReset", cmdPrefix)
	}

	var supportResetFlagDefault bool

	_ = cmd.PersistentFlags().Bool(supportResetFlagName, supportResetFlagDefault, supportResetDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCapabilitiesFlags(depth int, m *models.Capabilities, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, loadBalanceValuesAdded := retrieveCapabilitiesLoadBalanceValuesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || loadBalanceValuesAdded

	err, supportAutoEdgeAdded := retrieveCapabilitiesSupportAutoEdgeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || supportAutoEdgeAdded

	err, supportCableTestAdded := retrieveCapabilitiesSupportCableTestFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || supportCableTestAdded

	err, supportResetAdded := retrieveCapabilitiesSupportResetFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || supportResetAdded

	return nil, retAdded
}

func retrieveCapabilitiesLoadBalanceValuesFlags(depth int, m *models.Capabilities, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	loadBalanceValuesFlagName := fmt.Sprintf("%v.loadBalanceValues", cmdPrefix)
	if cmd.Flags().Changed(loadBalanceValuesFlagName) {
		// warning: loadBalanceValues array type LoadBalanceValues is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveCapabilitiesSupportAutoEdgeFlags(depth int, m *models.Capabilities, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	supportAutoEdgeFlagName := fmt.Sprintf("%v.supportAutoEdge", cmdPrefix)
	if cmd.Flags().Changed(supportAutoEdgeFlagName) {

		var supportAutoEdgeFlagName string
		if cmdPrefix == "" {
			supportAutoEdgeFlagName = "supportAutoEdge"
		} else {
			supportAutoEdgeFlagName = fmt.Sprintf("%v.supportAutoEdge", cmdPrefix)
		}

		supportAutoEdgeFlagValue, err := cmd.Flags().GetBool(supportAutoEdgeFlagName)
		if err != nil {
			return err, false
		}
		m.SupportAutoEdge = supportAutoEdgeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCapabilitiesSupportCableTestFlags(depth int, m *models.Capabilities, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	supportCableTestFlagName := fmt.Sprintf("%v.supportCableTest", cmdPrefix)
	if cmd.Flags().Changed(supportCableTestFlagName) {

		var supportCableTestFlagName string
		if cmdPrefix == "" {
			supportCableTestFlagName = "supportCableTest"
		} else {
			supportCableTestFlagName = fmt.Sprintf("%v.supportCableTest", cmdPrefix)
		}

		supportCableTestFlagValue, err := cmd.Flags().GetBool(supportCableTestFlagName)
		if err != nil {
			return err, false
		}
		m.SupportCableTest = supportCableTestFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCapabilitiesSupportResetFlags(depth int, m *models.Capabilities, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	supportResetFlagName := fmt.Sprintf("%v.supportReset", cmdPrefix)
	if cmd.Flags().Changed(supportResetFlagName) {

		var supportResetFlagName string
		if cmdPrefix == "" {
			supportResetFlagName = "supportReset"
		} else {
			supportResetFlagName = fmt.Sprintf("%v.supportReset", cmdPrefix)
		}

		supportResetFlagValue, err := cmd.Flags().GetBool(supportResetFlagName)
		if err != nil {
			return err, false
		}
		m.SupportReset = supportResetFlagValue

		retAdded = true
	}

	return nil, retAdded
}
