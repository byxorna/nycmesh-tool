// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for PayloadUnmsHostnameReachable

// register flags to command
func registerModelPayloadUnmsHostnameReachableFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPayloadUnmsHostnameReachableHostname(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPayloadUnmsHostnameReachablePort(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPayloadUnmsHostnameReachableHostname(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerPayloadUnmsHostnameReachablePort(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	portDescription := ``

	var portFlagName string
	if cmdPrefix == "" {
		portFlagName = "port"
	} else {
		portFlagName = fmt.Sprintf("%v.port", cmdPrefix)
	}

	var portFlagDefault int64

	_ = cmd.PersistentFlags().Int64(portFlagName, portFlagDefault, portDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPayloadUnmsHostnameReachableFlags(depth int, m *models.PayloadUnmsHostnameReachable, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, hostnameAdded := retrievePayloadUnmsHostnameReachableHostnameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || hostnameAdded

	err, portAdded := retrievePayloadUnmsHostnameReachablePortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || portAdded

	return nil, retAdded
}

func retrievePayloadUnmsHostnameReachableHostnameFlags(depth int, m *models.PayloadUnmsHostnameReachable, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrievePayloadUnmsHostnameReachablePortFlags(depth int, m *models.PayloadUnmsHostnameReachable, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	portFlagName := fmt.Sprintf("%v.port", cmdPrefix)
	if cmd.Flags().Changed(portFlagName) {

		var portFlagName string
		if cmdPrefix == "" {
			portFlagName = "port"
		} else {
			portFlagName = fmt.Sprintf("%v.port", cmdPrefix)
		}

		portFlagValue, err := cmd.Flags().GetInt64(portFlagName)
		if err != nil {
			return err, false
		}
		m.Port = &portFlagValue

		retAdded = true
	}

	return nil, retAdded
}
