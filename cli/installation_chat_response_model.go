// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for InstallationChatResponse

// register flags to command
func registerModelInstallationChatResponseFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerInstallationChatResponseCreated(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInstallationChatResponseID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInstallationChatResponseMessage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInstallationChatResponseUser(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerInstallationChatResponseCreated(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: primitive created strfmt.Date is not supported by go-swagger cli yet

	return nil
}

func registerInstallationChatResponseID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `Required. Id of the message.`

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

func registerInstallationChatResponseMessage(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerInstallationChatResponseUser(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var userFlagName string
	if cmdPrefix == "" {
		userFlagName = "user"
	} else {
		userFlagName = fmt.Sprintf("%v.user", cmdPrefix)
	}

	if err := registerModelUser3Flags(depth+1, userFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelInstallationChatResponseFlags(depth int, m *models.InstallationChatResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, createdAdded := retrieveInstallationChatResponseCreatedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || createdAdded

	err, idAdded := retrieveInstallationChatResponseIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, messageAdded := retrieveInstallationChatResponseMessageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || messageAdded

	err, userAdded := retrieveInstallationChatResponseUserFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || userAdded

	return nil, retAdded
}

func retrieveInstallationChatResponseCreatedFlags(depth int, m *models.InstallationChatResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	createdFlagName := fmt.Sprintf("%v.created", cmdPrefix)
	if cmd.Flags().Changed(createdFlagName) {

		// warning: primitive created strfmt.Date is not supported by go-swagger cli yet

		retAdded = true
	}

	return nil, retAdded
}

func retrieveInstallationChatResponseIDFlags(depth int, m *models.InstallationChatResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveInstallationChatResponseMessageFlags(depth int, m *models.InstallationChatResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveInstallationChatResponseUserFlags(depth int, m *models.InstallationChatResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	userFlagName := fmt.Sprintf("%v.user", cmdPrefix)
	if cmd.Flags().Changed(userFlagName) {
		// info: complex object user User3 is retrieved outside this Changed() block
	}
	userFlagValue := m.User
	if swag.IsZero(userFlagValue) {
		userFlagValue = &models.User3{}
	}

	err, userAdded := retrieveModelUser3Flags(depth+1, userFlagValue, userFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || userAdded
	if userAdded {
		m.User = userFlagValue
	}

	return nil, retAdded
}
