// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

// Schema cli for DeviceMeta

// register flags to command
func registerModelDeviceMetaFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDeviceMetaAlias(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceMetaCustomIPAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceMetaFailedMessageDecryption(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceMetaMaintenance(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceMetaNote(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceMetaRestartTimestamp(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceMetaAlias(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	aliasDescription := ``

	var aliasFlagName string
	if cmdPrefix == "" {
		aliasFlagName = "alias"
	} else {
		aliasFlagName = fmt.Sprintf("%v.alias", cmdPrefix)
	}

	var aliasFlagDefault string

	_ = cmd.PersistentFlags().String(aliasFlagName, aliasFlagDefault, aliasDescription)

	return nil
}

func registerDeviceMetaCustomIPAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	customIpAddressDescription := `Custom IP address in IPv4 or IPv6 format.`

	var customIpAddressFlagName string
	if cmdPrefix == "" {
		customIpAddressFlagName = "customIpAddress"
	} else {
		customIpAddressFlagName = fmt.Sprintf("%v.customIpAddress", cmdPrefix)
	}

	var customIpAddressFlagDefault string

	_ = cmd.PersistentFlags().String(customIpAddressFlagName, customIpAddressFlagDefault, customIpAddressDescription)

	return nil
}

func registerDeviceMetaFailedMessageDecryption(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	failedMessageDecryptionDescription := `Required. `

	var failedMessageDecryptionFlagName string
	if cmdPrefix == "" {
		failedMessageDecryptionFlagName = "failedMessageDecryption"
	} else {
		failedMessageDecryptionFlagName = fmt.Sprintf("%v.failedMessageDecryption", cmdPrefix)
	}

	var failedMessageDecryptionFlagDefault bool

	_ = cmd.PersistentFlags().Bool(failedMessageDecryptionFlagName, failedMessageDecryptionFlagDefault, failedMessageDecryptionDescription)

	return nil
}

func registerDeviceMetaMaintenance(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	maintenanceDescription := `Required. `

	var maintenanceFlagName string
	if cmdPrefix == "" {
		maintenanceFlagName = "maintenance"
	} else {
		maintenanceFlagName = fmt.Sprintf("%v.maintenance", cmdPrefix)
	}

	var maintenanceFlagDefault bool

	_ = cmd.PersistentFlags().Bool(maintenanceFlagName, maintenanceFlagDefault, maintenanceDescription)

	return nil
}

func registerDeviceMetaNote(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	noteDescription := ``

	var noteFlagName string
	if cmdPrefix == "" {
		noteFlagName = "note"
	} else {
		noteFlagName = fmt.Sprintf("%v.note", cmdPrefix)
	}

	var noteFlagDefault string

	_ = cmd.PersistentFlags().String(noteFlagName, noteFlagDefault, noteDescription)

	return nil
}

func registerDeviceMetaRestartTimestamp(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	restartTimestampDescription := `Required. `

	var restartTimestampFlagName string
	if cmdPrefix == "" {
		restartTimestampFlagName = "restartTimestamp"
	} else {
		restartTimestampFlagName = fmt.Sprintf("%v.restartTimestamp", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(restartTimestampFlagName, "", restartTimestampDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDeviceMetaFlags(depth int, m *models.DeviceMeta, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, aliasAdded := retrieveDeviceMetaAliasFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || aliasAdded

	err, customIpAddressAdded := retrieveDeviceMetaCustomIPAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || customIpAddressAdded

	err, failedMessageDecryptionAdded := retrieveDeviceMetaFailedMessageDecryptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || failedMessageDecryptionAdded

	err, maintenanceAdded := retrieveDeviceMetaMaintenanceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || maintenanceAdded

	err, noteAdded := retrieveDeviceMetaNoteFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || noteAdded

	err, restartTimestampAdded := retrieveDeviceMetaRestartTimestampFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || restartTimestampAdded

	return nil, retAdded
}

func retrieveDeviceMetaAliasFlags(depth int, m *models.DeviceMeta, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	aliasFlagName := fmt.Sprintf("%v.alias", cmdPrefix)
	if cmd.Flags().Changed(aliasFlagName) {

		var aliasFlagName string
		if cmdPrefix == "" {
			aliasFlagName = "alias"
		} else {
			aliasFlagName = fmt.Sprintf("%v.alias", cmdPrefix)
		}

		aliasFlagValue, err := cmd.Flags().GetString(aliasFlagName)
		if err != nil {
			return err, false
		}
		m.Alias = aliasFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceMetaCustomIPAddressFlags(depth int, m *models.DeviceMeta, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	customIpAddressFlagName := fmt.Sprintf("%v.customIpAddress", cmdPrefix)
	if cmd.Flags().Changed(customIpAddressFlagName) {

		var customIpAddressFlagName string
		if cmdPrefix == "" {
			customIpAddressFlagName = "customIpAddress"
		} else {
			customIpAddressFlagName = fmt.Sprintf("%v.customIpAddress", cmdPrefix)
		}

		customIpAddressFlagValue, err := cmd.Flags().GetString(customIpAddressFlagName)
		if err != nil {
			return err, false
		}
		m.CustomIPAddress = customIpAddressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceMetaFailedMessageDecryptionFlags(depth int, m *models.DeviceMeta, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	failedMessageDecryptionFlagName := fmt.Sprintf("%v.failedMessageDecryption", cmdPrefix)
	if cmd.Flags().Changed(failedMessageDecryptionFlagName) {

		var failedMessageDecryptionFlagName string
		if cmdPrefix == "" {
			failedMessageDecryptionFlagName = "failedMessageDecryption"
		} else {
			failedMessageDecryptionFlagName = fmt.Sprintf("%v.failedMessageDecryption", cmdPrefix)
		}

		failedMessageDecryptionFlagValue, err := cmd.Flags().GetBool(failedMessageDecryptionFlagName)
		if err != nil {
			return err, false
		}
		m.FailedMessageDecryption = &failedMessageDecryptionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceMetaMaintenanceFlags(depth int, m *models.DeviceMeta, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	maintenanceFlagName := fmt.Sprintf("%v.maintenance", cmdPrefix)
	if cmd.Flags().Changed(maintenanceFlagName) {

		var maintenanceFlagName string
		if cmdPrefix == "" {
			maintenanceFlagName = "maintenance"
		} else {
			maintenanceFlagName = fmt.Sprintf("%v.maintenance", cmdPrefix)
		}

		maintenanceFlagValue, err := cmd.Flags().GetBool(maintenanceFlagName)
		if err != nil {
			return err, false
		}
		m.Maintenance = &maintenanceFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceMetaNoteFlags(depth int, m *models.DeviceMeta, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	noteFlagName := fmt.Sprintf("%v.note", cmdPrefix)
	if cmd.Flags().Changed(noteFlagName) {

		var noteFlagName string
		if cmdPrefix == "" {
			noteFlagName = "note"
		} else {
			noteFlagName = fmt.Sprintf("%v.note", cmdPrefix)
		}

		noteFlagValue, err := cmd.Flags().GetString(noteFlagName)
		if err != nil {
			return err, false
		}
		m.Note = noteFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceMetaRestartTimestampFlags(depth int, m *models.DeviceMeta, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	restartTimestampFlagName := fmt.Sprintf("%v.restartTimestamp", cmdPrefix)
	if cmd.Flags().Changed(restartTimestampFlagName) {

		var restartTimestampFlagName string
		if cmdPrefix == "" {
			restartTimestampFlagName = "restartTimestamp"
		} else {
			restartTimestampFlagName = fmt.Sprintf("%v.restartTimestamp", cmdPrefix)
		}

		restartTimestampFlagValueStr, err := cmd.Flags().GetString(restartTimestampFlagName)
		if err != nil {
			return err, false
		}
		var restartTimestampFlagValue strfmt.DateTime
		if err := restartTimestampFlagValue.UnmarshalText([]byte(restartTimestampFlagValueStr)); err != nil {
			return err, false
		}
		m.RestartTimestamp = &restartTimestampFlagValue

		retAdded = true
	}

	return nil, retAdded
}
