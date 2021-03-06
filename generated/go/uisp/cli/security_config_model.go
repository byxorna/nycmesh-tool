// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for SecurityConfig

// register flags to command
func registerModelSecurityConfigFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSecurityConfigAccountingServerIP(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSecurityConfigAccountingServerPort(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSecurityConfigAccountingServerSecret(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSecurityConfigAuthServerIP(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSecurityConfigAuthServerPort(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSecurityConfigAuthServerSecret(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSecurityConfigBackupWPA(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSecurityConfigEapAnonymousIdentity(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSecurityConfigEapPassword(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSecurityConfigEapType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSecurityConfigEapTypeExt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSecurityConfigEapUsername(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSecurityConfigIsAAAEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSecurityConfigIsAccountingServerEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSecurityConfigIsWPASupplicantEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSecurityConfigPresharedKey(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSecurityConfigSecurity(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSecurityConfigWpaAuthentication(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSecurityConfigAccountingServerIP(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	accountingServerIpDescription := `Required. `

	var accountingServerIpFlagName string
	if cmdPrefix == "" {
		accountingServerIpFlagName = "accountingServerIP"
	} else {
		accountingServerIpFlagName = fmt.Sprintf("%v.accountingServerIP", cmdPrefix)
	}

	var accountingServerIpFlagDefault string

	_ = cmd.PersistentFlags().String(accountingServerIpFlagName, accountingServerIpFlagDefault, accountingServerIpDescription)

	return nil
}

func registerSecurityConfigAccountingServerPort(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	accountingServerPortDescription := `Required. `

	var accountingServerPortFlagName string
	if cmdPrefix == "" {
		accountingServerPortFlagName = "accountingServerPort"
	} else {
		accountingServerPortFlagName = fmt.Sprintf("%v.accountingServerPort", cmdPrefix)
	}

	var accountingServerPortFlagDefault float64

	_ = cmd.PersistentFlags().Float64(accountingServerPortFlagName, accountingServerPortFlagDefault, accountingServerPortDescription)

	return nil
}

func registerSecurityConfigAccountingServerSecret(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	accountingServerSecretDescription := `Required. `

	var accountingServerSecretFlagName string
	if cmdPrefix == "" {
		accountingServerSecretFlagName = "accountingServerSecret"
	} else {
		accountingServerSecretFlagName = fmt.Sprintf("%v.accountingServerSecret", cmdPrefix)
	}

	var accountingServerSecretFlagDefault string

	_ = cmd.PersistentFlags().String(accountingServerSecretFlagName, accountingServerSecretFlagDefault, accountingServerSecretDescription)

	return nil
}

func registerSecurityConfigAuthServerIP(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	authServerIpDescription := `Required. `

	var authServerIpFlagName string
	if cmdPrefix == "" {
		authServerIpFlagName = "authServerIP"
	} else {
		authServerIpFlagName = fmt.Sprintf("%v.authServerIP", cmdPrefix)
	}

	var authServerIpFlagDefault string

	_ = cmd.PersistentFlags().String(authServerIpFlagName, authServerIpFlagDefault, authServerIpDescription)

	return nil
}

func registerSecurityConfigAuthServerPort(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	authServerPortDescription := `Required. `

	var authServerPortFlagName string
	if cmdPrefix == "" {
		authServerPortFlagName = "authServerPort"
	} else {
		authServerPortFlagName = fmt.Sprintf("%v.authServerPort", cmdPrefix)
	}

	var authServerPortFlagDefault float64

	_ = cmd.PersistentFlags().Float64(authServerPortFlagName, authServerPortFlagDefault, authServerPortDescription)

	return nil
}

func registerSecurityConfigAuthServerSecret(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	authServerSecretDescription := `Required. `

	var authServerSecretFlagName string
	if cmdPrefix == "" {
		authServerSecretFlagName = "authServerSecret"
	} else {
		authServerSecretFlagName = fmt.Sprintf("%v.authServerSecret", cmdPrefix)
	}

	var authServerSecretFlagDefault string

	_ = cmd.PersistentFlags().String(authServerSecretFlagName, authServerSecretFlagDefault, authServerSecretDescription)

	return nil
}

func registerSecurityConfigBackupWPA(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var backupWPAFlagName string
	if cmdPrefix == "" {
		backupWPAFlagName = "backupWPA"
	} else {
		backupWPAFlagName = fmt.Sprintf("%v.backupWPA", cmdPrefix)
	}

	if err := registerModelBackupWPAFlags(depth+1, backupWPAFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerSecurityConfigEapAnonymousIdentity(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	eapAnonymousIdentityDescription := `Required. `

	var eapAnonymousIdentityFlagName string
	if cmdPrefix == "" {
		eapAnonymousIdentityFlagName = "eapAnonymousIdentity"
	} else {
		eapAnonymousIdentityFlagName = fmt.Sprintf("%v.eapAnonymousIdentity", cmdPrefix)
	}

	var eapAnonymousIdentityFlagDefault string

	_ = cmd.PersistentFlags().String(eapAnonymousIdentityFlagName, eapAnonymousIdentityFlagDefault, eapAnonymousIdentityDescription)

	return nil
}

func registerSecurityConfigEapPassword(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	eapPasswordDescription := `Required. `

	var eapPasswordFlagName string
	if cmdPrefix == "" {
		eapPasswordFlagName = "eapPassword"
	} else {
		eapPasswordFlagName = fmt.Sprintf("%v.eapPassword", cmdPrefix)
	}

	var eapPasswordFlagDefault string

	_ = cmd.PersistentFlags().String(eapPasswordFlagName, eapPasswordFlagDefault, eapPasswordDescription)

	return nil
}

func registerSecurityConfigEapType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	eapTypeDescription := `Enum: ["PEAP","TTLS"]. Required. `

	var eapTypeFlagName string
	if cmdPrefix == "" {
		eapTypeFlagName = "eapType"
	} else {
		eapTypeFlagName = fmt.Sprintf("%v.eapType", cmdPrefix)
	}

	var eapTypeFlagDefault string

	_ = cmd.PersistentFlags().String(eapTypeFlagName, eapTypeFlagDefault, eapTypeDescription)

	if err := cmd.RegisterFlagCompletionFunc(eapTypeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["PEAP","TTLS"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerSecurityConfigEapTypeExt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	eapTypeExtDescription := `Enum: ["auth=MSCHAPV2"]. Required. `

	var eapTypeExtFlagName string
	if cmdPrefix == "" {
		eapTypeExtFlagName = "eapTypeExt"
	} else {
		eapTypeExtFlagName = fmt.Sprintf("%v.eapTypeExt", cmdPrefix)
	}

	var eapTypeExtFlagDefault string

	_ = cmd.PersistentFlags().String(eapTypeExtFlagName, eapTypeExtFlagDefault, eapTypeExtDescription)

	if err := cmd.RegisterFlagCompletionFunc(eapTypeExtFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["auth=MSCHAPV2"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerSecurityConfigEapUsername(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	eapUsernameDescription := `Required. `

	var eapUsernameFlagName string
	if cmdPrefix == "" {
		eapUsernameFlagName = "eapUsername"
	} else {
		eapUsernameFlagName = fmt.Sprintf("%v.eapUsername", cmdPrefix)
	}

	var eapUsernameFlagDefault string

	_ = cmd.PersistentFlags().String(eapUsernameFlagName, eapUsernameFlagDefault, eapUsernameDescription)

	return nil
}

func registerSecurityConfigIsAAAEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	isAAAEnabledDescription := `Required. `

	var isAAAEnabledFlagName string
	if cmdPrefix == "" {
		isAAAEnabledFlagName = "isAAAEnabled"
	} else {
		isAAAEnabledFlagName = fmt.Sprintf("%v.isAAAEnabled", cmdPrefix)
	}

	var isAAAEnabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(isAAAEnabledFlagName, isAAAEnabledFlagDefault, isAAAEnabledDescription)

	return nil
}

func registerSecurityConfigIsAccountingServerEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	isAccountingServerEnabledDescription := `Required. `

	var isAccountingServerEnabledFlagName string
	if cmdPrefix == "" {
		isAccountingServerEnabledFlagName = "isAccountingServerEnabled"
	} else {
		isAccountingServerEnabledFlagName = fmt.Sprintf("%v.isAccountingServerEnabled", cmdPrefix)
	}

	var isAccountingServerEnabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(isAccountingServerEnabledFlagName, isAccountingServerEnabledFlagDefault, isAccountingServerEnabledDescription)

	return nil
}

func registerSecurityConfigIsWPASupplicantEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	isWPASupplicantEnabledDescription := `Required. `

	var isWPASupplicantEnabledFlagName string
	if cmdPrefix == "" {
		isWPASupplicantEnabledFlagName = "isWPASupplicantEnabled"
	} else {
		isWPASupplicantEnabledFlagName = fmt.Sprintf("%v.isWPASupplicantEnabled", cmdPrefix)
	}

	var isWPASupplicantEnabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(isWPASupplicantEnabledFlagName, isWPASupplicantEnabledFlagDefault, isWPASupplicantEnabledDescription)

	return nil
}

func registerSecurityConfigPresharedKey(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	presharedKeyDescription := `Required. `

	var presharedKeyFlagName string
	if cmdPrefix == "" {
		presharedKeyFlagName = "presharedKey"
	} else {
		presharedKeyFlagName = fmt.Sprintf("%v.presharedKey", cmdPrefix)
	}

	var presharedKeyFlagDefault string

	_ = cmd.PersistentFlags().String(presharedKeyFlagName, presharedKeyFlagDefault, presharedKeyDescription)

	return nil
}

func registerSecurityConfigSecurity(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	securityDescription := `Enum: ["wep","wpa","wpa2","wpaTKIP","wpa2TKIP","wpaAES","wpa2AES","aes256","none"]. Required. `

	var securityFlagName string
	if cmdPrefix == "" {
		securityFlagName = "security"
	} else {
		securityFlagName = fmt.Sprintf("%v.security", cmdPrefix)
	}

	var securityFlagDefault string

	_ = cmd.PersistentFlags().String(securityFlagName, securityFlagDefault, securityDescription)

	if err := cmd.RegisterFlagCompletionFunc(securityFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["wep","wpa","wpa2","wpaTKIP","wpa2TKIP","wpaAES","wpa2AES","aes256","none"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerSecurityConfigWpaAuthentication(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	wpaAuthenticationDescription := `Enum: ["psk","psk2","ent","none"]. Required. `

	var wpaAuthenticationFlagName string
	if cmdPrefix == "" {
		wpaAuthenticationFlagName = "wpaAuthentication"
	} else {
		wpaAuthenticationFlagName = fmt.Sprintf("%v.wpaAuthentication", cmdPrefix)
	}

	var wpaAuthenticationFlagDefault string

	_ = cmd.PersistentFlags().String(wpaAuthenticationFlagName, wpaAuthenticationFlagDefault, wpaAuthenticationDescription)

	if err := cmd.RegisterFlagCompletionFunc(wpaAuthenticationFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["psk","psk2","ent","none"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSecurityConfigFlags(depth int, m *models.SecurityConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, accountingServerIpAdded := retrieveSecurityConfigAccountingServerIPFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || accountingServerIpAdded

	err, accountingServerPortAdded := retrieveSecurityConfigAccountingServerPortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || accountingServerPortAdded

	err, accountingServerSecretAdded := retrieveSecurityConfigAccountingServerSecretFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || accountingServerSecretAdded

	err, authServerIpAdded := retrieveSecurityConfigAuthServerIPFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || authServerIpAdded

	err, authServerPortAdded := retrieveSecurityConfigAuthServerPortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || authServerPortAdded

	err, authServerSecretAdded := retrieveSecurityConfigAuthServerSecretFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || authServerSecretAdded

	err, backupWPAAdded := retrieveSecurityConfigBackupWPAFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || backupWPAAdded

	err, eapAnonymousIdentityAdded := retrieveSecurityConfigEapAnonymousIdentityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || eapAnonymousIdentityAdded

	err, eapPasswordAdded := retrieveSecurityConfigEapPasswordFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || eapPasswordAdded

	err, eapTypeAdded := retrieveSecurityConfigEapTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || eapTypeAdded

	err, eapTypeExtAdded := retrieveSecurityConfigEapTypeExtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || eapTypeExtAdded

	err, eapUsernameAdded := retrieveSecurityConfigEapUsernameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || eapUsernameAdded

	err, isAAAEnabledAdded := retrieveSecurityConfigIsAAAEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || isAAAEnabledAdded

	err, isAccountingServerEnabledAdded := retrieveSecurityConfigIsAccountingServerEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || isAccountingServerEnabledAdded

	err, isWPASupplicantEnabledAdded := retrieveSecurityConfigIsWPASupplicantEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || isWPASupplicantEnabledAdded

	err, presharedKeyAdded := retrieveSecurityConfigPresharedKeyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || presharedKeyAdded

	err, securityAdded := retrieveSecurityConfigSecurityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || securityAdded

	err, wpaAuthenticationAdded := retrieveSecurityConfigWpaAuthenticationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || wpaAuthenticationAdded

	return nil, retAdded
}

func retrieveSecurityConfigAccountingServerIPFlags(depth int, m *models.SecurityConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	accountingServerIpFlagName := fmt.Sprintf("%v.accountingServerIP", cmdPrefix)
	if cmd.Flags().Changed(accountingServerIpFlagName) {

		var accountingServerIpFlagName string
		if cmdPrefix == "" {
			accountingServerIpFlagName = "accountingServerIP"
		} else {
			accountingServerIpFlagName = fmt.Sprintf("%v.accountingServerIP", cmdPrefix)
		}

		accountingServerIpFlagValue, err := cmd.Flags().GetString(accountingServerIpFlagName)
		if err != nil {
			return err, false
		}
		m.AccountingServerIP = &accountingServerIpFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSecurityConfigAccountingServerPortFlags(depth int, m *models.SecurityConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	accountingServerPortFlagName := fmt.Sprintf("%v.accountingServerPort", cmdPrefix)
	if cmd.Flags().Changed(accountingServerPortFlagName) {

		var accountingServerPortFlagName string
		if cmdPrefix == "" {
			accountingServerPortFlagName = "accountingServerPort"
		} else {
			accountingServerPortFlagName = fmt.Sprintf("%v.accountingServerPort", cmdPrefix)
		}

		accountingServerPortFlagValue, err := cmd.Flags().GetFloat64(accountingServerPortFlagName)
		if err != nil {
			return err, false
		}
		m.AccountingServerPort = &accountingServerPortFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSecurityConfigAccountingServerSecretFlags(depth int, m *models.SecurityConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	accountingServerSecretFlagName := fmt.Sprintf("%v.accountingServerSecret", cmdPrefix)
	if cmd.Flags().Changed(accountingServerSecretFlagName) {

		var accountingServerSecretFlagName string
		if cmdPrefix == "" {
			accountingServerSecretFlagName = "accountingServerSecret"
		} else {
			accountingServerSecretFlagName = fmt.Sprintf("%v.accountingServerSecret", cmdPrefix)
		}

		accountingServerSecretFlagValue, err := cmd.Flags().GetString(accountingServerSecretFlagName)
		if err != nil {
			return err, false
		}
		m.AccountingServerSecret = &accountingServerSecretFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSecurityConfigAuthServerIPFlags(depth int, m *models.SecurityConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	authServerIpFlagName := fmt.Sprintf("%v.authServerIP", cmdPrefix)
	if cmd.Flags().Changed(authServerIpFlagName) {

		var authServerIpFlagName string
		if cmdPrefix == "" {
			authServerIpFlagName = "authServerIP"
		} else {
			authServerIpFlagName = fmt.Sprintf("%v.authServerIP", cmdPrefix)
		}

		authServerIpFlagValue, err := cmd.Flags().GetString(authServerIpFlagName)
		if err != nil {
			return err, false
		}
		m.AuthServerIP = &authServerIpFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSecurityConfigAuthServerPortFlags(depth int, m *models.SecurityConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	authServerPortFlagName := fmt.Sprintf("%v.authServerPort", cmdPrefix)
	if cmd.Flags().Changed(authServerPortFlagName) {

		var authServerPortFlagName string
		if cmdPrefix == "" {
			authServerPortFlagName = "authServerPort"
		} else {
			authServerPortFlagName = fmt.Sprintf("%v.authServerPort", cmdPrefix)
		}

		authServerPortFlagValue, err := cmd.Flags().GetFloat64(authServerPortFlagName)
		if err != nil {
			return err, false
		}
		m.AuthServerPort = &authServerPortFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSecurityConfigAuthServerSecretFlags(depth int, m *models.SecurityConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	authServerSecretFlagName := fmt.Sprintf("%v.authServerSecret", cmdPrefix)
	if cmd.Flags().Changed(authServerSecretFlagName) {

		var authServerSecretFlagName string
		if cmdPrefix == "" {
			authServerSecretFlagName = "authServerSecret"
		} else {
			authServerSecretFlagName = fmt.Sprintf("%v.authServerSecret", cmdPrefix)
		}

		authServerSecretFlagValue, err := cmd.Flags().GetString(authServerSecretFlagName)
		if err != nil {
			return err, false
		}
		m.AuthServerSecret = &authServerSecretFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSecurityConfigBackupWPAFlags(depth int, m *models.SecurityConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	backupWPAFlagName := fmt.Sprintf("%v.backupWPA", cmdPrefix)
	if cmd.Flags().Changed(backupWPAFlagName) {
		// info: complex object backupWPA BackupWPA is retrieved outside this Changed() block
	}
	backupWPAFlagValue := m.BackupWPA
	if swag.IsZero(backupWPAFlagValue) {
		backupWPAFlagValue = &models.BackupWPA{}
	}

	err, backupWPAAdded := retrieveModelBackupWPAFlags(depth+1, backupWPAFlagValue, backupWPAFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || backupWPAAdded
	if backupWPAAdded {
		m.BackupWPA = backupWPAFlagValue
	}

	return nil, retAdded
}

func retrieveSecurityConfigEapAnonymousIdentityFlags(depth int, m *models.SecurityConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	eapAnonymousIdentityFlagName := fmt.Sprintf("%v.eapAnonymousIdentity", cmdPrefix)
	if cmd.Flags().Changed(eapAnonymousIdentityFlagName) {

		var eapAnonymousIdentityFlagName string
		if cmdPrefix == "" {
			eapAnonymousIdentityFlagName = "eapAnonymousIdentity"
		} else {
			eapAnonymousIdentityFlagName = fmt.Sprintf("%v.eapAnonymousIdentity", cmdPrefix)
		}

		eapAnonymousIdentityFlagValue, err := cmd.Flags().GetString(eapAnonymousIdentityFlagName)
		if err != nil {
			return err, false
		}
		m.EapAnonymousIdentity = &eapAnonymousIdentityFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSecurityConfigEapPasswordFlags(depth int, m *models.SecurityConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	eapPasswordFlagName := fmt.Sprintf("%v.eapPassword", cmdPrefix)
	if cmd.Flags().Changed(eapPasswordFlagName) {

		var eapPasswordFlagName string
		if cmdPrefix == "" {
			eapPasswordFlagName = "eapPassword"
		} else {
			eapPasswordFlagName = fmt.Sprintf("%v.eapPassword", cmdPrefix)
		}

		eapPasswordFlagValue, err := cmd.Flags().GetString(eapPasswordFlagName)
		if err != nil {
			return err, false
		}
		m.EapPassword = &eapPasswordFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSecurityConfigEapTypeFlags(depth int, m *models.SecurityConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	eapTypeFlagName := fmt.Sprintf("%v.eapType", cmdPrefix)
	if cmd.Flags().Changed(eapTypeFlagName) {

		var eapTypeFlagName string
		if cmdPrefix == "" {
			eapTypeFlagName = "eapType"
		} else {
			eapTypeFlagName = fmt.Sprintf("%v.eapType", cmdPrefix)
		}

		eapTypeFlagValue, err := cmd.Flags().GetString(eapTypeFlagName)
		if err != nil {
			return err, false
		}
		m.EapType = &eapTypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSecurityConfigEapTypeExtFlags(depth int, m *models.SecurityConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	eapTypeExtFlagName := fmt.Sprintf("%v.eapTypeExt", cmdPrefix)
	if cmd.Flags().Changed(eapTypeExtFlagName) {

		var eapTypeExtFlagName string
		if cmdPrefix == "" {
			eapTypeExtFlagName = "eapTypeExt"
		} else {
			eapTypeExtFlagName = fmt.Sprintf("%v.eapTypeExt", cmdPrefix)
		}

		eapTypeExtFlagValue, err := cmd.Flags().GetString(eapTypeExtFlagName)
		if err != nil {
			return err, false
		}
		m.EapTypeExt = &eapTypeExtFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSecurityConfigEapUsernameFlags(depth int, m *models.SecurityConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	eapUsernameFlagName := fmt.Sprintf("%v.eapUsername", cmdPrefix)
	if cmd.Flags().Changed(eapUsernameFlagName) {

		var eapUsernameFlagName string
		if cmdPrefix == "" {
			eapUsernameFlagName = "eapUsername"
		} else {
			eapUsernameFlagName = fmt.Sprintf("%v.eapUsername", cmdPrefix)
		}

		eapUsernameFlagValue, err := cmd.Flags().GetString(eapUsernameFlagName)
		if err != nil {
			return err, false
		}
		m.EapUsername = &eapUsernameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSecurityConfigIsAAAEnabledFlags(depth int, m *models.SecurityConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	isAAAEnabledFlagName := fmt.Sprintf("%v.isAAAEnabled", cmdPrefix)
	if cmd.Flags().Changed(isAAAEnabledFlagName) {

		var isAAAEnabledFlagName string
		if cmdPrefix == "" {
			isAAAEnabledFlagName = "isAAAEnabled"
		} else {
			isAAAEnabledFlagName = fmt.Sprintf("%v.isAAAEnabled", cmdPrefix)
		}

		isAAAEnabledFlagValue, err := cmd.Flags().GetBool(isAAAEnabledFlagName)
		if err != nil {
			return err, false
		}
		m.IsAAAEnabled = &isAAAEnabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSecurityConfigIsAccountingServerEnabledFlags(depth int, m *models.SecurityConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	isAccountingServerEnabledFlagName := fmt.Sprintf("%v.isAccountingServerEnabled", cmdPrefix)
	if cmd.Flags().Changed(isAccountingServerEnabledFlagName) {

		var isAccountingServerEnabledFlagName string
		if cmdPrefix == "" {
			isAccountingServerEnabledFlagName = "isAccountingServerEnabled"
		} else {
			isAccountingServerEnabledFlagName = fmt.Sprintf("%v.isAccountingServerEnabled", cmdPrefix)
		}

		isAccountingServerEnabledFlagValue, err := cmd.Flags().GetBool(isAccountingServerEnabledFlagName)
		if err != nil {
			return err, false
		}
		m.IsAccountingServerEnabled = &isAccountingServerEnabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSecurityConfigIsWPASupplicantEnabledFlags(depth int, m *models.SecurityConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	isWPASupplicantEnabledFlagName := fmt.Sprintf("%v.isWPASupplicantEnabled", cmdPrefix)
	if cmd.Flags().Changed(isWPASupplicantEnabledFlagName) {

		var isWPASupplicantEnabledFlagName string
		if cmdPrefix == "" {
			isWPASupplicantEnabledFlagName = "isWPASupplicantEnabled"
		} else {
			isWPASupplicantEnabledFlagName = fmt.Sprintf("%v.isWPASupplicantEnabled", cmdPrefix)
		}

		isWPASupplicantEnabledFlagValue, err := cmd.Flags().GetBool(isWPASupplicantEnabledFlagName)
		if err != nil {
			return err, false
		}
		m.IsWPASupplicantEnabled = &isWPASupplicantEnabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSecurityConfigPresharedKeyFlags(depth int, m *models.SecurityConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	presharedKeyFlagName := fmt.Sprintf("%v.presharedKey", cmdPrefix)
	if cmd.Flags().Changed(presharedKeyFlagName) {

		var presharedKeyFlagName string
		if cmdPrefix == "" {
			presharedKeyFlagName = "presharedKey"
		} else {
			presharedKeyFlagName = fmt.Sprintf("%v.presharedKey", cmdPrefix)
		}

		presharedKeyFlagValue, err := cmd.Flags().GetString(presharedKeyFlagName)
		if err != nil {
			return err, false
		}
		m.PresharedKey = &presharedKeyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSecurityConfigSecurityFlags(depth int, m *models.SecurityConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	securityFlagName := fmt.Sprintf("%v.security", cmdPrefix)
	if cmd.Flags().Changed(securityFlagName) {

		var securityFlagName string
		if cmdPrefix == "" {
			securityFlagName = "security"
		} else {
			securityFlagName = fmt.Sprintf("%v.security", cmdPrefix)
		}

		securityFlagValue, err := cmd.Flags().GetString(securityFlagName)
		if err != nil {
			return err, false
		}
		m.Security = &securityFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSecurityConfigWpaAuthenticationFlags(depth int, m *models.SecurityConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	wpaAuthenticationFlagName := fmt.Sprintf("%v.wpaAuthentication", cmdPrefix)
	if cmd.Flags().Changed(wpaAuthenticationFlagName) {

		var wpaAuthenticationFlagName string
		if cmdPrefix == "" {
			wpaAuthenticationFlagName = "wpaAuthentication"
		} else {
			wpaAuthenticationFlagName = fmt.Sprintf("%v.wpaAuthentication", cmdPrefix)
		}

		wpaAuthenticationFlagValue, err := cmd.Flags().GetString(wpaAuthenticationFlagName)
		if err != nil {
			return err, false
		}
		m.WpaAuthentication = &wpaAuthenticationFlagValue

		retAdded = true
	}

	return nil, retAdded
}
