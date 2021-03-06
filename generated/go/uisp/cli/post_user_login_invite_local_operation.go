// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/authorization"
	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationAuthorizationPostUserLoginInviteLocalCmd returns a cmd to handle operation postUserLoginInviteLocal
func makeOperationAuthorizationPostUserLoginInviteLocalCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "postUserLoginInviteLocal",
		Short: ``,
		RunE:  runOperationAuthorizationPostUserLoginInviteLocal,
	}

	if err := registerOperationAuthorizationPostUserLoginInviteLocalParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationAuthorizationPostUserLoginInviteLocal uses cmd flags to call endpoint api
func runOperationAuthorizationPostUserLoginInviteLocal(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := authorization.NewPostUserLoginInviteLocalParams()
	if err, _ := retrieveOperationAuthorizationPostUserLoginInviteLocalBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationAuthorizationPostUserLoginInviteLocalResult(appCli.Authorization.PostUserLoginInviteLocal(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationAuthorizationPostUserLoginInviteLocalParamFlags registers all flags needed to fill params
func registerOperationAuthorizationPostUserLoginInviteLocalParamFlags(cmd *cobra.Command) error {
	if err := registerOperationAuthorizationPostUserLoginInviteLocalBodyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationAuthorizationPostUserLoginInviteLocalBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelLocalInvitationRequestFlags(0, "localInvitationRequest", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationAuthorizationPostUserLoginInviteLocalBodyFlag(m *authorization.PostUserLoginInviteLocalParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.LocalInvitationRequest{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.LocalInvitationRequest: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.LocalInvitationRequest{}
	}
	err, added := retrieveModelLocalInvitationRequestFlags(0, bodyValueModel, "localInvitationRequest", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Body = bodyValueModel
	}
	if dryRun && debug {

		bodyValueDebugBytes, err := json.Marshal(m.Body)
		if err != nil {
			return err, false
		}
		logDebugf("Body dry-run payload: %v", string(bodyValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}

// parseOperationAuthorizationPostUserLoginInviteLocalResult parses request result and return the string content
func parseOperationAuthorizationPostUserLoginInviteLocalResult(resp0 *authorization.PostUserLoginInviteLocalOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*authorization.PostUserLoginInviteLocalOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*authorization.PostUserLoginInviteLocalBadRequest)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*authorization.PostUserLoginInviteLocalUnauthorized)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp3 interface{} = respErr
		resp3, ok := iResp3.(*authorization.PostUserLoginInviteLocalInternalServerError)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
