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

// Schema cli for Model24

// register flags to command
func registerModelModel24Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel24Error(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel24Status(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel24Error(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	errorDescription := `Required. `

	var errorFlagName string
	if cmdPrefix == "" {
		errorFlagName = "error"
	} else {
		errorFlagName = fmt.Sprintf("%v.error", cmdPrefix)
	}

	var errorFlagDefault string

	_ = cmd.PersistentFlags().String(errorFlagName, errorFlagDefault, errorDescription)

	return nil
}

func registerModel24Status(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statusDescription := `Enum: ["running","complete","error","none"]. `

	var statusFlagName string
	if cmdPrefix == "" {
		statusFlagName = "status"
	} else {
		statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
	}

	var statusFlagDefault string

	_ = cmd.PersistentFlags().String(statusFlagName, statusFlagDefault, statusDescription)

	if err := cmd.RegisterFlagCompletionFunc(statusFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["running","complete","error","none"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModel24Flags(depth int, m *models.Model24, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, errorAdded := retrieveModel24ErrorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || errorAdded

	err, statusAdded := retrieveModel24StatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusAdded

	return nil, retAdded
}

func retrieveModel24ErrorFlags(depth int, m *models.Model24, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	errorFlagName := fmt.Sprintf("%v.error", cmdPrefix)
	if cmd.Flags().Changed(errorFlagName) {

		var errorFlagName string
		if cmdPrefix == "" {
			errorFlagName = "error"
		} else {
			errorFlagName = fmt.Sprintf("%v.error", cmdPrefix)
		}

		errorFlagValue, err := cmd.Flags().GetString(errorFlagName)
		if err != nil {
			return err, false
		}
		m.Error = &errorFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel24StatusFlags(depth int, m *models.Model24, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statusFlagName := fmt.Sprintf("%v.status", cmdPrefix)
	if cmd.Flags().Changed(statusFlagName) {

		var statusFlagName string
		if cmdPrefix == "" {
			statusFlagName = "status"
		} else {
			statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
		}

		statusFlagValue, err := cmd.Flags().GetString(statusFlagName)
		if err != nil {
			return err, false
		}
		m.Status = statusFlagValue

		retAdded = true
	}

	return nil, retAdded
}