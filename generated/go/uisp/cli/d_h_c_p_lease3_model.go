// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for DHCPLease3

// register flags to command
func registerModelDHCPLease3Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDHCPLease3Address(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDHCPLease3Expiration(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDHCPLease3Hostname(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDHCPLease3LeaseID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDHCPLease3Mac(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDHCPLease3ServerName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDHCPLease3Type(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDHCPLease3Address(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	addressDescription := `Required. `

	var addressFlagName string
	if cmdPrefix == "" {
		addressFlagName = "address"
	} else {
		addressFlagName = fmt.Sprintf("%v.address", cmdPrefix)
	}

	var addressFlagDefault string

	_ = cmd.PersistentFlags().String(addressFlagName, addressFlagDefault, addressDescription)

	return nil
}

func registerDHCPLease3Expiration(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	expirationDescription := `expiration date in ISO format (UNMS-420); Nullable string.`

	var expirationFlagName string
	if cmdPrefix == "" {
		expirationFlagName = "expiration"
	} else {
		expirationFlagName = fmt.Sprintf("%v.expiration", cmdPrefix)
	}

	var expirationFlagDefault string

	_ = cmd.PersistentFlags().String(expirationFlagName, expirationFlagDefault, expirationDescription)

	return nil
}

func registerDHCPLease3Hostname(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerDHCPLease3LeaseID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	leaseIdDescription := `Unique per dhcp server.`

	var leaseIdFlagName string
	if cmdPrefix == "" {
		leaseIdFlagName = "leaseId"
	} else {
		leaseIdFlagName = fmt.Sprintf("%v.leaseId", cmdPrefix)
	}

	var leaseIdFlagDefault string

	_ = cmd.PersistentFlags().String(leaseIdFlagName, leaseIdFlagDefault, leaseIdDescription)

	return nil
}

func registerDHCPLease3Mac(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	macDescription := `Required. `

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

func registerDHCPLease3ServerName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	serverNameDescription := `Required. `

	var serverNameFlagName string
	if cmdPrefix == "" {
		serverNameFlagName = "serverName"
	} else {
		serverNameFlagName = fmt.Sprintf("%v.serverName", cmdPrefix)
	}

	var serverNameFlagDefault string

	_ = cmd.PersistentFlags().String(serverNameFlagName, serverNameFlagDefault, serverNameDescription)

	return nil
}

func registerDHCPLease3Type(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := ``

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
func retrieveModelDHCPLease3Flags(depth int, m *models.DHCPLease3, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, addressAdded := retrieveDHCPLease3AddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || addressAdded

	err, expirationAdded := retrieveDHCPLease3ExpirationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || expirationAdded

	err, hostnameAdded := retrieveDHCPLease3HostnameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || hostnameAdded

	err, leaseIdAdded := retrieveDHCPLease3LeaseIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || leaseIdAdded

	err, macAdded := retrieveDHCPLease3MacFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || macAdded

	err, serverNameAdded := retrieveDHCPLease3ServerNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || serverNameAdded

	err, typeAdded := retrieveDHCPLease3TypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveDHCPLease3AddressFlags(depth int, m *models.DHCPLease3, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	addressFlagName := fmt.Sprintf("%v.address", cmdPrefix)
	if cmd.Flags().Changed(addressFlagName) {

		var addressFlagName string
		if cmdPrefix == "" {
			addressFlagName = "address"
		} else {
			addressFlagName = fmt.Sprintf("%v.address", cmdPrefix)
		}

		addressFlagValue, err := cmd.Flags().GetString(addressFlagName)
		if err != nil {
			return err, false
		}
		m.Address = &addressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDHCPLease3ExpirationFlags(depth int, m *models.DHCPLease3, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	expirationFlagName := fmt.Sprintf("%v.expiration", cmdPrefix)
	if cmd.Flags().Changed(expirationFlagName) {

		var expirationFlagName string
		if cmdPrefix == "" {
			expirationFlagName = "expiration"
		} else {
			expirationFlagName = fmt.Sprintf("%v.expiration", cmdPrefix)
		}

		expirationFlagValue, err := cmd.Flags().GetString(expirationFlagName)
		if err != nil {
			return err, false
		}
		m.Expiration = expirationFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDHCPLease3HostnameFlags(depth int, m *models.DHCPLease3, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDHCPLease3LeaseIDFlags(depth int, m *models.DHCPLease3, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	leaseIdFlagName := fmt.Sprintf("%v.leaseId", cmdPrefix)
	if cmd.Flags().Changed(leaseIdFlagName) {

		var leaseIdFlagName string
		if cmdPrefix == "" {
			leaseIdFlagName = "leaseId"
		} else {
			leaseIdFlagName = fmt.Sprintf("%v.leaseId", cmdPrefix)
		}

		leaseIdFlagValue, err := cmd.Flags().GetString(leaseIdFlagName)
		if err != nil {
			return err, false
		}
		m.LeaseID = leaseIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDHCPLease3MacFlags(depth int, m *models.DHCPLease3, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDHCPLease3ServerNameFlags(depth int, m *models.DHCPLease3, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	serverNameFlagName := fmt.Sprintf("%v.serverName", cmdPrefix)
	if cmd.Flags().Changed(serverNameFlagName) {

		var serverNameFlagName string
		if cmdPrefix == "" {
			serverNameFlagName = "serverName"
		} else {
			serverNameFlagName = fmt.Sprintf("%v.serverName", cmdPrefix)
		}

		serverNameFlagValue, err := cmd.Flags().GetString(serverNameFlagName)
		if err != nil {
			return err, false
		}
		m.ServerName = &serverNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDHCPLease3TypeFlags(depth int, m *models.DHCPLease3, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
