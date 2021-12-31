// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"
"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for Lag

// register flags to command
func registerModelLagFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerLagDhcpSnooping(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerLagIncludeVlans(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerLagLinkTrap(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerLagLoadBalance(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerLagMode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerLagPorts(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerLagStatic(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerLagStp(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerLagVlanNative(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerLagDhcpSnooping(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	dhcpSnoopingDescription := ``

	var dhcpSnoopingFlagName string
	if cmdPrefix == "" {
		dhcpSnoopingFlagName = "dhcpSnooping"
	} else {
		dhcpSnoopingFlagName = fmt.Sprintf("%v.dhcpSnooping", cmdPrefix)
	}

	var dhcpSnoopingFlagDefault bool

	_ = cmd.PersistentFlags().Bool(dhcpSnoopingFlagName, dhcpSnoopingFlagDefault, dhcpSnoopingDescription)

	return nil
}

func registerLagIncludeVlans(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	includeVlansDescription := ``

	var includeVlansFlagName string
	if cmdPrefix == "" {
		includeVlansFlagName = "includeVlans"
	} else {
		includeVlansFlagName = fmt.Sprintf("%v.includeVlans", cmdPrefix)
	}

	var includeVlansFlagDefault string

	_ = cmd.PersistentFlags().String(includeVlansFlagName, includeVlansFlagDefault, includeVlansDescription)

	return nil
}

func registerLagLinkTrap(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	linkTrapDescription := ``

	var linkTrapFlagName string
	if cmdPrefix == "" {
		linkTrapFlagName = "linkTrap"
	} else {
		linkTrapFlagName = fmt.Sprintf("%v.linkTrap", cmdPrefix)
	}

	var linkTrapFlagDefault bool

	_ = cmd.PersistentFlags().Bool(linkTrapFlagName, linkTrapFlagDefault, linkTrapDescription)

	return nil
}

func registerLagLoadBalance(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	loadBalanceDescription := ``

	var loadBalanceFlagName string
	if cmdPrefix == "" {
		loadBalanceFlagName = "loadBalance"
	} else {
		loadBalanceFlagName = fmt.Sprintf("%v.loadBalance", cmdPrefix)
	}

	var loadBalanceFlagDefault string

	_ = cmd.PersistentFlags().String(loadBalanceFlagName, loadBalanceFlagDefault, loadBalanceDescription)

	return nil
}

func registerLagMode(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerLagPorts(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: ports Ports array type is not supported by go-swagger cli yet

	return nil
}

func registerLagStatic(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	staticDescription := ``

	var staticFlagName string
	if cmdPrefix == "" {
		staticFlagName = "static"
	} else {
		staticFlagName = fmt.Sprintf("%v.static", cmdPrefix)
	}

	var staticFlagDefault bool

	_ = cmd.PersistentFlags().Bool(staticFlagName, staticFlagDefault, staticDescription)

	return nil
}

func registerLagStp(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var stpFlagName string
	if cmdPrefix == "" {
		stpFlagName = "stp"
	} else {
		stpFlagName = fmt.Sprintf("%v.stp", cmdPrefix)
	}

	if err := registerModelStpFlags(depth+1, stpFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerLagVlanNative(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	vlanNativeDescription := ``

	var vlanNativeFlagName string
	if cmdPrefix == "" {
		vlanNativeFlagName = "vlanNative"
	} else {
		vlanNativeFlagName = fmt.Sprintf("%v.vlanNative", cmdPrefix)
	}

	var vlanNativeFlagDefault float64

	_ = cmd.PersistentFlags().Float64(vlanNativeFlagName, vlanNativeFlagDefault, vlanNativeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelLagFlags(depth int, m *models.Lag, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dhcpSnoopingAdded := retrieveLagDhcpSnoopingFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dhcpSnoopingAdded

	err, includeVlansAdded := retrieveLagIncludeVlansFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || includeVlansAdded

	err, linkTrapAdded := retrieveLagLinkTrapFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || linkTrapAdded

	err, loadBalanceAdded := retrieveLagLoadBalanceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || loadBalanceAdded

	err, modeAdded := retrieveLagModeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || modeAdded

	err, portsAdded := retrieveLagPortsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || portsAdded

	err, staticAdded := retrieveLagStaticFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || staticAdded

	err, stpAdded := retrieveLagStpFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || stpAdded

	err, vlanNativeAdded := retrieveLagVlanNativeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || vlanNativeAdded

	return nil, retAdded
}

func retrieveLagDhcpSnoopingFlags(depth int, m *models.Lag, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dhcpSnoopingFlagName := fmt.Sprintf("%v.dhcpSnooping", cmdPrefix)
	if cmd.Flags().Changed(dhcpSnoopingFlagName) {

		var dhcpSnoopingFlagName string
		if cmdPrefix == "" {
			dhcpSnoopingFlagName = "dhcpSnooping"
		} else {
			dhcpSnoopingFlagName = fmt.Sprintf("%v.dhcpSnooping", cmdPrefix)
		}

		dhcpSnoopingFlagValue, err := cmd.Flags().GetBool(dhcpSnoopingFlagName)
		if err != nil {
			return err, false
		}
		m.DhcpSnooping = dhcpSnoopingFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveLagIncludeVlansFlags(depth int, m *models.Lag, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	includeVlansFlagName := fmt.Sprintf("%v.includeVlans", cmdPrefix)
	if cmd.Flags().Changed(includeVlansFlagName) {

		var includeVlansFlagName string
		if cmdPrefix == "" {
			includeVlansFlagName = "includeVlans"
		} else {
			includeVlansFlagName = fmt.Sprintf("%v.includeVlans", cmdPrefix)
		}

		includeVlansFlagValue, err := cmd.Flags().GetString(includeVlansFlagName)
		if err != nil {
			return err, false
		}
		m.IncludeVlans = includeVlansFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveLagLinkTrapFlags(depth int, m *models.Lag, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	linkTrapFlagName := fmt.Sprintf("%v.linkTrap", cmdPrefix)
	if cmd.Flags().Changed(linkTrapFlagName) {

		var linkTrapFlagName string
		if cmdPrefix == "" {
			linkTrapFlagName = "linkTrap"
		} else {
			linkTrapFlagName = fmt.Sprintf("%v.linkTrap", cmdPrefix)
		}

		linkTrapFlagValue, err := cmd.Flags().GetBool(linkTrapFlagName)
		if err != nil {
			return err, false
		}
		m.LinkTrap = linkTrapFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveLagLoadBalanceFlags(depth int, m *models.Lag, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	loadBalanceFlagName := fmt.Sprintf("%v.loadBalance", cmdPrefix)
	if cmd.Flags().Changed(loadBalanceFlagName) {

		var loadBalanceFlagName string
		if cmdPrefix == "" {
			loadBalanceFlagName = "loadBalance"
		} else {
			loadBalanceFlagName = fmt.Sprintf("%v.loadBalance", cmdPrefix)
		}

		loadBalanceFlagValue, err := cmd.Flags().GetString(loadBalanceFlagName)
		if err != nil {
			return err, false
		}
		m.LoadBalance = loadBalanceFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveLagModeFlags(depth int, m *models.Lag, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveLagPortsFlags(depth int, m *models.Lag, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	portsFlagName := fmt.Sprintf("%v.ports", cmdPrefix)
	if cmd.Flags().Changed(portsFlagName) {
		// warning: ports array type Ports is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveLagStaticFlags(depth int, m *models.Lag, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	staticFlagName := fmt.Sprintf("%v.static", cmdPrefix)
	if cmd.Flags().Changed(staticFlagName) {

		var staticFlagName string
		if cmdPrefix == "" {
			staticFlagName = "static"
		} else {
			staticFlagName = fmt.Sprintf("%v.static", cmdPrefix)
		}

		staticFlagValue, err := cmd.Flags().GetBool(staticFlagName)
		if err != nil {
			return err, false
		}
		m.Static = staticFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveLagStpFlags(depth int, m *models.Lag, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	stpFlagName := fmt.Sprintf("%v.stp", cmdPrefix)
	if cmd.Flags().Changed(stpFlagName) {
		// info: complex object stp Stp is retrieved outside this Changed() block
	}
	stpFlagValue := m.Stp
	if swag.IsZero(stpFlagValue) {
		stpFlagValue = &models.Stp{}
	}

	err, stpAdded := retrieveModelStpFlags(depth+1, stpFlagValue, stpFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || stpAdded
	if stpAdded {
		m.Stp = stpFlagValue
	}

	return nil, retAdded
}

func retrieveLagVlanNativeFlags(depth int, m *models.Lag, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	vlanNativeFlagName := fmt.Sprintf("%v.vlanNative", cmdPrefix)
	if cmd.Flags().Changed(vlanNativeFlagName) {

		var vlanNativeFlagName string
		if cmdPrefix == "" {
			vlanNativeFlagName = "vlanNative"
		} else {
			vlanNativeFlagName = fmt.Sprintf("%v.vlanNative", cmdPrefix)
		}

		vlanNativeFlagValue, err := cmd.Flags().GetFloat64(vlanNativeFlagName)
		if err != nil {
			return err, false
		}
		m.VlanNative = vlanNativeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
