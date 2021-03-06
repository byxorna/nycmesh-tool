// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/devices"
	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleCmd returns a cmd to handle operation putDevicesUisprsIdFirewallFiltersFilternameRule
func makeOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "putDevicesUisprsIdFirewallFiltersFilternameRule",
		Short: ``,
		RunE:  runOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRule,
	}

	if err := registerOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRule uses cmd flags to call endpoint api
func runOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRule(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := devices.NewPutDevicesUisprsIDFirewallFiltersFilternameRuleParams()
	if err, _ := retrieveOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleFilterNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleIsForceApplyFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleResult(appCli.Devices.PutDevicesUisprsIDFirewallFiltersFilternameRule(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleParamFlags registers all flags needed to fill params
func registerOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleFilterNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleIsForceApplyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelModel110Flags(0, "model110", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleFilterNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterNameDescription := `Required. `

	var filterNameFlagName string
	if cmdPrefix == "" {
		filterNameFlagName = "filterName"
	} else {
		filterNameFlagName = fmt.Sprintf("%v.filterName", cmdPrefix)
	}

	var filterNameFlagDefault string

	_ = cmd.PersistentFlags().String(filterNameFlagName, filterNameFlagDefault, filterNameDescription)

	return nil
}
func registerOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleIsForceApplyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	isForceApplyDescription := ``

	var isForceApplyFlagName string
	if cmdPrefix == "" {
		isForceApplyFlagName = "isForceApply"
	} else {
		isForceApplyFlagName = fmt.Sprintf("%v.isForceApply", cmdPrefix)
	}

	var isForceApplyFlagDefault string

	_ = cmd.PersistentFlags().String(isForceApplyFlagName, isForceApplyFlagDefault, isForceApplyDescription)

	return nil
}

func retrieveOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleBodyFlag(m *devices.PutDevicesUisprsIDFirewallFiltersFilternameRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.Model110{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.Model110: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.Model110{}
	}
	err, added := retrieveModelModel110Flags(0, bodyValueModel, "model110", cmd)
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
func retrieveOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleFilterNameFlag(m *devices.PutDevicesUisprsIDFirewallFiltersFilternameRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filterName") {

		var filterNameFlagName string
		if cmdPrefix == "" {
			filterNameFlagName = "filterName"
		} else {
			filterNameFlagName = fmt.Sprintf("%v.filterName", cmdPrefix)
		}

		filterNameFlagValue, err := cmd.Flags().GetString(filterNameFlagName)
		if err != nil {
			return err, false
		}
		m.FilterName = filterNameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleIDFlag(m *devices.PutDevicesUisprsIDFirewallFiltersFilternameRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleIsForceApplyFlag(m *devices.PutDevicesUisprsIDFirewallFiltersFilternameRuleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("isForceApply") {

		var isForceApplyFlagName string
		if cmdPrefix == "" {
			isForceApplyFlagName = "isForceApply"
		} else {
			isForceApplyFlagName = fmt.Sprintf("%v.isForceApply", cmdPrefix)
		}

		isForceApplyFlagValue, err := cmd.Flags().GetString(isForceApplyFlagName)
		if err != nil {
			return err, false
		}
		m.IsForceApply = &isForceApplyFlagValue

	}
	return nil, retAdded
}

// parseOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleResult parses request result and return the string content
func parseOperationDevicesPutDevicesUisprsIDFirewallFiltersFilternameRuleResult(resp0 *devices.PutDevicesUisprsIDFirewallFiltersFilternameRuleOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*devices.PutDevicesUisprsIDFirewallFiltersFilternameRuleOK)
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
		resp1, ok := iResp1.(*devices.PutDevicesUisprsIDFirewallFiltersFilternameRuleBadRequest)
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
		resp2, ok := iResp2.(*devices.PutDevicesUisprsIDFirewallFiltersFilternameRuleUnauthorized)
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
		resp3, ok := iResp3.(*devices.PutDevicesUisprsIDFirewallFiltersFilternameRuleNotFound)
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
		resp4, ok := iResp4.(*devices.PutDevicesUisprsIDFirewallFiltersFilternameRuleInternalServerError)
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
