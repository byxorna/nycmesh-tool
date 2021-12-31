// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"
  "github.com/byxorna/nycmesh-tool/models"

	"github.com/spf13/cobra"
)

// Schema cli for Admin3

// register flags to command
func registerModelAdmin3Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerAdmin3Login(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerAdmin3Login(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var loginFlagName string
	if cmdPrefix == "" {
		loginFlagName = "login"
	} else {
		loginFlagName = fmt.Sprintf("%v.login", cmdPrefix)
	}

	if err := registerModelLogin5Flags(depth+1, loginFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelAdmin3Flags(depth int, m *models.Admin3, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, loginAdded := retrieveAdmin3LoginFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || loginAdded

	return nil, retAdded
}

func retrieveAdmin3LoginFlags(depth int, m *models.Admin3, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	loginFlagName := fmt.Sprintf("%v.login", cmdPrefix)
	if cmd.Flags().Changed(loginFlagName) {
		// info: complex object login Login5 is retrieved outside this Changed() block
	}
	loginFlagValue := m.Login
	if swag.IsZero(loginFlagValue) {
		loginFlagValue = &models.Login5{}
	}

	err, loginAdded := retrieveModelLogin5Flags(depth+1, loginFlagValue, loginFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || loginAdded
	if loginAdded {
		m.Login = loginFlagValue
	}

	return nil, retAdded
}
