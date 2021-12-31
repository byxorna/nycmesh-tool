// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/client/authorization"
	"github.com/byxorna/nycmesh-tool/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationAuthorizationPostUserLastreleasenotesseenCmd returns a cmd to handle operation postUserLastreleasenotesseen
func makeOperationAuthorizationPostUserLastreleasenotesseenCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "postUserLastreleasenotesseen",
		Short: ``,
		RunE:  runOperationAuthorizationPostUserLastreleasenotesseen,
	}

	if err := registerOperationAuthorizationPostUserLastreleasenotesseenParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationAuthorizationPostUserLastreleasenotesseen uses cmd flags to call endpoint api
func runOperationAuthorizationPostUserLastreleasenotesseen(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := authorization.NewPostUserLastreleasenotesseenParams()
	if err, _ := retrieveOperationAuthorizationPostUserLastreleasenotesseenBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationAuthorizationPostUserLastreleasenotesseenResult(appCli.Authorization.PostUserLastreleasenotesseen(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationAuthorizationPostUserLastreleasenotesseenParamFlags registers all flags needed to fill params
func registerOperationAuthorizationPostUserLastreleasenotesseenParamFlags(cmd *cobra.Command) error {
	if err := registerOperationAuthorizationPostUserLastreleasenotesseenBodyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationAuthorizationPostUserLastreleasenotesseenBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelUserLastReleaseNotesSeenFlags(0, "userLastReleaseNotesSeen", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationAuthorizationPostUserLastreleasenotesseenBodyFlag(m *authorization.PostUserLastreleasenotesseenParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.UserLastReleaseNotesSeen{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.UserLastReleaseNotesSeen: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.UserLastReleaseNotesSeen{}
	}
	err, added := retrieveModelUserLastReleaseNotesSeenFlags(0, bodyValueModel, "userLastReleaseNotesSeen", cmd)
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

// parseOperationAuthorizationPostUserLastreleasenotesseenResult parses request result and return the string content
func parseOperationAuthorizationPostUserLastreleasenotesseenResult(resp0 *authorization.PostUserLastreleasenotesseenOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*authorization.PostUserLastreleasenotesseenOK)
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
		resp1, ok := iResp1.(*authorization.PostUserLastreleasenotesseenBadRequest)
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
		resp2, ok := iResp2.(*authorization.PostUserLastreleasenotesseenUnauthorized)
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
		resp3, ok := iResp3.(*authorization.PostUserLastreleasenotesseenForbidden)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp4 interface{} = respErr
		resp4, ok := iResp4.(*authorization.PostUserLastreleasenotesseenInternalServerError)
		if ok {
			if !swag.IsZero(resp4) && !swag.IsZero(resp4.Payload) {
				msgStr, err := json.Marshal(resp4.Payload)
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
