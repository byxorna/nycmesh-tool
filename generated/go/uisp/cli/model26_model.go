// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for Model26

// register flags to command
func registerModelModel26Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel26NetflowStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel26NetflowStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	netflowStatusDescription := `Enum: ["active","connecting","discovered","inactive","disabled","disconnected","unauthorized","proposed","unknown","unplaced","custom"]. `

	var netflowStatusFlagName string
	if cmdPrefix == "" {
		netflowStatusFlagName = "netflowStatus"
	} else {
		netflowStatusFlagName = fmt.Sprintf("%v.netflowStatus", cmdPrefix)
	}

	var netflowStatusFlagDefault string

	_ = cmd.PersistentFlags().String(netflowStatusFlagName, netflowStatusFlagDefault, netflowStatusDescription)

	if err := cmd.RegisterFlagCompletionFunc(netflowStatusFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["active","connecting","discovered","inactive","disabled","disconnected","unauthorized","proposed","unknown","unplaced","custom"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModel26Flags(depth int, m *models.Model26, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, netflowStatusAdded := retrieveModel26NetflowStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || netflowStatusAdded

	return nil, retAdded
}

func retrieveModel26NetflowStatusFlags(depth int, m *models.Model26, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	netflowStatusFlagName := fmt.Sprintf("%v.netflowStatus", cmdPrefix)
	if cmd.Flags().Changed(netflowStatusFlagName) {

		var netflowStatusFlagName string
		if cmdPrefix == "" {
			netflowStatusFlagName = "netflowStatus"
		} else {
			netflowStatusFlagName = fmt.Sprintf("%v.netflowStatus", cmdPrefix)
		}

		netflowStatusFlagValue, err := cmd.Flags().GetString(netflowStatusFlagName)
		if err != nil {
			return err, false
		}
		m.NetflowStatus = netflowStatusFlagValue

		retAdded = true
	}

	return nil, retAdded
}