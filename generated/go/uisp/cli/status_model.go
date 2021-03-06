// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for Status

// register flags to command
func registerModelStatusFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerStatusMessage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStatusResult(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerStatusMessage(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	messageDescription := `Required. `

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

func registerStatusResult(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	resultDescription := `Required. `

	var resultFlagName string
	if cmdPrefix == "" {
		resultFlagName = "result"
	} else {
		resultFlagName = fmt.Sprintf("%v.result", cmdPrefix)
	}

	var resultFlagDefault bool

	_ = cmd.PersistentFlags().Bool(resultFlagName, resultFlagDefault, resultDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelStatusFlags(depth int, m *models.Status, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, messageAdded := retrieveStatusMessageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || messageAdded

	err, resultAdded := retrieveStatusResultFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || resultAdded

	return nil, retAdded
}

func retrieveStatusMessageFlags(depth int, m *models.Status, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Message = &messageFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStatusResultFlags(depth int, m *models.Status, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	resultFlagName := fmt.Sprintf("%v.result", cmdPrefix)
	if cmd.Flags().Changed(resultFlagName) {

		var resultFlagName string
		if cmdPrefix == "" {
			resultFlagName = "result"
		} else {
			resultFlagName = fmt.Sprintf("%v.result", cmdPrefix)
		}

		resultFlagValue, err := cmd.Flags().GetBool(resultFlagName)
		if err != nil {
			return err, false
		}
		m.Result = &resultFlagValue

		retAdded = true
	}

	return nil, retAdded
}
