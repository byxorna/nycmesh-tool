// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for Discovery

// register flags to command
func registerModelDiscoveryFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDiscoveryConfigured(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDiscoveryError(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDiscoveryIsProcessing(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDiscoveryProtocol(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDiscoverySnmpCommunity(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDiscoveryStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDiscoveryVisibleBy(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDiscoveryConfigured(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	configuredDescription := ``

	var configuredFlagName string
	if cmdPrefix == "" {
		configuredFlagName = "configured"
	} else {
		configuredFlagName = fmt.Sprintf("%v.configured", cmdPrefix)
	}

	var configuredFlagDefault bool

	_ = cmd.PersistentFlags().Bool(configuredFlagName, configuredFlagDefault, configuredDescription)

	return nil
}

func registerDiscoveryError(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	errorDescription := ``

	var errorFlagName string
	if cmdPrefix == "" {
		errorFlagName = "error"
	} else {
		errorFlagName = fmt.Sprintf("%v.error", cmdPrefix)
	}

	var errorFlagDefault string

	_ = cmd.PersistentFlags().String(errorFlagName, errorFlagDefault, errorDescription)

	return nil
}

func registerDiscoveryIsProcessing(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	isProcessingDescription := ``

	var isProcessingFlagName string
	if cmdPrefix == "" {
		isProcessingFlagName = "isProcessing"
	} else {
		isProcessingFlagName = fmt.Sprintf("%v.isProcessing", cmdPrefix)
	}

	var isProcessingFlagDefault bool

	_ = cmd.PersistentFlags().Bool(isProcessingFlagName, isProcessingFlagDefault, isProcessingDescription)

	return nil
}

func registerDiscoveryProtocol(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	protocolDescription := ``

	var protocolFlagName string
	if cmdPrefix == "" {
		protocolFlagName = "protocol"
	} else {
		protocolFlagName = fmt.Sprintf("%v.protocol", cmdPrefix)
	}

	var protocolFlagDefault string

	_ = cmd.PersistentFlags().String(protocolFlagName, protocolFlagDefault, protocolDescription)

	return nil
}

func registerDiscoverySnmpCommunity(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	snmpCommunityDescription := ``

	var snmpCommunityFlagName string
	if cmdPrefix == "" {
		snmpCommunityFlagName = "snmpCommunity"
	} else {
		snmpCommunityFlagName = fmt.Sprintf("%v.snmpCommunity", cmdPrefix)
	}

	var snmpCommunityFlagDefault string

	_ = cmd.PersistentFlags().String(snmpCommunityFlagName, snmpCommunityFlagDefault, snmpCommunityDescription)

	return nil
}

func registerDiscoveryStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statusDescription := ``

	var statusFlagName string
	if cmdPrefix == "" {
		statusFlagName = "status"
	} else {
		statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
	}

	var statusFlagDefault string

	_ = cmd.PersistentFlags().String(statusFlagName, statusFlagDefault, statusDescription)

	return nil
}

func registerDiscoveryVisibleBy(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var visibleByFlagName string
	if cmdPrefix == "" {
		visibleByFlagName = "visibleBy"
	} else {
		visibleByFlagName = fmt.Sprintf("%v.visibleBy", cmdPrefix)
	}

	if err := registerModelVisibleByFlags(depth+1, visibleByFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDiscoveryFlags(depth int, m *models.Discovery, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, configuredAdded := retrieveDiscoveryConfiguredFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || configuredAdded

	err, errorAdded := retrieveDiscoveryErrorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || errorAdded

	err, isProcessingAdded := retrieveDiscoveryIsProcessingFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || isProcessingAdded

	err, protocolAdded := retrieveDiscoveryProtocolFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || protocolAdded

	err, snmpCommunityAdded := retrieveDiscoverySnmpCommunityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || snmpCommunityAdded

	err, statusAdded := retrieveDiscoveryStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusAdded

	err, visibleByAdded := retrieveDiscoveryVisibleByFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || visibleByAdded

	return nil, retAdded
}

func retrieveDiscoveryConfiguredFlags(depth int, m *models.Discovery, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	configuredFlagName := fmt.Sprintf("%v.configured", cmdPrefix)
	if cmd.Flags().Changed(configuredFlagName) {

		var configuredFlagName string
		if cmdPrefix == "" {
			configuredFlagName = "configured"
		} else {
			configuredFlagName = fmt.Sprintf("%v.configured", cmdPrefix)
		}

		configuredFlagValue, err := cmd.Flags().GetBool(configuredFlagName)
		if err != nil {
			return err, false
		}
		m.Configured = configuredFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDiscoveryErrorFlags(depth int, m *models.Discovery, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	errorFlagName := fmt.Sprintf("%v.error", cmdPrefix)
	if cmd.Flags().Changed(errorFlagName) {

		var errorFlagName string
		if cmdPrefix == "" {
			errorFlagName = "error"
		} else {
			errorFlagName = fmt.Sprintf("%v.error", cmdPrefix)
		}

		errorFlagValue, err := cmd.Flags().GetString(errorFlagName)
		if err != nil {
			return err, false
		}
		m.Error = errorFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDiscoveryIsProcessingFlags(depth int, m *models.Discovery, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	isProcessingFlagName := fmt.Sprintf("%v.isProcessing", cmdPrefix)
	if cmd.Flags().Changed(isProcessingFlagName) {

		var isProcessingFlagName string
		if cmdPrefix == "" {
			isProcessingFlagName = "isProcessing"
		} else {
			isProcessingFlagName = fmt.Sprintf("%v.isProcessing", cmdPrefix)
		}

		isProcessingFlagValue, err := cmd.Flags().GetBool(isProcessingFlagName)
		if err != nil {
			return err, false
		}
		m.IsProcessing = isProcessingFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDiscoveryProtocolFlags(depth int, m *models.Discovery, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	protocolFlagName := fmt.Sprintf("%v.protocol", cmdPrefix)
	if cmd.Flags().Changed(protocolFlagName) {

		var protocolFlagName string
		if cmdPrefix == "" {
			protocolFlagName = "protocol"
		} else {
			protocolFlagName = fmt.Sprintf("%v.protocol", cmdPrefix)
		}

		protocolFlagValue, err := cmd.Flags().GetString(protocolFlagName)
		if err != nil {
			return err, false
		}
		m.Protocol = protocolFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDiscoverySnmpCommunityFlags(depth int, m *models.Discovery, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	snmpCommunityFlagName := fmt.Sprintf("%v.snmpCommunity", cmdPrefix)
	if cmd.Flags().Changed(snmpCommunityFlagName) {

		var snmpCommunityFlagName string
		if cmdPrefix == "" {
			snmpCommunityFlagName = "snmpCommunity"
		} else {
			snmpCommunityFlagName = fmt.Sprintf("%v.snmpCommunity", cmdPrefix)
		}

		snmpCommunityFlagValue, err := cmd.Flags().GetString(snmpCommunityFlagName)
		if err != nil {
			return err, false
		}
		m.SnmpCommunity = snmpCommunityFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDiscoveryStatusFlags(depth int, m *models.Discovery, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statusFlagName := fmt.Sprintf("%v.status", cmdPrefix)
	if cmd.Flags().Changed(statusFlagName) {

		var statusFlagName string
		if cmdPrefix == "" {
			statusFlagName = "status"
		} else {
			statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
		}

		statusFlagValue, err := cmd.Flags().GetString(statusFlagName)
		if err != nil {
			return err, false
		}
		m.Status = statusFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDiscoveryVisibleByFlags(depth int, m *models.Discovery, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	visibleByFlagName := fmt.Sprintf("%v.visibleBy", cmdPrefix)
	if cmd.Flags().Changed(visibleByFlagName) {
		// info: complex object visibleBy VisibleBy is retrieved outside this Changed() block
	}
	visibleByFlagValue := m.VisibleBy
	if swag.IsZero(visibleByFlagValue) {
		visibleByFlagValue = &models.VisibleBy{}
	}

	err, visibleByAdded := retrieveModelVisibleByFlags(depth+1, visibleByFlagValue, visibleByFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || visibleByAdded
	if visibleByAdded {
		m.VisibleBy = visibleByFlagValue
	}

	return nil, retAdded
}
