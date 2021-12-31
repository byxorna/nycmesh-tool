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

// makeOperationServerGetNmsChangedCmd returns a cmd to handle operation getNmsChanged
func makeOperationServerGetNmsChangedCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getNmsChanged",
		Short: ``,
		RunE:  runOperationServerGetNmsChanged,
	}

	if err := registerOperationServerGetNmsChangedParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServerGetNmsChanged uses cmd flags to call endpoint api
func runOperationServerGetNmsChanged(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := serverops.NewGetNmsChangedParams()
	if err, _ := retrieveOperationServerGetNmsChangedSinceFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServerGetNmsChangedUcrmFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationServerGetNmsChangedResult(appCli.Server.GetNmsChanged(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationServerGetNmsChangedParamFlags registers all flags needed to fill params
func registerOperationServerGetNmsChangedParamFlags(cmd *cobra.Command) error {
	if err := registerOperationServerGetNmsChangedSinceParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServerGetNmsChangedUcrmParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationServerGetNmsChangedSinceParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	sinceDescription := `Required. Timestamp in milliseconds.`

	var sinceFlagName string
	if cmdPrefix == "" {
		sinceFlagName = "since"
	} else {
		sinceFlagName = fmt.Sprintf("%v.since", cmdPrefix)
	}

	var sinceFlagDefault int64

	_ = cmd.PersistentFlags().Int64(sinceFlagName, sinceFlagDefault, sinceDescription)

	return nil
}
func registerOperationServerGetNmsChangedUcrmParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	ucrmDescription := `Search items related to UCRM.`

	var ucrmFlagName string
	if cmdPrefix == "" {
		ucrmFlagName = "ucrm"
	} else {
		ucrmFlagName = fmt.Sprintf("%v.ucrm", cmdPrefix)
	}

	var ucrmFlagDefault bool

	_ = cmd.PersistentFlags().Bool(ucrmFlagName, ucrmFlagDefault, ucrmDescription)

	return nil
}

func retrieveOperationServerGetNmsChangedSinceFlag(m *serverops.GetNmsChangedParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("since") {

		var sinceFlagName string
		if cmdPrefix == "" {
			sinceFlagName = "since"
		} else {
			sinceFlagName = fmt.Sprintf("%v.since", cmdPrefix)
		}

		sinceFlagValue, err := cmd.Flags().GetInt64(sinceFlagName)
		if err != nil {
			return err, false
		}
		m.Since = sinceFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServerGetNmsChangedUcrmFlag(m *serverops.GetNmsChangedParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("ucrm") {

		var ucrmFlagName string
		if cmdPrefix == "" {
			ucrmFlagName = "ucrm"
		} else {
			ucrmFlagName = fmt.Sprintf("%v.ucrm", cmdPrefix)
		}

		ucrmFlagValue, err := cmd.Flags().GetBool(ucrmFlagName)
		if err != nil {
			return err, false
		}
		m.Ucrm = &ucrmFlagValue

	}
	return nil, retAdded
}

// parseOperationServerGetNmsChangedResult parses request result and return the string content
func parseOperationServerGetNmsChangedResult(resp0 *server.GetNmsChangedOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*server.GetNmsChangedOK)
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
		resp1, ok := iResp1.(*server.GetNmsChangedUnauthorized)
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
		resp2, ok := iResp2.(*server.GetNmsChangedForbidden)
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
		resp3, ok := iResp3.(*server.GetNmsChangedInternalServerError)
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
