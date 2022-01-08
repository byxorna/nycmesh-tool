// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/backups"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationBackupsGetNmsDownloadsTokenCmd returns a cmd to handle operation getNmsDownloadsToken
func makeOperationBackupsGetNmsDownloadsTokenCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getNmsDownloadsToken",
		Short: ``,
		RunE:  runOperationBackupsGetNmsDownloadsToken,
	}

	if err := registerOperationBackupsGetNmsDownloadsTokenParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBackupsGetNmsDownloadsToken uses cmd flags to call endpoint api
func runOperationBackupsGetNmsDownloadsToken(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := backups.NewGetNmsDownloadsTokenParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBackupsGetNmsDownloadsTokenResult(appCli.Backups.GetNmsDownloadsToken(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBackupsGetNmsDownloadsTokenParamFlags registers all flags needed to fill params
func registerOperationBackupsGetNmsDownloadsTokenParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationBackupsGetNmsDownloadsTokenResult parses request result and return the string content
func parseOperationBackupsGetNmsDownloadsTokenResult(respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*backups.GetNmsDownloadsTokenBadRequest)
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
		resp1, ok := iResp1.(*backups.GetNmsDownloadsTokenNotFound)
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
		resp2, ok := iResp2.(*backups.GetNmsDownloadsTokenInternalServerError)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}
	return "", nil
}