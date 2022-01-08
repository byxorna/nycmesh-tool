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

// makeOperationDevicesGetDevicesCmd returns a cmd to handle operation getDevices
func makeOperationDevicesGetDevicesCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getDevices",
		Short: ``,
		RunE:  runOperationDevicesGetDevices,
	}

	if err := registerOperationDevicesGetDevicesParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDevicesGetDevices uses cmd flags to call endpoint api
func runOperationDevicesGetDevices(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := devices.NewGetDevicesParams()
	if err, _ := retrieveOperationDevicesGetDevicesAuthorizedFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesGetDevicesRoleFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesGetDevicesSiteIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesGetDevicesTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesGetDevicesWithInterfacesFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDevicesGetDevicesResult(appCli.Devices.GetDevices(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDevicesGetDevicesParamFlags registers all flags needed to fill params
func registerOperationDevicesGetDevicesParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDevicesGetDevicesAuthorizedParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesGetDevicesRoleParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesGetDevicesSiteIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesGetDevicesTypeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesGetDevicesWithInterfacesParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDevicesGetDevicesAuthorizedParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	authorizedDescription := ``

	var authorizedFlagName string
	if cmdPrefix == "" {
		authorizedFlagName = "authorized"
	} else {
		authorizedFlagName = fmt.Sprintf("%v.authorized", cmdPrefix)
	}

	var authorizedFlagDefault bool

	_ = cmd.PersistentFlags().Bool(authorizedFlagName, authorizedFlagDefault, authorizedDescription)

	return nil
}
func registerOperationDevicesGetDevicesRoleParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	roleDescription := ``

	var roleFlagName string
	if cmdPrefix == "" {
		roleFlagName = "role"
	} else {
		roleFlagName = fmt.Sprintf("%v.role", cmdPrefix)
	}

	var roleFlagDefault []string

	_ = cmd.PersistentFlags().StringSlice(roleFlagName, roleFlagDefault, roleDescription)

	return nil
}
func registerOperationDevicesGetDevicesSiteIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	siteIdDescription := ``

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
func registerOperationDevicesGetDevicesTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	typeDescription := ``

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault []string

	_ = cmd.PersistentFlags().StringSlice(typeFlagName, typeFlagDefault, typeDescription)

	return nil
}
func registerOperationDevicesGetDevicesWithInterfacesParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	withInterfacesDescription := ``

	var withInterfacesFlagName string
	if cmdPrefix == "" {
		withInterfacesFlagName = "withInterfaces"
	} else {
		withInterfacesFlagName = fmt.Sprintf("%v.withInterfaces", cmdPrefix)
	}

	var withInterfacesFlagDefault bool

	_ = cmd.PersistentFlags().Bool(withInterfacesFlagName, withInterfacesFlagDefault, withInterfacesDescription)

	return nil
}

func retrieveOperationDevicesGetDevicesAuthorizedFlag(m *devices.GetDevicesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("authorized") {

		var authorizedFlagName string
		if cmdPrefix == "" {
			authorizedFlagName = "authorized"
		} else {
			authorizedFlagName = fmt.Sprintf("%v.authorized", cmdPrefix)
		}

		authorizedFlagValue, err := cmd.Flags().GetBool(authorizedFlagName)
		if err != nil {
			return err, false
		}
		m.Authorized = &authorizedFlagValue

	}
	return nil, retAdded
}
func retrieveOperationDevicesGetDevicesRoleFlag(m *devices.GetDevicesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("role") {

		var roleFlagName string
		if cmdPrefix == "" {
			roleFlagName = "role"
		} else {
			roleFlagName = fmt.Sprintf("%v.role", cmdPrefix)
		}

		roleFlagValues, err := cmd.Flags().GetStringSlice(roleFlagName)
		if err != nil {
			return err, false
		}
		m.Role = roleFlagValues

	}
	return nil, retAdded
}
func retrieveOperationDevicesGetDevicesSiteIDFlag(m *devices.GetDevicesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDevicesGetDevicesTypeFlag(m *devices.GetDevicesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("type") {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValues, err := cmd.Flags().GetStringSlice(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = typeFlagValues

	}
	return nil, retAdded
}
func retrieveOperationDevicesGetDevicesWithInterfacesFlag(m *devices.GetDevicesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("withInterfaces") {

		var withInterfacesFlagName string
		if cmdPrefix == "" {
			withInterfacesFlagName = "withInterfaces"
		} else {
			withInterfacesFlagName = fmt.Sprintf("%v.withInterfaces", cmdPrefix)
		}

		withInterfacesFlagValue, err := cmd.Flags().GetBool(withInterfacesFlagName)
		if err != nil {
			return err, false
		}
		m.WithInterfaces = &withInterfacesFlagValue

	}
	return nil, retAdded
}

// parseOperationDevicesGetDevicesResult parses request result and return the string content
func parseOperationDevicesGetDevicesResult(resp0 *devices.GetDevicesOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*devices.GetDevicesOK)
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
		resp1, ok := iResp1.(*devices.GetDevicesBadRequest)
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
		resp2, ok := iResp2.(*devices.GetDevicesUnauthorized)
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
		resp3, ok := iResp3.(*devices.GetDevicesForbidden)
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
		resp4, ok := iResp4.(*devices.GetDevicesInternalServerError)
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
