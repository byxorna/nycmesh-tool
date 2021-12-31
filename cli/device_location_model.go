// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for DeviceLocation

// register flags to command
func registerModelDeviceLocationFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDeviceLocationAltitude(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceLocationElevation(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceLocationHeading(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceLocationLatitude(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceLocationLongitude(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceLocationMagneticHeading(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceLocationRoll(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceLocationTilt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceLocationAltitude(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	altitudeDescription := `Required. Meters above the ground.`

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

func registerDeviceLocationElevation(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	elevationDescription := `Required. Meters above the sea level.`

	var elevationFlagName string
	if cmdPrefix == "" {
		elevationFlagName = "elevation"
	} else {
		elevationFlagName = fmt.Sprintf("%v.elevation", cmdPrefix)
	}

	var elevationFlagDefault float64

	_ = cmd.PersistentFlags().Float64(elevationFlagName, elevationFlagDefault, elevationDescription)

	return nil
}

func registerDeviceLocationHeading(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	headingDescription := `
    Value in degrees. The value is based on map, e.g. without magnetic declination.
    Ignored on udpate when magneticHeading is present. North is 0 degrees, East is 90, South is 180, West is 270.
  `

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

func registerDeviceLocationLatitude(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	latitudeDescription := `Required. `

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

func registerDeviceLocationLongitude(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	longitudeDescription := `Required. `

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

func registerDeviceLocationMagneticHeading(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	magneticHeadingDescription := `Required. 
    Value in degrees. Compass heading (magnetic north).
    It’s what you’re changing when you spin the dish around on a mast.
    North is 0 degrees, East is 90, South is 180, West is 270.
  `

	var magneticHeadingFlagName string
	if cmdPrefix == "" {
		magneticHeadingFlagName = "magneticHeading"
	} else {
		magneticHeadingFlagName = fmt.Sprintf("%v.magneticHeading", cmdPrefix)
	}

	var magneticHeadingFlagDefault float64

	_ = cmd.PersistentFlags().Float64(magneticHeadingFlagName, magneticHeadingFlagDefault, magneticHeadingDescription)

	return nil
}

func registerDeviceLocationRoll(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	rollDescription := `Required. 
    Value in degrees. How low one side of the dish is compared to the other.
    0 degrees no rotation, 90 rotation to the left, -90 rotation to the right and -180 upside down.
  `

	var rollFlagName string
	if cmdPrefix == "" {
		rollFlagName = "roll"
	} else {
		rollFlagName = fmt.Sprintf("%v.roll", cmdPrefix)
	}

	var rollFlagDefault float64

	_ = cmd.PersistentFlags().Float64(rollFlagName, rollFlagDefault, rollDescription)

	return nil
}

func registerDeviceLocationTilt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	tiltDescription := `Required. 
    Value in degrees. How high in the sky the dish is pointing.
    0 degrees has the dish pointing straight ahead, 90 has the dish pointing straight up
    and -90 has the dish pointing straight down.
  `

	var tiltFlagName string
	if cmdPrefix == "" {
		tiltFlagName = "tilt"
	} else {
		tiltFlagName = fmt.Sprintf("%v.tilt", cmdPrefix)
	}

	var tiltFlagDefault float64

	_ = cmd.PersistentFlags().Float64(tiltFlagName, tiltFlagDefault, tiltDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDeviceLocationFlags(depth int, m *models.DeviceLocation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, altitudeAdded := retrieveDeviceLocationAltitudeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || altitudeAdded

	err, elevationAdded := retrieveDeviceLocationElevationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || elevationAdded

	err, headingAdded := retrieveDeviceLocationHeadingFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || headingAdded

	err, latitudeAdded := retrieveDeviceLocationLatitudeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || latitudeAdded

	err, longitudeAdded := retrieveDeviceLocationLongitudeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || longitudeAdded

	err, magneticHeadingAdded := retrieveDeviceLocationMagneticHeadingFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || magneticHeadingAdded

	err, rollAdded := retrieveDeviceLocationRollFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || rollAdded

	err, tiltAdded := retrieveDeviceLocationTiltFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tiltAdded

	return nil, retAdded
}

func retrieveDeviceLocationAltitudeFlags(depth int, m *models.DeviceLocation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDeviceLocationElevationFlags(depth int, m *models.DeviceLocation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	elevationFlagName := fmt.Sprintf("%v.elevation", cmdPrefix)
	if cmd.Flags().Changed(elevationFlagName) {

		var elevationFlagName string
		if cmdPrefix == "" {
			elevationFlagName = "elevation"
		} else {
			elevationFlagName = fmt.Sprintf("%v.elevation", cmdPrefix)
		}

		elevationFlagValue, err := cmd.Flags().GetFloat64(elevationFlagName)
		if err != nil {
			return err, false
		}
		m.Elevation = &elevationFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceLocationHeadingFlags(depth int, m *models.DeviceLocation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDeviceLocationLatitudeFlags(depth int, m *models.DeviceLocation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDeviceLocationLongitudeFlags(depth int, m *models.DeviceLocation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDeviceLocationMagneticHeadingFlags(depth int, m *models.DeviceLocation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	magneticHeadingFlagName := fmt.Sprintf("%v.magneticHeading", cmdPrefix)
	if cmd.Flags().Changed(magneticHeadingFlagName) {

		var magneticHeadingFlagName string
		if cmdPrefix == "" {
			magneticHeadingFlagName = "magneticHeading"
		} else {
			magneticHeadingFlagName = fmt.Sprintf("%v.magneticHeading", cmdPrefix)
		}

		magneticHeadingFlagValue, err := cmd.Flags().GetFloat64(magneticHeadingFlagName)
		if err != nil {
			return err, false
		}
		m.MagneticHeading = &magneticHeadingFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceLocationRollFlags(depth int, m *models.DeviceLocation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	rollFlagName := fmt.Sprintf("%v.roll", cmdPrefix)
	if cmd.Flags().Changed(rollFlagName) {

		var rollFlagName string
		if cmdPrefix == "" {
			rollFlagName = "roll"
		} else {
			rollFlagName = fmt.Sprintf("%v.roll", cmdPrefix)
		}

		rollFlagValue, err := cmd.Flags().GetFloat64(rollFlagName)
		if err != nil {
			return err, false
		}
		m.Roll = &rollFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceLocationTiltFlags(depth int, m *models.DeviceLocation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	tiltFlagName := fmt.Sprintf("%v.tilt", cmdPrefix)
	if cmd.Flags().Changed(tiltFlagName) {

		var tiltFlagName string
		if cmdPrefix == "" {
			tiltFlagName = "tilt"
		} else {
			tiltFlagName = fmt.Sprintf("%v.tilt", cmdPrefix)
		}

		tiltFlagValue, err := cmd.Flags().GetFloat64(tiltFlagName)
		if err != nil {
			return err, false
		}
		m.Tilt = &tiltFlagValue

		retAdded = true
	}

	return nil, retAdded
}
