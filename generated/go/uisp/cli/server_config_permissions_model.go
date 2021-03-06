// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for ServerConfigPermissions

// register flags to command
func registerModelServerConfigPermissionsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerServerConfigPermissionsCanConfigureCertificate(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerConfigPermissionsCanConfigureDeviceProfiles(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerConfigPermissionsCanConfigureFirstDeviceInSetup(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerConfigPermissionsCanConfigureHostname(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerConfigPermissionsCanConfigureMaps(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerConfigPermissionsCanConfigureNetflow(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerConfigPermissionsCanConfigureSMTP(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerConfigPermissionsCanConfigureUcrm(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerConfigPermissionsCanDownloadSupportInfo(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerConfigPermissionsCanRunLocalDiscovery(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerConfigPermissionsCanSetupWithoutAuthentication(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerConfigPermissionsCanSkipPrivacyPolicyAgreement(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerConfigPermissionsCanSkipSetupSurvey(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerConfigPermissionsCanUpdateUnms(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerConfigPermissionsCanUseCloudVault(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerConfigPermissionsCanUseSsoLogin(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerServerConfigPermissionsCanConfigureCertificate(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	canConfigureCertificateDescription := `Required. Whether SSL certificate can be configured or not.`

	var canConfigureCertificateFlagName string
	if cmdPrefix == "" {
		canConfigureCertificateFlagName = "canConfigureCertificate"
	} else {
		canConfigureCertificateFlagName = fmt.Sprintf("%v.canConfigureCertificate", cmdPrefix)
	}

	var canConfigureCertificateFlagDefault bool

	_ = cmd.PersistentFlags().Bool(canConfigureCertificateFlagName, canConfigureCertificateFlagDefault, canConfigureCertificateDescription)

	return nil
}

func registerServerConfigPermissionsCanConfigureDeviceProfiles(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	canConfigureDeviceProfilesDescription := `Required. Whether device profiles can be configured or not.`

	var canConfigureDeviceProfilesFlagName string
	if cmdPrefix == "" {
		canConfigureDeviceProfilesFlagName = "canConfigureDeviceProfiles"
	} else {
		canConfigureDeviceProfilesFlagName = fmt.Sprintf("%v.canConfigureDeviceProfiles", cmdPrefix)
	}

	var canConfigureDeviceProfilesFlagDefault bool

	_ = cmd.PersistentFlags().Bool(canConfigureDeviceProfilesFlagName, canConfigureDeviceProfilesFlagDefault, canConfigureDeviceProfilesDescription)

	return nil
}

func registerServerConfigPermissionsCanConfigureFirstDeviceInSetup(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	canConfigureFirstDeviceInSetupDescription := `Required. Whether UISP setup wizard can show the First device screen.`

	var canConfigureFirstDeviceInSetupFlagName string
	if cmdPrefix == "" {
		canConfigureFirstDeviceInSetupFlagName = "canConfigureFirstDeviceInSetup"
	} else {
		canConfigureFirstDeviceInSetupFlagName = fmt.Sprintf("%v.canConfigureFirstDeviceInSetup", cmdPrefix)
	}

	var canConfigureFirstDeviceInSetupFlagDefault bool

	_ = cmd.PersistentFlags().Bool(canConfigureFirstDeviceInSetupFlagName, canConfigureFirstDeviceInSetupFlagDefault, canConfigureFirstDeviceInSetupDescription)

	return nil
}

func registerServerConfigPermissionsCanConfigureHostname(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	canConfigureHostnameDescription := `Required. Whether hostname can be configured or not.`

	var canConfigureHostnameFlagName string
	if cmdPrefix == "" {
		canConfigureHostnameFlagName = "canConfigureHostname"
	} else {
		canConfigureHostnameFlagName = fmt.Sprintf("%v.canConfigureHostname", cmdPrefix)
	}

	var canConfigureHostnameFlagDefault bool

	_ = cmd.PersistentFlags().Bool(canConfigureHostnameFlagName, canConfigureHostnameFlagDefault, canConfigureHostnameDescription)

	return nil
}

func registerServerConfigPermissionsCanConfigureMaps(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	canConfigureMapsDescription := `Required. Whether map engine is configurable or not.`

	var canConfigureMapsFlagName string
	if cmdPrefix == "" {
		canConfigureMapsFlagName = "canConfigureMaps"
	} else {
		canConfigureMapsFlagName = fmt.Sprintf("%v.canConfigureMaps", cmdPrefix)
	}

	var canConfigureMapsFlagDefault bool

	_ = cmd.PersistentFlags().Bool(canConfigureMapsFlagName, canConfigureMapsFlagDefault, canConfigureMapsDescription)

	return nil
}

func registerServerConfigPermissionsCanConfigureNetflow(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	canConfigureNetflowDescription := `Required. Whether netflow can be configured or not.`

	var canConfigureNetflowFlagName string
	if cmdPrefix == "" {
		canConfigureNetflowFlagName = "canConfigureNetflow"
	} else {
		canConfigureNetflowFlagName = fmt.Sprintf("%v.canConfigureNetflow", cmdPrefix)
	}

	var canConfigureNetflowFlagDefault bool

	_ = cmd.PersistentFlags().Bool(canConfigureNetflowFlagName, canConfigureNetflowFlagDefault, canConfigureNetflowDescription)

	return nil
}

func registerServerConfigPermissionsCanConfigureSMTP(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	canConfigureSmtpDescription := `Required. Whether SMTP can be configured or not.`

	var canConfigureSmtpFlagName string
	if cmdPrefix == "" {
		canConfigureSmtpFlagName = "canConfigureSmtp"
	} else {
		canConfigureSmtpFlagName = fmt.Sprintf("%v.canConfigureSmtp", cmdPrefix)
	}

	var canConfigureSmtpFlagDefault bool

	_ = cmd.PersistentFlags().Bool(canConfigureSmtpFlagName, canConfigureSmtpFlagDefault, canConfigureSmtpDescription)

	return nil
}

func registerServerConfigPermissionsCanConfigureUcrm(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	canConfigureUcrmDescription := `Required. Whether CRM can be configured from UISP UI or not.`

	var canConfigureUcrmFlagName string
	if cmdPrefix == "" {
		canConfigureUcrmFlagName = "canConfigureUcrm"
	} else {
		canConfigureUcrmFlagName = fmt.Sprintf("%v.canConfigureUcrm", cmdPrefix)
	}

	var canConfigureUcrmFlagDefault bool

	_ = cmd.PersistentFlags().Bool(canConfigureUcrmFlagName, canConfigureUcrmFlagDefault, canConfigureUcrmDescription)

	return nil
}

func registerServerConfigPermissionsCanDownloadSupportInfo(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	canDownloadSupportInfoDescription := `Required. Whether support info can be downloaded or not.`

	var canDownloadSupportInfoFlagName string
	if cmdPrefix == "" {
		canDownloadSupportInfoFlagName = "canDownloadSupportInfo"
	} else {
		canDownloadSupportInfoFlagName = fmt.Sprintf("%v.canDownloadSupportInfo", cmdPrefix)
	}

	var canDownloadSupportInfoFlagDefault bool

	_ = cmd.PersistentFlags().Bool(canDownloadSupportInfoFlagName, canDownloadSupportInfoFlagDefault, canDownloadSupportInfoDescription)

	return nil
}

func registerServerConfigPermissionsCanRunLocalDiscovery(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	canRunLocalDiscoveryDescription := `Required. Whether UISP can discover devices in local network.`

	var canRunLocalDiscoveryFlagName string
	if cmdPrefix == "" {
		canRunLocalDiscoveryFlagName = "canRunLocalDiscovery"
	} else {
		canRunLocalDiscoveryFlagName = fmt.Sprintf("%v.canRunLocalDiscovery", cmdPrefix)
	}

	var canRunLocalDiscoveryFlagDefault bool

	_ = cmd.PersistentFlags().Bool(canRunLocalDiscoveryFlagName, canRunLocalDiscoveryFlagDefault, canRunLocalDiscoveryDescription)

	return nil
}

func registerServerConfigPermissionsCanSetupWithoutAuthentication(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	canSetupWithoutAuthenticationDescription := `Required. Whether UISP setup wizard can be completed without authentication.`

	var canSetupWithoutAuthenticationFlagName string
	if cmdPrefix == "" {
		canSetupWithoutAuthenticationFlagName = "canSetupWithoutAuthentication"
	} else {
		canSetupWithoutAuthenticationFlagName = fmt.Sprintf("%v.canSetupWithoutAuthentication", cmdPrefix)
	}

	var canSetupWithoutAuthenticationFlagDefault bool

	_ = cmd.PersistentFlags().Bool(canSetupWithoutAuthenticationFlagName, canSetupWithoutAuthenticationFlagDefault, canSetupWithoutAuthenticationDescription)

	return nil
}

func registerServerConfigPermissionsCanSkipPrivacyPolicyAgreement(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	canSkipPrivacyPolicyAgreementDescription := `Required. Whether UISP setup can skip privacy policy agreement screen.`

	var canSkipPrivacyPolicyAgreementFlagName string
	if cmdPrefix == "" {
		canSkipPrivacyPolicyAgreementFlagName = "canSkipPrivacyPolicyAgreement"
	} else {
		canSkipPrivacyPolicyAgreementFlagName = fmt.Sprintf("%v.canSkipPrivacyPolicyAgreement", cmdPrefix)
	}

	var canSkipPrivacyPolicyAgreementFlagDefault bool

	_ = cmd.PersistentFlags().Bool(canSkipPrivacyPolicyAgreementFlagName, canSkipPrivacyPolicyAgreementFlagDefault, canSkipPrivacyPolicyAgreementDescription)

	return nil
}

func registerServerConfigPermissionsCanSkipSetupSurvey(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	canSkipSetupSurveyDescription := `Required. Whether UISP can skip setup survey after setup.`

	var canSkipSetupSurveyFlagName string
	if cmdPrefix == "" {
		canSkipSetupSurveyFlagName = "canSkipSetupSurvey"
	} else {
		canSkipSetupSurveyFlagName = fmt.Sprintf("%v.canSkipSetupSurvey", cmdPrefix)
	}

	var canSkipSetupSurveyFlagDefault bool

	_ = cmd.PersistentFlags().Bool(canSkipSetupSurveyFlagName, canSkipSetupSurveyFlagDefault, canSkipSetupSurveyDescription)

	return nil
}

func registerServerConfigPermissionsCanUpdateUnms(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	canUpdateUnmsDescription := `Required. Whether UISP can be updated from UISP UI or not.`

	var canUpdateUnmsFlagName string
	if cmdPrefix == "" {
		canUpdateUnmsFlagName = "canUpdateUnms"
	} else {
		canUpdateUnmsFlagName = fmt.Sprintf("%v.canUpdateUnms", cmdPrefix)
	}

	var canUpdateUnmsFlagDefault bool

	_ = cmd.PersistentFlags().Bool(canUpdateUnmsFlagName, canUpdateUnmsFlagDefault, canUpdateUnmsDescription)

	return nil
}

func registerServerConfigPermissionsCanUseCloudVault(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	canUseCloudVaultDescription := `Required. Whether vault key can be managed by cloud.`

	var canUseCloudVaultFlagName string
	if cmdPrefix == "" {
		canUseCloudVaultFlagName = "canUseCloudVault"
	} else {
		canUseCloudVaultFlagName = fmt.Sprintf("%v.canUseCloudVault", cmdPrefix)
	}

	var canUseCloudVaultFlagDefault bool

	_ = cmd.PersistentFlags().Bool(canUseCloudVaultFlagName, canUseCloudVaultFlagDefault, canUseCloudVaultDescription)

	return nil
}

func registerServerConfigPermissionsCanUseSsoLogin(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	canUseSsoLoginDescription := `Required. Whether users can login using Ubiquiti SSO.`

	var canUseSsoLoginFlagName string
	if cmdPrefix == "" {
		canUseSsoLoginFlagName = "canUseSsoLogin"
	} else {
		canUseSsoLoginFlagName = fmt.Sprintf("%v.canUseSsoLogin", cmdPrefix)
	}

	var canUseSsoLoginFlagDefault bool

	_ = cmd.PersistentFlags().Bool(canUseSsoLoginFlagName, canUseSsoLoginFlagDefault, canUseSsoLoginDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServerConfigPermissionsFlags(depth int, m *models.ServerConfigPermissions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, canConfigureCertificateAdded := retrieveServerConfigPermissionsCanConfigureCertificateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || canConfigureCertificateAdded

	err, canConfigureDeviceProfilesAdded := retrieveServerConfigPermissionsCanConfigureDeviceProfilesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || canConfigureDeviceProfilesAdded

	err, canConfigureFirstDeviceInSetupAdded := retrieveServerConfigPermissionsCanConfigureFirstDeviceInSetupFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || canConfigureFirstDeviceInSetupAdded

	err, canConfigureHostnameAdded := retrieveServerConfigPermissionsCanConfigureHostnameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || canConfigureHostnameAdded

	err, canConfigureMapsAdded := retrieveServerConfigPermissionsCanConfigureMapsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || canConfigureMapsAdded

	err, canConfigureNetflowAdded := retrieveServerConfigPermissionsCanConfigureNetflowFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || canConfigureNetflowAdded

	err, canConfigureSmtpAdded := retrieveServerConfigPermissionsCanConfigureSMTPFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || canConfigureSmtpAdded

	err, canConfigureUcrmAdded := retrieveServerConfigPermissionsCanConfigureUcrmFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || canConfigureUcrmAdded

	err, canDownloadSupportInfoAdded := retrieveServerConfigPermissionsCanDownloadSupportInfoFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || canDownloadSupportInfoAdded

	err, canRunLocalDiscoveryAdded := retrieveServerConfigPermissionsCanRunLocalDiscoveryFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || canRunLocalDiscoveryAdded

	err, canSetupWithoutAuthenticationAdded := retrieveServerConfigPermissionsCanSetupWithoutAuthenticationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || canSetupWithoutAuthenticationAdded

	err, canSkipPrivacyPolicyAgreementAdded := retrieveServerConfigPermissionsCanSkipPrivacyPolicyAgreementFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || canSkipPrivacyPolicyAgreementAdded

	err, canSkipSetupSurveyAdded := retrieveServerConfigPermissionsCanSkipSetupSurveyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || canSkipSetupSurveyAdded

	err, canUpdateUnmsAdded := retrieveServerConfigPermissionsCanUpdateUnmsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || canUpdateUnmsAdded

	err, canUseCloudVaultAdded := retrieveServerConfigPermissionsCanUseCloudVaultFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || canUseCloudVaultAdded

	err, canUseSsoLoginAdded := retrieveServerConfigPermissionsCanUseSsoLoginFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || canUseSsoLoginAdded

	return nil, retAdded
}

func retrieveServerConfigPermissionsCanConfigureCertificateFlags(depth int, m *models.ServerConfigPermissions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	canConfigureCertificateFlagName := fmt.Sprintf("%v.canConfigureCertificate", cmdPrefix)
	if cmd.Flags().Changed(canConfigureCertificateFlagName) {

		var canConfigureCertificateFlagName string
		if cmdPrefix == "" {
			canConfigureCertificateFlagName = "canConfigureCertificate"
		} else {
			canConfigureCertificateFlagName = fmt.Sprintf("%v.canConfigureCertificate", cmdPrefix)
		}

		canConfigureCertificateFlagValue, err := cmd.Flags().GetBool(canConfigureCertificateFlagName)
		if err != nil {
			return err, false
		}
		m.CanConfigureCertificate = &canConfigureCertificateFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerConfigPermissionsCanConfigureDeviceProfilesFlags(depth int, m *models.ServerConfigPermissions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	canConfigureDeviceProfilesFlagName := fmt.Sprintf("%v.canConfigureDeviceProfiles", cmdPrefix)
	if cmd.Flags().Changed(canConfigureDeviceProfilesFlagName) {

		var canConfigureDeviceProfilesFlagName string
		if cmdPrefix == "" {
			canConfigureDeviceProfilesFlagName = "canConfigureDeviceProfiles"
		} else {
			canConfigureDeviceProfilesFlagName = fmt.Sprintf("%v.canConfigureDeviceProfiles", cmdPrefix)
		}

		canConfigureDeviceProfilesFlagValue, err := cmd.Flags().GetBool(canConfigureDeviceProfilesFlagName)
		if err != nil {
			return err, false
		}
		m.CanConfigureDeviceProfiles = &canConfigureDeviceProfilesFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerConfigPermissionsCanConfigureFirstDeviceInSetupFlags(depth int, m *models.ServerConfigPermissions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	canConfigureFirstDeviceInSetupFlagName := fmt.Sprintf("%v.canConfigureFirstDeviceInSetup", cmdPrefix)
	if cmd.Flags().Changed(canConfigureFirstDeviceInSetupFlagName) {

		var canConfigureFirstDeviceInSetupFlagName string
		if cmdPrefix == "" {
			canConfigureFirstDeviceInSetupFlagName = "canConfigureFirstDeviceInSetup"
		} else {
			canConfigureFirstDeviceInSetupFlagName = fmt.Sprintf("%v.canConfigureFirstDeviceInSetup", cmdPrefix)
		}

		canConfigureFirstDeviceInSetupFlagValue, err := cmd.Flags().GetBool(canConfigureFirstDeviceInSetupFlagName)
		if err != nil {
			return err, false
		}
		m.CanConfigureFirstDeviceInSetup = &canConfigureFirstDeviceInSetupFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerConfigPermissionsCanConfigureHostnameFlags(depth int, m *models.ServerConfigPermissions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	canConfigureHostnameFlagName := fmt.Sprintf("%v.canConfigureHostname", cmdPrefix)
	if cmd.Flags().Changed(canConfigureHostnameFlagName) {

		var canConfigureHostnameFlagName string
		if cmdPrefix == "" {
			canConfigureHostnameFlagName = "canConfigureHostname"
		} else {
			canConfigureHostnameFlagName = fmt.Sprintf("%v.canConfigureHostname", cmdPrefix)
		}

		canConfigureHostnameFlagValue, err := cmd.Flags().GetBool(canConfigureHostnameFlagName)
		if err != nil {
			return err, false
		}
		m.CanConfigureHostname = &canConfigureHostnameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerConfigPermissionsCanConfigureMapsFlags(depth int, m *models.ServerConfigPermissions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	canConfigureMapsFlagName := fmt.Sprintf("%v.canConfigureMaps", cmdPrefix)
	if cmd.Flags().Changed(canConfigureMapsFlagName) {

		var canConfigureMapsFlagName string
		if cmdPrefix == "" {
			canConfigureMapsFlagName = "canConfigureMaps"
		} else {
			canConfigureMapsFlagName = fmt.Sprintf("%v.canConfigureMaps", cmdPrefix)
		}

		canConfigureMapsFlagValue, err := cmd.Flags().GetBool(canConfigureMapsFlagName)
		if err != nil {
			return err, false
		}
		m.CanConfigureMaps = &canConfigureMapsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerConfigPermissionsCanConfigureNetflowFlags(depth int, m *models.ServerConfigPermissions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	canConfigureNetflowFlagName := fmt.Sprintf("%v.canConfigureNetflow", cmdPrefix)
	if cmd.Flags().Changed(canConfigureNetflowFlagName) {

		var canConfigureNetflowFlagName string
		if cmdPrefix == "" {
			canConfigureNetflowFlagName = "canConfigureNetflow"
		} else {
			canConfigureNetflowFlagName = fmt.Sprintf("%v.canConfigureNetflow", cmdPrefix)
		}

		canConfigureNetflowFlagValue, err := cmd.Flags().GetBool(canConfigureNetflowFlagName)
		if err != nil {
			return err, false
		}
		m.CanConfigureNetflow = &canConfigureNetflowFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerConfigPermissionsCanConfigureSMTPFlags(depth int, m *models.ServerConfigPermissions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	canConfigureSmtpFlagName := fmt.Sprintf("%v.canConfigureSmtp", cmdPrefix)
	if cmd.Flags().Changed(canConfigureSmtpFlagName) {

		var canConfigureSmtpFlagName string
		if cmdPrefix == "" {
			canConfigureSmtpFlagName = "canConfigureSmtp"
		} else {
			canConfigureSmtpFlagName = fmt.Sprintf("%v.canConfigureSmtp", cmdPrefix)
		}

		canConfigureSmtpFlagValue, err := cmd.Flags().GetBool(canConfigureSmtpFlagName)
		if err != nil {
			return err, false
		}
		m.CanConfigureSMTP = &canConfigureSmtpFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerConfigPermissionsCanConfigureUcrmFlags(depth int, m *models.ServerConfigPermissions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	canConfigureUcrmFlagName := fmt.Sprintf("%v.canConfigureUcrm", cmdPrefix)
	if cmd.Flags().Changed(canConfigureUcrmFlagName) {

		var canConfigureUcrmFlagName string
		if cmdPrefix == "" {
			canConfigureUcrmFlagName = "canConfigureUcrm"
		} else {
			canConfigureUcrmFlagName = fmt.Sprintf("%v.canConfigureUcrm", cmdPrefix)
		}

		canConfigureUcrmFlagValue, err := cmd.Flags().GetBool(canConfigureUcrmFlagName)
		if err != nil {
			return err, false
		}
		m.CanConfigureUcrm = &canConfigureUcrmFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerConfigPermissionsCanDownloadSupportInfoFlags(depth int, m *models.ServerConfigPermissions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	canDownloadSupportInfoFlagName := fmt.Sprintf("%v.canDownloadSupportInfo", cmdPrefix)
	if cmd.Flags().Changed(canDownloadSupportInfoFlagName) {

		var canDownloadSupportInfoFlagName string
		if cmdPrefix == "" {
			canDownloadSupportInfoFlagName = "canDownloadSupportInfo"
		} else {
			canDownloadSupportInfoFlagName = fmt.Sprintf("%v.canDownloadSupportInfo", cmdPrefix)
		}

		canDownloadSupportInfoFlagValue, err := cmd.Flags().GetBool(canDownloadSupportInfoFlagName)
		if err != nil {
			return err, false
		}
		m.CanDownloadSupportInfo = &canDownloadSupportInfoFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerConfigPermissionsCanRunLocalDiscoveryFlags(depth int, m *models.ServerConfigPermissions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	canRunLocalDiscoveryFlagName := fmt.Sprintf("%v.canRunLocalDiscovery", cmdPrefix)
	if cmd.Flags().Changed(canRunLocalDiscoveryFlagName) {

		var canRunLocalDiscoveryFlagName string
		if cmdPrefix == "" {
			canRunLocalDiscoveryFlagName = "canRunLocalDiscovery"
		} else {
			canRunLocalDiscoveryFlagName = fmt.Sprintf("%v.canRunLocalDiscovery", cmdPrefix)
		}

		canRunLocalDiscoveryFlagValue, err := cmd.Flags().GetBool(canRunLocalDiscoveryFlagName)
		if err != nil {
			return err, false
		}
		m.CanRunLocalDiscovery = &canRunLocalDiscoveryFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerConfigPermissionsCanSetupWithoutAuthenticationFlags(depth int, m *models.ServerConfigPermissions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	canSetupWithoutAuthenticationFlagName := fmt.Sprintf("%v.canSetupWithoutAuthentication", cmdPrefix)
	if cmd.Flags().Changed(canSetupWithoutAuthenticationFlagName) {

		var canSetupWithoutAuthenticationFlagName string
		if cmdPrefix == "" {
			canSetupWithoutAuthenticationFlagName = "canSetupWithoutAuthentication"
		} else {
			canSetupWithoutAuthenticationFlagName = fmt.Sprintf("%v.canSetupWithoutAuthentication", cmdPrefix)
		}

		canSetupWithoutAuthenticationFlagValue, err := cmd.Flags().GetBool(canSetupWithoutAuthenticationFlagName)
		if err != nil {
			return err, false
		}
		m.CanSetupWithoutAuthentication = &canSetupWithoutAuthenticationFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerConfigPermissionsCanSkipPrivacyPolicyAgreementFlags(depth int, m *models.ServerConfigPermissions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	canSkipPrivacyPolicyAgreementFlagName := fmt.Sprintf("%v.canSkipPrivacyPolicyAgreement", cmdPrefix)
	if cmd.Flags().Changed(canSkipPrivacyPolicyAgreementFlagName) {

		var canSkipPrivacyPolicyAgreementFlagName string
		if cmdPrefix == "" {
			canSkipPrivacyPolicyAgreementFlagName = "canSkipPrivacyPolicyAgreement"
		} else {
			canSkipPrivacyPolicyAgreementFlagName = fmt.Sprintf("%v.canSkipPrivacyPolicyAgreement", cmdPrefix)
		}

		canSkipPrivacyPolicyAgreementFlagValue, err := cmd.Flags().GetBool(canSkipPrivacyPolicyAgreementFlagName)
		if err != nil {
			return err, false
		}
		m.CanSkipPrivacyPolicyAgreement = &canSkipPrivacyPolicyAgreementFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerConfigPermissionsCanSkipSetupSurveyFlags(depth int, m *models.ServerConfigPermissions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	canSkipSetupSurveyFlagName := fmt.Sprintf("%v.canSkipSetupSurvey", cmdPrefix)
	if cmd.Flags().Changed(canSkipSetupSurveyFlagName) {

		var canSkipSetupSurveyFlagName string
		if cmdPrefix == "" {
			canSkipSetupSurveyFlagName = "canSkipSetupSurvey"
		} else {
			canSkipSetupSurveyFlagName = fmt.Sprintf("%v.canSkipSetupSurvey", cmdPrefix)
		}

		canSkipSetupSurveyFlagValue, err := cmd.Flags().GetBool(canSkipSetupSurveyFlagName)
		if err != nil {
			return err, false
		}
		m.CanSkipSetupSurvey = &canSkipSetupSurveyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerConfigPermissionsCanUpdateUnmsFlags(depth int, m *models.ServerConfigPermissions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	canUpdateUnmsFlagName := fmt.Sprintf("%v.canUpdateUnms", cmdPrefix)
	if cmd.Flags().Changed(canUpdateUnmsFlagName) {

		var canUpdateUnmsFlagName string
		if cmdPrefix == "" {
			canUpdateUnmsFlagName = "canUpdateUnms"
		} else {
			canUpdateUnmsFlagName = fmt.Sprintf("%v.canUpdateUnms", cmdPrefix)
		}

		canUpdateUnmsFlagValue, err := cmd.Flags().GetBool(canUpdateUnmsFlagName)
		if err != nil {
			return err, false
		}
		m.CanUpdateUnms = &canUpdateUnmsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerConfigPermissionsCanUseCloudVaultFlags(depth int, m *models.ServerConfigPermissions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	canUseCloudVaultFlagName := fmt.Sprintf("%v.canUseCloudVault", cmdPrefix)
	if cmd.Flags().Changed(canUseCloudVaultFlagName) {

		var canUseCloudVaultFlagName string
		if cmdPrefix == "" {
			canUseCloudVaultFlagName = "canUseCloudVault"
		} else {
			canUseCloudVaultFlagName = fmt.Sprintf("%v.canUseCloudVault", cmdPrefix)
		}

		canUseCloudVaultFlagValue, err := cmd.Flags().GetBool(canUseCloudVaultFlagName)
		if err != nil {
			return err, false
		}
		m.CanUseCloudVault = &canUseCloudVaultFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerConfigPermissionsCanUseSsoLoginFlags(depth int, m *models.ServerConfigPermissions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	canUseSsoLoginFlagName := fmt.Sprintf("%v.canUseSsoLogin", cmdPrefix)
	if cmd.Flags().Changed(canUseSsoLoginFlagName) {

		var canUseSsoLoginFlagName string
		if cmdPrefix == "" {
			canUseSsoLoginFlagName = "canUseSsoLogin"
		} else {
			canUseSsoLoginFlagName = fmt.Sprintf("%v.canUseSsoLogin", cmdPrefix)
		}

		canUseSsoLoginFlagValue, err := cmd.Flags().GetBool(canUseSsoLoginFlagName)
		if err != nil {
			return err, false
		}
		m.CanUseSsoLogin = &canUseSsoLoginFlagValue

		retAdded = true
	}

	return nil, retAdded
}
