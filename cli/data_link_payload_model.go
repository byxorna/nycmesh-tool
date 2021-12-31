// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for DataLinkPayload

// register flags to command
func registerModelDataLinkPayloadFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDataLinkPayloadDeviceIDFrom(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDataLinkPayloadDeviceIDTo(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDataLinkPayloadID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDataLinkPayloadInterfaceIDFrom(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDataLinkPayloadInterfaceIDTo(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDataLinkPayloadDeviceIDFrom(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	deviceIdFromDescription := `Required. `

	var deviceIdFromFlagName string
	if cmdPrefix == "" {
		deviceIdFromFlagName = "deviceIdFrom"
	} else {
		deviceIdFromFlagName = fmt.Sprintf("%v.deviceIdFrom", cmdPrefix)
	}

	var deviceIdFromFlagDefault string

	_ = cmd.PersistentFlags().String(deviceIdFromFlagName, deviceIdFromFlagDefault, deviceIdFromDescription)

	return nil
}

func registerDataLinkPayloadDeviceIDTo(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	deviceIdToDescription := `Required. `

	var deviceIdToFlagName string
	if cmdPrefix == "" {
		deviceIdToFlagName = "deviceIdTo"
	} else {
		deviceIdToFlagName = fmt.Sprintf("%v.deviceIdTo", cmdPrefix)
	}

	var deviceIdToFlagDefault string

	_ = cmd.PersistentFlags().String(deviceIdToFlagName, deviceIdToFlagDefault, deviceIdToDescription)

	return nil
}

func registerDataLinkPayloadID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerDataLinkPayloadInterfaceIDFrom(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	interfaceIdFromDescription := `Required. `

	var interfaceIdFromFlagName string
	if cmdPrefix == "" {
		interfaceIdFromFlagName = "interfaceIdFrom"
	} else {
		interfaceIdFromFlagName = fmt.Sprintf("%v.interfaceIdFrom", cmdPrefix)
	}

	var interfaceIdFromFlagDefault string

	_ = cmd.PersistentFlags().String(interfaceIdFromFlagName, interfaceIdFromFlagDefault, interfaceIdFromDescription)

	return nil
}

func registerDataLinkPayloadInterfaceIDTo(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	interfaceIdToDescription := `Required. `

	var interfaceIdToFlagName string
	if cmdPrefix == "" {
		interfaceIdToFlagName = "interfaceIdTo"
	} else {
		interfaceIdToFlagName = fmt.Sprintf("%v.interfaceIdTo", cmdPrefix)
	}

	var interfaceIdToFlagDefault string

	_ = cmd.PersistentFlags().String(interfaceIdToFlagName, interfaceIdToFlagDefault, interfaceIdToDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDataLinkPayloadFlags(depth int, m *models.DataLinkPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, deviceIdFromAdded := retrieveDataLinkPayloadDeviceIDFromFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceIdFromAdded

	err, deviceIdToAdded := retrieveDataLinkPayloadDeviceIDToFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceIdToAdded

	err, idAdded := retrieveDataLinkPayloadIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, interfaceIdFromAdded := retrieveDataLinkPayloadInterfaceIDFromFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || interfaceIdFromAdded

	err, interfaceIdToAdded := retrieveDataLinkPayloadInterfaceIDToFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || interfaceIdToAdded

	return nil, retAdded
}

func retrieveDataLinkPayloadDeviceIDFromFlags(depth int, m *models.DataLinkPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deviceIdFromFlagName := fmt.Sprintf("%v.deviceIdFrom", cmdPrefix)
	if cmd.Flags().Changed(deviceIdFromFlagName) {

		var deviceIdFromFlagName string
		if cmdPrefix == "" {
			deviceIdFromFlagName = "deviceIdFrom"
		} else {
			deviceIdFromFlagName = fmt.Sprintf("%v.deviceIdFrom", cmdPrefix)
		}

		deviceIdFromFlagValue, err := cmd.Flags().GetString(deviceIdFromFlagName)
		if err != nil {
			return err, false
		}
		m.DeviceIDFrom = &deviceIdFromFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDataLinkPayloadDeviceIDToFlags(depth int, m *models.DataLinkPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deviceIdToFlagName := fmt.Sprintf("%v.deviceIdTo", cmdPrefix)
	if cmd.Flags().Changed(deviceIdToFlagName) {

		var deviceIdToFlagName string
		if cmdPrefix == "" {
			deviceIdToFlagName = "deviceIdTo"
		} else {
			deviceIdToFlagName = fmt.Sprintf("%v.deviceIdTo", cmdPrefix)
		}

		deviceIdToFlagValue, err := cmd.Flags().GetString(deviceIdToFlagName)
		if err != nil {
			return err, false
		}
		m.DeviceIDTo = &deviceIdToFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDataLinkPayloadIDFlags(depth int, m *models.DataLinkPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDataLinkPayloadInterfaceIDFromFlags(depth int, m *models.DataLinkPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	interfaceIdFromFlagName := fmt.Sprintf("%v.interfaceIdFrom", cmdPrefix)
	if cmd.Flags().Changed(interfaceIdFromFlagName) {

		var interfaceIdFromFlagName string
		if cmdPrefix == "" {
			interfaceIdFromFlagName = "interfaceIdFrom"
		} else {
			interfaceIdFromFlagName = fmt.Sprintf("%v.interfaceIdFrom", cmdPrefix)
		}

		interfaceIdFromFlagValue, err := cmd.Flags().GetString(interfaceIdFromFlagName)
		if err != nil {
			return err, false
		}
		m.InterfaceIDFrom = &interfaceIdFromFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDataLinkPayloadInterfaceIDToFlags(depth int, m *models.DataLinkPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	interfaceIdToFlagName := fmt.Sprintf("%v.interfaceIdTo", cmdPrefix)
	if cmd.Flags().Changed(interfaceIdToFlagName) {

		var interfaceIdToFlagName string
		if cmdPrefix == "" {
			interfaceIdToFlagName = "interfaceIdTo"
		} else {
			interfaceIdToFlagName = fmt.Sprintf("%v.interfaceIdTo", cmdPrefix)
		}

		interfaceIdToFlagValue, err := cmd.Flags().GetString(interfaceIdToFlagName)
		if err != nil {
			return err, false
		}
		m.InterfaceIDTo = &interfaceIdToFlagValue

		retAdded = true
	}

	return nil, retAdded
}
