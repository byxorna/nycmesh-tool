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

// makeOperationSitesGetSitesIDCmd returns a cmd to handle operation getSitesId
func makeOperationSitesGetSitesIDCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getSitesId",
		Short: ``,
		RunE:  runOperationSitesGetSitesID,
	}

	if err := registerOperationSitesGetSitesIDParamFlags_2(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSitesGetSitesID uses cmd flags to call endpoint api
func runOperationSitesGetSitesID(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := sites.NewGetSitesIDParams()
	if err, _ := retrieveOperationSitesGetSitesIDIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSitesGetSitesIDUcrmDetailsFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSitesGetSitesIDResult(appCli.Sites.GetSitesID(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSitesGetSitesIDParamFlags_2 registers all flags needed to fill params
func registerOperationSitesGetSitesIDParamFlags_2(cmd *cobra.Command) error {
	if err := registerOperationSitesGetSitesIDIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSitesGetSitesIDUcrmDetailsParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSitesGetSitesIDIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. `

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}
func registerOperationSitesGetSitesIDUcrmDetailsParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	ucrmDetailsDescription := `Include CRM client and service details.`

	var ucrmDetailsFlagName string
	if cmdPrefix == "" {
		ucrmDetailsFlagName = "ucrmDetails"
	} else {
		ucrmDetailsFlagName = fmt.Sprintf("%v.ucrmDetails", cmdPrefix)
	}

	var ucrmDetailsFlagDefault bool

	_ = cmd.PersistentFlags().Bool(ucrmDetailsFlagName, ucrmDetailsFlagDefault, ucrmDetailsDescription)

	return nil
}

func retrieveOperationSitesGetSitesIDIDFlag(m *sites.GetSitesIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("id") {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

	}
	return nil, retAdded
}
func retrieveOperationSitesGetSitesIDUcrmDetailsFlag(m *sites.GetSitesIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("ucrmDetails") {

		var ucrmDetailsFlagName string
		if cmdPrefix == "" {
			ucrmDetailsFlagName = "ucrmDetails"
		} else {
			ucrmDetailsFlagName = fmt.Sprintf("%v.ucrmDetails", cmdPrefix)
		}

		ucrmDetailsFlagValue, err := cmd.Flags().GetBool(ucrmDetailsFlagName)
		if err != nil {
			return err, false
		}
		m.UcrmDetails = &ucrmDetailsFlagValue

	}
	return nil, retAdded
}

// parseOperationSitesGetSitesIDResult parses request result and return the string content
func parseOperationSitesGetSitesIDResult(resp0 *sites.GetSitesIDOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*sites.GetSitesIDOK)
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
		resp1, ok := iResp1.(*sites.GetSitesIDBadRequest)
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
		resp2, ok := iResp2.(*sites.GetSitesIDUnauthorized)
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
		resp3, ok := iResp3.(*sites.GetSitesIDForbidden)
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
		resp4, ok := iResp4.(*sites.GetSitesIDNotFound)
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
		resp5, ok := iResp5.(*sites.GetSitesIDInternalServerError)
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
