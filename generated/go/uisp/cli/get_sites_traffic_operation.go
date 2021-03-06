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

// makeOperationSitesGetSitesTrafficCmd returns a cmd to handle operation getSitesTraffic
func makeOperationSitesGetSitesTrafficCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getSitesTraffic",
		Short: ``,
		RunE:  runOperationSitesGetSitesTraffic,
	}

	if err := registerOperationSitesGetSitesTrafficParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSitesGetSitesTraffic uses cmd flags to call endpoint api
func runOperationSitesGetSitesTraffic(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := sites.NewGetSitesTrafficParams()
	if err, _ := retrieveOperationSitesGetSitesTrafficFromFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSitesGetSitesTrafficGranularityFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSitesGetSitesTrafficToFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSitesGetSitesTrafficResult(appCli.Sites.GetSitesTraffic(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSitesGetSitesTrafficParamFlags registers all flags needed to fill params
func registerOperationSitesGetSitesTrafficParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSitesGetSitesTrafficFromParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSitesGetSitesTrafficGranularityParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSitesGetSitesTrafficToParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSitesGetSitesTrafficFromParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSitesGetSitesTrafficGranularityParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSitesGetSitesTrafficToParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationSitesGetSitesTrafficFromFlag(m *sites.GetSitesTrafficParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSitesGetSitesTrafficGranularityFlag(m *sites.GetSitesTrafficParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSitesGetSitesTrafficToFlag(m *sites.GetSitesTrafficParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationSitesGetSitesTrafficResult parses request result and return the string content
func parseOperationSitesGetSitesTrafficResult(resp0 *sites.GetSitesTrafficOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*sites.GetSitesTrafficOK)
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
		resp1, ok := iResp1.(*sites.GetSitesTrafficBadRequest)
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
		resp2, ok := iResp2.(*sites.GetSitesTrafficUnauthorized)
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
		resp3, ok := iResp3.(*sites.GetSitesTrafficForbidden)
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
		resp4, ok := iResp4.(*sites.GetSitesTrafficInternalServerError)
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
