// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for UispSettingMeta

// register flags to command
func registerModelUispSettingMetaFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUispSettingMetaAlias(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUispSettingMetaCustomIPAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUispSettingMetaMaintenance(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUispSettingMetaNote(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUispSettingMetaAlias(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	aliasDescription := `Display name used in UISP for the device.`

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

func registerUispSettingMetaCustomIPAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerUispSettingMetaMaintenance(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	maintenanceDescription := `If set to true outages are not reported for the device.`

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

func registerUispSettingMetaNote(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	noteDescription := `Custom device description.`

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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUispSettingMetaFlags(depth int, m *models.UispSettingMeta, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, aliasAdded := retrieveUispSettingMetaAliasFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || aliasAdded

	err, customIpAddressAdded := retrieveUispSettingMetaCustomIPAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || customIpAddressAdded

	err, maintenanceAdded := retrieveUispSettingMetaMaintenanceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || maintenanceAdded

	err, noteAdded := retrieveUispSettingMetaNoteFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || noteAdded

	return nil, retAdded
}

func retrieveUispSettingMetaAliasFlags(depth int, m *models.UispSettingMeta, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveUispSettingMetaCustomIPAddressFlags(depth int, m *models.UispSettingMeta, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveUispSettingMetaMaintenanceFlags(depth int, m *models.UispSettingMeta, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Maintenance = maintenanceFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUispSettingMetaNoteFlags(depth int, m *models.UispSettingMeta, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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