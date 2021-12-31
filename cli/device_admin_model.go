// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for DeviceAdmin

// register flags to command
func registerModelDeviceAdminFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDeviceAdminLogin(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceAdminLogin(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var loginFlagName string
	if cmdPrefix == "" {
		loginFlagName = "login"
	} else {
		loginFlagName = fmt.Sprintf("%v.login", cmdPrefix)
	}

	if err := registerModelLogin1Flags(depth+1, loginFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDeviceAdminFlags(depth int, m *models.DeviceAdmin, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, loginAdded := retrieveDeviceAdminLoginFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || loginAdded

	return nil, retAdded
}

func retrieveDeviceAdminLoginFlags(depth int, m *models.DeviceAdmin, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	loginFlagName := fmt.Sprintf("%v.login", cmdPrefix)
	if cmd.Flags().Changed(loginFlagName) {
		// info: complex object login Login1 is retrieved outside this Changed() block
	}
	loginFlagValue := m.Login
	if swag.IsZero(loginFlagValue) {
		loginFlagValue = &models.Login1{}
	}

	err, loginAdded := retrieveModelLogin1Flags(depth+1, loginFlagValue, loginFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || loginAdded
	if loginAdded {
		m.Login = loginFlagValue
	}

	return nil, retAdded
}
