// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for UpdateDHCPServer

// register flags to command
func registerModelUpdateDHCPServerFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUpdateDHCPServerAvailable(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServerDns1(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServerDns2(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServerDomain(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServerEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServerInterface(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServerLeaseTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServerName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServerPoolSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServerRangeEnd(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServerRangeStart(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServerRouter(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServerUnifiController(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUpdateDHCPServerAvailable(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	availableDescription := `Available leases in DHCP pool.`

	var availableFlagName string
	if cmdPrefix == "" {
		availableFlagName = "available"
	} else {
		availableFlagName = fmt.Sprintf("%v.available", cmdPrefix)
	}

	var availableFlagDefault float64

	_ = cmd.PersistentFlags().Float64(availableFlagName, availableFlagDefault, availableDescription)

	return nil
}

func registerUpdateDHCPServerDns1(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	dns1Description := `Primary DNS server address.`

	var dns1FlagName string
	if cmdPrefix == "" {
		dns1FlagName = "dns1"
	} else {
		dns1FlagName = fmt.Sprintf("%v.dns1", cmdPrefix)
	}

	var dns1FlagDefault string

	_ = cmd.PersistentFlags().String(dns1FlagName, dns1FlagDefault, dns1Description)

	return nil
}

func registerUpdateDHCPServerDns2(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	dns2Description := `Secondary DNS server address.`

	var dns2FlagName string
	if cmdPrefix == "" {
		dns2FlagName = "dns2"
	} else {
		dns2FlagName = fmt.Sprintf("%v.dns2", cmdPrefix)
	}

	var dns2FlagDefault string

	_ = cmd.PersistentFlags().String(dns2FlagName, dns2FlagDefault, dns2Description)

	return nil
}

func registerUpdateDHCPServerDomain(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	domainDescription := `Domain name.`

	var domainFlagName string
	if cmdPrefix == "" {
		domainFlagName = "domain"
	} else {
		domainFlagName = fmt.Sprintf("%v.domain", cmdPrefix)
	}

	var domainFlagDefault string

	_ = cmd.PersistentFlags().String(domainFlagName, domainFlagDefault, domainDescription)

	return nil
}

func registerUpdateDHCPServerEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	enabledDescription := `Set to TRUE to enable DHCP server.`

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

func registerUpdateDHCPServerInterface(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	interfaceDescription := `Required. Interface IP v4 address in CIDR format.`

	var interfaceFlagName string
	if cmdPrefix == "" {
		interfaceFlagName = "interface"
	} else {
		interfaceFlagName = fmt.Sprintf("%v.interface", cmdPrefix)
	}

	var interfaceFlagDefault string

	_ = cmd.PersistentFlags().String(interfaceFlagName, interfaceFlagDefault, interfaceDescription)

	return nil
}

func registerUpdateDHCPServerLeaseTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	leaseTimeDescription := `Required. DHCP lease time in seconds.`

	var leaseTimeFlagName string
	if cmdPrefix == "" {
		leaseTimeFlagName = "leaseTime"
	} else {
		leaseTimeFlagName = fmt.Sprintf("%v.leaseTime", cmdPrefix)
	}

	var leaseTimeFlagDefault float64

	_ = cmd.PersistentFlags().Float64(leaseTimeFlagName, leaseTimeFlagDefault, leaseTimeDescription)

	return nil
}

func registerUpdateDHCPServerName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Required. DHCP server custom name.`

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

func registerUpdateDHCPServerPoolSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	poolSizeDescription := `Total range of leases in DHCP pool.`

	var poolSizeFlagName string
	if cmdPrefix == "" {
		poolSizeFlagName = "poolSize"
	} else {
		poolSizeFlagName = fmt.Sprintf("%v.poolSize", cmdPrefix)
	}

	var poolSizeFlagDefault float64

	_ = cmd.PersistentFlags().Float64(poolSizeFlagName, poolSizeFlagDefault, poolSizeDescription)

	return nil
}

func registerUpdateDHCPServerRangeEnd(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	rangeEndDescription := `Required. DHCP addresses pool end in CIDR format.`

	var rangeEndFlagName string
	if cmdPrefix == "" {
		rangeEndFlagName = "rangeEnd"
	} else {
		rangeEndFlagName = fmt.Sprintf("%v.rangeEnd", cmdPrefix)
	}

	var rangeEndFlagDefault string

	_ = cmd.PersistentFlags().String(rangeEndFlagName, rangeEndFlagDefault, rangeEndDescription)

	return nil
}

func registerUpdateDHCPServerRangeStart(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	rangeStartDescription := `Required. DHCP addresses pool start in CIDR format.`

	var rangeStartFlagName string
	if cmdPrefix == "" {
		rangeStartFlagName = "rangeStart"
	} else {
		rangeStartFlagName = fmt.Sprintf("%v.rangeStart", cmdPrefix)
	}

	var rangeStartFlagDefault string

	_ = cmd.PersistentFlags().String(rangeStartFlagName, rangeStartFlagDefault, rangeStartDescription)

	return nil
}

func registerUpdateDHCPServerRouter(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	routerDescription := `Router IP v4 address.`

	var routerFlagName string
	if cmdPrefix == "" {
		routerFlagName = "router"
	} else {
		routerFlagName = fmt.Sprintf("%v.router", cmdPrefix)
	}

	var routerFlagDefault string

	_ = cmd.PersistentFlags().String(routerFlagName, routerFlagDefault, routerDescription)

	return nil
}

func registerUpdateDHCPServerUnifiController(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	unifiControllerDescription := `Unifi controller IP v4 address.`

	var unifiControllerFlagName string
	if cmdPrefix == "" {
		unifiControllerFlagName = "unifiController"
	} else {
		unifiControllerFlagName = fmt.Sprintf("%v.unifiController", cmdPrefix)
	}

	var unifiControllerFlagDefault string

	_ = cmd.PersistentFlags().String(unifiControllerFlagName, unifiControllerFlagDefault, unifiControllerDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUpdateDHCPServerFlags(depth int, m *models.UpdateDHCPServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, availableAdded := retrieveUpdateDHCPServerAvailableFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || availableAdded

	err, dns1Added := retrieveUpdateDHCPServerDns1Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dns1Added

	err, dns2Added := retrieveUpdateDHCPServerDns2Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dns2Added

	err, domainAdded := retrieveUpdateDHCPServerDomainFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || domainAdded

	err, enabledAdded := retrieveUpdateDHCPServerEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || enabledAdded

	err, interfaceAdded := retrieveUpdateDHCPServerInterfaceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || interfaceAdded

	err, leaseTimeAdded := retrieveUpdateDHCPServerLeaseTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || leaseTimeAdded

	err, nameAdded := retrieveUpdateDHCPServerNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, poolSizeAdded := retrieveUpdateDHCPServerPoolSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || poolSizeAdded

	err, rangeEndAdded := retrieveUpdateDHCPServerRangeEndFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || rangeEndAdded

	err, rangeStartAdded := retrieveUpdateDHCPServerRangeStartFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || rangeStartAdded

	err, routerAdded := retrieveUpdateDHCPServerRouterFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || routerAdded

	err, unifiControllerAdded := retrieveUpdateDHCPServerUnifiControllerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || unifiControllerAdded

	return nil, retAdded
}

func retrieveUpdateDHCPServerAvailableFlags(depth int, m *models.UpdateDHCPServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	availableFlagName := fmt.Sprintf("%v.available", cmdPrefix)
	if cmd.Flags().Changed(availableFlagName) {

		var availableFlagName string
		if cmdPrefix == "" {
			availableFlagName = "available"
		} else {
			availableFlagName = fmt.Sprintf("%v.available", cmdPrefix)
		}

		availableFlagValue, err := cmd.Flags().GetFloat64(availableFlagName)
		if err != nil {
			return err, false
		}
		m.Available = &availableFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUpdateDHCPServerDns1Flags(depth int, m *models.UpdateDHCPServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dns1FlagName := fmt.Sprintf("%v.dns1", cmdPrefix)
	if cmd.Flags().Changed(dns1FlagName) {

		var dns1FlagName string
		if cmdPrefix == "" {
			dns1FlagName = "dns1"
		} else {
			dns1FlagName = fmt.Sprintf("%v.dns1", cmdPrefix)
		}

		dns1FlagValue, err := cmd.Flags().GetString(dns1FlagName)
		if err != nil {
			return err, false
		}
		m.Dns1 = dns1FlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUpdateDHCPServerDns2Flags(depth int, m *models.UpdateDHCPServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dns2FlagName := fmt.Sprintf("%v.dns2", cmdPrefix)
	if cmd.Flags().Changed(dns2FlagName) {

		var dns2FlagName string
		if cmdPrefix == "" {
			dns2FlagName = "dns2"
		} else {
			dns2FlagName = fmt.Sprintf("%v.dns2", cmdPrefix)
		}

		dns2FlagValue, err := cmd.Flags().GetString(dns2FlagName)
		if err != nil {
			return err, false
		}
		m.Dns2 = dns2FlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUpdateDHCPServerDomainFlags(depth int, m *models.UpdateDHCPServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	domainFlagName := fmt.Sprintf("%v.domain", cmdPrefix)
	if cmd.Flags().Changed(domainFlagName) {

		var domainFlagName string
		if cmdPrefix == "" {
			domainFlagName = "domain"
		} else {
			domainFlagName = fmt.Sprintf("%v.domain", cmdPrefix)
		}

		domainFlagValue, err := cmd.Flags().GetString(domainFlagName)
		if err != nil {
			return err, false
		}
		m.Domain = domainFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUpdateDHCPServerEnabledFlags(depth int, m *models.UpdateDHCPServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Enabled = enabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUpdateDHCPServerInterfaceFlags(depth int, m *models.UpdateDHCPServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	interfaceFlagName := fmt.Sprintf("%v.interface", cmdPrefix)
	if cmd.Flags().Changed(interfaceFlagName) {

		var interfaceFlagName string
		if cmdPrefix == "" {
			interfaceFlagName = "interface"
		} else {
			interfaceFlagName = fmt.Sprintf("%v.interface", cmdPrefix)
		}

		interfaceFlagValue, err := cmd.Flags().GetString(interfaceFlagName)
		if err != nil {
			return err, false
		}
		m.Interface = &interfaceFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUpdateDHCPServerLeaseTimeFlags(depth int, m *models.UpdateDHCPServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	leaseTimeFlagName := fmt.Sprintf("%v.leaseTime", cmdPrefix)
	if cmd.Flags().Changed(leaseTimeFlagName) {

		var leaseTimeFlagName string
		if cmdPrefix == "" {
			leaseTimeFlagName = "leaseTime"
		} else {
			leaseTimeFlagName = fmt.Sprintf("%v.leaseTime", cmdPrefix)
		}

		leaseTimeFlagValue, err := cmd.Flags().GetFloat64(leaseTimeFlagName)
		if err != nil {
			return err, false
		}
		m.LeaseTime = &leaseTimeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUpdateDHCPServerNameFlags(depth int, m *models.UpdateDHCPServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveUpdateDHCPServerPoolSizeFlags(depth int, m *models.UpdateDHCPServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	poolSizeFlagName := fmt.Sprintf("%v.poolSize", cmdPrefix)
	if cmd.Flags().Changed(poolSizeFlagName) {

		var poolSizeFlagName string
		if cmdPrefix == "" {
			poolSizeFlagName = "poolSize"
		} else {
			poolSizeFlagName = fmt.Sprintf("%v.poolSize", cmdPrefix)
		}

		poolSizeFlagValue, err := cmd.Flags().GetFloat64(poolSizeFlagName)
		if err != nil {
			return err, false
		}
		m.PoolSize = &poolSizeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUpdateDHCPServerRangeEndFlags(depth int, m *models.UpdateDHCPServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	rangeEndFlagName := fmt.Sprintf("%v.rangeEnd", cmdPrefix)
	if cmd.Flags().Changed(rangeEndFlagName) {

		var rangeEndFlagName string
		if cmdPrefix == "" {
			rangeEndFlagName = "rangeEnd"
		} else {
			rangeEndFlagName = fmt.Sprintf("%v.rangeEnd", cmdPrefix)
		}

		rangeEndFlagValue, err := cmd.Flags().GetString(rangeEndFlagName)
		if err != nil {
			return err, false
		}
		m.RangeEnd = &rangeEndFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUpdateDHCPServerRangeStartFlags(depth int, m *models.UpdateDHCPServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	rangeStartFlagName := fmt.Sprintf("%v.rangeStart", cmdPrefix)
	if cmd.Flags().Changed(rangeStartFlagName) {

		var rangeStartFlagName string
		if cmdPrefix == "" {
			rangeStartFlagName = "rangeStart"
		} else {
			rangeStartFlagName = fmt.Sprintf("%v.rangeStart", cmdPrefix)
		}

		rangeStartFlagValue, err := cmd.Flags().GetString(rangeStartFlagName)
		if err != nil {
			return err, false
		}
		m.RangeStart = &rangeStartFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUpdateDHCPServerRouterFlags(depth int, m *models.UpdateDHCPServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	routerFlagName := fmt.Sprintf("%v.router", cmdPrefix)
	if cmd.Flags().Changed(routerFlagName) {

		var routerFlagName string
		if cmdPrefix == "" {
			routerFlagName = "router"
		} else {
			routerFlagName = fmt.Sprintf("%v.router", cmdPrefix)
		}

		routerFlagValue, err := cmd.Flags().GetString(routerFlagName)
		if err != nil {
			return err, false
		}
		m.Router = routerFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUpdateDHCPServerUnifiControllerFlags(depth int, m *models.UpdateDHCPServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	unifiControllerFlagName := fmt.Sprintf("%v.unifiController", cmdPrefix)
	if cmd.Flags().Changed(unifiControllerFlagName) {

		var unifiControllerFlagName string
		if cmdPrefix == "" {
			unifiControllerFlagName = "unifiController"
		} else {
			unifiControllerFlagName = fmt.Sprintf("%v.unifiController", cmdPrefix)
		}

		unifiControllerFlagValue, err := cmd.Flags().GetString(unifiControllerFlagName)
		if err != nil {
			return err, false
		}
		m.UnifiController = unifiControllerFlagValue

		retAdded = true
	}

	return nil, retAdded
}
