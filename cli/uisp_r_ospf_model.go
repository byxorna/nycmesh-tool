// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for UispROspf

// register flags to command
func registerModelUispROspfFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUispROspfRedistributeConnected(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUispROspfRedistributeDefaultRoute(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUispROspfRedistributeStatic(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUispROspfRouter(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUispROspfRedistributeConnected(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var redistributeConnectedFlagName string
	if cmdPrefix == "" {
		redistributeConnectedFlagName = "redistributeConnected"
	} else {
		redistributeConnectedFlagName = fmt.Sprintf("%v.redistributeConnected", cmdPrefix)
	}

	if err := registerModelRedistributeConnectedFlags(depth+1, redistributeConnectedFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerUispROspfRedistributeDefaultRoute(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var redistributeDefaultRouteFlagName string
	if cmdPrefix == "" {
		redistributeDefaultRouteFlagName = "redistributeDefaultRoute"
	} else {
		redistributeDefaultRouteFlagName = fmt.Sprintf("%v.redistributeDefaultRoute", cmdPrefix)
	}

	if err := registerModelRedistributeDefaultRouteFlags(depth+1, redistributeDefaultRouteFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerUispROspfRedistributeStatic(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var redistributeStaticFlagName string
	if cmdPrefix == "" {
		redistributeStaticFlagName = "redistributeStatic"
	} else {
		redistributeStaticFlagName = fmt.Sprintf("%v.redistributeStatic", cmdPrefix)
	}

	if err := registerModelRedistributeStaticFlags(depth+1, redistributeStaticFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerUispROspfRouter(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	routerDescription := ``

	var routerFlagName string
	if cmdPrefix == "" {
		routerFlagName = "router"
	} else {
		routerFlagName = fmt.Sprintf("%v.router", cmdPrefix)
	}

	var routerFlagDefault string

	_ = cmd.PersistentFlags().String(routerFlagName, routerFlagDefault, routerDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUispROspfFlags(depth int, m *models.UispROspf, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, redistributeConnectedAdded := retrieveUispROspfRedistributeConnectedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || redistributeConnectedAdded

	err, redistributeDefaultRouteAdded := retrieveUispROspfRedistributeDefaultRouteFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || redistributeDefaultRouteAdded

	err, redistributeStaticAdded := retrieveUispROspfRedistributeStaticFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || redistributeStaticAdded

	err, routerAdded := retrieveUispROspfRouterFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || routerAdded

	return nil, retAdded
}

func retrieveUispROspfRedistributeConnectedFlags(depth int, m *models.UispROspf, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	redistributeConnectedFlagName := fmt.Sprintf("%v.redistributeConnected", cmdPrefix)
	if cmd.Flags().Changed(redistributeConnectedFlagName) {
		// info: complex object redistributeConnected RedistributeConnected is retrieved outside this Changed() block
	}
	redistributeConnectedFlagValue := m.RedistributeConnected
	if swag.IsZero(redistributeConnectedFlagValue) {
		redistributeConnectedFlagValue = &models.RedistributeConnected{}
	}

	err, redistributeConnectedAdded := retrieveModelRedistributeConnectedFlags(depth+1, redistributeConnectedFlagValue, redistributeConnectedFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || redistributeConnectedAdded
	if redistributeConnectedAdded {
		m.RedistributeConnected = redistributeConnectedFlagValue
	}

	return nil, retAdded
}

func retrieveUispROspfRedistributeDefaultRouteFlags(depth int, m *models.UispROspf, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	redistributeDefaultRouteFlagName := fmt.Sprintf("%v.redistributeDefaultRoute", cmdPrefix)
	if cmd.Flags().Changed(redistributeDefaultRouteFlagName) {
		// info: complex object redistributeDefaultRoute RedistributeDefaultRoute is retrieved outside this Changed() block
	}
	redistributeDefaultRouteFlagValue := m.RedistributeDefaultRoute
	if swag.IsZero(redistributeDefaultRouteFlagValue) {
		redistributeDefaultRouteFlagValue = &models.RedistributeDefaultRoute{}
	}

	err, redistributeDefaultRouteAdded := retrieveModelRedistributeDefaultRouteFlags(depth+1, redistributeDefaultRouteFlagValue, redistributeDefaultRouteFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || redistributeDefaultRouteAdded
	if redistributeDefaultRouteAdded {
		m.RedistributeDefaultRoute = redistributeDefaultRouteFlagValue
	}

	return nil, retAdded
}

func retrieveUispROspfRedistributeStaticFlags(depth int, m *models.UispROspf, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	redistributeStaticFlagName := fmt.Sprintf("%v.redistributeStatic", cmdPrefix)
	if cmd.Flags().Changed(redistributeStaticFlagName) {
		// info: complex object redistributeStatic RedistributeStatic is retrieved outside this Changed() block
	}
	redistributeStaticFlagValue := m.RedistributeStatic
	if swag.IsZero(redistributeStaticFlagValue) {
		redistributeStaticFlagValue = &models.RedistributeStatic{}
	}

	err, redistributeStaticAdded := retrieveModelRedistributeStaticFlags(depth+1, redistributeStaticFlagValue, redistributeStaticFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || redistributeStaticAdded
	if redistributeStaticAdded {
		m.RedistributeStatic = redistributeStaticFlagValue
	}

	return nil, retAdded
}

func retrieveUispROspfRouterFlags(depth int, m *models.UispROspf, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	routerFlagName := fmt.Sprintf("%v.router", cmdPrefix)
	if cmd.Flags().Changed(routerFlagName) {

		var routerFlagName string
		if cmdPrefix == "" {
			routerFlagName = "router"
		} else {
			routerFlagName = fmt.Sprintf("%v.router", cmdPrefix)
		}

		routerFlagValue, err := cmd.Flags().GetString(routerFlagName)
		if err != nil {
			return err, false
		}
		m.Router = routerFlagValue

		retAdded = true
	}

	return nil, retAdded
}
