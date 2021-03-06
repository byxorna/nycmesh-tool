// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for DeviceOnu1

// register flags to command
func registerModelDeviceOnu1Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDeviceOnu1ID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOnu1Port(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOnu1PortForwards(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOnu1PppoeUser(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOnu1Profile(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOnu1ProfileName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOnu1ReceivePower(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOnu1RxRate(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOnu1TransmitPower(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOnu1TxRate(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceOnu1WanAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceOnu1ID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := ``

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

func registerDeviceOnu1Port(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	portDescription := ``

	var portFlagName string
	if cmdPrefix == "" {
		portFlagName = "port"
	} else {
		portFlagName = fmt.Sprintf("%v.port", cmdPrefix)
	}

	var portFlagDefault int64

	_ = cmd.PersistentFlags().Int64(portFlagName, portFlagDefault, portDescription)

	return nil
}

func registerDeviceOnu1PortForwards(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	portForwardsDescription := ``

	var portForwardsFlagName string
	if cmdPrefix == "" {
		portForwardsFlagName = "portForwards"
	} else {
		portForwardsFlagName = fmt.Sprintf("%v.portForwards", cmdPrefix)
	}

	var portForwardsFlagDefault string

	_ = cmd.PersistentFlags().String(portForwardsFlagName, portForwardsFlagDefault, portForwardsDescription)

	return nil
}

func registerDeviceOnu1PppoeUser(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	pppoeUserDescription := ``

	var pppoeUserFlagName string
	if cmdPrefix == "" {
		pppoeUserFlagName = "pppoeUser"
	} else {
		pppoeUserFlagName = fmt.Sprintf("%v.pppoeUser", cmdPrefix)
	}

	var pppoeUserFlagDefault string

	_ = cmd.PersistentFlags().String(pppoeUserFlagName, pppoeUserFlagDefault, pppoeUserDescription)

	return nil
}

func registerDeviceOnu1Profile(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	profileDescription := ``

	var profileFlagName string
	if cmdPrefix == "" {
		profileFlagName = "profile"
	} else {
		profileFlagName = fmt.Sprintf("%v.profile", cmdPrefix)
	}

	var profileFlagDefault string

	_ = cmd.PersistentFlags().String(profileFlagName, profileFlagDefault, profileDescription)

	return nil
}

func registerDeviceOnu1ProfileName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	profileNameDescription := ``

	var profileNameFlagName string
	if cmdPrefix == "" {
		profileNameFlagName = "profileName"
	} else {
		profileNameFlagName = fmt.Sprintf("%v.profileName", cmdPrefix)
	}

	var profileNameFlagDefault string

	_ = cmd.PersistentFlags().String(profileNameFlagName, profileNameFlagDefault, profileNameDescription)

	return nil
}

func registerDeviceOnu1ReceivePower(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	receivePowerDescription := ``

	var receivePowerFlagName string
	if cmdPrefix == "" {
		receivePowerFlagName = "receivePower"
	} else {
		receivePowerFlagName = fmt.Sprintf("%v.receivePower", cmdPrefix)
	}

	var receivePowerFlagDefault float64

	_ = cmd.PersistentFlags().Float64(receivePowerFlagName, receivePowerFlagDefault, receivePowerDescription)

	return nil
}

func registerDeviceOnu1RxRate(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	rxRateDescription := ``

	var rxRateFlagName string
	if cmdPrefix == "" {
		rxRateFlagName = "rxRate"
	} else {
		rxRateFlagName = fmt.Sprintf("%v.rxRate", cmdPrefix)
	}

	var rxRateFlagDefault float64

	_ = cmd.PersistentFlags().Float64(rxRateFlagName, rxRateFlagDefault, rxRateDescription)

	return nil
}

func registerDeviceOnu1TransmitPower(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	transmitPowerDescription := ``

	var transmitPowerFlagName string
	if cmdPrefix == "" {
		transmitPowerFlagName = "transmitPower"
	} else {
		transmitPowerFlagName = fmt.Sprintf("%v.transmitPower", cmdPrefix)
	}

	var transmitPowerFlagDefault float64

	_ = cmd.PersistentFlags().Float64(transmitPowerFlagName, transmitPowerFlagDefault, transmitPowerDescription)

	return nil
}

func registerDeviceOnu1TxRate(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	txRateDescription := ``

	var txRateFlagName string
	if cmdPrefix == "" {
		txRateFlagName = "txRate"
	} else {
		txRateFlagName = fmt.Sprintf("%v.txRate", cmdPrefix)
	}

	var txRateFlagDefault float64

	_ = cmd.PersistentFlags().Float64(txRateFlagName, txRateFlagDefault, txRateDescription)

	return nil
}

func registerDeviceOnu1WanAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	wanAddressDescription := ``

	var wanAddressFlagName string
	if cmdPrefix == "" {
		wanAddressFlagName = "wanAddress"
	} else {
		wanAddressFlagName = fmt.Sprintf("%v.wanAddress", cmdPrefix)
	}

	var wanAddressFlagDefault string

	_ = cmd.PersistentFlags().String(wanAddressFlagName, wanAddressFlagDefault, wanAddressDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDeviceOnu1Flags(depth int, m *models.DeviceOnu1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, idAdded := retrieveDeviceOnu1IDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, portAdded := retrieveDeviceOnu1PortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || portAdded

	err, portForwardsAdded := retrieveDeviceOnu1PortForwardsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || portForwardsAdded

	err, pppoeUserAdded := retrieveDeviceOnu1PppoeUserFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pppoeUserAdded

	err, profileAdded := retrieveDeviceOnu1ProfileFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || profileAdded

	err, profileNameAdded := retrieveDeviceOnu1ProfileNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || profileNameAdded

	err, receivePowerAdded := retrieveDeviceOnu1ReceivePowerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || receivePowerAdded

	err, rxRateAdded := retrieveDeviceOnu1RxRateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || rxRateAdded

	err, transmitPowerAdded := retrieveDeviceOnu1TransmitPowerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || transmitPowerAdded

	err, txRateAdded := retrieveDeviceOnu1TxRateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || txRateAdded

	err, wanAddressAdded := retrieveDeviceOnu1WanAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || wanAddressAdded

	return nil, retAdded
}

func retrieveDeviceOnu1IDFlags(depth int, m *models.DeviceOnu1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDeviceOnu1PortFlags(depth int, m *models.DeviceOnu1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	portFlagName := fmt.Sprintf("%v.port", cmdPrefix)
	if cmd.Flags().Changed(portFlagName) {

		var portFlagName string
		if cmdPrefix == "" {
			portFlagName = "port"
		} else {
			portFlagName = fmt.Sprintf("%v.port", cmdPrefix)
		}

		portFlagValue, err := cmd.Flags().GetInt64(portFlagName)
		if err != nil {
			return err, false
		}
		m.Port = portFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceOnu1PortForwardsFlags(depth int, m *models.DeviceOnu1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	portForwardsFlagName := fmt.Sprintf("%v.portForwards", cmdPrefix)
	if cmd.Flags().Changed(portForwardsFlagName) {

		var portForwardsFlagName string
		if cmdPrefix == "" {
			portForwardsFlagName = "portForwards"
		} else {
			portForwardsFlagName = fmt.Sprintf("%v.portForwards", cmdPrefix)
		}

		portForwardsFlagValue, err := cmd.Flags().GetString(portForwardsFlagName)
		if err != nil {
			return err, false
		}
		m.PortForwards = portForwardsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceOnu1PppoeUserFlags(depth int, m *models.DeviceOnu1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pppoeUserFlagName := fmt.Sprintf("%v.pppoeUser", cmdPrefix)
	if cmd.Flags().Changed(pppoeUserFlagName) {

		var pppoeUserFlagName string
		if cmdPrefix == "" {
			pppoeUserFlagName = "pppoeUser"
		} else {
			pppoeUserFlagName = fmt.Sprintf("%v.pppoeUser", cmdPrefix)
		}

		pppoeUserFlagValue, err := cmd.Flags().GetString(pppoeUserFlagName)
		if err != nil {
			return err, false
		}
		m.PppoeUser = pppoeUserFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceOnu1ProfileFlags(depth int, m *models.DeviceOnu1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	profileFlagName := fmt.Sprintf("%v.profile", cmdPrefix)
	if cmd.Flags().Changed(profileFlagName) {

		var profileFlagName string
		if cmdPrefix == "" {
			profileFlagName = "profile"
		} else {
			profileFlagName = fmt.Sprintf("%v.profile", cmdPrefix)
		}

		profileFlagValue, err := cmd.Flags().GetString(profileFlagName)
		if err != nil {
			return err, false
		}
		m.Profile = profileFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceOnu1ProfileNameFlags(depth int, m *models.DeviceOnu1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	profileNameFlagName := fmt.Sprintf("%v.profileName", cmdPrefix)
	if cmd.Flags().Changed(profileNameFlagName) {

		var profileNameFlagName string
		if cmdPrefix == "" {
			profileNameFlagName = "profileName"
		} else {
			profileNameFlagName = fmt.Sprintf("%v.profileName", cmdPrefix)
		}

		profileNameFlagValue, err := cmd.Flags().GetString(profileNameFlagName)
		if err != nil {
			return err, false
		}
		m.ProfileName = profileNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceOnu1ReceivePowerFlags(depth int, m *models.DeviceOnu1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	receivePowerFlagName := fmt.Sprintf("%v.receivePower", cmdPrefix)
	if cmd.Flags().Changed(receivePowerFlagName) {

		var receivePowerFlagName string
		if cmdPrefix == "" {
			receivePowerFlagName = "receivePower"
		} else {
			receivePowerFlagName = fmt.Sprintf("%v.receivePower", cmdPrefix)
		}

		receivePowerFlagValue, err := cmd.Flags().GetFloat64(receivePowerFlagName)
		if err != nil {
			return err, false
		}
		m.ReceivePower = receivePowerFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceOnu1RxRateFlags(depth int, m *models.DeviceOnu1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	rxRateFlagName := fmt.Sprintf("%v.rxRate", cmdPrefix)
	if cmd.Flags().Changed(rxRateFlagName) {

		var rxRateFlagName string
		if cmdPrefix == "" {
			rxRateFlagName = "rxRate"
		} else {
			rxRateFlagName = fmt.Sprintf("%v.rxRate", cmdPrefix)
		}

		rxRateFlagValue, err := cmd.Flags().GetFloat64(rxRateFlagName)
		if err != nil {
			return err, false
		}
		m.RxRate = rxRateFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceOnu1TransmitPowerFlags(depth int, m *models.DeviceOnu1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	transmitPowerFlagName := fmt.Sprintf("%v.transmitPower", cmdPrefix)
	if cmd.Flags().Changed(transmitPowerFlagName) {

		var transmitPowerFlagName string
		if cmdPrefix == "" {
			transmitPowerFlagName = "transmitPower"
		} else {
			transmitPowerFlagName = fmt.Sprintf("%v.transmitPower", cmdPrefix)
		}

		transmitPowerFlagValue, err := cmd.Flags().GetFloat64(transmitPowerFlagName)
		if err != nil {
			return err, false
		}
		m.TransmitPower = transmitPowerFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceOnu1TxRateFlags(depth int, m *models.DeviceOnu1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	txRateFlagName := fmt.Sprintf("%v.txRate", cmdPrefix)
	if cmd.Flags().Changed(txRateFlagName) {

		var txRateFlagName string
		if cmdPrefix == "" {
			txRateFlagName = "txRate"
		} else {
			txRateFlagName = fmt.Sprintf("%v.txRate", cmdPrefix)
		}

		txRateFlagValue, err := cmd.Flags().GetFloat64(txRateFlagName)
		if err != nil {
			return err, false
		}
		m.TxRate = txRateFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceOnu1WanAddressFlags(depth int, m *models.DeviceOnu1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	wanAddressFlagName := fmt.Sprintf("%v.wanAddress", cmdPrefix)
	if cmd.Flags().Changed(wanAddressFlagName) {

		var wanAddressFlagName string
		if cmdPrefix == "" {
			wanAddressFlagName = "wanAddress"
		} else {
			wanAddressFlagName = fmt.Sprintf("%v.wanAddress", cmdPrefix)
		}

		wanAddressFlagValue, err := cmd.Flags().GetString(wanAddressFlagName)
		if err != nil {
			return err, false
		}
		m.WanAddress = wanAddressFlagValue

		retAdded = true
	}

	return nil, retAdded
}
