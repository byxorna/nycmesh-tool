// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/logs"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationLogsGetLogsCmd returns a cmd to handle operation getLogs
func makeOperationLogsGetLogsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getLogs",
		Short: ``,
		RunE:  runOperationLogsGetLogs,
	}

	if err := registerOperationLogsGetLogsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationLogsGetLogs uses cmd flags to call endpoint api
func runOperationLogsGetLogs(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := logs.NewGetLogsParams()
	if err, _ := retrieveOperationLogsGetLogsCountFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationLogsGetLogsDeviceIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationLogsGetLogsLevelFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationLogsGetLogsPageFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationLogsGetLogsPeriodFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationLogsGetLogsQueryFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationLogsGetLogsSiteIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationLogsGetLogsTagFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationLogsGetLogsResult(appCli.Logs.GetLogs(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationLogsGetLogsParamFlags registers all flags needed to fill params
func registerOperationLogsGetLogsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationLogsGetLogsCountParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationLogsGetLogsDeviceIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationLogsGetLogsLevelParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationLogsGetLogsPageParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationLogsGetLogsPeriodParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationLogsGetLogsQueryParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationLogsGetLogsSiteIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationLogsGetLogsTagParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationLogsGetLogsCountParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	countDescription := `Required. `

	var countFlagName string
	if cmdPrefix == "" {
		countFlagName = "count"
	} else {
		countFlagName = fmt.Sprintf("%v.count", cmdPrefix)
	}

	var countFlagDefault float64

	_ = cmd.PersistentFlags().Float64(countFlagName, countFlagDefault, countDescription)

	return nil
}
func registerOperationLogsGetLogsDeviceIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	deviceIdDescription := ``

	var deviceIdFlagName string
	if cmdPrefix == "" {
		deviceIdFlagName = "deviceId"
	} else {
		deviceIdFlagName = fmt.Sprintf("%v.deviceId", cmdPrefix)
	}

	var deviceIdFlagDefault []string

	_ = cmd.PersistentFlags().StringSlice(deviceIdFlagName, deviceIdFlagDefault, deviceIdDescription)

	return nil
}
func registerOperationLogsGetLogsLevelParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	levelDescription := `Enum: ["info","warning","error"]. `

	var levelFlagName string
	if cmdPrefix == "" {
		levelFlagName = "level"
	} else {
		levelFlagName = fmt.Sprintf("%v.level", cmdPrefix)
	}

	var levelFlagDefault string

	_ = cmd.PersistentFlags().String(levelFlagName, levelFlagDefault, levelDescription)

	if err := cmd.RegisterFlagCompletionFunc(levelFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["info","warning","error"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}
func registerOperationLogsGetLogsPageParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	pageDescription := `Required. `

	var pageFlagName string
	if cmdPrefix == "" {
		pageFlagName = "page"
	} else {
		pageFlagName = fmt.Sprintf("%v.page", cmdPrefix)
	}

	var pageFlagDefault float64

	_ = cmd.PersistentFlags().Float64(pageFlagName, pageFlagDefault, pageDescription)

	return nil
}
func registerOperationLogsGetLogsPeriodParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	periodDescription := `Unix timestamp in milliseconds.`

	var periodFlagName string
	if cmdPrefix == "" {
		periodFlagName = "period"
	} else {
		periodFlagName = fmt.Sprintf("%v.period", cmdPrefix)
	}

	var periodFlagDefault float64

	_ = cmd.PersistentFlags().Float64(periodFlagName, periodFlagDefault, periodDescription)

	return nil
}
func registerOperationLogsGetLogsQueryParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	queryDescription := `Search pattern.`

	var queryFlagName string
	if cmdPrefix == "" {
		queryFlagName = "query"
	} else {
		queryFlagName = fmt.Sprintf("%v.query", cmdPrefix)
	}

	var queryFlagDefault string

	_ = cmd.PersistentFlags().String(queryFlagName, queryFlagDefault, queryDescription)

	return nil
}
func registerOperationLogsGetLogsSiteIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationLogsGetLogsTagParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	tagDescription := `Enum: ["login","device","email-dispatch","nms-backup","nms-update","nms-error","device-state","device-backup","device-upgrade","device-interface","site"]. `

	var tagFlagName string
	if cmdPrefix == "" {
		tagFlagName = "tag"
	} else {
		tagFlagName = fmt.Sprintf("%v.tag", cmdPrefix)
	}

	var tagFlagDefault string

	_ = cmd.PersistentFlags().String(tagFlagName, tagFlagDefault, tagDescription)

	if err := cmd.RegisterFlagCompletionFunc(tagFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["login","device","email-dispatch","nms-backup","nms-update","nms-error","device-state","device-backup","device-upgrade","device-interface","site"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func retrieveOperationLogsGetLogsCountFlag(m *logs.GetLogsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Count = countFlagValue

	}
	return nil, retAdded
}
func retrieveOperationLogsGetLogsDeviceIDFlag(m *logs.GetLogsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("deviceId") {

		var deviceIdFlagName string
		if cmdPrefix == "" {
			deviceIdFlagName = "deviceId"
		} else {
			deviceIdFlagName = fmt.Sprintf("%v.deviceId", cmdPrefix)
		}

		deviceIdFlagValues, err := cmd.Flags().GetStringSlice(deviceIdFlagName)
		if err != nil {
			return err, false
		}
		m.DeviceID = deviceIdFlagValues

	}
	return nil, retAdded
}
func retrieveOperationLogsGetLogsLevelFlag(m *logs.GetLogsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("level") {

		var levelFlagName string
		if cmdPrefix == "" {
			levelFlagName = "level"
		} else {
			levelFlagName = fmt.Sprintf("%v.level", cmdPrefix)
		}

		levelFlagValue, err := cmd.Flags().GetString(levelFlagName)
		if err != nil {
			return err, false
		}
		m.Level = &levelFlagValue

	}
	return nil, retAdded
}
func retrieveOperationLogsGetLogsPageFlag(m *logs.GetLogsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("page") {

		var pageFlagName string
		if cmdPrefix == "" {
			pageFlagName = "page"
		} else {
			pageFlagName = fmt.Sprintf("%v.page", cmdPrefix)
		}

		pageFlagValue, err := cmd.Flags().GetFloat64(pageFlagName)
		if err != nil {
			return err, false
		}
		m.Page = pageFlagValue

	}
	return nil, retAdded
}
func retrieveOperationLogsGetLogsPeriodFlag(m *logs.GetLogsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("period") {

		var periodFlagName string
		if cmdPrefix == "" {
			periodFlagName = "period"
		} else {
			periodFlagName = fmt.Sprintf("%v.period", cmdPrefix)
		}

		periodFlagValue, err := cmd.Flags().GetFloat64(periodFlagName)
		if err != nil {
			return err, false
		}
		m.Period = &periodFlagValue

	}
	return nil, retAdded
}
func retrieveOperationLogsGetLogsQueryFlag(m *logs.GetLogsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("query") {

		var queryFlagName string
		if cmdPrefix == "" {
			queryFlagName = "query"
		} else {
			queryFlagName = fmt.Sprintf("%v.query", cmdPrefix)
		}

		queryFlagValue, err := cmd.Flags().GetString(queryFlagName)
		if err != nil {
			return err, false
		}
		m.Query = &queryFlagValue

	}
	return nil, retAdded
}
func retrieveOperationLogsGetLogsSiteIDFlag(m *logs.GetLogsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationLogsGetLogsTagFlag(m *logs.GetLogsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("tag") {

		var tagFlagName string
		if cmdPrefix == "" {
			tagFlagName = "tag"
		} else {
			tagFlagName = fmt.Sprintf("%v.tag", cmdPrefix)
		}

		tagFlagValue, err := cmd.Flags().GetString(tagFlagName)
		if err != nil {
			return err, false
		}
		m.Tag = &tagFlagValue

	}
	return nil, retAdded
}

// parseOperationLogsGetLogsResult parses request result and return the string content
func parseOperationLogsGetLogsResult(resp0 *logs.GetLogsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*logs.GetLogsOK)
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
		resp1, ok := iResp1.(*logs.GetLogsBadRequest)
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
		resp2, ok := iResp2.(*logs.GetLogsUnauthorized)
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
		resp3, ok := iResp3.(*logs.GetLogsForbidden)
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
		resp4, ok := iResp4.(*logs.GetLogsInternalServerError)
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
