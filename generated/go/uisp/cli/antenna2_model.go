// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for Antenna2

// register flags to command
func registerModelAntenna2Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerAntenna2Angle(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAntenna2Custom(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAntenna2Gain(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAntenna2Name(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerAntenna2Angle(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	angleDescription := `Required. `

	var angleFlagName string
	if cmdPrefix == "" {
		angleFlagName = "angle"
	} else {
		angleFlagName = fmt.Sprintf("%v.angle", cmdPrefix)
	}

	var angleFlagDefault float64

	_ = cmd.PersistentFlags().Float64(angleFlagName, angleFlagDefault, angleDescription)

	return nil
}

func registerAntenna2Custom(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	customDescription := `Required. `

	var customFlagName string
	if cmdPrefix == "" {
		customFlagName = "custom"
	} else {
		customFlagName = fmt.Sprintf("%v.custom", cmdPrefix)
	}

	var customFlagDefault bool

	_ = cmd.PersistentFlags().Bool(customFlagName, customFlagDefault, customDescription)

	return nil
}

func registerAntenna2Gain(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	gainDescription := `Required. `

	var gainFlagName string
	if cmdPrefix == "" {
		gainFlagName = "gain"
	} else {
		gainFlagName = fmt.Sprintf("%v.gain", cmdPrefix)
	}

	var gainFlagDefault int64

	_ = cmd.PersistentFlags().Int64(gainFlagName, gainFlagDefault, gainDescription)

	return nil
}

func registerAntenna2Name(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Required. `

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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelAntenna2Flags(depth int, m *models.Antenna2, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, angleAdded := retrieveAntenna2AngleFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || angleAdded

	err, customAdded := retrieveAntenna2CustomFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || customAdded

	err, gainAdded := retrieveAntenna2GainFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || gainAdded

	err, nameAdded := retrieveAntenna2NameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	return nil, retAdded
}

func retrieveAntenna2AngleFlags(depth int, m *models.Antenna2, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	angleFlagName := fmt.Sprintf("%v.angle", cmdPrefix)
	if cmd.Flags().Changed(angleFlagName) {

		var angleFlagName string
		if cmdPrefix == "" {
			angleFlagName = "angle"
		} else {
			angleFlagName = fmt.Sprintf("%v.angle", cmdPrefix)
		}

		angleFlagValue, err := cmd.Flags().GetFloat64(angleFlagName)
		if err != nil {
			return err, false
		}
		m.Angle = &angleFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAntenna2CustomFlags(depth int, m *models.Antenna2, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	customFlagName := fmt.Sprintf("%v.custom", cmdPrefix)
	if cmd.Flags().Changed(customFlagName) {

		var customFlagName string
		if cmdPrefix == "" {
			customFlagName = "custom"
		} else {
			customFlagName = fmt.Sprintf("%v.custom", cmdPrefix)
		}

		customFlagValue, err := cmd.Flags().GetBool(customFlagName)
		if err != nil {
			return err, false
		}
		m.Custom = &customFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAntenna2GainFlags(depth int, m *models.Antenna2, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	gainFlagName := fmt.Sprintf("%v.gain", cmdPrefix)
	if cmd.Flags().Changed(gainFlagName) {

		var gainFlagName string
		if cmdPrefix == "" {
			gainFlagName = "gain"
		} else {
			gainFlagName = fmt.Sprintf("%v.gain", cmdPrefix)
		}

		gainFlagValue, err := cmd.Flags().GetInt64(gainFlagName)
		if err != nil {
			return err, false
		}
		m.Gain = &gainFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAntenna2NameFlags(depth int, m *models.Antenna2, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Name = &nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}
