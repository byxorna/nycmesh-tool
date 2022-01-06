// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for Device1

// register flags to command
func registerModelDevice1Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDevice1Authorized(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDevice1Connected(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDevice1Enabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDevice1ID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDevice1Identification(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDevice1IPAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDevice1Overview(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDevice1Authorized(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerDevice1Connected(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	connectedDescription := ``

	var connectedFlagName string
	if cmdPrefix == "" {
		connectedFlagName = "connected"
	} else {
		connectedFlagName = fmt.Sprintf("%v.connected", cmdPrefix)
	}

	var connectedFlagDefault bool

	_ = cmd.PersistentFlags().Bool(connectedFlagName, connectedFlagDefault, connectedDescription)

	return nil
}

func registerDevice1Enabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	enabledDescription := ``

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

func registerDevice1ID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `Required. `

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

func registerDevice1Identification(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var identificationFlagName string
	if cmdPrefix == "" {
		identificationFlagName = "identification"
	} else {
		identificationFlagName = fmt.Sprintf("%v.identification", cmdPrefix)
	}

	if err := registerModelIdentification3Flags(depth+1, identificationFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDevice1IPAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ipAddressDescription := `Custom IP address in IPv4 or IPv6 format.`

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

func registerDevice1Overview(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var overviewFlagName string
	if cmdPrefix == "" {
		overviewFlagName = "overview"
	} else {
		overviewFlagName = fmt.Sprintf("%v.overview", cmdPrefix)
	}

	if err := registerModelOverview1Flags(depth+1, overviewFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDevice1Flags(depth int, m *models.Device1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, authorizedAdded := retrieveDevice1AuthorizedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || authorizedAdded

	err, connectedAdded := retrieveDevice1ConnectedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || connectedAdded

	err, enabledAdded := retrieveDevice1EnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || enabledAdded

	err, idAdded := retrieveDevice1IDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, identificationAdded := retrieveDevice1IdentificationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || identificationAdded

	err, ipAddressAdded := retrieveDevice1IPAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipAddressAdded

	err, overviewAdded := retrieveDevice1OverviewFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || overviewAdded

	return nil, retAdded
}

func retrieveDevice1AuthorizedFlags(depth int, m *models.Device1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDevice1ConnectedFlags(depth int, m *models.Device1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	connectedFlagName := fmt.Sprintf("%v.connected", cmdPrefix)
	if cmd.Flags().Changed(connectedFlagName) {

		var connectedFlagName string
		if cmdPrefix == "" {
			connectedFlagName = "connected"
		} else {
			connectedFlagName = fmt.Sprintf("%v.connected", cmdPrefix)
		}

		connectedFlagValue, err := cmd.Flags().GetBool(connectedFlagName)
		if err != nil {
			return err, false
		}
		m.Connected = connectedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDevice1EnabledFlags(depth int, m *models.Device1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDevice1IDFlags(depth int, m *models.Device1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDevice1IdentificationFlags(depth int, m *models.Device1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	identificationFlagName := fmt.Sprintf("%v.identification", cmdPrefix)
	if cmd.Flags().Changed(identificationFlagName) {
		// info: complex object identification Identification3 is retrieved outside this Changed() block
	}
	identificationFlagValue := m.Identification
	if swag.IsZero(identificationFlagValue) {
		identificationFlagValue = &models.Identification3{}
	}

	err, identificationAdded := retrieveModelIdentification3Flags(depth+1, identificationFlagValue, identificationFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || identificationAdded
	if identificationAdded {
		m.Identification = identificationFlagValue
	}

	return nil, retAdded
}

func retrieveDevice1IPAddressFlags(depth int, m *models.Device1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.IPAddress = ipAddressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDevice1OverviewFlags(depth int, m *models.Device1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	overviewFlagName := fmt.Sprintf("%v.overview", cmdPrefix)
	if cmd.Flags().Changed(overviewFlagName) {
		// info: complex object overview Overview1 is retrieved outside this Changed() block
	}
	overviewFlagValue := m.Overview
	if swag.IsZero(overviewFlagValue) {
		overviewFlagValue = &models.Overview1{}
	}

	err, overviewAdded := retrieveModelOverview1Flags(depth+1, overviewFlagValue, overviewFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || overviewAdded
	if overviewAdded {
		m.Overview = overviewFlagValue
	}

	return nil, retAdded
}
