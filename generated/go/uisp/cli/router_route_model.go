// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for RouterRoute

// register flags to command
func registerModelRouterRouteFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerRouterRouteDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRouterRouteDestination(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRouterRouteDistance(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRouterRouteEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRouterRouteFib(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRouterRouteGateway(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRouterRouteGatewayStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRouterRouteInterface(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRouterRouteNextHop(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRouterRouteSelected(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRouterRouteStaticType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRouterRouteType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerRouterRouteDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := ``

	var descriptionFlagName string
	if cmdPrefix == "" {
		descriptionFlagName = "description"
	} else {
		descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
	}

	var descriptionFlagDefault string

	_ = cmd.PersistentFlags().String(descriptionFlagName, descriptionFlagDefault, descriptionDescription)

	return nil
}

func registerRouterRouteDestination(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	destinationDescription := ``

	var destinationFlagName string
	if cmdPrefix == "" {
		destinationFlagName = "destination"
	} else {
		destinationFlagName = fmt.Sprintf("%v.destination", cmdPrefix)
	}

	var destinationFlagDefault string

	_ = cmd.PersistentFlags().String(destinationFlagName, destinationFlagDefault, destinationDescription)

	return nil
}

func registerRouterRouteDistance(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	distanceDescription := ``

	var distanceFlagName string
	if cmdPrefix == "" {
		distanceFlagName = "distance"
	} else {
		distanceFlagName = fmt.Sprintf("%v.distance", cmdPrefix)
	}

	var distanceFlagDefault float64

	_ = cmd.PersistentFlags().Float64(distanceFlagName, distanceFlagDefault, distanceDescription)

	return nil
}

func registerRouterRouteEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	enabledDescription := ``

	var enabledFlagName string
	if cmdPrefix == "" {
		enabledFlagName = "enabled"
	} else {
		enabledFlagName = fmt.Sprintf("%v.enabled", cmdPrefix)
	}

	var enabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(enabledFlagName, enabledFlagDefault, enabledDescription)

	return nil
}

func registerRouterRouteFib(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	fibDescription := ``

	var fibFlagName string
	if cmdPrefix == "" {
		fibFlagName = "fib"
	} else {
		fibFlagName = fmt.Sprintf("%v.fib", cmdPrefix)
	}

	var fibFlagDefault bool

	_ = cmd.PersistentFlags().Bool(fibFlagName, fibFlagDefault, fibDescription)

	return nil
}

func registerRouterRouteGateway(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	gatewayDescription := ``

	var gatewayFlagName string
	if cmdPrefix == "" {
		gatewayFlagName = "gateway"
	} else {
		gatewayFlagName = fmt.Sprintf("%v.gateway", cmdPrefix)
	}

	var gatewayFlagDefault string

	_ = cmd.PersistentFlags().String(gatewayFlagName, gatewayFlagDefault, gatewayDescription)

	return nil
}

func registerRouterRouteGatewayStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	gatewayStatusDescription := ``

	var gatewayStatusFlagName string
	if cmdPrefix == "" {
		gatewayStatusFlagName = "gatewayStatus"
	} else {
		gatewayStatusFlagName = fmt.Sprintf("%v.gatewayStatus", cmdPrefix)
	}

	var gatewayStatusFlagDefault string

	_ = cmd.PersistentFlags().String(gatewayStatusFlagName, gatewayStatusFlagDefault, gatewayStatusDescription)

	return nil
}

func registerRouterRouteInterface(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	interfaceDescription := ``

	var interfaceFlagName string
	if cmdPrefix == "" {
		interfaceFlagName = "interface"
	} else {
		interfaceFlagName = fmt.Sprintf("%v.interface", cmdPrefix)
	}

	var interfaceFlagDefault string

	_ = cmd.PersistentFlags().String(interfaceFlagName, interfaceFlagDefault, interfaceDescription)

	return nil
}

func registerRouterRouteNextHop(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nextHopDescription := ``

	var nextHopFlagName string
	if cmdPrefix == "" {
		nextHopFlagName = "nextHop"
	} else {
		nextHopFlagName = fmt.Sprintf("%v.nextHop", cmdPrefix)
	}

	var nextHopFlagDefault string

	_ = cmd.PersistentFlags().String(nextHopFlagName, nextHopFlagDefault, nextHopDescription)

	return nil
}

func registerRouterRouteSelected(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	selectedDescription := ``

	var selectedFlagName string
	if cmdPrefix == "" {
		selectedFlagName = "selected"
	} else {
		selectedFlagName = fmt.Sprintf("%v.selected", cmdPrefix)
	}

	var selectedFlagDefault bool

	_ = cmd.PersistentFlags().Bool(selectedFlagName, selectedFlagDefault, selectedDescription)

	return nil
}

func registerRouterRouteStaticType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	staticTypeDescription := ``

	var staticTypeFlagName string
	if cmdPrefix == "" {
		staticTypeFlagName = "staticType"
	} else {
		staticTypeFlagName = fmt.Sprintf("%v.staticType", cmdPrefix)
	}

	var staticTypeFlagDefault string

	_ = cmd.PersistentFlags().String(staticTypeFlagName, staticTypeFlagDefault, staticTypeDescription)

	return nil
}

func registerRouterRouteType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := ``

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelRouterRouteFlags(depth int, m *models.RouterRoute, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, descriptionAdded := retrieveRouterRouteDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, destinationAdded := retrieveRouterRouteDestinationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || destinationAdded

	err, distanceAdded := retrieveRouterRouteDistanceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || distanceAdded

	err, enabledAdded := retrieveRouterRouteEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || enabledAdded

	err, fibAdded := retrieveRouterRouteFibFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || fibAdded

	err, gatewayAdded := retrieveRouterRouteGatewayFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || gatewayAdded

	err, gatewayStatusAdded := retrieveRouterRouteGatewayStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || gatewayStatusAdded

	err, interfaceAdded := retrieveRouterRouteInterfaceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || interfaceAdded

	err, nextHopAdded := retrieveRouterRouteNextHopFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nextHopAdded

	err, selectedAdded := retrieveRouterRouteSelectedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || selectedAdded

	err, staticTypeAdded := retrieveRouterRouteStaticTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || staticTypeAdded

	err, typeAdded := retrieveRouterRouteTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveRouterRouteDescriptionFlags(depth int, m *models.RouterRoute, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	descriptionFlagName := fmt.Sprintf("%v.description", cmdPrefix)
	if cmd.Flags().Changed(descriptionFlagName) {

		var descriptionFlagName string
		if cmdPrefix == "" {
			descriptionFlagName = "description"
		} else {
			descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
		}

		descriptionFlagValue, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			return err, false
		}
		m.Description = descriptionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRouterRouteDestinationFlags(depth int, m *models.RouterRoute, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	destinationFlagName := fmt.Sprintf("%v.destination", cmdPrefix)
	if cmd.Flags().Changed(destinationFlagName) {

		var destinationFlagName string
		if cmdPrefix == "" {
			destinationFlagName = "destination"
		} else {
			destinationFlagName = fmt.Sprintf("%v.destination", cmdPrefix)
		}

		destinationFlagValue, err := cmd.Flags().GetString(destinationFlagName)
		if err != nil {
			return err, false
		}
		m.Destination = destinationFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRouterRouteDistanceFlags(depth int, m *models.RouterRoute, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	distanceFlagName := fmt.Sprintf("%v.distance", cmdPrefix)
	if cmd.Flags().Changed(distanceFlagName) {

		var distanceFlagName string
		if cmdPrefix == "" {
			distanceFlagName = "distance"
		} else {
			distanceFlagName = fmt.Sprintf("%v.distance", cmdPrefix)
		}

		distanceFlagValue, err := cmd.Flags().GetFloat64(distanceFlagName)
		if err != nil {
			return err, false
		}
		m.Distance = &distanceFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRouterRouteEnabledFlags(depth int, m *models.RouterRoute, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	enabledFlagName := fmt.Sprintf("%v.enabled", cmdPrefix)
	if cmd.Flags().Changed(enabledFlagName) {

		var enabledFlagName string
		if cmdPrefix == "" {
			enabledFlagName = "enabled"
		} else {
			enabledFlagName = fmt.Sprintf("%v.enabled", cmdPrefix)
		}

		enabledFlagValue, err := cmd.Flags().GetBool(enabledFlagName)
		if err != nil {
			return err, false
		}
		m.Enabled = enabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRouterRouteFibFlags(depth int, m *models.RouterRoute, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	fibFlagName := fmt.Sprintf("%v.fib", cmdPrefix)
	if cmd.Flags().Changed(fibFlagName) {

		var fibFlagName string
		if cmdPrefix == "" {
			fibFlagName = "fib"
		} else {
			fibFlagName = fmt.Sprintf("%v.fib", cmdPrefix)
		}

		fibFlagValue, err := cmd.Flags().GetBool(fibFlagName)
		if err != nil {
			return err, false
		}
		m.Fib = fibFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRouterRouteGatewayFlags(depth int, m *models.RouterRoute, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	gatewayFlagName := fmt.Sprintf("%v.gateway", cmdPrefix)
	if cmd.Flags().Changed(gatewayFlagName) {

		var gatewayFlagName string
		if cmdPrefix == "" {
			gatewayFlagName = "gateway"
		} else {
			gatewayFlagName = fmt.Sprintf("%v.gateway", cmdPrefix)
		}

		gatewayFlagValue, err := cmd.Flags().GetString(gatewayFlagName)
		if err != nil {
			return err, false
		}
		m.Gateway = gatewayFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRouterRouteGatewayStatusFlags(depth int, m *models.RouterRoute, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	gatewayStatusFlagName := fmt.Sprintf("%v.gatewayStatus", cmdPrefix)
	if cmd.Flags().Changed(gatewayStatusFlagName) {

		var gatewayStatusFlagName string
		if cmdPrefix == "" {
			gatewayStatusFlagName = "gatewayStatus"
		} else {
			gatewayStatusFlagName = fmt.Sprintf("%v.gatewayStatus", cmdPrefix)
		}

		gatewayStatusFlagValue, err := cmd.Flags().GetString(gatewayStatusFlagName)
		if err != nil {
			return err, false
		}
		m.GatewayStatus = gatewayStatusFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRouterRouteInterfaceFlags(depth int, m *models.RouterRoute, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	interfaceFlagName := fmt.Sprintf("%v.interface", cmdPrefix)
	if cmd.Flags().Changed(interfaceFlagName) {

		var interfaceFlagName string
		if cmdPrefix == "" {
			interfaceFlagName = "interface"
		} else {
			interfaceFlagName = fmt.Sprintf("%v.interface", cmdPrefix)
		}

		interfaceFlagValue, err := cmd.Flags().GetString(interfaceFlagName)
		if err != nil {
			return err, false
		}
		m.Interface = interfaceFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRouterRouteNextHopFlags(depth int, m *models.RouterRoute, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nextHopFlagName := fmt.Sprintf("%v.nextHop", cmdPrefix)
	if cmd.Flags().Changed(nextHopFlagName) {

		var nextHopFlagName string
		if cmdPrefix == "" {
			nextHopFlagName = "nextHop"
		} else {
			nextHopFlagName = fmt.Sprintf("%v.nextHop", cmdPrefix)
		}

		nextHopFlagValue, err := cmd.Flags().GetString(nextHopFlagName)
		if err != nil {
			return err, false
		}
		m.NextHop = nextHopFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRouterRouteSelectedFlags(depth int, m *models.RouterRoute, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	selectedFlagName := fmt.Sprintf("%v.selected", cmdPrefix)
	if cmd.Flags().Changed(selectedFlagName) {

		var selectedFlagName string
		if cmdPrefix == "" {
			selectedFlagName = "selected"
		} else {
			selectedFlagName = fmt.Sprintf("%v.selected", cmdPrefix)
		}

		selectedFlagValue, err := cmd.Flags().GetBool(selectedFlagName)
		if err != nil {
			return err, false
		}
		m.Selected = selectedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRouterRouteStaticTypeFlags(depth int, m *models.RouterRoute, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	staticTypeFlagName := fmt.Sprintf("%v.staticType", cmdPrefix)
	if cmd.Flags().Changed(staticTypeFlagName) {

		var staticTypeFlagName string
		if cmdPrefix == "" {
			staticTypeFlagName = "staticType"
		} else {
			staticTypeFlagName = fmt.Sprintf("%v.staticType", cmdPrefix)
		}

		staticTypeFlagValue, err := cmd.Flags().GetString(staticTypeFlagName)
		if err != nil {
			return err, false
		}
		m.StaticType = staticTypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRouterRouteTypeFlags(depth int, m *models.RouterRoute, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Type = typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
