// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/client/devices"
	"github.com/byxorna/nycmesh-tool/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidCmd returns a cmd to handle operation putDevicesUisprsIdDhcpLeasesServernameLeaseid
func makeOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "putDevicesUisprsIdDhcpLeasesServernameLeaseid",
		Short: `
        Request does accept payload with "leaseId" property, to match api,
        but the id is ignored and it will be changed when mac or ip is updated.
      `,
		RunE: runOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseid,
	}

	if err := registerOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseid uses cmd flags to call endpoint api
func runOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseid(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := devices.NewPutDevicesUisprsIDDhcpLeasesServernameLeaseidParams()
	if err, _ := retrieveOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidLeaseIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidServerNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidResult(appCli.Devices.PutDevicesUisprsIDDhcpLeasesServernameLeaseid(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidParamFlags registers all flags needed to fill params
func registerOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidLeaseIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidServerNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelDHCPLease3Flags(0, "dHCPLease3", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidLeaseIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	leaseIdDescription := `Required. `

	var leaseIdFlagName string
	if cmdPrefix == "" {
		leaseIdFlagName = "leaseId"
	} else {
		leaseIdFlagName = fmt.Sprintf("%v.leaseId", cmdPrefix)
	}

	var leaseIdFlagDefault string

	_ = cmd.PersistentFlags().String(leaseIdFlagName, leaseIdFlagDefault, leaseIdDescription)

	return nil
}
func registerOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidServerNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidBodyFlag(m *devices.PutDevicesUisprsIDDhcpLeasesServernameLeaseidParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.DHCPLease3{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.DHCPLease3: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.DHCPLease3{}
	}
	err, added := retrieveModelDHCPLease3Flags(0, bodyValueModel, "dHCPLease3", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Body = bodyValueModel
	}
	if dryRun && debug {

		bodyValueDebugBytes, err := json.Marshal(m.Body)
		if err != nil {
			return err, false
		}
		logDebugf("Body dry-run payload: %v", string(bodyValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidIDFlag(m *devices.PutDevicesUisprsIDDhcpLeasesServernameLeaseidParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidLeaseIDFlag(m *devices.PutDevicesUisprsIDDhcpLeasesServernameLeaseidParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("leaseId") {

		var leaseIdFlagName string
		if cmdPrefix == "" {
			leaseIdFlagName = "leaseId"
		} else {
			leaseIdFlagName = fmt.Sprintf("%v.leaseId", cmdPrefix)
		}

		leaseIdFlagValue, err := cmd.Flags().GetString(leaseIdFlagName)
		if err != nil {
			return err, false
		}
		m.LeaseID = leaseIdFlagValue

	}
	return nil, retAdded
}
func retrieveOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidServerNameFlag(m *devices.PutDevicesUisprsIDDhcpLeasesServernameLeaseidParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidResult parses request result and return the string content
func parseOperationDevicesPutDevicesUisprsIDDhcpLeasesServernameLeaseidResult(resp0 *devices.PutDevicesUisprsIDDhcpLeasesServernameLeaseidOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*devices.PutDevicesUisprsIDDhcpLeasesServernameLeaseidOK)
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
		resp1, ok := iResp1.(*devices.PutDevicesUisprsIDDhcpLeasesServernameLeaseidBadRequest)
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
		resp2, ok := iResp2.(*devices.PutDevicesUisprsIDDhcpLeasesServernameLeaseidUnauthorized)
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
		resp3, ok := iResp3.(*devices.PutDevicesUisprsIDDhcpLeasesServernameLeaseidForbidden)
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
		resp4, ok := iResp4.(*devices.PutDevicesUisprsIDDhcpLeasesServernameLeaseidNotFound)
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
		resp5, ok := iResp5.(*devices.PutDevicesUisprsIDDhcpLeasesServernameLeaseidInternalServerError)
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