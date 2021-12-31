// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"
  "github.com/byxorna/nycmesh-tool/models"

	"github.com/spf13/cobra"
)

// Schema cli for Counts

// register flags to command
func registerModelCountsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCountsEndpoint(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCountsSite(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCountsUser(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCountsEndpoint(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var endpointFlagName string
	if cmdPrefix == "" {
		endpointFlagName = "endpoint"
	} else {
		endpointFlagName = fmt.Sprintf("%v.endpoint", cmdPrefix)
	}

	if err := registerModelEndpointFlags(depth+1, endpointFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerCountsSite(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var siteFlagName string
	if cmdPrefix == "" {
		siteFlagName = "site"
	} else {
		siteFlagName = fmt.Sprintf("%v.site", cmdPrefix)
	}

	if err := registerModelSite3Flags(depth+1, siteFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerCountsUser(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var userFlagName string
	if cmdPrefix == "" {
		userFlagName = "user"
	} else {
		userFlagName = fmt.Sprintf("%v.user", cmdPrefix)
	}

	if err := registerModelUser1Flags(depth+1, userFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCountsFlags(depth int, m *models.Counts, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, endpointAdded := retrieveCountsEndpointFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || endpointAdded

	err, siteAdded := retrieveCountsSiteFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || siteAdded

	err, userAdded := retrieveCountsUserFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || userAdded

	return nil, retAdded
}

func retrieveCountsEndpointFlags(depth int, m *models.Counts, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	endpointFlagName := fmt.Sprintf("%v.endpoint", cmdPrefix)
	if cmd.Flags().Changed(endpointFlagName) {
		// info: complex object endpoint Endpoint is retrieved outside this Changed() block
	}
	endpointFlagValue := m.Endpoint
	if swag.IsZero(endpointFlagValue) {
		endpointFlagValue = &models.Endpoint{}
	}

	err, endpointAdded := retrieveModelEndpointFlags(depth+1, endpointFlagValue, endpointFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || endpointAdded
	if endpointAdded {
		m.Endpoint = endpointFlagValue
	}

	return nil, retAdded
}

func retrieveCountsSiteFlags(depth int, m *models.Counts, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	siteFlagName := fmt.Sprintf("%v.site", cmdPrefix)
	if cmd.Flags().Changed(siteFlagName) {
		// info: complex object site Site3 is retrieved outside this Changed() block
	}
	siteFlagValue := m.Site
	if swag.IsZero(siteFlagValue) {
		siteFlagValue = &models.Site3{}
	}

	err, siteAdded := retrieveModelSite3Flags(depth+1, siteFlagValue, siteFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || siteAdded
	if siteAdded {
		m.Site = siteFlagValue
	}

	return nil, retAdded
}

func retrieveCountsUserFlags(depth int, m *models.Counts, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	userFlagName := fmt.Sprintf("%v.user", cmdPrefix)
	if cmd.Flags().Changed(userFlagName) {
		// info: complex object user User1 is retrieved outside this Changed() block
	}
	userFlagValue := m.User
	if swag.IsZero(userFlagValue) {
		userFlagValue = &models.User1{}
	}

	err, userAdded := retrieveModelUser1Flags(depth+1, userFlagValue, userFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || userAdded
	if userAdded {
		m.User = userFlagValue
	}

	return nil, retAdded
}