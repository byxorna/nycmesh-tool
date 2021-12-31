// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for CrackTimesSeconds

// register flags to command
func registerModelCrackTimesSecondsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCrackTimesSecondsOfflineFastHashing1e10PerSecond(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCrackTimesSecondsOfflineSlowHashing1e4PerSecond(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCrackTimesSecondsOnlineNoThrottling10PerSecond(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCrackTimesSecondsOnlineThrottling100PerHour(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCrackTimesSecondsOfflineFastHashing1e10PerSecond(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	offlineFastHashing1e10PerSecondDescription := `Required. `

	var offlineFastHashing1e10PerSecondFlagName string
	if cmdPrefix == "" {
		offlineFastHashing1e10PerSecondFlagName = "offline_fast_hashing_1e10_per_second"
	} else {
		offlineFastHashing1e10PerSecondFlagName = fmt.Sprintf("%v.offline_fast_hashing_1e10_per_second", cmdPrefix)
	}

	var offlineFastHashing1e10PerSecondFlagDefault float64

	_ = cmd.PersistentFlags().Float64(offlineFastHashing1e10PerSecondFlagName, offlineFastHashing1e10PerSecondFlagDefault, offlineFastHashing1e10PerSecondDescription)

	return nil
}

func registerCrackTimesSecondsOfflineSlowHashing1e4PerSecond(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	offlineSlowHashing1e4PerSecondDescription := `Required. `

	var offlineSlowHashing1e4PerSecondFlagName string
	if cmdPrefix == "" {
		offlineSlowHashing1e4PerSecondFlagName = "offline_slow_hashing_1e4_per_second"
	} else {
		offlineSlowHashing1e4PerSecondFlagName = fmt.Sprintf("%v.offline_slow_hashing_1e4_per_second", cmdPrefix)
	}

	var offlineSlowHashing1e4PerSecondFlagDefault float64

	_ = cmd.PersistentFlags().Float64(offlineSlowHashing1e4PerSecondFlagName, offlineSlowHashing1e4PerSecondFlagDefault, offlineSlowHashing1e4PerSecondDescription)

	return nil
}

func registerCrackTimesSecondsOnlineNoThrottling10PerSecond(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	onlineNoThrottling10PerSecondDescription := `Required. `

	var onlineNoThrottling10PerSecondFlagName string
	if cmdPrefix == "" {
		onlineNoThrottling10PerSecondFlagName = "online_no_throttling_10_per_second"
	} else {
		onlineNoThrottling10PerSecondFlagName = fmt.Sprintf("%v.online_no_throttling_10_per_second", cmdPrefix)
	}

	var onlineNoThrottling10PerSecondFlagDefault float64

	_ = cmd.PersistentFlags().Float64(onlineNoThrottling10PerSecondFlagName, onlineNoThrottling10PerSecondFlagDefault, onlineNoThrottling10PerSecondDescription)

	return nil
}

func registerCrackTimesSecondsOnlineThrottling100PerHour(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	onlineThrottling100PerHourDescription := `Required. `

	var onlineThrottling100PerHourFlagName string
	if cmdPrefix == "" {
		onlineThrottling100PerHourFlagName = "online_throttling_100_per_hour"
	} else {
		onlineThrottling100PerHourFlagName = fmt.Sprintf("%v.online_throttling_100_per_hour", cmdPrefix)
	}

	var onlineThrottling100PerHourFlagDefault float64

	_ = cmd.PersistentFlags().Float64(onlineThrottling100PerHourFlagName, onlineThrottling100PerHourFlagDefault, onlineThrottling100PerHourDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCrackTimesSecondsFlags(depth int, m *models.CrackTimesSeconds, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, offlineFastHashing1e10PerSecondAdded := retrieveCrackTimesSecondsOfflineFastHashing1e10PerSecondFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || offlineFastHashing1e10PerSecondAdded

	err, offlineSlowHashing1e4PerSecondAdded := retrieveCrackTimesSecondsOfflineSlowHashing1e4PerSecondFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || offlineSlowHashing1e4PerSecondAdded

	err, onlineNoThrottling10PerSecondAdded := retrieveCrackTimesSecondsOnlineNoThrottling10PerSecondFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || onlineNoThrottling10PerSecondAdded

	err, onlineThrottling100PerHourAdded := retrieveCrackTimesSecondsOnlineThrottling100PerHourFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || onlineThrottling100PerHourAdded

	return nil, retAdded
}

func retrieveCrackTimesSecondsOfflineFastHashing1e10PerSecondFlags(depth int, m *models.CrackTimesSeconds, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	offlineFastHashing1e10PerSecondFlagName := fmt.Sprintf("%v.offline_fast_hashing_1e10_per_second", cmdPrefix)
	if cmd.Flags().Changed(offlineFastHashing1e10PerSecondFlagName) {

		var offlineFastHashing1e10PerSecondFlagName string
		if cmdPrefix == "" {
			offlineFastHashing1e10PerSecondFlagName = "offline_fast_hashing_1e10_per_second"
		} else {
			offlineFastHashing1e10PerSecondFlagName = fmt.Sprintf("%v.offline_fast_hashing_1e10_per_second", cmdPrefix)
		}

		offlineFastHashing1e10PerSecondFlagValue, err := cmd.Flags().GetFloat64(offlineFastHashing1e10PerSecondFlagName)
		if err != nil {
			return err, false
		}
		m.OfflineFastHashing1e10PerSecond = &offlineFastHashing1e10PerSecondFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCrackTimesSecondsOfflineSlowHashing1e4PerSecondFlags(depth int, m *models.CrackTimesSeconds, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	offlineSlowHashing1e4PerSecondFlagName := fmt.Sprintf("%v.offline_slow_hashing_1e4_per_second", cmdPrefix)
	if cmd.Flags().Changed(offlineSlowHashing1e4PerSecondFlagName) {

		var offlineSlowHashing1e4PerSecondFlagName string
		if cmdPrefix == "" {
			offlineSlowHashing1e4PerSecondFlagName = "offline_slow_hashing_1e4_per_second"
		} else {
			offlineSlowHashing1e4PerSecondFlagName = fmt.Sprintf("%v.offline_slow_hashing_1e4_per_second", cmdPrefix)
		}

		offlineSlowHashing1e4PerSecondFlagValue, err := cmd.Flags().GetFloat64(offlineSlowHashing1e4PerSecondFlagName)
		if err != nil {
			return err, false
		}
		m.OfflineSlowHashing1e4PerSecond = &offlineSlowHashing1e4PerSecondFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCrackTimesSecondsOnlineNoThrottling10PerSecondFlags(depth int, m *models.CrackTimesSeconds, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	onlineNoThrottling10PerSecondFlagName := fmt.Sprintf("%v.online_no_throttling_10_per_second", cmdPrefix)
	if cmd.Flags().Changed(onlineNoThrottling10PerSecondFlagName) {

		var onlineNoThrottling10PerSecondFlagName string
		if cmdPrefix == "" {
			onlineNoThrottling10PerSecondFlagName = "online_no_throttling_10_per_second"
		} else {
			onlineNoThrottling10PerSecondFlagName = fmt.Sprintf("%v.online_no_throttling_10_per_second", cmdPrefix)
		}

		onlineNoThrottling10PerSecondFlagValue, err := cmd.Flags().GetFloat64(onlineNoThrottling10PerSecondFlagName)
		if err != nil {
			return err, false
		}
		m.OnlineNoThrottling10PerSecond = &onlineNoThrottling10PerSecondFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCrackTimesSecondsOnlineThrottling100PerHourFlags(depth int, m *models.CrackTimesSeconds, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	onlineThrottling100PerHourFlagName := fmt.Sprintf("%v.online_throttling_100_per_hour", cmdPrefix)
	if cmd.Flags().Changed(onlineThrottling100PerHourFlagName) {

		var onlineThrottling100PerHourFlagName string
		if cmdPrefix == "" {
			onlineThrottling100PerHourFlagName = "online_throttling_100_per_hour"
		} else {
			onlineThrottling100PerHourFlagName = fmt.Sprintf("%v.online_throttling_100_per_hour", cmdPrefix)
		}

		onlineThrottling100PerHourFlagValue, err := cmd.Flags().GetFloat64(onlineThrottling100PerHourFlagName)
		if err != nil {
			return err, false
		}
		m.OnlineThrottling100PerHour = &onlineThrottling100PerHourFlagValue

		retAdded = true
	}

	return nil, retAdded
}
