// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for Model116

// register flags to command
func registerModelModel116Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel116Admin(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel116Admin(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var adminFlagName string
	if cmdPrefix == "" {
		adminFlagName = "admin"
	} else {
		adminFlagName = fmt.Sprintf("%v.admin", cmdPrefix)
	}

	if err := registerModelAdmin1Flags(depth+1, adminFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModel116Flags(depth int, m *models.Model116, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, adminAdded := retrieveModel116AdminFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || adminAdded

	return nil, retAdded
}

func retrieveModel116AdminFlags(depth int, m *models.Model116, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	adminFlagName := fmt.Sprintf("%v.admin", cmdPrefix)
	if cmd.Flags().Changed(adminFlagName) {
		// info: complex object admin Admin1 is retrieved outside this Changed() block
	}
	adminFlagValue := m.Admin
	if swag.IsZero(adminFlagValue) {
		adminFlagValue = &models.Admin1{}
	}

	err, adminAdded := retrieveModelAdmin1Flags(depth+1, adminFlagValue, adminFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || adminAdded
	if adminAdded {
		m.Admin = adminFlagValue
	}

	return nil, retAdded
}
