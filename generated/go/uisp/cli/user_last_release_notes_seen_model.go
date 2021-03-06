// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for UserLastReleaseNotesSeen

// register flags to command
func registerModelUserLastReleaseNotesSeenFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUserLastReleaseNotesSeenLastReleaseNotesSeen(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUserLastReleaseNotesSeenLastReleaseNotesSeen(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	lastReleaseNotesSeenDescription := `Required. 1.0.0`

	var lastReleaseNotesSeenFlagName string
	if cmdPrefix == "" {
		lastReleaseNotesSeenFlagName = "lastReleaseNotesSeen"
	} else {
		lastReleaseNotesSeenFlagName = fmt.Sprintf("%v.lastReleaseNotesSeen", cmdPrefix)
	}

	var lastReleaseNotesSeenFlagDefault string

	_ = cmd.PersistentFlags().String(lastReleaseNotesSeenFlagName, lastReleaseNotesSeenFlagDefault, lastReleaseNotesSeenDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUserLastReleaseNotesSeenFlags(depth int, m *models.UserLastReleaseNotesSeen, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, lastReleaseNotesSeenAdded := retrieveUserLastReleaseNotesSeenLastReleaseNotesSeenFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || lastReleaseNotesSeenAdded

	return nil, retAdded
}

func retrieveUserLastReleaseNotesSeenLastReleaseNotesSeenFlags(depth int, m *models.UserLastReleaseNotesSeen, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	lastReleaseNotesSeenFlagName := fmt.Sprintf("%v.lastReleaseNotesSeen", cmdPrefix)
	if cmd.Flags().Changed(lastReleaseNotesSeenFlagName) {

		var lastReleaseNotesSeenFlagName string
		if cmdPrefix == "" {
			lastReleaseNotesSeenFlagName = "lastReleaseNotesSeen"
		} else {
			lastReleaseNotesSeenFlagName = fmt.Sprintf("%v.lastReleaseNotesSeen", cmdPrefix)
		}

		lastReleaseNotesSeenFlagValue, err := cmd.Flags().GetString(lastReleaseNotesSeenFlagName)
		if err != nil {
			return err, false
		}
		m.LastReleaseNotesSeen = &lastReleaseNotesSeenFlagValue

		retAdded = true
	}

	return nil, retAdded
}
