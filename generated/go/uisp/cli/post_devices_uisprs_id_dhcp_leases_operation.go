// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/devices"
	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationDevicesPostDevicesUisprsIDDhcpLeasesCmd returns a cmd to handle operation postDevicesUisprsIdDhcpLeases
func makeOperationDevicesPostDevicesUisprsIDDhcpLeasesCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "postDevicesUisprsIdDhcpLeases",
		Short: `
        Request does accept payload with "leaseId" property, to match api,
        but the id is ignored and it will be created from mac and ip.
      `,
		RunE: runOperationDevicesPostDevicesUisprsIDDhcpLeases,
	}

	if err := registerOperationDevicesPostDevicesUisprsIDDhcpLeasesParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDevicesPostDevicesUisprsIDDhcpLeases uses cmd flags to call endpoint api
func runOperationDevicesPostDevicesUisprsIDDhcpLeases(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := devices.NewPostDevicesUisprsIDDhcpLeasesParams()
	if err, _ := retrieveOperationDevicesPostDevicesUisprsIDDhcpLeasesBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDevicesPostDevicesUisprsIDDhcpLeasesIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDevicesPostDevicesUisprsIDDhcpLeasesResult(appCli.Devices.PostDevicesUisprsIDDhcpLeases(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDevicesPostDevicesUisprsIDDhcpLeasesParamFlags registers all flags needed to fill params
func registerOperationDevicesPostDevicesUisprsIDDhcpLeasesParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDevicesPostDevicesUisprsIDDhcpLeasesBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDevicesPostDevicesUisprsIDDhcpLeasesIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDevicesPostDevicesUisprsIDDhcpLeasesBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationDevicesPostDevicesUisprsIDDhcpLeasesIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationDevicesPostDevicesUisprsIDDhcpLeasesBodyFlag(m *devices.PostDevicesUisprsIDDhcpLeasesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDevicesPostDevicesUisprsIDDhcpLeasesIDFlag(m *devices.PostDevicesUisprsIDDhcpLeasesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationDevicesPostDevicesUisprsIDDhcpLeasesResult parses request result and return the string content
func parseOperationDevicesPostDevicesUisprsIDDhcpLeasesResult(resp0 *devices.PostDevicesUisprsIDDhcpLeasesOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*devices.PostDevicesUisprsIDDhcpLeasesOK)
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
		resp1, ok := iResp1.(*devices.PostDevicesUisprsIDDhcpLeasesBadRequest)
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
		resp2, ok := iResp2.(*devices.PostDevicesUisprsIDDhcpLeasesUnauthorized)
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
		resp3, ok := iResp3.(*devices.PostDevicesUisprsIDDhcpLeasesForbidden)
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
		resp4, ok := iResp4.(*devices.PostDevicesUisprsIDDhcpLeasesNotFound)
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
		resp5, ok := iResp5.(*devices.PostDevicesUisprsIDDhcpLeasesInternalServerError)
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
