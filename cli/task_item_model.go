// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for TaskItem

// register flags to command
func registerModelTaskItemFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTaskItemEndTimestamp(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskItemIdentification(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskItemProgress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskItemStartTimestamp(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskItemStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskItemTasks(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTaskItemEndTimestamp(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	endTimestampDescription := ``

	var endTimestampFlagName string
	if cmdPrefix == "" {
		endTimestampFlagName = "endTimestamp"
	} else {
		endTimestampFlagName = fmt.Sprintf("%v.endTimestamp", cmdPrefix)
	}

	var endTimestampFlagDefault int64

	_ = cmd.PersistentFlags().Int64(endTimestampFlagName, endTimestampFlagDefault, endTimestampDescription)

	return nil
}

func registerTaskItemIdentification(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var identificationFlagName string
	if cmdPrefix == "" {
		identificationFlagName = "identification"
	} else {
		identificationFlagName = fmt.Sprintf("%v.identification", cmdPrefix)
	}

	if err := registerModelIdentification2Flags(depth+1, identificationFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerTaskItemProgress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	progressDescription := ``

	var progressFlagName string
	if cmdPrefix == "" {
		progressFlagName = "progress"
	} else {
		progressFlagName = fmt.Sprintf("%v.progress", cmdPrefix)
	}

	var progressFlagDefault float64

	_ = cmd.PersistentFlags().Float64(progressFlagName, progressFlagDefault, progressDescription)

	return nil
}

func registerTaskItemStartTimestamp(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	startTimestampDescription := ``

	var startTimestampFlagName string
	if cmdPrefix == "" {
		startTimestampFlagName = "startTimestamp"
	} else {
		startTimestampFlagName = fmt.Sprintf("%v.startTimestamp", cmdPrefix)
	}

	var startTimestampFlagDefault int64

	_ = cmd.PersistentFlags().Int64(startTimestampFlagName, startTimestampFlagDefault, startTimestampDescription)

	return nil
}

func registerTaskItemStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statusDescription := ``

	var statusFlagName string
	if cmdPrefix == "" {
		statusFlagName = "status"
	} else {
		statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
	}

	var statusFlagDefault string

	_ = cmd.PersistentFlags().String(statusFlagName, statusFlagDefault, statusDescription)

	return nil
}

func registerTaskItemTasks(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var tasksFlagName string
	if cmdPrefix == "" {
		tasksFlagName = "tasks"
	} else {
		tasksFlagName = fmt.Sprintf("%v.tasks", cmdPrefix)
	}

	if err := registerModelTasksFlags(depth+1, tasksFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTaskItemFlags(depth int, m *models.TaskItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, endTimestampAdded := retrieveTaskItemEndTimestampFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || endTimestampAdded

	err, identificationAdded := retrieveTaskItemIdentificationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || identificationAdded

	err, progressAdded := retrieveTaskItemProgressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || progressAdded

	err, startTimestampAdded := retrieveTaskItemStartTimestampFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || startTimestampAdded

	err, statusAdded := retrieveTaskItemStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusAdded

	err, tasksAdded := retrieveTaskItemTasksFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tasksAdded

	return nil, retAdded
}

func retrieveTaskItemEndTimestampFlags(depth int, m *models.TaskItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	endTimestampFlagName := fmt.Sprintf("%v.endTimestamp", cmdPrefix)
	if cmd.Flags().Changed(endTimestampFlagName) {

		var endTimestampFlagName string
		if cmdPrefix == "" {
			endTimestampFlagName = "endTimestamp"
		} else {
			endTimestampFlagName = fmt.Sprintf("%v.endTimestamp", cmdPrefix)
		}

		endTimestampFlagValue, err := cmd.Flags().GetInt64(endTimestampFlagName)
		if err != nil {
			return err, false
		}
		m.EndTimestamp = &endTimestampFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTaskItemIdentificationFlags(depth int, m *models.TaskItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	identificationFlagName := fmt.Sprintf("%v.identification", cmdPrefix)
	if cmd.Flags().Changed(identificationFlagName) {
		// info: complex object identification Identification2 is retrieved outside this Changed() block
	}
	identificationFlagValue := m.Identification
	if swag.IsZero(identificationFlagValue) {
		identificationFlagValue = &models.Identification2{}
	}

	err, identificationAdded := retrieveModelIdentification2Flags(depth+1, identificationFlagValue, identificationFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || identificationAdded
	if identificationAdded {
		m.Identification = identificationFlagValue
	}

	return nil, retAdded
}

func retrieveTaskItemProgressFlags(depth int, m *models.TaskItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	progressFlagName := fmt.Sprintf("%v.progress", cmdPrefix)
	if cmd.Flags().Changed(progressFlagName) {

		var progressFlagName string
		if cmdPrefix == "" {
			progressFlagName = "progress"
		} else {
			progressFlagName = fmt.Sprintf("%v.progress", cmdPrefix)
		}

		progressFlagValue, err := cmd.Flags().GetFloat64(progressFlagName)
		if err != nil {
			return err, false
		}
		m.Progress = &progressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTaskItemStartTimestampFlags(depth int, m *models.TaskItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	startTimestampFlagName := fmt.Sprintf("%v.startTimestamp", cmdPrefix)
	if cmd.Flags().Changed(startTimestampFlagName) {

		var startTimestampFlagName string
		if cmdPrefix == "" {
			startTimestampFlagName = "startTimestamp"
		} else {
			startTimestampFlagName = fmt.Sprintf("%v.startTimestamp", cmdPrefix)
		}

		startTimestampFlagValue, err := cmd.Flags().GetInt64(startTimestampFlagName)
		if err != nil {
			return err, false
		}
		m.StartTimestamp = &startTimestampFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTaskItemStatusFlags(depth int, m *models.TaskItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statusFlagName := fmt.Sprintf("%v.status", cmdPrefix)
	if cmd.Flags().Changed(statusFlagName) {

		var statusFlagName string
		if cmdPrefix == "" {
			statusFlagName = "status"
		} else {
			statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
		}

		statusFlagValue, err := cmd.Flags().GetString(statusFlagName)
		if err != nil {
			return err, false
		}
		m.Status = statusFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTaskItemTasksFlags(depth int, m *models.TaskItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	tasksFlagName := fmt.Sprintf("%v.tasks", cmdPrefix)
	if cmd.Flags().Changed(tasksFlagName) {
		// info: complex object tasks Tasks is retrieved outside this Changed() block
	}
	tasksFlagValue := m.Tasks
	if swag.IsZero(tasksFlagValue) {
		tasksFlagValue = &models.Tasks{}
	}

	err, tasksAdded := retrieveModelTasksFlags(depth+1, tasksFlagValue, tasksFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tasksAdded
	if tasksAdded {
		m.Tasks = tasksFlagValue
	}

	return nil, retAdded
}
