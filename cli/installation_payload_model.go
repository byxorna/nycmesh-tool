// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for InstallationPayload

// register flags to command
func registerModelInstallationPayloadFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerInstallationPayloadMac1(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInstallationPayloadMac2(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerInstallationPayloadMac1(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	mac1Description := `Required. MAC address of the first installation device.`

	var mac1FlagName string
	if cmdPrefix == "" {
		mac1FlagName = "mac1"
	} else {
		mac1FlagName = fmt.Sprintf("%v.mac1", cmdPrefix)
	}

	var mac1FlagDefault string

	_ = cmd.PersistentFlags().String(mac1FlagName, mac1FlagDefault, mac1Description)

	return nil
}

func registerInstallationPayloadMac2(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	mac2Description := `Required. MAC address of the second installation device.`

	var mac2FlagName string
	if cmdPrefix == "" {
		mac2FlagName = "mac2"
	} else {
		mac2FlagName = fmt.Sprintf("%v.mac2", cmdPrefix)
	}

	var mac2FlagDefault string

	_ = cmd.PersistentFlags().String(mac2FlagName, mac2FlagDefault, mac2Description)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelInstallationPayloadFlags(depth int, m *models.InstallationPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, mac1Added := retrieveInstallationPayloadMac1Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || mac1Added

	err, mac2Added := retrieveInstallationPayloadMac2Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || mac2Added

	return nil, retAdded
}

func retrieveInstallationPayloadMac1Flags(depth int, m *models.InstallationPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	mac1FlagName := fmt.Sprintf("%v.mac1", cmdPrefix)
	if cmd.Flags().Changed(mac1FlagName) {

		var mac1FlagName string
		if cmdPrefix == "" {
			mac1FlagName = "mac1"
		} else {
			mac1FlagName = fmt.Sprintf("%v.mac1", cmdPrefix)
		}

		mac1FlagValue, err := cmd.Flags().GetString(mac1FlagName)
		if err != nil {
			return err, false
		}
		m.Mac1 = &mac1FlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveInstallationPayloadMac2Flags(depth int, m *models.InstallationPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	mac2FlagName := fmt.Sprintf("%v.mac2", cmdPrefix)
	if cmd.Flags().Changed(mac2FlagName) {

		var mac2FlagName string
		if cmdPrefix == "" {
			mac2FlagName = "mac2"
		} else {
			mac2FlagName = fmt.Sprintf("%v.mac2", cmdPrefix)
		}

		mac2FlagValue, err := cmd.Flags().GetString(mac2FlagName)
		if err != nil {
			return err, false
		}
		m.Mac2 = &mac2FlagValue

		retAdded = true
	}

	return nil, retAdded
}
