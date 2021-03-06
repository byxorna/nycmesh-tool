// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for ImportUsersResultSchema

// register flags to command
func registerModelImportUsersResultSchemaFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerImportUsersResultSchemaNewUsername(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImportUsersResultSchemaUcrmID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImportUsersResultSchemaUnmsID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerImportUsersResultSchemaNewUsername(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	newUsernameDescription := `New username in case that the original CRM username is already taken in UISP.`

	var newUsernameFlagName string
	if cmdPrefix == "" {
		newUsernameFlagName = "newUsername"
	} else {
		newUsernameFlagName = fmt.Sprintf("%v.newUsername", cmdPrefix)
	}

	var newUsernameFlagDefault string

	_ = cmd.PersistentFlags().String(newUsernameFlagName, newUsernameFlagDefault, newUsernameDescription)

	return nil
}

func registerImportUsersResultSchemaUcrmID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ucrmIdDescription := `Required. User's id in CRM.`

	var ucrmIdFlagName string
	if cmdPrefix == "" {
		ucrmIdFlagName = "ucrmId"
	} else {
		ucrmIdFlagName = fmt.Sprintf("%v.ucrmId", cmdPrefix)
	}

	var ucrmIdFlagDefault string

	_ = cmd.PersistentFlags().String(ucrmIdFlagName, ucrmIdFlagDefault, ucrmIdDescription)

	return nil
}

func registerImportUsersResultSchemaUnmsID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	unmsIdDescription := `Required. User's id in UISP.`

	var unmsIdFlagName string
	if cmdPrefix == "" {
		unmsIdFlagName = "unmsId"
	} else {
		unmsIdFlagName = fmt.Sprintf("%v.unmsId", cmdPrefix)
	}

	var unmsIdFlagDefault string

	_ = cmd.PersistentFlags().String(unmsIdFlagName, unmsIdFlagDefault, unmsIdDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelImportUsersResultSchemaFlags(depth int, m *models.ImportUsersResultSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, newUsernameAdded := retrieveImportUsersResultSchemaNewUsernameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || newUsernameAdded

	err, ucrmIdAdded := retrieveImportUsersResultSchemaUcrmIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ucrmIdAdded

	err, unmsIdAdded := retrieveImportUsersResultSchemaUnmsIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || unmsIdAdded

	return nil, retAdded
}

func retrieveImportUsersResultSchemaNewUsernameFlags(depth int, m *models.ImportUsersResultSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	newUsernameFlagName := fmt.Sprintf("%v.newUsername", cmdPrefix)
	if cmd.Flags().Changed(newUsernameFlagName) {

		var newUsernameFlagName string
		if cmdPrefix == "" {
			newUsernameFlagName = "newUsername"
		} else {
			newUsernameFlagName = fmt.Sprintf("%v.newUsername", cmdPrefix)
		}

		newUsernameFlagValue, err := cmd.Flags().GetString(newUsernameFlagName)
		if err != nil {
			return err, false
		}
		m.NewUsername = newUsernameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImportUsersResultSchemaUcrmIDFlags(depth int, m *models.ImportUsersResultSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ucrmIdFlagName := fmt.Sprintf("%v.ucrmId", cmdPrefix)
	if cmd.Flags().Changed(ucrmIdFlagName) {

		var ucrmIdFlagName string
		if cmdPrefix == "" {
			ucrmIdFlagName = "ucrmId"
		} else {
			ucrmIdFlagName = fmt.Sprintf("%v.ucrmId", cmdPrefix)
		}

		ucrmIdFlagValue, err := cmd.Flags().GetString(ucrmIdFlagName)
		if err != nil {
			return err, false
		}
		m.UcrmID = &ucrmIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImportUsersResultSchemaUnmsIDFlags(depth int, m *models.ImportUsersResultSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	unmsIdFlagName := fmt.Sprintf("%v.unmsId", cmdPrefix)
	if cmd.Flags().Changed(unmsIdFlagName) {

		var unmsIdFlagName string
		if cmdPrefix == "" {
			unmsIdFlagName = "unmsId"
		} else {
			unmsIdFlagName = fmt.Sprintf("%v.unmsId", cmdPrefix)
		}

		unmsIdFlagValue, err := cmd.Flags().GetString(unmsIdFlagName)
		if err != nil {
			return err, false
		}
		m.UnmsID = &unmsIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}
