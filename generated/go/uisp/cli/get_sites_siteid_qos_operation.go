// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/sites"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSitesGetSitesSiteidQosCmd returns a cmd to handle operation getSitesSiteidQos
func makeOperationSitesGetSitesSiteidQosCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getSitesSiteidQos",
		Short: ``,
		RunE:  runOperationSitesGetSitesSiteidQos,
	}

	if err := registerOperationSitesGetSitesSiteidQosParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSitesGetSitesSiteidQos uses cmd flags to call endpoint api
func runOperationSitesGetSitesSiteidQos(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := sites.NewGetSitesSiteidQosParams()
	if err, _ := retrieveOperationSitesGetSitesSiteidQosSiteIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSitesGetSitesSiteidQosResult(appCli.Sites.GetSitesSiteidQos(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSitesGetSitesSiteidQosParamFlags registers all flags needed to fill params
func registerOperationSitesGetSitesSiteidQosParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSitesGetSitesSiteidQosSiteIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSitesGetSitesSiteidQosSiteIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationSitesGetSitesSiteidQosSiteIDFlag(m *sites.GetSitesSiteidQosParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationSitesGetSitesSiteidQosResult parses request result and return the string content
func parseOperationSitesGetSitesSiteidQosResult(resp0 *sites.GetSitesSiteidQosOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*sites.GetSitesSiteidQosOK)
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
		resp1, ok := iResp1.(*sites.GetSitesSiteidQosBadRequest)
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
		resp2, ok := iResp2.(*sites.GetSitesSiteidQosUnauthorized)
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
		resp3, ok := iResp3.(*sites.GetSitesSiteidQosForbidden)
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
		resp4, ok := iResp4.(*sites.GetSitesSiteidQosNotFound)
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
		resp5, ok := iResp5.(*sites.GetSitesSiteidQosInternalServerError)
		if ok {
			if !swag.IsZero(resp5) && !swag.IsZero(resp5.Payload) {
				msgStr, err := json.Marshal(resp5.Payload)
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
