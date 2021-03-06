// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for ImageOrder

// register flags to command
func registerModelImageOrderFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerImageOrderCurrentOrder(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageOrderNextOrder(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerImageOrderCurrentOrder(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	currentOrderDescription := `Required. `

	var currentOrderFlagName string
	if cmdPrefix == "" {
		currentOrderFlagName = "currentOrder"
	} else {
		currentOrderFlagName = fmt.Sprintf("%v.currentOrder", cmdPrefix)
	}

	var currentOrderFlagDefault float64

	_ = cmd.PersistentFlags().Float64(currentOrderFlagName, currentOrderFlagDefault, currentOrderDescription)

	return nil
}

func registerImageOrderNextOrder(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nextOrderDescription := `Required. `

	var nextOrderFlagName string
	if cmdPrefix == "" {
		nextOrderFlagName = "nextOrder"
	} else {
		nextOrderFlagName = fmt.Sprintf("%v.nextOrder", cmdPrefix)
	}

	var nextOrderFlagDefault float64

	_ = cmd.PersistentFlags().Float64(nextOrderFlagName, nextOrderFlagDefault, nextOrderDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelImageOrderFlags(depth int, m *models.ImageOrder, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, currentOrderAdded := retrieveImageOrderCurrentOrderFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || currentOrderAdded

	err, nextOrderAdded := retrieveImageOrderNextOrderFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nextOrderAdded

	return nil, retAdded
}

func retrieveImageOrderCurrentOrderFlags(depth int, m *models.ImageOrder, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	currentOrderFlagName := fmt.Sprintf("%v.currentOrder", cmdPrefix)
	if cmd.Flags().Changed(currentOrderFlagName) {

		var currentOrderFlagName string
		if cmdPrefix == "" {
			currentOrderFlagName = "currentOrder"
		} else {
			currentOrderFlagName = fmt.Sprintf("%v.currentOrder", cmdPrefix)
		}

		currentOrderFlagValue, err := cmd.Flags().GetFloat64(currentOrderFlagName)
		if err != nil {
			return err, false
		}
		m.CurrentOrder = &currentOrderFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageOrderNextOrderFlags(depth int, m *models.ImageOrder, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nextOrderFlagName := fmt.Sprintf("%v.nextOrder", cmdPrefix)
	if cmd.Flags().Changed(nextOrderFlagName) {

		var nextOrderFlagName string
		if cmdPrefix == "" {
			nextOrderFlagName = "nextOrder"
		} else {
			nextOrderFlagName = fmt.Sprintf("%v.nextOrder", cmdPrefix)
		}

		nextOrderFlagValue, err := cmd.Flags().GetFloat64(nextOrderFlagName)
		if err != nil {
			return err, false
		}
		m.NextOrder = &nextOrderFlagValue

		retAdded = true
	}

	return nil, retAdded
}
