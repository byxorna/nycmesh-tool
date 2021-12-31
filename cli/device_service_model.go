// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for DeviceService

// register flags to command
func registerModelDeviceServiceFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDeviceServiceDiscovery(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceServiceNtpClient(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceServiceSnmpAgent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceServiceSSHServer(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceServiceSystemLog(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceServiceTelnetServer(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceServiceWebServer(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceServiceDiscovery(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var discoveryFlagName string
	if cmdPrefix == "" {
		discoveryFlagName = "discovery"
	} else {
		discoveryFlagName = fmt.Sprintf("%v.discovery", cmdPrefix)
	}

	if err := registerModelDiscovery1Flags(depth+1, discoveryFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceServiceNtpClient(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var ntpClientFlagName string
	if cmdPrefix == "" {
		ntpClientFlagName = "ntpClient"
	} else {
		ntpClientFlagName = fmt.Sprintf("%v.ntpClient", cmdPrefix)
	}

	if err := registerModelNtpClientFlags(depth+1, ntpClientFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceServiceSnmpAgent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var snmpAgentFlagName string
	if cmdPrefix == "" {
		snmpAgentFlagName = "snmpAgent"
	} else {
		snmpAgentFlagName = fmt.Sprintf("%v.snmpAgent", cmdPrefix)
	}

	if err := registerModelSnmpAgentFlags(depth+1, snmpAgentFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceServiceSSHServer(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var sshServerFlagName string
	if cmdPrefix == "" {
		sshServerFlagName = "sshServer"
	} else {
		sshServerFlagName = fmt.Sprintf("%v.sshServer", cmdPrefix)
	}

	if err := registerModelSSHServerFlags(depth+1, sshServerFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceServiceSystemLog(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var systemLogFlagName string
	if cmdPrefix == "" {
		systemLogFlagName = "systemLog"
	} else {
		systemLogFlagName = fmt.Sprintf("%v.systemLog", cmdPrefix)
	}

	if err := registerModelSystemLogFlags(depth+1, systemLogFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceServiceTelnetServer(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var telnetServerFlagName string
	if cmdPrefix == "" {
		telnetServerFlagName = "telnetServer"
	} else {
		telnetServerFlagName = fmt.Sprintf("%v.telnetServer", cmdPrefix)
	}

	if err := registerModelTelnetServerFlags(depth+1, telnetServerFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceServiceWebServer(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var webServerFlagName string
	if cmdPrefix == "" {
		webServerFlagName = "webServer"
	} else {
		webServerFlagName = fmt.Sprintf("%v.webServer", cmdPrefix)
	}

	if err := registerModelWebServerFlags(depth+1, webServerFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDeviceServiceFlags(depth int, m *models.DeviceService, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, discoveryAdded := retrieveDeviceServiceDiscoveryFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || discoveryAdded

	err, ntpClientAdded := retrieveDeviceServiceNtpClientFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ntpClientAdded

	err, snmpAgentAdded := retrieveDeviceServiceSnmpAgentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || snmpAgentAdded

	err, sshServerAdded := retrieveDeviceServiceSSHServerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sshServerAdded

	err, systemLogAdded := retrieveDeviceServiceSystemLogFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || systemLogAdded

	err, telnetServerAdded := retrieveDeviceServiceTelnetServerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || telnetServerAdded

	err, webServerAdded := retrieveDeviceServiceWebServerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || webServerAdded

	return nil, retAdded
}

func retrieveDeviceServiceDiscoveryFlags(depth int, m *models.DeviceService, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	discoveryFlagName := fmt.Sprintf("%v.discovery", cmdPrefix)
	if cmd.Flags().Changed(discoveryFlagName) {
		// info: complex object discovery Discovery1 is retrieved outside this Changed() block
	}
	discoveryFlagValue := m.Discovery
	if swag.IsZero(discoveryFlagValue) {
		discoveryFlagValue = &models.Discovery1{}
	}

	err, discoveryAdded := retrieveModelDiscovery1Flags(depth+1, discoveryFlagValue, discoveryFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || discoveryAdded
	if discoveryAdded {
		m.Discovery = discoveryFlagValue
	}

	return nil, retAdded
}

func retrieveDeviceServiceNtpClientFlags(depth int, m *models.DeviceService, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ntpClientFlagName := fmt.Sprintf("%v.ntpClient", cmdPrefix)
	if cmd.Flags().Changed(ntpClientFlagName) {
		// info: complex object ntpClient NtpClient is retrieved outside this Changed() block
	}
	ntpClientFlagValue := m.NtpClient
	if swag.IsZero(ntpClientFlagValue) {
		ntpClientFlagValue = &models.NtpClient{}
	}

	err, ntpClientAdded := retrieveModelNtpClientFlags(depth+1, ntpClientFlagValue, ntpClientFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ntpClientAdded
	if ntpClientAdded {
		m.NtpClient = ntpClientFlagValue
	}

	return nil, retAdded
}

func retrieveDeviceServiceSnmpAgentFlags(depth int, m *models.DeviceService, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	snmpAgentFlagName := fmt.Sprintf("%v.snmpAgent", cmdPrefix)
	if cmd.Flags().Changed(snmpAgentFlagName) {
		// info: complex object snmpAgent SnmpAgent is retrieved outside this Changed() block
	}
	snmpAgentFlagValue := m.SnmpAgent
	if swag.IsZero(snmpAgentFlagValue) {
		snmpAgentFlagValue = &models.SnmpAgent{}
	}

	err, snmpAgentAdded := retrieveModelSnmpAgentFlags(depth+1, snmpAgentFlagValue, snmpAgentFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || snmpAgentAdded
	if snmpAgentAdded {
		m.SnmpAgent = snmpAgentFlagValue
	}

	return nil, retAdded
}

func retrieveDeviceServiceSSHServerFlags(depth int, m *models.DeviceService, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sshServerFlagName := fmt.Sprintf("%v.sshServer", cmdPrefix)
	if cmd.Flags().Changed(sshServerFlagName) {
		// info: complex object sshServer SSHServer is retrieved outside this Changed() block
	}
	sshServerFlagValue := m.SSHServer
	if swag.IsZero(sshServerFlagValue) {
		sshServerFlagValue = &models.SSHServer{}
	}

	err, sshServerAdded := retrieveModelSSHServerFlags(depth+1, sshServerFlagValue, sshServerFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sshServerAdded
	if sshServerAdded {
		m.SSHServer = sshServerFlagValue
	}

	return nil, retAdded
}

func retrieveDeviceServiceSystemLogFlags(depth int, m *models.DeviceService, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	systemLogFlagName := fmt.Sprintf("%v.systemLog", cmdPrefix)
	if cmd.Flags().Changed(systemLogFlagName) {
		// info: complex object systemLog SystemLog is retrieved outside this Changed() block
	}
	systemLogFlagValue := m.SystemLog
	if swag.IsZero(systemLogFlagValue) {
		systemLogFlagValue = &models.SystemLog{}
	}

	err, systemLogAdded := retrieveModelSystemLogFlags(depth+1, systemLogFlagValue, systemLogFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || systemLogAdded
	if systemLogAdded {
		m.SystemLog = systemLogFlagValue
	}

	return nil, retAdded
}

func retrieveDeviceServiceTelnetServerFlags(depth int, m *models.DeviceService, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	telnetServerFlagName := fmt.Sprintf("%v.telnetServer", cmdPrefix)
	if cmd.Flags().Changed(telnetServerFlagName) {
		// info: complex object telnetServer TelnetServer is retrieved outside this Changed() block
	}
	telnetServerFlagValue := m.TelnetServer
	if swag.IsZero(telnetServerFlagValue) {
		telnetServerFlagValue = &models.TelnetServer{}
	}

	err, telnetServerAdded := retrieveModelTelnetServerFlags(depth+1, telnetServerFlagValue, telnetServerFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || telnetServerAdded
	if telnetServerAdded {
		m.TelnetServer = telnetServerFlagValue
	}

	return nil, retAdded
}

func retrieveDeviceServiceWebServerFlags(depth int, m *models.DeviceService, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	webServerFlagName := fmt.Sprintf("%v.webServer", cmdPrefix)
	if cmd.Flags().Changed(webServerFlagName) {
		// info: complex object webServer WebServer is retrieved outside this Changed() block
	}
	webServerFlagValue := m.WebServer
	if swag.IsZero(webServerFlagValue) {
		webServerFlagValue = &models.WebServer{}
	}

	err, webServerAdded := retrieveModelWebServerFlags(depth+1, webServerFlagValue, webServerFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || webServerAdded
	if webServerAdded {
		m.WebServer = webServerFlagValue
	}

	return nil, retAdded
}
