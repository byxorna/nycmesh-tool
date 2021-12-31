// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for Error

// register flags to command
func registerModelErrorFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerErrorError(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerErrorMessage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerErrorStatusCode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerErrorValidation(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerErrorError(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerErrorMessage(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	messageDescription := ``

	var messageFlagName string
	if cmdPrefix == "" {
		messageFlagName = "message"
	} else {
		messageFlagName = fmt.Sprintf("%v.message", cmdPrefix)
	}

	var messageFlagDefault string

	_ = cmd.PersistentFlags().String(messageFlagName, messageFlagDefault, messageDescription)

	return nil
}

func registerErrorStatusCode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statusCodeDescription := `Required. `

	var statusCodeFlagName string
	if cmdPrefix == "" {
		statusCodeFlagName = "statusCode"
	} else {
		statusCodeFlagName = fmt.Sprintf("%v.statusCode", cmdPrefix)
	}

	var statusCodeFlagDefault float64

	_ = cmd.PersistentFlags().Float64(statusCodeFlagName, statusCodeFlagDefault, statusCodeDescription)

	return nil
}

func registerErrorValidation(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: validation Validation map type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelErrorFlags(depth int, m *models.Error, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, errorAdded := retrieveErrorErrorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || errorAdded

	err, messageAdded := retrieveErrorMessageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || messageAdded

	err, statusCodeAdded := retrieveErrorStatusCodeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusCodeAdded

	err, validationAdded := retrieveErrorValidationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || validationAdded

	return nil, retAdded
}

func retrieveErrorErrorFlags(depth int, m *models.Error, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveErrorMessageFlags(depth int, m *models.Error, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	messageFlagName := fmt.Sprintf("%v.message", cmdPrefix)
	if cmd.Flags().Changed(messageFlagName) {

		var messageFlagName string
		if cmdPrefix == "" {
			messageFlagName = "message"
		} else {
			messageFlagName = fmt.Sprintf("%v.message", cmdPrefix)
		}

		messageFlagValue, err := cmd.Flags().GetString(messageFlagName)
		if err != nil {
			return err, false
		}
		m.Message = messageFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveErrorStatusCodeFlags(depth int, m *models.Error, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statusCodeFlagName := fmt.Sprintf("%v.statusCode", cmdPrefix)
	if cmd.Flags().Changed(statusCodeFlagName) {

		var statusCodeFlagName string
		if cmdPrefix == "" {
			statusCodeFlagName = "statusCode"
		} else {
			statusCodeFlagName = fmt.Sprintf("%v.statusCode", cmdPrefix)
		}

		statusCodeFlagValue, err := cmd.Flags().GetFloat64(statusCodeFlagName)
		if err != nil {
			return err, false
		}
		m.StatusCode = &statusCodeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveErrorValidationFlags(depth int, m *models.Error, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	validationFlagName := fmt.Sprintf("%v.validation", cmdPrefix)
	if cmd.Flags().Changed(validationFlagName) {
		// warning: validation map type Validation is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
