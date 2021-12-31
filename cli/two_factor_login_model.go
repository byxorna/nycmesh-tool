// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for TwoFactorLogin

// register flags to command
func registerModelTwoFactorLoginFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTwoFactorLoginPassword(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTwoFactorLoginSessionTimeout(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTwoFactorLoginToken(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTwoFactorLoginVerificationCode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTwoFactorLoginPassword(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	passwordDescription := `Ignored, kept for backward compatibility.`

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

func registerTwoFactorLoginSessionTimeout(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sessionTimeoutDescription := `Ignored, kept for backward compatibility.`

	var sessionTimeoutFlagName string
	if cmdPrefix == "" {
		sessionTimeoutFlagName = "sessionTimeout"
	} else {
		sessionTimeoutFlagName = fmt.Sprintf("%v.sessionTimeout", cmdPrefix)
	}

	var sessionTimeoutFlagDefault float64

	_ = cmd.PersistentFlags().Float64(sessionTimeoutFlagName, sessionTimeoutFlagDefault, sessionTimeoutDescription)

	return nil
}

func registerTwoFactorLoginToken(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	tokenDescription := `Required. Token from successful login.`

	var tokenFlagName string
	if cmdPrefix == "" {
		tokenFlagName = "token"
	} else {
		tokenFlagName = fmt.Sprintf("%v.token", cmdPrefix)
	}

	var tokenFlagDefault string

	_ = cmd.PersistentFlags().String(tokenFlagName, tokenFlagDefault, tokenDescription)

	return nil
}

func registerTwoFactorLoginVerificationCode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	verificationCodeDescription := `Required. 6 digit code from Authenticator.`

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
func retrieveModelTwoFactorLoginFlags(depth int, m *models.TwoFactorLogin, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, passwordAdded := retrieveTwoFactorLoginPasswordFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || passwordAdded

	err, sessionTimeoutAdded := retrieveTwoFactorLoginSessionTimeoutFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sessionTimeoutAdded

	err, tokenAdded := retrieveTwoFactorLoginTokenFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tokenAdded

	err, verificationCodeAdded := retrieveTwoFactorLoginVerificationCodeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || verificationCodeAdded

	return nil, retAdded
}

func retrieveTwoFactorLoginPasswordFlags(depth int, m *models.TwoFactorLogin, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Password = passwordFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTwoFactorLoginSessionTimeoutFlags(depth int, m *models.TwoFactorLogin, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sessionTimeoutFlagName := fmt.Sprintf("%v.sessionTimeout", cmdPrefix)
	if cmd.Flags().Changed(sessionTimeoutFlagName) {

		var sessionTimeoutFlagName string
		if cmdPrefix == "" {
			sessionTimeoutFlagName = "sessionTimeout"
		} else {
			sessionTimeoutFlagName = fmt.Sprintf("%v.sessionTimeout", cmdPrefix)
		}

		sessionTimeoutFlagValue, err := cmd.Flags().GetFloat64(sessionTimeoutFlagName)
		if err != nil {
			return err, false
		}
		m.SessionTimeout = sessionTimeoutFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTwoFactorLoginTokenFlags(depth int, m *models.TwoFactorLogin, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	tokenFlagName := fmt.Sprintf("%v.token", cmdPrefix)
	if cmd.Flags().Changed(tokenFlagName) {

		var tokenFlagName string
		if cmdPrefix == "" {
			tokenFlagName = "token"
		} else {
			tokenFlagName = fmt.Sprintf("%v.token", cmdPrefix)
		}

		tokenFlagValue, err := cmd.Flags().GetString(tokenFlagName)
		if err != nil {
			return err, false
		}
		m.Token = &tokenFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTwoFactorLoginVerificationCodeFlags(depth int, m *models.TwoFactorLogin, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.VerificationCode = &verificationCodeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
