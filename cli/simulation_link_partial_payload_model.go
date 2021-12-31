// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for SimulationLinkPartialPayload

// register flags to command
func registerModelSimulationLinkPartialPayloadFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSimulationLinkPartialPayloadQuality(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSimulationLinkPartialPayloadTerrain(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSimulationLinkPartialPayloadQuality(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	qualityDescription := ``

	var qualityFlagName string
	if cmdPrefix == "" {
		qualityFlagName = "quality"
	} else {
		qualityFlagName = fmt.Sprintf("%v.quality", cmdPrefix)
	}

	var qualityFlagDefault string

	_ = cmd.PersistentFlags().String(qualityFlagName, qualityFlagDefault, qualityDescription)

	return nil
}

func registerSimulationLinkPartialPayloadTerrain(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var terrainFlagName string
	if cmdPrefix == "" {
		terrainFlagName = "terrain"
	} else {
		terrainFlagName = fmt.Sprintf("%v.terrain", cmdPrefix)
	}

	if err := registerModelTerrainFlags(depth+1, terrainFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSimulationLinkPartialPayloadFlags(depth int, m *models.SimulationLinkPartialPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, qualityAdded := retrieveSimulationLinkPartialPayloadQualityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || qualityAdded

	err, terrainAdded := retrieveSimulationLinkPartialPayloadTerrainFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || terrainAdded

	return nil, retAdded
}

func retrieveSimulationLinkPartialPayloadQualityFlags(depth int, m *models.SimulationLinkPartialPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	qualityFlagName := fmt.Sprintf("%v.quality", cmdPrefix)
	if cmd.Flags().Changed(qualityFlagName) {

		var qualityFlagName string
		if cmdPrefix == "" {
			qualityFlagName = "quality"
		} else {
			qualityFlagName = fmt.Sprintf("%v.quality", cmdPrefix)
		}

		qualityFlagValue, err := cmd.Flags().GetString(qualityFlagName)
		if err != nil {
			return err, false
		}
		m.Quality = qualityFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSimulationLinkPartialPayloadTerrainFlags(depth int, m *models.SimulationLinkPartialPayload, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	terrainFlagName := fmt.Sprintf("%v.terrain", cmdPrefix)
	if cmd.Flags().Changed(terrainFlagName) {
		// info: complex object terrain Terrain is retrieved outside this Changed() block
	}
	terrainFlagValue := m.Terrain
	if swag.IsZero(terrainFlagValue) {
		terrainFlagValue = &models.Terrain{}
	}

	err, terrainAdded := retrieveModelTerrainFlags(depth+1, terrainFlagValue, terrainFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || terrainAdded
	if terrainAdded {
		m.Terrain = terrainFlagValue
	}

	return nil, retAdded
}
