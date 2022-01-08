// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for UpdateDHCPServer1

// register flags to command
func registerModelUpdateDHCPServer1Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUpdateDHCPServer1Available(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServer1Dns1(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServer1Dns2(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServer1Domain(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServer1Enabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServer1Interface(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServer1LeaseTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServer1Name(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServer1PoolSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServer1RangeEnd(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServer1RangeStart(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServer1Router(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateDHCPServer1UnifiController(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUpdateDHCPServer1Available(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerUpdateDHCPServer1Dns1(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerUpdateDHCPServer1Dns2(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerUpdateDHCPServer1Domain(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerUpdateDHCPServer1Enabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerUpdateDHCPServer1Interface(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	interfaceDescription := `Required. Interface id.`

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

func registerUpdateDHCPServer1LeaseTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	leaseTimeDescription := `DHCP lease time in seconds.`

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

func registerUpdateDHCPServer1Name(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerUpdateDHCPServer1PoolSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerUpdateDHCPServer1RangeEnd(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerUpdateDHCPServer1RangeStart(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerUpdateDHCPServer1Router(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerUpdateDHCPServer1UnifiController(depth int, cmdPrefix string, cmd *cobra.Command) error {
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
func retrieveModelUpdateDHCPServer1Flags(depth int, m *models.UpdateDHCPServer1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, availableAdded := retrieveUpdateDHCPServer1AvailableFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || availableAdded

	err, dns1Added := retrieveUpdateDHCPServer1Dns1Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dns1Added

	err, dns2Added := retrieveUpdateDHCPServer1Dns2Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dns2Added

	err, domainAdded := retrieveUpdateDHCPServer1DomainFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || domainAdded

	err, enabledAdded := retrieveUpdateDHCPServer1EnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || enabledAdded

	err, interfaceAdded := retrieveUpdateDHCPServer1InterfaceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || interfaceAdded

	err, leaseTimeAdded := retrieveUpdateDHCPServer1LeaseTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || leaseTimeAdded

	err, nameAdded := retrieveUpdateDHCPServer1NameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, poolSizeAdded := retrieveUpdateDHCPServer1PoolSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || poolSizeAdded

	err, rangeEndAdded := retrieveUpdateDHCPServer1RangeEndFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || rangeEndAdded

	err, rangeStartAdded := retrieveUpdateDHCPServer1RangeStartFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || rangeStartAdded

	err, routerAdded := retrieveUpdateDHCPServer1RouterFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || routerAdded

	err, unifiControllerAdded := retrieveUpdateDHCPServer1UnifiControllerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || unifiControllerAdded

	return nil, retAdded
}

func retrieveUpdateDHCPServer1AvailableFlags(depth int, m *models.UpdateDHCPServer1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveUpdateDHCPServer1Dns1Flags(depth int, m *models.UpdateDHCPServer1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveUpdateDHCPServer1Dns2Flags(depth int, m *models.UpdateDHCPServer1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveUpdateDHCPServer1DomainFlags(depth int, m *models.UpdateDHCPServer1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveUpdateDHCPServer1EnabledFlags(depth int, m *models.UpdateDHCPServer1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveUpdateDHCPServer1InterfaceFlags(depth int, m *models.UpdateDHCPServer1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveUpdateDHCPServer1LeaseTimeFlags(depth int, m *models.UpdateDHCPServer1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.LeaseTime = leaseTimeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUpdateDHCPServer1NameFlags(depth int, m *models.UpdateDHCPServer1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveUpdateDHCPServer1PoolSizeFlags(depth int, m *models.UpdateDHCPServer1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveUpdateDHCPServer1RangeEndFlags(depth int, m *models.UpdateDHCPServer1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveUpdateDHCPServer1RangeStartFlags(depth int, m *models.UpdateDHCPServer1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveUpdateDHCPServer1RouterFlags(depth int, m *models.UpdateDHCPServer1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveUpdateDHCPServer1UnifiControllerFlags(depth int, m *models.UpdateDHCPServer1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
