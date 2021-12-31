// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for Coordinates

// register flags to command
func registerModelCoordinatesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCoordinatesX(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCoordinatesY(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCoordinatesX(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	xDescription := ``

	var xFlagName string
	if cmdPrefix == "" {
		xFlagName = "x"
	} else {
		xFlagName = fmt.Sprintf("%v.x", cmdPrefix)
	}

	var xFlagDefault float64

	_ = cmd.PersistentFlags().Float64(xFlagName, xFlagDefault, xDescription)

	return nil
}

func registerCoordinatesY(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	yDescription := ``

	var yFlagName string
	if cmdPrefix == "" {
		yFlagName = "y"
	} else {
		yFlagName = fmt.Sprintf("%v.y", cmdPrefix)
	}

	var yFlagDefault float64

	_ = cmd.PersistentFlags().Float64(yFlagName, yFlagDefault, yDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCoordinatesFlags(depth int, m *models.Coordinates, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, xAdded := retrieveCoordinatesXFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || xAdded

	err, yAdded := retrieveCoordinatesYFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || yAdded

	return nil, retAdded
}

func retrieveCoordinatesXFlags(depth int, m *models.Coordinates, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	xFlagName := fmt.Sprintf("%v.x", cmdPrefix)
	if cmd.Flags().Changed(xFlagName) {

		var xFlagName string
		if cmdPrefix == "" {
			xFlagName = "x"
		} else {
			xFlagName = fmt.Sprintf("%v.x", cmdPrefix)
		}

		xFlagValue, err := cmd.Flags().GetFloat64(xFlagName)
		if err != nil {
			return err, false
		}
		m.X = xFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCoordinatesYFlags(depth int, m *models.Coordinates, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	yFlagName := fmt.Sprintf("%v.y", cmdPrefix)
	if cmd.Flags().Changed(yFlagName) {

		var yFlagName string
		if cmdPrefix == "" {
			yFlagName = "y"
		} else {
			yFlagName = fmt.Sprintf("%v.y", cmdPrefix)
		}

		yFlagValue, err := cmd.Flags().GetFloat64(yFlagName)
		if err != nil {
			return err, false
		}
		m.Y = yFlagValue

		retAdded = true
	}

	return nil, retAdded
}
