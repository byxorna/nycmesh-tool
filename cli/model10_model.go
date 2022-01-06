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

// Schema cli for Model10

// register flags to command
func registerModelModel10Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel10Aggregation(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel10Items(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel10Pagination(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel10Aggregation(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var aggregationFlagName string
	if cmdPrefix == "" {
		aggregationFlagName = "aggregation"
	} else {
		aggregationFlagName = fmt.Sprintf("%v.aggregation", cmdPrefix)
	}

	if err := registerModelAggregationFlags(depth+1, aggregationFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel10Items(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: items Items array type is not supported by go-swagger cli yet

	return nil
}

func registerModel10Pagination(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var paginationFlagName string
	if cmdPrefix == "" {
		paginationFlagName = "pagination"
	} else {
		paginationFlagName = fmt.Sprintf("%v.pagination", cmdPrefix)
	}

	if err := registerModelPaginationFlags(depth+1, paginationFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModel10Flags(depth int, m *models.Model10, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, aggregationAdded := retrieveModel10AggregationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || aggregationAdded

	err, itemsAdded := retrieveModel10ItemsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || itemsAdded

	err, paginationAdded := retrieveModel10PaginationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || paginationAdded

	return nil, retAdded
}

func retrieveModel10AggregationFlags(depth int, m *models.Model10, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	aggregationFlagName := fmt.Sprintf("%v.aggregation", cmdPrefix)
	if cmd.Flags().Changed(aggregationFlagName) {
		// info: complex object aggregation Aggregation is retrieved outside this Changed() block
	}
	aggregationFlagValue := m.Aggregation
	if swag.IsZero(aggregationFlagValue) {
		aggregationFlagValue = &models.Aggregation{}
	}

	err, aggregationAdded := retrieveModelAggregationFlags(depth+1, aggregationFlagValue, aggregationFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || aggregationAdded
	if aggregationAdded {
		m.Aggregation = aggregationFlagValue
	}

	return nil, retAdded
}

func retrieveModel10ItemsFlags(depth int, m *models.Model10, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	itemsFlagName := fmt.Sprintf("%v.items", cmdPrefix)
	if cmd.Flags().Changed(itemsFlagName) {
		// warning: items array type Items is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveModel10PaginationFlags(depth int, m *models.Model10, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	paginationFlagName := fmt.Sprintf("%v.pagination", cmdPrefix)
	if cmd.Flags().Changed(paginationFlagName) {
		// info: complex object pagination Pagination is retrieved outside this Changed() block
	}
	paginationFlagValue := m.Pagination
	if swag.IsZero(paginationFlagValue) {
		paginationFlagValue = &models.Pagination{}
	}

	err, paginationAdded := retrieveModelPaginationFlags(depth+1, paginationFlagValue, paginationFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || paginationAdded
	if paginationAdded {
		m.Pagination = paginationFlagValue
	}

	return nil, retAdded
}
