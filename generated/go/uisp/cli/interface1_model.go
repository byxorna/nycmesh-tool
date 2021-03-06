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

// Schema cli for Interface1

// register flags to command
func registerModelInterface1Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerInterface1Sfp1(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInterface1Sfp2(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInterface1SfpPlus1(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInterface1SfpPlus2(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerInterface1Sfp1(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var sfp1FlagName string
	if cmdPrefix == "" {
		sfp1FlagName = "sfp1"
	} else {
		sfp1FlagName = fmt.Sprintf("%v.sfp1", cmdPrefix)
	}

	if err := registerModelSfp1Flags(depth+1, sfp1FlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerInterface1Sfp2(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var sfp2FlagName string
	if cmdPrefix == "" {
		sfp2FlagName = "sfp2"
	} else {
		sfp2FlagName = fmt.Sprintf("%v.sfp2", cmdPrefix)
	}

	if err := registerModelSfp2Flags(depth+1, sfp2FlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerInterface1SfpPlus1(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var sfpPlus1FlagName string
	if cmdPrefix == "" {
		sfpPlus1FlagName = "sfpPlus1"
	} else {
		sfpPlus1FlagName = fmt.Sprintf("%v.sfpPlus1", cmdPrefix)
	}

	if err := registerModelSfpPlus1Flags(depth+1, sfpPlus1FlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerInterface1SfpPlus2(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var sfpPlus2FlagName string
	if cmdPrefix == "" {
		sfpPlus2FlagName = "sfpPlus2"
	} else {
		sfpPlus2FlagName = fmt.Sprintf("%v.sfpPlus2", cmdPrefix)
	}

	if err := registerModelSfpPlus2Flags(depth+1, sfpPlus2FlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelInterface1Flags(depth int, m *models.Interface1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, sfp1Added := retrieveInterface1Sfp1Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sfp1Added

	err, sfp2Added := retrieveInterface1Sfp2Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sfp2Added

	err, sfpPlus1Added := retrieveInterface1SfpPlus1Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sfpPlus1Added

	err, sfpPlus2Added := retrieveInterface1SfpPlus2Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sfpPlus2Added

	return nil, retAdded
}

func retrieveInterface1Sfp1Flags(depth int, m *models.Interface1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sfp1FlagName := fmt.Sprintf("%v.sfp1", cmdPrefix)
	if cmd.Flags().Changed(sfp1FlagName) {
		// info: complex object sfp1 Sfp1 is retrieved outside this Changed() block
	}
	sfp1FlagValue := m.Sfp1
	if swag.IsZero(sfp1FlagValue) {
		sfp1FlagValue = &models.Sfp1{}
	}

	err, sfp1Added := retrieveModelSfp1Flags(depth+1, sfp1FlagValue, sfp1FlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sfp1Added
	if sfp1Added {
		m.Sfp1 = sfp1FlagValue
	}

	return nil, retAdded
}

func retrieveInterface1Sfp2Flags(depth int, m *models.Interface1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sfp2FlagName := fmt.Sprintf("%v.sfp2", cmdPrefix)
	if cmd.Flags().Changed(sfp2FlagName) {
		// info: complex object sfp2 Sfp2 is retrieved outside this Changed() block
	}
	sfp2FlagValue := m.Sfp2
	if swag.IsZero(sfp2FlagValue) {
		sfp2FlagValue = &models.Sfp2{}
	}

	err, sfp2Added := retrieveModelSfp2Flags(depth+1, sfp2FlagValue, sfp2FlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sfp2Added
	if sfp2Added {
		m.Sfp2 = sfp2FlagValue
	}

	return nil, retAdded
}

func retrieveInterface1SfpPlus1Flags(depth int, m *models.Interface1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sfpPlus1FlagName := fmt.Sprintf("%v.sfpPlus1", cmdPrefix)
	if cmd.Flags().Changed(sfpPlus1FlagName) {
		// info: complex object sfpPlus1 SfpPlus1 is retrieved outside this Changed() block
	}
	sfpPlus1FlagValue := m.SfpPlus1
	if swag.IsZero(sfpPlus1FlagValue) {
		sfpPlus1FlagValue = &models.SfpPlus1{}
	}

	err, sfpPlus1Added := retrieveModelSfpPlus1Flags(depth+1, sfpPlus1FlagValue, sfpPlus1FlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sfpPlus1Added
	if sfpPlus1Added {
		m.SfpPlus1 = sfpPlus1FlagValue
	}

	return nil, retAdded
}

func retrieveInterface1SfpPlus2Flags(depth int, m *models.Interface1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sfpPlus2FlagName := fmt.Sprintf("%v.sfpPlus2", cmdPrefix)
	if cmd.Flags().Changed(sfpPlus2FlagName) {
		// info: complex object sfpPlus2 SfpPlus2 is retrieved outside this Changed() block
	}
	sfpPlus2FlagValue := m.SfpPlus2
	if swag.IsZero(sfpPlus2FlagValue) {
		sfpPlus2FlagValue = &models.SfpPlus2{}
	}

	err, sfpPlus2Added := retrieveModelSfpPlus2Flags(depth+1, sfpPlus2FlagValue, sfpPlus2FlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sfpPlus2Added
	if sfpPlus2Added {
		m.SfpPlus2 = sfpPlus2FlagValue
	}

	return nil, retAdded
}
