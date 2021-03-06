// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for ActiveQuestionnaire

// register flags to command
func registerModelActiveQuestionnaireFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerActiveQuestionnaireID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerActiveQuestionnaireID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `Required. ID of the currently active questionnaire. Null if no questionnaire is active.`

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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelActiveQuestionnaireFlags(depth int, m *models.ActiveQuestionnaire, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, idAdded := retrieveActiveQuestionnaireIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	return nil, retAdded
}

func retrieveActiveQuestionnaireIDFlags(depth int, m *models.ActiveQuestionnaire, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.ID = &idFlagValue

		retAdded = true
	}

	return nil, retAdded
}
