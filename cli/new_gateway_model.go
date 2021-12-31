// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for NewGateway

// register flags to command
func registerModelNewGatewayFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerNewGatewayConnectivityIPQueue(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNewGatewayDefaultIPQueue(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNewGatewayDevice(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNewGatewayForceOverwrite(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNewGatewayNetflowAlerts(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNewGatewayNetflowEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNewGatewayQosEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNewGatewaySuspend(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNewGatewaySuspendAllowedIps(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerNewGatewayConnectivityIPQueue(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var connectivityIpQueueFlagName string
	if cmdPrefix == "" {
		connectivityIpQueueFlagName = "connectivityIpQueue"
	} else {
		connectivityIpQueueFlagName = fmt.Sprintf("%v.connectivityIpQueue", cmdPrefix)
	}

	if err := registerModelConnectivityIPQueueFlags(depth+1, connectivityIpQueueFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerNewGatewayDefaultIPQueue(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var defaultIpQueueFlagName string
	if cmdPrefix == "" {
		defaultIpQueueFlagName = "defaultIpQueue"
	} else {
		defaultIpQueueFlagName = fmt.Sprintf("%v.defaultIpQueue", cmdPrefix)
	}

	if err := registerModelDefaultIPQueueFlags(depth+1, defaultIpQueueFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerNewGatewayDevice(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var deviceFlagName string
	if cmdPrefix == "" {
		deviceFlagName = "device"
	} else {
		deviceFlagName = fmt.Sprintf("%v.device", cmdPrefix)
	}

	if err := registerModelNewGatewayDeviceFlags(depth+1, deviceFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerNewGatewayForceOverwrite(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	forceOverwriteDescription := `Set to true to force write new NetFlow settings to device over current device's settings.`

	var forceOverwriteFlagName string
	if cmdPrefix == "" {
		forceOverwriteFlagName = "forceOverwrite"
	} else {
		forceOverwriteFlagName = fmt.Sprintf("%v.forceOverwrite", cmdPrefix)
	}

	var forceOverwriteFlagDefault bool

	_ = cmd.PersistentFlags().Bool(forceOverwriteFlagName, forceOverwriteFlagDefault, forceOverwriteDescription)

	return nil
}

func registerNewGatewayNetflowAlerts(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	netflowAlertsDescription := `Required. Set to true to enable alerts.`

	var netflowAlertsFlagName string
	if cmdPrefix == "" {
		netflowAlertsFlagName = "netflowAlerts"
	} else {
		netflowAlertsFlagName = fmt.Sprintf("%v.netflowAlerts", cmdPrefix)
	}

	var netflowAlertsFlagDefault bool

	_ = cmd.PersistentFlags().Bool(netflowAlertsFlagName, netflowAlertsFlagDefault, netflowAlertsDescription)

	return nil
}

func registerNewGatewayNetflowEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	netflowEnabledDescription := `Required. Set to true to enable NetFlow.`

	var netflowEnabledFlagName string
	if cmdPrefix == "" {
		netflowEnabledFlagName = "netflowEnabled"
	} else {
		netflowEnabledFlagName = fmt.Sprintf("%v.netflowEnabled", cmdPrefix)
	}

	var netflowEnabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(netflowEnabledFlagName, netflowEnabledFlagDefault, netflowEnabledDescription)

	return nil
}

func registerNewGatewayQosEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	qosEnabledDescription := `Required. Set to true to enable traffic shaping.`

	var qosEnabledFlagName string
	if cmdPrefix == "" {
		qosEnabledFlagName = "qosEnabled"
	} else {
		qosEnabledFlagName = fmt.Sprintf("%v.qosEnabled", cmdPrefix)
	}

	var qosEnabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(qosEnabledFlagName, qosEnabledFlagDefault, qosEnabledDescription)

	return nil
}

func registerNewGatewaySuspend(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	suspendDescription := `Required. Set to true to enable suspend.`

	var suspendFlagName string
	if cmdPrefix == "" {
		suspendFlagName = "suspend"
	} else {
		suspendFlagName = fmt.Sprintf("%v.suspend", cmdPrefix)
	}

	var suspendFlagDefault bool

	_ = cmd.PersistentFlags().Bool(suspendFlagName, suspendFlagDefault, suspendDescription)

	return nil
}

func registerNewGatewaySuspendAllowedIps(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: suspendAllowedIps SuspendAllowedIps array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelNewGatewayFlags(depth int, m *models.NewGateway, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, connectivityIpQueueAdded := retrieveNewGatewayConnectivityIPQueueFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || connectivityIpQueueAdded

	err, defaultIpQueueAdded := retrieveNewGatewayDefaultIPQueueFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || defaultIpQueueAdded

	err, deviceAdded := retrieveNewGatewayDeviceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceAdded

	err, forceOverwriteAdded := retrieveNewGatewayForceOverwriteFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || forceOverwriteAdded

	err, netflowAlertsAdded := retrieveNewGatewayNetflowAlertsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || netflowAlertsAdded

	err, netflowEnabledAdded := retrieveNewGatewayNetflowEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || netflowEnabledAdded

	err, qosEnabledAdded := retrieveNewGatewayQosEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || qosEnabledAdded

	err, suspendAdded := retrieveNewGatewaySuspendFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || suspendAdded

	err, suspendAllowedIpsAdded := retrieveNewGatewaySuspendAllowedIpsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || suspendAllowedIpsAdded

	return nil, retAdded
}

func retrieveNewGatewayConnectivityIPQueueFlags(depth int, m *models.NewGateway, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	connectivityIpQueueFlagName := fmt.Sprintf("%v.connectivityIpQueue", cmdPrefix)
	if cmd.Flags().Changed(connectivityIpQueueFlagName) {
		// info: complex object connectivityIpQueue ConnectivityIPQueue is retrieved outside this Changed() block
	}
	connectivityIpQueueFlagValue := m.ConnectivityIPQueue
	if swag.IsZero(connectivityIpQueueFlagValue) {
		connectivityIpQueueFlagValue = &models.ConnectivityIPQueue{}
	}

	err, connectivityIpQueueAdded := retrieveModelConnectivityIPQueueFlags(depth+1, connectivityIpQueueFlagValue, connectivityIpQueueFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || connectivityIpQueueAdded
	if connectivityIpQueueAdded {
		m.ConnectivityIPQueue = connectivityIpQueueFlagValue
	}

	return nil, retAdded
}

func retrieveNewGatewayDefaultIPQueueFlags(depth int, m *models.NewGateway, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	defaultIpQueueFlagName := fmt.Sprintf("%v.defaultIpQueue", cmdPrefix)
	if cmd.Flags().Changed(defaultIpQueueFlagName) {
		// info: complex object defaultIpQueue DefaultIPQueue is retrieved outside this Changed() block
	}
	defaultIpQueueFlagValue := m.DefaultIPQueue
	if swag.IsZero(defaultIpQueueFlagValue) {
		defaultIpQueueFlagValue = &models.DefaultIPQueue{}
	}

	err, defaultIpQueueAdded := retrieveModelDefaultIPQueueFlags(depth+1, defaultIpQueueFlagValue, defaultIpQueueFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || defaultIpQueueAdded
	if defaultIpQueueAdded {
		m.DefaultIPQueue = defaultIpQueueFlagValue
	}

	return nil, retAdded
}

func retrieveNewGatewayDeviceFlags(depth int, m *models.NewGateway, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deviceFlagName := fmt.Sprintf("%v.device", cmdPrefix)
	if cmd.Flags().Changed(deviceFlagName) {
		// info: complex object device NewGatewayDevice is retrieved outside this Changed() block
	}
	deviceFlagValue := m.Device
	if swag.IsZero(deviceFlagValue) {
		deviceFlagValue = &models.NewGatewayDevice{}
	}

	err, deviceAdded := retrieveModelNewGatewayDeviceFlags(depth+1, deviceFlagValue, deviceFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceAdded
	if deviceAdded {
		m.Device = deviceFlagValue
	}

	return nil, retAdded
}

func retrieveNewGatewayForceOverwriteFlags(depth int, m *models.NewGateway, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	forceOverwriteFlagName := fmt.Sprintf("%v.forceOverwrite", cmdPrefix)
	if cmd.Flags().Changed(forceOverwriteFlagName) {

		var forceOverwriteFlagName string
		if cmdPrefix == "" {
			forceOverwriteFlagName = "forceOverwrite"
		} else {
			forceOverwriteFlagName = fmt.Sprintf("%v.forceOverwrite", cmdPrefix)
		}

		forceOverwriteFlagValue, err := cmd.Flags().GetBool(forceOverwriteFlagName)
		if err != nil {
			return err, false
		}
		m.ForceOverwrite = forceOverwriteFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveNewGatewayNetflowAlertsFlags(depth int, m *models.NewGateway, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	netflowAlertsFlagName := fmt.Sprintf("%v.netflowAlerts", cmdPrefix)
	if cmd.Flags().Changed(netflowAlertsFlagName) {

		var netflowAlertsFlagName string
		if cmdPrefix == "" {
			netflowAlertsFlagName = "netflowAlerts"
		} else {
			netflowAlertsFlagName = fmt.Sprintf("%v.netflowAlerts", cmdPrefix)
		}

		netflowAlertsFlagValue, err := cmd.Flags().GetBool(netflowAlertsFlagName)
		if err != nil {
			return err, false
		}
		m.NetflowAlerts = &netflowAlertsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveNewGatewayNetflowEnabledFlags(depth int, m *models.NewGateway, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	netflowEnabledFlagName := fmt.Sprintf("%v.netflowEnabled", cmdPrefix)
	if cmd.Flags().Changed(netflowEnabledFlagName) {

		var netflowEnabledFlagName string
		if cmdPrefix == "" {
			netflowEnabledFlagName = "netflowEnabled"
		} else {
			netflowEnabledFlagName = fmt.Sprintf("%v.netflowEnabled", cmdPrefix)
		}

		netflowEnabledFlagValue, err := cmd.Flags().GetBool(netflowEnabledFlagName)
		if err != nil {
			return err, false
		}
		m.NetflowEnabled = &netflowEnabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveNewGatewayQosEnabledFlags(depth int, m *models.NewGateway, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	qosEnabledFlagName := fmt.Sprintf("%v.qosEnabled", cmdPrefix)
	if cmd.Flags().Changed(qosEnabledFlagName) {

		var qosEnabledFlagName string
		if cmdPrefix == "" {
			qosEnabledFlagName = "qosEnabled"
		} else {
			qosEnabledFlagName = fmt.Sprintf("%v.qosEnabled", cmdPrefix)
		}

		qosEnabledFlagValue, err := cmd.Flags().GetBool(qosEnabledFlagName)
		if err != nil {
			return err, false
		}
		m.QosEnabled = &qosEnabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveNewGatewaySuspendFlags(depth int, m *models.NewGateway, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	suspendFlagName := fmt.Sprintf("%v.suspend", cmdPrefix)
	if cmd.Flags().Changed(suspendFlagName) {

		var suspendFlagName string
		if cmdPrefix == "" {
			suspendFlagName = "suspend"
		} else {
			suspendFlagName = fmt.Sprintf("%v.suspend", cmdPrefix)
		}

		suspendFlagValue, err := cmd.Flags().GetBool(suspendFlagName)
		if err != nil {
			return err, false
		}
		m.Suspend = &suspendFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveNewGatewaySuspendAllowedIpsFlags(depth int, m *models.NewGateway, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	suspendAllowedIpsFlagName := fmt.Sprintf("%v.suspendAllowedIps", cmdPrefix)
	if cmd.Flags().Changed(suspendAllowedIpsFlagName) {
		// warning: suspendAllowedIps array type SuspendAllowedIps is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
