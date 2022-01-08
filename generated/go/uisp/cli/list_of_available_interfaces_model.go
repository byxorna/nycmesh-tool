// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for ListOfAvailableInterfaces

// register flags to command
func registerModelListOfAvailableInterfacesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerListOfAvailableInterfacesDataLinkID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerListOfAvailableInterfacesDeviceID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerListOfAvailableInterfacesInterfaces(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerListOfAvailableInterfacesDataLinkID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	dataLinkIdDescription := ``

	var dataLinkIdFlagName string
	if cmdPrefix == "" {
		dataLinkIdFlagName = "dataLinkId"
	} else {
		dataLinkIdFlagName = fmt.Sprintf("%v.dataLinkId", cmdPrefix)
	}

	var dataLinkIdFlagDefault string

	_ = cmd.PersistentFlags().String(dataLinkIdFlagName, dataLinkIdFlagDefault, dataLinkIdDescription)

	return nil
}

func registerListOfAvailableInterfacesDeviceID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	deviceIdDescription := `Required. Device ID.`

	var deviceIdFlagName string
	if cmdPrefix == "" {
		deviceIdFlagName = "deviceId"
	} else {
		deviceIdFlagName = fmt.Sprintf("%v.deviceId", cmdPrefix)
	}

	var deviceIdFlagDefault string

	_ = cmd.PersistentFlags().String(deviceIdFlagName, deviceIdFlagDefault, deviceIdDescription)

	return nil
}

func registerListOfAvailableInterfacesInterfaces(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: interfaces InterfaceList array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelListOfAvailableInterfacesFlags(depth int, m *models.ListOfAvailableInterfaces, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataLinkIdAdded := retrieveListOfAvailableInterfacesDataLinkIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataLinkIdAdded

	err, deviceIdAdded := retrieveListOfAvailableInterfacesDeviceIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceIdAdded

	err, interfacesAdded := retrieveListOfAvailableInterfacesInterfacesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || interfacesAdded

	return nil, retAdded
}

func retrieveListOfAvailableInterfacesDataLinkIDFlags(depth int, m *models.ListOfAvailableInterfaces, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataLinkIdFlagName := fmt.Sprintf("%v.dataLinkId", cmdPrefix)
	if cmd.Flags().Changed(dataLinkIdFlagName) {

		var dataLinkIdFlagName string
		if cmdPrefix == "" {
			dataLinkIdFlagName = "dataLinkId"
		} else {
			dataLinkIdFlagName = fmt.Sprintf("%v.dataLinkId", cmdPrefix)
		}

		dataLinkIdFlagValue, err := cmd.Flags().GetString(dataLinkIdFlagName)
		if err != nil {
			return err, false
		}
		m.DataLinkID = dataLinkIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveListOfAvailableInterfacesDeviceIDFlags(depth int, m *models.ListOfAvailableInterfaces, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deviceIdFlagName := fmt.Sprintf("%v.deviceId", cmdPrefix)
	if cmd.Flags().Changed(deviceIdFlagName) {

		var deviceIdFlagName string
		if cmdPrefix == "" {
			deviceIdFlagName = "deviceId"
		} else {
			deviceIdFlagName = fmt.Sprintf("%v.deviceId", cmdPrefix)
		}

		deviceIdFlagValue, err := cmd.Flags().GetString(deviceIdFlagName)
		if err != nil {
			return err, false
		}
		m.DeviceID = &deviceIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveListOfAvailableInterfacesInterfacesFlags(depth int, m *models.ListOfAvailableInterfaces, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	interfacesFlagName := fmt.Sprintf("%v.interfaces", cmdPrefix)
	if cmd.Flags().Changed(interfacesFlagName) {
		// warning: interfaces array type InterfaceList is not supported by go-swagger cli yet
	}

	return nil, retAdded
}