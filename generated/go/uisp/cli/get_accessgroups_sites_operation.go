// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/authorization"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationAuthorizationGetAccessgroupsSitesCmd returns a cmd to handle operation getAccessgroupsSites
func makeOperationAuthorizationGetAccessgroupsSitesCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getAccessgroupsSites",
		Short: ``,
		RunE:  runOperationAuthorizationGetAccessgroupsSites,
	}

	if err := registerOperationAuthorizationGetAccessgroupsSitesParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationAuthorizationGetAccessgroupsSites uses cmd flags to call endpoint api
func runOperationAuthorizationGetAccessgroupsSites(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := authorization.NewGetAccessgroupsSitesParams()
	if err, _ := retrieveOperationAuthorizationGetAccessgroupsSitesGroupIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationAuthorizationGetAccessgroupsSitesWithInternalFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationAuthorizationGetAccessgroupsSitesWithOverviewFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationAuthorizationGetAccessgroupsSitesResult(appCli.Authorization.GetAccessgroupsSites(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationAuthorizationGetAccessgroupsSitesParamFlags registers all flags needed to fill params
func registerOperationAuthorizationGetAccessgroupsSitesParamFlags(cmd *cobra.Command) error {
	if err := registerOperationAuthorizationGetAccessgroupsSitesGroupIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationAuthorizationGetAccessgroupsSitesWithInternalParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationAuthorizationGetAccessgroupsSitesWithOverviewParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationAuthorizationGetAccessgroupsSitesGroupIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	groupIdDescription := ``

	var groupIdFlagName string
	if cmdPrefix == "" {
		groupIdFlagName = "groupId"
	} else {
		groupIdFlagName = fmt.Sprintf("%v.groupId", cmdPrefix)
	}

	var groupIdFlagDefault []string

	_ = cmd.PersistentFlags().StringSlice(groupIdFlagName, groupIdFlagDefault, groupIdDescription)

	return nil
}
func registerOperationAuthorizationGetAccessgroupsSitesWithInternalParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	withInternalDescription := `Whether to include internal access groups.`

	var withInternalFlagName string
	if cmdPrefix == "" {
		withInternalFlagName = "withInternal"
	} else {
		withInternalFlagName = fmt.Sprintf("%v.withInternal", cmdPrefix)
	}

	var withInternalFlagDefault bool

	_ = cmd.PersistentFlags().Bool(withInternalFlagName, withInternalFlagDefault, withInternalDescription)

	return nil
}
func registerOperationAuthorizationGetAccessgroupsSitesWithOverviewParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	withOverviewDescription := `Whether to include site and user counts.`

	var withOverviewFlagName string
	if cmdPrefix == "" {
		withOverviewFlagName = "withOverview"
	} else {
		withOverviewFlagName = fmt.Sprintf("%v.withOverview", cmdPrefix)
	}

	var withOverviewFlagDefault bool

	_ = cmd.PersistentFlags().Bool(withOverviewFlagName, withOverviewFlagDefault, withOverviewDescription)

	return nil
}

func retrieveOperationAuthorizationGetAccessgroupsSitesGroupIDFlag(m *authorization.GetAccessgroupsSitesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("groupId") {

		var groupIdFlagName string
		if cmdPrefix == "" {
			groupIdFlagName = "groupId"
		} else {
			groupIdFlagName = fmt.Sprintf("%v.groupId", cmdPrefix)
		}

		groupIdFlagValues, err := cmd.Flags().GetStringSlice(groupIdFlagName)
		if err != nil {
			return err, false
		}
		m.GroupID = groupIdFlagValues

	}
	return nil, retAdded
}
func retrieveOperationAuthorizationGetAccessgroupsSitesWithInternalFlag(m *authorization.GetAccessgroupsSitesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("withInternal") {

		var withInternalFlagName string
		if cmdPrefix == "" {
			withInternalFlagName = "withInternal"
		} else {
			withInternalFlagName = fmt.Sprintf("%v.withInternal", cmdPrefix)
		}

		withInternalFlagValue, err := cmd.Flags().GetBool(withInternalFlagName)
		if err != nil {
			return err, false
		}
		m.WithInternal = &withInternalFlagValue

	}
	return nil, retAdded
}
func retrieveOperationAuthorizationGetAccessgroupsSitesWithOverviewFlag(m *authorization.GetAccessgroupsSitesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("withOverview") {

		var withOverviewFlagName string
		if cmdPrefix == "" {
			withOverviewFlagName = "withOverview"
		} else {
			withOverviewFlagName = fmt.Sprintf("%v.withOverview", cmdPrefix)
		}

		withOverviewFlagValue, err := cmd.Flags().GetBool(withOverviewFlagName)
		if err != nil {
			return err, false
		}
		m.WithOverview = &withOverviewFlagValue

	}
	return nil, retAdded
}

// parseOperationAuthorizationGetAccessgroupsSitesResult parses request result and return the string content
func parseOperationAuthorizationGetAccessgroupsSitesResult(resp0 *authorization.GetAccessgroupsSitesOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*authorization.GetAccessgroupsSitesOK)
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
		resp1, ok := iResp1.(*authorization.GetAccessgroupsSitesUnauthorized)
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
		resp2, ok := iResp2.(*authorization.GetAccessgroupsSitesForbidden)
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
		resp3, ok := iResp3.(*authorization.GetAccessgroupsSitesInternalServerError)
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
