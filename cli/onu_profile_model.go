// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for OnuProfile

// register flags to command
func registerModelOnuProfileFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerOnuProfileAdminPassword(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuProfileBandwidthLimitDown(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuProfileBandwidthLimitEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuProfileBandwidthLimitUp(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuProfileBridge(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuProfileHTTPPort(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuProfileID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuProfileLanAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuProfileLanProvisioned(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuProfileLinkSpeed(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuProfileMode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuProfileName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuProfileOnuCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuProfileRouter(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuProfileSSHEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuProfileSSHPort(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuProfileTelnetEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuProfileTelnetPort(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuProfileUbntDiscoveryEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerOnuProfileAdminPassword(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	adminPasswordDescription := ``

	var adminPasswordFlagName string
	if cmdPrefix == "" {
		adminPasswordFlagName = "adminPassword"
	} else {
		adminPasswordFlagName = fmt.Sprintf("%v.adminPassword", cmdPrefix)
	}

	var adminPasswordFlagDefault string

	_ = cmd.PersistentFlags().String(adminPasswordFlagName, adminPasswordFlagDefault, adminPasswordDescription)

	return nil
}

func registerOnuProfileBandwidthLimitDown(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	bandwidthLimitDownDescription := ``

	var bandwidthLimitDownFlagName string
	if cmdPrefix == "" {
		bandwidthLimitDownFlagName = "bandwidthLimitDown"
	} else {
		bandwidthLimitDownFlagName = fmt.Sprintf("%v.bandwidthLimitDown", cmdPrefix)
	}

	var bandwidthLimitDownFlagDefault int64

	_ = cmd.PersistentFlags().Int64(bandwidthLimitDownFlagName, bandwidthLimitDownFlagDefault, bandwidthLimitDownDescription)

	return nil
}

func registerOnuProfileBandwidthLimitEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	bandwidthLimitEnabledDescription := ``

	var bandwidthLimitEnabledFlagName string
	if cmdPrefix == "" {
		bandwidthLimitEnabledFlagName = "bandwidthLimitEnabled"
	} else {
		bandwidthLimitEnabledFlagName = fmt.Sprintf("%v.bandwidthLimitEnabled", cmdPrefix)
	}

	var bandwidthLimitEnabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(bandwidthLimitEnabledFlagName, bandwidthLimitEnabledFlagDefault, bandwidthLimitEnabledDescription)

	return nil
}

func registerOnuProfileBandwidthLimitUp(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	bandwidthLimitUpDescription := ``

	var bandwidthLimitUpFlagName string
	if cmdPrefix == "" {
		bandwidthLimitUpFlagName = "bandwidthLimitUp"
	} else {
		bandwidthLimitUpFlagName = fmt.Sprintf("%v.bandwidthLimitUp", cmdPrefix)
	}

	var bandwidthLimitUpFlagDefault int64

	_ = cmd.PersistentFlags().Int64(bandwidthLimitUpFlagName, bandwidthLimitUpFlagDefault, bandwidthLimitUpDescription)

	return nil
}

func registerOnuProfileBridge(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var bridgeFlagName string
	if cmdPrefix == "" {
		bridgeFlagName = "bridge"
	} else {
		bridgeFlagName = fmt.Sprintf("%v.bridge", cmdPrefix)
	}

	if err := registerModelBridgeFlags(depth+1, bridgeFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerOnuProfileHTTPPort(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	httpPortDescription := ``

	var httpPortFlagName string
	if cmdPrefix == "" {
		httpPortFlagName = "httpPort"
	} else {
		httpPortFlagName = fmt.Sprintf("%v.httpPort", cmdPrefix)
	}

	var httpPortFlagDefault int64

	_ = cmd.PersistentFlags().Int64(httpPortFlagName, httpPortFlagDefault, httpPortDescription)

	return nil
}

func registerOnuProfileID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := ``

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

func registerOnuProfileLanAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	lanAddressDescription := ``

	var lanAddressFlagName string
	if cmdPrefix == "" {
		lanAddressFlagName = "lanAddress"
	} else {
		lanAddressFlagName = fmt.Sprintf("%v.lanAddress", cmdPrefix)
	}

	var lanAddressFlagDefault string

	_ = cmd.PersistentFlags().String(lanAddressFlagName, lanAddressFlagDefault, lanAddressDescription)

	return nil
}

func registerOnuProfileLanProvisioned(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	lanProvisionedDescription := ``

	var lanProvisionedFlagName string
	if cmdPrefix == "" {
		lanProvisionedFlagName = "lanProvisioned"
	} else {
		lanProvisionedFlagName = fmt.Sprintf("%v.lanProvisioned", cmdPrefix)
	}

	var lanProvisionedFlagDefault bool

	_ = cmd.PersistentFlags().Bool(lanProvisionedFlagName, lanProvisionedFlagDefault, lanProvisionedDescription)

	return nil
}

func registerOnuProfileLinkSpeed(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: linkSpeed LinkSpeed map type is not supported by go-swagger cli yet

	return nil
}

func registerOnuProfileMode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	modeDescription := `Enum: ["bridge","router","soho"]. `

	var modeFlagName string
	if cmdPrefix == "" {
		modeFlagName = "mode"
	} else {
		modeFlagName = fmt.Sprintf("%v.mode", cmdPrefix)
	}

	var modeFlagDefault string

	_ = cmd.PersistentFlags().String(modeFlagName, modeFlagDefault, modeDescription)

	if err := cmd.RegisterFlagCompletionFunc(modeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["bridge","router","soho"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerOnuProfileName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerOnuProfileOnuCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	onuCountDescription := ``

	var onuCountFlagName string
	if cmdPrefix == "" {
		onuCountFlagName = "onuCount"
	} else {
		onuCountFlagName = fmt.Sprintf("%v.onuCount", cmdPrefix)
	}

	var onuCountFlagDefault float64

	_ = cmd.PersistentFlags().Float64(onuCountFlagName, onuCountFlagDefault, onuCountDescription)

	return nil
}

func registerOnuProfileRouter(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var routerFlagName string
	if cmdPrefix == "" {
		routerFlagName = "router"
	} else {
		routerFlagName = fmt.Sprintf("%v.router", cmdPrefix)
	}

	if err := registerModelRouterFlags(depth+1, routerFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerOnuProfileSSHEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sshEnabledDescription := ``

	var sshEnabledFlagName string
	if cmdPrefix == "" {
		sshEnabledFlagName = "sshEnabled"
	} else {
		sshEnabledFlagName = fmt.Sprintf("%v.sshEnabled", cmdPrefix)
	}

	var sshEnabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(sshEnabledFlagName, sshEnabledFlagDefault, sshEnabledDescription)

	return nil
}

func registerOnuProfileSSHPort(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sshPortDescription := ``

	var sshPortFlagName string
	if cmdPrefix == "" {
		sshPortFlagName = "sshPort"
	} else {
		sshPortFlagName = fmt.Sprintf("%v.sshPort", cmdPrefix)
	}

	var sshPortFlagDefault int64

	_ = cmd.PersistentFlags().Int64(sshPortFlagName, sshPortFlagDefault, sshPortDescription)

	return nil
}

func registerOnuProfileTelnetEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	telnetEnabledDescription := ``

	var telnetEnabledFlagName string
	if cmdPrefix == "" {
		telnetEnabledFlagName = "telnetEnabled"
	} else {
		telnetEnabledFlagName = fmt.Sprintf("%v.telnetEnabled", cmdPrefix)
	}

	var telnetEnabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(telnetEnabledFlagName, telnetEnabledFlagDefault, telnetEnabledDescription)

	return nil
}

func registerOnuProfileTelnetPort(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	telnetPortDescription := ``

	var telnetPortFlagName string
	if cmdPrefix == "" {
		telnetPortFlagName = "telnetPort"
	} else {
		telnetPortFlagName = fmt.Sprintf("%v.telnetPort", cmdPrefix)
	}

	var telnetPortFlagDefault int64

	_ = cmd.PersistentFlags().Int64(telnetPortFlagName, telnetPortFlagDefault, telnetPortDescription)

	return nil
}

func registerOnuProfileUbntDiscoveryEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ubntDiscoveryEnabledDescription := ``

	var ubntDiscoveryEnabledFlagName string
	if cmdPrefix == "" {
		ubntDiscoveryEnabledFlagName = "ubntDiscoveryEnabled"
	} else {
		ubntDiscoveryEnabledFlagName = fmt.Sprintf("%v.ubntDiscoveryEnabled", cmdPrefix)
	}

	var ubntDiscoveryEnabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(ubntDiscoveryEnabledFlagName, ubntDiscoveryEnabledFlagDefault, ubntDiscoveryEnabledDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelOnuProfileFlags(depth int, m *models.OnuProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, adminPasswordAdded := retrieveOnuProfileAdminPasswordFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || adminPasswordAdded

	err, bandwidthLimitDownAdded := retrieveOnuProfileBandwidthLimitDownFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || bandwidthLimitDownAdded

	err, bandwidthLimitEnabledAdded := retrieveOnuProfileBandwidthLimitEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || bandwidthLimitEnabledAdded

	err, bandwidthLimitUpAdded := retrieveOnuProfileBandwidthLimitUpFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || bandwidthLimitUpAdded

	err, bridgeAdded := retrieveOnuProfileBridgeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || bridgeAdded

	err, httpPortAdded := retrieveOnuProfileHTTPPortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || httpPortAdded

	err, idAdded := retrieveOnuProfileIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, lanAddressAdded := retrieveOnuProfileLanAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || lanAddressAdded

	err, lanProvisionedAdded := retrieveOnuProfileLanProvisionedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || lanProvisionedAdded

	err, linkSpeedAdded := retrieveOnuProfileLinkSpeedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || linkSpeedAdded

	err, modeAdded := retrieveOnuProfileModeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || modeAdded

	err, nameAdded := retrieveOnuProfileNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, onuCountAdded := retrieveOnuProfileOnuCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || onuCountAdded

	err, routerAdded := retrieveOnuProfileRouterFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || routerAdded

	err, sshEnabledAdded := retrieveOnuProfileSSHEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sshEnabledAdded

	err, sshPortAdded := retrieveOnuProfileSSHPortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sshPortAdded

	err, telnetEnabledAdded := retrieveOnuProfileTelnetEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || telnetEnabledAdded

	err, telnetPortAdded := retrieveOnuProfileTelnetPortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || telnetPortAdded

	err, ubntDiscoveryEnabledAdded := retrieveOnuProfileUbntDiscoveryEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ubntDiscoveryEnabledAdded

	return nil, retAdded
}

func retrieveOnuProfileAdminPasswordFlags(depth int, m *models.OnuProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	adminPasswordFlagName := fmt.Sprintf("%v.adminPassword", cmdPrefix)
	if cmd.Flags().Changed(adminPasswordFlagName) {

		var adminPasswordFlagName string
		if cmdPrefix == "" {
			adminPasswordFlagName = "adminPassword"
		} else {
			adminPasswordFlagName = fmt.Sprintf("%v.adminPassword", cmdPrefix)
		}

		adminPasswordFlagValue, err := cmd.Flags().GetString(adminPasswordFlagName)
		if err != nil {
			return err, false
		}
		m.AdminPassword = adminPasswordFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuProfileBandwidthLimitDownFlags(depth int, m *models.OnuProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	bandwidthLimitDownFlagName := fmt.Sprintf("%v.bandwidthLimitDown", cmdPrefix)
	if cmd.Flags().Changed(bandwidthLimitDownFlagName) {

		var bandwidthLimitDownFlagName string
		if cmdPrefix == "" {
			bandwidthLimitDownFlagName = "bandwidthLimitDown"
		} else {
			bandwidthLimitDownFlagName = fmt.Sprintf("%v.bandwidthLimitDown", cmdPrefix)
		}

		bandwidthLimitDownFlagValue, err := cmd.Flags().GetInt64(bandwidthLimitDownFlagName)
		if err != nil {
			return err, false
		}
		m.BandwidthLimitDown = bandwidthLimitDownFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuProfileBandwidthLimitEnabledFlags(depth int, m *models.OnuProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	bandwidthLimitEnabledFlagName := fmt.Sprintf("%v.bandwidthLimitEnabled", cmdPrefix)
	if cmd.Flags().Changed(bandwidthLimitEnabledFlagName) {

		var bandwidthLimitEnabledFlagName string
		if cmdPrefix == "" {
			bandwidthLimitEnabledFlagName = "bandwidthLimitEnabled"
		} else {
			bandwidthLimitEnabledFlagName = fmt.Sprintf("%v.bandwidthLimitEnabled", cmdPrefix)
		}

		bandwidthLimitEnabledFlagValue, err := cmd.Flags().GetBool(bandwidthLimitEnabledFlagName)
		if err != nil {
			return err, false
		}
		m.BandwidthLimitEnabled = bandwidthLimitEnabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuProfileBandwidthLimitUpFlags(depth int, m *models.OnuProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	bandwidthLimitUpFlagName := fmt.Sprintf("%v.bandwidthLimitUp", cmdPrefix)
	if cmd.Flags().Changed(bandwidthLimitUpFlagName) {

		var bandwidthLimitUpFlagName string
		if cmdPrefix == "" {
			bandwidthLimitUpFlagName = "bandwidthLimitUp"
		} else {
			bandwidthLimitUpFlagName = fmt.Sprintf("%v.bandwidthLimitUp", cmdPrefix)
		}

		bandwidthLimitUpFlagValue, err := cmd.Flags().GetInt64(bandwidthLimitUpFlagName)
		if err != nil {
			return err, false
		}
		m.BandwidthLimitUp = bandwidthLimitUpFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuProfileBridgeFlags(depth int, m *models.OnuProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	bridgeFlagName := fmt.Sprintf("%v.bridge", cmdPrefix)
	if cmd.Flags().Changed(bridgeFlagName) {
		// info: complex object bridge Bridge is retrieved outside this Changed() block
	}
	bridgeFlagValue := m.Bridge
	if swag.IsZero(bridgeFlagValue) {
		bridgeFlagValue = &models.Bridge{}
	}

	err, bridgeAdded := retrieveModelBridgeFlags(depth+1, bridgeFlagValue, bridgeFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || bridgeAdded
	if bridgeAdded {
		m.Bridge = bridgeFlagValue
	}

	return nil, retAdded
}

func retrieveOnuProfileHTTPPortFlags(depth int, m *models.OnuProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	httpPortFlagName := fmt.Sprintf("%v.httpPort", cmdPrefix)
	if cmd.Flags().Changed(httpPortFlagName) {

		var httpPortFlagName string
		if cmdPrefix == "" {
			httpPortFlagName = "httpPort"
		} else {
			httpPortFlagName = fmt.Sprintf("%v.httpPort", cmdPrefix)
		}

		httpPortFlagValue, err := cmd.Flags().GetInt64(httpPortFlagName)
		if err != nil {
			return err, false
		}
		m.HTTPPort = httpPortFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuProfileIDFlags(depth int, m *models.OnuProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuProfileLanAddressFlags(depth int, m *models.OnuProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	lanAddressFlagName := fmt.Sprintf("%v.lanAddress", cmdPrefix)
	if cmd.Flags().Changed(lanAddressFlagName) {

		var lanAddressFlagName string
		if cmdPrefix == "" {
			lanAddressFlagName = "lanAddress"
		} else {
			lanAddressFlagName = fmt.Sprintf("%v.lanAddress", cmdPrefix)
		}

		lanAddressFlagValue, err := cmd.Flags().GetString(lanAddressFlagName)
		if err != nil {
			return err, false
		}
		m.LanAddress = lanAddressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuProfileLanProvisionedFlags(depth int, m *models.OnuProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	lanProvisionedFlagName := fmt.Sprintf("%v.lanProvisioned", cmdPrefix)
	if cmd.Flags().Changed(lanProvisionedFlagName) {

		var lanProvisionedFlagName string
		if cmdPrefix == "" {
			lanProvisionedFlagName = "lanProvisioned"
		} else {
			lanProvisionedFlagName = fmt.Sprintf("%v.lanProvisioned", cmdPrefix)
		}

		lanProvisionedFlagValue, err := cmd.Flags().GetBool(lanProvisionedFlagName)
		if err != nil {
			return err, false
		}
		m.LanProvisioned = lanProvisionedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuProfileLinkSpeedFlags(depth int, m *models.OnuProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	linkSpeedFlagName := fmt.Sprintf("%v.linkSpeed", cmdPrefix)
	if cmd.Flags().Changed(linkSpeedFlagName) {
		// warning: linkSpeed map type LinkSpeed is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveOnuProfileModeFlags(depth int, m *models.OnuProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveOnuProfileNameFlags(depth int, m *models.OnuProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveOnuProfileOnuCountFlags(depth int, m *models.OnuProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	onuCountFlagName := fmt.Sprintf("%v.onuCount", cmdPrefix)
	if cmd.Flags().Changed(onuCountFlagName) {

		var onuCountFlagName string
		if cmdPrefix == "" {
			onuCountFlagName = "onuCount"
		} else {
			onuCountFlagName = fmt.Sprintf("%v.onuCount", cmdPrefix)
		}

		onuCountFlagValue, err := cmd.Flags().GetFloat64(onuCountFlagName)
		if err != nil {
			return err, false
		}
		m.OnuCount = onuCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuProfileRouterFlags(depth int, m *models.OnuProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	routerFlagName := fmt.Sprintf("%v.router", cmdPrefix)
	if cmd.Flags().Changed(routerFlagName) {
		// info: complex object router Router is retrieved outside this Changed() block
	}
	routerFlagValue := m.Router
	if swag.IsZero(routerFlagValue) {
		routerFlagValue = &models.Router{}
	}

	err, routerAdded := retrieveModelRouterFlags(depth+1, routerFlagValue, routerFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || routerAdded
	if routerAdded {
		m.Router = routerFlagValue
	}

	return nil, retAdded
}

func retrieveOnuProfileSSHEnabledFlags(depth int, m *models.OnuProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sshEnabledFlagName := fmt.Sprintf("%v.sshEnabled", cmdPrefix)
	if cmd.Flags().Changed(sshEnabledFlagName) {

		var sshEnabledFlagName string
		if cmdPrefix == "" {
			sshEnabledFlagName = "sshEnabled"
		} else {
			sshEnabledFlagName = fmt.Sprintf("%v.sshEnabled", cmdPrefix)
		}

		sshEnabledFlagValue, err := cmd.Flags().GetBool(sshEnabledFlagName)
		if err != nil {
			return err, false
		}
		m.SSHEnabled = sshEnabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuProfileSSHPortFlags(depth int, m *models.OnuProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sshPortFlagName := fmt.Sprintf("%v.sshPort", cmdPrefix)
	if cmd.Flags().Changed(sshPortFlagName) {

		var sshPortFlagName string
		if cmdPrefix == "" {
			sshPortFlagName = "sshPort"
		} else {
			sshPortFlagName = fmt.Sprintf("%v.sshPort", cmdPrefix)
		}

		sshPortFlagValue, err := cmd.Flags().GetInt64(sshPortFlagName)
		if err != nil {
			return err, false
		}
		m.SSHPort = sshPortFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuProfileTelnetEnabledFlags(depth int, m *models.OnuProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	telnetEnabledFlagName := fmt.Sprintf("%v.telnetEnabled", cmdPrefix)
	if cmd.Flags().Changed(telnetEnabledFlagName) {

		var telnetEnabledFlagName string
		if cmdPrefix == "" {
			telnetEnabledFlagName = "telnetEnabled"
		} else {
			telnetEnabledFlagName = fmt.Sprintf("%v.telnetEnabled", cmdPrefix)
		}

		telnetEnabledFlagValue, err := cmd.Flags().GetBool(telnetEnabledFlagName)
		if err != nil {
			return err, false
		}
		m.TelnetEnabled = telnetEnabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuProfileTelnetPortFlags(depth int, m *models.OnuProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	telnetPortFlagName := fmt.Sprintf("%v.telnetPort", cmdPrefix)
	if cmd.Flags().Changed(telnetPortFlagName) {

		var telnetPortFlagName string
		if cmdPrefix == "" {
			telnetPortFlagName = "telnetPort"
		} else {
			telnetPortFlagName = fmt.Sprintf("%v.telnetPort", cmdPrefix)
		}

		telnetPortFlagValue, err := cmd.Flags().GetInt64(telnetPortFlagName)
		if err != nil {
			return err, false
		}
		m.TelnetPort = telnetPortFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuProfileUbntDiscoveryEnabledFlags(depth int, m *models.OnuProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ubntDiscoveryEnabledFlagName := fmt.Sprintf("%v.ubntDiscoveryEnabled", cmdPrefix)
	if cmd.Flags().Changed(ubntDiscoveryEnabledFlagName) {

		var ubntDiscoveryEnabledFlagName string
		if cmdPrefix == "" {
			ubntDiscoveryEnabledFlagName = "ubntDiscoveryEnabled"
		} else {
			ubntDiscoveryEnabledFlagName = fmt.Sprintf("%v.ubntDiscoveryEnabled", cmdPrefix)
		}

		ubntDiscoveryEnabledFlagValue, err := cmd.Flags().GetBool(ubntDiscoveryEnabledFlagName)
		if err != nil {
			return err, false
		}
		m.UbntDiscoveryEnabled = ubntDiscoveryEnabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}
