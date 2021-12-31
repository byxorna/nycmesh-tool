// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/client/devices"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationDevicesPutDevicesUisprsIDFirewallNatsRulesCmd returns a cmd to handle operation putDevicesUisprsIdFirewallNatsRules
func makeOperationDevicesPutDevicesUisprsIDFirewallNatsRulesCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "putDevicesUisprsIdFirewallNatsRules",
		Short: ``,
		RunE:  runOperationDevicesPutDevicesUisprsIDFirewallNatsRules,
	}

	if err := registerOperationDevicesPutDevicesUisprsIDFirewallNatsRulesParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDevicesPutDevicesUisprsIDFirewallNatsRules uses cmd flags to call endpoint api
func runOperationDevicesPutDevicesUisprsIDFirewallNatsRules(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := devices.NewPutDevicesUisprsIDFirewallNatsRulesParams()
	if err, _ := retrieveOperationDevicesPutDevicesUisprsIDFirewallNatsRulesBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesPutDevicesUisprsIDFirewallNatsRulesIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDevicesPutDevicesUisprsIDFirewallNatsRulesResult(appCli.Devices.PutDevicesUisprsIDFirewallNatsRules(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDevicesPutDevicesUisprsIDFirewallNatsRulesParamFlags registers all flags needed to fill params
func registerOperationDevicesPutDevicesUisprsIDFirewallNatsRulesParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDevicesPutDevicesUisprsIDFirewallNatsRulesBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesPutDevicesUisprsIDFirewallNatsRulesIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDevicesPutDevicesUisprsIDFirewallNatsRulesBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	// warning: array body models.Model128 is not supported by go-swagger cli yet

	return nil
}
func registerOperationDevicesPutDevicesUisprsIDFirewallNatsRulesIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationDevicesPutDevicesUisprsIDFirewallNatsRulesBodyFlag(m *devices.PutDevicesUisprsIDFirewallNatsRulesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {

		// warning: array body models.Model128 is not supported by go-swagger cli yet

	}
	return nil, retAdded
}
func retrieveOperationDevicesPutDevicesUisprsIDFirewallNatsRulesIDFlag(m *devices.PutDevicesUisprsIDFirewallNatsRulesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationDevicesPutDevicesUisprsIDFirewallNatsRulesResult parses request result and return the string content
func parseOperationDevicesPutDevicesUisprsIDFirewallNatsRulesResult(resp0 *devices.PutDevicesUisprsIDFirewallNatsRulesOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*devices.PutDevicesUisprsIDFirewallNatsRulesOK)
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
		resp1, ok := iResp1.(*devices.PutDevicesUisprsIDFirewallNatsRulesBadRequest)
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
		resp2, ok := iResp2.(*devices.PutDevicesUisprsIDFirewallNatsRulesUnauthorized)
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
		resp3, ok := iResp3.(*devices.PutDevicesUisprsIDFirewallNatsRulesNotFound)
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
		resp4, ok := iResp4.(*devices.PutDevicesUisprsIDFirewallNatsRulesInternalServerError)
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
