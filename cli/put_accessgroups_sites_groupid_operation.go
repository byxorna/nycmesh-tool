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

// makeOperationAuthorizationPutAccessgroupsSitesGroupidCmd returns a cmd to handle operation putAccessgroupsSitesGroupid
func makeOperationAuthorizationPutAccessgroupsSitesGroupidCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "putAccessgroupsSitesGroupid",
		Short: ``,
		RunE:  runOperationAuthorizationPutAccessgroupsSitesGroupid,
	}

	if err := registerOperationAuthorizationPutAccessgroupsSitesGroupidParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationAuthorizationPutAccessgroupsSitesGroupid uses cmd flags to call endpoint api
func runOperationAuthorizationPutAccessgroupsSitesGroupid(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := authorization.NewPutAccessgroupsSitesGroupidParams()
	if err, _ := retrieveOperationAuthorizationPutAccessgroupsSitesGroupidBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationAuthorizationPutAccessgroupsSitesGroupidGroupIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationAuthorizationPutAccessgroupsSitesGroupidResult(appCli.Authorization.PutAccessgroupsSitesGroupid(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationAuthorizationPutAccessgroupsSitesGroupidParamFlags registers all flags needed to fill params
func registerOperationAuthorizationPutAccessgroupsSitesGroupidParamFlags(cmd *cobra.Command) error {
	if err := registerOperationAuthorizationPutAccessgroupsSitesGroupidBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationAuthorizationPutAccessgroupsSitesGroupidGroupIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationAuthorizationPutAccessgroupsSitesGroupidBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelRequestSiteAccessGroupFlags(0, "requestSiteAccessGroup", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationAuthorizationPutAccessgroupsSitesGroupidGroupIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	groupIdDescription := `Required. `

	var groupIdFlagName string
	if cmdPrefix == "" {
		groupIdFlagName = "groupId"
	} else {
		groupIdFlagName = fmt.Sprintf("%v.groupId", cmdPrefix)
	}

	var groupIdFlagDefault string

	_ = cmd.PersistentFlags().String(groupIdFlagName, groupIdFlagDefault, groupIdDescription)

	return nil
}

func retrieveOperationAuthorizationPutAccessgroupsSitesGroupidBodyFlag(m *authorization.PutAccessgroupsSitesGroupidParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.RequestSiteAccessGroup{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.RequestSiteAccessGroup: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.RequestSiteAccessGroup{}
	}
	err, added := retrieveModelRequestSiteAccessGroupFlags(0, bodyValueModel, "requestSiteAccessGroup", cmd)
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
func retrieveOperationAuthorizationPutAccessgroupsSitesGroupidGroupIDFlag(m *authorization.PutAccessgroupsSitesGroupidParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("groupId") {

		var groupIdFlagName string
		if cmdPrefix == "" {
			groupIdFlagName = "groupId"
		} else {
			groupIdFlagName = fmt.Sprintf("%v.groupId", cmdPrefix)
		}

		groupIdFlagValue, err := cmd.Flags().GetString(groupIdFlagName)
		if err != nil {
			return err, false
		}
		m.GroupID = groupIdFlagValue

	}
	return nil, retAdded
}

// parseOperationAuthorizationPutAccessgroupsSitesGroupidResult parses request result and return the string content
func parseOperationAuthorizationPutAccessgroupsSitesGroupidResult(resp0 *authorization.PutAccessgroupsSitesGroupidOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*authorization.PutAccessgroupsSitesGroupidOK)
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
		resp1, ok := iResp1.(*authorization.PutAccessgroupsSitesGroupidBadRequest)
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
		resp2, ok := iResp2.(*authorization.PutAccessgroupsSitesGroupidUnauthorized)
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
		resp3, ok := iResp3.(*authorization.PutAccessgroupsSitesGroupidForbidden)
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
		resp4, ok := iResp4.(*authorization.PutAccessgroupsSitesGroupidNotFound)
		if ok {
			if !swag.IsZero(resp4) && !swag.IsZero(resp4.Payload) {
				msgStr, err := json.Marshal(resp4.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp5 interface{} = respErr
		resp5, ok := iResp5.(*authorization.PutAccessgroupsSitesGroupidConflict)
		if ok {
			if !swag.IsZero(resp5) && !swag.IsZero(resp5.Payload) {
				msgStr, err := json.Marshal(resp5.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp6 interface{} = respErr
		resp6, ok := iResp6.(*authorization.PutAccessgroupsSitesGroupidInternalServerError)
		if ok {
			if !swag.IsZero(resp6) && !swag.IsZero(resp6.Payload) {
				msgStr, err := json.Marshal(resp6.Payload)
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