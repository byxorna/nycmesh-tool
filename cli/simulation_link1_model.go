// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for SimulationLink1

// register flags to command
func registerModelSimulationLink1Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSimulationLink1ApDeviceID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSimulationLink1CpeDeviceID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSimulationLink1CreatedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSimulationLink1ID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSimulationLink1Quality(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSimulationLink1Terrain(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSimulationLink1UpdatedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSimulationLink1ApDeviceID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	apDeviceIdDescription := `Required. `

	var apDeviceIdFlagName string
	if cmdPrefix == "" {
		apDeviceIdFlagName = "apDeviceId"
	} else {
		apDeviceIdFlagName = fmt.Sprintf("%v.apDeviceId", cmdPrefix)
	}

	var apDeviceIdFlagDefault string

	_ = cmd.PersistentFlags().String(apDeviceIdFlagName, apDeviceIdFlagDefault, apDeviceIdDescription)

	return nil
}

func registerSimulationLink1CpeDeviceID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	cpeDeviceIdDescription := `Required. `

	var cpeDeviceIdFlagName string
	if cmdPrefix == "" {
		cpeDeviceIdFlagName = "cpeDeviceId"
	} else {
		cpeDeviceIdFlagName = fmt.Sprintf("%v.cpeDeviceId", cmdPrefix)
	}

	var cpeDeviceIdFlagDefault string

	_ = cmd.PersistentFlags().String(cpeDeviceIdFlagName, cpeDeviceIdFlagDefault, cpeDeviceIdDescription)

	return nil
}

func registerSimulationLink1CreatedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: primitive createdAt strfmt.Date is not supported by go-swagger cli yet

	return nil
}

func registerSimulationLink1ID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `Required. `

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerSimulationLink1Quality(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	qualityDescription := `Required. `

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

func registerSimulationLink1Terrain(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerSimulationLink1UpdatedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: primitive updatedAt strfmt.Date is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSimulationLink1Flags(depth int, m *models.SimulationLink1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, apDeviceIdAdded := retrieveSimulationLink1ApDeviceIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || apDeviceIdAdded

	err, cpeDeviceIdAdded := retrieveSimulationLink1CpeDeviceIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || cpeDeviceIdAdded

	err, createdAtAdded := retrieveSimulationLink1CreatedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || createdAtAdded

	err, idAdded := retrieveSimulationLink1IDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, qualityAdded := retrieveSimulationLink1QualityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || qualityAdded

	err, terrainAdded := retrieveSimulationLink1TerrainFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || terrainAdded

	err, updatedAtAdded := retrieveSimulationLink1UpdatedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || updatedAtAdded

	return nil, retAdded
}

func retrieveSimulationLink1ApDeviceIDFlags(depth int, m *models.SimulationLink1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	apDeviceIdFlagName := fmt.Sprintf("%v.apDeviceId", cmdPrefix)
	if cmd.Flags().Changed(apDeviceIdFlagName) {

		var apDeviceIdFlagName string
		if cmdPrefix == "" {
			apDeviceIdFlagName = "apDeviceId"
		} else {
			apDeviceIdFlagName = fmt.Sprintf("%v.apDeviceId", cmdPrefix)
		}

		apDeviceIdFlagValue, err := cmd.Flags().GetString(apDeviceIdFlagName)
		if err != nil {
			return err, false
		}
		m.ApDeviceID = &apDeviceIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSimulationLink1CpeDeviceIDFlags(depth int, m *models.SimulationLink1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	cpeDeviceIdFlagName := fmt.Sprintf("%v.cpeDeviceId", cmdPrefix)
	if cmd.Flags().Changed(cpeDeviceIdFlagName) {

		var cpeDeviceIdFlagName string
		if cmdPrefix == "" {
			cpeDeviceIdFlagName = "cpeDeviceId"
		} else {
			cpeDeviceIdFlagName = fmt.Sprintf("%v.cpeDeviceId", cmdPrefix)
		}

		cpeDeviceIdFlagValue, err := cmd.Flags().GetString(cpeDeviceIdFlagName)
		if err != nil {
			return err, false
		}
		m.CpeDeviceID = &cpeDeviceIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSimulationLink1CreatedAtFlags(depth int, m *models.SimulationLink1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	createdAtFlagName := fmt.Sprintf("%v.createdAt", cmdPrefix)
	if cmd.Flags().Changed(createdAtFlagName) {

		// warning: primitive createdAt strfmt.Date is not supported by go-swagger cli yet

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSimulationLink1IDFlags(depth int, m *models.SimulationLink1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = &idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSimulationLink1QualityFlags(depth int, m *models.SimulationLink1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Quality = &qualityFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSimulationLink1TerrainFlags(depth int, m *models.SimulationLink1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveSimulationLink1UpdatedAtFlags(depth int, m *models.SimulationLink1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	updatedAtFlagName := fmt.Sprintf("%v.updatedAt", cmdPrefix)
	if cmd.Flags().Changed(updatedAtFlagName) {

		// warning: primitive updatedAt strfmt.Date is not supported by go-swagger cli yet

		retAdded = true
	}

	return nil, retAdded
}
