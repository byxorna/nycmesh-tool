// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for InstallationChatPayload

// register flags to command
func registerModelInstallationChatPayloadFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerInstallationChatPayloadMessage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerInstallationChatPayloadMessage(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	messageDescription := `Required. Chat message text.`

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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelInstallationChatPayloadFlags(depth int, m *models.InstallationChatPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, messageAdded := retrieveInstallationChatPayloadMessageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || messageAdded

	return nil, retAdded
}

func retrieveInstallationChatPayloadMessageFlags(depth int, m *models.InstallationChatPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
