// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for Latest

// register flags to command
func registerModelLatestFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerLatestMajor(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerLatestMinor(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerLatestOrder(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerLatestPatch(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerLatestPrerelease(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerLatestMajor(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerLatestMinor(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerLatestOrder(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	orderDescription := `Required. `

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

func registerLatestPatch(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerLatestPrerelease(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: prerelease Prerelease array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelLatestFlags(depth int, m *models.Latest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, majorAdded := retrieveLatestMajorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || majorAdded

	err, minorAdded := retrieveLatestMinorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || minorAdded

	err, orderAdded := retrieveLatestOrderFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || orderAdded

	err, patchAdded := retrieveLatestPatchFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || patchAdded

	err, prereleaseAdded := retrieveLatestPrereleaseFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || prereleaseAdded

	return nil, retAdded
}

func retrieveLatestMajorFlags(depth int, m *models.Latest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveLatestMinorFlags(depth int, m *models.Latest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveLatestOrderFlags(depth int, m *models.Latest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Order = &orderFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveLatestPatchFlags(depth int, m *models.Latest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveLatestPrereleaseFlags(depth int, m *models.Latest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	prereleaseFlagName := fmt.Sprintf("%v.prerelease", cmdPrefix)
	if cmd.Flags().Changed(prereleaseFlagName) {
		// warning: prerelease array type Prerelease is not supported by go-swagger cli yet
	}

	return nil, retAdded
}