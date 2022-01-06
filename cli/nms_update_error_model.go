// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// Schema cli for NmsUpdateError

// register flags to command
func registerModelNmsUpdateErrorFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerNmsUpdateErrorError(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNmsUpdateErrorMessage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNmsUpdateErrorMetadata(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNmsUpdateErrorTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerNmsUpdateErrorError(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	errorDescription := `Enum: ["DiskSpace","InstallationPackage","DockerStop","OldDocker","OldDockerCompose","OldUnms","Timeout","PullImages","Unknown"]. Required. Type of update error.`

	var errorFlagName string
	if cmdPrefix == "" {
		errorFlagName = "error"
	} else {
		errorFlagName = fmt.Sprintf("%v.error", cmdPrefix)
	}

	var errorFlagDefault string

	_ = cmd.PersistentFlags().String(errorFlagName, errorFlagDefault, errorDescription)

	if err := cmd.RegisterFlagCompletionFunc(errorFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["DiskSpace","InstallationPackage","DockerStop","OldDocker","OldDockerCompose","OldUnms","Timeout","PullImages","Unknown"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerNmsUpdateErrorMessage(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	messageDescription := `Required. Brief description of the update error.`

	var messageFlagName string
	if cmdPrefix == "" {
		messageFlagName = "message"
	} else {
		messageFlagName = fmt.Sprintf("%v.message", cmdPrefix)
	}

	var messageFlagDefault string

	_ = cmd.PersistentFlags().String(messageFlagName, messageFlagDefault, messageDescription)

	return nil
}

func registerNmsUpdateErrorMetadata(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var metadataFlagName string
	if cmdPrefix == "" {
		metadataFlagName = "metadata"
	} else {
		metadataFlagName = fmt.Sprintf("%v.metadata", cmdPrefix)
	}

	if err := registerModelMetadataFlags(depth+1, metadataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerNmsUpdateErrorTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: primitive time strfmt.Date is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelNmsUpdateErrorFlags(depth int, m *models.NmsUpdateError, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, errorAdded := retrieveNmsUpdateErrorErrorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || errorAdded

	err, messageAdded := retrieveNmsUpdateErrorMessageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || messageAdded

	err, metadataAdded := retrieveNmsUpdateErrorMetadataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || metadataAdded

	err, timeAdded := retrieveNmsUpdateErrorTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || timeAdded

	return nil, retAdded
}

func retrieveNmsUpdateErrorErrorFlags(depth int, m *models.NmsUpdateError, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	errorFlagName := fmt.Sprintf("%v.error", cmdPrefix)
	if cmd.Flags().Changed(errorFlagName) {

		var errorFlagName string
		if cmdPrefix == "" {
			errorFlagName = "error"
		} else {
			errorFlagName = fmt.Sprintf("%v.error", cmdPrefix)
		}

		errorFlagValue, err := cmd.Flags().GetString(errorFlagName)
		if err != nil {
			return err, false
		}
		m.Error = &errorFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveNmsUpdateErrorMessageFlags(depth int, m *models.NmsUpdateError, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	messageFlagName := fmt.Sprintf("%v.message", cmdPrefix)
	if cmd.Flags().Changed(messageFlagName) {

		var messageFlagName string
		if cmdPrefix == "" {
			messageFlagName = "message"
		} else {
			messageFlagName = fmt.Sprintf("%v.message", cmdPrefix)
		}

		messageFlagValue, err := cmd.Flags().GetString(messageFlagName)
		if err != nil {
			return err, false
		}
		m.Message = &messageFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveNmsUpdateErrorMetadataFlags(depth int, m *models.NmsUpdateError, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	metadataFlagName := fmt.Sprintf("%v.metadata", cmdPrefix)
	if cmd.Flags().Changed(metadataFlagName) {
		// info: complex object metadata Metadata is retrieved outside this Changed() block
	}
	metadataFlagValue := m.Metadata
	if swag.IsZero(metadataFlagValue) {
		metadataFlagValue = &models.Metadata{}
	}

	err, metadataAdded := retrieveModelMetadataFlags(depth+1, metadataFlagValue, metadataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || metadataAdded
	if metadataAdded {
		m.Metadata = metadataFlagValue
	}

	return nil, retAdded
}

func retrieveNmsUpdateErrorTimeFlags(depth int, m *models.NmsUpdateError, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	timeFlagName := fmt.Sprintf("%v.time", cmdPrefix)
	if cmd.Flags().Changed(timeFlagName) {

		// warning: primitive time strfmt.Date is not supported by go-swagger cli yet

		retAdded = true
	}

	return nil, retAdded
}
