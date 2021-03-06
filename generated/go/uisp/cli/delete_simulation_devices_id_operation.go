// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/simulation"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSimulationDeleteSimulationDevicesIDCmd returns a cmd to handle operation deleteSimulationDevicesId
func makeOperationSimulationDeleteSimulationDevicesIDCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteSimulationDevicesId",
		Short: ``,
		RunE:  runOperationSimulationDeleteSimulationDevicesID,
	}

	if err := registerOperationSimulationDeleteSimulationDevicesIDParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSimulationDeleteSimulationDevicesID uses cmd flags to call endpoint api
func runOperationSimulationDeleteSimulationDevicesID(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := simulation.NewDeleteSimulationDevicesIDParams()
	if err, _ := retrieveOperationSimulationDeleteSimulationDevicesIDIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSimulationDeleteSimulationDevicesIDXAuthTokenFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSimulationDeleteSimulationDevicesIDResult(appCli.Simulation.DeleteSimulationDevicesID(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSimulationDeleteSimulationDevicesIDParamFlags registers all flags needed to fill params
func registerOperationSimulationDeleteSimulationDevicesIDParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSimulationDeleteSimulationDevicesIDIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSimulationDeleteSimulationDevicesIDXAuthTokenParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSimulationDeleteSimulationDevicesIDIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSimulationDeleteSimulationDevicesIDXAuthTokenParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	xAuthTokenDescription := `Token for user authorization.`

	var xAuthTokenFlagName string
	if cmdPrefix == "" {
		xAuthTokenFlagName = "x-auth-token"
	} else {
		xAuthTokenFlagName = fmt.Sprintf("%v.x-auth-token", cmdPrefix)
	}

	var xAuthTokenFlagDefault string

	_ = cmd.PersistentFlags().String(xAuthTokenFlagName, xAuthTokenFlagDefault, xAuthTokenDescription)

	return nil
}

func retrieveOperationSimulationDeleteSimulationDevicesIDIDFlag(m *simulation.DeleteSimulationDevicesIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSimulationDeleteSimulationDevicesIDXAuthTokenFlag(m *simulation.DeleteSimulationDevicesIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("x-auth-token") {

		var xAuthTokenFlagName string
		if cmdPrefix == "" {
			xAuthTokenFlagName = "x-auth-token"
		} else {
			xAuthTokenFlagName = fmt.Sprintf("%v.x-auth-token", cmdPrefix)
		}

		xAuthTokenFlagValue, err := cmd.Flags().GetString(xAuthTokenFlagName)
		if err != nil {
			return err, false
		}
		m.XAuthToken = &xAuthTokenFlagValue

	}
	return nil, retAdded
}

// parseOperationSimulationDeleteSimulationDevicesIDResult parses request result and return the string content
func parseOperationSimulationDeleteSimulationDevicesIDResult(resp0 *simulation.DeleteSimulationDevicesIDOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*simulation.DeleteSimulationDevicesIDOK)
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
		resp1, ok := iResp1.(*simulation.DeleteSimulationDevicesIDBadRequest)
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
		resp2, ok := iResp2.(*simulation.DeleteSimulationDevicesIDUnauthorized)
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
		resp3, ok := iResp3.(*simulation.DeleteSimulationDevicesIDNotFound)
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
		resp4, ok := iResp4.(*simulation.DeleteSimulationDevicesIDInternalServerError)
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
