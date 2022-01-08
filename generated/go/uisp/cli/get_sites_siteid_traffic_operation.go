// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/sites"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSitesGetSitesSiteidTrafficCmd returns a cmd to handle operation getSitesSiteidTraffic
func makeOperationSitesGetSitesSiteidTrafficCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getSitesSiteidTraffic",
		Short: ``,
		RunE:  runOperationSitesGetSitesSiteidTraffic,
	}

	if err := registerOperationSitesGetSitesSiteidTrafficParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSitesGetSitesSiteidTraffic uses cmd flags to call endpoint api
func runOperationSitesGetSitesSiteidTraffic(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := sites.NewGetSitesSiteidTrafficParams()
	if err, _ := retrieveOperationSitesGetSitesSiteidTrafficFromFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSitesGetSitesSiteidTrafficGranularityFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSitesGetSitesSiteidTrafficSiteIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSitesGetSitesSiteidTrafficToFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSitesGetSitesSiteidTrafficResult(appCli.Sites.GetSitesSiteidTraffic(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSitesGetSitesSiteidTrafficParamFlags registers all flags needed to fill params
func registerOperationSitesGetSitesSiteidTrafficParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSitesGetSitesSiteidTrafficFromParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSitesGetSitesSiteidTrafficGranularityParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSitesGetSitesSiteidTrafficSiteIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSitesGetSitesSiteidTrafficToParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSitesGetSitesSiteidTrafficFromParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	fromDescription := `Required. Timestamp of the beginning of the interval in milliseconds`

	var fromFlagName string
	if cmdPrefix == "" {
		fromFlagName = "from"
	} else {
		fromFlagName = fmt.Sprintf("%v.from", cmdPrefix)
	}

	var fromFlagDefault float64

	_ = cmd.PersistentFlags().Float64(fromFlagName, fromFlagDefault, fromDescription)

	return nil
}
func registerOperationSitesGetSitesSiteidTrafficGranularityParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	granularityDescription := `Enum: ["fiveMinutes","oneHour"]. Required. Granularity of the traffic stats`

	var granularityFlagName string
	if cmdPrefix == "" {
		granularityFlagName = "granularity"
	} else {
		granularityFlagName = fmt.Sprintf("%v.granularity", cmdPrefix)
	}

	var granularityFlagDefault string

	_ = cmd.PersistentFlags().String(granularityFlagName, granularityFlagDefault, granularityDescription)

	if err := cmd.RegisterFlagCompletionFunc(granularityFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["fiveMinutes","oneHour"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}
func registerOperationSitesGetSitesSiteidTrafficSiteIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSitesGetSitesSiteidTrafficToParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	toDescription := `Required. Timestamp of the end of the interval in milliseconds`

	var toFlagName string
	if cmdPrefix == "" {
		toFlagName = "to"
	} else {
		toFlagName = fmt.Sprintf("%v.to", cmdPrefix)
	}

	var toFlagDefault float64

	_ = cmd.PersistentFlags().Float64(toFlagName, toFlagDefault, toDescription)

	return nil
}

func retrieveOperationSitesGetSitesSiteidTrafficFromFlag(m *sites.GetSitesSiteidTrafficParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("from") {

		var fromFlagName string
		if cmdPrefix == "" {
			fromFlagName = "from"
		} else {
			fromFlagName = fmt.Sprintf("%v.from", cmdPrefix)
		}

		fromFlagValue, err := cmd.Flags().GetFloat64(fromFlagName)
		if err != nil {
			return err, false
		}
		m.From = fromFlagValue

	}
	return nil, retAdded
}
func retrieveOperationSitesGetSitesSiteidTrafficGranularityFlag(m *sites.GetSitesSiteidTrafficParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("granularity") {

		var granularityFlagName string
		if cmdPrefix == "" {
			granularityFlagName = "granularity"
		} else {
			granularityFlagName = fmt.Sprintf("%v.granularity", cmdPrefix)
		}

		granularityFlagValue, err := cmd.Flags().GetString(granularityFlagName)
		if err != nil {
			return err, false
		}
		m.Granularity = granularityFlagValue

	}
	return nil, retAdded
}
func retrieveOperationSitesGetSitesSiteidTrafficSiteIDFlag(m *sites.GetSitesSiteidTrafficParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSitesGetSitesSiteidTrafficToFlag(m *sites.GetSitesSiteidTrafficParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("to") {

		var toFlagName string
		if cmdPrefix == "" {
			toFlagName = "to"
		} else {
			toFlagName = fmt.Sprintf("%v.to", cmdPrefix)
		}

		toFlagValue, err := cmd.Flags().GetFloat64(toFlagName)
		if err != nil {
			return err, false
		}
		m.To = toFlagValue

	}
	return nil, retAdded
}

// parseOperationSitesGetSitesSiteidTrafficResult parses request result and return the string content
func parseOperationSitesGetSitesSiteidTrafficResult(resp0 *sites.GetSitesSiteidTrafficOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*sites.GetSitesSiteidTrafficOK)
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
		resp1, ok := iResp1.(*sites.GetSitesSiteidTrafficBadRequest)
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
		resp2, ok := iResp2.(*sites.GetSitesSiteidTrafficUnauthorized)
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
		resp3, ok := iResp3.(*sites.GetSitesSiteidTrafficForbidden)
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
		resp4, ok := iResp4.(*sites.GetSitesSiteidTrafficInternalServerError)
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