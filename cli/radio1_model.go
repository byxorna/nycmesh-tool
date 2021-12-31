// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for Radio1

// register flags to command
func registerModelRadio1Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerRadio1AntennaList(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRadio1CcodeList(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRadio1CcodeLocked(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRadio1ChannelWidthList(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRadio1DefaultAntennaID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRadio1DistanceLimit(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRadio1IeeeBitmask(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRadio1SupportedBands(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRadio1SupportedIeeeStandards(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRadio1SupportsApPtmp(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRadio1SupportsApPtp(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRadio1SupportsStaPtmp(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRadio1SupportsStaPtp(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRadio1TxPowerRange(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerRadio1AntennaList(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: antennaList AntennaList array type is not supported by go-swagger cli yet

	return nil
}

func registerRadio1CcodeList(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ccodeListDescription := `Required. `

	var ccodeListFlagName string
	if cmdPrefix == "" {
		ccodeListFlagName = "ccodeList"
	} else {
		ccodeListFlagName = fmt.Sprintf("%v.ccodeList", cmdPrefix)
	}

	var ccodeListFlagDefault float64

	_ = cmd.PersistentFlags().Float64(ccodeListFlagName, ccodeListFlagDefault, ccodeListDescription)

	return nil
}

func registerRadio1CcodeLocked(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ccodeLockedDescription := `Required. `

	var ccodeLockedFlagName string
	if cmdPrefix == "" {
		ccodeLockedFlagName = "ccodeLocked"
	} else {
		ccodeLockedFlagName = fmt.Sprintf("%v.ccodeLocked", cmdPrefix)
	}

	var ccodeLockedFlagDefault float64

	_ = cmd.PersistentFlags().Float64(ccodeLockedFlagName, ccodeLockedFlagDefault, ccodeLockedDescription)

	return nil
}

func registerRadio1ChannelWidthList(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: channelWidthList ChannelWidthList array type is not supported by go-swagger cli yet

	return nil
}

func registerRadio1DefaultAntennaID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	defaultAntennaIdDescription := `Required. `

	var defaultAntennaIdFlagName string
	if cmdPrefix == "" {
		defaultAntennaIdFlagName = "defaultAntennaId"
	} else {
		defaultAntennaIdFlagName = fmt.Sprintf("%v.defaultAntennaId", cmdPrefix)
	}

	var defaultAntennaIdFlagDefault float64

	_ = cmd.PersistentFlags().Float64(defaultAntennaIdFlagName, defaultAntennaIdFlagDefault, defaultAntennaIdDescription)

	return nil
}

func registerRadio1DistanceLimit(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	distanceLimitDescription := `Required. `

	var distanceLimitFlagName string
	if cmdPrefix == "" {
		distanceLimitFlagName = "distanceLimit"
	} else {
		distanceLimitFlagName = fmt.Sprintf("%v.distanceLimit", cmdPrefix)
	}

	var distanceLimitFlagDefault float64

	_ = cmd.PersistentFlags().Float64(distanceLimitFlagName, distanceLimitFlagDefault, distanceLimitDescription)

	return nil
}

func registerRadio1IeeeBitmask(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ieeeBitmaskDescription := `Required. `

	var ieeeBitmaskFlagName string
	if cmdPrefix == "" {
		ieeeBitmaskFlagName = "ieeeBitmask"
	} else {
		ieeeBitmaskFlagName = fmt.Sprintf("%v.ieeeBitmask", cmdPrefix)
	}

	var ieeeBitmaskFlagDefault float64

	_ = cmd.PersistentFlags().Float64(ieeeBitmaskFlagName, ieeeBitmaskFlagDefault, ieeeBitmaskDescription)

	return nil
}

func registerRadio1SupportedBands(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: supportedBands SupportedBands array type is not supported by go-swagger cli yet

	return nil
}

func registerRadio1SupportedIeeeStandards(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: supportedIeeeStandards SupportedIeeeStandards array type is not supported by go-swagger cli yet

	return nil
}

func registerRadio1SupportsApPtmp(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	supportsApPtmpDescription := `Required. `

	var supportsApPtmpFlagName string
	if cmdPrefix == "" {
		supportsApPtmpFlagName = "supportsApPtmp"
	} else {
		supportsApPtmpFlagName = fmt.Sprintf("%v.supportsApPtmp", cmdPrefix)
	}

	var supportsApPtmpFlagDefault bool

	_ = cmd.PersistentFlags().Bool(supportsApPtmpFlagName, supportsApPtmpFlagDefault, supportsApPtmpDescription)

	return nil
}

func registerRadio1SupportsApPtp(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	supportsApPtpDescription := `Required. `

	var supportsApPtpFlagName string
	if cmdPrefix == "" {
		supportsApPtpFlagName = "supportsApPtp"
	} else {
		supportsApPtpFlagName = fmt.Sprintf("%v.supportsApPtp", cmdPrefix)
	}

	var supportsApPtpFlagDefault bool

	_ = cmd.PersistentFlags().Bool(supportsApPtpFlagName, supportsApPtpFlagDefault, supportsApPtpDescription)

	return nil
}

func registerRadio1SupportsStaPtmp(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	supportsStaPtmpDescription := `Required. `

	var supportsStaPtmpFlagName string
	if cmdPrefix == "" {
		supportsStaPtmpFlagName = "supportsStaPtmp"
	} else {
		supportsStaPtmpFlagName = fmt.Sprintf("%v.supportsStaPtmp", cmdPrefix)
	}

	var supportsStaPtmpFlagDefault bool

	_ = cmd.PersistentFlags().Bool(supportsStaPtmpFlagName, supportsStaPtmpFlagDefault, supportsStaPtmpDescription)

	return nil
}

func registerRadio1SupportsStaPtp(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	supportsStaPtpDescription := `Required. `

	var supportsStaPtpFlagName string
	if cmdPrefix == "" {
		supportsStaPtpFlagName = "supportsStaPtp"
	} else {
		supportsStaPtpFlagName = fmt.Sprintf("%v.supportsStaPtp", cmdPrefix)
	}

	var supportsStaPtpFlagDefault bool

	_ = cmd.PersistentFlags().Bool(supportsStaPtpFlagName, supportsStaPtpFlagDefault, supportsStaPtpDescription)

	return nil
}

func registerRadio1TxPowerRange(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var txPowerRangeFlagName string
	if cmdPrefix == "" {
		txPowerRangeFlagName = "txPowerRange"
	} else {
		txPowerRangeFlagName = fmt.Sprintf("%v.txPowerRange", cmdPrefix)
	}

	if err := registerModelTxPowerRangeFlags(depth+1, txPowerRangeFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelRadio1Flags(depth int, m *models.Radio1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, antennaListAdded := retrieveRadio1AntennaListFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || antennaListAdded

	err, ccodeListAdded := retrieveRadio1CcodeListFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ccodeListAdded

	err, ccodeLockedAdded := retrieveRadio1CcodeLockedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ccodeLockedAdded

	err, channelWidthListAdded := retrieveRadio1ChannelWidthListFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || channelWidthListAdded

	err, defaultAntennaIdAdded := retrieveRadio1DefaultAntennaIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || defaultAntennaIdAdded

	err, distanceLimitAdded := retrieveRadio1DistanceLimitFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || distanceLimitAdded

	err, ieeeBitmaskAdded := retrieveRadio1IeeeBitmaskFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ieeeBitmaskAdded

	err, supportedBandsAdded := retrieveRadio1SupportedBandsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || supportedBandsAdded

	err, supportedIeeeStandardsAdded := retrieveRadio1SupportedIeeeStandardsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || supportedIeeeStandardsAdded

	err, supportsApPtmpAdded := retrieveRadio1SupportsApPtmpFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || supportsApPtmpAdded

	err, supportsApPtpAdded := retrieveRadio1SupportsApPtpFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || supportsApPtpAdded

	err, supportsStaPtmpAdded := retrieveRadio1SupportsStaPtmpFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || supportsStaPtmpAdded

	err, supportsStaPtpAdded := retrieveRadio1SupportsStaPtpFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || supportsStaPtpAdded

	err, txPowerRangeAdded := retrieveRadio1TxPowerRangeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || txPowerRangeAdded

	return nil, retAdded
}

func retrieveRadio1AntennaListFlags(depth int, m *models.Radio1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	antennaListFlagName := fmt.Sprintf("%v.antennaList", cmdPrefix)
	if cmd.Flags().Changed(antennaListFlagName) {
		// warning: antennaList array type AntennaList is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveRadio1CcodeListFlags(depth int, m *models.Radio1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ccodeListFlagName := fmt.Sprintf("%v.ccodeList", cmdPrefix)
	if cmd.Flags().Changed(ccodeListFlagName) {

		var ccodeListFlagName string
		if cmdPrefix == "" {
			ccodeListFlagName = "ccodeList"
		} else {
			ccodeListFlagName = fmt.Sprintf("%v.ccodeList", cmdPrefix)
		}

		ccodeListFlagValue, err := cmd.Flags().GetFloat64(ccodeListFlagName)
		if err != nil {
			return err, false
		}
		m.CcodeList = &ccodeListFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRadio1CcodeLockedFlags(depth int, m *models.Radio1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ccodeLockedFlagName := fmt.Sprintf("%v.ccodeLocked", cmdPrefix)
	if cmd.Flags().Changed(ccodeLockedFlagName) {

		var ccodeLockedFlagName string
		if cmdPrefix == "" {
			ccodeLockedFlagName = "ccodeLocked"
		} else {
			ccodeLockedFlagName = fmt.Sprintf("%v.ccodeLocked", cmdPrefix)
		}

		ccodeLockedFlagValue, err := cmd.Flags().GetFloat64(ccodeLockedFlagName)
		if err != nil {
			return err, false
		}
		m.CcodeLocked = &ccodeLockedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRadio1ChannelWidthListFlags(depth int, m *models.Radio1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	channelWidthListFlagName := fmt.Sprintf("%v.channelWidthList", cmdPrefix)
	if cmd.Flags().Changed(channelWidthListFlagName) {
		// warning: channelWidthList array type ChannelWidthList is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveRadio1DefaultAntennaIDFlags(depth int, m *models.Radio1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	defaultAntennaIdFlagName := fmt.Sprintf("%v.defaultAntennaId", cmdPrefix)
	if cmd.Flags().Changed(defaultAntennaIdFlagName) {

		var defaultAntennaIdFlagName string
		if cmdPrefix == "" {
			defaultAntennaIdFlagName = "defaultAntennaId"
		} else {
			defaultAntennaIdFlagName = fmt.Sprintf("%v.defaultAntennaId", cmdPrefix)
		}

		defaultAntennaIdFlagValue, err := cmd.Flags().GetFloat64(defaultAntennaIdFlagName)
		if err != nil {
			return err, false
		}
		m.DefaultAntennaID = &defaultAntennaIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRadio1DistanceLimitFlags(depth int, m *models.Radio1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	distanceLimitFlagName := fmt.Sprintf("%v.distanceLimit", cmdPrefix)
	if cmd.Flags().Changed(distanceLimitFlagName) {

		var distanceLimitFlagName string
		if cmdPrefix == "" {
			distanceLimitFlagName = "distanceLimit"
		} else {
			distanceLimitFlagName = fmt.Sprintf("%v.distanceLimit", cmdPrefix)
		}

		distanceLimitFlagValue, err := cmd.Flags().GetFloat64(distanceLimitFlagName)
		if err != nil {
			return err, false
		}
		m.DistanceLimit = &distanceLimitFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRadio1IeeeBitmaskFlags(depth int, m *models.Radio1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ieeeBitmaskFlagName := fmt.Sprintf("%v.ieeeBitmask", cmdPrefix)
	if cmd.Flags().Changed(ieeeBitmaskFlagName) {

		var ieeeBitmaskFlagName string
		if cmdPrefix == "" {
			ieeeBitmaskFlagName = "ieeeBitmask"
		} else {
			ieeeBitmaskFlagName = fmt.Sprintf("%v.ieeeBitmask", cmdPrefix)
		}

		ieeeBitmaskFlagValue, err := cmd.Flags().GetFloat64(ieeeBitmaskFlagName)
		if err != nil {
			return err, false
		}
		m.IeeeBitmask = &ieeeBitmaskFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRadio1SupportedBandsFlags(depth int, m *models.Radio1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	supportedBandsFlagName := fmt.Sprintf("%v.supportedBands", cmdPrefix)
	if cmd.Flags().Changed(supportedBandsFlagName) {
		// warning: supportedBands array type SupportedBands is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveRadio1SupportedIeeeStandardsFlags(depth int, m *models.Radio1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	supportedIeeeStandardsFlagName := fmt.Sprintf("%v.supportedIeeeStandards", cmdPrefix)
	if cmd.Flags().Changed(supportedIeeeStandardsFlagName) {
		// warning: supportedIeeeStandards array type SupportedIeeeStandards is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveRadio1SupportsApPtmpFlags(depth int, m *models.Radio1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	supportsApPtmpFlagName := fmt.Sprintf("%v.supportsApPtmp", cmdPrefix)
	if cmd.Flags().Changed(supportsApPtmpFlagName) {

		var supportsApPtmpFlagName string
		if cmdPrefix == "" {
			supportsApPtmpFlagName = "supportsApPtmp"
		} else {
			supportsApPtmpFlagName = fmt.Sprintf("%v.supportsApPtmp", cmdPrefix)
		}

		supportsApPtmpFlagValue, err := cmd.Flags().GetBool(supportsApPtmpFlagName)
		if err != nil {
			return err, false
		}
		m.SupportsApPtmp = &supportsApPtmpFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRadio1SupportsApPtpFlags(depth int, m *models.Radio1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	supportsApPtpFlagName := fmt.Sprintf("%v.supportsApPtp", cmdPrefix)
	if cmd.Flags().Changed(supportsApPtpFlagName) {

		var supportsApPtpFlagName string
		if cmdPrefix == "" {
			supportsApPtpFlagName = "supportsApPtp"
		} else {
			supportsApPtpFlagName = fmt.Sprintf("%v.supportsApPtp", cmdPrefix)
		}

		supportsApPtpFlagValue, err := cmd.Flags().GetBool(supportsApPtpFlagName)
		if err != nil {
			return err, false
		}
		m.SupportsApPtp = &supportsApPtpFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRadio1SupportsStaPtmpFlags(depth int, m *models.Radio1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	supportsStaPtmpFlagName := fmt.Sprintf("%v.supportsStaPtmp", cmdPrefix)
	if cmd.Flags().Changed(supportsStaPtmpFlagName) {

		var supportsStaPtmpFlagName string
		if cmdPrefix == "" {
			supportsStaPtmpFlagName = "supportsStaPtmp"
		} else {
			supportsStaPtmpFlagName = fmt.Sprintf("%v.supportsStaPtmp", cmdPrefix)
		}

		supportsStaPtmpFlagValue, err := cmd.Flags().GetBool(supportsStaPtmpFlagName)
		if err != nil {
			return err, false
		}
		m.SupportsStaPtmp = &supportsStaPtmpFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRadio1SupportsStaPtpFlags(depth int, m *models.Radio1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	supportsStaPtpFlagName := fmt.Sprintf("%v.supportsStaPtp", cmdPrefix)
	if cmd.Flags().Changed(supportsStaPtpFlagName) {

		var supportsStaPtpFlagName string
		if cmdPrefix == "" {
			supportsStaPtpFlagName = "supportsStaPtp"
		} else {
			supportsStaPtpFlagName = fmt.Sprintf("%v.supportsStaPtp", cmdPrefix)
		}

		supportsStaPtpFlagValue, err := cmd.Flags().GetBool(supportsStaPtpFlagName)
		if err != nil {
			return err, false
		}
		m.SupportsStaPtp = &supportsStaPtpFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRadio1TxPowerRangeFlags(depth int, m *models.Radio1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	txPowerRangeFlagName := fmt.Sprintf("%v.txPowerRange", cmdPrefix)
	if cmd.Flags().Changed(txPowerRangeFlagName) {
		// info: complex object txPowerRange TxPowerRange is retrieved outside this Changed() block
	}
	txPowerRangeFlagValue := m.TxPowerRange
	if swag.IsZero(txPowerRangeFlagValue) {
		txPowerRangeFlagValue = &models.TxPowerRange{}
	}

	err, txPowerRangeAdded := retrieveModelTxPowerRangeFlags(depth+1, txPowerRangeFlagValue, txPowerRangeFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || txPowerRangeAdded
	if txPowerRangeAdded {
		m.TxPowerRange = txPowerRangeFlagValue
	}

	return nil, retAdded
}
