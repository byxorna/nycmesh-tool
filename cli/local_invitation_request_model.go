// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for LocalInvitationRequest

// register flags to command
func registerModelLocalInvitationRequestFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerLocalInvitationRequestInvitationToken(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerLocalInvitationRequestPassword(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerLocalInvitationRequestInvitationToken(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	invitationTokenDescription := `Required. User invitation token.`

	var invitationTokenFlagName string
	if cmdPrefix == "" {
		invitationTokenFlagName = "invitationToken"
	} else {
		invitationTokenFlagName = fmt.Sprintf("%v.invitationToken", cmdPrefix)
	}

	var invitationTokenFlagDefault string

	_ = cmd.PersistentFlags().String(invitationTokenFlagName, invitationTokenFlagDefault, invitationTokenDescription)

	return nil
}

func registerLocalInvitationRequestPassword(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	passwordDescription := `Required. New user password.`

	var passwordFlagName string
	if cmdPrefix == "" {
		passwordFlagName = "password"
	} else {
		passwordFlagName = fmt.Sprintf("%v.password", cmdPrefix)
	}

	var passwordFlagDefault string

	_ = cmd.PersistentFlags().String(passwordFlagName, passwordFlagDefault, passwordDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelLocalInvitationRequestFlags(depth int, m *models.LocalInvitationRequest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, invitationTokenAdded := retrieveLocalInvitationRequestInvitationTokenFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || invitationTokenAdded

	err, passwordAdded := retrieveLocalInvitationRequestPasswordFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || passwordAdded

	return nil, retAdded
}

func retrieveLocalInvitationRequestInvitationTokenFlags(depth int, m *models.LocalInvitationRequest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	invitationTokenFlagName := fmt.Sprintf("%v.invitationToken", cmdPrefix)
	if cmd.Flags().Changed(invitationTokenFlagName) {

		var invitationTokenFlagName string
		if cmdPrefix == "" {
			invitationTokenFlagName = "invitationToken"
		} else {
			invitationTokenFlagName = fmt.Sprintf("%v.invitationToken", cmdPrefix)
		}

		invitationTokenFlagValue, err := cmd.Flags().GetString(invitationTokenFlagName)
		if err != nil {
			return err, false
		}
		m.InvitationToken = &invitationTokenFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveLocalInvitationRequestPasswordFlags(depth int, m *models.LocalInvitationRequest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	passwordFlagName := fmt.Sprintf("%v.password", cmdPrefix)
	if cmd.Flags().Changed(passwordFlagName) {

		var passwordFlagName string
		if cmdPrefix == "" {
			passwordFlagName = "password"
		} else {
			passwordFlagName = fmt.Sprintf("%v.password", cmdPrefix)
		}

		passwordFlagValue, err := cmd.Flags().GetString(passwordFlagName)
		if err != nil {
			return err, false
		}
		m.Password = &passwordFlagValue

		retAdded = true
	}

	return nil, retAdded
}