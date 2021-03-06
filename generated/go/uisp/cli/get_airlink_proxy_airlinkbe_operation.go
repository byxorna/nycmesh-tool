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

// makeOperationDevicesGetAirlinkProxyAirlinkbeCmd returns a cmd to handle operation getAirlinkProxyAirlinkbe
func makeOperationDevicesGetAirlinkProxyAirlinkbeCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getAirlinkProxyAirlinkbe",
		Short: ``,
		RunE:  runOperationDevicesGetAirlinkProxyAirlinkbe,
	}

	if err := registerOperationDevicesGetAirlinkProxyAirlinkbeParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDevicesGetAirlinkProxyAirlinkbe uses cmd flags to call endpoint api
func runOperationDevicesGetAirlinkProxyAirlinkbe(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := devices.NewGetAirlinkProxyAirlinkbeParams()
	if err, _ := retrieveOperationDevicesGetAirlinkProxyAirlinkbeLat1Flag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesGetAirlinkProxyAirlinkbeLat2Flag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesGetAirlinkProxyAirlinkbeLon1Flag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesGetAirlinkProxyAirlinkbeLon2Flag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesGetAirlinkProxyAirlinkbeSamplesFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDevicesGetAirlinkProxyAirlinkbeResult(appCli.Devices.GetAirlinkProxyAirlinkbe(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDevicesGetAirlinkProxyAirlinkbeParamFlags registers all flags needed to fill params
func registerOperationDevicesGetAirlinkProxyAirlinkbeParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDevicesGetAirlinkProxyAirlinkbeLat1ParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesGetAirlinkProxyAirlinkbeLat2ParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesGetAirlinkProxyAirlinkbeLon1ParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesGetAirlinkProxyAirlinkbeLon2ParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesGetAirlinkProxyAirlinkbeSamplesParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDevicesGetAirlinkProxyAirlinkbeLat1ParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	lat1Description := `Required. Latitude of the first point.`

	var lat1FlagName string
	if cmdPrefix == "" {
		lat1FlagName = "lat1"
	} else {
		lat1FlagName = fmt.Sprintf("%v.lat1", cmdPrefix)
	}

	var lat1FlagDefault float64

	_ = cmd.PersistentFlags().Float64(lat1FlagName, lat1FlagDefault, lat1Description)

	return nil
}
func registerOperationDevicesGetAirlinkProxyAirlinkbeLat2ParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	lat2Description := `Required. Latitude of the second point.`

	var lat2FlagName string
	if cmdPrefix == "" {
		lat2FlagName = "lat2"
	} else {
		lat2FlagName = fmt.Sprintf("%v.lat2", cmdPrefix)
	}

	var lat2FlagDefault float64

	_ = cmd.PersistentFlags().Float64(lat2FlagName, lat2FlagDefault, lat2Description)

	return nil
}
func registerOperationDevicesGetAirlinkProxyAirlinkbeLon1ParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	lon1Description := `Required. Longitude of the first point.`

	var lon1FlagName string
	if cmdPrefix == "" {
		lon1FlagName = "lon1"
	} else {
		lon1FlagName = fmt.Sprintf("%v.lon1", cmdPrefix)
	}

	var lon1FlagDefault float64

	_ = cmd.PersistentFlags().Float64(lon1FlagName, lon1FlagDefault, lon1Description)

	return nil
}
func registerOperationDevicesGetAirlinkProxyAirlinkbeLon2ParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	lon2Description := `Required. Longitude of the second point.`

	var lon2FlagName string
	if cmdPrefix == "" {
		lon2FlagName = "lon2"
	} else {
		lon2FlagName = fmt.Sprintf("%v.lon2", cmdPrefix)
	}

	var lon2FlagDefault float64

	_ = cmd.PersistentFlags().Float64(lon2FlagName, lon2FlagDefault, lon2Description)

	return nil
}
func registerOperationDevicesGetAirlinkProxyAirlinkbeSamplesParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	samplesDescription := `Number of elevations to get.`

	var samplesFlagName string
	if cmdPrefix == "" {
		samplesFlagName = "samples"
	} else {
		samplesFlagName = fmt.Sprintf("%v.samples", cmdPrefix)
	}

	var samplesFlagDefault float64 = 256

	_ = cmd.PersistentFlags().Float64(samplesFlagName, samplesFlagDefault, samplesDescription)

	return nil
}

