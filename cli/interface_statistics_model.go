// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for InterfaceStatistics

// register flags to command
func registerModelInterfaceStatisticsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerInterfaceStatisticsDropped(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInterfaceStatisticsErrors(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInterfaceStatisticsPoePower(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInterfaceStatisticsRxbytes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInterfaceStatisticsRxrate(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInterfaceStatisticsTxbytes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInterfaceStatisticsTxrate(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerInterfaceStatisticsDropped(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	droppedDescription := ``

	var droppedFlagName string
	if cmdPrefix == "" {
		droppedFlagName = "dropped"
	} else {
		droppedFlagName = fmt.Sprintf("%v.dropped", cmdPrefix)
	}

	var droppedFlagDefault float64

	_ = cmd.PersistentFlags().Float64(droppedFlagName, droppedFlagDefault, droppedDescription)

	return nil
}

func registerInterfaceStatisticsErrors(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	errorsDescription := ``

	var errorsFlagName string
	if cmdPrefix == "" {
		errorsFlagName = "errors"
	} else {
		errorsFlagName = fmt.Sprintf("%v.errors", cmdPrefix)
	}

	var errorsFlagDefault float64

	_ = cmd.PersistentFlags().Float64(errorsFlagName, errorsFlagDefault, errorsDescription)

	return nil
}

func registerInterfaceStatisticsPoePower(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	poePowerDescription := ``

	var poePowerFlagName string
	if cmdPrefix == "" {
		poePowerFlagName = "poePower"
	} else {
		poePowerFlagName = fmt.Sprintf("%v.poePower", cmdPrefix)
	}

	var poePowerFlagDefault float64

	_ = cmd.PersistentFlags().Float64(poePowerFlagName, poePowerFlagDefault, poePowerDescription)

	return nil
}

func registerInterfaceStatisticsRxbytes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	rxbytesDescription := ``

	var rxbytesFlagName string
	if cmdPrefix == "" {
		rxbytesFlagName = "rxbytes"
	} else {
		rxbytesFlagName = fmt.Sprintf("%v.rxbytes", cmdPrefix)
	}

	var rxbytesFlagDefault float64

	_ = cmd.PersistentFlags().Float64(rxbytesFlagName, rxbytesFlagDefault, rxbytesDescription)

	return nil
}

func registerInterfaceStatisticsRxrate(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	rxrateDescription := ``

	var rxrateFlagName string
	if cmdPrefix == "" {
		rxrateFlagName = "rxrate"
	} else {
		rxrateFlagName = fmt.Sprintf("%v.rxrate", cmdPrefix)
	}

	var rxrateFlagDefault float64

	_ = cmd.PersistentFlags().Float64(rxrateFlagName, rxrateFlagDefault, rxrateDescription)

	return nil
}

func registerInterfaceStatisticsTxbytes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	txbytesDescription := ``

	var txbytesFlagName string
	if cmdPrefix == "" {
		txbytesFlagName = "txbytes"
	} else {
		txbytesFlagName = fmt.Sprintf("%v.txbytes", cmdPrefix)
	}

	var txbytesFlagDefault float64

	_ = cmd.PersistentFlags().Float64(txbytesFlagName, txbytesFlagDefault, txbytesDescription)

	return nil
}

func registerInterfaceStatisticsTxrate(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	txrateDescription := ``

	var txrateFlagName string
	if cmdPrefix == "" {
		txrateFlagName = "txrate"
	} else {
		txrateFlagName = fmt.Sprintf("%v.txrate", cmdPrefix)
	}

	var txrateFlagDefault float64

	_ = cmd.PersistentFlags().Float64(txrateFlagName, txrateFlagDefault, txrateDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelInterfaceStatisticsFlags(depth int, m *models.InterfaceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, droppedAdded := retrieveInterfaceStatisticsDroppedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || droppedAdded

	err, errorsAdded := retrieveInterfaceStatisticsErrorsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || errorsAdded

	err, poePowerAdded := retrieveInterfaceStatisticsPoePowerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || poePowerAdded

	err, rxbytesAdded := retrieveInterfaceStatisticsRxbytesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || rxbytesAdded

	err, rxrateAdded := retrieveInterfaceStatisticsRxrateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || rxrateAdded

	err, txbytesAdded := retrieveInterfaceStatisticsTxbytesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || txbytesAdded

	err, txrateAdded := retrieveInterfaceStatisticsTxrateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || txrateAdded

	return nil, retAdded
}

func retrieveInterfaceStatisticsDroppedFlags(depth int, m *models.InterfaceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	droppedFlagName := fmt.Sprintf("%v.dropped", cmdPrefix)
	if cmd.Flags().Changed(droppedFlagName) {

		var droppedFlagName string
		if cmdPrefix == "" {
			droppedFlagName = "dropped"
		} else {
			droppedFlagName = fmt.Sprintf("%v.dropped", cmdPrefix)
		}

		droppedFlagValue, err := cmd.Flags().GetFloat64(droppedFlagName)
		if err != nil {
			return err, false
		}
		m.Dropped = droppedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveInterfaceStatisticsErrorsFlags(depth int, m *models.InterfaceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	errorsFlagName := fmt.Sprintf("%v.errors", cmdPrefix)
	if cmd.Flags().Changed(errorsFlagName) {

		var errorsFlagName string
		if cmdPrefix == "" {
			errorsFlagName = "errors"
		} else {
			errorsFlagName = fmt.Sprintf("%v.errors", cmdPrefix)
		}

		errorsFlagValue, err := cmd.Flags().GetFloat64(errorsFlagName)
		if err != nil {
			return err, false
		}
		m.Errors = errorsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveInterfaceStatisticsPoePowerFlags(depth int, m *models.InterfaceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	poePowerFlagName := fmt.Sprintf("%v.poePower", cmdPrefix)
	if cmd.Flags().Changed(poePowerFlagName) {

		var poePowerFlagName string
		if cmdPrefix == "" {
			poePowerFlagName = "poePower"
		} else {
			poePowerFlagName = fmt.Sprintf("%v.poePower", cmdPrefix)
		}

		poePowerFlagValue, err := cmd.Flags().GetFloat64(poePowerFlagName)
		if err != nil {
			return err, false
		}
		m.PoePower = poePowerFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveInterfaceStatisticsRxbytesFlags(depth int, m *models.InterfaceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	rxbytesFlagName := fmt.Sprintf("%v.rxbytes", cmdPrefix)
	if cmd.Flags().Changed(rxbytesFlagName) {

		var rxbytesFlagName string
		if cmdPrefix == "" {
			rxbytesFlagName = "rxbytes"
		} else {
			rxbytesFlagName = fmt.Sprintf("%v.rxbytes", cmdPrefix)
		}

		rxbytesFlagValue, err := cmd.Flags().GetFloat64(rxbytesFlagName)
		if err != nil {
			return err, false
		}
		m.Rxbytes = rxbytesFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveInterfaceStatisticsRxrateFlags(depth int, m *models.InterfaceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	rxrateFlagName := fmt.Sprintf("%v.rxrate", cmdPrefix)
	if cmd.Flags().Changed(rxrateFlagName) {

		var rxrateFlagName string
		if cmdPrefix == "" {
			rxrateFlagName = "rxrate"
		} else {
			rxrateFlagName = fmt.Sprintf("%v.rxrate", cmdPrefix)
		}

		rxrateFlagValue, err := cmd.Flags().GetFloat64(rxrateFlagName)
		if err != nil {
			return err, false
		}
		m.Rxrate = rxrateFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveInterfaceStatisticsTxbytesFlags(depth int, m *models.InterfaceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	txbytesFlagName := fmt.Sprintf("%v.txbytes", cmdPrefix)
	if cmd.Flags().Changed(txbytesFlagName) {

		var txbytesFlagName string
		if cmdPrefix == "" {
			txbytesFlagName = "txbytes"
		} else {
			txbytesFlagName = fmt.Sprintf("%v.txbytes", cmdPrefix)
		}

		txbytesFlagValue, err := cmd.Flags().GetFloat64(txbytesFlagName)
		if err != nil {
			return err, false
		}
		m.Txbytes = txbytesFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveInterfaceStatisticsTxrateFlags(depth int, m *models.InterfaceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	txrateFlagName := fmt.Sprintf("%v.txrate", cmdPrefix)
	if cmd.Flags().Changed(txrateFlagName) {

		var txrateFlagName string
		if cmdPrefix == "" {
			txrateFlagName = "txrate"
		} else {
			txrateFlagName = fmt.Sprintf("%v.txrate", cmdPrefix)
		}

		txrateFlagValue, err := cmd.Flags().GetFloat64(txrateFlagName)
		if err != nil {
			return err, false
		}
		m.Txrate = txrateFlagValue

		retAdded = true
	}

	return nil, retAdded
}