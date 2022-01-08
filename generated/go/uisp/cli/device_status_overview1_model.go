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

// Schema cli for DeviceStatusOverview1

// register flags to command
func registerModelDeviceStatusOverview1Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDeviceStatusOverview1Airfiber(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatusOverview1Airmax(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatusOverview1Attributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatusOverview1Discovery(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatusOverview1Enabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatusOverview1Eswitch(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatusOverview1Firmware(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatusOverview1Identification(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatusOverview1Interfaces(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatusOverview1IPAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatusOverview1LatestBackup(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatusOverview1Location(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatusOverview1Meta(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatusOverview1Mode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatusOverview1Overview(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatusOverview1Uisps(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatusOverview1Upgrade(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceStatusOverview1Airfiber(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var airfiberFlagName string
	if cmdPrefix == "" {
		airfiberFlagName = "airfiber"
	} else {
		airfiberFlagName = fmt.Sprintf("%v.airfiber", cmdPrefix)
	}

	if err := registerModelDEPRECATEDDeviceAirFiberPartiallyMovedPropertiesToInterfaceOrStationFlags(depth+1, airfiberFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceStatusOverview1Airmax(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var airmaxFlagName string
	if cmdPrefix == "" {
		airmaxFlagName = "airmax"
	} else {
		airmaxFlagName = fmt.Sprintf("%v.airmax", cmdPrefix)
	}

	if err := registerModelDEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationFlags(depth+1, airmaxFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceStatusOverview1Attributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelDeviceAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceStatusOverview1Discovery(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var discoveryFlagName string
	if cmdPrefix == "" {
		discoveryFlagName = "discovery"
	} else {
		discoveryFlagName = fmt.Sprintf("%v.discovery", cmdPrefix)
	}

	if err := registerModelDiscoveryFlags(depth+1, discoveryFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceStatusOverview1Enabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	enabledDescription := `Required. `

	var enabledFlagName string
	if cmdPrefix == "" {
		enabledFlagName = "enabled"
	} else {
		enabledFlagName = fmt.Sprintf("%v.enabled", cmdPrefix)
	}

	var enabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(enabledFlagName, enabledFlagDefault, enabledDescription)

	return nil
}

func registerDeviceStatusOverview1Eswitch(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var eswitchFlagName string
	if cmdPrefix == "" {
		eswitchFlagName = "eswitch"
	} else {
		eswitchFlagName = fmt.Sprintf("%v.eswitch", cmdPrefix)
	}

	if err := registerModelEswitchFlags(depth+1, eswitchFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceStatusOverview1Firmware(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var firmwareFlagName string
	if cmdPrefix == "" {
		firmwareFlagName = "firmware"
	} else {
		firmwareFlagName = fmt.Sprintf("%v.firmware", cmdPrefix)
	}

	if err := registerModelDeviceFirmwareFlags(depth+1, firmwareFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceStatusOverview1Identification(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var identificationFlagName string
	if cmdPrefix == "" {
		identificationFlagName = "identification"
	} else {
		identificationFlagName = fmt.Sprintf("%v.identification", cmdPrefix)
	}

	if err := registerModelDeviceIdentificationFlags(depth+1, identificationFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceStatusOverview1Interfaces(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: interfaces DeviceInterfaceListSchema array type is not supported by go-swagger cli yet

	return nil
}

func registerDeviceStatusOverview1IPAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ipAddressDescription := `Required. Custom IP address in IPv4 or IPv6 format.`

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

func registerDeviceStatusOverview1LatestBackup(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var latestBackupFlagName string
	if cmdPrefix == "" {
		latestBackupFlagName = "latestBackup"
	} else {
		latestBackupFlagName = fmt.Sprintf("%v.latestBackup", cmdPrefix)
	}

	if err := registerModelLatestBackupFlags(depth+1, latestBackupFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceStatusOverview1Location(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var locationFlagName string
	if cmdPrefix == "" {
		locationFlagName = "location"
	} else {
		locationFlagName = fmt.Sprintf("%v.location", cmdPrefix)
	}

	if err := registerModelDeviceLocationFlags(depth+1, locationFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceStatusOverview1Meta(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var metaFlagName string
	if cmdPrefix == "" {
		metaFlagName = "meta"
	} else {
		metaFlagName = fmt.Sprintf("%v.meta", cmdPrefix)
	}

	if err := registerModelDeviceMetaFlags(depth+1, metaFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceStatusOverview1Mode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	modeDescription := ``

	var modeFlagName string
	if cmdPrefix == "" {
		modeFlagName = "mode"
	} else {
		modeFlagName = fmt.Sprintf("%v.mode", cmdPrefix)
	}

	var modeFlagDefault string

	_ = cmd.PersistentFlags().String(modeFlagName, modeFlagDefault, modeDescription)

	return nil
}

func registerDeviceStatusOverview1Overview(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var overviewFlagName string
	if cmdPrefix == "" {
		overviewFlagName = "overview"
	} else {
		overviewFlagName = fmt.Sprintf("%v.overview", cmdPrefix)
	}

	if err := registerModelDeviceOverviewFlags(depth+1, overviewFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceStatusOverview1Uisps(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var uispsFlagName string
	if cmdPrefix == "" {
		uispsFlagName = "uisps"
	} else {
		uispsFlagName = fmt.Sprintf("%v.uisps", cmdPrefix)
	}

	if err := registerModelUispsFlags(depth+1, uispsFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceStatusOverview1Upgrade(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var upgradeFlagName string
	if cmdPrefix == "" {
		upgradeFlagName = "upgrade"
	} else {
		upgradeFlagName = fmt.Sprintf("%v.upgrade", cmdPrefix)
	}

	if err := registerModelDeviceUpgradeFlags(depth+1, upgradeFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDeviceStatusOverview1Flags(depth int, m *models.DeviceStatusOverview1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, airfiberAdded := retrieveDeviceStatusOverview1AirfiberFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || airfiberAdded

	err, airmaxAdded := retrieveDeviceStatusOverview1AirmaxFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || airmaxAdded

	err, attributesAdded := retrieveDeviceStatusOverview1AttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, discoveryAdded := retrieveDeviceStatusOverview1DiscoveryFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || discoveryAdded

	err, enabledAdded := retrieveDeviceStatusOverview1EnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || enabledAdded

	err, eswitchAdded := retrieveDeviceStatusOverview1EswitchFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || eswitchAdded

	err, firmwareAdded := retrieveDeviceStatusOverview1FirmwareFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || firmwareAdded

	err, identificationAdded := retrieveDeviceStatusOverview1IdentificationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || identificationAdded

	err, interfacesAdded := retrieveDeviceStatusOverview1InterfacesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || interfacesAdded

	err, ipAddressAdded := retrieveDeviceStatusOverview1IPAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipAddressAdded

	err, latestBackupAdded := retrieveDeviceStatusOverview1LatestBackupFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || latestBackupAdded

	err, locationAdded := retrieveDeviceStatusOverview1LocationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || locationAdded

	err, metaAdded := retrieveDeviceStatusOverview1MetaFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || metaAdded

	err, modeAdded := retrieveDeviceStatusOverview1ModeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || modeAdded

	err, overviewAdded := retrieveDeviceStatusOverview1OverviewFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || overviewAdded

	err, uispsAdded := retrieveDeviceStatusOverview1UispsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || uispsAdded

	err, upgradeAdded := retrieveDeviceStatusOverview1UpgradeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || upgradeAdded

	return nil, retAdded
}

func retrieveDeviceStatusOverview1AirfiberFlags(depth int, m *models.DeviceStatusOverview1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	airfiberFlagName := fmt.Sprintf("%v.airfiber", cmdPrefix)
	if cmd.Flags().Changed(airfiberFlagName) {
		// info: complex object airfiber DEPRECATEDDeviceAirFiberPartiallyMovedPropertiesToInterfaceOrStation is retrieved outside this Changed() block
	}
	airfiberFlagValue := m.Airfiber
	if swag.IsZero(airfiberFlagValue) {
		airfiberFlagValue = &models.DEPRECATEDDeviceAirFiberPartiallyMovedPropertiesToInterfaceOrStation{}
	}

	err, airfiberAdded := retrieveModelDEPRECATEDDeviceAirFiberPartiallyMovedPropertiesToInterfaceOrStationFlags(depth+1, airfiberFlagValue, airfiberFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || airfiberAdded
	if airfiberAdded {
		m.Airfiber = airfiberFlagValue
	}

	return nil, retAdded
}

func retrieveDeviceStatusOverview1AirmaxFlags(depth int, m *models.DeviceStatusOverview1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	airmaxFlagName := fmt.Sprintf("%v.airmax", cmdPrefix)
	if cmd.Flags().Changed(airmaxFlagName) {
		// info: complex object airmax DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation is retrieved outside this Changed() block
	}
	airmaxFlagValue := m.Airmax
	if swag.IsZero(airmaxFlagValue) {
		airmaxFlagValue = &models.DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation{}
	}

	err, airmaxAdded := retrieveModelDEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationFlags(depth+1, airmaxFlagValue, airmaxFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || airmaxAdded
	if airmaxAdded {
		m.Airmax = airmaxFlagValue
	}

	return nil, retAdded
}

func retrieveDeviceStatusOverview1AttributesFlags(depth int, m *models.DeviceStatusOverview1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes DeviceAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &models.DeviceAttributes{}
	}

	err, attributesAdded := retrieveModelDeviceAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrieveDeviceStatusOverview1DiscoveryFlags(depth int, m *models.DeviceStatusOverview1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	discoveryFlagName := fmt.Sprintf("%v.discovery", cmdPrefix)
	if cmd.Flags().Changed(discoveryFlagName) {
		// info: complex object discovery Discovery is retrieved outside this Changed() block
	}
	discoveryFlagValue := m.Discovery
	if swag.IsZero(discoveryFlagValue) {
		discoveryFlagValue = &models.Discovery{}
	}

	err, discoveryAdded := retrieveModelDiscoveryFlags(depth+1, discoveryFlagValue, discoveryFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || discoveryAdded
	if discoveryAdded {
		m.Discovery = discoveryFlagValue
	}

	return nil, retAdded
}

func retrieveDeviceStatusOverview1EnabledFlags(depth int, m *models.DeviceStatusOverview1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	enabledFlagName := fmt.Sprintf("%v.enabled", cmdPrefix)
	if cmd.Flags().Changed(enabledFlagName) {

		var enabledFlagName string
		if cmdPrefix == "" {
			enabledFlagName = "enabled"
		} else {
			enabledFlagName = fmt.Sprintf("%v.enabled", cmdPrefix)
		}

		enabledFlagValue, err := cmd.Flags().GetBool(enabledFlagName)
		if err != nil {
			return err, false
		}
		m.Enabled = &enabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceStatusOverview1EswitchFlags(depth int, m *models.DeviceStatusOverview1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	eswitchFlagName := fmt.Sprintf("%v.eswitch", cmdPrefix)
	if cmd.Flags().Changed(eswitchFlagName) {
		// info: complex object eswitch Eswitch is retrieved outside this Changed() block
	}
	eswitchFlagValue := m.Eswitch
	if swag.IsZero(eswitchFlagValue) {
		eswitchFlagValue = &models.Eswitch{}
	}

	err, eswitchAdded := retrieveModelEswitchFlags(depth+1, eswitchFlagValue, eswitchFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || eswitchAdded
	if eswitchAdded {
		m.Eswitch = eswitchFlagValue
	}

	return nil, retAdded
}

func retrieveDeviceStatusOverview1FirmwareFlags(depth int, m *models.DeviceStatusOverview1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	firmwareFlagName := fmt.Sprintf("%v.firmware", cmdPrefix)
	if cmd.Flags().Changed(firmwareFlagName) {
		// info: complex object firmware DeviceFirmware is retrieved outside this Changed() block
	}
	firmwareFlagValue := m.Firmware
	if swag.IsZero(firmwareFlagValue) {
		firmwareFlagValue = &models.DeviceFirmware{}
	}

	err, firmwareAdded := retrieveModelDeviceFirmwareFlags(depth+1, firmwareFlagValue, firmwareFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || firmwareAdded
	if firmwareAdded {
		m.Firmware = firmwareFlagValue
	}

	return nil, retAdded
}

func retrieveDeviceStatusOverview1IdentificationFlags(depth int, m *models.DeviceStatusOverview1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	identificationFlagName := fmt.Sprintf("%v.identification", cmdPrefix)
	if cmd.Flags().Changed(identificationFlagName) {
		// info: complex object identification DeviceIdentification is retrieved outside this Changed() block
	}
	identificationFlagValue := m.Identification
	if swag.IsZero(identificationFlagValue) {
		identificationFlagValue = &models.DeviceIdentification{}
	}

	err, identificationAdded := retrieveModelDeviceIdentificationFlags(depth+1, identificationFlagValue, identificationFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || identificationAdded
	if identificationAdded {
		m.Identification = identificationFlagValue
	}

	return nil, retAdded
}

func retrieveDeviceStatusOverview1InterfacesFlags(depth int, m *models.DeviceStatusOverview1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	interfacesFlagName := fmt.Sprintf("%v.interfaces", cmdPrefix)
	if cmd.Flags().Changed(interfacesFlagName) {
		// warning: interfaces array type DeviceInterfaceListSchema is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDeviceStatusOverview1IPAddressFlags(depth int, m *models.DeviceStatusOverview1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.IPAddress = &ipAddressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceStatusOverview1LatestBackupFlags(depth int, m *models.DeviceStatusOverview1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	latestBackupFlagName := fmt.Sprintf("%v.latestBackup", cmdPrefix)
	if cmd.Flags().Changed(latestBackupFlagName) {
		// info: complex object latestBackup LatestBackup is retrieved outside this Changed() block
	}
	latestBackupFlagValue := m.LatestBackup
	if swag.IsZero(latestBackupFlagValue) {
		latestBackupFlagValue = &models.LatestBackup{}
	}

	err, latestBackupAdded := retrieveModelLatestBackupFlags(depth+1, latestBackupFlagValue, latestBackupFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || latestBackupAdded
	if latestBackupAdded {
		m.LatestBackup = latestBackupFlagValue
	}

	return nil, retAdded
}

func retrieveDeviceStatusOverview1LocationFlags(depth int, m *models.DeviceStatusOverview1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	locationFlagName := fmt.Sprintf("%v.location", cmdPrefix)
	if cmd.Flags().Changed(locationFlagName) {
		// info: complex object location DeviceLocation is retrieved outside this Changed() block
	}
	locationFlagValue := m.Location
	if swag.IsZero(locationFlagValue) {
		locationFlagValue = &models.DeviceLocation{}
	}

	err, locationAdded := retrieveModelDeviceLocationFlags(depth+1, locationFlagValue, locationFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || locationAdded
	if locationAdded {
		m.Location = locationFlagValue
	}

	return nil, retAdded
}

func retrieveDeviceStatusOverview1MetaFlags(depth int, m *models.DeviceStatusOverview1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	metaFlagName := fmt.Sprintf("%v.meta", cmdPrefix)
	if cmd.Flags().Changed(metaFlagName) {
		// info: complex object meta DeviceMeta is retrieved outside this Changed() block
	}
	metaFlagValue := m.Meta
	if swag.IsZero(metaFlagValue) {
		metaFlagValue = &models.DeviceMeta{}
	}

	err, metaAdded := retrieveModelDeviceMetaFlags(depth+1, metaFlagValue, metaFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || metaAdded
	if metaAdded {
		m.Meta = metaFlagValue
	}

	return nil, retAdded
}

func retrieveDeviceStatusOverview1ModeFlags(depth int, m *models.DeviceStatusOverview1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	modeFlagName := fmt.Sprintf("%v.mode", cmdPrefix)
	if cmd.Flags().Changed(modeFlagName) {

		var modeFlagName string
		if cmdPrefix == "" {
			modeFlagName = "mode"
		} else {
			modeFlagName = fmt.Sprintf("%v.mode", cmdPrefix)
		}

		modeFlagValue, err := cmd.Flags().GetString(modeFlagName)
		if err != nil {
			return err, false
		}
		m.Mode = modeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceStatusOverview1OverviewFlags(depth int, m *models.DeviceStatusOverview1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	overviewFlagName := fmt.Sprintf("%v.overview", cmdPrefix)
	if cmd.Flags().Changed(overviewFlagName) {
		// info: complex object overview DeviceOverview is retrieved outside this Changed() block
	}
	overviewFlagValue := m.Overview
	if swag.IsZero(overviewFlagValue) {
		overviewFlagValue = &models.DeviceOverview{}
	}

	err, overviewAdded := retrieveModelDeviceOverviewFlags(depth+1, overviewFlagValue, overviewFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || overviewAdded
	if overviewAdded {
		m.Overview = overviewFlagValue
	}

	return nil, retAdded
}

func retrieveDeviceStatusOverview1UispsFlags(depth int, m *models.DeviceStatusOverview1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	uispsFlagName := fmt.Sprintf("%v.uisps", cmdPrefix)
	if cmd.Flags().Changed(uispsFlagName) {
		// info: complex object uisps Uisps is retrieved outside this Changed() block
	}
	uispsFlagValue := m.Uisps
	if swag.IsZero(uispsFlagValue) {
		uispsFlagValue = &models.Uisps{}
	}

	err, uispsAdded := retrieveModelUispsFlags(depth+1, uispsFlagValue, uispsFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || uispsAdded
	if uispsAdded {
		m.Uisps = uispsFlagValue
	}

	return nil, retAdded
}

func retrieveDeviceStatusOverview1UpgradeFlags(depth int, m *models.DeviceStatusOverview1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	upgradeFlagName := fmt.Sprintf("%v.upgrade", cmdPrefix)
	if cmd.Flags().Changed(upgradeFlagName) {
		// info: complex object upgrade DeviceUpgrade is retrieved outside this Changed() block
	}
	upgradeFlagValue := m.Upgrade
	if swag.IsZero(upgradeFlagValue) {
		upgradeFlagValue = &models.DeviceUpgrade{}
	}

	err, upgradeAdded := retrieveModelDeviceUpgradeFlags(depth+1, upgradeFlagValue, upgradeFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || upgradeAdded
	if upgradeAdded {
		m.Upgrade = upgradeFlagValue
	}

	return nil, retAdded
}
