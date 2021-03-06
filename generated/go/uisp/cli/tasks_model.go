// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for Tasks

// register flags to command
func registerModelTasksFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTasksCanceled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTasksFailed(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTasksInProgress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTasksQueued(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTasksSuccessful(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTasksTotal(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTasksCanceled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	canceledDescription := ``

	var canceledFlagName string
	if cmdPrefix == "" {
		canceledFlagName = "canceled"
	} else {
		canceledFlagName = fmt.Sprintf("%v.canceled", cmdPrefix)
	}

	var canceledFlagDefault int64

	_ = cmd.PersistentFlags().Int64(canceledFlagName, canceledFlagDefault, canceledDescription)

	return nil
}

func registerTasksFailed(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	failedDescription := ``

	var failedFlagName string
	if cmdPrefix == "" {
		failedFlagName = "failed"
	} else {
		failedFlagName = fmt.Sprintf("%v.failed", cmdPrefix)
	}

	var failedFlagDefault int64

	_ = cmd.PersistentFlags().Int64(failedFlagName, failedFlagDefault, failedDescription)

	return nil
}

func registerTasksInProgress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	inProgressDescription := ``

	var inProgressFlagName string
	if cmdPrefix == "" {
		inProgressFlagName = "inProgress"
	} else {
		inProgressFlagName = fmt.Sprintf("%v.inProgress", cmdPrefix)
	}

	var inProgressFlagDefault int64

	_ = cmd.PersistentFlags().Int64(inProgressFlagName, inProgressFlagDefault, inProgressDescription)

	return nil
}

func registerTasksQueued(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	queuedDescription := ``

	var queuedFlagName string
	if cmdPrefix == "" {
		queuedFlagName = "queued"
	} else {
		queuedFlagName = fmt.Sprintf("%v.queued", cmdPrefix)
	}

	var queuedFlagDefault int64

	_ = cmd.PersistentFlags().Int64(queuedFlagName, queuedFlagDefault, queuedDescription)

	return nil
}

func registerTasksSuccessful(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	successfulDescription := ``

	var successfulFlagName string
	if cmdPrefix == "" {
		successfulFlagName = "successful"
	} else {
		successfulFlagName = fmt.Sprintf("%v.successful", cmdPrefix)
	}

	var successfulFlagDefault int64

	_ = cmd.PersistentFlags().Int64(successfulFlagName, successfulFlagDefault, successfulDescription)

	return nil
}

func registerTasksTotal(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalDescription := ``

	var totalFlagName string
	if cmdPrefix == "" {
		totalFlagName = "total"
	} else {
		totalFlagName = fmt.Sprintf("%v.total", cmdPrefix)
	}

	var totalFlagDefault int64

	_ = cmd.PersistentFlags().Int64(totalFlagName, totalFlagDefault, totalDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTasksFlags(depth int, m *models.Tasks, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, canceledAdded := retrieveTasksCanceledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || canceledAdded

	err, failedAdded := retrieveTasksFailedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || failedAdded

	err, inProgressAdded := retrieveTasksInProgressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || inProgressAdded

	err, queuedAdded := retrieveTasksQueuedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || queuedAdded

	err, successfulAdded := retrieveTasksSuccessfulFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || successfulAdded

	err, totalAdded := retrieveTasksTotalFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalAdded

	return nil, retAdded
}

func retrieveTasksCanceledFlags(depth int, m *models.Tasks, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	canceledFlagName := fmt.Sprintf("%v.canceled", cmdPrefix)
	if cmd.Flags().Changed(canceledFlagName) {

		var canceledFlagName string
		if cmdPrefix == "" {
			canceledFlagName = "canceled"
		} else {
			canceledFlagName = fmt.Sprintf("%v.canceled", cmdPrefix)
		}

		canceledFlagValue, err := cmd.Flags().GetInt64(canceledFlagName)
		if err != nil {
			return err, false
		}
		m.Canceled = &canceledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTasksFailedFlags(depth int, m *models.Tasks, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	failedFlagName := fmt.Sprintf("%v.failed", cmdPrefix)
	if cmd.Flags().Changed(failedFlagName) {

		var failedFlagName string
		if cmdPrefix == "" {
			failedFlagName = "failed"
		} else {
			failedFlagName = fmt.Sprintf("%v.failed", cmdPrefix)
		}

		failedFlagValue, err := cmd.Flags().GetInt64(failedFlagName)
		if err != nil {
			return err, false
		}
		m.Failed = &failedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTasksInProgressFlags(depth int, m *models.Tasks, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	inProgressFlagName := fmt.Sprintf("%v.inProgress", cmdPrefix)
	if cmd.Flags().Changed(inProgressFlagName) {

		var inProgressFlagName string
		if cmdPrefix == "" {
			inProgressFlagName = "inProgress"
		} else {
			inProgressFlagName = fmt.Sprintf("%v.inProgress", cmdPrefix)
		}

		inProgressFlagValue, err := cmd.Flags().GetInt64(inProgressFlagName)
		if err != nil {
			return err, false
		}
		m.InProgress = &inProgressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTasksQueuedFlags(depth int, m *models.Tasks, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	queuedFlagName := fmt.Sprintf("%v.queued", cmdPrefix)
	if cmd.Flags().Changed(queuedFlagName) {

		var queuedFlagName string
		if cmdPrefix == "" {
			queuedFlagName = "queued"
		} else {
			queuedFlagName = fmt.Sprintf("%v.queued", cmdPrefix)
		}

		queuedFlagValue, err := cmd.Flags().GetInt64(queuedFlagName)
		if err != nil {
			return err, false
		}
		m.Queued = &queuedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTasksSuccessfulFlags(depth int, m *models.Tasks, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	successfulFlagName := fmt.Sprintf("%v.successful", cmdPrefix)
	if cmd.Flags().Changed(successfulFlagName) {

		var successfulFlagName string
		if cmdPrefix == "" {
			successfulFlagName = "successful"
		} else {
			successfulFlagName = fmt.Sprintf("%v.successful", cmdPrefix)
		}

		successfulFlagValue, err := cmd.Flags().GetInt64(successfulFlagName)
		if err != nil {
			return err, false
		}
		m.Successful = &successfulFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTasksTotalFlags(depth int, m *models.Tasks, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalFlagName := fmt.Sprintf("%v.total", cmdPrefix)
	if cmd.Flags().Changed(totalFlagName) {

		var totalFlagName string
		if cmdPrefix == "" {
			totalFlagName = "total"
		} else {
			totalFlagName = fmt.Sprintf("%v.total", cmdPrefix)
		}

		totalFlagValue, err := cmd.Flags().GetInt64(totalFlagName)
		if err != nil {
			return err, false
		}
		m.Total = &totalFlagValue

		retAdded = true
	}

	return nil, retAdded
}
