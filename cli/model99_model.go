// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for Model99

// register flags to command
func registerModelModel99Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel99Auth(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel99ID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel99Networks(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel99Type(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel99Auth(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	authDescription := `Enum: ["md5","plaintext-password","off"]. `

	var authFlagName string
	if cmdPrefix == "" {
		authFlagName = "auth"
	} else {
		authFlagName = fmt.Sprintf("%v.auth", cmdPrefix)
	}

	var authFlagDefault string

	_ = cmd.PersistentFlags().String(authFlagName, authFlagDefault, authDescription)

	if err := cmd.RegisterFlagCompletionFunc(authFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["md5","plaintext-password","off"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerModel99ID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := ``

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

func registerModel99Networks(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: networks Networks array type is not supported by go-swagger cli yet

	return nil
}

func registerModel99Type(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["normal","nssa","stub"]. `

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
			if err := json.Unmarshal([]byte(`["normal","nssa","stub"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModel99Flags(depth int, m *models.Model99, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, authAdded := retrieveModel99AuthFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || authAdded

	err, idAdded := retrieveModel99IDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, networksAdded := retrieveModel99NetworksFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || networksAdded

	err, typeAdded := retrieveModel99TypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveModel99AuthFlags(depth int, m *models.Model99, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	authFlagName := fmt.Sprintf("%v.auth", cmdPrefix)
	if cmd.Flags().Changed(authFlagName) {

		var authFlagName string
		if cmdPrefix == "" {
			authFlagName = "auth"
		} else {
			authFlagName = fmt.Sprintf("%v.auth", cmdPrefix)
		}

		authFlagValue, err := cmd.Flags().GetString(authFlagName)
		if err != nil {
			return err, false
		}
		m.Auth = authFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel99IDFlags(depth int, m *models.Model99, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel99NetworksFlags(depth int, m *models.Model99, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	networksFlagName := fmt.Sprintf("%v.networks", cmdPrefix)
	if cmd.Flags().Changed(networksFlagName) {
		// warning: networks array type Networks is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveModel99TypeFlags(depth int, m *models.Model99, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Type = typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}