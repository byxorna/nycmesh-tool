// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/client/users"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationUsersDeleteUsersIDCmd returns a cmd to handle operation deleteUsersId
func makeOperationUsersDeleteUsersIDCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteUsersId",
		Short: ``,
		RunE:  runOperationUsersDeleteUsersID,
	}

	if err := registerOperationUsersDeleteUsersIDParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationUsersDeleteUsersID uses cmd flags to call endpoint api
func runOperationUsersDeleteUsersID(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := users.NewDeleteUsersIDParams()
	if err, _ := retrieveOperationUsersDeleteUsersIDIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationUsersDeleteUsersIDNotifyUserFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationUsersDeleteUsersIDResult(appCli.Users.DeleteUsersID(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationUsersDeleteUsersIDParamFlags registers all flags needed to fill params
func registerOperationUsersDeleteUsersIDParamFlags(cmd *cobra.Command) error {
	if err := registerOperationUsersDeleteUsersIDIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationUsersDeleteUsersIDNotifyUserParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationUsersDeleteUsersIDIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationUsersDeleteUsersIDNotifyUserParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	notifyUserDescription := ``

	var notifyUserFlagName string
	if cmdPrefix == "" {
		notifyUserFlagName = "notifyUser"
	} else {
		notifyUserFlagName = fmt.Sprintf("%v.notifyUser", cmdPrefix)
	}

	var notifyUserFlagDefault bool

	_ = cmd.PersistentFlags().Bool(notifyUserFlagName, notifyUserFlagDefault, notifyUserDescription)

	return nil
}

func retrieveOperationUsersDeleteUsersIDIDFlag(m *users.DeleteUsersIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationUsersDeleteUsersIDNotifyUserFlag(m *users.DeleteUsersIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("notifyUser") {

		var notifyUserFlagName string
		if cmdPrefix == "" {
			notifyUserFlagName = "notifyUser"
		} else {
			notifyUserFlagName = fmt.Sprintf("%v.notifyUser", cmdPrefix)
		}

		notifyUserFlagValue, err := cmd.Flags().GetBool(notifyUserFlagName)
		if err != nil {
			return err, false
		}
		m.NotifyUser = &notifyUserFlagValue

	}
	return nil, retAdded
}

// parseOperationUsersDeleteUsersIDResult parses request result and return the string content
func parseOperationUsersDeleteUsersIDResult(resp0 *users.DeleteUsersIDOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*users.DeleteUsersIDOK)
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
		resp1, ok := iResp1.(*users.DeleteUsersIDBadRequest)
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
		resp2, ok := iResp2.(*users.DeleteUsersIDUnauthorized)
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
		resp3, ok := iResp3.(*users.DeleteUsersIDForbidden)
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
		resp4, ok := iResp4.(*users.DeleteUsersIDNotFound)
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
		resp5, ok := iResp5.(*users.DeleteUsersIDInternalServerError)
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
