// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for InterfacePoe

// register flags to command
func registerModelInterfacePoeFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerInterfacePoeCapacities(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInterfacePoeOutput(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerInterfacePoeCapacities(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: capacities Capacities array type is not supported by go-swagger cli yet

	return nil
}

func registerInterfacePoeOutput(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	outputDescription := `Enum: ["off","active","24v","48v","54v","24v-4pair","54v-4pair","pthru"]. `

	var outputFlagName string
	if cmdPrefix == "" {
		outputFlagName = "output"
	} else {
		outputFlagName = fmt.Sprintf("%v.output", cmdPrefix)
	}

	var outputFlagDefault string

	_ = cmd.PersistentFlags().String(outputFlagName, outputFlagDefault, outputDescription)

	if err := cmd.RegisterFlagCompletionFunc(outputFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["off","active","24v","48v","54v","24v-4pair","54v-4pair","pthru"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelInterfacePoeFlags(depth int, m *models.InterfacePoe, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, capacitiesAdded := retrieveInterfacePoeCapacitiesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || capacitiesAdded

	err, outputAdded := retrieveInterfacePoeOutputFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || outputAdded

	return nil, retAdded
}

func retrieveInterfacePoeCapacitiesFlags(depth int, m *models.InterfacePoe, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	capacitiesFlagName := fmt.Sprintf("%v.capacities", cmdPrefix)
	if cmd.Flags().Changed(capacitiesFlagName) {
		// warning: capacities array type Capacities is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveInterfacePoeOutputFlags(depth int, m *models.InterfacePoe, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	outputFlagName := fmt.Sprintf("%v.output", cmdPrefix)
	if cmd.Flags().Changed(outputFlagName) {

		var outputFlagName string
		if cmdPrefix == "" {
			outputFlagName = "output"
		} else {
			outputFlagName = fmt.Sprintf("%v.output", cmdPrefix)
		}

		outputFlagValue, err := cmd.Flags().GetString(outputFlagName)
		if err != nil {
			return err, false
		}
		m.Output = outputFlagValue

		retAdded = true
	}

	return nil, retAdded
}