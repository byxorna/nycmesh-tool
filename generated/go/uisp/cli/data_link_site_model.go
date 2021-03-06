// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for DataLinkSite

// register flags to command
func registerModelDataLinkSiteFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDataLinkSiteIdentification(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDataLinkSiteIdentification(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var identificationFlagName string
	if cmdPrefix == "" {
		identificationFlagName = "identification"
	} else {
		identificationFlagName = fmt.Sprintf("%v.identification", cmdPrefix)
	}

	if err := registerModelIdentificationFlags(depth+1, identificationFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDataLinkSiteFlags(depth int, m *models.DataLinkSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, identificationAdded := retrieveDataLinkSiteIdentificationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || identificationAdded

	return nil, retAdded
}

func retrieveDataLinkSiteIdentificationFlags(depth int, m *models.DataLinkSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	identificationFlagName := fmt.Sprintf("%v.identification", cmdPrefix)
	if cmd.Flags().Changed(identificationFlagName) {
		// info: complex object identification Identification is retrieved outside this Changed() block
	}
	identificationFlagValue := m.Identification
	if swag.IsZero(identificationFlagValue) {
		identificationFlagValue = &models.Identification{}
	}

	err, identificationAdded := retrieveModelIdentificationFlags(depth+1, identificationFlagValue, identificationFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || identificationAdded
	if identificationAdded {
		m.Identification = identificationFlagValue
	}

	return nil, retAdded
}
