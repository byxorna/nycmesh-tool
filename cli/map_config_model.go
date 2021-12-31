// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for MapConfig

// register flags to command
func registerModelMapConfigFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerMapConfigShowClientSites(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerMapConfigShowClientSites(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	showClientSitesDescription := ``

	var showClientSitesFlagName string
	if cmdPrefix == "" {
		showClientSitesFlagName = "showClientSites"
	} else {
		showClientSitesFlagName = fmt.Sprintf("%v.showClientSites", cmdPrefix)
	}

	var showClientSitesFlagDefault bool

	_ = cmd.PersistentFlags().Bool(showClientSitesFlagName, showClientSitesFlagDefault, showClientSitesDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelMapConfigFlags(depth int, m *models.MapConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, showClientSitesAdded := retrieveMapConfigShowClientSitesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || showClientSitesAdded

	return nil, retAdded
}

func retrieveMapConfigShowClientSitesFlags(depth int, m *models.MapConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	showClientSitesFlagName := fmt.Sprintf("%v.showClientSites", cmdPrefix)
	if cmd.Flags().Changed(showClientSitesFlagName) {

		var showClientSitesFlagName string
		if cmdPrefix == "" {
			showClientSitesFlagName = "showClientSites"
		} else {
			showClientSitesFlagName = fmt.Sprintf("%v.showClientSites", cmdPrefix)
		}

		showClientSitesFlagValue, err := cmd.Flags().GetBool(showClientSitesFlagName)
		if err != nil {
			return err, false
		}
		m.ShowClientSites = showClientSitesFlagValue

		retAdded = true
	}

	return nil, retAdded
}
