// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for NetworkStatisticInterval

// register flags to command
func registerModelNetworkStatisticIntervalFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerNetworkStatisticIntervalEnd(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkStatisticIntervalStart(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerNetworkStatisticIntervalEnd(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	endDescription := ``

	var endFlagName string
	if cmdPrefix == "" {
		endFlagName = "end"
	} else {
		endFlagName = fmt.Sprintf("%v.end", cmdPrefix)
	}

	var endFlagDefault string

	_ = cmd.PersistentFlags().String(endFlagName, endFlagDefault, endDescription)

	return nil
}

func registerNetworkStatisticIntervalStart(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	startDescription := ``

	var startFlagName string
	if cmdPrefix == "" {
		startFlagName = "start"
	} else {
		startFlagName = fmt.Sprintf("%v.start", cmdPrefix)
	}

	var startFlagDefault string

	_ = cmd.PersistentFlags().String(startFlagName, startFlagDefault, startDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelNetworkStatisticIntervalFlags(depth int, m *models.NetworkStatisticInterval, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, endAdded := retrieveNetworkStatisticIntervalEndFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || endAdded

	err, startAdded := retrieveNetworkStatisticIntervalStartFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || startAdded

	return nil, retAdded
}

func retrieveNetworkStatisticIntervalEndFlags(depth int, m *models.NetworkStatisticInterval, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	endFlagName := fmt.Sprintf("%v.end", cmdPrefix)
	if cmd.Flags().Changed(endFlagName) {

		var endFlagName string
		if cmdPrefix == "" {
			endFlagName = "end"
		} else {
			endFlagName = fmt.Sprintf("%v.end", cmdPrefix)
		}

		endFlagValue, err := cmd.Flags().GetString(endFlagName)
		if err != nil {
			return err, false
		}
		m.End = endFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveNetworkStatisticIntervalStartFlags(depth int, m *models.NetworkStatisticInterval, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	startFlagName := fmt.Sprintf("%v.start", cmdPrefix)
	if cmd.Flags().Changed(startFlagName) {

		var startFlagName string
		if cmdPrefix == "" {
			startFlagName = "start"
		} else {
			startFlagName = fmt.Sprintf("%v.start", cmdPrefix)
		}

		startFlagValue, err := cmd.Flags().GetString(startFlagName)
		if err != nil {
			return err, false
		}
		m.Start = startFlagValue

		retAdded = true
	}

	return nil, retAdded
}
