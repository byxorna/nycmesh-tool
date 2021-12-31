// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for RedistributeStatic

// register flags to command
func registerModelRedistributeStaticFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerRedistributeStaticEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRedistributeStaticMetric(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerRedistributeStaticEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	enabledDescription := `Required. `

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

func registerRedistributeStaticMetric(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	metricDescription := `Required. `

	var metricFlagName string
	if cmdPrefix == "" {
		metricFlagName = "metric"
	} else {
		metricFlagName = fmt.Sprintf("%v.metric", cmdPrefix)
	}

	var metricFlagDefault float64

	_ = cmd.PersistentFlags().Float64(metricFlagName, metricFlagDefault, metricDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelRedistributeStaticFlags(depth int, m *models.RedistributeStatic, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, enabledAdded := retrieveRedistributeStaticEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || enabledAdded

	err, metricAdded := retrieveRedistributeStaticMetricFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || metricAdded

	return nil, retAdded
}

func retrieveRedistributeStaticEnabledFlags(depth int, m *models.RedistributeStatic, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveRedistributeStaticMetricFlags(depth int, m *models.RedistributeStatic, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	metricFlagName := fmt.Sprintf("%v.metric", cmdPrefix)
	if cmd.Flags().Changed(metricFlagName) {

		var metricFlagName string
		if cmdPrefix == "" {
			metricFlagName = "metric"
		} else {
			metricFlagName = fmt.Sprintf("%v.metric", cmdPrefix)
		}

		metricFlagValue, err := cmd.Flags().GetFloat64(metricFlagName)
		if err != nil {
			return err, false
		}
		m.Metric = &metricFlagValue

		retAdded = true
	}

	return nil, retAdded
}
