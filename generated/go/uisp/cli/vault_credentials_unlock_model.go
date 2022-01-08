// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for VaultCredentialsUnlock

// register flags to command
func registerModelVaultCredentialsUnlockFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerVaultCredentialsUnlockPassphrase(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerVaultCredentialsUnlockPassphrase(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	passphraseDescription := `Required. `

	var passphraseFlagName string
	if cmdPrefix == "" {
		passphraseFlagName = "passphrase"
	} else {
		passphraseFlagName = fmt.Sprintf("%v.passphrase", cmdPrefix)
	}

	var passphraseFlagDefault string

	_ = cmd.PersistentFlags().String(passphraseFlagName, passphraseFlagDefault, passphraseDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelVaultCredentialsUnlockFlags(depth int, m *models.VaultCredentialsUnlock, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, passphraseAdded := retrieveVaultCredentialsUnlockPassphraseFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || passphraseAdded

	return nil, retAdded
}

func retrieveVaultCredentialsUnlockPassphraseFlags(depth int, m *models.VaultCredentialsUnlock, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	passphraseFlagName := fmt.Sprintf("%v.passphrase", cmdPrefix)
	if cmd.Flags().Changed(passphraseFlagName) {

		var passphraseFlagName string
		if cmdPrefix == "" {
			passphraseFlagName = "passphrase"
		} else {
			passphraseFlagName = fmt.Sprintf("%v.passphrase", cmdPrefix)
		}

		passphraseFlagValue, err := cmd.Flags().GetString(passphraseFlagName)
		if err != nil {
			return err, false
		}
		m.Passphrase = &passphraseFlagValue

		retAdded = true
	}

	return nil, retAdded
}
