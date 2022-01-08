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

// Schema cli for Wan1

// register flags to command
func registerModelWan1Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerWan1DNSResolvers(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWan1Gateway(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWan1PppoeMode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWan1PppoePassword(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWan1PppoeUser(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWan1WanAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWan1WanMode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerWan1DNSResolvers(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: dnsResolvers DNSResolvers array type is not supported by go-swagger cli yet

	return nil
}

func registerWan1Gateway(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	gatewayDescription := `Required. IPv4 address of gateway`

	var gatewayFlagName string
	if cmdPrefix == "" {
		gatewayFlagName = "gateway"
	} else {
		gatewayFlagName = fmt.Sprintf("%v.gateway", cmdPrefix)
	}

	var gatewayFlagDefault string

	_ = cmd.PersistentFlags().String(gatewayFlagName, gatewayFlagDefault, gatewayDescription)

	return nil
}

func registerWan1PppoeMode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	pppoeModeDescription := `Enum: ["auto","pap","chap"]. Required. PPPoE mode, required if WAN mode set to PPPoE`

	var pppoeModeFlagName string
	if cmdPrefix == "" {
		pppoeModeFlagName = "pppoeMode"
	} else {
		pppoeModeFlagName = fmt.Sprintf("%v.pppoeMode", cmdPrefix)
	}

	var pppoeModeFlagDefault string

	_ = cmd.PersistentFlags().String(pppoeModeFlagName, pppoeModeFlagDefault, pppoeModeDescription)

	if err := cmd.RegisterFlagCompletionFunc(pppoeModeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["auto","pap","chap"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerWan1PppoePassword(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	pppoePasswordDescription := `Required. PPPoE password`

	var pppoePasswordFlagName string
	if cmdPrefix == "" {
		pppoePasswordFlagName = "pppoePassword"
	} else {
		pppoePasswordFlagName = fmt.Sprintf("%v.pppoePassword", cmdPrefix)
	}

	var pppoePasswordFlagDefault string

	_ = cmd.PersistentFlags().String(pppoePasswordFlagName, pppoePasswordFlagDefault, pppoePasswordDescription)

	return nil
}

func registerWan1PppoeUser(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	pppoeUserDescription := `Required. PPPoE username`

	var pppoeUserFlagName string
	if cmdPrefix == "" {
		pppoeUserFlagName = "pppoeUser"
	} else {
		pppoeUserFlagName = fmt.Sprintf("%v.pppoeUser", cmdPrefix)
	}

	var pppoeUserFlagDefault string

	_ = cmd.PersistentFlags().String(pppoeUserFlagName, pppoeUserFlagDefault, pppoeUserDescription)

	return nil
}

func registerWan1WanAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	wanAddressDescription := `Required. WAN address of ONU in CIDR format - required if wanMode is set to "static"`

	var wanAddressFlagName string
	if cmdPrefix == "" {
		wanAddressFlagName = "wanAddress"
	} else {
		wanAddressFlagName = fmt.Sprintf("%v.wanAddress", cmdPrefix)
	}

	var wanAddressFlagDefault string

	_ = cmd.PersistentFlags().String(wanAddressFlagName, wanAddressFlagDefault, wanAddressDescription)

	return nil
}

func registerWan1WanMode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	wanModeDescription := `Enum: ["static","dhcp","pppoe"]. Required. WAN mode`

	var wanModeFlagName string
	if cmdPrefix == "" {
		wanModeFlagName = "wanMode"
	} else {
		wanModeFlagName = fmt.Sprintf("%v.wanMode", cmdPrefix)
	}

	var wanModeFlagDefault string

	_ = cmd.PersistentFlags().String(wanModeFlagName, wanModeFlagDefault, wanModeDescription)

	if err := cmd.RegisterFlagCompletionFunc(wanModeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["static","dhcp","pppoe"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelWan1Flags(depth int, m *models.Wan1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dnsResolversAdded := retrieveWan1DNSResolversFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dnsResolversAdded

	err, gatewayAdded := retrieveWan1GatewayFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || gatewayAdded

	err, pppoeModeAdded := retrieveWan1PppoeModeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pppoeModeAdded

	err, pppoePasswordAdded := retrieveWan1PppoePasswordFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pppoePasswordAdded

	err, pppoeUserAdded := retrieveWan1PppoeUserFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pppoeUserAdded

	err, wanAddressAdded := retrieveWan1WanAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || wanAddressAdded

	err, wanModeAdded := retrieveWan1WanModeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || wanModeAdded

	return nil, retAdded
}

func retrieveWan1DNSResolversFlags(depth int, m *models.Wan1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dnsResolversFlagName := fmt.Sprintf("%v.dnsResolvers", cmdPrefix)
	if cmd.Flags().Changed(dnsResolversFlagName) {
		// warning: dnsResolvers array type DNSResolvers is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveWan1GatewayFlags(depth int, m *models.Wan1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	gatewayFlagName := fmt.Sprintf("%v.gateway", cmdPrefix)
	if cmd.Flags().Changed(gatewayFlagName) {

		var gatewayFlagName string
		if cmdPrefix == "" {
			gatewayFlagName = "gateway"
		} else {
			gatewayFlagName = fmt.Sprintf("%v.gateway", cmdPrefix)
		}

		gatewayFlagValue, err := cmd.Flags().GetString(gatewayFlagName)
		if err != nil {
			return err, false
		}
		m.Gateway = &gatewayFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWan1PppoeModeFlags(depth int, m *models.Wan1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pppoeModeFlagName := fmt.Sprintf("%v.pppoeMode", cmdPrefix)
	if cmd.Flags().Changed(pppoeModeFlagName) {

		var pppoeModeFlagName string
		if cmdPrefix == "" {
			pppoeModeFlagName = "pppoeMode"
		} else {
			pppoeModeFlagName = fmt.Sprintf("%v.pppoeMode", cmdPrefix)
		}

		pppoeModeFlagValue, err := cmd.Flags().GetString(pppoeModeFlagName)
		if err != nil {
			return err, false
		}
		m.PppoeMode = &pppoeModeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWan1PppoePasswordFlags(depth int, m *models.Wan1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pppoePasswordFlagName := fmt.Sprintf("%v.pppoePassword", cmdPrefix)
	if cmd.Flags().Changed(pppoePasswordFlagName) {

		var pppoePasswordFlagName string
		if cmdPrefix == "" {
			pppoePasswordFlagName = "pppoePassword"
		} else {
			pppoePasswordFlagName = fmt.Sprintf("%v.pppoePassword", cmdPrefix)
		}

		pppoePasswordFlagValue, err := cmd.Flags().GetString(pppoePasswordFlagName)
		if err != nil {
			return err, false
		}
		m.PppoePassword = &pppoePasswordFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWan1PppoeUserFlags(depth int, m *models.Wan1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pppoeUserFlagName := fmt.Sprintf("%v.pppoeUser", cmdPrefix)
	if cmd.Flags().Changed(pppoeUserFlagName) {

		var pppoeUserFlagName string
		if cmdPrefix == "" {
			pppoeUserFlagName = "pppoeUser"
		} else {
			pppoeUserFlagName = fmt.Sprintf("%v.pppoeUser", cmdPrefix)
		}

		pppoeUserFlagValue, err := cmd.Flags().GetString(pppoeUserFlagName)
		if err != nil {
			return err, false
		}
		m.PppoeUser = &pppoeUserFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWan1WanAddressFlags(depth int, m *models.Wan1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	wanAddressFlagName := fmt.Sprintf("%v.wanAddress", cmdPrefix)
	if cmd.Flags().Changed(wanAddressFlagName) {

		var wanAddressFlagName string
		if cmdPrefix == "" {
			wanAddressFlagName = "wanAddress"
		} else {
			wanAddressFlagName = fmt.Sprintf("%v.wanAddress", cmdPrefix)
		}

		wanAddressFlagValue, err := cmd.Flags().GetString(wanAddressFlagName)
		if err != nil {
			return err, false
		}
		m.WanAddress = &wanAddressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWan1WanModeFlags(depth int, m *models.Wan1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	wanModeFlagName := fmt.Sprintf("%v.wanMode", cmdPrefix)
	if cmd.Flags().Changed(wanModeFlagName) {

		var wanModeFlagName string
		if cmdPrefix == "" {
			wanModeFlagName = "wanMode"
		} else {
			wanModeFlagName = fmt.Sprintf("%v.wanMode", cmdPrefix)
		}

		wanModeFlagValue, err := cmd.Flags().GetString(wanModeFlagName)
		if err != nil {
			return err, false
		}
		m.WanMode = &wanModeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
