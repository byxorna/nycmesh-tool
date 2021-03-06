// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/speed_test"
	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSpeedTestPostSpeedtestsStartCmd returns a cmd to handle operation postSpeedtestsStart
func makeOperationSpeedTestPostSpeedtestsStartCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "postSpeedtestsStart",
		Short: ``,
		RunE:  runOperationSpeedTestPostSpeedtestsStart,
	}

	if err := registerOperationSpeedTestPostSpeedtestsStartParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSpeedTestPostSpeedtestsStart uses cmd flags to call endpoint api
func runOperationSpeedTestPostSpeedtestsStart(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := speed_test.NewPostSpeedtestsStartParams()
	if err, _ := retrieveOperationSpeedTestPostSpeedtestsStartBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSpeedTestPostSpeedtestsStartResult(appCli.SpeedTest.PostSpeedtestsStart(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSpeedTestPostSpeedtestsStartParamFlags registers all flags needed to fill params
func registerOperationSpeedTestPostSpeedtestsStartParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSpeedTestPostSpeedtestsStartBodyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSpeedTestPostSpeedtestsStartBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelPayloadSpeedTestFlags(0, "payloadSpeedTest", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationSpeedTestPostSpeedtestsStartBodyFlag(m *speed_test.PostSpeedtestsStartParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.PayloadSpeedTest{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.PayloadSpeedTest: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.PayloadSpeedTest{}
	}
	err, added := retrieveModelPayloadSpeedTestFlags(0, bodyValueModel, "payloadSpeedTest", cmd)
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

// parseOperationSpeedTestPostSpeedtestsStartResult parses request result and return the string content
func parseOperationSpeedTestPostSpeedtestsStartResult(resp0 *speed_test.PostSpeedtestsStartOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*speed_test.PostSpeedtestsStartOK)
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
		resp1, ok := iResp1.(*speed_test.PostSpeedtestsStartBadRequest)
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
		resp2, ok := iResp2.(*speed_test.PostSpeedtestsStartUnauthorized)
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
		resp3, ok := iResp3.(*speed_test.PostSpeedtestsStartForbidden)
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
		resp4, ok := iResp4.(*speed_test.PostSpeedtestsStartInternalServerError)
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
