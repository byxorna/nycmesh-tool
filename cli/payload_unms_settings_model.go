// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
  "github.com/byxorna/nycmesh-tool/models"

	"github.com/spf13/cobra"
)

// Schema cli for PayloadUnmsSettings

// register flags to command
func registerModelPayloadUnmsSettingsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPayloadUnmsSettingsDeviceGracePeriodOutage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPayloadUnmsSettingsDevicePingAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPayloadUnmsSettingsDevicePingIntervalNormal(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPayloadUnmsSettingsDevicePingIntervalOutage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPayloadUnmsSettingsDeviceTransmissionFrequency(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPayloadUnmsSettingsMeta(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPayloadUnmsSettingsOverrideGlobal(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPayloadUnmsSettingsDeviceGracePeriodOutage(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	deviceGracePeriodOutageDescription := ``

	var deviceGracePeriodOutageFlagName string
	if cmdPrefix == "" {
		deviceGracePeriodOutageFlagName = "deviceGracePeriodOutage"
	} else {
		deviceGracePeriodOutageFlagName = fmt.Sprintf("%v.deviceGracePeriodOutage", cmdPrefix)
	}

	var deviceGracePeriodOutageFlagDefault float64

	_ = cmd.PersistentFlags().Float64(deviceGracePeriodOutageFlagName, deviceGracePeriodOutageFlagDefault, deviceGracePeriodOutageDescription)

	return nil
}

func registerPayloadUnmsSettingsDevicePingAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	devicePingAddressDescription := ``

	var devicePingAddressFlagName string
	if cmdPrefix == "" {
		devicePingAddressFlagName = "devicePingAddress"
	} else {
		devicePingAddressFlagName = fmt.Sprintf("%v.devicePingAddress", cmdPrefix)
	}

	var devicePingAddressFlagDefault string

	_ = cmd.PersistentFlags().String(devicePingAddressFlagName, devicePingAddressFlagDefault, devicePingAddressDescription)

	return nil
}

func registerPayloadUnmsSettingsDevicePingIntervalNormal(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	devicePingIntervalNormalDescription := `In milliseconds.`

	var devicePingIntervalNormalFlagName string
	if cmdPrefix == "" {
		devicePingIntervalNormalFlagName = "devicePingIntervalNormal"
	} else {
		devicePingIntervalNormalFlagName = fmt.Sprintf("%v.devicePingIntervalNormal", cmdPrefix)
	}

	var devicePingIntervalNormalFlagDefault float64

	_ = cmd.PersistentFlags().Float64(devicePingIntervalNormalFlagName, devicePingIntervalNormalFlagDefault, devicePingIntervalNormalDescription)

	return nil
}

func registerPayloadUnmsSettingsDevicePingIntervalOutage(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	devicePingIntervalOutageDescription := `In milliseconds.`

	var devicePingIntervalOutageFlagName string
	if cmdPrefix == "" {
		devicePingIntervalOutageFlagName = "devicePingIntervalOutage"
	} else {
		devicePingIntervalOutageFlagName = fmt.Sprintf("%v.devicePingIntervalOutage", cmdPrefix)
	}

	var devicePingIntervalOutageFlagDefault float64

	_ = cmd.PersistentFlags().Float64(devicePingIntervalOutageFlagName, devicePingIntervalOutageFlagDefault, devicePingIntervalOutageDescription)

	return nil
}

func registerPayloadUnmsSettingsDeviceTransmissionFrequency(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	deviceTransmissionFrequencyDescription := `Enum: ["minimal","low","medium","high","realtime"]. `

	var deviceTransmissionFrequencyFlagName string
	if cmdPrefix == "" {
		deviceTransmissionFrequencyFlagName = "deviceTransmissionFrequency"
	} else {
		deviceTransmissionFrequencyFlagName = fmt.Sprintf("%v.deviceTransmissionFrequency", cmdPrefix)
	}

	var deviceTransmissionFrequencyFlagDefault string

	_ = cmd.PersistentFlags().String(deviceTransmissionFrequencyFlagName, deviceTransmissionFrequencyFlagDefault, deviceTransmissionFrequencyDescription)

	if err := cmd.RegisterFlagCompletionFunc(deviceTransmissionFrequencyFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["minimal","low","medium","high","realtime"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerPayloadUnmsSettingsMeta(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var metaFlagName string
	if cmdPrefix == "" {
		metaFlagName = "meta"
	} else {
		metaFlagName = fmt.Sprintf("%v.meta", cmdPrefix)
	}

	if err := registerModelMeta1Flags(depth+1, metaFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerPayloadUnmsSettingsOverrideGlobal(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	overrideGlobalDescription := `Required. `

	var overrideGlobalFlagName string
	if cmdPrefix == "" {
		overrideGlobalFlagName = "overrideGlobal"
	} else {
		overrideGlobalFlagName = fmt.Sprintf("%v.overrideGlobal", cmdPrefix)
	}

	var overrideGlobalFlagDefault bool

	_ = cmd.PersistentFlags().Bool(overrideGlobalFlagName, overrideGlobalFlagDefault, overrideGlobalDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPayloadUnmsSettingsFlags(depth int, m *models.PayloadUnmsSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, deviceGracePeriodOutageAdded := retrievePayloadUnmsSettingsDeviceGracePeriodOutageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceGracePeriodOutageAdded

	err, devicePingAddressAdded := retrievePayloadUnmsSettingsDevicePingAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || devicePingAddressAdded

	err, devicePingIntervalNormalAdded := retrievePayloadUnmsSettingsDevicePingIntervalNormalFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || devicePingIntervalNormalAdded

	err, devicePingIntervalOutageAdded := retrievePayloadUnmsSettingsDevicePingIntervalOutageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || devicePingIntervalOutageAdded

	err, deviceTransmissionFrequencyAdded := retrievePayloadUnmsSettingsDeviceTransmissionFrequencyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceTransmissionFrequencyAdded

	err, metaAdded := retrievePayloadUnmsSettingsMetaFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || metaAdded

	err, overrideGlobalAdded := retrievePayloadUnmsSettingsOverrideGlobalFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || overrideGlobalAdded

	return nil, retAdded
}

func retrievePayloadUnmsSettingsDeviceGracePeriodOutageFlags(depth int, m *models.PayloadUnmsSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deviceGracePeriodOutageFlagName := fmt.Sprintf("%v.deviceGracePeriodOutage", cmdPrefix)
	if cmd.Flags().Changed(deviceGracePeriodOutageFlagName) {

		var deviceGracePeriodOutageFlagName string
		if cmdPrefix == "" {
			deviceGracePeriodOutageFlagName = "deviceGracePeriodOutage"
		} else {
			deviceGracePeriodOutageFlagName = fmt.Sprintf("%v.deviceGracePeriodOutage", cmdPrefix)
		}

		deviceGracePeriodOutageFlagValue, err := cmd.Flags().GetFloat64(deviceGracePeriodOutageFlagName)
		if err != nil {
			return err, false
		}
		m.DeviceGracePeriodOutage = deviceGracePeriodOutageFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePayloadUnmsSettingsDevicePingAddressFlags(depth int, m *models.PayloadUnmsSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	devicePingAddressFlagName := fmt.Sprintf("%v.devicePingAddress", cmdPrefix)
	if cmd.Flags().Changed(devicePingAddressFlagName) {

		var devicePingAddressFlagName string
		if cmdPrefix == "" {
			devicePingAddressFlagName = "devicePingAddress"
		} else {
			devicePingAddressFlagName = fmt.Sprintf("%v.devicePingAddress", cmdPrefix)
		}

		devicePingAddressFlagValue, err := cmd.Flags().GetString(devicePingAddressFlagName)
		if err != nil {
			return err, false
		}
		m.DevicePingAddress = devicePingAddressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePayloadUnmsSettingsDevicePingIntervalNormalFlags(depth int, m *models.PayloadUnmsSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	devicePingIntervalNormalFlagName := fmt.Sprintf("%v.devicePingIntervalNormal", cmdPrefix)
	if cmd.Flags().Changed(devicePingIntervalNormalFlagName) {

		var devicePingIntervalNormalFlagName string
		if cmdPrefix == "" {
			devicePingIntervalNormalFlagName = "devicePingIntervalNormal"
		} else {
			devicePingIntervalNormalFlagName = fmt.Sprintf("%v.devicePingIntervalNormal", cmdPrefix)
		}

		devicePingIntervalNormalFlagValue, err := cmd.Flags().GetFloat64(devicePingIntervalNormalFlagName)
		if err != nil {
			return err, false
		}
		m.DevicePingIntervalNormal = devicePingIntervalNormalFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePayloadUnmsSettingsDevicePingIntervalOutageFlags(depth int, m *models.PayloadUnmsSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	devicePingIntervalOutageFlagName := fmt.Sprintf("%v.devicePingIntervalOutage", cmdPrefix)
	if cmd.Flags().Changed(devicePingIntervalOutageFlagName) {

		var devicePingIntervalOutageFlagName string
		if cmdPrefix == "" {
			devicePingIntervalOutageFlagName = "devicePingIntervalOutage"
		} else {
			devicePingIntervalOutageFlagName = fmt.Sprintf("%v.devicePingIntervalOutage", cmdPrefix)
		}

		devicePingIntervalOutageFlagValue, err := cmd.Flags().GetFloat64(devicePingIntervalOutageFlagName)
		if err != nil {
			return err, false
		}
		m.DevicePingIntervalOutage = devicePingIntervalOutageFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePayloadUnmsSettingsDeviceTransmissionFrequencyFlags(depth int, m *models.PayloadUnmsSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deviceTransmissionFrequencyFlagName := fmt.Sprintf("%v.deviceTransmissionFrequency", cmdPrefix)
	if cmd.Flags().Changed(deviceTransmissionFrequencyFlagName) {

		var deviceTransmissionFrequencyFlagName string
		if cmdPrefix == "" {
			deviceTransmissionFrequencyFlagName = "deviceTransmissionFrequency"
		} else {
			deviceTransmissionFrequencyFlagName = fmt.Sprintf("%v.deviceTransmissionFrequency", cmdPrefix)
		}

		deviceTransmissionFrequencyFlagValue, err := cmd.Flags().GetString(deviceTransmissionFrequencyFlagName)
		if err != nil {
			return err, false
		}
		m.DeviceTransmissionFrequency = deviceTransmissionFrequencyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePayloadUnmsSettingsMetaFlags(depth int, m *models.PayloadUnmsSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	metaFlagName := fmt.Sprintf("%v.meta", cmdPrefix)
	if cmd.Flags().Changed(metaFlagName) {
		// info: complex object meta Meta1 is retrieved outside this Changed() block
	}
	metaFlagValue := m.Meta
	if swag.IsZero(metaFlagValue) {
		metaFlagValue = &models.Meta1{}
	}

	err, metaAdded := retrieveModelMeta1Flags(depth+1, metaFlagValue, metaFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || metaAdded
	if metaAdded {
		m.Meta = metaFlagValue
	}

	return nil, retAdded
}

func retrievePayloadUnmsSettingsOverrideGlobalFlags(depth int, m *models.PayloadUnmsSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	overrideGlobalFlagName := fmt.Sprintf("%v.overrideGlobal", cmdPrefix)
	if cmd.Flags().Changed(overrideGlobalFlagName) {

		var overrideGlobalFlagName string
		if cmdPrefix == "" {
			overrideGlobalFlagName = "overrideGlobal"
		} else {
			overrideGlobalFlagName = fmt.Sprintf("%v.overrideGlobal", cmdPrefix)
		}

		overrideGlobalFlagValue, err := cmd.Flags().GetBool(overrideGlobalFlagName)
		if err != nil {
			return err, false
		}
		m.OverrideGlobal = &overrideGlobalFlagValue

		retAdded = true
	}

	return nil, retAdded
}