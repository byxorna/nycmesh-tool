// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/installations"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationInstallationsGetInstallationsDevicesCmd returns a cmd to handle operation getInstallationsDevices
func makeOperationInstallationsGetInstallationsDevicesCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getInstallationsDevices",
		Short: ``,
		RunE:  runOperationInstallationsGetInstallationsDevices,
	}

	if err := registerOperationInstallationsGetInstallationsDevicesParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationInstallationsGetInstallationsDevices uses cmd flags to call endpoint api
func runOperationInstallationsGetInstallationsDevices(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := installations.NewGetInstallationsDevicesParams()
	if err, _ := retrieveOperationInstallationsGetInstallationsDevicesCountFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationInstallationsGetInstallationsDevicesLatitudeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationInstallationsGetInstallationsDevicesLongitudeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationInstallationsGetInstallationsDevicesUserIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationInstallationsGetInstallationsDevicesResult(appCli.Installations.GetInstallationsDevices(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationInstallationsGetInstallationsDevicesParamFlags registers all flags needed to fill params
func registerOperationInstallationsGetInstallationsDevicesParamFlags(cmd *cobra.Command) error {
	if err := registerOperationInstallationsGetInstallationsDevicesCountParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationInstallationsGetInstallationsDevicesLatitudeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationInstallationsGetInstallationsDevicesLongitudeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationInstallationsGetInstallationsDevicesUserIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationInstallationsGetInstallationsDevicesCountParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	countDescription := `How many devices to get.`

	var countFlagName string
	if cmdPrefix == "" {
		countFlagName = "count"
	} else {
		countFlagName = fmt.Sprintf("%v.count", cmdPrefix)
	}

	var countFlagDefault float64 = 100

	_ = cmd.PersistentFlags().Float64(countFlagName, countFlagDefault, countDescription)

	return nil
}
func registerOperationInstallationsGetInstallationsDevicesLatitudeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	latitudeDescription := ``

	var latitudeFlagName string
	if cmdPrefix == "" {
		latitudeFlagName = "latitude"
	} else {
		latitudeFlagName = fmt.Sprintf("%v.latitude", cmdPrefix)
	}

	var latitudeFlagDefault float64

	_ = cmd.PersistentFlags().Float64(latitudeFlagName, latitudeFlagDefault, latitudeDescription)

	return nil
}
func registerOperationInstallationsGetInstallationsDevicesLongitudeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	longitudeDescription := ``

	var longitudeFlagName string
	if cmdPrefix == "" {
		longitudeFlagName = "longitude"
	} else {
		longitudeFlagName = fmt.Sprintf("%v.longitude", cmdPrefix)
	}

	var longitudeFlagDefault float64

	_ = cmd.PersistentFlags().Float64(longitudeFlagName, longitudeFlagDefault, longitudeDescription)

	return nil
}
func registerOperationInstallationsGetInstallationsDevicesUserIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	userIdDescription := `If specified, get only devices registered by given users.`

	var userIdFlagName string
	if cmdPrefix == "" {
		userIdFlagName = "userId"
	} else {
		userIdFlagName = fmt.Sprintf("%v.userId", cmdPrefix)
	}

	var userIdFlagDefault []string

	_ = cmd.PersistentFlags().StringSlice(userIdFlagName, userIdFlagDefault, userIdDescription)

	return nil
}

func retrieveOperationInstallationsGetInstallationsDevicesCountFlag(m *installations.GetInstallationsDevicesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("count") {

		var countFlagName string
		if cmdPrefix == "" {
			countFlagName = "count"
		} else {
			countFlagName = fmt.Sprintf("%v.count", cmdPrefix)
		}

		countFlagValue, err := cmd.Flags().GetFloat64(countFlagName)
		if err != nil {
			return err, false
		}
		m.Count = &countFlagValue

	}
	return nil, retAdded
}
func retrieveOperationInstallationsGetInstallationsDevicesLatitudeFlag(m *installations.GetInstallationsDevicesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("latitude") {

		var latitudeFlagName string
		if cmdPrefix == "" {
			latitudeFlagName = "latitude"
		} else {
			latitudeFlagName = fmt.Sprintf("%v.latitude", cmdPrefix)
		}

		latitudeFlagValue, err := cmd.Flags().GetFloat64(latitudeFlagName)
		if err != nil {
			return err, false
		}
		m.Latitude = &latitudeFlagValue

	}
	return nil, retAdded
}
func retrieveOperationInstallationsGetInstallationsDevicesLongitudeFlag(m *installations.GetInstallationsDevicesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("longitude") {

		var longitudeFlagName string
		if cmdPrefix == "" {
			longitudeFlagName = "longitude"
		} else {
			longitudeFlagName = fmt.Sprintf("%v.longitude", cmdPrefix)
		}

		longitudeFlagValue, err := cmd.Flags().GetFloat64(longitudeFlagName)
		if err != nil {
			return err, false
		}
		m.Longitude = &longitudeFlagValue

	}
	return nil, retAdded
}
func retrieveOperationInstallationsGetInstallationsDevicesUserIDFlag(m *installations.GetInstallationsDevicesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("userId") {

		var userIdFlagName string
		if cmdPrefix == "" {
			userIdFlagName = "userId"
		} else {
			userIdFlagName = fmt.Sprintf("%v.userId", cmdPrefix)
		}

		userIdFlagValues, err := cmd.Flags().GetStringSlice(userIdFlagName)
		if err != nil {
			return err, false
		}
		m.UserID = userIdFlagValues

	}
	return nil, retAdded
}

// parseOperationInstallationsGetInstallationsDevicesResult parses request result and return the string content
func parseOperationInstallationsGetInstallationsDevicesResult(resp0 *installations.GetInstallationsDevicesOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*installations.GetInstallationsDevicesOK)
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
		resp1, ok := iResp1.(*installations.GetInstallationsDevicesUnauthorized)
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
		resp2, ok := iResp2.(*installations.GetInstallationsDevicesForbidden)
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
		resp3, ok := iResp3.(*installations.GetInstallationsDevicesInternalServerError)
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
		resp4, ok := iResp4.(*installations.GetInstallationsDevicesServiceUnavailable)
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
