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

// Schema cli for SiteAccessGroupSingleSite

// register flags to command
func registerModelSiteAccessGroupSingleSiteFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSiteAccessGroupSingleSiteID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteAccessGroupSingleSiteIsReadOnly(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteAccessGroupSingleSiteName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteAccessGroupSingleSiteType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSiteAccessGroupSingleSiteID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

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

func registerSiteAccessGroupSingleSiteIsReadOnly(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	isReadOnlyDescription := `Required. If true, access to site is read-only. If false, access to site is read-write.`

	var isReadOnlyFlagName string
	if cmdPrefix == "" {
		isReadOnlyFlagName = "isReadOnly"
	} else {
		isReadOnlyFlagName = fmt.Sprintf("%v.isReadOnly", cmdPrefix)
	}

	var isReadOnlyFlagDefault bool

	_ = cmd.PersistentFlags().Bool(isReadOnlyFlagName, isReadOnlyFlagDefault, isReadOnlyDescription)

	return nil
}

func registerSiteAccessGroupSingleSiteName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Required. Name of the site.`

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerSiteAccessGroupSingleSiteType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["site","endpoint","client"]. Required. Type of the site.`

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	if err := cmd.RegisterFlagCompletionFunc(typeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["site","endpoint","client"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSiteAccessGroupSingleSiteFlags(depth int, m *models.SiteAccessGroupSingleSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, idAdded := retrieveSiteAccessGroupSingleSiteIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, isReadOnlyAdded := retrieveSiteAccessGroupSingleSiteIsReadOnlyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || isReadOnlyAdded

	err, nameAdded := retrieveSiteAccessGroupSingleSiteNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, typeAdded := retrieveSiteAccessGroupSingleSiteTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveSiteAccessGroupSingleSiteIDFlags(depth int, m *models.SiteAccessGroupSingleSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

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
		m.ID = &idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSiteAccessGroupSingleSiteIsReadOnlyFlags(depth int, m *models.SiteAccessGroupSingleSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	isReadOnlyFlagName := fmt.Sprintf("%v.isReadOnly", cmdPrefix)
	if cmd.Flags().Changed(isReadOnlyFlagName) {

		var isReadOnlyFlagName string
		if cmdPrefix == "" {
			isReadOnlyFlagName = "isReadOnly"
		} else {
			isReadOnlyFlagName = fmt.Sprintf("%v.isReadOnly", cmdPrefix)
		}

		isReadOnlyFlagValue, err := cmd.Flags().GetBool(isReadOnlyFlagName)
		if err != nil {
			return err, false
		}
		m.IsReadOnly = &isReadOnlyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSiteAccessGroupSingleSiteNameFlags(depth int, m *models.SiteAccessGroupSingleSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = &nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSiteAccessGroupSingleSiteTypeFlags(depth int, m *models.SiteAccessGroupSingleSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	typeFlagName := fmt.Sprintf("%v.type", cmdPrefix)
	if cmd.Flags().Changed(typeFlagName) {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = &typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
