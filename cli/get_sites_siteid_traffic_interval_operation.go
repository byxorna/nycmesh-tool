// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/client/sites"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSitesGetSitesSiteidTrafficIntervalCmd returns a cmd to handle operation getSitesSiteidTrafficInterval
func makeOperationSitesGetSitesSiteidTrafficIntervalCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getSitesSiteidTrafficInterval",
		Short: ``,
		RunE:  runOperationSitesGetSitesSiteidTrafficInterval,
	}

	if err := registerOperationSitesGetSitesSiteidTrafficIntervalParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSitesGetSitesSiteidTrafficInterval uses cmd flags to call endpoint api
func runOperationSitesGetSitesSiteidTrafficInterval(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := sites.NewGetSitesSiteidTrafficIntervalParams()
	if err, _ := retrieveOperationSitesGetSitesSiteidTrafficIntervalGranularityFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSitesGetSitesSiteidTrafficIntervalIntervalFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSitesGetSitesSiteidTrafficIntervalSiteIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSitesGetSitesSiteidTrafficIntervalResult(appCli.Sites.GetSitesSiteidTrafficInterval(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSitesGetSitesSiteidTrafficIntervalParamFlags registers all flags needed to fill params
func registerOperationSitesGetSitesSiteidTrafficIntervalParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSitesGetSitesSiteidTrafficIntervalGranularityParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSitesGetSitesSiteidTrafficIntervalIntervalParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSitesGetSitesSiteidTrafficIntervalSiteIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSitesGetSitesSiteidTrafficIntervalGranularityParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSitesGetSitesSiteidTrafficIntervalIntervalParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	intervalDescription := `Enum: ["hour","day","month"]. Required. Interval for which to fetch traffic stats`

	var intervalFlagName string
	if cmdPrefix == "" {
		intervalFlagName = "interval"
	} else {
		intervalFlagName = fmt.Sprintf("%v.interval", cmdPrefix)
	}

	var intervalFlagDefault string

	_ = cmd.PersistentFlags().String(intervalFlagName, intervalFlagDefault, intervalDescription)

	if err := cmd.RegisterFlagCompletionFunc(intervalFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["hour","day","month"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}
func registerOperationSitesGetSitesSiteidTrafficIntervalSiteIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationSitesGetSitesSiteidTrafficIntervalGranularityFlag(m *sites.GetSitesSiteidTrafficIntervalParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSitesGetSitesSiteidTrafficIntervalIntervalFlag(m *sites.GetSitesSiteidTrafficIntervalParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("interval") {

		var intervalFlagName string
		if cmdPrefix == "" {
			intervalFlagName = "interval"
		} else {
			intervalFlagName = fmt.Sprintf("%v.interval", cmdPrefix)
		}

		intervalFlagValue, err := cmd.Flags().GetString(intervalFlagName)
		if err != nil {
			return err, false
		}
		m.Interval = intervalFlagValue

	}
	return nil, retAdded
}
func retrieveOperationSitesGetSitesSiteidTrafficIntervalSiteIDFlag(m *sites.GetSitesSiteidTrafficIntervalParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationSitesGetSitesSiteidTrafficIntervalResult parses request result and return the string content
func parseOperationSitesGetSitesSiteidTrafficIntervalResult(resp0 *sites.GetSitesSiteidTrafficIntervalOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*sites.GetSitesSiteidTrafficIntervalOK)
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
		resp1, ok := iResp1.(*sites.GetSitesSiteidTrafficIntervalBadRequest)
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
		resp2, ok := iResp2.(*sites.GetSitesSiteidTrafficIntervalUnauthorized)
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
		resp3, ok := iResp3.(*sites.GetSitesSiteidTrafficIntervalForbidden)
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
		resp4, ok := iResp4.(*sites.GetSitesSiteidTrafficIntervalInternalServerError)
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
