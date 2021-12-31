// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/client/devices"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationDevicesPostDevicesEroutersIDDhcpServersServernameBlockCmd returns a cmd to handle operation postDevicesEroutersIdDhcpServersServernameBlock
func makeOperationDevicesPostDevicesEroutersIDDhcpServersServernameBlockCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "postDevicesEroutersIdDhcpServersServernameBlock",
		Short: ``,
		RunE:  runOperationDevicesPostDevicesEroutersIDDhcpServersServernameBlock,
	}

	if err := registerOperationDevicesPostDevicesEroutersIDDhcpServersServernameBlockParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDevicesPostDevicesEroutersIDDhcpServersServernameBlock uses cmd flags to call endpoint api
func runOperationDevicesPostDevicesEroutersIDDhcpServersServernameBlock(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := devices.NewPostDevicesEroutersIDDhcpServersServernameBlockParams()
	if err, _ := retrieveOperationDevicesPostDevicesEroutersIDDhcpServersServernameBlockIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesPostDevicesEroutersIDDhcpServersServernameBlockServerNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDevicesPostDevicesEroutersIDDhcpServersServernameBlockResult(appCli.Devices.PostDevicesEroutersIDDhcpServersServernameBlock(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDevicesPostDevicesEroutersIDDhcpServersServernameBlockParamFlags registers all flags needed to fill params
func registerOperationDevicesPostDevicesEroutersIDDhcpServersServernameBlockParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDevicesPostDevicesEroutersIDDhcpServersServernameBlockIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesPostDevicesEroutersIDDhcpServersServernameBlockServerNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDevicesPostDevicesEroutersIDDhcpServersServernameBlockIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationDevicesPostDevicesEroutersIDDhcpServersServernameBlockServerNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationDevicesPostDevicesEroutersIDDhcpServersServernameBlockIDFlag(m *devices.PostDevicesEroutersIDDhcpServersServernameBlockParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDevicesPostDevicesEroutersIDDhcpServersServernameBlockServerNameFlag(m *devices.PostDevicesEroutersIDDhcpServersServernameBlockParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationDevicesPostDevicesEroutersIDDhcpServersServernameBlockResult parses request result and return the string content
func parseOperationDevicesPostDevicesEroutersIDDhcpServersServernameBlockResult(resp0 *devices.PostDevicesEroutersIDDhcpServersServernameBlockOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*devices.PostDevicesEroutersIDDhcpServersServernameBlockOK)
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
		resp1, ok := iResp1.(*devices.PostDevicesEroutersIDDhcpServersServernameBlockBadRequest)
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
		resp2, ok := iResp2.(*devices.PostDevicesEroutersIDDhcpServersServernameBlockUnauthorized)
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
		resp3, ok := iResp3.(*devices.PostDevicesEroutersIDDhcpServersServernameBlockForbidden)
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
		resp4, ok := iResp4.(*devices.PostDevicesEroutersIDDhcpServersServernameBlockNotFound)
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
		resp5, ok := iResp5.(*devices.PostDevicesEroutersIDDhcpServersServernameBlockInternalServerError)
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
