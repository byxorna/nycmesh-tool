// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for VaultCredentials

// register flags to command
func registerModelVaultCredentialsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerVaultCredentialsIsCloudVaultEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVaultCredentialsIsVaultEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVaultCredentialsPassphrase(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVaultCredentialsRegeneratePgpKeys(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerVaultCredentialsIsCloudVaultEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	isCloudVaultEnabledDescription := `Required. `

	var isCloudVaultEnabledFlagName string
	if cmdPrefix == "" {
		isCloudVaultEnabledFlagName = "isCloudVaultEnabled"
	} else {
		isCloudVaultEnabledFlagName = fmt.Sprintf("%v.isCloudVaultEnabled", cmdPrefix)
	}

	var isCloudVaultEnabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(isCloudVaultEnabledFlagName, isCloudVaultEnabledFlagDefault, isCloudVaultEnabledDescription)

	return nil
}

func registerVaultCredentialsIsVaultEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	isVaultEnabledDescription := `Required. `

	var isVaultEnabledFlagName string
	if cmdPrefix == "" {
		isVaultEnabledFlagName = "isVaultEnabled"
	} else {
		isVaultEnabledFlagName = fmt.Sprintf("%v.isVaultEnabled", cmdPrefix)
	}

	var isVaultEnabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(isVaultEnabledFlagName, isVaultEnabledFlagDefault, isVaultEnabledDescription)

	return nil
}

func registerVaultCredentialsPassphrase(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	passphraseDescription := ``

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

func registerVaultCredentialsRegeneratePgpKeys(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	regeneratePgpKeysDescription := ``

	var regeneratePgpKeysFlagName string
	if cmdPrefix == "" {
		regeneratePgpKeysFlagName = "regeneratePgpKeys"
	} else {
		regeneratePgpKeysFlagName = fmt.Sprintf("%v.regeneratePgpKeys", cmdPrefix)
	}

	var regeneratePgpKeysFlagDefault bool

	_ = cmd.PersistentFlags().Bool(regeneratePgpKeysFlagName, regeneratePgpKeysFlagDefault, regeneratePgpKeysDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelVaultCredentialsFlags(depth int, m *models.VaultCredentials, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, isCloudVaultEnabledAdded := retrieveVaultCredentialsIsCloudVaultEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || isCloudVaultEnabledAdded

	err, isVaultEnabledAdded := retrieveVaultCredentialsIsVaultEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || isVaultEnabledAdded

	err, passphraseAdded := retrieveVaultCredentialsPassphraseFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || passphraseAdded

	err, regeneratePgpKeysAdded := retrieveVaultCredentialsRegeneratePgpKeysFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || regeneratePgpKeysAdded

	return nil, retAdded
}

func retrieveVaultCredentialsIsCloudVaultEnabledFlags(depth int, m *models.VaultCredentials, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	isCloudVaultEnabledFlagName := fmt.Sprintf("%v.isCloudVaultEnabled", cmdPrefix)
	if cmd.Flags().Changed(isCloudVaultEnabledFlagName) {

		var isCloudVaultEnabledFlagName string
		if cmdPrefix == "" {
			isCloudVaultEnabledFlagName = "isCloudVaultEnabled"
		} else {
			isCloudVaultEnabledFlagName = fmt.Sprintf("%v.isCloudVaultEnabled", cmdPrefix)
		}

		isCloudVaultEnabledFlagValue, err := cmd.Flags().GetBool(isCloudVaultEnabledFlagName)
		if err != nil {
			return err, false
		}
		m.IsCloudVaultEnabled = &isCloudVaultEnabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVaultCredentialsIsVaultEnabledFlags(depth int, m *models.VaultCredentials, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	isVaultEnabledFlagName := fmt.Sprintf("%v.isVaultEnabled", cmdPrefix)
	if cmd.Flags().Changed(isVaultEnabledFlagName) {

		var isVaultEnabledFlagName string
		if cmdPrefix == "" {
			isVaultEnabledFlagName = "isVaultEnabled"
		} else {
			isVaultEnabledFlagName = fmt.Sprintf("%v.isVaultEnabled", cmdPrefix)
		}

		isVaultEnabledFlagValue, err := cmd.Flags().GetBool(isVaultEnabledFlagName)
		if err != nil {
			return err, false
		}
		m.IsVaultEnabled = &isVaultEnabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVaultCredentialsPassphraseFlags(depth int, m *models.VaultCredentials, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Passphrase = passphraseFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVaultCredentialsRegeneratePgpKeysFlags(depth int, m *models.VaultCredentials, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	regeneratePgpKeysFlagName := fmt.Sprintf("%v.regeneratePgpKeys", cmdPrefix)
	if cmd.Flags().Changed(regeneratePgpKeysFlagName) {

		var regeneratePgpKeysFlagName string
		if cmdPrefix == "" {
			regeneratePgpKeysFlagName = "regeneratePgpKeys"
		} else {
			regeneratePgpKeysFlagName = fmt.Sprintf("%v.regeneratePgpKeys", cmdPrefix)
		}

		regeneratePgpKeysFlagValue, err := cmd.Flags().GetBool(regeneratePgpKeysFlagName)
		if err != nil {
			return err, false
		}
		m.RegeneratePgpKeys = regeneratePgpKeysFlagValue

		retAdded = true
	}

	return nil, retAdded
}
