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

// Schema cli for DeviceAttributes

// register flags to command
func registerModelDeviceAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDeviceAttributesApDevice(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceAttributesCountry(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceAttributesDyingGasp(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceAttributesGatewayID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceAttributesHasGatewayInterfaceAvailable(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceAttributesIsGatewaySupported(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceAttributesSeries(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceAttributesSsid(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceAttributesApDevice(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var apDeviceFlagName string
	if cmdPrefix == "" {
		apDeviceFlagName = "apDevice"
	} else {
		apDeviceFlagName = fmt.Sprintf("%v.apDevice", cmdPrefix)
	}

	if err := registerModelApDevice2Flags(depth+1, apDeviceFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceAttributesCountry(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	countryDescription := `Enum: ["Unknown","Licensed","Afghanistan","Åland Islands","Albania","Algeria","American Samoa","Andorra","Angola","Anguilla","Antarctica","Antigua and Barbuda","Argentina","Armenia","Aruba","Australia","Austria","Azerbaijan","Bahamas","Bahrain","Bangladesh","Barbados","Belarus","Belgium","Belize","Benin","Bermuda","Bhutan","Bolivia","Bonaire, Sint Eustatius and Saba","Bosnia and Herzegovina","Botswana","Bouvet Island","Brazil","British Indian Ocean Territory","Brunei Darussalam","Bulgaria","Burkina Faso","Burundi","Cabo Verde","Cambodia","Cameroon","Canada","Cayman Islands","Central Africa Republic","Chad  ","Chile","China","Christmas Island","Cocos Islands","Colombia","Comoros","Cook Islands","Costa Rica","Côte d'Ivoire","Croatia","Curaçao","Cyprus","Czech Republic","Democratic Republic of the Congo","Denmark","Djibouti","Dominica","Dominican Republic","Ecuador","Egypt","El Salvador","Equatorial Guinea","Eritrea","Estonia","Eswatini","Ethiopia","Falkland Islands","Faroe Islands","Fiji","Finland","France","French Guiana","French Polynesia","French Southern Territories","Gabon","Gambia","Georgia","Germany","Ghana","Gibraltar","Greece","Greenland","Grenada","Guadeloupe","Guam","Guatemala","Guernsey","Guinea","Guinea-Bissau","Guyana","Haiti","Heard Island and McDonald Islands","Honduras","Hong Kong","Hungary","Iceland","India","Indonesia","Iraq","Ireland","Isle of Man","Israel","Italy","Jamaica","Japan","Jersey","Jordan","Kazakhstan","Kenya","Kiribati","Kuwait","Kyrgyzstan","Lao People's Democratic Republic","Latvia","Lebanon","Lesotho","Liberia","Libya","Liechtenstein","Lithuania","Luxembourg","Macau Sar","Macedonia, Fyro","Madagascar","Malawi","Malaysia","Maldives","Mali","Malta","Marshall Islands","Martinique","Mauritania","Mauritius","Mayotte","Mexico","Micronesia","Moldova","Monaco","Mongolia","Montenegro","Montserrat","Morocco","Mozambique","Myanmar","Namibia","Nauru","Nepal","Netherlands","New Caledonia","New Zealand","Nicaragua","Niger","Nigeria","Niue","Norfolk Island","Northern Mariana Islands","Norway","Oman","Pakistan","Palau","Panama","Papua New Guinea","Paraguay","Peru","Philippines","Pitcairn","Poland","Portugal","Puerto Rico","Qatar","Republic of Korea","Republic of Serbia","Republic of Seychelles","Republic of the Congo","Reunion","Romania","Russia","Rwanda","Saint Barthélemy","Saint Helena, Ascension and Tristan da Cunh","Saint Kitts and Nevis","Saint Lucia","Saint Martin","Saint Pierre Miquelon","Saint Vincent Grenadiens","Samoa","San Marino","Sao Tome and Principe","Saudi Arabia","Senegal","Sierra Leone","Singapore","Sint Maarten","Slovakia","Slovenia","Solomon Islands","Somalia","South Africa","South Georgia and the South Sandwich Islands","South Sudan","Spain","Sri Lanka","State of Palestine","Suriname","Svalbard and Jan Mayen","Sweden","Switzerland","Taiwan","Tajikistan","Tanzania","Thailand","Timor-Leste","Togo","Tokelau","Tonga","Trinidad and Tobago","Tunisia","Turkey","Turkmenistan","Turks Caicos","Tuvalu","Uganda","Ukraine","United Arab Emirates","United Kingdom","United States","Uruguay","Uzbekistan","Vanuatu","Vatican City State","Venezuela","Vietnam","Virgin Islands (British)","Virgin Islands (U.S.)","Wallis Futuna","Western Sahara","Yemen","Zambia","Zimbabwe"]. `

	var countryFlagName string
	if cmdPrefix == "" {
		countryFlagName = "country"
	} else {
		countryFlagName = fmt.Sprintf("%v.country", cmdPrefix)
	}

	var countryFlagDefault string

	_ = cmd.PersistentFlags().String(countryFlagName, countryFlagDefault, countryDescription)

	if err := cmd.RegisterFlagCompletionFunc(countryFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["Unknown","Licensed","Afghanistan","Åland Islands","Albania","Algeria","American Samoa","Andorra","Angola","Anguilla","Antarctica","Antigua and Barbuda","Argentina","Armenia","Aruba","Australia","Austria","Azerbaijan","Bahamas","Bahrain","Bangladesh","Barbados","Belarus","Belgium","Belize","Benin","Bermuda","Bhutan","Bolivia","Bonaire, Sint Eustatius and Saba","Bosnia and Herzegovina","Botswana","Bouvet Island","Brazil","British Indian Ocean Territory","Brunei Darussalam","Bulgaria","Burkina Faso","Burundi","Cabo Verde","Cambodia","Cameroon","Canada","Cayman Islands","Central Africa Republic","Chad  ","Chile","China","Christmas Island","Cocos Islands","Colombia","Comoros","Cook Islands","Costa Rica","Côte d'Ivoire","Croatia","Curaçao","Cyprus","Czech Republic","Democratic Republic of the Congo","Denmark","Djibouti","Dominica","Dominican Republic","Ecuador","Egypt","El Salvador","Equatorial Guinea","Eritrea","Estonia","Eswatini","Ethiopia","Falkland Islands","Faroe Islands","Fiji","Finland","France","French Guiana","French Polynesia","French Southern Territories","Gabon","Gambia","Georgia","Germany","Ghana","Gibraltar","Greece","Greenland","Grenada","Guadeloupe","Guam","Guatemala","Guernsey","Guinea","Guinea-Bissau","Guyana","Haiti","Heard Island and McDonald Islands","Honduras","Hong Kong","Hungary","Iceland","India","Indonesia","Iraq","Ireland","Isle of Man","Israel","Italy","Jamaica","Japan","Jersey","Jordan","Kazakhstan","Kenya","Kiribati","Kuwait","Kyrgyzstan","Lao People's Democratic Republic","Latvia","Lebanon","Lesotho","Liberia","Libya","Liechtenstein","Lithuania","Luxembourg","Macau Sar","Macedonia, Fyro","Madagascar","Malawi","Malaysia","Maldives","Mali","Malta","Marshall Islands","Martinique","Mauritania","Mauritius","Mayotte","Mexico","Micronesia","Moldova","Monaco","Mongolia","Montenegro","Montserrat","Morocco","Mozambique","Myanmar","Namibia","Nauru","Nepal","Netherlands","New Caledonia","New Zealand","Nicaragua","Niger","Nigeria","Niue","Norfolk Island","Northern Mariana Islands","Norway","Oman","Pakistan","Palau","Panama","Papua New Guinea","Paraguay","Peru","Philippines","Pitcairn","Poland","Portugal","Puerto Rico","Qatar","Republic of Korea","Republic of Serbia","Republic of Seychelles","Republic of the Congo","Reunion","Romania","Russia","Rwanda","Saint Barthélemy","Saint Helena, Ascension and Tristan da Cunh","Saint Kitts and Nevis","Saint Lucia","Saint Martin","Saint Pierre Miquelon","Saint Vincent Grenadiens","Samoa","San Marino","Sao Tome and Principe","Saudi Arabia","Senegal","Sierra Leone","Singapore","Sint Maarten","Slovakia","Slovenia","Solomon Islands","Somalia","South Africa","South Georgia and the South Sandwich Islands","South Sudan","Spain","Sri Lanka","State of Palestine","Suriname","Svalbard and Jan Mayen","Sweden","Switzerland","Taiwan","Tajikistan","Tanzania","Thailand","Timor-Leste","Togo","Tokelau","Tonga","Trinidad and Tobago","Tunisia","Turkey","Turkmenistan","Turks Caicos","Tuvalu","Uganda","Ukraine","United Arab Emirates","United Kingdom","United States","Uruguay","Uzbekistan","Vanuatu","Vatican City State","Venezuela","Vietnam","Virgin Islands (British)","Virgin Islands (U.S.)","Wallis Futuna","Western Sahara","Yemen","Zambia","Zimbabwe"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerDeviceAttributesDyingGasp(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	dyingGaspDescription := `Set to true if ONU is disconnected due to power outage`

	var dyingGaspFlagName string
	if cmdPrefix == "" {
		dyingGaspFlagName = "dyingGasp"
	} else {
		dyingGaspFlagName = fmt.Sprintf("%v.dyingGasp", cmdPrefix)
	}

	var dyingGaspFlagDefault bool

	_ = cmd.PersistentFlags().Bool(dyingGaspFlagName, dyingGaspFlagDefault, dyingGaspDescription)

	return nil
}

func registerDeviceAttributesGatewayID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	gatewayIdDescription := ``

	var gatewayIdFlagName string
	if cmdPrefix == "" {
		gatewayIdFlagName = "gatewayId"
	} else {
		gatewayIdFlagName = fmt.Sprintf("%v.gatewayId", cmdPrefix)
	}

	var gatewayIdFlagDefault string

	_ = cmd.PersistentFlags().String(gatewayIdFlagName, gatewayIdFlagDefault, gatewayIdDescription)

	return nil
}

func registerDeviceAttributesHasGatewayInterfaceAvailable(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	hasGatewayInterfaceAvailableDescription := ``

	var hasGatewayInterfaceAvailableFlagName string
	if cmdPrefix == "" {
		hasGatewayInterfaceAvailableFlagName = "hasGatewayInterfaceAvailable"
	} else {
		hasGatewayInterfaceAvailableFlagName = fmt.Sprintf("%v.hasGatewayInterfaceAvailable", cmdPrefix)
	}

	var hasGatewayInterfaceAvailableFlagDefault bool

	_ = cmd.PersistentFlags().Bool(hasGatewayInterfaceAvailableFlagName, hasGatewayInterfaceAvailableFlagDefault, hasGatewayInterfaceAvailableDescription)

	return nil
}

func registerDeviceAttributesIsGatewaySupported(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	isGatewaySupportedDescription := ``

	var isGatewaySupportedFlagName string
	if cmdPrefix == "" {
		isGatewaySupportedFlagName = "isGatewaySupported"
	} else {
		isGatewaySupportedFlagName = fmt.Sprintf("%v.isGatewaySupported", cmdPrefix)
	}

	var isGatewaySupportedFlagDefault bool

	_ = cmd.PersistentFlags().Bool(isGatewaySupportedFlagName, isGatewaySupportedFlagDefault, isGatewaySupportedDescription)

	return nil
}

func registerDeviceAttributesSeries(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	seriesDescription := ``

	var seriesFlagName string
	if cmdPrefix == "" {
		seriesFlagName = "series"
	} else {
		seriesFlagName = fmt.Sprintf("%v.series", cmdPrefix)
	}

	var seriesFlagDefault string

	_ = cmd.PersistentFlags().String(seriesFlagName, seriesFlagDefault, seriesDescription)

	return nil
}

func registerDeviceAttributesSsid(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ssidDescription := `SSID`

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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDeviceAttributesFlags(depth int, m *models.DeviceAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, apDeviceAdded := retrieveDeviceAttributesApDeviceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || apDeviceAdded

	err, countryAdded := retrieveDeviceAttributesCountryFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || countryAdded

	err, dyingGaspAdded := retrieveDeviceAttributesDyingGaspFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dyingGaspAdded

	err, gatewayIdAdded := retrieveDeviceAttributesGatewayIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || gatewayIdAdded

	err, hasGatewayInterfaceAvailableAdded := retrieveDeviceAttributesHasGatewayInterfaceAvailableFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || hasGatewayInterfaceAvailableAdded

	err, isGatewaySupportedAdded := retrieveDeviceAttributesIsGatewaySupportedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || isGatewaySupportedAdded

	err, seriesAdded := retrieveDeviceAttributesSeriesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || seriesAdded

	err, ssidAdded := retrieveDeviceAttributesSsidFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ssidAdded

	return nil, retAdded
}

func retrieveDeviceAttributesApDeviceFlags(depth int, m *models.DeviceAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	apDeviceFlagName := fmt.Sprintf("%v.apDevice", cmdPrefix)
	if cmd.Flags().Changed(apDeviceFlagName) {
		// info: complex object apDevice ApDevice2 is retrieved outside this Changed() block
	}
	apDeviceFlagValue := m.ApDevice
	if swag.IsZero(apDeviceFlagValue) {
		apDeviceFlagValue = &models.ApDevice2{}
	}

	err, apDeviceAdded := retrieveModelApDevice2Flags(depth+1, apDeviceFlagValue, apDeviceFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || apDeviceAdded
	if apDeviceAdded {
		m.ApDevice = apDeviceFlagValue
	}

	return nil, retAdded
}

func retrieveDeviceAttributesCountryFlags(depth int, m *models.DeviceAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDeviceAttributesDyingGaspFlags(depth int, m *models.DeviceAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dyingGaspFlagName := fmt.Sprintf("%v.dyingGasp", cmdPrefix)
	if cmd.Flags().Changed(dyingGaspFlagName) {

		var dyingGaspFlagName string
		if cmdPrefix == "" {
			dyingGaspFlagName = "dyingGasp"
		} else {
			dyingGaspFlagName = fmt.Sprintf("%v.dyingGasp", cmdPrefix)
		}

		dyingGaspFlagValue, err := cmd.Flags().GetBool(dyingGaspFlagName)
		if err != nil {
			return err, false
		}
		m.DyingGasp = dyingGaspFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceAttributesGatewayIDFlags(depth int, m *models.DeviceAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	gatewayIdFlagName := fmt.Sprintf("%v.gatewayId", cmdPrefix)
	if cmd.Flags().Changed(gatewayIdFlagName) {

		var gatewayIdFlagName string
		if cmdPrefix == "" {
			gatewayIdFlagName = "gatewayId"
		} else {
			gatewayIdFlagName = fmt.Sprintf("%v.gatewayId", cmdPrefix)
		}

		gatewayIdFlagValue, err := cmd.Flags().GetString(gatewayIdFlagName)
		if err != nil {
			return err, false
		}
		m.GatewayID = gatewayIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceAttributesHasGatewayInterfaceAvailableFlags(depth int, m *models.DeviceAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	hasGatewayInterfaceAvailableFlagName := fmt.Sprintf("%v.hasGatewayInterfaceAvailable", cmdPrefix)
	if cmd.Flags().Changed(hasGatewayInterfaceAvailableFlagName) {

		var hasGatewayInterfaceAvailableFlagName string
		if cmdPrefix == "" {
			hasGatewayInterfaceAvailableFlagName = "hasGatewayInterfaceAvailable"
		} else {
			hasGatewayInterfaceAvailableFlagName = fmt.Sprintf("%v.hasGatewayInterfaceAvailable", cmdPrefix)
		}

		hasGatewayInterfaceAvailableFlagValue, err := cmd.Flags().GetBool(hasGatewayInterfaceAvailableFlagName)
		if err != nil {
			return err, false
		}
		m.HasGatewayInterfaceAvailable = hasGatewayInterfaceAvailableFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceAttributesIsGatewaySupportedFlags(depth int, m *models.DeviceAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	isGatewaySupportedFlagName := fmt.Sprintf("%v.isGatewaySupported", cmdPrefix)
	if cmd.Flags().Changed(isGatewaySupportedFlagName) {

		var isGatewaySupportedFlagName string
		if cmdPrefix == "" {
			isGatewaySupportedFlagName = "isGatewaySupported"
		} else {
			isGatewaySupportedFlagName = fmt.Sprintf("%v.isGatewaySupported", cmdPrefix)
		}

		isGatewaySupportedFlagValue, err := cmd.Flags().GetBool(isGatewaySupportedFlagName)
		if err != nil {
			return err, false
		}
		m.IsGatewaySupported = isGatewaySupportedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceAttributesSeriesFlags(depth int, m *models.DeviceAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	seriesFlagName := fmt.Sprintf("%v.series", cmdPrefix)
	if cmd.Flags().Changed(seriesFlagName) {

		var seriesFlagName string
		if cmdPrefix == "" {
			seriesFlagName = "series"
		} else {
			seriesFlagName = fmt.Sprintf("%v.series", cmdPrefix)
		}

		seriesFlagValue, err := cmd.Flags().GetString(seriesFlagName)
		if err != nil {
			return err, false
		}
		m.Series = seriesFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceAttributesSsidFlags(depth int, m *models.DeviceAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Ssid = ssidFlagValue

		retAdded = true
	}

	return nil, retAdded
}
