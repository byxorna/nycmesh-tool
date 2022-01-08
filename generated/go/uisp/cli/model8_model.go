// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for Model8

// register flags to command
func registerModelModel8Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel8ConnectivityIPQueue(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel8DefaultIPQueue(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel8Device(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel8ForceOverwrite(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel8ID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel8LastTrafficSeen(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel8NetflowAlerts(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel8NetflowEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel8NetflowStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel8QosEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel8Suspend(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel8SuspendAllowedIps(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel8ConnectivityIPQueue(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModel8DefaultIPQueue(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModel8Device(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var deviceFlagName string
	if cmdPrefix == "" {
		deviceFlagName = "device"
	} else {
		deviceFlagName = fmt.Sprintf("%v.device", cmdPrefix)
	}

	if err := registerModelDeviceFlags(depth+1, deviceFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel8ForceOverwrite(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModel8ID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `Required. Gateway id`

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

func registerModel8LastTrafficSeen(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	lastTrafficSeenDescription := `Required. Date and time of last traffic in ISO format`

	var lastTrafficSeenFlagName string
	if cmdPrefix == "" {
		lastTrafficSeenFlagName = "lastTrafficSeen"
	} else {
		lastTrafficSeenFlagName = fmt.Sprintf("%v.lastTrafficSeen", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(lastTrafficSeenFlagName, "", lastTrafficSeenDescription)

	return nil
}

func registerModel8NetflowAlerts(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModel8NetflowEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModel8NetflowStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	netflowStatusDescription := `Enum: ["active","connecting","discovered","inactive","disabled","disconnected","unauthorized","proposed","unknown","unplaced","custom"]. Netflow status`

	var netflowStatusFlagName string
	if cmdPrefix == "" {
		netflowStatusFlagName = "netflowStatus"
	} else {
		netflowStatusFlagName = fmt.Sprintf("%v.netflowStatus", cmdPrefix)
	}

	var netflowStatusFlagDefault string

	_ = cmd.PersistentFlags().String(netflowStatusFlagName, netflowStatusFlagDefault, netflowStatusDescription)

	if err := cmd.RegisterFlagCompletionFunc(netflowStatusFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["active","connecting","discovered","inactive","disabled","disconnected","unauthorized","proposed","unknown","unplaced","custom"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerModel8QosEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModel8Suspend(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModel8SuspendAllowedIps(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: suspendAllowedIps SuspendAllowedIps array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModel8Flags(depth int, m *models.Model8, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, connectivityIpQueueAdded := retrieveModel8ConnectivityIPQueueFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || connectivityIpQueueAdded

	err, defaultIpQueueAdded := retrieveModel8DefaultIPQueueFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || defaultIpQueueAdded

	err, deviceAdded := retrieveModel8DeviceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceAdded

	err, forceOverwriteAdded := retrieveModel8ForceOverwriteFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || forceOverwriteAdded

	err, idAdded := retrieveModel8IDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, lastTrafficSeenAdded := retrieveModel8LastTrafficSeenFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || lastTrafficSeenAdded

	err, netflowAlertsAdded := retrieveModel8NetflowAlertsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || netflowAlertsAdded

	err, netflowEnabledAdded := retrieveModel8NetflowEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || netflowEnabledAdded

	err, netflowStatusAdded := retrieveModel8NetflowStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || netflowStatusAdded

	err, qosEnabledAdded := retrieveModel8QosEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || qosEnabledAdded

	err, suspendAdded := retrieveModel8SuspendFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || suspendAdded

	err, suspendAllowedIpsAdded := retrieveModel8SuspendAllowedIpsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || suspendAllowedIpsAdded

	return nil, retAdded
}

func retrieveModel8ConnectivityIPQueueFlags(depth int, m *models.Model8, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel8DefaultIPQueueFlags(depth int, m *models.Model8, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel8DeviceFlags(depth int, m *models.Model8, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deviceFlagName := fmt.Sprintf("%v.device", cmdPrefix)
	if cmd.Flags().Changed(deviceFlagName) {
		// info: complex object device Device is retrieved outside this Changed() block
	}
	deviceFlagValue := m.Device
	if swag.IsZero(deviceFlagValue) {
		deviceFlagValue = &models.Device{}
	}

	err, deviceAdded := retrieveModelDeviceFlags(depth+1, deviceFlagValue, deviceFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceAdded
	if deviceAdded {
		m.Device = deviceFlagValue
	}

	return nil, retAdded
}

func retrieveModel8ForceOverwriteFlags(depth int, m *models.Model8, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel8IDFlags(depth int, m *models.Model8, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel8LastTrafficSeenFlags(depth int, m *models.Model8, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	lastTrafficSeenFlagName := fmt.Sprintf("%v.lastTrafficSeen", cmdPrefix)
	if cmd.Flags().Changed(lastTrafficSeenFlagName) {

		var lastTrafficSeenFlagName string
		if cmdPrefix == "" {
			lastTrafficSeenFlagName = "lastTrafficSeen"
		} else {
			lastTrafficSeenFlagName = fmt.Sprintf("%v.lastTrafficSeen", cmdPrefix)
		}

		lastTrafficSeenFlagValueStr, err := cmd.Flags().GetString(lastTrafficSeenFlagName)
		if err != nil {
			return err, false
		}
		var lastTrafficSeenFlagValue strfmt.DateTime
		if err := lastTrafficSeenFlagValue.UnmarshalText([]byte(lastTrafficSeenFlagValueStr)); err != nil {
			return err, false
		}
		m.LastTrafficSeen = &lastTrafficSeenFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel8NetflowAlertsFlags(depth int, m *models.Model8, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel8NetflowEnabledFlags(depth int, m *models.Model8, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel8NetflowStatusFlags(depth int, m *models.Model8, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	netflowStatusFlagName := fmt.Sprintf("%v.netflowStatus", cmdPrefix)
	if cmd.Flags().Changed(netflowStatusFlagName) {

		var netflowStatusFlagName string
		if cmdPrefix == "" {
			netflowStatusFlagName = "netflowStatus"
		} else {
			netflowStatusFlagName = fmt.Sprintf("%v.netflowStatus", cmdPrefix)
		}

		netflowStatusFlagValue, err := cmd.Flags().GetString(netflowStatusFlagName)
		if err != nil {
			return err, false
		}
		m.NetflowStatus = netflowStatusFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel8QosEnabledFlags(depth int, m *models.Model8, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel8SuspendFlags(depth int, m *models.Model8, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel8SuspendAllowedIpsFlags(depth int, m *models.Model8, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
