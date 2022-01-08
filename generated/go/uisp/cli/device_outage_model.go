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

// Schema cli for DeviceOutage

// register flags to command
func registerModelDeviceOutageFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDeviceOutageAuthorized(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOutageCategory(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOutageDisplayName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOutageFirmwareVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOutageID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOutageIP(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOutageMac(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOutageModel(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOutageModelName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOutageName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOutagePlatformID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOutagePlatformName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOutageRole(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOutageSerialNumber(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOutageSite(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOutageType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceOutageAuthorized(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	authorizedDescription := `Required. Device is added to UISP`

	var authorizedFlagName string
	if cmdPrefix == "" {
		authorizedFlagName = "authorized"
	} else {
		authorizedFlagName = fmt.Sprintf("%v.authorized", cmdPrefix)
	}

	var authorizedFlagDefault bool

	_ = cmd.PersistentFlags().Bool(authorizedFlagName, authorizedFlagDefault, authorizedDescription)

	return nil
}

func registerDeviceOutageCategory(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	categoryDescription := `Enum: ["optical","wired","wireless","accessories"]. Required. `

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

func registerDeviceOutageDisplayName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	displayNameDescription := `Required. `

	var displayNameFlagName string
	if cmdPrefix == "" {
		displayNameFlagName = "displayName"
	} else {
		displayNameFlagName = fmt.Sprintf("%v.displayName", cmdPrefix)
	}

	var displayNameFlagDefault string

	_ = cmd.PersistentFlags().String(displayNameFlagName, displayNameFlagDefault, displayNameDescription)

	return nil
}

func registerDeviceOutageFirmwareVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	firmwareVersionDescription := `Required. `

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

func registerDeviceOutageID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `Required. Technical ID only for UISP`

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

func registerDeviceOutageIP(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ipDescription := `Required. `

	var ipFlagName string
	if cmdPrefix == "" {
		ipFlagName = "ip"
	} else {
		ipFlagName = fmt.Sprintf("%v.ip", cmdPrefix)
	}

	var ipFlagDefault string

	_ = cmd.PersistentFlags().String(ipFlagName, ipFlagDefault, ipDescription)

	return nil
}

func registerDeviceOutageMac(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	macDescription := `Required. Available only for clients (ONU)`

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

func registerDeviceOutageModel(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	modelDescription := `Required. `

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

func registerDeviceOutageModelName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	modelNameDescription := `Required. `

	var modelNameFlagName string
	if cmdPrefix == "" {
		modelNameFlagName = "modelName"
	} else {
		modelNameFlagName = fmt.Sprintf("%v.modelName", cmdPrefix)
	}

	var modelNameFlagDefault string

	_ = cmd.PersistentFlags().String(modelNameFlagName, modelNameFlagDefault, modelNameDescription)

	return nil
}

func registerDeviceOutageName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerDeviceOutagePlatformID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	platformIdDescription := `Required. `

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

func registerDeviceOutagePlatformName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	platformNameDescription := `Required. `

	var platformNameFlagName string
	if cmdPrefix == "" {
		platformNameFlagName = "platformName"
	} else {
		platformNameFlagName = fmt.Sprintf("%v.platformName", cmdPrefix)
	}

	var platformNameFlagDefault string

	_ = cmd.PersistentFlags().String(platformNameFlagName, platformNameFlagDefault, platformNameDescription)

	return nil
}

func registerDeviceOutageRole(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	roleDescription := `Enum: ["router","switch","gpon","ap","station","other","ups","server","wireless","convertor","gateway"]. Required. `

	var roleFlagName string
	if cmdPrefix == "" {
		roleFlagName = "role"
	} else {
		roleFlagName = fmt.Sprintf("%v.role", cmdPrefix)
	}

	var roleFlagDefault string

	_ = cmd.PersistentFlags().String(roleFlagName, roleFlagDefault, roleDescription)

	if err := cmd.RegisterFlagCompletionFunc(roleFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["router","switch","gpon","ap","station","other","ups","server","wireless","convertor","gateway"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerDeviceOutageSerialNumber(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	serialNumberDescription := `Required. `

	var serialNumberFlagName string
	if cmdPrefix == "" {
		serialNumberFlagName = "serialNumber"
	} else {
		serialNumberFlagName = fmt.Sprintf("%v.serialNumber", cmdPrefix)
	}

	var serialNumberFlagDefault string

	_ = cmd.PersistentFlags().String(serialNumberFlagName, serialNumberFlagDefault, serialNumberDescription)

	return nil
}

func registerDeviceOutageSite(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var siteFlagName string
	if cmdPrefix == "" {
		siteFlagName = "site"
	} else {
		siteFlagName = fmt.Sprintf("%v.site", cmdPrefix)
	}

	if err := registerModelSite2Flags(depth+1, siteFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceOutageType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["onu","olt","uispp","uispr","uisps","uispLte","erouter","eswitch","epower","airCube","airMax","airFiber","toughSwitch","solarBeam","wave","blackBox"]. Required. `

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
			if err := json.Unmarshal([]byte(`["onu","olt","uispp","uispr","uisps","uispLte","erouter","eswitch","epower","airCube","airMax","airFiber","toughSwitch","solarBeam","wave","blackBox"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDeviceOutageFlags(depth int, m *models.DeviceOutage, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, authorizedAdded := retrieveDeviceOutageAuthorizedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || authorizedAdded

	err, categoryAdded := retrieveDeviceOutageCategoryFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || categoryAdded

	err, displayNameAdded := retrieveDeviceOutageDisplayNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || displayNameAdded

	err, firmwareVersionAdded := retrieveDeviceOutageFirmwareVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || firmwareVersionAdded

	err, idAdded := retrieveDeviceOutageIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, ipAdded := retrieveDeviceOutageIPFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipAdded

	err, macAdded := retrieveDeviceOutageMacFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || macAdded

	err, modelAdded := retrieveDeviceOutageModelFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || modelAdded

	err, modelNameAdded := retrieveDeviceOutageModelNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || modelNameAdded

	err, nameAdded := retrieveDeviceOutageNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, platformIdAdded := retrieveDeviceOutagePlatformIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || platformIdAdded

	err, platformNameAdded := retrieveDeviceOutagePlatformNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || platformNameAdded

	err, roleAdded := retrieveDeviceOutageRoleFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || roleAdded

	err, serialNumberAdded := retrieveDeviceOutageSerialNumberFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || serialNumberAdded

	err, siteAdded := retrieveDeviceOutageSiteFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || siteAdded

	err, typeAdded := retrieveDeviceOutageTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveDeviceOutageAuthorizedFlags(depth int, m *models.DeviceOutage, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	authorizedFlagName := fmt.Sprintf("%v.authorized", cmdPrefix)
	if cmd.Flags().Changed(authorizedFlagName) {

		var authorizedFlagName string
		if cmdPrefix == "" {
			authorizedFlagName = "authorized"
		} else {
			authorizedFlagName = fmt.Sprintf("%v.authorized", cmdPrefix)
		}

		authorizedFlagValue, err := cmd.Flags().GetBool(authorizedFlagName)
		if err != nil {
			return err, false
		}
		m.Authorized = &authorizedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceOutageCategoryFlags(depth int, m *models.DeviceOutage, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Category = &categoryFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceOutageDisplayNameFlags(depth int, m *models.DeviceOutage, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	displayNameFlagName := fmt.Sprintf("%v.displayName", cmdPrefix)
	if cmd.Flags().Changed(displayNameFlagName) {

		var displayNameFlagName string
		if cmdPrefix == "" {
			displayNameFlagName = "displayName"
		} else {
			displayNameFlagName = fmt.Sprintf("%v.displayName", cmdPrefix)
		}

		displayNameFlagValue, err := cmd.Flags().GetString(displayNameFlagName)
		if err != nil {
			return err, false
		}
		m.DisplayName = &displayNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceOutageFirmwareVersionFlags(depth int, m *models.DeviceOutage, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.FirmwareVersion = &firmwareVersionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceOutageIDFlags(depth int, m *models.DeviceOutage, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDeviceOutageIPFlags(depth int, m *models.DeviceOutage, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ipFlagName := fmt.Sprintf("%v.ip", cmdPrefix)
	if cmd.Flags().Changed(ipFlagName) {

		var ipFlagName string
		if cmdPrefix == "" {
			ipFlagName = "ip"
		} else {
			ipFlagName = fmt.Sprintf("%v.ip", cmdPrefix)
		}

		ipFlagValue, err := cmd.Flags().GetString(ipFlagName)
		if err != nil {
			return err, false
		}
		m.IP = &ipFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceOutageMacFlags(depth int, m *models.DeviceOutage, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDeviceOutageModelFlags(depth int, m *models.DeviceOutage, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Model = &modelFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceOutageModelNameFlags(depth int, m *models.DeviceOutage, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	modelNameFlagName := fmt.Sprintf("%v.modelName", cmdPrefix)
	if cmd.Flags().Changed(modelNameFlagName) {

		var modelNameFlagName string
		if cmdPrefix == "" {
			modelNameFlagName = "modelName"
		} else {
			modelNameFlagName = fmt.Sprintf("%v.modelName", cmdPrefix)
		}

		modelNameFlagValue, err := cmd.Flags().GetString(modelNameFlagName)
		if err != nil {
			return err, false
		}
		m.ModelName = &modelNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceOutageNameFlags(depth int, m *models.DeviceOutage, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDeviceOutagePlatformIDFlags(depth int, m *models.DeviceOutage, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.PlatformID = &platformIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceOutagePlatformNameFlags(depth int, m *models.DeviceOutage, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	platformNameFlagName := fmt.Sprintf("%v.platformName", cmdPrefix)
	if cmd.Flags().Changed(platformNameFlagName) {

		var platformNameFlagName string
		if cmdPrefix == "" {
			platformNameFlagName = "platformName"
		} else {
			platformNameFlagName = fmt.Sprintf("%v.platformName", cmdPrefix)
		}

		platformNameFlagValue, err := cmd.Flags().GetString(platformNameFlagName)
		if err != nil {
			return err, false
		}
		m.PlatformName = &platformNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceOutageRoleFlags(depth int, m *models.DeviceOutage, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	roleFlagName := fmt.Sprintf("%v.role", cmdPrefix)
	if cmd.Flags().Changed(roleFlagName) {

		var roleFlagName string
		if cmdPrefix == "" {
			roleFlagName = "role"
		} else {
			roleFlagName = fmt.Sprintf("%v.role", cmdPrefix)
		}

		roleFlagValue, err := cmd.Flags().GetString(roleFlagName)
		if err != nil {
			return err, false
		}
		m.Role = &roleFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceOutageSerialNumberFlags(depth int, m *models.DeviceOutage, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	serialNumberFlagName := fmt.Sprintf("%v.serialNumber", cmdPrefix)
	if cmd.Flags().Changed(serialNumberFlagName) {

		var serialNumberFlagName string
		if cmdPrefix == "" {
			serialNumberFlagName = "serialNumber"
		} else {
			serialNumberFlagName = fmt.Sprintf("%v.serialNumber", cmdPrefix)
		}

		serialNumberFlagValue, err := cmd.Flags().GetString(serialNumberFlagName)
		if err != nil {
			return err, false
		}
		m.SerialNumber = &serialNumberFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceOutageSiteFlags(depth int, m *models.DeviceOutage, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	siteFlagName := fmt.Sprintf("%v.site", cmdPrefix)
	if cmd.Flags().Changed(siteFlagName) {
		// info: complex object site Site2 is retrieved outside this Changed() block
	}
	siteFlagValue := m.Site
	if swag.IsZero(siteFlagValue) {
		siteFlagValue = &models.Site2{}
	}

	err, siteAdded := retrieveModelSite2Flags(depth+1, siteFlagValue, siteFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || siteAdded
	if siteAdded {
		m.Site = siteFlagValue
	}

	return nil, retAdded
}

func retrieveDeviceOutageTypeFlags(depth int, m *models.DeviceOutage, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
