// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

  "github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for Outages

// register flags to command
func registerModelOutagesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerOutagesAggregation(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOutagesItems(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOutagesPagination(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerOutagesAggregation(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var aggregationFlagName string
	if cmdPrefix == "" {
		aggregationFlagName = "aggregation"
	} else {
		aggregationFlagName = fmt.Sprintf("%v.aggregation", cmdPrefix)
	}

	if err := registerModelAggregation1Flags(depth+1, aggregationFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerOutagesItems(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: items OutageList array type is not supported by go-swagger cli yet

	return nil
}

func registerOutagesPagination(depth int, cmdPrefix string, cmd *cobra.Command) error {
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
func retrieveModelOutagesFlags(depth int, m *models.Outages, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, aggregationAdded := retrieveOutagesAggregationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || aggregationAdded

	err, itemsAdded := retrieveOutagesItemsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || itemsAdded

	err, paginationAdded := retrieveOutagesPaginationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || paginationAdded

	return nil, retAdded
}

func retrieveOutagesAggregationFlags(depth int, m *models.Outages, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	aggregationFlagName := fmt.Sprintf("%v.aggregation", cmdPrefix)
	if cmd.Flags().Changed(aggregationFlagName) {
		// info: complex object aggregation Aggregation1 is retrieved outside this Changed() block
	}
	aggregationFlagValue := m.Aggregation
	if swag.IsZero(aggregationFlagValue) {
		aggregationFlagValue = &models.Aggregation1{}
	}

	err, aggregationAdded := retrieveModelAggregation1Flags(depth+1, aggregationFlagValue, aggregationFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || aggregationAdded
	if aggregationAdded {
		m.Aggregation = aggregationFlagValue
	}

	return nil, retAdded
}

func retrieveOutagesItemsFlags(depth int, m *models.Outages, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	itemsFlagName := fmt.Sprintf("%v.items", cmdPrefix)
	if cmd.Flags().Changed(itemsFlagName) {
		// warning: items array type OutageList is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveOutagesPaginationFlags(depth int, m *models.Outages, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
