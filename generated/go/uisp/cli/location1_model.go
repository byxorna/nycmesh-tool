// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for Location1

// register flags to command
func registerModelLocation1Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerLocation1Lat(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerLocation1Lng(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerLocation1Lat(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	latDescription := `Required. `

	var latFlagName string
	if cmdPrefix == "" {
		latFlagName = "lat"
	} else {
		latFlagName = fmt.Sprintf("%v.lat", cmdPrefix)
	}

	var latFlagDefault float64

	_ = cmd.PersistentFlags().Float64(latFlagName, latFlagDefault, latDescription)

	return nil
}

func registerLocation1Lng(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	lngDescription := `Required. `

	var lngFlagName string
	if cmdPrefix == "" {
		lngFlagName = "lng"
	} else {
		lngFlagName = fmt.Sprintf("%v.lng", cmdPrefix)
	}

	var lngFlagDefault float64

	_ = cmd.PersistentFlags().Float64(lngFlagName, lngFlagDefault, lngDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelLocation1Flags(depth int, m *models.Location1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, latAdded := retrieveLocation1LatFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || latAdded

	err, lngAdded := retrieveLocation1LngFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || lngAdded

	return nil, retAdded
}

func retrieveLocation1LatFlags(depth int, m *models.Location1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	latFlagName := fmt.Sprintf("%v.lat", cmdPrefix)
	if cmd.Flags().Changed(latFlagName) {

		var latFlagName string
		if cmdPrefix == "" {
			latFlagName = "lat"
		} else {
			latFlagName = fmt.Sprintf("%v.lat", cmdPrefix)
		}

		latFlagValue, err := cmd.Flags().GetFloat64(latFlagName)
		if err != nil {
			return err, false
		}
		m.Lat = &latFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveLocation1LngFlags(depth int, m *models.Location1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	lngFlagName := fmt.Sprintf("%v.lng", cmdPrefix)
	if cmd.Flags().Changed(lngFlagName) {

		var lngFlagName string
		if cmdPrefix == "" {
			lngFlagName = "lng"
		} else {
			lngFlagName = fmt.Sprintf("%v.lng", cmdPrefix)
		}

		lngFlagValue, err := cmd.Flags().GetFloat64(lngFlagName)
		if err != nil {
			return err, false
		}
		m.Lng = &lngFlagValue

		retAdded = true
	}

	return nil, retAdded
}
