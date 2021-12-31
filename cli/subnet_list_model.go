// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for SubnetList

// register flags to command
func registerModelSubnetListFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSubnetListInvalidSubnets(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSubnetListValidSubnets(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSubnetListInvalidSubnets(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: invalidSubnets InvalidSubnets array type is not supported by go-swagger cli yet

	return nil
}

func registerSubnetListValidSubnets(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: validSubnets ValidSubnets array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSubnetListFlags(depth int, m *models.SubnetList, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, invalidSubnetsAdded := retrieveSubnetListInvalidSubnetsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || invalidSubnetsAdded

	err, validSubnetsAdded := retrieveSubnetListValidSubnetsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || validSubnetsAdded

	return nil, retAdded
}

func retrieveSubnetListInvalidSubnetsFlags(depth int, m *models.SubnetList, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	invalidSubnetsFlagName := fmt.Sprintf("%v.invalidSubnets", cmdPrefix)
	if cmd.Flags().Changed(invalidSubnetsFlagName) {
		// warning: invalidSubnets array type InvalidSubnets is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveSubnetListValidSubnetsFlags(depth int, m *models.SubnetList, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	validSubnetsFlagName := fmt.Sprintf("%v.validSubnets", cmdPrefix)
	if cmd.Flags().Changed(validSubnetsFlagName) {
		// warning: validSubnets array type ValidSubnets is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
