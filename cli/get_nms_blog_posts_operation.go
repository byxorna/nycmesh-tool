// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/client/server"
	serverops "github.com/byxorna/nycmesh-tool/client/server"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationServerGetNmsBlogPostsCmd returns a cmd to handle operation getNmsBlogPosts
func makeOperationServerGetNmsBlogPostsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getNmsBlogPosts",
		Short: ``,
		RunE:  runOperationServerGetNmsBlogPosts,
	}

	if err := registerOperationServerGetNmsBlogPostsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServerGetNmsBlogPosts uses cmd flags to call endpoint api
func runOperationServerGetNmsBlogPosts(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := serverops.NewGetNmsBlogPostsParams()
	if err, _ := retrieveOperationServerGetNmsBlogPostsCountFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServerGetNmsBlogPostsPageFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationServerGetNmsBlogPostsResult(appCli.Server.GetNmsBlogPosts(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationServerGetNmsBlogPostsParamFlags registers all flags needed to fill params
func registerOperationServerGetNmsBlogPostsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationServerGetNmsBlogPostsCountParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServerGetNmsBlogPostsPageParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationServerGetNmsBlogPostsCountParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	countDescription := `items count`

	var countFlagName string
	if cmdPrefix == "" {
		countFlagName = "count"
	} else {
		countFlagName = fmt.Sprintf("%v.count", cmdPrefix)
	}

	var countFlagDefault int64

	_ = cmd.PersistentFlags().Int64(countFlagName, countFlagDefault, countDescription)

	return nil
}
func registerOperationServerGetNmsBlogPostsPageParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	pageDescription := `number page`

	var pageFlagName string
	if cmdPrefix == "" {
		pageFlagName = "page"
	} else {
		pageFlagName = fmt.Sprintf("%v.page", cmdPrefix)
	}

	var pageFlagDefault int64

	_ = cmd.PersistentFlags().Int64(pageFlagName, pageFlagDefault, pageDescription)

	return nil
}

func retrieveOperationServerGetNmsBlogPostsCountFlag(m *serverops.GetNmsBlogPostsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("count") {

		var countFlagName string
		if cmdPrefix == "" {
			countFlagName = "count"
		} else {
			countFlagName = fmt.Sprintf("%v.count", cmdPrefix)
		}

		countFlagValue, err := cmd.Flags().GetInt64(countFlagName)
		if err != nil {
			return err, false
		}
		m.Count = &countFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServerGetNmsBlogPostsPageFlag(m *serverops.GetNmsBlogPostsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("page") {

		var pageFlagName string
		if cmdPrefix == "" {
			pageFlagName = "page"
		} else {
			pageFlagName = fmt.Sprintf("%v.page", cmdPrefix)
		}

		pageFlagValue, err := cmd.Flags().GetInt64(pageFlagName)
		if err != nil {
			return err, false
		}
		m.Page = &pageFlagValue

	}
	return nil, retAdded
}

// parseOperationServerGetNmsBlogPostsResult parses request result and return the string content
func parseOperationServerGetNmsBlogPostsResult(resp0 *server.GetNmsBlogPostsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*server.GetNmsBlogPostsOK)
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
		resp1, ok := iResp1.(*server.GetNmsBlogPostsBadRequest)
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
		resp2, ok := iResp2.(*server.GetNmsBlogPostsUnauthorized)
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
		resp3, ok := iResp3.(*server.GetNmsBlogPostsForbidden)
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
		resp4, ok := iResp4.(*server.GetNmsBlogPostsInternalServerError)
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
