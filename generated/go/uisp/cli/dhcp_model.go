// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for Dhcp

// register flags to command
func registerModelDhcpFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDhcpIgnore(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDhcpInterface(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDhcpLeaseTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDhcpRangeEnd(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDhcpRangeStart(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDhcpIgnore(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ignoreDescription := `Required. `

	var ignoreFlagName string
	if cmdPrefix == "" {
		ignoreFlagName = "ignore"
	} else {
		ignoreFlagName = fmt.Sprintf("%v.ignore", cmdPrefix)
	}

	var ignoreFlagDefault bool

	_ = cmd.PersistentFlags().Bool(ignoreFlagName, ignoreFlagDefault, ignoreDescription)

	return nil
}

func registerDhcpInterface(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	interfaceDescription := ``

	var interfaceFlagName string
	if cmdPrefix == "" {
		interfaceFlagName = "interface"
	} else {
		interfaceFlagName = fmt.Sprintf("%v.interface", cmdPrefix)
	}

	var interfaceFlagDefault string

	_ = cmd.PersistentFlags().String(interfaceFlagName, interfaceFlagDefault, interfaceDescription)

	return nil
}

func registerDhcpLeaseTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	leaseTimeDescription := ``

	var leaseTimeFlagName string
	if cmdPrefix == "" {
		leaseTimeFlagName = "leaseTime"
	} else {
		leaseTimeFlagName = fmt.Sprintf("%v.leaseTime", cmdPrefix)
	}

	var leaseTimeFlagDefault string

	_ = cmd.PersistentFlags().String(leaseTimeFlagName, leaseTimeFlagDefault, leaseTimeDescription)

	return nil
}

func registerDhcpRangeEnd(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	rangeEndDescription := `Required. `

	var rangeEndFlagName string
	if cmdPrefix == "" {
		rangeEndFlagName = "rangeEnd"
	} else {
		rangeEndFlagName = fmt.Sprintf("%v.rangeEnd", cmdPrefix)
	}

	var rangeEndFlagDefault string

	_ = cmd.PersistentFlags().String(rangeEndFlagName, rangeEndFlagDefault, rangeEndDescription)

	return nil
}

func registerDhcpRangeStart(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	rangeStartDescription := `Required. `

	var rangeStartFlagName string
	if cmdPrefix == "" {
		rangeStartFlagName = "rangeStart"
	} else {
		rangeStartFlagName = fmt.Sprintf("%v.rangeStart", cmdPrefix)
	}

	var rangeStartFlagDefault string

	_ = cmd.PersistentFlags().String(rangeStartFlagName, rangeStartFlagDefault, rangeStartDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDhcpFlags(depth int, m *models.Dhcp, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, ignoreAdded := retrieveDhcpIgnoreFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ignoreAdded

	err, interfaceAdded := retrieveDhcpInterfaceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || interfaceAdded

	err, leaseTimeAdded := retrieveDhcpLeaseTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || leaseTimeAdded

	err, rangeEndAdded := retrieveDhcpRangeEndFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || rangeEndAdded

	err, rangeStartAdded := retrieveDhcpRangeStartFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || rangeStartAdded

	return nil, retAdded
}

func retrieveDhcpIgnoreFlags(depth int, m *models.Dhcp, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ignoreFlagName := fmt.Sprintf("%v.ignore", cmdPrefix)
	if cmd.Flags().Changed(ignoreFlagName) {

		var ignoreFlagName string
		if cmdPrefix == "" {
			ignoreFlagName = "ignore"
		} else {
			ignoreFlagName = fmt.Sprintf("%v.ignore", cmdPrefix)
		}

		ignoreFlagValue, err := cmd.Flags().GetBool(ignoreFlagName)
		if err != nil {
			return err, false
		}
		m.Ignore = &ignoreFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDhcpInterfaceFlags(depth int, m *models.Dhcp, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	interfaceFlagName := fmt.Sprintf("%v.interface", cmdPrefix)
	if cmd.Flags().Changed(interfaceFlagName) {

		var interfaceFlagName string
		if cmdPrefix == "" {
			interfaceFlagName = "interface"
		} else {
			interfaceFlagName = fmt.Sprintf("%v.interface", cmdPrefix)
		}

		interfaceFlagValue, err := cmd.Flags().GetString(interfaceFlagName)
		if err != nil {
			return err, false
		}
		m.Interface = interfaceFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDhcpLeaseTimeFlags(depth int, m *models.Dhcp, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	leaseTimeFlagName := fmt.Sprintf("%v.leaseTime", cmdPrefix)
	if cmd.Flags().Changed(leaseTimeFlagName) {

		var leaseTimeFlagName string
		if cmdPrefix == "" {
			leaseTimeFlagName = "leaseTime"
		} else {
			leaseTimeFlagName = fmt.Sprintf("%v.leaseTime", cmdPrefix)
		}

		leaseTimeFlagValue, err := cmd.Flags().GetString(leaseTimeFlagName)
		if err != nil {
			return err, false
		}
		m.LeaseTime = leaseTimeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDhcpRangeEndFlags(depth int, m *models.Dhcp, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	rangeEndFlagName := fmt.Sprintf("%v.rangeEnd", cmdPrefix)
	if cmd.Flags().Changed(rangeEndFlagName) {

		var rangeEndFlagName string
		if cmdPrefix == "" {
			rangeEndFlagName = "rangeEnd"
		} else {
			rangeEndFlagName = fmt.Sprintf("%v.rangeEnd", cmdPrefix)
		}

		rangeEndFlagValue, err := cmd.Flags().GetString(rangeEndFlagName)
		if err != nil {
			return err, false
		}
		m.RangeEnd = &rangeEndFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDhcpRangeStartFlags(depth int, m *models.Dhcp, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	rangeStartFlagName := fmt.Sprintf("%v.rangeStart", cmdPrefix)
	if cmd.Flags().Changed(rangeStartFlagName) {

		var rangeStartFlagName string
		if cmdPrefix == "" {
			rangeStartFlagName = "rangeStart"
		} else {
			rangeStartFlagName = fmt.Sprintf("%v.rangeStart", cmdPrefix)
		}

		rangeStartFlagValue, err := cmd.Flags().GetString(rangeStartFlagName)
		if err != nil {
			return err, false
		}
		m.RangeStart = &rangeStartFlagValue

		retAdded = true
	}

	return nil, retAdded
}
