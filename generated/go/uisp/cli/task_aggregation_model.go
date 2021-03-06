// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for TaskAggregation

// register flags to command
func registerModelTaskAggregationFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTaskAggregationCanceled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskAggregationFailed(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskAggregationInProgress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskAggregationQueued(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskAggregationSuccess(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTaskAggregationCanceled(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerTaskAggregationFailed(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerTaskAggregationInProgress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	inProgressDescription := ``

	var inProgressFlagName string
	if cmdPrefix == "" {
		inProgressFlagName = "in-progress"
	} else {
		inProgressFlagName = fmt.Sprintf("%v.in-progress", cmdPrefix)
	}

	var inProgressFlagDefault int64

	_ = cmd.PersistentFlags().Int64(inProgressFlagName, inProgressFlagDefault, inProgressDescription)

	return nil
}

func registerTaskAggregationQueued(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerTaskAggregationSuccess(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	successDescription := ``

	var successFlagName string
	if cmdPrefix == "" {
		successFlagName = "success"
	} else {
		successFlagName = fmt.Sprintf("%v.success", cmdPrefix)
	}

	var successFlagDefault int64

	_ = cmd.PersistentFlags().Int64(successFlagName, successFlagDefault, successDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTaskAggregationFlags(depth int, m *models.TaskAggregation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, canceledAdded := retrieveTaskAggregationCanceledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || canceledAdded

	err, failedAdded := retrieveTaskAggregationFailedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || failedAdded

	err, inProgressAdded := retrieveTaskAggregationInProgressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || inProgressAdded

	err, queuedAdded := retrieveTaskAggregationQueuedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || queuedAdded

	err, successAdded := retrieveTaskAggregationSuccessFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || successAdded

	return nil, retAdded
}

func retrieveTaskAggregationCanceledFlags(depth int, m *models.TaskAggregation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveTaskAggregationFailedFlags(depth int, m *models.TaskAggregation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveTaskAggregationInProgressFlags(depth int, m *models.TaskAggregation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	inProgressFlagName := fmt.Sprintf("%v.in-progress", cmdPrefix)
	if cmd.Flags().Changed(inProgressFlagName) {

		var inProgressFlagName string
		if cmdPrefix == "" {
			inProgressFlagName = "in-progress"
		} else {
			inProgressFlagName = fmt.Sprintf("%v.in-progress", cmdPrefix)
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

func retrieveTaskAggregationQueuedFlags(depth int, m *models.TaskAggregation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveTaskAggregationSuccessFlags(depth int, m *models.TaskAggregation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	successFlagName := fmt.Sprintf("%v.success", cmdPrefix)
	if cmd.Flags().Changed(successFlagName) {

		var successFlagName string
		if cmdPrefix == "" {
			successFlagName = "success"
		} else {
			successFlagName = fmt.Sprintf("%v.success", cmdPrefix)
		}

		successFlagValue, err := cmd.Flags().GetInt64(successFlagName)
		if err != nil {
			return err, false
		}
		m.Success = &successFlagValue

		retAdded = true
	}

	return nil, retAdded
}
