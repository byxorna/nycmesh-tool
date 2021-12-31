// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for Wifi5Ghz

// register flags to command
func registerModelWifi5GhzFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerWifi5GhzAuthentication(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWifi5GhzAvailable(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWifi5GhzChannel(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWifi5GhzChannelWidth(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWifi5GhzCountry(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWifi5GhzEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWifi5GhzEncryption(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWifi5GhzFrequency(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWifi5GhzIsChannelAuto(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWifi5GhzIsWPA2PSKEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWifi5GhzKey(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWifi5GhzMac(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWifi5GhzMode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWifi5GhzSsid(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWifi5GhzTxPower(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerWifi5GhzAuthentication(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	authenticationDescription := `Enum: ["psk","psk2","ent","none"]. Required. `

	var authenticationFlagName string
	if cmdPrefix == "" {
		authenticationFlagName = "authentication"
	} else {
		authenticationFlagName = fmt.Sprintf("%v.authentication", cmdPrefix)
	}

	var authenticationFlagDefault string

	_ = cmd.PersistentFlags().String(authenticationFlagName, authenticationFlagDefault, authenticationDescription)

	if err := cmd.RegisterFlagCompletionFunc(authenticationFlagName,
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

func registerWifi5GhzAvailable(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	availableDescription := `Required. `

	var availableFlagName string
	if cmdPrefix == "" {
		availableFlagName = "available"
	} else {
		availableFlagName = fmt.Sprintf("%v.available", cmdPrefix)
	}

	var availableFlagDefault bool

	_ = cmd.PersistentFlags().Bool(availableFlagName, availableFlagDefault, availableDescription)

	return nil
}

func registerWifi5GhzChannel(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	channelDescription := `Required. `

	var channelFlagName string
	if cmdPrefix == "" {
		channelFlagName = "channel"
	} else {
		channelFlagName = fmt.Sprintf("%v.channel", cmdPrefix)
	}

	var channelFlagDefault float64

	_ = cmd.PersistentFlags().Float64(channelFlagName, channelFlagDefault, channelDescription)

	return nil
}

func registerWifi5GhzChannelWidth(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	channelWidthDescription := `Required. `

	var channelWidthFlagName string
	if cmdPrefix == "" {
		channelWidthFlagName = "channelWidth"
	} else {
		channelWidthFlagName = fmt.Sprintf("%v.channelWidth", cmdPrefix)
	}

	var channelWidthFlagDefault float64

	_ = cmd.PersistentFlags().Float64(channelWidthFlagName, channelWidthFlagDefault, channelWidthDescription)

	return nil
}

func registerWifi5GhzCountry(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	countryDescription := ``

	var countryFlagName string
	if cmdPrefix == "" {
		countryFlagName = "country"
	} else {
		countryFlagName = fmt.Sprintf("%v.country", cmdPrefix)
	}

	var countryFlagDefault string

	_ = cmd.PersistentFlags().String(countryFlagName, countryFlagDefault, countryDescription)

	return nil
}

func registerWifi5GhzEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	enabledDescription := `Required. `

	var enabledFlagName string
	if cmdPrefix == "" {
		enabledFlagName = "enabled"
	} else {
		enabledFlagName = fmt.Sprintf("%v.enabled", cmdPrefix)
	}

	var enabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(enabledFlagName, enabledFlagDefault, enabledDescription)

	return nil
}

func registerWifi5GhzEncryption(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	encryptionDescription := `Enum: ["wep","wpa","wpa-psk","wpa2","enabled","none"]. Required. `

	var encryptionFlagName string
	if cmdPrefix == "" {
		encryptionFlagName = "encryption"
	} else {
		encryptionFlagName = fmt.Sprintf("%v.encryption", cmdPrefix)
	}

	var encryptionFlagDefault string

	_ = cmd.PersistentFlags().String(encryptionFlagName, encryptionFlagDefault, encryptionDescription)

	if err := cmd.RegisterFlagCompletionFunc(encryptionFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["wep","wpa","wpa-psk","wpa2","enabled","none"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerWifi5GhzFrequency(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	frequencyDescription := ``

	var frequencyFlagName string
	if cmdPrefix == "" {
		frequencyFlagName = "frequency"
	} else {
		frequencyFlagName = fmt.Sprintf("%v.frequency", cmdPrefix)
	}

	var frequencyFlagDefault float64

	_ = cmd.PersistentFlags().Float64(frequencyFlagName, frequencyFlagDefault, frequencyDescription)

	return nil
}

func registerWifi5GhzIsChannelAuto(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	isChannelAutoDescription := ``

	var isChannelAutoFlagName string
	if cmdPrefix == "" {
		isChannelAutoFlagName = "isChannelAuto"
	} else {
		isChannelAutoFlagName = fmt.Sprintf("%v.isChannelAuto", cmdPrefix)
	}

	var isChannelAutoFlagDefault bool

	_ = cmd.PersistentFlags().Bool(isChannelAutoFlagName, isChannelAutoFlagDefault, isChannelAutoDescription)

	return nil
}

func registerWifi5GhzIsWPA2PSKEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	isWPA2PSKEnabledDescription := ``

	var isWPA2PSKEnabledFlagName string
	if cmdPrefix == "" {
		isWPA2PSKEnabledFlagName = "isWPA2PSKEnabled"
	} else {
		isWPA2PSKEnabledFlagName = fmt.Sprintf("%v.isWPA2PSKEnabled", cmdPrefix)
	}

	var isWPA2PSKEnabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(isWPA2PSKEnabledFlagName, isWPA2PSKEnabledFlagDefault, isWPA2PSKEnabledDescription)

	return nil
}

func registerWifi5GhzKey(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	keyDescription := `Required. `

	var keyFlagName string
	if cmdPrefix == "" {
		keyFlagName = "key"
	} else {
		keyFlagName = fmt.Sprintf("%v.key", cmdPrefix)
	}

	var keyFlagDefault string

	_ = cmd.PersistentFlags().String(keyFlagName, keyFlagDefault, keyDescription)

	return nil
}

func registerWifi5GhzMac(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	macDescription := ``

	var macFlagName string
	if cmdPrefix == "" {
		macFlagName = "mac"
	} else {
		macFlagName = fmt.Sprintf("%v.mac", cmdPrefix)
	}

	var macFlagDefault string

	_ = cmd.PersistentFlags().String(macFlagName, macFlagDefault, macDescription)

	return nil
}

func registerWifi5GhzMode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	modeDescription := `Enum: ["ap","ap-ptp","ap-ptmp","ap-ptmp-airmax","ap-ptmp-airmax-mixed","ap-ptmp-airmax-ac","sta","sta-ptp","sta-ptmp","aprepeater","repeater","mesh"]. Required. `

	var modeFlagName string
	if cmdPrefix == "" {
		modeFlagName = "mode"
	} else {
		modeFlagName = fmt.Sprintf("%v.mode", cmdPrefix)
	}

	var modeFlagDefault string

	_ = cmd.PersistentFlags().String(modeFlagName, modeFlagDefault, modeDescription)

	if err := cmd.RegisterFlagCompletionFunc(modeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["ap","ap-ptp","ap-ptmp","ap-ptmp-airmax","ap-ptmp-airmax-mixed","ap-ptmp-airmax-ac","sta","sta-ptp","sta-ptmp","aprepeater","repeater","mesh"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerWifi5GhzSsid(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ssidDescription := `Required. `

	var ssidFlagName string
	if cmdPrefix == "" {
		ssidFlagName = "ssid"
	} else {
		ssidFlagName = fmt.Sprintf("%v.ssid", cmdPrefix)
	}

	var ssidFlagDefault string

	_ = cmd.PersistentFlags().String(ssidFlagName, ssidFlagDefault, ssidDescription)

	return nil
}

func registerWifi5GhzTxPower(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	txPowerDescription := `Required. `

	var txPowerFlagName string
	if cmdPrefix == "" {
		txPowerFlagName = "txPower"
	} else {
		txPowerFlagName = fmt.Sprintf("%v.txPower", cmdPrefix)
	}

	var txPowerFlagDefault float64

	_ = cmd.PersistentFlags().Float64(txPowerFlagName, txPowerFlagDefault, txPowerDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelWifi5GhzFlags(depth int, m *models.Wifi5Ghz, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, authenticationAdded := retrieveWifi5GhzAuthenticationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || authenticationAdded

	err, availableAdded := retrieveWifi5GhzAvailableFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || availableAdded

	err, channelAdded := retrieveWifi5GhzChannelFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || channelAdded

	err, channelWidthAdded := retrieveWifi5GhzChannelWidthFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || channelWidthAdded

	err, countryAdded := retrieveWifi5GhzCountryFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || countryAdded

	err, enabledAdded := retrieveWifi5GhzEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || enabledAdded

	err, encryptionAdded := retrieveWifi5GhzEncryptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || encryptionAdded

	err, frequencyAdded := retrieveWifi5GhzFrequencyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || frequencyAdded

	err, isChannelAutoAdded := retrieveWifi5GhzIsChannelAutoFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || isChannelAutoAdded

	err, isWPA2PSKEnabledAdded := retrieveWifi5GhzIsWPA2PSKEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || isWPA2PSKEnabledAdded

	err, keyAdded := retrieveWifi5GhzKeyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || keyAdded

	err, macAdded := retrieveWifi5GhzMacFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || macAdded

	err, modeAdded := retrieveWifi5GhzModeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || modeAdded

	err, ssidAdded := retrieveWifi5GhzSsidFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ssidAdded

	err, txPowerAdded := retrieveWifi5GhzTxPowerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || txPowerAdded

	return nil, retAdded
}

func retrieveWifi5GhzAuthenticationFlags(depth int, m *models.Wifi5Ghz, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	authenticationFlagName := fmt.Sprintf("%v.authentication", cmdPrefix)
	if cmd.Flags().Changed(authenticationFlagName) {

		var authenticationFlagName string
		if cmdPrefix == "" {
			authenticationFlagName = "authentication"
		} else {
			authenticationFlagName = fmt.Sprintf("%v.authentication", cmdPrefix)
		}

		authenticationFlagValue, err := cmd.Flags().GetString(authenticationFlagName)
		if err != nil {
			return err, false
		}
		m.Authentication = &authenticationFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWifi5GhzAvailableFlags(depth int, m *models.Wifi5Ghz, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	availableFlagName := fmt.Sprintf("%v.available", cmdPrefix)
	if cmd.Flags().Changed(availableFlagName) {

		var availableFlagName string
		if cmdPrefix == "" {
			availableFlagName = "available"
		} else {
			availableFlagName = fmt.Sprintf("%v.available", cmdPrefix)
		}

		availableFlagValue, err := cmd.Flags().GetBool(availableFlagName)
		if err != nil {
			return err, false
		}
		m.Available = &availableFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWifi5GhzChannelFlags(depth int, m *models.Wifi5Ghz, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	channelFlagName := fmt.Sprintf("%v.channel", cmdPrefix)
	if cmd.Flags().Changed(channelFlagName) {

		var channelFlagName string
		if cmdPrefix == "" {
			channelFlagName = "channel"
		} else {
			channelFlagName = fmt.Sprintf("%v.channel", cmdPrefix)
		}

		channelFlagValue, err := cmd.Flags().GetFloat64(channelFlagName)
		if err != nil {
			return err, false
		}
		m.Channel = &channelFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWifi5GhzChannelWidthFlags(depth int, m *models.Wifi5Ghz, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	channelWidthFlagName := fmt.Sprintf("%v.channelWidth", cmdPrefix)
	if cmd.Flags().Changed(channelWidthFlagName) {

		var channelWidthFlagName string
		if cmdPrefix == "" {
			channelWidthFlagName = "channelWidth"
		} else {
			channelWidthFlagName = fmt.Sprintf("%v.channelWidth", cmdPrefix)
		}

		channelWidthFlagValue, err := cmd.Flags().GetFloat64(channelWidthFlagName)
		if err != nil {
			return err, false
		}
		m.ChannelWidth = &channelWidthFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWifi5GhzCountryFlags(depth int, m *models.Wifi5Ghz, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	countryFlagName := fmt.Sprintf("%v.country", cmdPrefix)
	if cmd.Flags().Changed(countryFlagName) {

		var countryFlagName string
		if cmdPrefix == "" {
			countryFlagName = "country"
		} else {
			countryFlagName = fmt.Sprintf("%v.country", cmdPrefix)
		}

		countryFlagValue, err := cmd.Flags().GetString(countryFlagName)
		if err != nil {
			return err, false
		}
		m.Country = countryFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWifi5GhzEnabledFlags(depth int, m *models.Wifi5Ghz, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	enabledFlagName := fmt.Sprintf("%v.enabled", cmdPrefix)
	if cmd.Flags().Changed(enabledFlagName) {

		var enabledFlagName string
		if cmdPrefix == "" {
			enabledFlagName = "enabled"
		} else {
			enabledFlagName = fmt.Sprintf("%v.enabled", cmdPrefix)
		}

		enabledFlagValue, err := cmd.Flags().GetBool(enabledFlagName)
		if err != nil {
			return err, false
		}
		m.Enabled = &enabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWifi5GhzEncryptionFlags(depth int, m *models.Wifi5Ghz, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	encryptionFlagName := fmt.Sprintf("%v.encryption", cmdPrefix)
	if cmd.Flags().Changed(encryptionFlagName) {

		var encryptionFlagName string
		if cmdPrefix == "" {
			encryptionFlagName = "encryption"
		} else {
			encryptionFlagName = fmt.Sprintf("%v.encryption", cmdPrefix)
		}

		encryptionFlagValue, err := cmd.Flags().GetString(encryptionFlagName)
		if err != nil {
			return err, false
		}
		m.Encryption = &encryptionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWifi5GhzFrequencyFlags(depth int, m *models.Wifi5Ghz, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	frequencyFlagName := fmt.Sprintf("%v.frequency", cmdPrefix)
	if cmd.Flags().Changed(frequencyFlagName) {

		var frequencyFlagName string
		if cmdPrefix == "" {
			frequencyFlagName = "frequency"
		} else {
			frequencyFlagName = fmt.Sprintf("%v.frequency", cmdPrefix)
		}

		frequencyFlagValue, err := cmd.Flags().GetFloat64(frequencyFlagName)
		if err != nil {
			return err, false
		}
		m.Frequency = frequencyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWifi5GhzIsChannelAutoFlags(depth int, m *models.Wifi5Ghz, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	isChannelAutoFlagName := fmt.Sprintf("%v.isChannelAuto", cmdPrefix)
	if cmd.Flags().Changed(isChannelAutoFlagName) {

		var isChannelAutoFlagName string
		if cmdPrefix == "" {
			isChannelAutoFlagName = "isChannelAuto"
		} else {
			isChannelAutoFlagName = fmt.Sprintf("%v.isChannelAuto", cmdPrefix)
		}

		isChannelAutoFlagValue, err := cmd.Flags().GetBool(isChannelAutoFlagName)
		if err != nil {
			return err, false
		}
		m.IsChannelAuto = isChannelAutoFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWifi5GhzIsWPA2PSKEnabledFlags(depth int, m *models.Wifi5Ghz, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	isWPA2PSKEnabledFlagName := fmt.Sprintf("%v.isWPA2PSKEnabled", cmdPrefix)
	if cmd.Flags().Changed(isWPA2PSKEnabledFlagName) {

		var isWPA2PSKEnabledFlagName string
		if cmdPrefix == "" {
			isWPA2PSKEnabledFlagName = "isWPA2PSKEnabled"
		} else {
			isWPA2PSKEnabledFlagName = fmt.Sprintf("%v.isWPA2PSKEnabled", cmdPrefix)
		}

		isWPA2PSKEnabledFlagValue, err := cmd.Flags().GetBool(isWPA2PSKEnabledFlagName)
		if err != nil {
			return err, false
		}
		m.IsWPA2PSKEnabled = isWPA2PSKEnabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWifi5GhzKeyFlags(depth int, m *models.Wifi5Ghz, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	keyFlagName := fmt.Sprintf("%v.key", cmdPrefix)
	if cmd.Flags().Changed(keyFlagName) {

		var keyFlagName string
		if cmdPrefix == "" {
			keyFlagName = "key"
		} else {
			keyFlagName = fmt.Sprintf("%v.key", cmdPrefix)
		}

		keyFlagValue, err := cmd.Flags().GetString(keyFlagName)
		if err != nil {
			return err, false
		}
		m.Key = &keyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWifi5GhzMacFlags(depth int, m *models.Wifi5Ghz, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	macFlagName := fmt.Sprintf("%v.mac", cmdPrefix)
	if cmd.Flags().Changed(macFlagName) {

		var macFlagName string
		if cmdPrefix == "" {
			macFlagName = "mac"
		} else {
			macFlagName = fmt.Sprintf("%v.mac", cmdPrefix)
		}

		macFlagValue, err := cmd.Flags().GetString(macFlagName)
		if err != nil {
			return err, false
		}
		m.Mac = macFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWifi5GhzModeFlags(depth int, m *models.Wifi5Ghz, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	modeFlagName := fmt.Sprintf("%v.mode", cmdPrefix)
	if cmd.Flags().Changed(modeFlagName) {

		var modeFlagName string
		if cmdPrefix == "" {
			modeFlagName = "mode"
		} else {
			modeFlagName = fmt.Sprintf("%v.mode", cmdPrefix)
		}

		modeFlagValue, err := cmd.Flags().GetString(modeFlagName)
		if err != nil {
			return err, false
		}
		m.Mode = &modeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWifi5GhzSsidFlags(depth int, m *models.Wifi5Ghz, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ssidFlagName := fmt.Sprintf("%v.ssid", cmdPrefix)
	if cmd.Flags().Changed(ssidFlagName) {

		var ssidFlagName string
		if cmdPrefix == "" {
			ssidFlagName = "ssid"
		} else {
			ssidFlagName = fmt.Sprintf("%v.ssid", cmdPrefix)
		}

		ssidFlagValue, err := cmd.Flags().GetString(ssidFlagName)
		if err != nil {
			return err, false
		}
		m.Ssid = &ssidFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWifi5GhzTxPowerFlags(depth int, m *models.Wifi5Ghz, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	txPowerFlagName := fmt.Sprintf("%v.txPower", cmdPrefix)
	if cmd.Flags().Changed(txPowerFlagName) {

		var txPowerFlagName string
		if cmdPrefix == "" {
			txPowerFlagName = "txPower"
		} else {
			txPowerFlagName = fmt.Sprintf("%v.txPower", cmdPrefix)
		}

		txPowerFlagValue, err := cmd.Flags().GetFloat64(txPowerFlagName)
		if err != nil {
			return err, false
		}
		m.TxPower = &txPowerFlagValue

		retAdded = true
	}

	return nil, retAdded
}