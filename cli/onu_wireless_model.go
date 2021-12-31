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

// Schema cli for OnuWireless

// register flags to command
func registerModelOnuWirelessFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerOnuWirelessChannel(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuWirelessChannelWidth(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuWirelessCountry(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuWirelessCountryListID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuWirelessEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuWirelessKey(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuWirelessProvisioned(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuWirelessSsid(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOnuWirelessTxPower(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerOnuWirelessChannel(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	channelDescription := `Enum: ["auto","1","2","3","4","5","6","7","8","9","10","11","12","13"]. Required. Wireless channel, set to 'auto' or number 1-13`

	var channelFlagName string
	if cmdPrefix == "" {
		channelFlagName = "channel"
	} else {
		channelFlagName = fmt.Sprintf("%v.channel", cmdPrefix)
	}

	var channelFlagDefault string

	_ = cmd.PersistentFlags().String(channelFlagName, channelFlagDefault, channelDescription)

	if err := cmd.RegisterFlagCompletionFunc(channelFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["auto","1","2","3","4","5","6","7","8","9","10","11","12","13"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerOnuWirelessChannelWidth(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	channelWidthDescription := `Enum: ["20","40"]. Required. Wireless channel width in MHz`

	var channelWidthFlagName string
	if cmdPrefix == "" {
		channelWidthFlagName = "channelWidth"
	} else {
		channelWidthFlagName = fmt.Sprintf("%v.channelWidth", cmdPrefix)
	}

	var channelWidthFlagDefault string

	_ = cmd.PersistentFlags().String(channelWidthFlagName, channelWidthFlagDefault, channelWidthDescription)

	if err := cmd.RegisterFlagCompletionFunc(channelWidthFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["20","40"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerOnuWirelessCountry(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	countryDescription := `Required. Country for regulatory domain - ISO 3166-2 alpha code or WO for 'World wide'`

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

func registerOnuWirelessCountryListID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	countryListIdDescription := `Required. Country list id`

	var countryListIdFlagName string
	if cmdPrefix == "" {
		countryListIdFlagName = "countryListId"
	} else {
		countryListIdFlagName = fmt.Sprintf("%v.countryListId", cmdPrefix)
	}

	var countryListIdFlagDefault string

	_ = cmd.PersistentFlags().String(countryListIdFlagName, countryListIdFlagDefault, countryListIdDescription)

	return nil
}

func registerOnuWirelessEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	enabledDescription := `Required. Set to true if wireless is enabled`

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

func registerOnuWirelessKey(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	keyDescription := `Required. WPA pre-shared key`

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

func registerOnuWirelessProvisioned(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	provisionedDescription := `Required. When false, configuration of wireless part can be done directly from ONU only.`

	var provisionedFlagName string
	if cmdPrefix == "" {
		provisionedFlagName = "provisioned"
	} else {
		provisionedFlagName = fmt.Sprintf("%v.provisioned", cmdPrefix)
	}

	var provisionedFlagDefault bool

	_ = cmd.PersistentFlags().Bool(provisionedFlagName, provisionedFlagDefault, provisionedDescription)

	return nil
}

func registerOnuWirelessSsid(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ssidDescription := `Required. Wireless network SSID`

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

func registerOnuWirelessTxPower(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	txPowerDescription := `Required. TX power in %`

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
func retrieveModelOnuWirelessFlags(depth int, m *models.OnuWireless, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, channelAdded := retrieveOnuWirelessChannelFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || channelAdded

	err, channelWidthAdded := retrieveOnuWirelessChannelWidthFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || channelWidthAdded

	err, countryAdded := retrieveOnuWirelessCountryFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || countryAdded

	err, countryListIdAdded := retrieveOnuWirelessCountryListIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || countryListIdAdded

	err, enabledAdded := retrieveOnuWirelessEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || enabledAdded

	err, keyAdded := retrieveOnuWirelessKeyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || keyAdded

	err, provisionedAdded := retrieveOnuWirelessProvisionedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || provisionedAdded

	err, ssidAdded := retrieveOnuWirelessSsidFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ssidAdded

	err, txPowerAdded := retrieveOnuWirelessTxPowerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || txPowerAdded

	return nil, retAdded
}

func retrieveOnuWirelessChannelFlags(depth int, m *models.OnuWireless, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

		channelFlagValue, err := cmd.Flags().GetString(channelFlagName)
		if err != nil {
			return err, false
		}
		m.Channel = &channelFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuWirelessChannelWidthFlags(depth int, m *models.OnuWireless, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

		channelWidthFlagValue, err := cmd.Flags().GetString(channelWidthFlagName)
		if err != nil {
			return err, false
		}
		m.ChannelWidth = &channelWidthFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuWirelessCountryFlags(depth int, m *models.OnuWireless, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Country = &countryFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuWirelessCountryListIDFlags(depth int, m *models.OnuWireless, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	countryListIdFlagName := fmt.Sprintf("%v.countryListId", cmdPrefix)
	if cmd.Flags().Changed(countryListIdFlagName) {

		var countryListIdFlagName string
		if cmdPrefix == "" {
			countryListIdFlagName = "countryListId"
		} else {
			countryListIdFlagName = fmt.Sprintf("%v.countryListId", cmdPrefix)
		}

		countryListIdFlagValue, err := cmd.Flags().GetString(countryListIdFlagName)
		if err != nil {
			return err, false
		}
		m.CountryListID = &countryListIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuWirelessEnabledFlags(depth int, m *models.OnuWireless, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveOnuWirelessKeyFlags(depth int, m *models.OnuWireless, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveOnuWirelessProvisionedFlags(depth int, m *models.OnuWireless, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	provisionedFlagName := fmt.Sprintf("%v.provisioned", cmdPrefix)
	if cmd.Flags().Changed(provisionedFlagName) {

		var provisionedFlagName string
		if cmdPrefix == "" {
			provisionedFlagName = "provisioned"
		} else {
			provisionedFlagName = fmt.Sprintf("%v.provisioned", cmdPrefix)
		}

		provisionedFlagValue, err := cmd.Flags().GetBool(provisionedFlagName)
		if err != nil {
			return err, false
		}
		m.Provisioned = &provisionedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveOnuWirelessSsidFlags(depth int, m *models.OnuWireless, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveOnuWirelessTxPowerFlags(depth int, m *models.OnuWireless, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
