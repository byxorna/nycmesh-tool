// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for UserCredentials

// register flags to command
func registerModelUserCredentialsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUserCredentialsPassword(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserCredentialsUsername(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUserCredentialsPassword(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	passwordDescription := `Required. `

	var passwordFlagName string
	if cmdPrefix == "" {
		passwordFlagName = "password"
	} else {
		passwordFlagName = fmt.Sprintf("%v.password", cmdPrefix)
	}

	var passwordFlagDefault string

	_ = cmd.PersistentFlags().String(passwordFlagName, passwordFlagDefault, passwordDescription)

	return nil
}

func registerUserCredentialsUsername(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUserCredentialsFlags(depth int, m *models.UserCredentials, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, passwordAdded := retrieveUserCredentialsPasswordFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || passwordAdded

	err, usernameAdded := retrieveUserCredentialsUsernameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || usernameAdded

	return nil, retAdded
}

func retrieveUserCredentialsPasswordFlags(depth int, m *models.UserCredentials, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	passwordFlagName := fmt.Sprintf("%v.password", cmdPrefix)
	if cmd.Flags().Changed(passwordFlagName) {

		var passwordFlagName string
		if cmdPrefix == "" {
			passwordFlagName = "password"
		} else {
			passwordFlagName = fmt.Sprintf("%v.password", cmdPrefix)
		}

		passwordFlagValue, err := cmd.Flags().GetString(passwordFlagName)
		if err != nil {
			return err, false
		}
		m.Password = &passwordFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUserCredentialsUsernameFlags(depth int, m *models.UserCredentials, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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