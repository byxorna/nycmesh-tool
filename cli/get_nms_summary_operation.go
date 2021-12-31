// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/client/server"
	serverops "github.com/byxorna/nycmesh-tool/client/server"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationServerGetNmsSummaryCmd returns a cmd to handle operation getNmsSummary
func makeOperationServerGetNmsSummaryCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getNmsSummary",
		Short: ``,
		RunE:  runOperationServerGetNmsSummary,
	}

	if err := registerOperationServerGetNmsSummaryParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServerGetNmsSummary uses cmd flags to call endpoint api
func runOperationServerGetNmsSummary(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := serverops.NewGetNmsSummaryParams()
	if err, _ := retrieveOperationServerGetNmsSummaryFirmwaresTimestampFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServerGetNmsSummaryLogsLevelFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServerGetNmsSummaryLogsTimestampFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServerGetNmsSummaryOutagesTimestampFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationServerGetNmsSummaryResult(appCli.Server.GetNmsSummary(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationServerGetNmsSummaryParamFlags registers all flags needed to fill params
func registerOperationServerGetNmsSummaryParamFlags(cmd *cobra.Command) error {
	if err := registerOperationServerGetNmsSummaryFirmwaresTimestampParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServerGetNmsSummaryLogsLevelParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServerGetNmsSummaryLogsTimestampParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServerGetNmsSummaryOutagesTimestampParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationServerGetNmsSummaryFirmwaresTimestampParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	firmwaresTimestampDescription := ``

	var firmwaresTimestampFlagName string
	if cmdPrefix == "" {
		firmwaresTimestampFlagName = "firmwaresTimestamp"
	} else {
		firmwaresTimestampFlagName = fmt.Sprintf("%v.firmwaresTimestamp", cmdPrefix)
	}

	var firmwaresTimestampFlagDefault float64

	_ = cmd.PersistentFlags().Float64(firmwaresTimestampFlagName, firmwaresTimestampFlagDefault, firmwaresTimestampDescription)

	return nil
}
func registerOperationServerGetNmsSummaryLogsLevelParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	logsLevelDescription := `Required. `

	var logsLevelFlagName string
	if cmdPrefix == "" {
		logsLevelFlagName = "logsLevel"
	} else {
		logsLevelFlagName = fmt.Sprintf("%v.logsLevel", cmdPrefix)
	}

	var logsLevelFlagDefault []string

	_ = cmd.PersistentFlags().StringSlice(logsLevelFlagName, logsLevelFlagDefault, logsLevelDescription)

	return nil
}
func registerOperationServerGetNmsSummaryLogsTimestampParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	logsTimestampDescription := ``

	var logsTimestampFlagName string
	if cmdPrefix == "" {
		logsTimestampFlagName = "logsTimestamp"
	} else {
		logsTimestampFlagName = fmt.Sprintf("%v.logsTimestamp", cmdPrefix)
	}

	var logsTimestampFlagDefault float64

	_ = cmd.PersistentFlags().Float64(logsTimestampFlagName, logsTimestampFlagDefault, logsTimestampDescription)

	return nil
}
func registerOperationServerGetNmsSummaryOutagesTimestampParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	outagesTimestampDescription := ``

	var outagesTimestampFlagName string
	if cmdPrefix == "" {
		outagesTimestampFlagName = "outagesTimestamp"
	} else {
		outagesTimestampFlagName = fmt.Sprintf("%v.outagesTimestamp", cmdPrefix)
	}

	var outagesTimestampFlagDefault float64

	_ = cmd.PersistentFlags().Float64(outagesTimestampFlagName, outagesTimestampFlagDefault, outagesTimestampDescription)

	return nil
}

func retrieveOperationServerGetNmsSummaryFirmwaresTimestampFlag(m *serverops.GetNmsSummaryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("firmwaresTimestamp") {

		var firmwaresTimestampFlagName string
		if cmdPrefix == "" {
			firmwaresTimestampFlagName = "firmwaresTimestamp"
		} else {
			firmwaresTimestampFlagName = fmt.Sprintf("%v.firmwaresTimestamp", cmdPrefix)
		}

		firmwaresTimestampFlagValue, err := cmd.Flags().GetFloat64(firmwaresTimestampFlagName)
		if err != nil {
			return err, false
		}
		m.FirmwaresTimestamp = &firmwaresTimestampFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServerGetNmsSummaryLogsLevelFlag(m *serverops.GetNmsSummaryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("logsLevel") {

		var logsLevelFlagName string
		if cmdPrefix == "" {
			logsLevelFlagName = "logsLevel"
		} else {
			logsLevelFlagName = fmt.Sprintf("%v.logsLevel", cmdPrefix)
		}

		logsLevelFlagValues, err := cmd.Flags().GetStringSlice(logsLevelFlagName)
		if err != nil {
			return err, false
		}
		m.LogsLevel = logsLevelFlagValues

	}
	return nil, retAdded
}
func retrieveOperationServerGetNmsSummaryLogsTimestampFlag(m *serverops.GetNmsSummaryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("logsTimestamp") {

		var logsTimestampFlagName string
		if cmdPrefix == "" {
			logsTimestampFlagName = "logsTimestamp"
		} else {
			logsTimestampFlagName = fmt.Sprintf("%v.logsTimestamp", cmdPrefix)
		}

		logsTimestampFlagValue, err := cmd.Flags().GetFloat64(logsTimestampFlagName)
		if err != nil {
			return err, false
		}
		m.LogsTimestamp = &logsTimestampFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServerGetNmsSummaryOutagesTimestampFlag(m *serverops.GetNmsSummaryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("outagesTimestamp") {

		var outagesTimestampFlagName string
		if cmdPrefix == "" {
			outagesTimestampFlagName = "outagesTimestamp"
		} else {
			outagesTimestampFlagName = fmt.Sprintf("%v.outagesTimestamp", cmdPrefix)
		}

		outagesTimestampFlagValue, err := cmd.Flags().GetFloat64(outagesTimestampFlagName)
		if err != nil {
			return err, false
		}
		m.OutagesTimestamp = &outagesTimestampFlagValue

	}
	return nil, retAdded
}

// parseOperationServerGetNmsSummaryResult parses request result and return the string content
func parseOperationServerGetNmsSummaryResult(resp0 *server.GetNmsSummaryOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*server.GetNmsSummaryOK)
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
		resp1, ok := iResp1.(*server.GetNmsSummaryUnauthorized)
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
		resp2, ok := iResp2.(*server.GetNmsSummaryForbidden)
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
		resp3, ok := iResp3.(*server.GetNmsSummaryInternalServerError)
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