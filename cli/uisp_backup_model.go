// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
  "github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for UispBackup

// register flags to command
func registerModelUispBackupFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUispBackupCompatible(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUispBackupCreatedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUispBackupID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUispBackupOrigin(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUispBackupSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUispBackupState(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUispBackupUnmsVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUispBackupCompatible(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	compatibleDescription := `Required. Flag if the backup is compatible with current UISP version.`

	var compatibleFlagName string
	if cmdPrefix == "" {
		compatibleFlagName = "compatible"
	} else {
		compatibleFlagName = fmt.Sprintf("%v.compatible", cmdPrefix)
	}

	var compatibleFlagDefault bool

	_ = cmd.PersistentFlags().Bool(compatibleFlagName, compatibleFlagDefault, compatibleDescription)

	return nil
}

func registerUispBackupCreatedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: primitive createdAt strfmt.Date is not supported by go-swagger cli yet

	return nil
}

func registerUispBackupID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerUispBackupOrigin(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	originDescription := `Enum: ["auto","manual"]. Required. `

	var originFlagName string
	if cmdPrefix == "" {
		originFlagName = "origin"
	} else {
		originFlagName = fmt.Sprintf("%v.origin", cmdPrefix)
	}

	var originFlagDefault string

	_ = cmd.PersistentFlags().String(originFlagName, originFlagDefault, originDescription)

	if err := cmd.RegisterFlagCompletionFunc(originFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["auto","manual"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerUispBackupSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sizeDescription := `Required. Size in bytes.`

	var sizeFlagName string
	if cmdPrefix == "" {
		sizeFlagName = "size"
	} else {
		sizeFlagName = fmt.Sprintf("%v.size", cmdPrefix)
	}

	var sizeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(sizeFlagName, sizeFlagDefault, sizeDescription)

	return nil
}

func registerUispBackupState(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	stateDescription := `Enum: ["success","in-progress","failed"]. Required. `

	var stateFlagName string
	if cmdPrefix == "" {
		stateFlagName = "state"
	} else {
		stateFlagName = fmt.Sprintf("%v.state", cmdPrefix)
	}

	var stateFlagDefault string

	_ = cmd.PersistentFlags().String(stateFlagName, stateFlagDefault, stateDescription)

	if err := cmd.RegisterFlagCompletionFunc(stateFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["success","in-progress","failed"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerUispBackupUnmsVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var unmsVersionFlagName string
	if cmdPrefix == "" {
		unmsVersionFlagName = "unmsVersion"
	} else {
		unmsVersionFlagName = fmt.Sprintf("%v.unmsVersion", cmdPrefix)
	}

	if err := registerModelUnmsVersionFlags(depth+1, unmsVersionFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUispBackupFlags(depth int, m *models.UispBackup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, compatibleAdded := retrieveUispBackupCompatibleFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || compatibleAdded

	err, createdAtAdded := retrieveUispBackupCreatedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || createdAtAdded

	err, idAdded := retrieveUispBackupIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, originAdded := retrieveUispBackupOriginFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || originAdded

	err, sizeAdded := retrieveUispBackupSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sizeAdded

	err, stateAdded := retrieveUispBackupStateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || stateAdded

	err, unmsVersionAdded := retrieveUispBackupUnmsVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || unmsVersionAdded

	return nil, retAdded
}

func retrieveUispBackupCompatibleFlags(depth int, m *models.UispBackup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	compatibleFlagName := fmt.Sprintf("%v.compatible", cmdPrefix)
	if cmd.Flags().Changed(compatibleFlagName) {

		var compatibleFlagName string
		if cmdPrefix == "" {
			compatibleFlagName = "compatible"
		} else {
			compatibleFlagName = fmt.Sprintf("%v.compatible", cmdPrefix)
		}

		compatibleFlagValue, err := cmd.Flags().GetBool(compatibleFlagName)
		if err != nil {
			return err, false
		}
		m.Compatible = &compatibleFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUispBackupCreatedAtFlags(depth int, m *models.UispBackup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveUispBackupIDFlags(depth int, m *models.UispBackup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveUispBackupOriginFlags(depth int, m *models.UispBackup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	originFlagName := fmt.Sprintf("%v.origin", cmdPrefix)
	if cmd.Flags().Changed(originFlagName) {

		var originFlagName string
		if cmdPrefix == "" {
			originFlagName = "origin"
		} else {
			originFlagName = fmt.Sprintf("%v.origin", cmdPrefix)
		}

		originFlagValue, err := cmd.Flags().GetString(originFlagName)
		if err != nil {
			return err, false
		}
		m.Origin = &originFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUispBackupSizeFlags(depth int, m *models.UispBackup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sizeFlagName := fmt.Sprintf("%v.size", cmdPrefix)
	if cmd.Flags().Changed(sizeFlagName) {

		var sizeFlagName string
		if cmdPrefix == "" {
			sizeFlagName = "size"
		} else {
			sizeFlagName = fmt.Sprintf("%v.size", cmdPrefix)
		}

		sizeFlagValue, err := cmd.Flags().GetInt64(sizeFlagName)
		if err != nil {
			return err, false
		}
		m.Size = &sizeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUispBackupStateFlags(depth int, m *models.UispBackup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	stateFlagName := fmt.Sprintf("%v.state", cmdPrefix)
	if cmd.Flags().Changed(stateFlagName) {

		var stateFlagName string
		if cmdPrefix == "" {
			stateFlagName = "state"
		} else {
			stateFlagName = fmt.Sprintf("%v.state", cmdPrefix)
		}

		stateFlagValue, err := cmd.Flags().GetString(stateFlagName)
		if err != nil {
			return err, false
		}
		m.State = &stateFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUispBackupUnmsVersionFlags(depth int, m *models.UispBackup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	unmsVersionFlagName := fmt.Sprintf("%v.unmsVersion", cmdPrefix)
	if cmd.Flags().Changed(unmsVersionFlagName) {
		// info: complex object unmsVersion UnmsVersion is retrieved outside this Changed() block
	}
	unmsVersionFlagValue := m.UnmsVersion
	if swag.IsZero(unmsVersionFlagValue) {
		unmsVersionFlagValue = &models.UnmsVersion{}
	}

	err, unmsVersionAdded := retrieveModelUnmsVersionFlags(depth+1, unmsVersionFlagValue, unmsVersionFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || unmsVersionAdded
	if unmsVersionAdded {
		m.UnmsVersion = unmsVersionFlagValue
	}

	return nil, retAdded
}