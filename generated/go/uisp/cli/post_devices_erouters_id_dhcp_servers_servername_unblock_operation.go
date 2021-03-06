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

// makeOperationDevicesPostDevicesEroutersIDDhcpServersServernameUnblockCmd returns a cmd to handle operation postDevicesEroutersIdDhcpServersServernameUnblock
func makeOperationDevicesPostDevicesEroutersIDDhcpServersServernameUnblockCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "postDevicesEroutersIdDhcpServersServernameUnblock",
		Short: ``,
		RunE:  runOperationDevicesPostDevicesEroutersIDDhcpServersServernameUnblock,
	}

	if err := registerOperationDevicesPostDevicesEroutersIDDhcpServersServernameUnblockParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDevicesPostDevicesEroutersIDDhcpServersServernameUnblock uses cmd flags to call endpoint api
func runOperationDevicesPostDevicesEroutersIDDhcpServersServernameUnblock(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := devices.NewPostDevicesEroutersIDDhcpServersServernameUnblockParams()
	if err, _ := retrieveOperationDevicesPostDevicesEroutersIDDhcpServersServernameUnblockIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesPostDevicesEroutersIDDhcpServersServernameUnblockServerNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDevicesPostDevicesEroutersIDDhcpServersServernameUnblockResult(appCli.Devices.PostDevicesEroutersIDDhcpServersServernameUnblock(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDevicesPostDevicesEroutersIDDhcpServersServernameUnblockParamFlags registers all flags needed to fill params
func registerOperationDevicesPostDevicesEroutersIDDhcpServersServernameUnblockParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDevicesPostDevicesEroutersIDDhcpServersServernameUnblockIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesPostDevicesEroutersIDDhcpServersServernameUnblockServerNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDevicesPostDevicesEroutersIDDhcpServersServernameUnblockIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationDevicesPostDevicesEroutersIDDhcpServersServernameUnblockServerNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	serverNameDescription := `Required. `

	var serverNameFlagName string
	if cmdPrefix == "" {
		serverNameFlagName = "serverName"
	} else {
		serverNameFlagName = fmt.Sprintf("%v.serverName", cmdPrefix)
	}

	var serverNameFlagDefault string

	_ = cmd.PersistentFlags().String(serverNameFlagName, serverNameFlagDefault, serverNameDescription)

	return nil
}

func retrieveOperationDevicesPostDevicesEroutersIDDhcpServersServernameUnblockIDFlag(m *devices.PostDevicesEroutersIDDhcpServersServernameUnblockParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDevicesPostDevicesEroutersIDDhcpServersServernameUnblockServerNameFlag(m *devices.PostDevicesEroutersIDDhcpServersServernameUnblockParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("serverName") {

		var serverNameFlagName string
		if cmdPrefix == "" {
			serverNameFlagName = "serverName"
		} else {
			serverNameFlagName = fmt.Sprintf("%v.serverName", cmdPrefix)
		}

		serverNameFlagValue, err := cmd.Flags().GetString(serverNameFlagName)
		if err != nil {
			return err, false
		}
		m.ServerName = serverNameFlagValue

	}
	return nil, retAdded
}

// parseOperationDevicesPostDevicesEroutersIDDhcpServersServernameUnblockResult parses request result and return the string content
func parseOperationDevicesPostDevicesEroutersIDDhcpServersServernameUnblockResult(resp0 *devices.PostDevicesEroutersIDDhcpServersServernameUnblockOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*devices.PostDevicesEroutersIDDhcpServersServernameUnblockOK)
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
		resp1, ok := iResp1.(*devices.PostDevicesEroutersIDDhcpServersServernameUnblockBadRequest)
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
		resp2, ok := iResp2.(*devices.PostDevicesEroutersIDDhcpServersServernameUnblockUnauthorized)
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
		resp3, ok := iResp3.(*devices.PostDevicesEroutersIDDhcpServersServernameUnblockForbidden)
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
		resp4, ok := iResp4.(*devices.PostDevicesEroutersIDDhcpServersServernameUnblockNotFound)
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
		resp5, ok := iResp5.(*devices.PostDevicesEroutersIDDhcpServersServernameUnblockInternalServerError)
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
