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

// Schema cli for ServerLogVerbosityPayloadSchema

// register flags to command
func registerModelServerLogVerbosityPayloadSchemaFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerServerLogVerbosityPayloadSchemaDuration(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerLogVerbosityPayloadSchemaVerbosity(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerServerLogVerbosityPayloadSchemaDuration(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	durationDescription := `Enum: [-1]. Required. Duration in milliseconds.`

	var durationFlagName string
	if cmdPrefix == "" {
		durationFlagName = "duration"
	} else {
		durationFlagName = fmt.Sprintf("%v.duration", cmdPrefix)
	}

	var durationFlagDefault int64

	_ = cmd.PersistentFlags().Int64(durationFlagName, durationFlagDefault, durationDescription)

	if err := cmd.RegisterFlagCompletionFunc(durationFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`[-1]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerServerLogVerbosityPayloadSchemaVerbosity(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	verbosityDescription := `Enum: ["trace","debug","info","warn","error","fatal"]. Required. `

	var verbosityFlagName string
	if cmdPrefix == "" {
		verbosityFlagName = "verbosity"
	} else {
		verbosityFlagName = fmt.Sprintf("%v.verbosity", cmdPrefix)
	}

	var verbosityFlagDefault string

	_ = cmd.PersistentFlags().String(verbosityFlagName, verbosityFlagDefault, verbosityDescription)

	if err := cmd.RegisterFlagCompletionFunc(verbosityFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["trace","debug","info","warn","error","fatal"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServerLogVerbosityPayloadSchemaFlags(depth int, m *models.ServerLogVerbosityPayloadSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, durationAdded := retrieveServerLogVerbosityPayloadSchemaDurationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || durationAdded

	err, verbosityAdded := retrieveServerLogVerbosityPayloadSchemaVerbosityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || verbosityAdded

	return nil, retAdded
}

func retrieveServerLogVerbosityPayloadSchemaDurationFlags(depth int, m *models.ServerLogVerbosityPayloadSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

		durationFlagValue, err := cmd.Flags().GetInt64(durationFlagName)
		if err != nil {
			return err, false
		}
		m.Duration = &durationFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerLogVerbosityPayloadSchemaVerbosityFlags(depth int, m *models.ServerLogVerbosityPayloadSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	verbosityFlagName := fmt.Sprintf("%v.verbosity", cmdPrefix)
	if cmd.Flags().Changed(verbosityFlagName) {

		var verbosityFlagName string
		if cmdPrefix == "" {
			verbosityFlagName = "verbosity"
		} else {
			verbosityFlagName = fmt.Sprintf("%v.verbosity", cmdPrefix)
		}

		verbosityFlagValue, err := cmd.Flags().GetString(verbosityFlagName)
		if err != nil {
			return err, false
		}
		m.Verbosity = &verbosityFlagValue

		retAdded = true
	}

	return nil, retAdded
}
