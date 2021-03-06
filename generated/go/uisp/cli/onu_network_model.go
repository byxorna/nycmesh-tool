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

// Schema cli for OnuNetwork

// register flags to command
func registerModelOnuNetworkFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerOnuNetworkDownloadLimit(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuNetworkDownloadLimitEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuNetworkFirewallEnabled6(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuNetworkIPV4(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuNetworkIPV6(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuNetworkIsIPV6Enabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuNetworkLanAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuNetworkLanProvisioned(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuNetworkMode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuNetworkNatFtp(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuNetworkNatPptp(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuNetworkNatRtsp(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuNetworkNatSip(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuNetworkPortForwards(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuNetworkPorts(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuNetworkUploadLimit(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuNetworkUploadLimitEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuNetworkUpnpEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuNetworkVlans(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuNetworkWanVlan(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerOnuNetworkDownloadLimit(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	downloadLimitDescription := `Required. Download bandwidth limit in Mbps`

	var downloadLimitFlagName string
	if cmdPrefix == "" {
		downloadLimitFlagName = "downloadLimit"
	} else {
		downloadLimitFlagName = fmt.Sprintf("%v.downloadLimit", cmdPrefix)
	}

	var downloadLimitFlagDefault int64

	_ = cmd.PersistentFlags().Int64(downloadLimitFlagName, downloadLimitFlagDefault, downloadLimitDescription)

	return nil
}

func registerOnuNetworkDownloadLimitEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	downloadLimitEnabledDescription := `Required. Set to true to enable download bandwidth limit`

	var downloadLimitEnabledFlagName string
	if cmdPrefix == "" {
		downloadLimitEnabledFlagName = "downloadLimitEnabled"
	} else {
		downloadLimitEnabledFlagName = fmt.Sprintf("%v.downloadLimitEnabled", cmdPrefix)
	}

	var downloadLimitEnabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(downloadLimitEnabledFlagName, downloadLimitEnabledFlagDefault, downloadLimitEnabledDescription)

	return nil
}

func registerOnuNetworkFirewallEnabled6(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	firewallEnabled6Description := `Required. Set to true to enable IPv6 statefull firewall`

	var firewallEnabled6FlagName string
	if cmdPrefix == "" {
		firewallEnabled6FlagName = "firewallEnabled6"
	} else {
		firewallEnabled6FlagName = fmt.Sprintf("%v.firewallEnabled6", cmdPrefix)
	}

	var firewallEnabled6FlagDefault bool

	_ = cmd.PersistentFlags().Bool(firewallEnabled6FlagName, firewallEnabled6FlagDefault, firewallEnabled6Description)

	return nil
}

func registerOnuNetworkIPV4(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var ipv4FlagName string
	if cmdPrefix == "" {
		ipv4FlagName = "ipv4"
	} else {
		ipv4FlagName = fmt.Sprintf("%v.ipv4", cmdPrefix)
	}

	if err := registerModelIPV4Flags(depth+1, ipv4FlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerOnuNetworkIPV6(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var ipv6FlagName string
	if cmdPrefix == "" {
		ipv6FlagName = "ipv6"
	} else {
		ipv6FlagName = fmt.Sprintf("%v.ipv6", cmdPrefix)
	}

	if err := registerModelIPV6Flags(depth+1, ipv6FlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerOnuNetworkIsIPV6Enabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	isIpv6EnabledDescription := `Required. Set to true to enable IPv6`

	var isIpv6EnabledFlagName string
	if cmdPrefix == "" {
		isIpv6EnabledFlagName = "isIpv6Enabled"
	} else {
		isIpv6EnabledFlagName = fmt.Sprintf("%v.isIpv6Enabled", cmdPrefix)
	}

	var isIpv6EnabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(isIpv6EnabledFlagName, isIpv6EnabledFlagDefault, isIpv6EnabledDescription)

	return nil
}

func registerOnuNetworkLanAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	lanAddressDescription := `Required. `

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

func registerOnuNetworkLanProvisioned(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	lanProvisionedDescription := `Required. Set to true to enable setting LAN from OLT`

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

func registerOnuNetworkMode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	modeDescription := `Enum: ["bridge","router","soho"]. Required. Onu network mode`

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

func registerOnuNetworkNatFtp(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	natFtpDescription := `Required. Set to true to enable NAT for FTP protocol`

	var natFtpFlagName string
	if cmdPrefix == "" {
		natFtpFlagName = "natFtp"
	} else {
		natFtpFlagName = fmt.Sprintf("%v.natFtp", cmdPrefix)
	}

	var natFtpFlagDefault bool

	_ = cmd.PersistentFlags().Bool(natFtpFlagName, natFtpFlagDefault, natFtpDescription)

	return nil
}

func registerOnuNetworkNatPptp(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	natPptpDescription := `Required. Set to true to enable NAT for PPTP protocol`

	var natPptpFlagName string
	if cmdPrefix == "" {
		natPptpFlagName = "natPptp"
	} else {
		natPptpFlagName = fmt.Sprintf("%v.natPptp", cmdPrefix)
	}

	var natPptpFlagDefault bool

	_ = cmd.PersistentFlags().Bool(natPptpFlagName, natPptpFlagDefault, natPptpDescription)

	return nil
}

func registerOnuNetworkNatRtsp(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	natRtspDescription := `Required. Set to true to enable NAT for RTSP protocol`

	var natRtspFlagName string
	if cmdPrefix == "" {
		natRtspFlagName = "natRtsp"
	} else {
		natRtspFlagName = fmt.Sprintf("%v.natRtsp", cmdPrefix)
	}

	var natRtspFlagDefault bool

	_ = cmd.PersistentFlags().Bool(natRtspFlagName, natRtspFlagDefault, natRtspDescription)

	return nil
}

func registerOnuNetworkNatSip(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	natSipDescription := `Required. Set to true to enable NAT for SIP protocol`

	var natSipFlagName string
	if cmdPrefix == "" {
		natSipFlagName = "natSip"
	} else {
		natSipFlagName = fmt.Sprintf("%v.natSip", cmdPrefix)
	}

	var natSipFlagDefault bool

	_ = cmd.PersistentFlags().Bool(natSipFlagName, natSipFlagDefault, natSipDescription)

	return nil
}

func registerOnuNetworkPortForwards(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: portForwards PortForwards array type is not supported by go-swagger cli yet

	return nil
}

func registerOnuNetworkPorts(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: ports Ports5 array type is not supported by go-swagger cli yet

	return nil
}

func registerOnuNetworkUploadLimit(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	uploadLimitDescription := `Required. Upload bandwidth limit in Mbps`

	var uploadLimitFlagName string
	if cmdPrefix == "" {
		uploadLimitFlagName = "uploadLimit"
	} else {
		uploadLimitFlagName = fmt.Sprintf("%v.uploadLimit", cmdPrefix)
	}

	var uploadLimitFlagDefault int64

	_ = cmd.PersistentFlags().Int64(uploadLimitFlagName, uploadLimitFlagDefault, uploadLimitDescription)

	return nil
}

func registerOnuNetworkUploadLimitEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	uploadLimitEnabledDescription := `Required. Set to true to enable upload bandwidth limit`

	var uploadLimitEnabledFlagName string
	if cmdPrefix == "" {
		uploadLimitEnabledFlagName = "uploadLimitEnabled"
	} else {
		uploadLimitEnabledFlagName = fmt.Sprintf("%v.uploadLimitEnabled", cmdPrefix)
	}

	var uploadLimitEnabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(uploadLimitEnabledFlagName, uploadLimitEnabledFlagDefault, uploadLimitEnabledDescription)

	return nil
}

func registerOnuNetworkUpnpEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	upnpEnabledDescription := `Required. Set to true to enable UPnP`

	var upnpEnabledFlagName string
	if cmdPrefix == "" {
		upnpEnabledFlagName = "upnpEnabled"
	} else {
		upnpEnabledFlagName = fmt.Sprintf("%v.upnpEnabled", cmdPrefix)
	}

	var upnpEnabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(upnpEnabledFlagName, upnpEnabledFlagDefault, upnpEnabledDescription)

	return nil
}

func registerOnuNetworkVlans(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: vlans Vlans3 array type is not supported by go-swagger cli yet

	return nil
}

func registerOnuNetworkWanVlan(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	wanVlanDescription := `Required. WAN VLAN identification`

	var wanVlanFlagName string
	if cmdPrefix == "" {
		wanVlanFlagName = "wanVlan"
	} else {
		wanVlanFlagName = fmt.Sprintf("%v.wanVlan", cmdPrefix)
	}

	var wanVlanFlagDefault int64

	_ = cmd.PersistentFlags().Int64(wanVlanFlagName, wanVlanFlagDefault, wanVlanDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelOnuNetworkFlags(depth int, m *models.OnuNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, downloadLimitAdded := retrieveOnuNetworkDownloadLimitFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || downloadLimitAdded

	err, downloadLimitEnabledAdded := retrieveOnuNetworkDownloadLimitEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || downloadLimitEnabledAdded

	err, firewallEnabled6Added := retrieveOnuNetworkFirewallEnabled6Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || firewallEnabled6Added

	err, ipv4Added := retrieveOnuNetworkIPV4Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipv4Added

	err, ipv6Added := retrieveOnuNetworkIPV6Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipv6Added

	err, isIpv6EnabledAdded := retrieveOnuNetworkIsIPV6EnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || isIpv6EnabledAdded

	err, lanAddressAdded := retrieveOnuNetworkLanAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || lanAddressAdded

	err, lanProvisionedAdded := retrieveOnuNetworkLanProvisionedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || lanProvisionedAdded

	err, modeAdded := retrieveOnuNetworkModeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || modeAdded

	err, natFtpAdded := retrieveOnuNetworkNatFtpFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || natFtpAdded

	err, natPptpAdded := retrieveOnuNetworkNatPptpFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || natPptpAdded

	err, natRtspAdded := retrieveOnuNetworkNatRtspFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || natRtspAdded

	err, natSipAdded := retrieveOnuNetworkNatSipFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || natSipAdded

	err, portForwardsAdded := retrieveOnuNetworkPortForwardsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || portForwardsAdded

	err, portsAdded := retrieveOnuNetworkPortsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || portsAdded

	err, uploadLimitAdded := retrieveOnuNetworkUploadLimitFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || uploadLimitAdded

	err, uploadLimitEnabledAdded := retrieveOnuNetworkUploadLimitEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || uploadLimitEnabledAdded

	err, upnpEnabledAdded := retrieveOnuNetworkUpnpEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || upnpEnabledAdded

	err, vlansAdded := retrieveOnuNetworkVlansFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || vlansAdded

	err, wanVlanAdded := retrieveOnuNetworkWanVlanFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || wanVlanAdded

	return nil, retAdded
}

func retrieveOnuNetworkDownloadLimitFlags(depth int, m *models.OnuNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	downloadLimitFlagName := fmt.Sprintf("%v.downloadLimit", cmdPrefix)
	if cmd.Flags().Changed(downloadLimitFlagName) {

		var downloadLimitFlagName string
		if cmdPrefix == "" {
			downloadLimitFlagName = "downloadLimit"
		} else {
			downloadLimitFlagName = fmt.Sprintf("%v.downloadLimit", cmdPrefix)
		}

		downloadLimitFlagValue, err := cmd.Flags().GetInt64(downloadLimitFlagName)
		if err != nil {
			return err, false
		}
		m.DownloadLimit = &downloadLimitFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuNetworkDownloadLimitEnabledFlags(depth int, m *models.OnuNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	downloadLimitEnabledFlagName := fmt.Sprintf("%v.downloadLimitEnabled", cmdPrefix)
	if cmd.Flags().Changed(downloadLimitEnabledFlagName) {

		var downloadLimitEnabledFlagName string
		if cmdPrefix == "" {
			downloadLimitEnabledFlagName = "downloadLimitEnabled"
		} else {
			downloadLimitEnabledFlagName = fmt.Sprintf("%v.downloadLimitEnabled", cmdPrefix)
		}

		downloadLimitEnabledFlagValue, err := cmd.Flags().GetBool(downloadLimitEnabledFlagName)
		if err != nil {
			return err, false
		}
		m.DownloadLimitEnabled = &downloadLimitEnabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuNetworkFirewallEnabled6Flags(depth int, m *models.OnuNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	firewallEnabled6FlagName := fmt.Sprintf("%v.firewallEnabled6", cmdPrefix)
	if cmd.Flags().Changed(firewallEnabled6FlagName) {

		var firewallEnabled6FlagName string
		if cmdPrefix == "" {
			firewallEnabled6FlagName = "firewallEnabled6"
		} else {
			firewallEnabled6FlagName = fmt.Sprintf("%v.firewallEnabled6", cmdPrefix)
		}

		firewallEnabled6FlagValue, err := cmd.Flags().GetBool(firewallEnabled6FlagName)
		if err != nil {
			return err, false
		}
		m.FirewallEnabled6 = &firewallEnabled6FlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuNetworkIPV4Flags(depth int, m *models.OnuNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ipv4FlagName := fmt.Sprintf("%v.ipv4", cmdPrefix)
	if cmd.Flags().Changed(ipv4FlagName) {
		// info: complex object ipv4 IPV4 is retrieved outside this Changed() block
	}
	ipv4FlagValue := m.IPV4
	if swag.IsZero(ipv4FlagValue) {
		ipv4FlagValue = &models.IPV4{}
	}

	err, ipv4Added := retrieveModelIPV4Flags(depth+1, ipv4FlagValue, ipv4FlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipv4Added
	if ipv4Added {
		m.IPV4 = ipv4FlagValue
	}

	return nil, retAdded
}

func retrieveOnuNetworkIPV6Flags(depth int, m *models.OnuNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ipv6FlagName := fmt.Sprintf("%v.ipv6", cmdPrefix)
	if cmd.Flags().Changed(ipv6FlagName) {
		// info: complex object ipv6 IPV6 is retrieved outside this Changed() block
	}
	ipv6FlagValue := m.IPV6
	if swag.IsZero(ipv6FlagValue) {
		ipv6FlagValue = &models.IPV6{}
	}

	err, ipv6Added := retrieveModelIPV6Flags(depth+1, ipv6FlagValue, ipv6FlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipv6Added
	if ipv6Added {
		m.IPV6 = ipv6FlagValue
	}

	return nil, retAdded
}

func retrieveOnuNetworkIsIPV6EnabledFlags(depth int, m *models.OnuNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	isIpv6EnabledFlagName := fmt.Sprintf("%v.isIpv6Enabled", cmdPrefix)
	if cmd.Flags().Changed(isIpv6EnabledFlagName) {

		var isIpv6EnabledFlagName string
		if cmdPrefix == "" {
			isIpv6EnabledFlagName = "isIpv6Enabled"
		} else {
			isIpv6EnabledFlagName = fmt.Sprintf("%v.isIpv6Enabled", cmdPrefix)
		}

		isIpv6EnabledFlagValue, err := cmd.Flags().GetBool(isIpv6EnabledFlagName)
		if err != nil {
			return err, false
		}
		m.IsIPV6Enabled = &isIpv6EnabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuNetworkLanAddressFlags(depth int, m *models.OnuNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.LanAddress = &lanAddressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuNetworkLanProvisionedFlags(depth int, m *models.OnuNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.LanProvisioned = &lanProvisionedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuNetworkModeFlags(depth int, m *models.OnuNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Mode = &modeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuNetworkNatFtpFlags(depth int, m *models.OnuNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	natFtpFlagName := fmt.Sprintf("%v.natFtp", cmdPrefix)
	if cmd.Flags().Changed(natFtpFlagName) {

		var natFtpFlagName string
		if cmdPrefix == "" {
			natFtpFlagName = "natFtp"
		} else {
			natFtpFlagName = fmt.Sprintf("%v.natFtp", cmdPrefix)
		}

		natFtpFlagValue, err := cmd.Flags().GetBool(natFtpFlagName)
		if err != nil {
			return err, false
		}
		m.NatFtp = &natFtpFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuNetworkNatPptpFlags(depth int, m *models.OnuNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	natPptpFlagName := fmt.Sprintf("%v.natPptp", cmdPrefix)
	if cmd.Flags().Changed(natPptpFlagName) {

		var natPptpFlagName string
		if cmdPrefix == "" {
			natPptpFlagName = "natPptp"
		} else {
			natPptpFlagName = fmt.Sprintf("%v.natPptp", cmdPrefix)
		}

		natPptpFlagValue, err := cmd.Flags().GetBool(natPptpFlagName)
		if err != nil {
			return err, false
		}
		m.NatPptp = &natPptpFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuNetworkNatRtspFlags(depth int, m *models.OnuNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	natRtspFlagName := fmt.Sprintf("%v.natRtsp", cmdPrefix)
	if cmd.Flags().Changed(natRtspFlagName) {

		var natRtspFlagName string
		if cmdPrefix == "" {
			natRtspFlagName = "natRtsp"
		} else {
			natRtspFlagName = fmt.Sprintf("%v.natRtsp", cmdPrefix)
		}

		natRtspFlagValue, err := cmd.Flags().GetBool(natRtspFlagName)
		if err != nil {
			return err, false
		}
		m.NatRtsp = &natRtspFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuNetworkNatSipFlags(depth int, m *models.OnuNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	natSipFlagName := fmt.Sprintf("%v.natSip", cmdPrefix)
	if cmd.Flags().Changed(natSipFlagName) {

		var natSipFlagName string
		if cmdPrefix == "" {
			natSipFlagName = "natSip"
		} else {
			natSipFlagName = fmt.Sprintf("%v.natSip", cmdPrefix)
		}

		natSipFlagValue, err := cmd.Flags().GetBool(natSipFlagName)
		if err != nil {
			return err, false
		}
		m.NatSip = &natSipFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuNetworkPortForwardsFlags(depth int, m *models.OnuNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	portForwardsFlagName := fmt.Sprintf("%v.portForwards", cmdPrefix)
	if cmd.Flags().Changed(portForwardsFlagName) {
		// warning: portForwards array type PortForwards is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveOnuNetworkPortsFlags(depth int, m *models.OnuNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	portsFlagName := fmt.Sprintf("%v.ports", cmdPrefix)
	if cmd.Flags().Changed(portsFlagName) {
		// warning: ports array type Ports5 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveOnuNetworkUploadLimitFlags(depth int, m *models.OnuNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	uploadLimitFlagName := fmt.Sprintf("%v.uploadLimit", cmdPrefix)
	if cmd.Flags().Changed(uploadLimitFlagName) {

		var uploadLimitFlagName string
		if cmdPrefix == "" {
			uploadLimitFlagName = "uploadLimit"
		} else {
			uploadLimitFlagName = fmt.Sprintf("%v.uploadLimit", cmdPrefix)
		}

		uploadLimitFlagValue, err := cmd.Flags().GetInt64(uploadLimitFlagName)
		if err != nil {
			return err, false
		}
		m.UploadLimit = &uploadLimitFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuNetworkUploadLimitEnabledFlags(depth int, m *models.OnuNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	uploadLimitEnabledFlagName := fmt.Sprintf("%v.uploadLimitEnabled", cmdPrefix)
	if cmd.Flags().Changed(uploadLimitEnabledFlagName) {

		var uploadLimitEnabledFlagName string
		if cmdPrefix == "" {
			uploadLimitEnabledFlagName = "uploadLimitEnabled"
		} else {
			uploadLimitEnabledFlagName = fmt.Sprintf("%v.uploadLimitEnabled", cmdPrefix)
		}

		uploadLimitEnabledFlagValue, err := cmd.Flags().GetBool(uploadLimitEnabledFlagName)
		if err != nil {
			return err, false
		}
		m.UploadLimitEnabled = &uploadLimitEnabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuNetworkUpnpEnabledFlags(depth int, m *models.OnuNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	upnpEnabledFlagName := fmt.Sprintf("%v.upnpEnabled", cmdPrefix)
	if cmd.Flags().Changed(upnpEnabledFlagName) {

		var upnpEnabledFlagName string
		if cmdPrefix == "" {
			upnpEnabledFlagName = "upnpEnabled"
		} else {
			upnpEnabledFlagName = fmt.Sprintf("%v.upnpEnabled", cmdPrefix)
		}

		upnpEnabledFlagValue, err := cmd.Flags().GetBool(upnpEnabledFlagName)
		if err != nil {
			return err, false
		}
		m.UpnpEnabled = &upnpEnabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuNetworkVlansFlags(depth int, m *models.OnuNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	vlansFlagName := fmt.Sprintf("%v.vlans", cmdPrefix)
	if cmd.Flags().Changed(vlansFlagName) {
		// warning: vlans array type Vlans3 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveOnuNetworkWanVlanFlags(depth int, m *models.OnuNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	wanVlanFlagName := fmt.Sprintf("%v.wanVlan", cmdPrefix)
	if cmd.Flags().Changed(wanVlanFlagName) {

		var wanVlanFlagName string
		if cmdPrefix == "" {
			wanVlanFlagName = "wanVlan"
		} else {
			wanVlanFlagName = fmt.Sprintf("%v.wanVlan", cmdPrefix)
		}

		wanVlanFlagValue, err := cmd.Flags().GetInt64(wanVlanFlagName)
		if err != nil {
			return err, false
		}
		m.WanVlan = &wanVlanFlagValue

		retAdded = true
	}

	return nil, retAdded
}
