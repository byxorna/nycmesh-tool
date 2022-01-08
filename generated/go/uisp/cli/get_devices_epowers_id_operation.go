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

// makeOperationDevicesGetDevicesEpowersIDCmd returns a cmd to handle operation getDevicesEpowersId
func makeOperationDevicesGetDevicesEpowersIDCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getDevicesEpowersId",
		Short: ``,
		RunE:  runOperationDevicesGetDevicesEpowersID,
	}

	if err := registerOperationDevicesGetDevicesEpowersIDParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDevicesGetDevicesEpowersID uses cmd flags to call endpoint api
func runOperationDevicesGetDevicesEpowersID(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := devices.NewGetDevicesEpowersIDParams()
	if err, _ := retrieveOperationDevicesGetDevicesEpowersIDIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDevicesGetDevicesEpowersIDResult(appCli.Devices.GetDevicesEpowersID(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDevicesGetDevicesEpowersIDParamFlags registers all flags needed to fill params
func registerOperationDevicesGetDevicesEpowersIDParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDevicesGetDevicesEpowersIDIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDevicesGetDevicesEpowersIDIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationDevicesGetDevicesEpowersIDIDFlag(m *devices.GetDevicesEpowersIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationDevicesGetDevicesEpowersIDResult parses request result and return the string content
func parseOperationDevicesGetDevicesEpowersIDResult(respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*devices.GetDevicesEpowersIDDefault)
		if ok {
			if !swag.IsZero(respD) && !swag.IsZero(respD.Payload) {
				msgStr, err := json.Marshal(respD.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}
	return "", nil
}