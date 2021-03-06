// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/devices"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidCmd returns a cmd to handle operation deleteDevicesUisprsIdFirewallFiltersFilternameRuleRuleid
func makeOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteDevicesUisprsIdFirewallFiltersFilternameRuleRuleid",
		Short: ``,
		RunE:  runOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleid,
	}

	if err := registerOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleid uses cmd flags to call endpoint api
func runOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleid(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := devices.NewDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidParams()
	if err, _ := retrieveOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidFilterNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidIsForceApplyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidRuleIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidResult(appCli.Devices.DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleid(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidParamFlags registers all flags needed to fill params
func registerOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidFilterNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidIsForceApplyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidRuleIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidFilterNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidIsForceApplyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidRuleIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	ruleIdDescription := `Required. `

	var ruleIdFlagName string
	if cmdPrefix == "" {
		ruleIdFlagName = "ruleId"
	} else {
		ruleIdFlagName = fmt.Sprintf("%v.ruleId", cmdPrefix)
	}

	var ruleIdFlagDefault string

	_ = cmd.PersistentFlags().String(ruleIdFlagName, ruleIdFlagDefault, ruleIdDescription)

	return nil
}

func retrieveOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidFilterNameFlag(m *devices.DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidIDFlag(m *devices.DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidIsForceApplyFlag(m *devices.DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidRuleIDFlag(m *devices.DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("ruleId") {

		var ruleIdFlagName string
		if cmdPrefix == "" {
			ruleIdFlagName = "ruleId"
		} else {
			ruleIdFlagName = fmt.Sprintf("%v.ruleId", cmdPrefix)
		}

		ruleIdFlagValue, err := cmd.Flags().GetString(ruleIdFlagName)
		if err != nil {
			return err, false
		}
		m.RuleID = ruleIdFlagValue

	}
	return nil, retAdded
}

// parseOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidResult parses request result and return the string content
func parseOperationDevicesDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidResult(resp0 *devices.DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*devices.DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidOK)
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
		resp1, ok := iResp1.(*devices.DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidBadRequest)
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
		resp2, ok := iResp2.(*devices.DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidUnauthorized)
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
		resp3, ok := iResp3.(*devices.DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidNotFound)
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
		resp4, ok := iResp4.(*devices.DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidInternalServerError)
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
