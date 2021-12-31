// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for SiteAccessGroup

// register flags to command
func registerModelSiteAccessGroupFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSiteAccessGroupCounts(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteAccessGroupID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteAccessGroupIsAuto(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteAccessGroupIsInternal(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteAccessGroupName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteAccessGroupWhitelist(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSiteAccessGroupCounts(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var countsFlagName string
	if cmdPrefix == "" {
		countsFlagName = "counts"
	} else {
		countsFlagName = fmt.Sprintf("%v.counts", cmdPrefix)
	}

	if err := registerModelCountsFlags(depth+1, countsFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerSiteAccessGroupID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `Access group id.`

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

func registerSiteAccessGroupIsAuto(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	isAutoDescription := `Required. If true the access group is automatically generated and cannot be edited.`

	var isAutoFlagName string
	if cmdPrefix == "" {
		isAutoFlagName = "isAuto"
	} else {
		isAutoFlagName = fmt.Sprintf("%v.isAuto", cmdPrefix)
	}

	var isAutoFlagDefault bool

	_ = cmd.PersistentFlags().Bool(isAutoFlagName, isAutoFlagDefault, isAutoDescription)

	return nil
}

func registerSiteAccessGroupIsInternal(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	isInternalDescription := `Required. If true the access group is internal. User should not be able to select it`

	var isInternalFlagName string
	if cmdPrefix == "" {
		isInternalFlagName = "isInternal"
	} else {
		isInternalFlagName = fmt.Sprintf("%v.isInternal", cmdPrefix)
	}

	var isInternalFlagDefault bool

	_ = cmd.PersistentFlags().Bool(isInternalFlagName, isInternalFlagDefault, isInternalDescription)

	return nil
}

func registerSiteAccessGroupName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Required. `

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

func registerSiteAccessGroupWhitelist(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: whitelist Whitelist array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSiteAccessGroupFlags(depth int, m *models.SiteAccessGroup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, countsAdded := retrieveSiteAccessGroupCountsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || countsAdded

	err, idAdded := retrieveSiteAccessGroupIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, isAutoAdded := retrieveSiteAccessGroupIsAutoFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || isAutoAdded

	err, isInternalAdded := retrieveSiteAccessGroupIsInternalFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || isInternalAdded

	err, nameAdded := retrieveSiteAccessGroupNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, whitelistAdded := retrieveSiteAccessGroupWhitelistFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || whitelistAdded

	return nil, retAdded
}

func retrieveSiteAccessGroupCountsFlags(depth int, m *models.SiteAccessGroup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	countsFlagName := fmt.Sprintf("%v.counts", cmdPrefix)
	if cmd.Flags().Changed(countsFlagName) {
		// info: complex object counts Counts is retrieved outside this Changed() block
	}
	countsFlagValue := m.Counts
	if swag.IsZero(countsFlagValue) {
		countsFlagValue = &models.Counts{}
	}

	err, countsAdded := retrieveModelCountsFlags(depth+1, countsFlagValue, countsFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || countsAdded
	if countsAdded {
		m.Counts = countsFlagValue
	}

	return nil, retAdded
}

func retrieveSiteAccessGroupIDFlags(depth int, m *models.SiteAccessGroup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSiteAccessGroupIsAutoFlags(depth int, m *models.SiteAccessGroup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	isAutoFlagName := fmt.Sprintf("%v.isAuto", cmdPrefix)
	if cmd.Flags().Changed(isAutoFlagName) {

		var isAutoFlagName string
		if cmdPrefix == "" {
			isAutoFlagName = "isAuto"
		} else {
			isAutoFlagName = fmt.Sprintf("%v.isAuto", cmdPrefix)
		}

		isAutoFlagValue, err := cmd.Flags().GetBool(isAutoFlagName)
		if err != nil {
			return err, false
		}
		m.IsAuto = &isAutoFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSiteAccessGroupIsInternalFlags(depth int, m *models.SiteAccessGroup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	isInternalFlagName := fmt.Sprintf("%v.isInternal", cmdPrefix)
	if cmd.Flags().Changed(isInternalFlagName) {

		var isInternalFlagName string
		if cmdPrefix == "" {
			isInternalFlagName = "isInternal"
		} else {
			isInternalFlagName = fmt.Sprintf("%v.isInternal", cmdPrefix)
		}

		isInternalFlagValue, err := cmd.Flags().GetBool(isInternalFlagName)
		if err != nil {
			return err, false
		}
		m.IsInternal = &isInternalFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSiteAccessGroupNameFlags(depth int, m *models.SiteAccessGroup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveSiteAccessGroupWhitelistFlags(depth int, m *models.SiteAccessGroup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	whitelistFlagName := fmt.Sprintf("%v.whitelist", cmdPrefix)
	if cmd.Flags().Changed(whitelistFlagName) {
		// warning: whitelist array type Whitelist is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
