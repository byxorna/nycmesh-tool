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

// Schema cli for EnableSso

// register flags to command
func registerModelEnableSsoFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerEnableSsoPassword(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEnableSsoSsoUser(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEnableSsoVerificationCode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerEnableSsoPassword(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	passwordDescription := `Required. UISP user password.`

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

func registerEnableSsoSsoUser(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var ssoUserFlagName string
	if cmdPrefix == "" {
		ssoUserFlagName = "ssoUser"
	} else {
		ssoUserFlagName = fmt.Sprintf("%v.ssoUser", cmdPrefix)
	}

	if err := registerModelSsoUserSchema1Flags(depth+1, ssoUserFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerEnableSsoVerificationCode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	verificationCodeDescription := `UISP user 2FA verification code. Not required when 2FA is not enabled for the account.`

	var verificationCodeFlagName string
	if cmdPrefix == "" {
		verificationCodeFlagName = "verificationCode"
	} else {
		verificationCodeFlagName = fmt.Sprintf("%v.verificationCode", cmdPrefix)
	}

	var verificationCodeFlagDefault string

	_ = cmd.PersistentFlags().String(verificationCodeFlagName, verificationCodeFlagDefault, verificationCodeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelEnableSsoFlags(depth int, m *models.EnableSso, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, passwordAdded := retrieveEnableSsoPasswordFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || passwordAdded

	err, ssoUserAdded := retrieveEnableSsoSsoUserFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ssoUserAdded

	err, verificationCodeAdded := retrieveEnableSsoVerificationCodeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || verificationCodeAdded

	return nil, retAdded
}

func retrieveEnableSsoPasswordFlags(depth int, m *models.EnableSso, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveEnableSsoSsoUserFlags(depth int, m *models.EnableSso, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ssoUserFlagName := fmt.Sprintf("%v.ssoUser", cmdPrefix)
	if cmd.Flags().Changed(ssoUserFlagName) {
		// info: complex object ssoUser SsoUserSchema1 is retrieved outside this Changed() block
	}
	ssoUserFlagValue := m.SsoUser
	if swag.IsZero(ssoUserFlagValue) {
		ssoUserFlagValue = &models.SsoUserSchema1{}
	}

	err, ssoUserAdded := retrieveModelSsoUserSchema1Flags(depth+1, ssoUserFlagValue, ssoUserFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ssoUserAdded
	if ssoUserAdded {
		m.SsoUser = ssoUserFlagValue
	}

	return nil, retAdded
}

func retrieveEnableSsoVerificationCodeFlags(depth int, m *models.EnableSso, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	verificationCodeFlagName := fmt.Sprintf("%v.verificationCode", cmdPrefix)
	if cmd.Flags().Changed(verificationCodeFlagName) {

		var verificationCodeFlagName string
		if cmdPrefix == "" {
			verificationCodeFlagName = "verificationCode"
		} else {
			verificationCodeFlagName = fmt.Sprintf("%v.verificationCode", cmdPrefix)
		}

		verificationCodeFlagValue, err := cmd.Flags().GetString(verificationCodeFlagName)
		if err != nil {
			return err, false
		}
		m.VerificationCode = verificationCodeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
