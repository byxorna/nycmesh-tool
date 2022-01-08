// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for SsoInvitationRequest

// register flags to command
func registerModelSsoInvitationRequestFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSsoInvitationRequestCode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSsoInvitationRequestInvitationToken(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSsoInvitationRequestState(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSsoInvitationRequestCode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	codeDescription := `Required. Code from successful SSO authentication.`

	var codeFlagName string
	if cmdPrefix == "" {
		codeFlagName = "code"
	} else {
		codeFlagName = fmt.Sprintf("%v.code", cmdPrefix)
	}

	var codeFlagDefault string

	_ = cmd.PersistentFlags().String(codeFlagName, codeFlagDefault, codeDescription)

	return nil
}

func registerSsoInvitationRequestInvitationToken(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	invitationTokenDescription := `User invitation token.`

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

func registerSsoInvitationRequestState(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	stateDescription := `Required. State previously sent to SSO. Same value must be in nms-sso-login cookie.`

	var stateFlagName string
	if cmdPrefix == "" {
		stateFlagName = "state"
	} else {
		stateFlagName = fmt.Sprintf("%v.state", cmdPrefix)
	}

	var stateFlagDefault string

	_ = cmd.PersistentFlags().String(stateFlagName, stateFlagDefault, stateDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSsoInvitationRequestFlags(depth int, m *models.SsoInvitationRequest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, codeAdded := retrieveSsoInvitationRequestCodeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || codeAdded

	err, invitationTokenAdded := retrieveSsoInvitationRequestInvitationTokenFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || invitationTokenAdded

	err, stateAdded := retrieveSsoInvitationRequestStateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || stateAdded

	return nil, retAdded
}

func retrieveSsoInvitationRequestCodeFlags(depth int, m *models.SsoInvitationRequest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	codeFlagName := fmt.Sprintf("%v.code", cmdPrefix)
	if cmd.Flags().Changed(codeFlagName) {

		var codeFlagName string
		if cmdPrefix == "" {
			codeFlagName = "code"
		} else {
			codeFlagName = fmt.Sprintf("%v.code", cmdPrefix)
		}

		codeFlagValue, err := cmd.Flags().GetString(codeFlagName)
		if err != nil {
			return err, false
		}
		m.Code = &codeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSsoInvitationRequestInvitationTokenFlags(depth int, m *models.SsoInvitationRequest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.InvitationToken = invitationTokenFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSsoInvitationRequestStateFlags(depth int, m *models.SsoInvitationRequest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	stateFlagName := fmt.Sprintf("%v.state", cmdPrefix)
	if cmd.Flags().Changed(stateFlagName) {

		var stateFlagName string
		if cmdPrefix == "" {
			stateFlagName = "state"
		} else {
			stateFlagName = fmt.Sprintf("%v.state", cmdPrefix)
		}

		stateFlagValue, err := cmd.Flags().GetString(stateFlagName)
		if err != nil {
			return err, false
		}
		m.State = &stateFlagValue

		retAdded = true
	}

	return nil, retAdded
}
