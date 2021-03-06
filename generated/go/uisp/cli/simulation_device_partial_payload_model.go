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

// Schema cli for SimulationDevicePartialPayload

// register flags to command
func registerModelSimulationDevicePartialPayloadFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSimulationDevicePartialPayloadAltitude(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSimulationDevicePartialPayloadAntenna(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSimulationDevicePartialPayloadApType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSimulationDevicePartialPayloadChannelWidth(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSimulationDevicePartialPayloadColor(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSimulationDevicePartialPayloadCoverageCpeHeight(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSimulationDevicePartialPayloadDeviceID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSimulationDevicePartialPayloadDeviceName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSimulationDevicePartialPayloadEirp(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSimulationDevicePartialPayloadFrequency(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSimulationDevicePartialPayloadHeading(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSimulationDevicePartialPayloadHeight(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSimulationDevicePartialPayloadLatitude(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSimulationDevicePartialPayloadLongitude(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSimulationDevicePartialPayloadName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSimulationDevicePartialPayloadRadius(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSimulationDevicePartialPayloadRole(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSimulationDevicePartialPayloadAltitude(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	altitudeDescription := ``

	var altitudeFlagName string
	if cmdPrefix == "" {
		altitudeFlagName = "altitude"
	} else {
		altitudeFlagName = fmt.Sprintf("%v.altitude", cmdPrefix)
	}

	var altitudeFlagDefault float64

	_ = cmd.PersistentFlags().Float64(altitudeFlagName, altitudeFlagDefault, altitudeDescription)

	return nil
}

func registerSimulationDevicePartialPayloadAntenna(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var antennaFlagName string
	if cmdPrefix == "" {
		antennaFlagName = "antenna"
	} else {
		antennaFlagName = fmt.Sprintf("%v.antenna", cmdPrefix)
	}

	if err := registerModelAntenna3Flags(depth+1, antennaFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerSimulationDevicePartialPayloadApType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	apTypeDescription := `Enum: ["ptp","ptmp"]. `

	var apTypeFlagName string
	if cmdPrefix == "" {
		apTypeFlagName = "apType"
	} else {
		apTypeFlagName = fmt.Sprintf("%v.apType", cmdPrefix)
	}

	var apTypeFlagDefault string

	_ = cmd.PersistentFlags().String(apTypeFlagName, apTypeFlagDefault, apTypeDescription)

	if err := cmd.RegisterFlagCompletionFunc(apTypeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["ptp","ptmp"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerSimulationDevicePartialPayloadChannelWidth(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	channelWidthDescription := `Enum: [0]. `

	var channelWidthFlagName string
	if cmdPrefix == "" {
		channelWidthFlagName = "channelWidth"
	} else {
		channelWidthFlagName = fmt.Sprintf("%v.channelWidth", cmdPrefix)
	}

	var channelWidthFlagDefault float64

	_ = cmd.PersistentFlags().Float64(channelWidthFlagName, channelWidthFlagDefault, channelWidthDescription)

	if err := cmd.RegisterFlagCompletionFunc(channelWidthFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`[0]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerSimulationDevicePartialPayloadColor(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	colorDescription := ``

	var colorFlagName string
	if cmdPrefix == "" {
		colorFlagName = "color"
	} else {
		colorFlagName = fmt.Sprintf("%v.color", cmdPrefix)
	}

	var colorFlagDefault string

	_ = cmd.PersistentFlags().String(colorFlagName, colorFlagDefault, colorDescription)

	return nil
}

func registerSimulationDevicePartialPayloadCoverageCpeHeight(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	coverageCpeHeightDescription := `Enum: [0]. `

	var coverageCpeHeightFlagName string
	if cmdPrefix == "" {
		coverageCpeHeightFlagName = "coverageCpeHeight"
	} else {
		coverageCpeHeightFlagName = fmt.Sprintf("%v.coverageCpeHeight", cmdPrefix)
	}

	var coverageCpeHeightFlagDefault float64

	_ = cmd.PersistentFlags().Float64(coverageCpeHeightFlagName, coverageCpeHeightFlagDefault, coverageCpeHeightDescription)

	if err := cmd.RegisterFlagCompletionFunc(coverageCpeHeightFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`[0]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerSimulationDevicePartialPayloadDeviceID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	deviceIdDescription := ``

	var deviceIdFlagName string
	if cmdPrefix == "" {
		deviceIdFlagName = "deviceId"
	} else {
		deviceIdFlagName = fmt.Sprintf("%v.deviceId", cmdPrefix)
	}

	var deviceIdFlagDefault string

	_ = cmd.PersistentFlags().String(deviceIdFlagName, deviceIdFlagDefault, deviceIdDescription)

	return nil
}

func registerSimulationDevicePartialPayloadDeviceName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	deviceNameDescription := ``

	var deviceNameFlagName string
	if cmdPrefix == "" {
		deviceNameFlagName = "deviceName"
	} else {
		deviceNameFlagName = fmt.Sprintf("%v.deviceName", cmdPrefix)
	}

	var deviceNameFlagDefault string

	_ = cmd.PersistentFlags().String(deviceNameFlagName, deviceNameFlagDefault, deviceNameDescription)

	return nil
}

func registerSimulationDevicePartialPayloadEirp(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	eirpDescription := `Enum: [0]. `

	var eirpFlagName string
	if cmdPrefix == "" {
		eirpFlagName = "eirp"
	} else {
		eirpFlagName = fmt.Sprintf("%v.eirp", cmdPrefix)
	}

	var eirpFlagDefault int64

	_ = cmd.PersistentFlags().Int64(eirpFlagName, eirpFlagDefault, eirpDescription)

	if err := cmd.RegisterFlagCompletionFunc(eirpFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`[0]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerSimulationDevicePartialPayloadFrequency(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	frequencyDescription := `Enum: [0]. `

	var frequencyFlagName string
	if cmdPrefix == "" {
		frequencyFlagName = "frequency"
	} else {
		frequencyFlagName = fmt.Sprintf("%v.frequency", cmdPrefix)
	}

	var frequencyFlagDefault int64

	_ = cmd.PersistentFlags().Int64(frequencyFlagName, frequencyFlagDefault, frequencyDescription)

	if err := cmd.RegisterFlagCompletionFunc(frequencyFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`[0]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerSimulationDevicePartialPayloadHeading(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	headingDescription := ``

	var headingFlagName string
	if cmdPrefix == "" {
		headingFlagName = "heading"
	} else {
		headingFlagName = fmt.Sprintf("%v.heading", cmdPrefix)
	}

	var headingFlagDefault float64

	_ = cmd.PersistentFlags().Float64(headingFlagName, headingFlagDefault, headingDescription)

	return nil
}

func registerSimulationDevicePartialPayloadHeight(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	heightDescription := `Enum: [0]. `

	var heightFlagName string
	if cmdPrefix == "" {
		heightFlagName = "height"
	} else {
		heightFlagName = fmt.Sprintf("%v.height", cmdPrefix)
	}

	var heightFlagDefault float64

	_ = cmd.PersistentFlags().Float64(heightFlagName, heightFlagDefault, heightDescription)

	if err := cmd.RegisterFlagCompletionFunc(heightFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`[0]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerSimulationDevicePartialPayloadLatitude(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	latitudeDescription := ``

	var latitudeFlagName string
	if cmdPrefix == "" {
		latitudeFlagName = "latitude"
	} else {
		latitudeFlagName = fmt.Sprintf("%v.latitude", cmdPrefix)
	}

	var latitudeFlagDefault float64

	_ = cmd.PersistentFlags().Float64(latitudeFlagName, latitudeFlagDefault, latitudeDescription)

	return nil
}

func registerSimulationDevicePartialPayloadLongitude(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	longitudeDescription := ``

	var longitudeFlagName string
	if cmdPrefix == "" {
		longitudeFlagName = "longitude"
	} else {
		longitudeFlagName = fmt.Sprintf("%v.longitude", cmdPrefix)
	}

	var longitudeFlagDefault float64

	_ = cmd.PersistentFlags().Float64(longitudeFlagName, longitudeFlagDefault, longitudeDescription)

	return nil
}

func registerSimulationDevicePartialPayloadName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := ``

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerSimulationDevicePartialPayloadRadius(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	radiusDescription := `Enum: [0]. `

	var radiusFlagName string
	if cmdPrefix == "" {
		radiusFlagName = "radius"
	} else {
		radiusFlagName = fmt.Sprintf("%v.radius", cmdPrefix)
	}

	var radiusFlagDefault float64

	_ = cmd.PersistentFlags().Float64(radiusFlagName, radiusFlagDefault, radiusDescription)

	if err := cmd.RegisterFlagCompletionFunc(radiusFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`[0]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerSimulationDevicePartialPayloadRole(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	roleDescription := `Enum: ["ap","cpe"]. `

	var roleFlagName string
	if cmdPrefix == "" {
		roleFlagName = "role"
	} else {
		roleFlagName = fmt.Sprintf("%v.role", cmdPrefix)
	}

	var roleFlagDefault string

	_ = cmd.PersistentFlags().String(roleFlagName, roleFlagDefault, roleDescription)

	if err := cmd.RegisterFlagCompletionFunc(roleFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["ap","cpe"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSimulationDevicePartialPayloadFlags(depth int, m *models.SimulationDevicePartialPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, altitudeAdded := retrieveSimulationDevicePartialPayloadAltitudeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || altitudeAdded

	err, antennaAdded := retrieveSimulationDevicePartialPayloadAntennaFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || antennaAdded

	err, apTypeAdded := retrieveSimulationDevicePartialPayloadApTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || apTypeAdded

	err, channelWidthAdded := retrieveSimulationDevicePartialPayloadChannelWidthFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || channelWidthAdded

	err, colorAdded := retrieveSimulationDevicePartialPayloadColorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || colorAdded

	err, coverageCpeHeightAdded := retrieveSimulationDevicePartialPayloadCoverageCpeHeightFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || coverageCpeHeightAdded

	err, deviceIdAdded := retrieveSimulationDevicePartialPayloadDeviceIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceIdAdded

	err, deviceNameAdded := retrieveSimulationDevicePartialPayloadDeviceNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceNameAdded

	err, eirpAdded := retrieveSimulationDevicePartialPayloadEirpFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || eirpAdded

	err, frequencyAdded := retrieveSimulationDevicePartialPayloadFrequencyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || frequencyAdded

	err, headingAdded := retrieveSimulationDevicePartialPayloadHeadingFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || headingAdded

	err, heightAdded := retrieveSimulationDevicePartialPayloadHeightFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || heightAdded

	err, latitudeAdded := retrieveSimulationDevicePartialPayloadLatitudeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || latitudeAdded

	err, longitudeAdded := retrieveSimulationDevicePartialPayloadLongitudeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || longitudeAdded

	err, nameAdded := retrieveSimulationDevicePartialPayloadNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, radiusAdded := retrieveSimulationDevicePartialPayloadRadiusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || radiusAdded

	err, roleAdded := retrieveSimulationDevicePartialPayloadRoleFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || roleAdded

	return nil, retAdded
}

func retrieveSimulationDevicePartialPayloadAltitudeFlags(depth int, m *models.SimulationDevicePartialPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	altitudeFlagName := fmt.Sprintf("%v.altitude", cmdPrefix)
	if cmd.Flags().Changed(altitudeFlagName) {

		var altitudeFlagName string
		if cmdPrefix == "" {
			altitudeFlagName = "altitude"
		} else {
			altitudeFlagName = fmt.Sprintf("%v.altitude", cmdPrefix)
		}

		altitudeFlagValue, err := cmd.Flags().GetFloat64(altitudeFlagName)
		if err != nil {
			return err, false
		}
		m.Altitude = &altitudeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSimulationDevicePartialPayloadAntennaFlags(depth int, m *models.SimulationDevicePartialPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	antennaFlagName := fmt.Sprintf("%v.antenna", cmdPrefix)
	if cmd.Flags().Changed(antennaFlagName) {
		// info: complex object antenna Antenna3 is retrieved outside this Changed() block
	}
	antennaFlagValue := m.Antenna
	if swag.IsZero(antennaFlagValue) {
		antennaFlagValue = &models.Antenna3{}
	}

	err, antennaAdded := retrieveModelAntenna3Flags(depth+1, antennaFlagValue, antennaFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || antennaAdded
	if antennaAdded {
		m.Antenna = antennaFlagValue
	}

	return nil, retAdded
}

func retrieveSimulationDevicePartialPayloadApTypeFlags(depth int, m *models.SimulationDevicePartialPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	apTypeFlagName := fmt.Sprintf("%v.apType", cmdPrefix)
	if cmd.Flags().Changed(apTypeFlagName) {

		var apTypeFlagName string
		if cmdPrefix == "" {
			apTypeFlagName = "apType"
		} else {
			apTypeFlagName = fmt.Sprintf("%v.apType", cmdPrefix)
		}

		apTypeFlagValue, err := cmd.Flags().GetString(apTypeFlagName)
		if err != nil {
			return err, false
		}
		m.ApType = apTypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSimulationDevicePartialPayloadChannelWidthFlags(depth int, m *models.SimulationDevicePartialPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.ChannelWidth = channelWidthFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSimulationDevicePartialPayloadColorFlags(depth int, m *models.SimulationDevicePartialPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	colorFlagName := fmt.Sprintf("%v.color", cmdPrefix)
	if cmd.Flags().Changed(colorFlagName) {

		var colorFlagName string
		if cmdPrefix == "" {
			colorFlagName = "color"
		} else {
			colorFlagName = fmt.Sprintf("%v.color", cmdPrefix)
		}

		colorFlagValue, err := cmd.Flags().GetString(colorFlagName)
		if err != nil {
			return err, false
		}
		m.Color = colorFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSimulationDevicePartialPayloadCoverageCpeHeightFlags(depth int, m *models.SimulationDevicePartialPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	coverageCpeHeightFlagName := fmt.Sprintf("%v.coverageCpeHeight", cmdPrefix)
	if cmd.Flags().Changed(coverageCpeHeightFlagName) {

		var coverageCpeHeightFlagName string
		if cmdPrefix == "" {
			coverageCpeHeightFlagName = "coverageCpeHeight"
		} else {
			coverageCpeHeightFlagName = fmt.Sprintf("%v.coverageCpeHeight", cmdPrefix)
		}

		coverageCpeHeightFlagValue, err := cmd.Flags().GetFloat64(coverageCpeHeightFlagName)
		if err != nil {
			return err, false
		}
		m.CoverageCpeHeight = coverageCpeHeightFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSimulationDevicePartialPayloadDeviceIDFlags(depth int, m *models.SimulationDevicePartialPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deviceIdFlagName := fmt.Sprintf("%v.deviceId", cmdPrefix)
	if cmd.Flags().Changed(deviceIdFlagName) {

		var deviceIdFlagName string
		if cmdPrefix == "" {
			deviceIdFlagName = "deviceId"
		} else {
			deviceIdFlagName = fmt.Sprintf("%v.deviceId", cmdPrefix)
		}

		deviceIdFlagValue, err := cmd.Flags().GetString(deviceIdFlagName)
		if err != nil {
			return err, false
		}
		m.DeviceID = deviceIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSimulationDevicePartialPayloadDeviceNameFlags(depth int, m *models.SimulationDevicePartialPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deviceNameFlagName := fmt.Sprintf("%v.deviceName", cmdPrefix)
	if cmd.Flags().Changed(deviceNameFlagName) {

		var deviceNameFlagName string
		if cmdPrefix == "" {
			deviceNameFlagName = "deviceName"
		} else {
			deviceNameFlagName = fmt.Sprintf("%v.deviceName", cmdPrefix)
		}

		deviceNameFlagValue, err := cmd.Flags().GetString(deviceNameFlagName)
		if err != nil {
			return err, false
		}
		m.DeviceName = deviceNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSimulationDevicePartialPayloadEirpFlags(depth int, m *models.SimulationDevicePartialPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	eirpFlagName := fmt.Sprintf("%v.eirp", cmdPrefix)
	if cmd.Flags().Changed(eirpFlagName) {

		var eirpFlagName string
		if cmdPrefix == "" {
			eirpFlagName = "eirp"
		} else {
			eirpFlagName = fmt.Sprintf("%v.eirp", cmdPrefix)
		}

		eirpFlagValue, err := cmd.Flags().GetInt64(eirpFlagName)
		if err != nil {
			return err, false
		}
		m.Eirp = eirpFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSimulationDevicePartialPayloadFrequencyFlags(depth int, m *models.SimulationDevicePartialPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

		frequencyFlagValue, err := cmd.Flags().GetInt64(frequencyFlagName)
		if err != nil {
			return err, false
		}
		m.Frequency = frequencyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSimulationDevicePartialPayloadHeadingFlags(depth int, m *models.SimulationDevicePartialPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	headingFlagName := fmt.Sprintf("%v.heading", cmdPrefix)
	if cmd.Flags().Changed(headingFlagName) {

		var headingFlagName string
		if cmdPrefix == "" {
			headingFlagName = "heading"
		} else {
			headingFlagName = fmt.Sprintf("%v.heading", cmdPrefix)
		}

		headingFlagValue, err := cmd.Flags().GetFloat64(headingFlagName)
		if err != nil {
			return err, false
		}
		m.Heading = headingFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSimulationDevicePartialPayloadHeightFlags(depth int, m *models.SimulationDevicePartialPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	heightFlagName := fmt.Sprintf("%v.height", cmdPrefix)
	if cmd.Flags().Changed(heightFlagName) {

		var heightFlagName string
		if cmdPrefix == "" {
			heightFlagName = "height"
		} else {
			heightFlagName = fmt.Sprintf("%v.height", cmdPrefix)
		}

		heightFlagValue, err := cmd.Flags().GetFloat64(heightFlagName)
		if err != nil {
			return err, false
		}
		m.Height = heightFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSimulationDevicePartialPayloadLatitudeFlags(depth int, m *models.SimulationDevicePartialPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	latitudeFlagName := fmt.Sprintf("%v.latitude", cmdPrefix)
	if cmd.Flags().Changed(latitudeFlagName) {

		var latitudeFlagName string
		if cmdPrefix == "" {
			latitudeFlagName = "latitude"
		} else {
			latitudeFlagName = fmt.Sprintf("%v.latitude", cmdPrefix)
		}

		latitudeFlagValue, err := cmd.Flags().GetFloat64(latitudeFlagName)
		if err != nil {
			return err, false
		}
		m.Latitude = &latitudeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSimulationDevicePartialPayloadLongitudeFlags(depth int, m *models.SimulationDevicePartialPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	longitudeFlagName := fmt.Sprintf("%v.longitude", cmdPrefix)
	if cmd.Flags().Changed(longitudeFlagName) {

		var longitudeFlagName string
		if cmdPrefix == "" {
			longitudeFlagName = "longitude"
		} else {
			longitudeFlagName = fmt.Sprintf("%v.longitude", cmdPrefix)
		}

		longitudeFlagValue, err := cmd.Flags().GetFloat64(longitudeFlagName)
		if err != nil {
			return err, false
		}
		m.Longitude = &longitudeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSimulationDevicePartialPayloadNameFlags(depth int, m *models.SimulationDevicePartialPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSimulationDevicePartialPayloadRadiusFlags(depth int, m *models.SimulationDevicePartialPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	radiusFlagName := fmt.Sprintf("%v.radius", cmdPrefix)
	if cmd.Flags().Changed(radiusFlagName) {

		var radiusFlagName string
		if cmdPrefix == "" {
			radiusFlagName = "radius"
		} else {
			radiusFlagName = fmt.Sprintf("%v.radius", cmdPrefix)
		}

		radiusFlagValue, err := cmd.Flags().GetFloat64(radiusFlagName)
		if err != nil {
			return err, false
		}
		m.Radius = radiusFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSimulationDevicePartialPayloadRoleFlags(depth int, m *models.SimulationDevicePartialPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	roleFlagName := fmt.Sprintf("%v.role", cmdPrefix)
	if cmd.Flags().Changed(roleFlagName) {

		var roleFlagName string
		if cmdPrefix == "" {
			roleFlagName = "role"
		} else {
			roleFlagName = fmt.Sprintf("%v.role", cmdPrefix)
		}

		roleFlagValue, err := cmd.Flags().GetString(roleFlagName)
		if err != nil {
			return err, false
		}
		m.Role = roleFlagValue

		retAdded = true
	}

	return nil, retAdded
}
