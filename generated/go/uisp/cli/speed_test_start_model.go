// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for SpeedTestStart

// register flags to command
func registerModelSpeedTestStartFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSpeedTestStartID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSpeedTestStartMessage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSpeedTestStartResult(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSpeedTestStartID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `Required. `

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerSpeedTestStartMessage(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerSpeedTestStartResult(depth int, cmdPrefix string, cmd *cobra.Command) error {
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
func retrieveModelSpeedTestStartFlags(depth int, m *models.SpeedTestStart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, idAdded := retrieveSpeedTestStartIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, messageAdded := retrieveSpeedTestStartMessageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || messageAdded

	err, resultAdded := retrieveSpeedTestStartResultFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || resultAdded

	return nil, retAdded
}

func retrieveSpeedTestStartIDFlags(depth int, m *models.SpeedTestStart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = &idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSpeedTestStartMessageFlags(depth int, m *models.SpeedTestStart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveSpeedTestStartResultFlags(depth int, m *models.SpeedTestStart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
