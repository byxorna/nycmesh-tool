// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/sites"
	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSitesPatchSitesSiteidImagesImageidCmd returns a cmd to handle operation patchSitesSiteidImagesImageid
func makeOperationSitesPatchSitesSiteidImagesImageidCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "patchSitesSiteidImagesImageid",
		Short: ``,
		RunE:  runOperationSitesPatchSitesSiteidImagesImageid,
	}

	if err := registerOperationSitesPatchSitesSiteidImagesImageidParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSitesPatchSitesSiteidImagesImageid uses cmd flags to call endpoint api
func runOperationSitesPatchSitesSiteidImagesImageid(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := sites.NewPatchSitesSiteidImagesImageidParams()
	if err, _ := retrieveOperationSitesPatchSitesSiteidImagesImageidBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSitesPatchSitesSiteidImagesImageidImageIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSitesPatchSitesSiteidImagesImageidSiteIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSitesPatchSitesSiteidImagesImageidResult(appCli.Sites.PatchSitesSiteidImagesImageid(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSitesPatchSitesSiteidImagesImageidParamFlags registers all flags needed to fill params
func registerOperationSitesPatchSitesSiteidImagesImageidParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSitesPatchSitesSiteidImagesImageidBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSitesPatchSitesSiteidImagesImageidImageIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSitesPatchSitesSiteidImagesImageidSiteIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSitesPatchSitesSiteidImagesImageidBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelImageDataFlags(0, "imageData", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationSitesPatchSitesSiteidImagesImageidImageIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	imageIdDescription := `Required. `

	var imageIdFlagName string
	if cmdPrefix == "" {
		imageIdFlagName = "imageId"
	} else {
		imageIdFlagName = fmt.Sprintf("%v.imageId", cmdPrefix)
	}

	var imageIdFlagDefault string

	_ = cmd.PersistentFlags().String(imageIdFlagName, imageIdFlagDefault, imageIdDescription)

	return nil
}
func registerOperationSitesPatchSitesSiteidImagesImageidSiteIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	siteIdDescription := `Required. `

	var siteIdFlagName string
	if cmdPrefix == "" {
		siteIdFlagName = "siteId"
	} else {
		siteIdFlagName = fmt.Sprintf("%v.siteId", cmdPrefix)
	}

	var siteIdFlagDefault string

	_ = cmd.PersistentFlags().String(siteIdFlagName, siteIdFlagDefault, siteIdDescription)

	return nil
}

func retrieveOperationSitesPatchSitesSiteidImagesImageidBodyFlag(m *sites.PatchSitesSiteidImagesImageidParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.ImageData{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.ImageData: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.ImageData{}
	}
	err, added := retrieveModelImageDataFlags(0, bodyValueModel, "imageData", cmd)
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
func retrieveOperationSitesPatchSitesSiteidImagesImageidImageIDFlag(m *sites.PatchSitesSiteidImagesImageidParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("imageId") {

		var imageIdFlagName string
		if cmdPrefix == "" {
			imageIdFlagName = "imageId"
		} else {
			imageIdFlagName = fmt.Sprintf("%v.imageId", cmdPrefix)
		}

		imageIdFlagValue, err := cmd.Flags().GetString(imageIdFlagName)
		if err != nil {
			return err, false
		}
		m.ImageID = imageIdFlagValue

	}
	return nil, retAdded
}
func retrieveOperationSitesPatchSitesSiteidImagesImageidSiteIDFlag(m *sites.PatchSitesSiteidImagesImageidParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("siteId") {

		var siteIdFlagName string
		if cmdPrefix == "" {
			siteIdFlagName = "siteId"
		} else {
			siteIdFlagName = fmt.Sprintf("%v.siteId", cmdPrefix)
		}

		siteIdFlagValue, err := cmd.Flags().GetString(siteIdFlagName)
		if err != nil {
			return err, false
		}
		m.SiteID = siteIdFlagValue

	}
	return nil, retAdded
}

// parseOperationSitesPatchSitesSiteidImagesImageidResult parses request result and return the string content
func parseOperationSitesPatchSitesSiteidImagesImageidResult(resp0 *sites.PatchSitesSiteidImagesImageidOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*sites.PatchSitesSiteidImagesImageidOK)
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
		resp1, ok := iResp1.(*sites.PatchSitesSiteidImagesImageidBadRequest)
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
		resp2, ok := iResp2.(*sites.PatchSitesSiteidImagesImageidUnauthorized)
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
		resp3, ok := iResp3.(*sites.PatchSitesSiteidImagesImageidForbidden)
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
		resp4, ok := iResp4.(*sites.PatchSitesSiteidImagesImageidInternalServerError)
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
