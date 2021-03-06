// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for VerifyPasswordAction

// register flags to command
func registerModelVerifyPasswordActionFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerVerifyPasswordActionPassword(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVerifyPasswordActionTotpAuthEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerVerifyPasswordActionPassword(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerVerifyPasswordActionTotpAuthEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totpAuthEnabledDescription := `Required. `

	var totpAuthEnabledFlagName string
	if cmdPrefix == "" {
		totpAuthEnabledFlagName = "totpAuthEnabled"
	} else {
		totpAuthEnabledFlagName = fmt.Sprintf("%v.totpAuthEnabled", cmdPrefix)
	}

	var totpAuthEnabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(totpAuthEnabledFlagName, totpAuthEnabledFlagDefault, totpAuthEnabledDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelVerifyPasswordActionFlags(depth int, m *models.VerifyPasswordAction, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, passwordAdded := retrieveVerifyPasswordActionPasswordFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || passwordAdded

	err, totpAuthEnabledAdded := retrieveVerifyPasswordActionTotpAuthEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totpAuthEnabledAdded

	return nil, retAdded
}

func retrieveVerifyPasswordActionPasswordFlags(depth int, m *models.VerifyPasswordAction, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveVerifyPasswordActionTotpAuthEnabledFlags(depth int, m *models.VerifyPasswordAction, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totpAuthEnabledFlagName := fmt.Sprintf("%v.totpAuthEnabled", cmdPrefix)
	if cmd.Flags().Changed(totpAuthEnabledFlagName) {

		var totpAuthEnabledFlagName string
		if cmdPrefix == "" {
			totpAuthEnabledFlagName = "totpAuthEnabled"
		} else {
			totpAuthEnabledFlagName = fmt.Sprintf("%v.totpAuthEnabled", cmdPrefix)
		}

		totpAuthEnabledFlagValue, err := cmd.Flags().GetBool(totpAuthEnabledFlagName)
		if err != nil {
			return err, false
		}
		m.TotpAuthEnabled = &totpAuthEnabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}
