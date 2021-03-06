// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

// Schema cli for DeviceBackup

// register flags to command
func registerModelDeviceBackupFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDeviceBackupExtension(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceBackupFilename(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceBackupFirmwareVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceBackupID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceBackupNote(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceBackupPinned(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceBackupTimestamp(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceBackupType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceBackupExtension(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	extensionDescription := `Required. `

	var extensionFlagName string
	if cmdPrefix == "" {
		extensionFlagName = "extension"
	} else {
		extensionFlagName = fmt.Sprintf("%v.extension", cmdPrefix)
	}

	var extensionFlagDefault string

	_ = cmd.PersistentFlags().String(extensionFlagName, extensionFlagDefault, extensionDescription)

	return nil
}

func registerDeviceBackupFilename(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	filenameDescription := ``

	var filenameFlagName string
	if cmdPrefix == "" {
		filenameFlagName = "filename"
	} else {
		filenameFlagName = fmt.Sprintf("%v.filename", cmdPrefix)
	}

	var filenameFlagDefault string

	_ = cmd.PersistentFlags().String(filenameFlagName, filenameFlagDefault, filenameDescription)

	return nil
}

func registerDeviceBackupFirmwareVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: firmwareVersion Interface map type is not supported by go-swagger cli yet

	return nil
}

func registerDeviceBackupID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerDeviceBackupNote(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	noteDescription := ``

	var noteFlagName string
	if cmdPrefix == "" {
		noteFlagName = "note"
	} else {
		noteFlagName = fmt.Sprintf("%v.note", cmdPrefix)
	}

	var noteFlagDefault string

	_ = cmd.PersistentFlags().String(noteFlagName, noteFlagDefault, noteDescription)

	return nil
}

func registerDeviceBackupPinned(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	pinnedDescription := ``

	var pinnedFlagName string
	if cmdPrefix == "" {
		pinnedFlagName = "pinned"
	} else {
		pinnedFlagName = fmt.Sprintf("%v.pinned", cmdPrefix)
	}

	var pinnedFlagDefault bool

	_ = cmd.PersistentFlags().Bool(pinnedFlagName, pinnedFlagDefault, pinnedDescription)

	return nil
}

func registerDeviceBackupTimestamp(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	timestampDescription := `Required. `

	var timestampFlagName string
	if cmdPrefix == "" {
		timestampFlagName = "timestamp"
	} else {
		timestampFlagName = fmt.Sprintf("%v.timestamp", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(timestampFlagName, "", timestampDescription)

	return nil
}

func registerDeviceBackupType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["cfg","tar"]. Required. `

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	if err := cmd.RegisterFlagCompletionFunc(typeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["cfg","tar"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDeviceBackupFlags(depth int, m *models.DeviceBackup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, extensionAdded := retrieveDeviceBackupExtensionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || extensionAdded

	err, filenameAdded := retrieveDeviceBackupFilenameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || filenameAdded

	err, firmwareVersionAdded := retrieveDeviceBackupFirmwareVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || firmwareVersionAdded

	err, idAdded := retrieveDeviceBackupIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, noteAdded := retrieveDeviceBackupNoteFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || noteAdded

	err, pinnedAdded := retrieveDeviceBackupPinnedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pinnedAdded

	err, timestampAdded := retrieveDeviceBackupTimestampFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || timestampAdded

	err, typeAdded := retrieveDeviceBackupTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveDeviceBackupExtensionFlags(depth int, m *models.DeviceBackup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	extensionFlagName := fmt.Sprintf("%v.extension", cmdPrefix)
	if cmd.Flags().Changed(extensionFlagName) {

		var extensionFlagName string
		if cmdPrefix == "" {
			extensionFlagName = "extension"
		} else {
			extensionFlagName = fmt.Sprintf("%v.extension", cmdPrefix)
		}

		extensionFlagValue, err := cmd.Flags().GetString(extensionFlagName)
		if err != nil {
			return err, false
		}
		m.Extension = &extensionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceBackupFilenameFlags(depth int, m *models.DeviceBackup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	filenameFlagName := fmt.Sprintf("%v.filename", cmdPrefix)
	if cmd.Flags().Changed(filenameFlagName) {

		var filenameFlagName string
		if cmdPrefix == "" {
			filenameFlagName = "filename"
		} else {
			filenameFlagName = fmt.Sprintf("%v.filename", cmdPrefix)
		}

		filenameFlagValue, err := cmd.Flags().GetString(filenameFlagName)
		if err != nil {
			return err, false
		}
		m.Filename = filenameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceBackupFirmwareVersionFlags(depth int, m *models.DeviceBackup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	firmwareVersionFlagName := fmt.Sprintf("%v.firmwareVersion", cmdPrefix)
	if cmd.Flags().Changed(firmwareVersionFlagName) {
		// warning: firmwareVersion map type Interface is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDeviceBackupIDFlags(depth int, m *models.DeviceBackup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDeviceBackupNoteFlags(depth int, m *models.DeviceBackup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	noteFlagName := fmt.Sprintf("%v.note", cmdPrefix)
	if cmd.Flags().Changed(noteFlagName) {

		var noteFlagName string
		if cmdPrefix == "" {
			noteFlagName = "note"
		} else {
			noteFlagName = fmt.Sprintf("%v.note", cmdPrefix)
		}

		noteFlagValue, err := cmd.Flags().GetString(noteFlagName)
		if err != nil {
			return err, false
		}
		m.Note = noteFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceBackupPinnedFlags(depth int, m *models.DeviceBackup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pinnedFlagName := fmt.Sprintf("%v.pinned", cmdPrefix)
	if cmd.Flags().Changed(pinnedFlagName) {

		var pinnedFlagName string
		if cmdPrefix == "" {
			pinnedFlagName = "pinned"
		} else {
			pinnedFlagName = fmt.Sprintf("%v.pinned", cmdPrefix)
		}

		pinnedFlagValue, err := cmd.Flags().GetBool(pinnedFlagName)
		if err != nil {
			return err, false
		}
		m.Pinned = pinnedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceBackupTimestampFlags(depth int, m *models.DeviceBackup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	timestampFlagName := fmt.Sprintf("%v.timestamp", cmdPrefix)
	if cmd.Flags().Changed(timestampFlagName) {

		var timestampFlagName string
		if cmdPrefix == "" {
			timestampFlagName = "timestamp"
		} else {
			timestampFlagName = fmt.Sprintf("%v.timestamp", cmdPrefix)
		}

		timestampFlagValueStr, err := cmd.Flags().GetString(timestampFlagName)
		if err != nil {
			return err, false
		}
		var timestampFlagValue strfmt.DateTime
		if err := timestampFlagValue.UnmarshalText([]byte(timestampFlagValueStr)); err != nil {
			return err, false
		}
		m.Timestamp = &timestampFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceBackupTypeFlags(depth int, m *models.DeviceBackup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	typeFlagName := fmt.Sprintf("%v.type", cmdPrefix)
	if cmd.Flags().Changed(typeFlagName) {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = &typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
