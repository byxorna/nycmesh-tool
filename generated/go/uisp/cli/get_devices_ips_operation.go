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

// makeOperationDevicesGetDevicesIpsCmd returns a cmd to handle operation getDevicesIps
func makeOperationDevicesGetDevicesIpsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getDevicesIps",
		Short: ``,
		RunE:  runOperationDevicesGetDevicesIps,
	}

	if err := registerOperationDevicesGetDevicesIpsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDevicesGetDevicesIps uses cmd flags to call endpoint api
func runOperationDevicesGetDevicesIps(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := devices.NewGetDevicesIpsParams()
	if err, _ := retrieveOperationDevicesGetDevicesIpsManagementFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesGetDevicesIpsSiteIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesGetDevicesIpsSuspendedFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesGetDevicesIpsUcrmIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDevicesGetDevicesIpsResult(appCli.Devices.GetDevicesIps(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDevicesGetDevicesIpsParamFlags registers all flags needed to fill params
func registerOperationDevicesGetDevicesIpsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDevicesGetDevicesIpsManagementParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesGetDevicesIpsSiteIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesGetDevicesIpsSuspendedParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesGetDevicesIpsUcrmIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDevicesGetDevicesIpsManagementParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	managementDescription := `Return also management IPs.`

	var managementFlagName string
	if cmdPrefix == "" {
		managementFlagName = "management"
	} else {
		managementFlagName = fmt.Sprintf("%v.management", cmdPrefix)
	}

	var managementFlagDefault bool

	_ = cmd.PersistentFlags().Bool(managementFlagName, managementFlagDefault, managementDescription)

	return nil
}
func registerOperationDevicesGetDevicesIpsSiteIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	siteIdDescription := `Return only IPs for this UISP site.`

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
func registerOperationDevicesGetDevicesIpsSuspendedParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	suspendedDescription := `Return only IPs which are blocked.`

	var suspendedFlagName string
	if cmdPrefix == "" {
		suspendedFlagName = "suspended"
	} else {
		suspendedFlagName = fmt.Sprintf("%v.suspended", cmdPrefix)
	}

	var suspendedFlagDefault bool

	_ = cmd.PersistentFlags().Bool(suspendedFlagName, suspendedFlagDefault, suspendedDescription)

	return nil
}
func registerOperationDevicesGetDevicesIpsUcrmIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	ucrmIdDescription := `Return only IPs for this UCRM service.`

	var ucrmIdFlagName string
	if cmdPrefix == "" {
		ucrmIdFlagName = "ucrmId"
	} else {
		ucrmIdFlagName = fmt.Sprintf("%v.ucrmId", cmdPrefix)
	}

	var ucrmIdFlagDefault string

	_ = cmd.PersistentFlags().String(ucrmIdFlagName, ucrmIdFlagDefault, ucrmIdDescription)

	return nil
}

func retrieveOperationDevicesGetDevicesIpsManagementFlag(m *devices.GetDevicesIpsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("management") {

		var managementFlagName string
		if cmdPrefix == "" {
			managementFlagName = "management"
		} else {
			managementFlagName = fmt.Sprintf("%v.management", cmdPrefix)
		}

		managementFlagValue, err := cmd.Flags().GetBool(managementFlagName)
		if err != nil {
			return err, false
		}
		m.Management = &managementFlagValue

	}
	return nil, retAdded
}
func retrieveOperationDevicesGetDevicesIpsSiteIDFlag(m *devices.GetDevicesIpsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.SiteID = &siteIdFlagValue

	}
	return nil, retAdded
}
func retrieveOperationDevicesGetDevicesIpsSuspendedFlag(m *devices.GetDevicesIpsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("suspended") {

		var suspendedFlagName string
		if cmdPrefix == "" {
			suspendedFlagName = "suspended"
		} else {
			suspendedFlagName = fmt.Sprintf("%v.suspended", cmdPrefix)
		}

		suspendedFlagValue, err := cmd.Flags().GetBool(suspendedFlagName)
		if err != nil {
			return err, false
		}
		m.Suspended = &suspendedFlagValue

	}
	return nil, retAdded
}
func retrieveOperationDevicesGetDevicesIpsUcrmIDFlag(m *devices.GetDevicesIpsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("ucrmId") {

		var ucrmIdFlagName string
		if cmdPrefix == "" {
			ucrmIdFlagName = "ucrmId"
		} else {
			ucrmIdFlagName = fmt.Sprintf("%v.ucrmId", cmdPrefix)
		}

		ucrmIdFlagValue, err := cmd.Flags().GetString(ucrmIdFlagName)
		if err != nil {
			return err, false
		}
		m.UcrmID = &ucrmIdFlagValue

	}
	return nil, retAdded
}

// parseOperationDevicesGetDevicesIpsResult parses request result and return the string content
func parseOperationDevicesGetDevicesIpsResult(resp0 *devices.GetDevicesIpsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*devices.GetDevicesIpsOK)
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
		resp1, ok := iResp1.(*devices.GetDevicesIpsBadRequest)
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
		resp2, ok := iResp2.(*devices.GetDevicesIpsUnauthorized)
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
		resp3, ok := iResp3.(*devices.GetDevicesIpsForbidden)
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
		resp4, ok := iResp4.(*devices.GetDevicesIpsInternalServerError)
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
