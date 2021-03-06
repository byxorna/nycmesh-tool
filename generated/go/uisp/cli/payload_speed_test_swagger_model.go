// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for PayloadSpeedTest

// register flags to command
func registerModelPayloadSpeedTestFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPayloadSpeedTestDirection(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPayloadSpeedTestDuration(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPayloadSpeedTestSource(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPayloadSpeedTestTarget(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPayloadSpeedTestDirection(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	directionDescription := `Enum: ["uplink","downlink","bidirectional"]. Required. `

	var directionFlagName string
	if cmdPrefix == "" {
		directionFlagName = "direction"
	} else {
		directionFlagName = fmt.Sprintf("%v.direction", cmdPrefix)
	}

	var directionFlagDefault string

	_ = cmd.PersistentFlags().String(directionFlagName, directionFlagDefault, directionDescription)

	if err := cmd.RegisterFlagCompletionFunc(directionFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["uplink","downlink","bidirectional"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerPayloadSpeedTestDuration(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	durationDescription := `Required. milliseconds`

	var durationFlagName string
	if cmdPrefix == "" {
		durationFlagName = "duration"
	} else {
		durationFlagName = fmt.Sprintf("%v.duration", cmdPrefix)
	}

	var durationFlagDefault float64

	_ = cmd.PersistentFlags().Float64(durationFlagName, durationFlagDefault, durationDescription)

	return nil
}

func registerPayloadSpeedTestSource(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sourceDescription := `Required. `

	var sourceFlagName string
	if cmdPrefix == "" {
		sourceFlagName = "source"
	} else {
		sourceFlagName = fmt.Sprintf("%v.source", cmdPrefix)
	}

	var sourceFlagDefault string

	_ = cmd.PersistentFlags().String(sourceFlagName, sourceFlagDefault, sourceDescription)

	return nil
}

func registerPayloadSpeedTestTarget(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	targetDescription := `Required. `

	var targetFlagName string
	if cmdPrefix == "" {
		targetFlagName = "target"
	} else {
		targetFlagName = fmt.Sprintf("%v.target", cmdPrefix)
	}

	var targetFlagDefault string

	_ = cmd.PersistentFlags().String(targetFlagName, targetFlagDefault, targetDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPayloadSpeedTestFlags(depth int, m *models.PayloadSpeedTest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, directionAdded := retrievePayloadSpeedTestDirectionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || directionAdded

	err, durationAdded := retrievePayloadSpeedTestDurationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || durationAdded

	err, sourceAdded := retrievePayloadSpeedTestSourceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sourceAdded

	err, targetAdded := retrievePayloadSpeedTestTargetFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || targetAdded

	return nil, retAdded
}

func retrievePayloadSpeedTestDirectionFlags(depth int, m *models.PayloadSpeedTest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	directionFlagName := fmt.Sprintf("%v.direction", cmdPrefix)
	if cmd.Flags().Changed(directionFlagName) {

		var directionFlagName string
		if cmdPrefix == "" {
			directionFlagName = "direction"
		} else {
			directionFlagName = fmt.Sprintf("%v.direction", cmdPrefix)
		}

		directionFlagValue, err := cmd.Flags().GetString(directionFlagName)
		if err != nil {
			return err, false
		}
		m.Direction = &directionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePayloadSpeedTestDurationFlags(depth int, m *models.PayloadSpeedTest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	durationFlagName := fmt.Sprintf("%v.duration", cmdPrefix)
	if cmd.Flags().Changed(durationFlagName) {

		var durationFlagName string
		if cmdPrefix == "" {
			durationFlagName = "duration"
		} else {
			durationFlagName = fmt.Sprintf("%v.duration", cmdPrefix)
		}

		durationFlagValue, err := cmd.Flags().GetFloat64(durationFlagName)
		if err != nil {
			return err, false
		}
		m.Duration = &durationFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePayloadSpeedTestSourceFlags(depth int, m *models.PayloadSpeedTest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sourceFlagName := fmt.Sprintf("%v.source", cmdPrefix)
	if cmd.Flags().Changed(sourceFlagName) {

		var sourceFlagName string
		if cmdPrefix == "" {
			sourceFlagName = "source"
		} else {
			sourceFlagName = fmt.Sprintf("%v.source", cmdPrefix)
		}

		sourceFlagValue, err := cmd.Flags().GetString(sourceFlagName)
		if err != nil {
			return err, false
		}
		m.Source = &sourceFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePayloadSpeedTestTargetFlags(depth int, m *models.PayloadSpeedTest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	targetFlagName := fmt.Sprintf("%v.target", cmdPrefix)
	if cmd.Flags().Changed(targetFlagName) {

		var targetFlagName string
		if cmdPrefix == "" {
			targetFlagName = "target"
		} else {
			targetFlagName = fmt.Sprintf("%v.target", cmdPrefix)
		}

		targetFlagValue, err := cmd.Flags().GetString(targetFlagName)
		if err != nil {
			return err, false
		}
		m.Target = &targetFlagValue

		retAdded = true
	}

	return nil, retAdded
}
