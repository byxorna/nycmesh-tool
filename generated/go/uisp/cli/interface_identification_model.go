// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for InterfaceIdentification

// register flags to command
func registerModelInterfaceIdentificationFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerInterfaceIdentificationDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInterfaceIdentificationDisplayName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInterfaceIdentificationMac(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInterfaceIdentificationName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInterfaceIdentificationPosition(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInterfaceIdentificationType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerInterfaceIdentificationDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := `Nullable string.`

	var descriptionFlagName string
	if cmdPrefix == "" {
		descriptionFlagName = "description"
	} else {
		descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
	}

	var descriptionFlagDefault string

	_ = cmd.PersistentFlags().String(descriptionFlagName, descriptionFlagDefault, descriptionDescription)

	return nil
}

func registerInterfaceIdentificationDisplayName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	displayNameDescription := `Computed display name from name and description`

	var displayNameFlagName string
	if cmdPrefix == "" {
		displayNameFlagName = "displayName"
	} else {
		displayNameFlagName = fmt.Sprintf("%v.displayName", cmdPrefix)
	}

	var displayNameFlagDefault string

	_ = cmd.PersistentFlags().String(displayNameFlagName, displayNameFlagDefault, displayNameDescription)

	return nil
}

func registerInterfaceIdentificationMac(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	macDescription := ``

	var macFlagName string
	if cmdPrefix == "" {
		macFlagName = "mac"
	} else {
		macFlagName = fmt.Sprintf("%v.mac", cmdPrefix)
	}

	var macFlagDefault string

	_ = cmd.PersistentFlags().String(macFlagName, macFlagDefault, macDescription)

	return nil
}

func registerInterfaceIdentificationName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Interface name.`

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerInterfaceIdentificationPosition(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	positionDescription := `Physical port position.`

	var positionFlagName string
	if cmdPrefix == "" {
		positionFlagName = "position"
	} else {
		positionFlagName = fmt.Sprintf("%v.position", cmdPrefix)
	}

	var positionFlagDefault float64

	_ = cmd.PersistentFlags().Float64(positionFlagName, positionFlagDefault, positionDescription)

	return nil
}

func registerInterfaceIdentificationType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := ``

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelInterfaceIdentificationFlags(depth int, m *models.InterfaceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, descriptionAdded := retrieveInterfaceIdentificationDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, displayNameAdded := retrieveInterfaceIdentificationDisplayNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || displayNameAdded

	err, macAdded := retrieveInterfaceIdentificationMacFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || macAdded

	err, nameAdded := retrieveInterfaceIdentificationNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, positionAdded := retrieveInterfaceIdentificationPositionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || positionAdded

	err, typeAdded := retrieveInterfaceIdentificationTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveInterfaceIdentificationDescriptionFlags(depth int, m *models.InterfaceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	descriptionFlagName := fmt.Sprintf("%v.description", cmdPrefix)
	if cmd.Flags().Changed(descriptionFlagName) {

		var descriptionFlagName string
		if cmdPrefix == "" {
			descriptionFlagName = "description"
		} else {
			descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
		}

		descriptionFlagValue, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			return err, false
		}
		m.Description = descriptionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveInterfaceIdentificationDisplayNameFlags(depth int, m *models.InterfaceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	displayNameFlagName := fmt.Sprintf("%v.displayName", cmdPrefix)
	if cmd.Flags().Changed(displayNameFlagName) {

		var displayNameFlagName string
		if cmdPrefix == "" {
			displayNameFlagName = "displayName"
		} else {
			displayNameFlagName = fmt.Sprintf("%v.displayName", cmdPrefix)
		}

		displayNameFlagValue, err := cmd.Flags().GetString(displayNameFlagName)
		if err != nil {
			return err, false
		}
		m.DisplayName = displayNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveInterfaceIdentificationMacFlags(depth int, m *models.InterfaceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	macFlagName := fmt.Sprintf("%v.mac", cmdPrefix)
	if cmd.Flags().Changed(macFlagName) {

		var macFlagName string
		if cmdPrefix == "" {
			macFlagName = "mac"
		} else {
			macFlagName = fmt.Sprintf("%v.mac", cmdPrefix)
		}

		macFlagValue, err := cmd.Flags().GetString(macFlagName)
		if err != nil {
			return err, false
		}
		m.Mac = macFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveInterfaceIdentificationNameFlags(depth int, m *models.InterfaceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveInterfaceIdentificationPositionFlags(depth int, m *models.InterfaceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	positionFlagName := fmt.Sprintf("%v.position", cmdPrefix)
	if cmd.Flags().Changed(positionFlagName) {

		var positionFlagName string
		if cmdPrefix == "" {
			positionFlagName = "position"
		} else {
			positionFlagName = fmt.Sprintf("%v.position", cmdPrefix)
		}

		positionFlagValue, err := cmd.Flags().GetFloat64(positionFlagName)
		if err != nil {
			return err, false
		}
		m.Position = positionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveInterfaceIdentificationTypeFlags(depth int, m *models.InterfaceIdentification, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	typeFlagName := fmt.Sprintf("%v.type", cmdPrefix)
	if cmd.Flags().Changed(typeFlagName) {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
