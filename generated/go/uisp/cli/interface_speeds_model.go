// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for InterfaceSpeeds

// register flags to command
func registerModelInterfaceSpeedsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerInterfaceSpeedsCapacities(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInterfaceSpeedsSpeed(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerInterfaceSpeedsCapacities(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: capacities Capacities1 array type is not supported by go-swagger cli yet

	return nil
}

func registerInterfaceSpeedsSpeed(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	speedDescription := ``

	var speedFlagName string
	if cmdPrefix == "" {
		speedFlagName = "speed"
	} else {
		speedFlagName = fmt.Sprintf("%v.speed", cmdPrefix)
	}

	var speedFlagDefault string

	_ = cmd.PersistentFlags().String(speedFlagName, speedFlagDefault, speedDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelInterfaceSpeedsFlags(depth int, m *models.InterfaceSpeeds, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, capacitiesAdded := retrieveInterfaceSpeedsCapacitiesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || capacitiesAdded

	err, speedAdded := retrieveInterfaceSpeedsSpeedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || speedAdded

	return nil, retAdded
}

func retrieveInterfaceSpeedsCapacitiesFlags(depth int, m *models.InterfaceSpeeds, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	capacitiesFlagName := fmt.Sprintf("%v.capacities", cmdPrefix)
	if cmd.Flags().Changed(capacitiesFlagName) {
		// warning: capacities array type Capacities1 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveInterfaceSpeedsSpeedFlags(depth int, m *models.InterfaceSpeeds, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	speedFlagName := fmt.Sprintf("%v.speed", cmdPrefix)
	if cmd.Flags().Changed(speedFlagName) {

		var speedFlagName string
		if cmdPrefix == "" {
			speedFlagName = "speed"
		} else {
			speedFlagName = fmt.Sprintf("%v.speed", cmdPrefix)
		}

		speedFlagValue, err := cmd.Flags().GetString(speedFlagName)
		if err != nil {
			return err, false
		}
		m.Speed = speedFlagValue

		retAdded = true
	}

	return nil, retAdded
}
