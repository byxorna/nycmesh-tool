// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"
  "github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for PasswordStrengthMetadata

// register flags to command
func registerModelPasswordStrengthMetadataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPasswordStrengthMetadataCalcTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPasswordStrengthMetadataCrackTimesDisplay(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPasswordStrengthMetadataCrackTimesSeconds(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPasswordStrengthMetadataFeedback(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPasswordStrengthMetadataGuesses(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPasswordStrengthMetadataGuessesLog10(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPasswordStrengthMetadataPassword(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPasswordStrengthMetadataScore(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPasswordStrengthMetadataSequence(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPasswordStrengthMetadataCalcTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	calcTimeDescription := `Required. `

	var calcTimeFlagName string
	if cmdPrefix == "" {
		calcTimeFlagName = "calc_time"
	} else {
		calcTimeFlagName = fmt.Sprintf("%v.calc_time", cmdPrefix)
	}

	var calcTimeFlagDefault float64

	_ = cmd.PersistentFlags().Float64(calcTimeFlagName, calcTimeFlagDefault, calcTimeDescription)

	return nil
}

func registerPasswordStrengthMetadataCrackTimesDisplay(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var crackTimesDisplayFlagName string
	if cmdPrefix == "" {
		crackTimesDisplayFlagName = "crack_times_display"
	} else {
		crackTimesDisplayFlagName = fmt.Sprintf("%v.crack_times_display", cmdPrefix)
	}

	if err := registerModelCrackTimesDisplayFlags(depth+1, crackTimesDisplayFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerPasswordStrengthMetadataCrackTimesSeconds(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var crackTimesSecondsFlagName string
	if cmdPrefix == "" {
		crackTimesSecondsFlagName = "crack_times_seconds"
	} else {
		crackTimesSecondsFlagName = fmt.Sprintf("%v.crack_times_seconds", cmdPrefix)
	}

	if err := registerModelCrackTimesSecondsFlags(depth+1, crackTimesSecondsFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerPasswordStrengthMetadataFeedback(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var feedbackFlagName string
	if cmdPrefix == "" {
		feedbackFlagName = "feedback"
	} else {
		feedbackFlagName = fmt.Sprintf("%v.feedback", cmdPrefix)
	}

	if err := registerModelFeedbackFlags(depth+1, feedbackFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerPasswordStrengthMetadataGuesses(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	guessesDescription := `Required. `

	var guessesFlagName string
	if cmdPrefix == "" {
		guessesFlagName = "guesses"
	} else {
		guessesFlagName = fmt.Sprintf("%v.guesses", cmdPrefix)
	}

	var guessesFlagDefault float64

	_ = cmd.PersistentFlags().Float64(guessesFlagName, guessesFlagDefault, guessesDescription)

	return nil
}

func registerPasswordStrengthMetadataGuessesLog10(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	guessesLog10Description := `Required. `

	var guessesLog10FlagName string
	if cmdPrefix == "" {
		guessesLog10FlagName = "guesses_log10"
	} else {
		guessesLog10FlagName = fmt.Sprintf("%v.guesses_log10", cmdPrefix)
	}

	var guessesLog10FlagDefault float64

	_ = cmd.PersistentFlags().Float64(guessesLog10FlagName, guessesLog10FlagDefault, guessesLog10Description)

	return nil
}

func registerPasswordStrengthMetadataPassword(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	passwordDescription := `Required. `

	var passwordFlagName string
	if cmdPrefix == "" {
		passwordFlagName = "password"
	} else {
		passwordFlagName = fmt.Sprintf("%v.password", cmdPrefix)
	}

	var passwordFlagDefault string

	_ = cmd.PersistentFlags().String(passwordFlagName, passwordFlagDefault, passwordDescription)

	return nil
}

func registerPasswordStrengthMetadataScore(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	scoreDescription := `Required. `

	var scoreFlagName string
	if cmdPrefix == "" {
		scoreFlagName = "score"
	} else {
		scoreFlagName = fmt.Sprintf("%v.score", cmdPrefix)
	}

	var scoreFlagDefault float64

	_ = cmd.PersistentFlags().Float64(scoreFlagName, scoreFlagDefault, scoreDescription)

	return nil
}

func registerPasswordStrengthMetadataSequence(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: sequence Sequence array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPasswordStrengthMetadataFlags(depth int, m *models.PasswordStrengthMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, calcTimeAdded := retrievePasswordStrengthMetadataCalcTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || calcTimeAdded

	err, crackTimesDisplayAdded := retrievePasswordStrengthMetadataCrackTimesDisplayFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || crackTimesDisplayAdded

	err, crackTimesSecondsAdded := retrievePasswordStrengthMetadataCrackTimesSecondsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || crackTimesSecondsAdded

	err, feedbackAdded := retrievePasswordStrengthMetadataFeedbackFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || feedbackAdded

	err, guessesAdded := retrievePasswordStrengthMetadataGuessesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || guessesAdded

	err, guessesLog10Added := retrievePasswordStrengthMetadataGuessesLog10Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || guessesLog10Added

	err, passwordAdded := retrievePasswordStrengthMetadataPasswordFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || passwordAdded

	err, scoreAdded := retrievePasswordStrengthMetadataScoreFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || scoreAdded

	err, sequenceAdded := retrievePasswordStrengthMetadataSequenceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sequenceAdded

	return nil, retAdded
}

func retrievePasswordStrengthMetadataCalcTimeFlags(depth int, m *models.PasswordStrengthMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	calcTimeFlagName := fmt.Sprintf("%v.calc_time", cmdPrefix)
	if cmd.Flags().Changed(calcTimeFlagName) {

		var calcTimeFlagName string
		if cmdPrefix == "" {
			calcTimeFlagName = "calc_time"
		} else {
			calcTimeFlagName = fmt.Sprintf("%v.calc_time", cmdPrefix)
		}

		calcTimeFlagValue, err := cmd.Flags().GetFloat64(calcTimeFlagName)
		if err != nil {
			return err, false
		}
		m.CalcTime = &calcTimeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePasswordStrengthMetadataCrackTimesDisplayFlags(depth int, m *models.PasswordStrengthMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	crackTimesDisplayFlagName := fmt.Sprintf("%v.crack_times_display", cmdPrefix)
	if cmd.Flags().Changed(crackTimesDisplayFlagName) {
		// info: complex object crack_times_display CrackTimesDisplay is retrieved outside this Changed() block
	}
	crackTimesDisplayFlagValue := m.CrackTimesDisplay
	if swag.IsZero(crackTimesDisplayFlagValue) {
		crackTimesDisplayFlagValue = &models.CrackTimesDisplay{}
	}

	err, crackTimesDisplayAdded := retrieveModelCrackTimesDisplayFlags(depth+1, crackTimesDisplayFlagValue, crackTimesDisplayFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || crackTimesDisplayAdded
	if crackTimesDisplayAdded {
		m.CrackTimesDisplay = crackTimesDisplayFlagValue
	}

	return nil, retAdded
}

func retrievePasswordStrengthMetadataCrackTimesSecondsFlags(depth int, m *models.PasswordStrengthMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	crackTimesSecondsFlagName := fmt.Sprintf("%v.crack_times_seconds", cmdPrefix)
	if cmd.Flags().Changed(crackTimesSecondsFlagName) {
		// info: complex object crack_times_seconds CrackTimesSeconds is retrieved outside this Changed() block
	}
	crackTimesSecondsFlagValue := m.CrackTimesSeconds
	if swag.IsZero(crackTimesSecondsFlagValue) {
		crackTimesSecondsFlagValue = &models.CrackTimesSeconds{}
	}

	err, crackTimesSecondsAdded := retrieveModelCrackTimesSecondsFlags(depth+1, crackTimesSecondsFlagValue, crackTimesSecondsFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || crackTimesSecondsAdded
	if crackTimesSecondsAdded {
		m.CrackTimesSeconds = crackTimesSecondsFlagValue
	}

	return nil, retAdded
}

func retrievePasswordStrengthMetadataFeedbackFlags(depth int, m *models.PasswordStrengthMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	feedbackFlagName := fmt.Sprintf("%v.feedback", cmdPrefix)
	if cmd.Flags().Changed(feedbackFlagName) {
		// info: complex object feedback Feedback is retrieved outside this Changed() block
	}
	feedbackFlagValue := m.Feedback
	if swag.IsZero(feedbackFlagValue) {
		feedbackFlagValue = &models.Feedback{}
	}

	err, feedbackAdded := retrieveModelFeedbackFlags(depth+1, feedbackFlagValue, feedbackFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || feedbackAdded
	if feedbackAdded {
		m.Feedback = feedbackFlagValue
	}

	return nil, retAdded
}

func retrievePasswordStrengthMetadataGuessesFlags(depth int, m *models.PasswordStrengthMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	guessesFlagName := fmt.Sprintf("%v.guesses", cmdPrefix)
	if cmd.Flags().Changed(guessesFlagName) {

		var guessesFlagName string
		if cmdPrefix == "" {
			guessesFlagName = "guesses"
		} else {
			guessesFlagName = fmt.Sprintf("%v.guesses", cmdPrefix)
		}

		guessesFlagValue, err := cmd.Flags().GetFloat64(guessesFlagName)
		if err != nil {
			return err, false
		}
		m.Guesses = &guessesFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePasswordStrengthMetadataGuessesLog10Flags(depth int, m *models.PasswordStrengthMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	guessesLog10FlagName := fmt.Sprintf("%v.guesses_log10", cmdPrefix)
	if cmd.Flags().Changed(guessesLog10FlagName) {

		var guessesLog10FlagName string
		if cmdPrefix == "" {
			guessesLog10FlagName = "guesses_log10"
		} else {
			guessesLog10FlagName = fmt.Sprintf("%v.guesses_log10", cmdPrefix)
		}

		guessesLog10FlagValue, err := cmd.Flags().GetFloat64(guessesLog10FlagName)
		if err != nil {
			return err, false
		}
		m.GuessesLog10 = &guessesLog10FlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePasswordStrengthMetadataPasswordFlags(depth int, m *models.PasswordStrengthMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	passwordFlagName := fmt.Sprintf("%v.password", cmdPrefix)
	if cmd.Flags().Changed(passwordFlagName) {

		var passwordFlagName string
		if cmdPrefix == "" {
			passwordFlagName = "password"
		} else {
			passwordFlagName = fmt.Sprintf("%v.password", cmdPrefix)
		}

		passwordFlagValue, err := cmd.Flags().GetString(passwordFlagName)
		if err != nil {
			return err, false
		}
		m.Password = &passwordFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePasswordStrengthMetadataScoreFlags(depth int, m *models.PasswordStrengthMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	scoreFlagName := fmt.Sprintf("%v.score", cmdPrefix)
	if cmd.Flags().Changed(scoreFlagName) {

		var scoreFlagName string
		if cmdPrefix == "" {
			scoreFlagName = "score"
		} else {
			scoreFlagName = fmt.Sprintf("%v.score", cmdPrefix)
		}

		scoreFlagValue, err := cmd.Flags().GetFloat64(scoreFlagName)
		if err != nil {
			return err, false
		}
		m.Score = &scoreFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePasswordStrengthMetadataSequenceFlags(depth int, m *models.PasswordStrengthMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sequenceFlagName := fmt.Sprintf("%v.sequence", cmdPrefix)
	if cmd.Flags().Changed(sequenceFlagName) {
		// warning: sequence array type Sequence is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
