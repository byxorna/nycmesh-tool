// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for Admin1

// register flags to command
func registerModelAdmin1Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerAdmin1Login(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerAdmin1Login(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var loginFlagName string
	if cmdPrefix == "" {
		loginFlagName = "login"
	} else {
		loginFlagName = fmt.Sprintf("%v.login", cmdPrefix)
	}

	if err := registerModelLogin3Flags(depth+1, loginFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelAdmin1Flags(depth int, m *models.Admin1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, loginAdded := retrieveAdmin1LoginFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || loginAdded

	return nil, retAdded
}

func retrieveAdmin1LoginFlags(depth int, m *models.Admin1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	loginFlagName := fmt.Sprintf("%v.login", cmdPrefix)
	if cmd.Flags().Changed(loginFlagName) {
		// info: complex object login Login3 is retrieved outside this Changed() block
	}
	loginFlagValue := m.Login
	if swag.IsZero(loginFlagValue) {
		loginFlagValue = &models.Login3{}
	}

	err, loginAdded := retrieveModelLogin3Flags(depth+1, loginFlagValue, loginFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || loginAdded
	if loginAdded {
		m.Login = loginFlagValue
	}

	return nil, retAdded
}
