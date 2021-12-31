// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
 "github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for DevicesAPSProfile

// register flags to command
func registerModelDevicesAPSProfileFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDevicesAPSProfileAirfiber(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDevicesAPSProfileAirmax(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDevicesAPSProfileAuthentication(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDevicesAPSProfileBackupRadio(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDevicesAPSProfileChannelWidth(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDevicesAPSProfileCountryCode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDevicesAPSProfileDevice(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDevicesAPSProfileFrameLength(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDevicesAPSProfileFrequency(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDevicesAPSProfileKey(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDevicesAPSProfileLocation(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDevicesAPSProfileMac(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDevicesAPSProfileSecurity(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDevicesAPSProfileSsid(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDevicesAPSProfileWirelessMode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDevicesAPSProfileAirfiber(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var airfiberFlagName string
	if cmdPrefix == "" {
		airfiberFlagName = "airfiber"
	} else {
		airfiberFlagName = fmt.Sprintf("%v.airfiber", cmdPrefix)
	}

	if err := registerModelAirfiberFlags(depth+1, airfiberFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDevicesAPSProfileAirmax(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var airmaxFlagName string
	if cmdPrefix == "" {
		airmaxFlagName = "airmax"
	} else {
		airmaxFlagName = fmt.Sprintf("%v.airmax", cmdPrefix)
	}

	if err := registerModelAirmaxFlags(depth+1, airmaxFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDevicesAPSProfileAuthentication(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerDevicesAPSProfileBackupRadio(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var backupRadioFlagName string
	if cmdPrefix == "" {
		backupRadioFlagName = "backupRadio"
	} else {
		backupRadioFlagName = fmt.Sprintf("%v.backupRadio", cmdPrefix)
	}

	if err := registerModelBackupRadioFlags(depth+1, backupRadioFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDevicesAPSProfileChannelWidth(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

	var channelWidthFlagDefault int64

	_ = cmd.PersistentFlags().Int64(channelWidthFlagName, channelWidthFlagDefault, channelWidthDescription)

	return nil
}

func registerDevicesAPSProfileCountryCode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	countryCodeDescription := `Required. `

	var countryCodeFlagName string
	if cmdPrefix == "" {
		countryCodeFlagName = "countryCode"
	} else {
		countryCodeFlagName = fmt.Sprintf("%v.countryCode", cmdPrefix)
	}

	var countryCodeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(countryCodeFlagName, countryCodeFlagDefault, countryCodeDescription)

	return nil
}

func registerDevicesAPSProfileDevice(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var deviceFlagName string
	if cmdPrefix == "" {
		deviceFlagName = "device"
	} else {
		deviceFlagName = fmt.Sprintf("%v.device", cmdPrefix)
	}

	if err := registerModelDeviceFlags(depth+1, deviceFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDevicesAPSProfileFrameLength(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	frameLengthDescription := ``

	var frameLengthFlagName string
	if cmdPrefix == "" {
		frameLengthFlagName = "frameLength"
	} else {
		frameLengthFlagName = fmt.Sprintf("%v.frameLength", cmdPrefix)
	}

	var frameLengthFlagDefault float64

	_ = cmd.PersistentFlags().Float64(frameLengthFlagName, frameLengthFlagDefault, frameLengthDescription)

	return nil
}

func registerDevicesAPSProfileFrequency(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	frequencyDescription := `Required. `

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

func registerDevicesAPSProfileKey(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	keyDescription := `Required. Pre shared key`

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

func registerDevicesAPSProfileLocation(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var locationFlagName string
	if cmdPrefix == "" {
		locationFlagName = "location"
	} else {
		locationFlagName = fmt.Sprintf("%v.location", cmdPrefix)
	}

	if err := registerModelLocation2Flags(depth+1, locationFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDevicesAPSProfileMac(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	macDescription := `Required. MAC address`

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

func registerDevicesAPSProfileSecurity(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	securityDescription := `Enum: ["wep","wpa","wpa-psk","wpa2","enabled","none"]. Required. `

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
			if err := json.Unmarshal([]byte(`["wep","wpa","wpa-psk","wpa2","enabled","none"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerDevicesAPSProfileSsid(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ssidDescription := `Required. SSID`

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

func registerDevicesAPSProfileWirelessMode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	wirelessModeDescription := `Enum: ["ap","ap-ptp","ap-ptmp","ap-ptmp-airmax","ap-ptmp-airmax-mixed","ap-ptmp-airmax-ac","sta","sta-ptp","sta-ptmp","aprepeater","repeater","mesh"]. `

	var wirelessModeFlagName string
	if cmdPrefix == "" {
		wirelessModeFlagName = "wirelessMode"
	} else {
		wirelessModeFlagName = fmt.Sprintf("%v.wirelessMode", cmdPrefix)
	}

	var wirelessModeFlagDefault string

	_ = cmd.PersistentFlags().String(wirelessModeFlagName, wirelessModeFlagDefault, wirelessModeDescription)

	if err := cmd.RegisterFlagCompletionFunc(wirelessModeFlagName,
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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDevicesAPSProfileFlags(depth int, m *models.DevicesAPSProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, airfiberAdded := retrieveDevicesAPSProfileAirfiberFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || airfiberAdded

	err, airmaxAdded := retrieveDevicesAPSProfileAirmaxFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || airmaxAdded

	err, authenticationAdded := retrieveDevicesAPSProfileAuthenticationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || authenticationAdded

	err, backupRadioAdded := retrieveDevicesAPSProfileBackupRadioFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || backupRadioAdded

	err, channelWidthAdded := retrieveDevicesAPSProfileChannelWidthFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || channelWidthAdded

	err, countryCodeAdded := retrieveDevicesAPSProfileCountryCodeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || countryCodeAdded

	err, deviceAdded := retrieveDevicesAPSProfileDeviceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceAdded

	err, frameLengthAdded := retrieveDevicesAPSProfileFrameLengthFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || frameLengthAdded

	err, frequencyAdded := retrieveDevicesAPSProfileFrequencyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || frequencyAdded

	err, keyAdded := retrieveDevicesAPSProfileKeyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || keyAdded

	err, locationAdded := retrieveDevicesAPSProfileLocationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || locationAdded

	err, macAdded := retrieveDevicesAPSProfileMacFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || macAdded

	err, securityAdded := retrieveDevicesAPSProfileSecurityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || securityAdded

	err, ssidAdded := retrieveDevicesAPSProfileSsidFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ssidAdded

	err, wirelessModeAdded := retrieveDevicesAPSProfileWirelessModeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || wirelessModeAdded

	return nil, retAdded
}

func retrieveDevicesAPSProfileAirfiberFlags(depth int, m *models.DevicesAPSProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	airfiberFlagName := fmt.Sprintf("%v.airfiber", cmdPrefix)
	if cmd.Flags().Changed(airfiberFlagName) {
		// info: complex object airfiber Airfiber is retrieved outside this Changed() block
	}
	airfiberFlagValue := m.Airfiber
	if swag.IsZero(airfiberFlagValue) {
		airfiberFlagValue = &models.Airfiber{}
	}

	err, airfiberAdded := retrieveModelAirfiberFlags(depth+1, airfiberFlagValue, airfiberFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || airfiberAdded
	if airfiberAdded {
		m.Airfiber = airfiberFlagValue
	}

	return nil, retAdded
}

func retrieveDevicesAPSProfileAirmaxFlags(depth int, m *models.DevicesAPSProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	airmaxFlagName := fmt.Sprintf("%v.airmax", cmdPrefix)
	if cmd.Flags().Changed(airmaxFlagName) {
		// info: complex object airmax Airmax is retrieved outside this Changed() block
	}
	airmaxFlagValue := m.Airmax
	if swag.IsZero(airmaxFlagValue) {
		airmaxFlagValue = &models.Airmax{}
	}

	err, airmaxAdded := retrieveModelAirmaxFlags(depth+1, airmaxFlagValue, airmaxFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || airmaxAdded
	if airmaxAdded {
		m.Airmax = airmaxFlagValue
	}

	return nil, retAdded
}

func retrieveDevicesAPSProfileAuthenticationFlags(depth int, m *models.DevicesAPSProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDevicesAPSProfileBackupRadioFlags(depth int, m *models.DevicesAPSProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	backupRadioFlagName := fmt.Sprintf("%v.backupRadio", cmdPrefix)
	if cmd.Flags().Changed(backupRadioFlagName) {
		// info: complex object backupRadio BackupRadio is retrieved outside this Changed() block
	}
	backupRadioFlagValue := m.BackupRadio
	if swag.IsZero(backupRadioFlagValue) {
		backupRadioFlagValue = &models.BackupRadio{}
	}

	err, backupRadioAdded := retrieveModelBackupRadioFlags(depth+1, backupRadioFlagValue, backupRadioFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || backupRadioAdded
	if backupRadioAdded {
		m.BackupRadio = backupRadioFlagValue
	}

	return nil, retAdded
}

func retrieveDevicesAPSProfileChannelWidthFlags(depth int, m *models.DevicesAPSProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

		channelWidthFlagValue, err := cmd.Flags().GetInt64(channelWidthFlagName)
		if err != nil {
			return err, false
		}
		m.ChannelWidth = &channelWidthFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDevicesAPSProfileCountryCodeFlags(depth int, m *models.DevicesAPSProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	countryCodeFlagName := fmt.Sprintf("%v.countryCode", cmdPrefix)
	if cmd.Flags().Changed(countryCodeFlagName) {

		var countryCodeFlagName string
		if cmdPrefix == "" {
			countryCodeFlagName = "countryCode"
		} else {
			countryCodeFlagName = fmt.Sprintf("%v.countryCode", cmdPrefix)
		}

		countryCodeFlagValue, err := cmd.Flags().GetInt64(countryCodeFlagName)
		if err != nil {
			return err, false
		}
		m.CountryCode = &countryCodeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDevicesAPSProfileDeviceFlags(depth int, m *models.DevicesAPSProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deviceFlagName := fmt.Sprintf("%v.device", cmdPrefix)
	if cmd.Flags().Changed(deviceFlagName) {
		// info: complex object device Device is retrieved outside this Changed() block
	}
	deviceFlagValue := m.Device
	if swag.IsZero(deviceFlagValue) {
		deviceFlagValue = &models.Device{}
	}

	err, deviceAdded := retrieveModelDeviceFlags(depth+1, deviceFlagValue, deviceFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceAdded
	if deviceAdded {
		m.Device = deviceFlagValue
	}

	return nil, retAdded
}

func retrieveDevicesAPSProfileFrameLengthFlags(depth int, m *models.DevicesAPSProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	frameLengthFlagName := fmt.Sprintf("%v.frameLength", cmdPrefix)
	if cmd.Flags().Changed(frameLengthFlagName) {

		var frameLengthFlagName string
		if cmdPrefix == "" {
			frameLengthFlagName = "frameLength"
		} else {
			frameLengthFlagName = fmt.Sprintf("%v.frameLength", cmdPrefix)
		}

		frameLengthFlagValue, err := cmd.Flags().GetFloat64(frameLengthFlagName)
		if err != nil {
			return err, false
		}
		m.FrameLength = &frameLengthFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDevicesAPSProfileFrequencyFlags(depth int, m *models.DevicesAPSProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Frequency = &frequencyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDevicesAPSProfileKeyFlags(depth int, m *models.DevicesAPSProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDevicesAPSProfileLocationFlags(depth int, m *models.DevicesAPSProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	locationFlagName := fmt.Sprintf("%v.location", cmdPrefix)
	if cmd.Flags().Changed(locationFlagName) {
		// info: complex object location Location2 is retrieved outside this Changed() block
	}
	locationFlagValue := m.Location
	if swag.IsZero(locationFlagValue) {
		locationFlagValue = &models.Location2{}
	}

	err, locationAdded := retrieveModelLocation2Flags(depth+1, locationFlagValue, locationFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || locationAdded
	if locationAdded {
		m.Location = locationFlagValue
	}

	return nil, retAdded
}

func retrieveDevicesAPSProfileMacFlags(depth int, m *models.DevicesAPSProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Mac = &macFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDevicesAPSProfileSecurityFlags(depth int, m *models.DevicesAPSProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDevicesAPSProfileSsidFlags(depth int, m *models.DevicesAPSProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDevicesAPSProfileWirelessModeFlags(depth int, m *models.DevicesAPSProfile, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	wirelessModeFlagName := fmt.Sprintf("%v.wirelessMode", cmdPrefix)
	if cmd.Flags().Changed(wirelessModeFlagName) {

		var wirelessModeFlagName string
		if cmdPrefix == "" {
			wirelessModeFlagName = "wirelessMode"
		} else {
			wirelessModeFlagName = fmt.Sprintf("%v.wirelessMode", cmdPrefix)
		}

		wirelessModeFlagValue, err := cmd.Flags().GetString(wirelessModeFlagName)
		if err != nil {
			return err, false
		}
		m.WirelessMode = wirelessModeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
