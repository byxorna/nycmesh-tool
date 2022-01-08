// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for Firmware

// register flags to command
func registerModelFirmwareFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerFirmwareMajor(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFirmwareMinor(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFirmwareOrder(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFirmwarePatch(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFirmwarePrerelease(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFirmwareUpgradeRecommendedToVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerFirmwareMajor(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	majorDescription := `Required. `

	var majorFlagName string
	if cmdPrefix == "" {
		majorFlagName = "major"
	} else {
		majorFlagName = fmt.Sprintf("%v.major", cmdPrefix)
	}

	var majorFlagDefault float64

	_ = cmd.PersistentFlags().Float64(majorFlagName, majorFlagDefault, majorDescription)

	return nil
}

func registerFirmwareMinor(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	minorDescription := `Required. `

	var minorFlagName string
	if cmdPrefix == "" {
		minorFlagName = "minor"
	} else {
		minorFlagName = fmt.Sprintf("%v.minor", cmdPrefix)
	}

	var minorFlagDefault float64

	_ = cmd.PersistentFlags().Float64(minorFlagName, minorFlagDefault, minorDescription)

	return nil
}

func registerFirmwareOrder(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	orderDescription := ``

	var orderFlagName string
	if cmdPrefix == "" {
		orderFlagName = "order"
	} else {
		orderFlagName = fmt.Sprintf("%v.order", cmdPrefix)
	}

	var orderFlagDefault string

	_ = cmd.PersistentFlags().String(orderFlagName, orderFlagDefault, orderDescription)

	return nil
}

func registerFirmwarePatch(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	patchDescription := `Required. `

	var patchFlagName string
	if cmdPrefix == "" {
		patchFlagName = "patch"
	} else {
		patchFlagName = fmt.Sprintf("%v.patch", cmdPrefix)
	}

	var patchFlagDefault float64

	_ = cmd.PersistentFlags().Float64(patchFlagName, patchFlagDefault, patchDescription)

	return nil
}

func registerFirmwarePrerelease(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: prerelease Prerelease1 array type is not supported by go-swagger cli yet

	return nil
}

func registerFirmwareUpgradeRecommendedToVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	upgradeRecommendedToVersionDescription := ``

	var upgradeRecommendedToVersionFlagName string
	if cmdPrefix == "" {
		upgradeRecommendedToVersionFlagName = "upgradeRecommendedToVersion"
	} else {
		upgradeRecommendedToVersionFlagName = fmt.Sprintf("%v.upgradeRecommendedToVersion", cmdPrefix)
	}

	var upgradeRecommendedToVersionFlagDefault string

	_ = cmd.PersistentFlags().String(upgradeRecommendedToVersionFlagName, upgradeRecommendedToVersionFlagDefault, upgradeRecommendedToVersionDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelFirmwareFlags(depth int, m *models.Firmware, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, majorAdded := retrieveFirmwareMajorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || majorAdded

	err, minorAdded := retrieveFirmwareMinorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || minorAdded

	err, orderAdded := retrieveFirmwareOrderFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || orderAdded

	err, patchAdded := retrieveFirmwarePatchFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || patchAdded

	err, prereleaseAdded := retrieveFirmwarePrereleaseFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || prereleaseAdded

	err, upgradeRecommendedToVersionAdded := retrieveFirmwareUpgradeRecommendedToVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || upgradeRecommendedToVersionAdded

	return nil, retAdded
}

func retrieveFirmwareMajorFlags(depth int, m *models.Firmware, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	majorFlagName := fmt.Sprintf("%v.major", cmdPrefix)
	if cmd.Flags().Changed(majorFlagName) {

		var majorFlagName string
		if cmdPrefix == "" {
			majorFlagName = "major"
		} else {
			majorFlagName = fmt.Sprintf("%v.major", cmdPrefix)
		}

		majorFlagValue, err := cmd.Flags().GetFloat64(majorFlagName)
		if err != nil {
			return err, false
		}
		m.Major = &majorFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveFirmwareMinorFlags(depth int, m *models.Firmware, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	minorFlagName := fmt.Sprintf("%v.minor", cmdPrefix)
	if cmd.Flags().Changed(minorFlagName) {

		var minorFlagName string
		if cmdPrefix == "" {
			minorFlagName = "minor"
		} else {
			minorFlagName = fmt.Sprintf("%v.minor", cmdPrefix)
		}

		minorFlagValue, err := cmd.Flags().GetFloat64(minorFlagName)
		if err != nil {
			return err, false
		}
		m.Minor = &minorFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveFirmwareOrderFlags(depth int, m *models.Firmware, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	orderFlagName := fmt.Sprintf("%v.order", cmdPrefix)
	if cmd.Flags().Changed(orderFlagName) {

		var orderFlagName string
		if cmdPrefix == "" {
			orderFlagName = "order"
		} else {
			orderFlagName = fmt.Sprintf("%v.order", cmdPrefix)
		}

		orderFlagValue, err := cmd.Flags().GetString(orderFlagName)
		if err != nil {
			return err, false
		}
		m.Order = orderFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveFirmwarePatchFlags(depth int, m *models.Firmware, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	patchFlagName := fmt.Sprintf("%v.patch", cmdPrefix)
	if cmd.Flags().Changed(patchFlagName) {

		var patchFlagName string
		if cmdPrefix == "" {
			patchFlagName = "patch"
		} else {
			patchFlagName = fmt.Sprintf("%v.patch", cmdPrefix)
		}

		patchFlagValue, err := cmd.Flags().GetFloat64(patchFlagName)
		if err != nil {
			return err, false
		}
		m.Patch = &patchFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveFirmwarePrereleaseFlags(depth int, m *models.Firmware, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	prereleaseFlagName := fmt.Sprintf("%v.prerelease", cmdPrefix)
	if cmd.Flags().Changed(prereleaseFlagName) {
		// warning: prerelease array type Prerelease1 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveFirmwareUpgradeRecommendedToVersionFlags(depth int, m *models.Firmware, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	upgradeRecommendedToVersionFlagName := fmt.Sprintf("%v.upgradeRecommendedToVersion", cmdPrefix)
	if cmd.Flags().Changed(upgradeRecommendedToVersionFlagName) {

		var upgradeRecommendedToVersionFlagName string
		if cmdPrefix == "" {
			upgradeRecommendedToVersionFlagName = "upgradeRecommendedToVersion"
		} else {
			upgradeRecommendedToVersionFlagName = fmt.Sprintf("%v.upgradeRecommendedToVersion", cmdPrefix)
		}

		upgradeRecommendedToVersionFlagValue, err := cmd.Flags().GetString(upgradeRecommendedToVersionFlagName)
		if err != nil {
			return err, false
		}
		m.UpgradeRecommendedToVersion = upgradeRecommendedToVersionFlagValue

		retAdded = true
	}

	return nil, retAdded
}
