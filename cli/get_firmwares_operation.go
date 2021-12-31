// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/client/firmware"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationFirmwareGetFirmwaresCmd returns a cmd to handle operation getFirmwares
func makeOperationFirmwareGetFirmwaresCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getFirmwares",
		Short: ``,
		RunE:  runOperationFirmwareGetFirmwares,
	}

	if err := registerOperationFirmwareGetFirmwaresParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationFirmwareGetFirmwares uses cmd flags to call endpoint api
func runOperationFirmwareGetFirmwares(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := firmware.NewGetFirmwaresParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationFirmwareGetFirmwaresResult(appCli.Firmware.GetFirmwares(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationFirmwareGetFirmwaresParamFlags registers all flags needed to fill params
func registerOperationFirmwareGetFirmwaresParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationFirmwareGetFirmwaresResult parses request result and return the string content
func parseOperationFirmwareGetFirmwaresResult(resp0 *firmware.GetFirmwaresOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*firmware.GetFirmwaresOK)
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
		resp1, ok := iResp1.(*firmware.GetFirmwaresUnauthorized)
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
		resp2, ok := iResp2.(*firmware.GetFirmwaresForbidden)
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
		resp3, ok := iResp3.(*firmware.GetFirmwaresInternalServerError)
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
