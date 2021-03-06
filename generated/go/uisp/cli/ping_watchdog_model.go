// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for PingWatchdog

// register flags to command
func registerModelPingWatchdogFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPingWatchdogAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPingWatchdogEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPingWatchdogFailureCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPingWatchdogInterval(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPingWatchdogOffDelay(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPingWatchdogStartDelay(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPingWatchdogAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	addressDescription := `Custom IP address in IPv4 or IPv6 format.`

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

func registerPingWatchdogEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerPingWatchdogFailureCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	failureCountDescription := ``

	var failureCountFlagName string
	if cmdPrefix == "" {
		failureCountFlagName = "failureCount"
	} else {
		failureCountFlagName = fmt.Sprintf("%v.failureCount", cmdPrefix)
	}

	var failureCountFlagDefault float64

	_ = cmd.PersistentFlags().Float64(failureCountFlagName, failureCountFlagDefault, failureCountDescription)

	return nil
}

func registerPingWatchdogInterval(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	intervalDescription := ``

	var intervalFlagName string
	if cmdPrefix == "" {
		intervalFlagName = "interval"
	} else {
		intervalFlagName = fmt.Sprintf("%v.interval", cmdPrefix)
	}

	var intervalFlagDefault float64

	_ = cmd.PersistentFlags().Float64(intervalFlagName, intervalFlagDefault, intervalDescription)

	return nil
}

func registerPingWatchdogOffDelay(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	offDelayDescription := ``

	var offDelayFlagName string
	if cmdPrefix == "" {
		offDelayFlagName = "offDelay"
	} else {
		offDelayFlagName = fmt.Sprintf("%v.offDelay", cmdPrefix)
	}

	var offDelayFlagDefault float64

	_ = cmd.PersistentFlags().Float64(offDelayFlagName, offDelayFlagDefault, offDelayDescription)

	return nil
}

func registerPingWatchdogStartDelay(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	startDelayDescription := ``

	var startDelayFlagName string
	if cmdPrefix == "" {
		startDelayFlagName = "startDelay"
	} else {
		startDelayFlagName = fmt.Sprintf("%v.startDelay", cmdPrefix)
	}

	var startDelayFlagDefault float64

	_ = cmd.PersistentFlags().Float64(startDelayFlagName, startDelayFlagDefault, startDelayDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPingWatchdogFlags(depth int, m *models.PingWatchdog, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, addressAdded := retrievePingWatchdogAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || addressAdded

	err, enabledAdded := retrievePingWatchdogEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || enabledAdded

	err, failureCountAdded := retrievePingWatchdogFailureCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || failureCountAdded

	err, intervalAdded := retrievePingWatchdogIntervalFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || intervalAdded

	err, offDelayAdded := retrievePingWatchdogOffDelayFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || offDelayAdded

	err, startDelayAdded := retrievePingWatchdogStartDelayFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || startDelayAdded

	return nil, retAdded
}

func retrievePingWatchdogAddressFlags(depth int, m *models.PingWatchdog, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Address = addressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePingWatchdogEnabledFlags(depth int, m *models.PingWatchdog, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrievePingWatchdogFailureCountFlags(depth int, m *models.PingWatchdog, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	failureCountFlagName := fmt.Sprintf("%v.failureCount", cmdPrefix)
	if cmd.Flags().Changed(failureCountFlagName) {

		var failureCountFlagName string
		if cmdPrefix == "" {
			failureCountFlagName = "failureCount"
		} else {
			failureCountFlagName = fmt.Sprintf("%v.failureCount", cmdPrefix)
		}

		failureCountFlagValue, err := cmd.Flags().GetFloat64(failureCountFlagName)
		if err != nil {
			return err, false
		}
		m.FailureCount = failureCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePingWatchdogIntervalFlags(depth int, m *models.PingWatchdog, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	intervalFlagName := fmt.Sprintf("%v.interval", cmdPrefix)
	if cmd.Flags().Changed(intervalFlagName) {

		var intervalFlagName string
		if cmdPrefix == "" {
			intervalFlagName = "interval"
		} else {
			intervalFlagName = fmt.Sprintf("%v.interval", cmdPrefix)
		}

		intervalFlagValue, err := cmd.Flags().GetFloat64(intervalFlagName)
		if err != nil {
			return err, false
		}
		m.Interval = intervalFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePingWatchdogOffDelayFlags(depth int, m *models.PingWatchdog, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	offDelayFlagName := fmt.Sprintf("%v.offDelay", cmdPrefix)
	if cmd.Flags().Changed(offDelayFlagName) {

		var offDelayFlagName string
		if cmdPrefix == "" {
			offDelayFlagName = "offDelay"
		} else {
			offDelayFlagName = fmt.Sprintf("%v.offDelay", cmdPrefix)
		}

		offDelayFlagValue, err := cmd.Flags().GetFloat64(offDelayFlagName)
		if err != nil {
			return err, false
		}
		m.OffDelay = offDelayFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePingWatchdogStartDelayFlags(depth int, m *models.PingWatchdog, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	startDelayFlagName := fmt.Sprintf("%v.startDelay", cmdPrefix)
	if cmd.Flags().Changed(startDelayFlagName) {

		var startDelayFlagName string
		if cmdPrefix == "" {
			startDelayFlagName = "startDelay"
		} else {
			startDelayFlagName = fmt.Sprintf("%v.startDelay", cmdPrefix)
		}

		startDelayFlagValue, err := cmd.Flags().GetFloat64(startDelayFlagName)
		if err != nil {
			return err, false
		}
		m.StartDelay = startDelayFlagValue

		retAdded = true
	}

	return nil, retAdded
}
