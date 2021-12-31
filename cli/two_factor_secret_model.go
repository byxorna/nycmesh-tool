// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for TwoFactorSecret

// register flags to command
func registerModelTwoFactorSecretFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTwoFactorSecretBase32(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTwoFactorSecretOtpauthURL(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTwoFactorSecretBase32(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	base32Description := ``

	var base32FlagName string
	if cmdPrefix == "" {
		base32FlagName = "base32"
	} else {
		base32FlagName = fmt.Sprintf("%v.base32", cmdPrefix)
	}

	var base32FlagDefault string

	_ = cmd.PersistentFlags().String(base32FlagName, base32FlagDefault, base32Description)

	return nil
}

func registerTwoFactorSecretOtpauthURL(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	otpauthUrlDescription := ``

	var otpauthUrlFlagName string
	if cmdPrefix == "" {
		otpauthUrlFlagName = "otpauth_url"
	} else {
		otpauthUrlFlagName = fmt.Sprintf("%v.otpauth_url", cmdPrefix)
	}

	var otpauthUrlFlagDefault string

	_ = cmd.PersistentFlags().String(otpauthUrlFlagName, otpauthUrlFlagDefault, otpauthUrlDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTwoFactorSecretFlags(depth int, m *models.TwoFactorSecret, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, base32Added := retrieveTwoFactorSecretBase32Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || base32Added

	err, otpauthUrlAdded := retrieveTwoFactorSecretOtpauthURLFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || otpauthUrlAdded

	return nil, retAdded
}

func retrieveTwoFactorSecretBase32Flags(depth int, m *models.TwoFactorSecret, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	base32FlagName := fmt.Sprintf("%v.base32", cmdPrefix)
	if cmd.Flags().Changed(base32FlagName) {

		var base32FlagName string
		if cmdPrefix == "" {
			base32FlagName = "base32"
		} else {
			base32FlagName = fmt.Sprintf("%v.base32", cmdPrefix)
		}

		base32FlagValue, err := cmd.Flags().GetString(base32FlagName)
		if err != nil {
			return err, false
		}
		m.Base32 = base32FlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTwoFactorSecretOtpauthURLFlags(depth int, m *models.TwoFactorSecret, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	otpauthUrlFlagName := fmt.Sprintf("%v.otpauth_url", cmdPrefix)
	if cmd.Flags().Changed(otpauthUrlFlagName) {

		var otpauthUrlFlagName string
		if cmdPrefix == "" {
			otpauthUrlFlagName = "otpauth_url"
		} else {
			otpauthUrlFlagName = fmt.Sprintf("%v.otpauth_url", cmdPrefix)
		}

		otpauthUrlFlagValue, err := cmd.Flags().GetString(otpauthUrlFlagName)
		if err != nil {
			return err, false
		}
		m.OtpauthURL = otpauthUrlFlagValue

		retAdded = true
	}

	return nil, retAdded
}