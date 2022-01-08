// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for DeviceIdentification

// register flags to command
func registerModelDeviceIdentificationFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDeviceIdentificationAuthorized(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceIdentificationCategory(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceIdentificationDisplayName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceIdentificationFirmwareVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceIdentificationHostname(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceIdentificationID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceIdentificationIP(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceIdentificationMac(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceIdentificationModel(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceIdentificationModelName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceIdentificationName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceIdentificationPlatformID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceIdentificationPlatformName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceIdentificationRole(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceIdentificationSerialNumber(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceIdentificationSite(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceIdentificationStarted(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceIdentificationStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceIdentificationType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceIdentificationUpdated(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceIdentificationWanInterfaceID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceIdentificationAuthorized(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	authorizedDescription := `Device is added to UISP.`

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

func registerDeviceIdentificationCategory(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	categoryDescription := `Enum: ["optical","wired","wireless","accessories"]. `

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

func registerDeviceIdentificationDisplayName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	displayNameDescription := `UISP device alias or real name.`

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

func registerDeviceIdentificationFirmwareVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	firmwareVersionDescription := `In SemVer format.`

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

func registerDeviceIdentificationHostname(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	hostnameDescription := ``

	var hostnameFlagName string
	if cmdPrefix == "" {
		hostnameFlagName = "hostname"
	} else {
		hostnameFlagName = fmt.Sprintf("%v.hostname", cmdPrefix)
	}

	var hostnameFlagDefault string

	_ = cmd.PersistentFlags().String(hostnameFlagName, hostnameFlagDefault, hostnameDescription)

	return nil
}

func registerDeviceIdentificationID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `Required. Device ID.`

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

func registerDeviceIdentificationIP(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ipDescription := `Custom IP address in IPv4 or IPv6 format.`

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

func registerDeviceIdentificationMac(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerDeviceIdentificationModel(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	modelDescription := `Enum: ["UF-Nano","UF-Loco","UF-Wifi","UF-Instant","UF-OLT","UF-OLT4","UISP-R-Pro","UISP-R-Lite","UNMS-S-Lite","UISP-S-Lite","UISP-S-Pro","UISP-P-Lite","UISP-LTE","ER-X","ER-X-SFP","ERLite-3","ERPoe-5","ERPro-8","ER-8","ER-8-XG","ER-4","ER-6P","ER-12","ER-12P","ER-10X","EP-R8","EP-R6","EP-S16","ES-12F","ES-16-150W","ES-24-250W","ES-24-500W","ES-24-Lite","ES-48-500W","ES-48-750W","ES-48-Lite","ES-8-150W","ES-16-XG","ES-10XP","ES-10X","ES-18X","ES-26X","EP-54V-150W","EP-24V-72W","EP-54V-72W","TSW-PoE","TSW-PoE PRO","ACB-AC","ACB-ISP","ACB-LOCO","AF11FX","AF24","AF24HD","AF2X","AF3X","AF4X","AF5","AF5U","AF5X","AF-5XHD","AF-LTU","LTU-LITE","AF-LTU5","LTU-Rocket","AFLTULR","AF60","AF60-LR","WaveAP","WaveCPE","GBE-LR","GBE","GBE-Plus","GBE-AP","R2N","R2T","R5N","R6N","R36-GPS","RM3-GPS","R2N-GPS","R5N-GPS","R9N-GPS","R5T-GPS","RM3","R36","R9N","N2N","N5N","N6N","NS3","N36","N9N","N9S","LM2","LM5","B2N","B2T","B5N","B5T","BAC","AG2","AG2-HP","AG5","AG5-HP","p2N","p5N","M25","P2B-400","P5B-300","P5B-300-ISO","P5B-400","P5B-400-ISO","P5B-620","LB5-120","LB5","N5B","N5B-16","N5B-19","N5B-300","N5B-400","N5B-Client","N2B","N2B-13","N2B-400","PAP","LAP-HP","LAP","AGW","AGW-LR","AGW-Pro","AGW-Installer","PB5","PB3","P36","PBM10","NB5","NB2","NB3","B36","NB9","SM5","WM5","IS-M5","Loco5AC","NS-5AC","R5AC-PTMP","R5AC-PTP","R5AC-Lite","R5AC-PRISM","R2AC-Prism","R2AC-Gen2","RP-5AC-Gen2","NBE-2AC-13","NBE-5AC-16","NBE-5AC-19","NBE-5AC-Gen2","PBE-5AC-300","PBE-5AC-300-ISO","PBE-5AC-400","PBE-5AC-400-ISO","PBE-5AC-500","PBE-5AC-500-ISO","PBE-5AC-620","PBE-5AC-620-ISO","PBE-2AC-400","PBE-2AC-400-ISO","PBE-5AC-X-Gen2","PBE-5AC-Gen2","PBE-5AC-ISO-Gen2","PBE-5AC-400-ISO-Gen2","LBE-5AC-16-120","LAP-120","LBE-5AC-23","LBE-5AC-Gen2","LBE-5AC-LR","LAP-GPS","IS-5AC","PS-5AC","SolarSwitch","SolarPoint","BulletAC-IP67","B-DB-AC","UNKNOWN"]. Short names, for example UF-OLT.`

	var modelFlagName string
	if cmdPrefix == "" {
		modelFlagName = "model"
	} else {
		modelFlagName = fmt.Sprintf("%v.model", cmdPrefix)
	}

	var modelFlagDefault string

	_ = cmd.PersistentFlags().String(modelFlagName, modelFlagDefault, modelDescription)

	if err := cmd.RegisterFlagCompletionFunc(modelFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["UF-Nano","UF-Loco","UF-Wifi","UF-Instant","UF-OLT","UF-OLT4","UISP-R-Pro","UISP-R-Lite","UNMS-S-Lite","UISP-S-Lite","UISP-S-Pro","UISP-P-Lite","UISP-LTE","ER-X","ER-X-SFP","ERLite-3","ERPoe-5","ERPro-8","ER-8","ER-8-XG","ER-4","ER-6P","ER-12","ER-12P","ER-10X","EP-R8","EP-R6","EP-S16","ES-12F","ES-16-150W","ES-24-250W","ES-24-500W","ES-24-Lite","ES-48-500W","ES-48-750W","ES-48-Lite","ES-8-150W","ES-16-XG","ES-10XP","ES-10X","ES-18X","ES-26X","EP-54V-150W","EP-24V-72W","EP-54V-72W","TSW-PoE","TSW-PoE PRO","ACB-AC","ACB-ISP","ACB-LOCO","AF11FX","AF24","AF24HD","AF2X","AF3X","AF4X","AF5","AF5U","AF5X","AF-5XHD","AF-LTU","LTU-LITE","AF-LTU5","LTU-Rocket","AFLTULR","AF60","AF60-LR","WaveAP","WaveCPE","GBE-LR","GBE","GBE-Plus","GBE-AP","R2N","R2T","R5N","R6N","R36-GPS","RM3-GPS","R2N-GPS","R5N-GPS","R9N-GPS","R5T-GPS","RM3","R36","R9N","N2N","N5N","N6N","NS3","N36","N9N","N9S","LM2","LM5","B2N","B2T","B5N","B5T","BAC","AG2","AG2-HP","AG5","AG5-HP","p2N","p5N","M25","P2B-400","P5B-300","P5B-300-ISO","P5B-400","P5B-400-ISO","P5B-620","LB5-120","LB5","N5B","N5B-16","N5B-19","N5B-300","N5B-400","N5B-Client","N2B","N2B-13","N2B-400","PAP","LAP-HP","LAP","AGW","AGW-LR","AGW-Pro","AGW-Installer","PB5","PB3","P36","PBM10","NB5","NB2","NB3","B36","NB9","SM5","WM5","IS-M5","Loco5AC","NS-5AC","R5AC-PTMP","R5AC-PTP","R5AC-Lite","R5AC-PRISM","R2AC-Prism","R2AC-Gen2","RP-5AC-Gen2","NBE-2AC-13","NBE-5AC-16","NBE-5AC-19","NBE-5AC-Gen2","PBE-5AC-300","PBE-5AC-300-ISO","PBE-5AC-400","PBE-5AC-400-ISO","PBE-5AC-500","PBE-5AC-500-ISO","PBE-5AC-620","PBE-5AC-620-ISO","PBE-2AC-400","PBE-2AC-400-ISO","PBE-5AC-X-Gen2","PBE-5AC-Gen2","PBE-5AC-ISO-Gen2","PBE-5AC-400-ISO-Gen2","LBE-5AC-16-120","LAP-120","LBE-5AC-23","LBE-5AC-Gen2","LBE-5AC-LR","LAP-GPS","IS-5AC","PS-5AC","SolarSwitch","SolarPoint","BulletAC-IP67","B-DB-AC","UNKNOWN"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerDeviceIdentificationModelName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	modelNameDescription := `Full names, for example UFiber OLT.`

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

func registerDeviceIdentificationName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerDeviceIdentificationPlatformID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	platformIdDescription := `Short name, for example e600.`

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

func registerDeviceIdentificationPlatformName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	platformNameDescription := `Enum: ["UFiber NanoG","UFiber Loco","UFiber Wifi","UFiber Instant","UISPRPro","UISPRLite","UISPS","UISPSPro","UNMSS","UISPLTE","UISPPLite","e50","e100","e200","e300","e600","e1000","u50","u100","u200","u300","u1000","ESWH","ESGH","ES","ESX","EP","EPX","SW","ACB","WA","2WA","XC","2XC","XW","XM","TI","GBE","AirGW","AirGWP","AF","AF02","AF06","AF07","AF08","AF09","af5xhd","afltu","lturocket","GP","ltu60","SunMax","UNKNOWN"]. `

	var platformNameFlagName string
	if cmdPrefix == "" {
		platformNameFlagName = "platformName"
	} else {
		platformNameFlagName = fmt.Sprintf("%v.platformName", cmdPrefix)
	}

	var platformNameFlagDefault string

	_ = cmd.PersistentFlags().String(platformNameFlagName, platformNameFlagDefault, platformNameDescription)

	if err := cmd.RegisterFlagCompletionFunc(platformNameFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["UFiber NanoG","UFiber Loco","UFiber Wifi","UFiber Instant","UISPRPro","UISPRLite","UISPS","UISPSPro","UNMSS","UISPLTE","UISPPLite","e50","e100","e200","e300","e600","e1000","u50","u100","u200","u300","u1000","ESWH","ESGH","ES","ESX","EP","EPX","SW","ACB","WA","2WA","XC","2XC","XW","XM","TI","GBE","AirGW","AirGWP","AF","AF02","AF06","AF07","AF08","AF09","af5xhd","afltu","lturocket","GP","ltu60","SunMax","UNKNOWN"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerDeviceIdentificationRole(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	roleDescription := `Enum: ["router","switch","gpon","ap","station","other","ups","server","wireless","convertor","gateway"]. `

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

func registerDeviceIdentificationSerialNumber(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	serialNumberDescription := ``

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

func registerDeviceIdentificationSite(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var siteFlagName string
	if cmdPrefix == "" {
		siteFlagName = "site"
	} else {
		siteFlagName = fmt.Sprintf("%v.site", cmdPrefix)
	}

	if err := registerModelSiteFlags(depth+1, siteFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceIdentificationStarted(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	startedDescription := ``

	var startedFlagName string
	if cmdPrefix == "" {
		startedFlagName = "started"
	} else {
		startedFlagName = fmt.Sprintf("%v.started", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(startedFlagName, "", startedDescription)

	return nil
}

func registerDeviceIdentificationStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statusDescription := `Enum: ["active","connecting","discovered","inactive","disabled","disconnected","unauthorized","proposed","unknown","unplaced","custom"]. Status of the station.`

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

func registerDeviceIdentificationType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["onu","olt","uispp","uispr","uisps","uispLte","erouter","eswitch","epower","airCube","airMax","airFiber","toughSwitch","solarBeam","wave","blackBox"]. `

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

func registerDeviceIdentificationUpdated(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	updatedDescription := ``

	var updatedFlagName string
	if cmdPrefix == "" {
		updatedFlagName = "updated"
	} else {
		updatedFlagName = fmt.Sprintf("%v.updated", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(updatedFlagName, "", updatedDescription)

	return nil
}

func registerDeviceIdentificationWanInterfaceID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	wanInterfaceIdDescription := ``

	var wanInterfaceIdFlagName string
	if cmdPrefix == "" {
		wanInterfaceIdFlagName = "wanInterfaceId"
	} else {
		wanInterfaceIdFlagName = fmt.Sprintf("%v.wanInterfaceId", cmdPrefix)
	}

	var wanInterfaceIdFlagDefault string

	_ = cmd.PersistentFlags().String(wanInterfaceIdFlagName, wanInterfaceIdFlagDefault, wanInterfaceIdDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDeviceIdentificationFlags(depth int, m *models.DeviceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, authorizedAdded := retrieveDeviceIdentificationAuthorizedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || authorizedAdded

	err, categoryAdded := retrieveDeviceIdentificationCategoryFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || categoryAdded

	err, displayNameAdded := retrieveDeviceIdentificationDisplayNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || displayNameAdded

	err, firmwareVersionAdded := retrieveDeviceIdentificationFirmwareVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || firmwareVersionAdded

	err, hostnameAdded := retrieveDeviceIdentificationHostnameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || hostnameAdded

	err, idAdded := retrieveDeviceIdentificationIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, ipAdded := retrieveDeviceIdentificationIPFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipAdded

	err, macAdded := retrieveDeviceIdentificationMacFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || macAdded

	err, modelAdded := retrieveDeviceIdentificationModelFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || modelAdded

	err, modelNameAdded := retrieveDeviceIdentificationModelNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || modelNameAdded

	err, nameAdded := retrieveDeviceIdentificationNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, platformIdAdded := retrieveDeviceIdentificationPlatformIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || platformIdAdded

	err, platformNameAdded := retrieveDeviceIdentificationPlatformNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || platformNameAdded

	err, roleAdded := retrieveDeviceIdentificationRoleFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || roleAdded

	err, serialNumberAdded := retrieveDeviceIdentificationSerialNumberFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || serialNumberAdded

	err, siteAdded := retrieveDeviceIdentificationSiteFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || siteAdded

	err, startedAdded := retrieveDeviceIdentificationStartedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || startedAdded

	err, statusAdded := retrieveDeviceIdentificationStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusAdded

	err, typeAdded := retrieveDeviceIdentificationTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	err, updatedAdded := retrieveDeviceIdentificationUpdatedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || updatedAdded

	err, wanInterfaceIdAdded := retrieveDeviceIdentificationWanInterfaceIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || wanInterfaceIdAdded

	return nil, retAdded
}

func retrieveDeviceIdentificationAuthorizedFlags(depth int, m *models.DeviceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Authorized = authorizedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceIdentificationCategoryFlags(depth int, m *models.DeviceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDeviceIdentificationDisplayNameFlags(depth int, m *models.DeviceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.DisplayName = displayNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceIdentificationFirmwareVersionFlags(depth int, m *models.DeviceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDeviceIdentificationHostnameFlags(depth int, m *models.DeviceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	hostnameFlagName := fmt.Sprintf("%v.hostname", cmdPrefix)
	if cmd.Flags().Changed(hostnameFlagName) {

		var hostnameFlagName string
		if cmdPrefix == "" {
			hostnameFlagName = "hostname"
		} else {
			hostnameFlagName = fmt.Sprintf("%v.hostname", cmdPrefix)
		}

		hostnameFlagValue, err := cmd.Flags().GetString(hostnameFlagName)
		if err != nil {
			return err, false
		}
		m.Hostname = hostnameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceIdentificationIDFlags(depth int, m *models.DeviceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDeviceIdentificationIPFlags(depth int, m *models.DeviceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.IP = ipFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceIdentificationMacFlags(depth int, m *models.DeviceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDeviceIdentificationModelFlags(depth int, m *models.DeviceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDeviceIdentificationModelNameFlags(depth int, m *models.DeviceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.ModelName = modelNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceIdentificationNameFlags(depth int, m *models.DeviceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDeviceIdentificationPlatformIDFlags(depth int, m *models.DeviceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDeviceIdentificationPlatformNameFlags(depth int, m *models.DeviceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.PlatformName = platformNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceIdentificationRoleFlags(depth int, m *models.DeviceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Role = roleFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceIdentificationSerialNumberFlags(depth int, m *models.DeviceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.SerialNumber = serialNumberFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceIdentificationSiteFlags(depth int, m *models.DeviceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	siteFlagName := fmt.Sprintf("%v.site", cmdPrefix)
	if cmd.Flags().Changed(siteFlagName) {
		// info: complex object site Site is retrieved outside this Changed() block
	}
	siteFlagValue := m.Site
	if swag.IsZero(siteFlagValue) {
		siteFlagValue = &models.Site{}
	}

	err, siteAdded := retrieveModelSiteFlags(depth+1, siteFlagValue, siteFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || siteAdded
	if siteAdded {
		m.Site = siteFlagValue
	}

	return nil, retAdded
}

func retrieveDeviceIdentificationStartedFlags(depth int, m *models.DeviceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	startedFlagName := fmt.Sprintf("%v.started", cmdPrefix)
	if cmd.Flags().Changed(startedFlagName) {

		var startedFlagName string
		if cmdPrefix == "" {
			startedFlagName = "started"
		} else {
			startedFlagName = fmt.Sprintf("%v.started", cmdPrefix)
		}

		startedFlagValueStr, err := cmd.Flags().GetString(startedFlagName)
		if err != nil {
			return err, false
		}
		var startedFlagValue strfmt.DateTime
		if err := startedFlagValue.UnmarshalText([]byte(startedFlagValueStr)); err != nil {
			return err, false
		}
		m.Started = startedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceIdentificationStatusFlags(depth int, m *models.DeviceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDeviceIdentificationTypeFlags(depth int, m *models.DeviceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDeviceIdentificationUpdatedFlags(depth int, m *models.DeviceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	updatedFlagName := fmt.Sprintf("%v.updated", cmdPrefix)
	if cmd.Flags().Changed(updatedFlagName) {

		var updatedFlagName string
		if cmdPrefix == "" {
			updatedFlagName = "updated"
		} else {
			updatedFlagName = fmt.Sprintf("%v.updated", cmdPrefix)
		}

		updatedFlagValueStr, err := cmd.Flags().GetString(updatedFlagName)
		if err != nil {
			return err, false
		}
		var updatedFlagValue strfmt.DateTime
		if err := updatedFlagValue.UnmarshalText([]byte(updatedFlagValueStr)); err != nil {
			return err, false
		}
		m.Updated = updatedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceIdentificationWanInterfaceIDFlags(depth int, m *models.DeviceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	wanInterfaceIdFlagName := fmt.Sprintf("%v.wanInterfaceId", cmdPrefix)
	if cmd.Flags().Changed(wanInterfaceIdFlagName) {

		var wanInterfaceIdFlagName string
		if cmdPrefix == "" {
			wanInterfaceIdFlagName = "wanInterfaceId"
		} else {
			wanInterfaceIdFlagName = fmt.Sprintf("%v.wanInterfaceId", cmdPrefix)
		}

		wanInterfaceIdFlagValue, err := cmd.Flags().GetString(wanInterfaceIdFlagName)
		if err != nil {
			return err, false
		}
		m.WanInterfaceID = wanInterfaceIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}
