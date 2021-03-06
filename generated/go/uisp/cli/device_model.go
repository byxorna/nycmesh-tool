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

// Schema cli for Device

// register flags to command
func registerModelDeviceFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDeviceCategory(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceFirmwareVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceIPAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceMac(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceModel(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDevicePlatformID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceCategory(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	categoryDescription := `Enum: ["optical","wired","wireless","accessories"]. Device category`

	var categoryFlagName string
	if cmdPrefix == "" {
		categoryFlagName = "category"
	} else {
		categoryFlagName = fmt.Sprintf("%v.category", cmdPrefix)
	}

	var categoryFlagDefault string

	_ = cmd.PersistentFlags().String(categoryFlagName, categoryFlagDefault, categoryDescription)

	if err := cmd.RegisterFlagCompletionFunc(categoryFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["optical","wired","wireless","accessories"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerDeviceFirmwareVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	firmwareVersionDescription := `Device firmware version`

	var firmwareVersionFlagName string
	if cmdPrefix == "" {
		firmwareVersionFlagName = "firmwareVersion"
	} else {
		firmwareVersionFlagName = fmt.Sprintf("%v.firmwareVersion", cmdPrefix)
	}

	var firmwareVersionFlagDefault string

	_ = cmd.PersistentFlags().String(firmwareVersionFlagName, firmwareVersionFlagDefault, firmwareVersionDescription)

	return nil
}

func registerDeviceID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `Required. Id of gateway device`

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

func registerDeviceIPAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ipAddressDescription := `Device IP address in CIDR format`

	var ipAddressFlagName string
	if cmdPrefix == "" {
		ipAddressFlagName = "ipAddress"
	} else {
		ipAddressFlagName = fmt.Sprintf("%v.ipAddress", cmdPrefix)
	}

	var ipAddressFlagDefault string

	_ = cmd.PersistentFlags().String(ipAddressFlagName, ipAddressFlagDefault, ipAddressDescription)

	return nil
}

func registerDeviceMac(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	macDescription := `Required. Device MAC address`

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

func registerDeviceModel(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	modelDescription := `Device model`

	var modelFlagName string
	if cmdPrefix == "" {
		modelFlagName = "model"
	} else {
		modelFlagName = fmt.Sprintf("%v.model", cmdPrefix)
	}

	var modelFlagDefault string

	_ = cmd.PersistentFlags().String(modelFlagName, modelFlagDefault, modelDescription)

	return nil
}

func registerDeviceName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Device name`

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

func registerDevicePlatformID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	platformIdDescription := `Device platform id`

	var platformIdFlagName string
	if cmdPrefix == "" {
		platformIdFlagName = "platformId"
	} else {
		platformIdFlagName = fmt.Sprintf("%v.platformId", cmdPrefix)
	}

	var platformIdFlagDefault string

	_ = cmd.PersistentFlags().String(platformIdFlagName, platformIdFlagDefault, platformIdDescription)

	return nil
}

func registerDeviceStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statusDescription := `Enum: ["active","connecting","discovered","inactive","disabled","disconnected","unauthorized","proposed","unknown","unplaced","custom"]. Required. Device status`

	var statusFlagName string
	if cmdPrefix == "" {
		statusFlagName = "status"
	} else {
		statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
	}

	var statusFlagDefault string

	_ = cmd.PersistentFlags().String(statusFlagName, statusFlagDefault, statusDescription)

	if err := cmd.RegisterFlagCompletionFunc(statusFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["active","connecting","discovered","inactive","disabled","disconnected","unauthorized","proposed","unknown","unplaced","custom"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerDeviceType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Device type`

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDeviceFlags(depth int, m *models.Device, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, categoryAdded := retrieveDeviceCategoryFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || categoryAdded

	err, firmwareVersionAdded := retrieveDeviceFirmwareVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || firmwareVersionAdded

	err, idAdded := retrieveDeviceIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, ipAddressAdded := retrieveDeviceIPAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipAddressAdded

	err, macAdded := retrieveDeviceMacFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || macAdded

	err, modelAdded := retrieveDeviceModelFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || modelAdded

	err, nameAdded := retrieveDeviceNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, platformIdAdded := retrieveDevicePlatformIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || platformIdAdded

	err, statusAdded := retrieveDeviceStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusAdded

	err, typeAdded := retrieveDeviceTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveDeviceCategoryFlags(depth int, m *models.Device, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	categoryFlagName := fmt.Sprintf("%v.category", cmdPrefix)
	if cmd.Flags().Changed(categoryFlagName) {

		var categoryFlagName string
		if cmdPrefix == "" {
			categoryFlagName = "category"
		} else {
			categoryFlagName = fmt.Sprintf("%v.category", cmdPrefix)
		}

		categoryFlagValue, err := cmd.Flags().GetString(categoryFlagName)
		if err != nil {
			return err, false
		}
		m.Category = categoryFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceFirmwareVersionFlags(depth int, m *models.Device, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	firmwareVersionFlagName := fmt.Sprintf("%v.firmwareVersion", cmdPrefix)
	if cmd.Flags().Changed(firmwareVersionFlagName) {

		var firmwareVersionFlagName string
		if cmdPrefix == "" {
			firmwareVersionFlagName = "firmwareVersion"
		} else {
			firmwareVersionFlagName = fmt.Sprintf("%v.firmwareVersion", cmdPrefix)
		}

		firmwareVersionFlagValue, err := cmd.Flags().GetString(firmwareVersionFlagName)
		if err != nil {
			return err, false
		}
		m.FirmwareVersion = firmwareVersionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceIDFlags(depth int, m *models.Device, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.ID = &idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceIPAddressFlags(depth int, m *models.Device, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ipAddressFlagName := fmt.Sprintf("%v.ipAddress", cmdPrefix)
	if cmd.Flags().Changed(ipAddressFlagName) {

		var ipAddressFlagName string
		if cmdPrefix == "" {
			ipAddressFlagName = "ipAddress"
		} else {
			ipAddressFlagName = fmt.Sprintf("%v.ipAddress", cmdPrefix)
		}

		ipAddressFlagValue, err := cmd.Flags().GetString(ipAddressFlagName)
		if err != nil {
			return err, false
		}
		m.IPAddress = ipAddressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceMacFlags(depth int, m *models.Device, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Mac = &macFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceModelFlags(depth int, m *models.Device, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	modelFlagName := fmt.Sprintf("%v.model", cmdPrefix)
	if cmd.Flags().Changed(modelFlagName) {

		var modelFlagName string
		if cmdPrefix == "" {
			modelFlagName = "model"
		} else {
			modelFlagName = fmt.Sprintf("%v.model", cmdPrefix)
		}

		modelFlagValue, err := cmd.Flags().GetString(modelFlagName)
		if err != nil {
			return err, false
		}
		m.Model = modelFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceNameFlags(depth int, m *models.Device, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDevicePlatformIDFlags(depth int, m *models.Device, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	platformIdFlagName := fmt.Sprintf("%v.platformId", cmdPrefix)
	if cmd.Flags().Changed(platformIdFlagName) {

		var platformIdFlagName string
		if cmdPrefix == "" {
			platformIdFlagName = "platformId"
		} else {
			platformIdFlagName = fmt.Sprintf("%v.platformId", cmdPrefix)
		}

		platformIdFlagValue, err := cmd.Flags().GetString(platformIdFlagName)
		if err != nil {
			return err, false
		}
		m.PlatformID = platformIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceStatusFlags(depth int, m *models.Device, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Status = &statusFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceTypeFlags(depth int, m *models.Device, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
