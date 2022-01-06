// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// Schema cli for Model102

// register flags to command
func registerModelModel102Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel102Description(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel102Entries(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel102Identification(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel102Description(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := ``

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

func registerModel102Entries(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: entries Entries array type is not supported by go-swagger cli yet

	return nil
}

func registerModel102Identification(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var identificationFlagName string
	if cmdPrefix == "" {
		identificationFlagName = "identification"
	} else {
		identificationFlagName = fmt.Sprintf("%v.identification", cmdPrefix)
	}

	if err := registerModelIdentification4Flags(depth+1, identificationFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModel102Flags(depth int, m *models.Model102, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, descriptionAdded := retrieveModel102DescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, entriesAdded := retrieveModel102EntriesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || entriesAdded

	err, identificationAdded := retrieveModel102IdentificationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || identificationAdded

	return nil, retAdded
}

func retrieveModel102DescriptionFlags(depth int, m *models.Model102, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel102EntriesFlags(depth int, m *models.Model102, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	entriesFlagName := fmt.Sprintf("%v.entries", cmdPrefix)
	if cmd.Flags().Changed(entriesFlagName) {
		// warning: entries array type Entries is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveModel102IdentificationFlags(depth int, m *models.Model102, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	identificationFlagName := fmt.Sprintf("%v.identification", cmdPrefix)
	if cmd.Flags().Changed(identificationFlagName) {
		// info: complex object identification Identification4 is retrieved outside this Changed() block
	}
	identificationFlagValue := m.Identification
	if swag.IsZero(identificationFlagValue) {
		identificationFlagValue = &models.Identification4{}
	}

	err, identificationAdded := retrieveModelIdentification4Flags(depth+1, identificationFlagValue, identificationFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || identificationAdded
	if identificationAdded {
		m.Identification = identificationFlagValue
	}

	return nil, retAdded
}
