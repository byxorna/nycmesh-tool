// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for SiteBinding

// register flags to command
func registerModelSiteBindingFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSiteBindingUcrmID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSiteBindingUcrmID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ucrmIdDescription := `Required. `

	var ucrmIdFlagName string
	if cmdPrefix == "" {
		ucrmIdFlagName = "ucrmId"
	} else {
		ucrmIdFlagName = fmt.Sprintf("%v.ucrmId", cmdPrefix)
	}

	var ucrmIdFlagDefault string

	_ = cmd.PersistentFlags().String(ucrmIdFlagName, ucrmIdFlagDefault, ucrmIdDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSiteBindingFlags(depth int, m *models.SiteBinding, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, ucrmIdAdded := retrieveSiteBindingUcrmIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ucrmIdAdded

	return nil, retAdded
}

func retrieveSiteBindingUcrmIDFlags(depth int, m *models.SiteBinding, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ucrmIdFlagName := fmt.Sprintf("%v.ucrmId", cmdPrefix)
	if cmd.Flags().Changed(ucrmIdFlagName) {

		var ucrmIdFlagName string
		if cmdPrefix == "" {
			ucrmIdFlagName = "ucrmId"
		} else {
			ucrmIdFlagName = fmt.Sprintf("%v.ucrmId", cmdPrefix)
		}

		ucrmIdFlagValue, err := cmd.Flags().GetString(ucrmIdFlagName)
		if err != nil {
			return err, false
		}
		m.UcrmID = &ucrmIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}
