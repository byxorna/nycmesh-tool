// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/client/authorization"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationAuthorizationDeleteAccessgroupsSitesGroupidSiteidCmd returns a cmd to handle operation deleteAccessgroupsSitesGroupidSiteid
func makeOperationAuthorizationDeleteAccessgroupsSitesGroupidSiteidCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteAccessgroupsSitesGroupidSiteid",
		Short: ``,
		RunE:  runOperationAuthorizationDeleteAccessgroupsSitesGroupidSiteid,
	}

	if err := registerOperationAuthorizationDeleteAccessgroupsSitesGroupidSiteidParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationAuthorizationDeleteAccessgroupsSitesGroupidSiteid uses cmd flags to call endpoint api
func runOperationAuthorizationDeleteAccessgroupsSitesGroupidSiteid(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := authorization.NewDeleteAccessgroupsSitesGroupidSiteidParams()
	if err, _ := retrieveOperationAuthorizationDeleteAccessgroupsSitesGroupidSiteidGroupIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationAuthorizationDeleteAccessgroupsSitesGroupidSiteidSiteIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationAuthorizationDeleteAccessgroupsSitesGroupidSiteidResult(appCli.Authorization.DeleteAccessgroupsSitesGroupidSiteid(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationAuthorizationDeleteAccessgroupsSitesGroupidSiteidParamFlags registers all flags needed to fill params
func registerOperationAuthorizationDeleteAccessgroupsSitesGroupidSiteidParamFlags(cmd *cobra.Command) error {
	if err := registerOperationAuthorizationDeleteAccessgroupsSitesGroupidSiteidGroupIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationAuthorizationDeleteAccessgroupsSitesGroupidSiteidSiteIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationAuthorizationDeleteAccessgroupsSitesGroupidSiteidGroupIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationAuthorizationDeleteAccessgroupsSitesGroupidSiteidSiteIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	siteIdDescription := `Required. `

	var siteIdFlagName string
	if cmdPrefix == "" {
		siteIdFlagName = "siteId"
	} else {
		siteIdFlagName = fmt.Sprintf("%v.siteId", cmdPrefix)
	}

	var siteIdFlagDefault string

	_ = cmd.PersistentFlags().String(siteIdFlagName, siteIdFlagDefault, siteIdDescription)

	return nil
}

func retrieveOperationAuthorizationDeleteAccessgroupsSitesGroupidSiteidGroupIDFlag(m *authorization.DeleteAccessgroupsSitesGroupidSiteidParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationAuthorizationDeleteAccessgroupsSitesGroupidSiteidSiteIDFlag(m *authorization.DeleteAccessgroupsSitesGroupidSiteidParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("siteId") {

		var siteIdFlagName string
		if cmdPrefix == "" {
			siteIdFlagName = "siteId"
		} else {
			siteIdFlagName = fmt.Sprintf("%v.siteId", cmdPrefix)
		}

		siteIdFlagValue, err := cmd.Flags().GetString(siteIdFlagName)
		if err != nil {
			return err, false
		}
		m.SiteID = siteIdFlagValue

	}
	return nil, retAdded
}

// parseOperationAuthorizationDeleteAccessgroupsSitesGroupidSiteidResult parses request result and return the string content
func parseOperationAuthorizationDeleteAccessgroupsSitesGroupidSiteidResult(resp0 *authorization.DeleteAccessgroupsSitesGroupidSiteidOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*authorization.DeleteAccessgroupsSitesGroupidSiteidOK)
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
		resp1, ok := iResp1.(*authorization.DeleteAccessgroupsSitesGroupidSiteidBadRequest)
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
		resp2, ok := iResp2.(*authorization.DeleteAccessgroupsSitesGroupidSiteidUnauthorized)
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
		resp3, ok := iResp3.(*authorization.DeleteAccessgroupsSitesGroupidSiteidForbidden)
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
		resp4, ok := iResp4.(*authorization.DeleteAccessgroupsSitesGroupidSiteidNotFound)
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
		resp5, ok := iResp5.(*authorization.DeleteAccessgroupsSitesGroupidSiteidConflict)
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
		resp6, ok := iResp6.(*authorization.DeleteAccessgroupsSitesGroupidSiteidInternalServerError)
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
