// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/sites"
	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSitesPostSitesSiteidUcrmBindCmd returns a cmd to handle operation postSitesSiteidUcrmBind
func makeOperationSitesPostSitesSiteidUcrmBindCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "postSitesSiteidUcrmBind",
		Short: ``,
		RunE:  runOperationSitesPostSitesSiteidUcrmBind,
	}

	if err := registerOperationSitesPostSitesSiteidUcrmBindParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSitesPostSitesSiteidUcrmBind uses cmd flags to call endpoint api
func runOperationSitesPostSitesSiteidUcrmBind(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := sites.NewPostSitesSiteidUcrmBindParams()
	if err, _ := retrieveOperationSitesPostSitesSiteidUcrmBindBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSitesPostSitesSiteidUcrmBindSiteIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSitesPostSitesSiteidUcrmBindResult(appCli.Sites.PostSitesSiteidUcrmBind(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSitesPostSitesSiteidUcrmBindParamFlags registers all flags needed to fill params
func registerOperationSitesPostSitesSiteidUcrmBindParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSitesPostSitesSiteidUcrmBindBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSitesPostSitesSiteidUcrmBindSiteIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSitesPostSitesSiteidUcrmBindBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelSiteBindingFlags(0, "siteBinding", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationSitesPostSitesSiteidUcrmBindSiteIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationSitesPostSitesSiteidUcrmBindBodyFlag(m *sites.PostSitesSiteidUcrmBindParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.SiteBinding{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.SiteBinding: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.SiteBinding{}
	}
	err, added := retrieveModelSiteBindingFlags(0, bodyValueModel, "siteBinding", cmd)
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
func retrieveOperationSitesPostSitesSiteidUcrmBindSiteIDFlag(m *sites.PostSitesSiteidUcrmBindParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationSitesPostSitesSiteidUcrmBindResult parses request result and return the string content
func parseOperationSitesPostSitesSiteidUcrmBindResult(resp0 *sites.PostSitesSiteidUcrmBindOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*sites.PostSitesSiteidUcrmBindOK)
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
		resp1, ok := iResp1.(*sites.PostSitesSiteidUcrmBindBadRequest)
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
		resp2, ok := iResp2.(*sites.PostSitesSiteidUcrmBindUnauthorized)
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
		resp3, ok := iResp3.(*sites.PostSitesSiteidUcrmBindForbidden)
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
		resp4, ok := iResp4.(*sites.PostSitesSiteidUcrmBindInternalServerError)
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
