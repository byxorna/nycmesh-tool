// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/data_links"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationDataLinksDeleteDatalinksIDCmd returns a cmd to handle operation deleteDatalinksId
func makeOperationDataLinksDeleteDatalinksIDCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteDatalinksId",
		Short: ``,
		RunE:  runOperationDataLinksDeleteDatalinksID,
	}

	if err := registerOperationDataLinksDeleteDatalinksIDParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDataLinksDeleteDatalinksID uses cmd flags to call endpoint api
func runOperationDataLinksDeleteDatalinksID(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := data_links.NewDeleteDatalinksIDParams()
	if err, _ := retrieveOperationDataLinksDeleteDatalinksIDIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDataLinksDeleteDatalinksIDResult(appCli.DataLinks.DeleteDatalinksID(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDataLinksDeleteDatalinksIDParamFlags registers all flags needed to fill params
func registerOperationDataLinksDeleteDatalinksIDParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDataLinksDeleteDatalinksIDIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDataLinksDeleteDatalinksIDIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationDataLinksDeleteDatalinksIDIDFlag(m *data_links.DeleteDatalinksIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationDataLinksDeleteDatalinksIDResult parses request result and return the string content
func parseOperationDataLinksDeleteDatalinksIDResult(resp0 *data_links.DeleteDatalinksIDOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*data_links.DeleteDatalinksIDOK)
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
		resp1, ok := iResp1.(*data_links.DeleteDatalinksIDBadRequest)
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
		resp2, ok := iResp2.(*data_links.DeleteDatalinksIDUnauthorized)
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
		resp3, ok := iResp3.(*data_links.DeleteDatalinksIDForbidden)
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
		resp4, ok := iResp4.(*data_links.DeleteDatalinksIDNotFound)
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
		resp5, ok := iResp5.(*data_links.DeleteDatalinksIDInternalServerError)
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
		msgStr := fmt.Sprintf("%v", resp0.Payload)
		return string(msgStr), nil
	}

	return "", nil
}
