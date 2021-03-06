// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for Model55

// register flags to command
func registerModelModel55Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel55DeviceName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel55NewPassword(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel55Timezone(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel55Username(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel55Zonename(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel55DeviceName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	deviceNameDescription := `Required. `

	var deviceNameFlagName string
	if cmdPrefix == "" {
		deviceNameFlagName = "deviceName"
	} else {
		deviceNameFlagName = fmt.Sprintf("%v.deviceName", cmdPrefix)
	}

	var deviceNameFlagDefault string

	_ = cmd.PersistentFlags().String(deviceNameFlagName, deviceNameFlagDefault, deviceNameDescription)

	return nil
}

func registerModel55NewPassword(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	newPasswordDescription := ``

	var newPasswordFlagName string
	if cmdPrefix == "" {
		newPasswordFlagName = "newPassword"
	} else {
		newPasswordFlagName = fmt.Sprintf("%v.newPassword", cmdPrefix)
	}

	var newPasswordFlagDefault string

	_ = cmd.PersistentFlags().String(newPasswordFlagName, newPasswordFlagDefault, newPasswordDescription)

	return nil
}

func registerModel55Timezone(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	timezoneDescription := `Required. `

	var timezoneFlagName string
	if cmdPrefix == "" {
		timezoneFlagName = "timezone"
	} else {
		timezoneFlagName = fmt.Sprintf("%v.timezone", cmdPrefix)
	}

	var timezoneFlagDefault string

	_ = cmd.PersistentFlags().String(timezoneFlagName, timezoneFlagDefault, timezoneDescription)

	return nil
}

func registerModel55Username(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	usernameDescription := `Required. `

	var usernameFlagName string
	if cmdPrefix == "" {
		usernameFlagName = "username"
	} else {
		usernameFlagName = fmt.Sprintf("%v.username", cmdPrefix)
	}

	var usernameFlagDefault string

	_ = cmd.PersistentFlags().String(usernameFlagName, usernameFlagDefault, usernameDescription)

	return nil
}

func registerModel55Zonename(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	zonenameDescription := `Required. `

	var zonenameFlagName string
	if cmdPrefix == "" {
		zonenameFlagName = "zonename"
	} else {
		zonenameFlagName = fmt.Sprintf("%v.zonename", cmdPrefix)
	}

	var zonenameFlagDefault string

	_ = cmd.PersistentFlags().String(zonenameFlagName, zonenameFlagDefault, zonenameDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModel55Flags(depth int, m *models.Model55, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, deviceNameAdded := retrieveModel55DeviceNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceNameAdded

	err, newPasswordAdded := retrieveModel55NewPasswordFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || newPasswordAdded

	err, timezoneAdded := retrieveModel55TimezoneFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || timezoneAdded

	err, usernameAdded := retrieveModel55UsernameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || usernameAdded

	err, zonenameAdded := retrieveModel55ZonenameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || zonenameAdded

	return nil, retAdded
}

func retrieveModel55DeviceNameFlags(depth int, m *models.Model55, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deviceNameFlagName := fmt.Sprintf("%v.deviceName", cmdPrefix)
	if cmd.Flags().Changed(deviceNameFlagName) {

		var deviceNameFlagName string
		if cmdPrefix == "" {
			deviceNameFlagName = "deviceName"
		} else {
			deviceNameFlagName = fmt.Sprintf("%v.deviceName", cmdPrefix)
		}

		deviceNameFlagValue, err := cmd.Flags().GetString(deviceNameFlagName)
		if err != nil {
			return err, false
		}
		m.DeviceName = &deviceNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel55NewPasswordFlags(depth int, m *models.Model55, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	newPasswordFlagName := fmt.Sprintf("%v.newPassword", cmdPrefix)
	if cmd.Flags().Changed(newPasswordFlagName) {

		var newPasswordFlagName string
		if cmdPrefix == "" {
			newPasswordFlagName = "newPassword"
		} else {
			newPasswordFlagName = fmt.Sprintf("%v.newPassword", cmdPrefix)
		}

		newPasswordFlagValue, err := cmd.Flags().GetString(newPasswordFlagName)
		if err != nil {
			return err, false
		}
		m.NewPassword = newPasswordFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel55TimezoneFlags(depth int, m *models.Model55, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	timezoneFlagName := fmt.Sprintf("%v.timezone", cmdPrefix)
	if cmd.Flags().Changed(timezoneFlagName) {

		var timezoneFlagName string
		if cmdPrefix == "" {
			timezoneFlagName = "timezone"
		} else {
			timezoneFlagName = fmt.Sprintf("%v.timezone", cmdPrefix)
		}

		timezoneFlagValue, err := cmd.Flags().GetString(timezoneFlagName)
		if err != nil {
			return err, false
		}
		m.Timezone = &timezoneFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel55UsernameFlags(depth int, m *models.Model55, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	usernameFlagName := fmt.Sprintf("%v.username", cmdPrefix)
	if cmd.Flags().Changed(usernameFlagName) {

		var usernameFlagName string
		if cmdPrefix == "" {
			usernameFlagName = "username"
		} else {
			usernameFlagName = fmt.Sprintf("%v.username", cmdPrefix)
		}

		usernameFlagValue, err := cmd.Flags().GetString(usernameFlagName)
		if err != nil {
			return err, false
		}
		m.Username = &usernameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel55ZonenameFlags(depth int, m *models.Model55, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	zonenameFlagName := fmt.Sprintf("%v.zonename", cmdPrefix)
	if cmd.Flags().Changed(zonenameFlagName) {

		var zonenameFlagName string
		if cmdPrefix == "" {
			zonenameFlagName = "zonename"
		} else {
			zonenameFlagName = fmt.Sprintf("%v.zonename", cmdPrefix)
		}

		zonenameFlagValue, err := cmd.Flags().GetString(zonenameFlagName)
		if err != nil {
			return err, false
		}
		m.Zonename = &zonenameFlagValue

		retAdded = true
	}

	return nil, retAdded
}
