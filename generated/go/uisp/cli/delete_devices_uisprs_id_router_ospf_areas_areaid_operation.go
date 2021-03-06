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

// makeOperationDevicesDeleteDevicesUisprsIDRouterOspfAreasAreaidCmd returns a cmd to handle operation deleteDevicesUisprsIdRouterOspfAreasAreaid
func makeOperationDevicesDeleteDevicesUisprsIDRouterOspfAreasAreaidCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteDevicesUisprsIdRouterOspfAreasAreaid",
		Short: ``,
		RunE:  runOperationDevicesDeleteDevicesUisprsIDRouterOspfAreasAreaid,
	}

	if err := registerOperationDevicesDeleteDevicesUisprsIDRouterOspfAreasAreaidParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDevicesDeleteDevicesUisprsIDRouterOspfAreasAreaid uses cmd flags to call endpoint api
func runOperationDevicesDeleteDevicesUisprsIDRouterOspfAreasAreaid(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := devices.NewDeleteDevicesUisprsIDRouterOspfAreasAreaidParams()
	if err, _ := retrieveOperationDevicesDeleteDevicesUisprsIDRouterOspfAreasAreaidAreaIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesDeleteDevicesUisprsIDRouterOspfAreasAreaidIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDevicesDeleteDevicesUisprsIDRouterOspfAreasAreaidResult(appCli.Devices.DeleteDevicesUisprsIDRouterOspfAreasAreaid(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDevicesDeleteDevicesUisprsIDRouterOspfAreasAreaidParamFlags registers all flags needed to fill params
func registerOperationDevicesDeleteDevicesUisprsIDRouterOspfAreasAreaidParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDevicesDeleteDevicesUisprsIDRouterOspfAreasAreaidAreaIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesDeleteDevicesUisprsIDRouterOspfAreasAreaidIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDevicesDeleteDevicesUisprsIDRouterOspfAreasAreaidAreaIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	areaIdDescription := `Required. `

	var areaIdFlagName string
	if cmdPrefix == "" {
		areaIdFlagName = "areaId"
	} else {
		areaIdFlagName = fmt.Sprintf("%v.areaId", cmdPrefix)
	}

	var areaIdFlagDefault string

	_ = cmd.PersistentFlags().String(areaIdFlagName, areaIdFlagDefault, areaIdDescription)

	return nil
}
func registerOperationDevicesDeleteDevicesUisprsIDRouterOspfAreasAreaidIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationDevicesDeleteDevicesUisprsIDRouterOspfAreasAreaidAreaIDFlag(m *devices.DeleteDevicesUisprsIDRouterOspfAreasAreaidParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("areaId") {

		var areaIdFlagName string
		if cmdPrefix == "" {
			areaIdFlagName = "areaId"
		} else {
			areaIdFlagName = fmt.Sprintf("%v.areaId", cmdPrefix)
		}

		areaIdFlagValue, err := cmd.Flags().GetString(areaIdFlagName)
		if err != nil {
			return err, false
		}
		m.AreaID = areaIdFlagValue

	}
	return nil, retAdded
}
func retrieveOperationDevicesDeleteDevicesUisprsIDRouterOspfAreasAreaidIDFlag(m *devices.DeleteDevicesUisprsIDRouterOspfAreasAreaidParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationDevicesDeleteDevicesUisprsIDRouterOspfAreasAreaidResult parses request result and return the string content
func parseOperationDevicesDeleteDevicesUisprsIDRouterOspfAreasAreaidResult(resp0 *devices.DeleteDevicesUisprsIDRouterOspfAreasAreaidOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*devices.DeleteDevicesUisprsIDRouterOspfAreasAreaidOK)
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
		resp1, ok := iResp1.(*devices.DeleteDevicesUisprsIDRouterOspfAreasAreaidBadRequest)
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
		resp2, ok := iResp2.(*devices.DeleteDevicesUisprsIDRouterOspfAreasAreaidUnauthorized)
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
		resp3, ok := iResp3.(*devices.DeleteDevicesUisprsIDRouterOspfAreasAreaidForbidden)
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
		resp4, ok := iResp4.(*devices.DeleteDevicesUisprsIDRouterOspfAreasAreaidNotFound)
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
		resp5, ok := iResp5.(*devices.DeleteDevicesUisprsIDRouterOspfAreasAreaidInternalServerError)
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