func retrieveOperationDevicesGetAirlinkProxyAirlinkbeLat1Flag(m *devices.GetAirlinkProxyAirlinkbeParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("lat1") {

		var lat1FlagName string
		if cmdPrefix == "" {
			lat1FlagName = "lat1"
		} else {
			lat1FlagName = fmt.Sprintf("%v.lat1", cmdPrefix)
		}

		lat1FlagValue, err := cmd.Flags().GetFloat64(lat1FlagName)
		if err != nil {
			return err, false
		}
		m.Lat1 = lat1FlagValue

	}
	return nil, retAdded
}
func retrieveOperationDevicesGetAirlinkProxyAirlinkbeLat2Flag(m *devices.GetAirlinkProxyAirlinkbeParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("lat2") {

		var lat2FlagName string
		if cmdPrefix == "" {
			lat2FlagName = "lat2"
		} else {
			lat2FlagName = fmt.Sprintf("%v.lat2", cmdPrefix)
		}

		lat2FlagValue, err := cmd.Flags().GetFloat64(lat2FlagName)
		if err != nil {
			return err, false
		}
		m.Lat2 = lat2FlagValue

	}
	return nil, retAdded
}
func retrieveOperationDevicesGetAirlinkProxyAirlinkbeLon1Flag(m *devices.GetAirlinkProxyAirlinkbeParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("lon1") {

		var lon1FlagName string
		if cmdPrefix == "" {
			lon1FlagName = "lon1"
		} else {
			lon1FlagName = fmt.Sprintf("%v.lon1", cmdPrefix)
		}

		lon1FlagValue, err := cmd.Flags().GetFloat64(lon1FlagName)
		if err != nil {
			return err, false
		}
		m.Lon1 = lon1FlagValue

	}
	return nil, retAdded
}
func retrieveOperationDevicesGetAirlinkProxyAirlinkbeLon2Flag(m *devices.GetAirlinkProxyAirlinkbeParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("lon2") {

		var lon2FlagName string
		if cmdPrefix == "" {
			lon2FlagName = "lon2"
		} else {
			lon2FlagName = fmt.Sprintf("%v.lon2", cmdPrefix)
		}

		lon2FlagValue, err := cmd.Flags().GetFloat64(lon2FlagName)
		if err != nil {
			return err, false
		}
		m.Lon2 = lon2FlagValue

	}
	return nil, retAdded
}
func retrieveOperationDevicesGetAirlinkProxyAirlinkbeSamplesFlag(m *devices.GetAirlinkProxyAirlinkbeParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("samples") {

		var samplesFlagName string
		if cmdPrefix == "" {
			samplesFlagName = "samples"
		} else {
			samplesFlagName = fmt.Sprintf("%v.samples", cmdPrefix)
		}

		samplesFlagValue, err := cmd.Flags().GetFloat64(samplesFlagName)
		if err != nil {
			return err, false
		}
		m.Samples = &samplesFlagValue

	}
	return nil, retAdded
}

// parseOperationDevicesGetAirlinkProxyAirlinkbeResult parses request result and return the string content
func parseOperationDevicesGetAirlinkProxyAirlinkbeResult(resp0 *devices.GetAirlinkProxyAirlinkbeOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*devices.GetAirlinkProxyAirlinkbeOK)
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
		resp1, ok := iResp1.(*devices.GetAirlinkProxyAirlinkbeBadRequest)
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
		resp2, ok := iResp2.(*devices.GetAirlinkProxyAirlinkbeUnauthorized)
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
		resp3, ok := iResp3.(*devices.GetAirlinkProxyAirlinkbeForbidden)
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
		resp4, ok := iResp4.(*devices.GetAirlinkProxyAirlinkbeInternalServerError)
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
		resp5, ok := iResp5.(*devices.GetAirlinkProxyAirlinkbeBadGateway)
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
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
