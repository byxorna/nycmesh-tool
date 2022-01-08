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

// Schema cli for Semver

// register flags to command
func registerModelSemverFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSemverCurrent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSemverLatest(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSemverLatestOnCurrentMajorVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSemverLatestOver(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSemverCurrent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var currentFlagName string
	if cmdPrefix == "" {
		currentFlagName = "current"
	} else {
		currentFlagName = fmt.Sprintf("%v.current", cmdPrefix)
	}

	if err := registerModelCurrentFlags(depth+1, currentFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerSemverLatest(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var latestFlagName string
	if cmdPrefix == "" {
		latestFlagName = "latest"
	} else {
		latestFlagName = fmt.Sprintf("%v.latest", cmdPrefix)
	}

	if err := registerModelLatestFlags(depth+1, latestFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerSemverLatestOnCurrentMajorVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var latestOnCurrentMajorVersionFlagName string
	if cmdPrefix == "" {
		latestOnCurrentMajorVersionFlagName = "latestOnCurrentMajorVersion"
	} else {
		latestOnCurrentMajorVersionFlagName = fmt.Sprintf("%v.latestOnCurrentMajorVersion", cmdPrefix)
	}

	if err := registerModelLatestOnCurrentMajorVersionFlags(depth+1, latestOnCurrentMajorVersionFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerSemverLatestOver(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var latestOverFlagName string
	if cmdPrefix == "" {
		latestOverFlagName = "latestOver"
	} else {
		latestOverFlagName = fmt.Sprintf("%v.latestOver", cmdPrefix)
	}

	if err := registerModelLatestOverFlags(depth+1, latestOverFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSemverFlags(depth int, m *models.Semver, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, currentAdded := retrieveSemverCurrentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || currentAdded

	err, latestAdded := retrieveSemverLatestFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || latestAdded

	err, latestOnCurrentMajorVersionAdded := retrieveSemverLatestOnCurrentMajorVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || latestOnCurrentMajorVersionAdded

	err, latestOverAdded := retrieveSemverLatestOverFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || latestOverAdded

	return nil, retAdded
}

func retrieveSemverCurrentFlags(depth int, m *models.Semver, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	currentFlagName := fmt.Sprintf("%v.current", cmdPrefix)
	if cmd.Flags().Changed(currentFlagName) {
		// info: complex object current Current is retrieved outside this Changed() block
	}
	currentFlagValue := m.Current
	if swag.IsZero(currentFlagValue) {
		currentFlagValue = &models.Current{}
	}

	err, currentAdded := retrieveModelCurrentFlags(depth+1, currentFlagValue, currentFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || currentAdded
	if currentAdded {
		m.Current = currentFlagValue
	}

	return nil, retAdded
}

func retrieveSemverLatestFlags(depth int, m *models.Semver, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	latestFlagName := fmt.Sprintf("%v.latest", cmdPrefix)
	if cmd.Flags().Changed(latestFlagName) {
		// info: complex object latest Latest is retrieved outside this Changed() block
	}
	latestFlagValue := m.Latest
	if swag.IsZero(latestFlagValue) {
		latestFlagValue = &models.Latest{}
	}

	err, latestAdded := retrieveModelLatestFlags(depth+1, latestFlagValue, latestFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || latestAdded
	if latestAdded {
		m.Latest = latestFlagValue
	}

	return nil, retAdded
}

func retrieveSemverLatestOnCurrentMajorVersionFlags(depth int, m *models.Semver, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	latestOnCurrentMajorVersionFlagName := fmt.Sprintf("%v.latestOnCurrentMajorVersion", cmdPrefix)
	if cmd.Flags().Changed(latestOnCurrentMajorVersionFlagName) {
		// info: complex object latestOnCurrentMajorVersion LatestOnCurrentMajorVersion is retrieved outside this Changed() block
	}
	latestOnCurrentMajorVersionFlagValue := m.LatestOnCurrentMajorVersion
	if swag.IsZero(latestOnCurrentMajorVersionFlagValue) {
		latestOnCurrentMajorVersionFlagValue = &models.LatestOnCurrentMajorVersion{}
	}

	err, latestOnCurrentMajorVersionAdded := retrieveModelLatestOnCurrentMajorVersionFlags(depth+1, latestOnCurrentMajorVersionFlagValue, latestOnCurrentMajorVersionFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || latestOnCurrentMajorVersionAdded
	if latestOnCurrentMajorVersionAdded {
		m.LatestOnCurrentMajorVersion = latestOnCurrentMajorVersionFlagValue
	}

	return nil, retAdded
}

func retrieveSemverLatestOverFlags(depth int, m *models.Semver, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	latestOverFlagName := fmt.Sprintf("%v.latestOver", cmdPrefix)
	if cmd.Flags().Changed(latestOverFlagName) {
		// info: complex object latestOver LatestOver is retrieved outside this Changed() block
	}
	latestOverFlagValue := m.LatestOver
	if swag.IsZero(latestOverFlagValue) {
		latestOverFlagValue = &models.LatestOver{}
	}

	err, latestOverAdded := retrieveModelLatestOverFlags(depth+1, latestOverFlagValue, latestOverFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || latestOverAdded
	if latestOverAdded {
		m.LatestOver = latestOverFlagValue
	}

	return nil, retAdded
}
