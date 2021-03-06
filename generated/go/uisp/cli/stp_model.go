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

// Schema cli for Stp

// register flags to command
func registerModelStpFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerStpEdgePort(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStpEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStpPathCost(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStpPortPriority(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerStpEdgePort(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	edgePortDescription := `Enum: ["auto","enable","disable"]. `

	var edgePortFlagName string
	if cmdPrefix == "" {
		edgePortFlagName = "edgePort"
	} else {
		edgePortFlagName = fmt.Sprintf("%v.edgePort", cmdPrefix)
	}

	var edgePortFlagDefault string

	_ = cmd.PersistentFlags().String(edgePortFlagName, edgePortFlagDefault, edgePortDescription)

	if err := cmd.RegisterFlagCompletionFunc(edgePortFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["auto","enable","disable"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerStpEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerStpPathCost(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	pathCostDescription := ``

	var pathCostFlagName string
	if cmdPrefix == "" {
		pathCostFlagName = "pathCost"
	} else {
		pathCostFlagName = fmt.Sprintf("%v.pathCost", cmdPrefix)
	}

	var pathCostFlagDefault float64

	_ = cmd.PersistentFlags().Float64(pathCostFlagName, pathCostFlagDefault, pathCostDescription)

	return nil
}

func registerStpPortPriority(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	portPriorityDescription := ``

	var portPriorityFlagName string
	if cmdPrefix == "" {
		portPriorityFlagName = "portPriority"
	} else {
		portPriorityFlagName = fmt.Sprintf("%v.portPriority", cmdPrefix)
	}

	var portPriorityFlagDefault float64

	_ = cmd.PersistentFlags().Float64(portPriorityFlagName, portPriorityFlagDefault, portPriorityDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelStpFlags(depth int, m *models.Stp, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, edgePortAdded := retrieveStpEdgePortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || edgePortAdded

	err, enabledAdded := retrieveStpEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || enabledAdded

	err, pathCostAdded := retrieveStpPathCostFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pathCostAdded

	err, portPriorityAdded := retrieveStpPortPriorityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || portPriorityAdded

	return nil, retAdded
}

func retrieveStpEdgePortFlags(depth int, m *models.Stp, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	edgePortFlagName := fmt.Sprintf("%v.edgePort", cmdPrefix)
	if cmd.Flags().Changed(edgePortFlagName) {

		var edgePortFlagName string
		if cmdPrefix == "" {
			edgePortFlagName = "edgePort"
		} else {
			edgePortFlagName = fmt.Sprintf("%v.edgePort", cmdPrefix)
		}

		edgePortFlagValue, err := cmd.Flags().GetString(edgePortFlagName)
		if err != nil {
			return err, false
		}
		m.EdgePort = edgePortFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStpEnabledFlags(depth int, m *models.Stp, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveStpPathCostFlags(depth int, m *models.Stp, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pathCostFlagName := fmt.Sprintf("%v.pathCost", cmdPrefix)
	if cmd.Flags().Changed(pathCostFlagName) {

		var pathCostFlagName string
		if cmdPrefix == "" {
			pathCostFlagName = "pathCost"
		} else {
			pathCostFlagName = fmt.Sprintf("%v.pathCost", cmdPrefix)
		}

		pathCostFlagValue, err := cmd.Flags().GetFloat64(pathCostFlagName)
		if err != nil {
			return err, false
		}
		m.PathCost = pathCostFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStpPortPriorityFlags(depth int, m *models.Stp, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	portPriorityFlagName := fmt.Sprintf("%v.portPriority", cmdPrefix)
	if cmd.Flags().Changed(portPriorityFlagName) {

		var portPriorityFlagName string
		if cmdPrefix == "" {
			portPriorityFlagName = "portPriority"
		} else {
			portPriorityFlagName = fmt.Sprintf("%v.portPriority", cmdPrefix)
		}

		portPriorityFlagValue, err := cmd.Flags().GetFloat64(portPriorityFlagName)
		if err != nil {
			return err, false
		}
		m.PortPriority = portPriorityFlagValue

		retAdded = true
	}

	return nil, retAdded
}
