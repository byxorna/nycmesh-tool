// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for Statistics

// register flags to command
func registerModelStatisticsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerStatisticsAirTimeScore(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStatisticsLinkScore(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStatisticsLinkScoreHint(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStatisticsScore(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStatisticsScoreMax(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStatisticsTheoreticalDownlinkCapacity(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStatisticsTheoreticalUplinkCapacity(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerStatisticsAirTimeScore(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	airTimeScoreDescription := `Score airTime 0 - 1, for example 0.33 is 33%.`

	var airTimeScoreFlagName string
	if cmdPrefix == "" {
		airTimeScoreFlagName = "airTimeScore"
	} else {
		airTimeScoreFlagName = fmt.Sprintf("%v.airTimeScore", cmdPrefix)
	}

	var airTimeScoreFlagDefault float64

	_ = cmd.PersistentFlags().Float64(airTimeScoreFlagName, airTimeScoreFlagDefault, airTimeScoreDescription)

	return nil
}

func registerStatisticsLinkScore(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	linkScoreDescription := `Result score 0 - 1, for example 0.33 is 33%.`

	var linkScoreFlagName string
	if cmdPrefix == "" {
		linkScoreFlagName = "linkScore"
	} else {
		linkScoreFlagName = fmt.Sprintf("%v.linkScore", cmdPrefix)
	}

	var linkScoreFlagDefault float64

	_ = cmd.PersistentFlags().Float64(linkScoreFlagName, linkScoreFlagDefault, linkScoreDescription)

	return nil
}

func registerStatisticsLinkScoreHint(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	linkScoreHintDescription := `Hint for better score.`

	var linkScoreHintFlagName string
	if cmdPrefix == "" {
		linkScoreHintFlagName = "linkScoreHint"
	} else {
		linkScoreHintFlagName = fmt.Sprintf("%v.linkScoreHint", cmdPrefix)
	}

	var linkScoreHintFlagDefault string

	_ = cmd.PersistentFlags().String(linkScoreHintFlagName, linkScoreHintFlagDefault, linkScoreHintDescription)

	return nil
}

func registerStatisticsScore(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	scoreDescription := `Score 0 - 1, for example 0.33 is 33%.`

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

func registerStatisticsScoreMax(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	scoreMaxDescription := `Score max 0 - 1, for example 0.33 is 33%.`

	var scoreMaxFlagName string
	if cmdPrefix == "" {
		scoreMaxFlagName = "scoreMax"
	} else {
		scoreMaxFlagName = fmt.Sprintf("%v.scoreMax", cmdPrefix)
	}

	var scoreMaxFlagDefault float64

	_ = cmd.PersistentFlags().Float64(scoreMaxFlagName, scoreMaxFlagDefault, scoreMaxDescription)

	return nil
}

func registerStatisticsTheoreticalDownlinkCapacity(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	theoreticalDownlinkCapacityDescription := `Theoretical downlink capacity.`

	var theoreticalDownlinkCapacityFlagName string
	if cmdPrefix == "" {
		theoreticalDownlinkCapacityFlagName = "theoreticalDownlinkCapacity"
	} else {
		theoreticalDownlinkCapacityFlagName = fmt.Sprintf("%v.theoreticalDownlinkCapacity", cmdPrefix)
	}

	var theoreticalDownlinkCapacityFlagDefault int64

	_ = cmd.PersistentFlags().Int64(theoreticalDownlinkCapacityFlagName, theoreticalDownlinkCapacityFlagDefault, theoreticalDownlinkCapacityDescription)

	return nil
}

func registerStatisticsTheoreticalUplinkCapacity(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	theoreticalUplinkCapacityDescription := `Theoretical uplink capacity.`

	var theoreticalUplinkCapacityFlagName string
	if cmdPrefix == "" {
		theoreticalUplinkCapacityFlagName = "theoreticalUplinkCapacity"
	} else {
		theoreticalUplinkCapacityFlagName = fmt.Sprintf("%v.theoreticalUplinkCapacity", cmdPrefix)
	}

	var theoreticalUplinkCapacityFlagDefault int64

	_ = cmd.PersistentFlags().Int64(theoreticalUplinkCapacityFlagName, theoreticalUplinkCapacityFlagDefault, theoreticalUplinkCapacityDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelStatisticsFlags(depth int, m *models.Statistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, airTimeScoreAdded := retrieveStatisticsAirTimeScoreFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || airTimeScoreAdded

	err, linkScoreAdded := retrieveStatisticsLinkScoreFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || linkScoreAdded

	err, linkScoreHintAdded := retrieveStatisticsLinkScoreHintFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || linkScoreHintAdded

	err, scoreAdded := retrieveStatisticsScoreFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || scoreAdded

	err, scoreMaxAdded := retrieveStatisticsScoreMaxFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || scoreMaxAdded

	err, theoreticalDownlinkCapacityAdded := retrieveStatisticsTheoreticalDownlinkCapacityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || theoreticalDownlinkCapacityAdded

	err, theoreticalUplinkCapacityAdded := retrieveStatisticsTheoreticalUplinkCapacityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || theoreticalUplinkCapacityAdded

	return nil, retAdded
}

func retrieveStatisticsAirTimeScoreFlags(depth int, m *models.Statistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	airTimeScoreFlagName := fmt.Sprintf("%v.airTimeScore", cmdPrefix)
	if cmd.Flags().Changed(airTimeScoreFlagName) {

		var airTimeScoreFlagName string
		if cmdPrefix == "" {
			airTimeScoreFlagName = "airTimeScore"
		} else {
			airTimeScoreFlagName = fmt.Sprintf("%v.airTimeScore", cmdPrefix)
		}

		airTimeScoreFlagValue, err := cmd.Flags().GetFloat64(airTimeScoreFlagName)
		if err != nil {
			return err, false
		}
		m.AirTimeScore = &airTimeScoreFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStatisticsLinkScoreFlags(depth int, m *models.Statistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	linkScoreFlagName := fmt.Sprintf("%v.linkScore", cmdPrefix)
	if cmd.Flags().Changed(linkScoreFlagName) {

		var linkScoreFlagName string
		if cmdPrefix == "" {
			linkScoreFlagName = "linkScore"
		} else {
			linkScoreFlagName = fmt.Sprintf("%v.linkScore", cmdPrefix)
		}

		linkScoreFlagValue, err := cmd.Flags().GetFloat64(linkScoreFlagName)
		if err != nil {
			return err, false
		}
		m.LinkScore = &linkScoreFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStatisticsLinkScoreHintFlags(depth int, m *models.Statistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	linkScoreHintFlagName := fmt.Sprintf("%v.linkScoreHint", cmdPrefix)
	if cmd.Flags().Changed(linkScoreHintFlagName) {

		var linkScoreHintFlagName string
		if cmdPrefix == "" {
			linkScoreHintFlagName = "linkScoreHint"
		} else {
			linkScoreHintFlagName = fmt.Sprintf("%v.linkScoreHint", cmdPrefix)
		}

		linkScoreHintFlagValue, err := cmd.Flags().GetString(linkScoreHintFlagName)
		if err != nil {
			return err, false
		}
		m.LinkScoreHint = linkScoreHintFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStatisticsScoreFlags(depth int, m *models.Statistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveStatisticsScoreMaxFlags(depth int, m *models.Statistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	scoreMaxFlagName := fmt.Sprintf("%v.scoreMax", cmdPrefix)
	if cmd.Flags().Changed(scoreMaxFlagName) {

		var scoreMaxFlagName string
		if cmdPrefix == "" {
			scoreMaxFlagName = "scoreMax"
		} else {
			scoreMaxFlagName = fmt.Sprintf("%v.scoreMax", cmdPrefix)
		}

		scoreMaxFlagValue, err := cmd.Flags().GetFloat64(scoreMaxFlagName)
		if err != nil {
			return err, false
		}
		m.ScoreMax = &scoreMaxFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStatisticsTheoreticalDownlinkCapacityFlags(depth int, m *models.Statistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	theoreticalDownlinkCapacityFlagName := fmt.Sprintf("%v.theoreticalDownlinkCapacity", cmdPrefix)
	if cmd.Flags().Changed(theoreticalDownlinkCapacityFlagName) {

		var theoreticalDownlinkCapacityFlagName string
		if cmdPrefix == "" {
			theoreticalDownlinkCapacityFlagName = "theoreticalDownlinkCapacity"
		} else {
			theoreticalDownlinkCapacityFlagName = fmt.Sprintf("%v.theoreticalDownlinkCapacity", cmdPrefix)
		}

		theoreticalDownlinkCapacityFlagValue, err := cmd.Flags().GetInt64(theoreticalDownlinkCapacityFlagName)
		if err != nil {
			return err, false
		}
		m.TheoreticalDownlinkCapacity = theoreticalDownlinkCapacityFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStatisticsTheoreticalUplinkCapacityFlags(depth int, m *models.Statistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	theoreticalUplinkCapacityFlagName := fmt.Sprintf("%v.theoreticalUplinkCapacity", cmdPrefix)
	if cmd.Flags().Changed(theoreticalUplinkCapacityFlagName) {

		var theoreticalUplinkCapacityFlagName string
		if cmdPrefix == "" {
			theoreticalUplinkCapacityFlagName = "theoreticalUplinkCapacity"
		} else {
			theoreticalUplinkCapacityFlagName = fmt.Sprintf("%v.theoreticalUplinkCapacity", cmdPrefix)
		}

		theoreticalUplinkCapacityFlagValue, err := cmd.Flags().GetInt64(theoreticalUplinkCapacityFlagName)
		if err != nil {
			return err, false
		}
		m.TheoreticalUplinkCapacity = theoreticalUplinkCapacityFlagValue

		retAdded = true
	}

	return nil, retAdded
}
