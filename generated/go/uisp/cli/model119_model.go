// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for Model119

// register flags to command
func registerModelModel119Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel119RedistributeConnected(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel119RedistributeDefaultRoute(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel119RedistributeStatic(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel119Router(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel119RedistributeConnected(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModel119RedistributeDefaultRoute(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var redistributeDefaultRouteFlagName string
	if cmdPrefix == "" {
		redistributeDefaultRouteFlagName = "redistributeDefaultRoute"
	} else {
		redistributeDefaultRouteFlagName = fmt.Sprintf("%v.redistributeDefaultRoute", cmdPrefix)
	}

	if err := registerModelRedistributeDefaultRoute2Flags(depth+1, redistributeDefaultRouteFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel119RedistributeStatic(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModel119Router(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	routerDescription := `Required. `

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
func retrieveModelModel119Flags(depth int, m *models.Model119, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, redistributeConnectedAdded := retrieveModel119RedistributeConnectedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || redistributeConnectedAdded

	err, redistributeDefaultRouteAdded := retrieveModel119RedistributeDefaultRouteFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || redistributeDefaultRouteAdded

	err, redistributeStaticAdded := retrieveModel119RedistributeStaticFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || redistributeStaticAdded

	err, routerAdded := retrieveModel119RouterFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || routerAdded

	return nil, retAdded
}

func retrieveModel119RedistributeConnectedFlags(depth int, m *models.Model119, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel119RedistributeDefaultRouteFlags(depth int, m *models.Model119, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	redistributeDefaultRouteFlagName := fmt.Sprintf("%v.redistributeDefaultRoute", cmdPrefix)
	if cmd.Flags().Changed(redistributeDefaultRouteFlagName) {
		// info: complex object redistributeDefaultRoute RedistributeDefaultRoute2 is retrieved outside this Changed() block
	}
	redistributeDefaultRouteFlagValue := m.RedistributeDefaultRoute
	if swag.IsZero(redistributeDefaultRouteFlagValue) {
		redistributeDefaultRouteFlagValue = &models.RedistributeDefaultRoute2{}
	}

	err, redistributeDefaultRouteAdded := retrieveModelRedistributeDefaultRoute2Flags(depth+1, redistributeDefaultRouteFlagValue, redistributeDefaultRouteFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || redistributeDefaultRouteAdded
	if redistributeDefaultRouteAdded {
		m.RedistributeDefaultRoute = redistributeDefaultRouteFlagValue
	}

	return nil, retAdded
}

func retrieveModel119RedistributeStaticFlags(depth int, m *models.Model119, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel119RouterFlags(depth int, m *models.Model119, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Router = &routerFlagValue

		retAdded = true
	}

	return nil, retAdded
}