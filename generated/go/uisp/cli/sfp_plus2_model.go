// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for SfpPlus2

// register flags to command
func registerModelSfpPlus2Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSfpPlus2VlanNative(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSfpPlus2VlanNative(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	vlanNativeDescription := ``

	var vlanNativeFlagName string
	if cmdPrefix == "" {
		vlanNativeFlagName = "vlanNative"
	} else {
		vlanNativeFlagName = fmt.Sprintf("%v.vlanNative", cmdPrefix)
	}

	var vlanNativeFlagDefault float64

	_ = cmd.PersistentFlags().Float64(vlanNativeFlagName, vlanNativeFlagDefault, vlanNativeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSfpPlus2Flags(depth int, m *models.SfpPlus2, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, vlanNativeAdded := retrieveSfpPlus2VlanNativeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || vlanNativeAdded

	return nil, retAdded
}

func retrieveSfpPlus2VlanNativeFlags(depth int, m *models.SfpPlus2, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	vlanNativeFlagName := fmt.Sprintf("%v.vlanNative", cmdPrefix)
	if cmd.Flags().Changed(vlanNativeFlagName) {

		var vlanNativeFlagName string
		if cmdPrefix == "" {
			vlanNativeFlagName = "vlanNative"
		} else {
			vlanNativeFlagName = fmt.Sprintf("%v.vlanNative", cmdPrefix)
		}

		vlanNativeFlagValue, err := cmd.Flags().GetFloat64(vlanNativeFlagName)
		if err != nil {
			return err, false
		}
		m.VlanNative = vlanNativeFlagValue

		retAdded = true
	}

	return nil, retAdded
}