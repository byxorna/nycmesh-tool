// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for Model45

// register flags to command
func registerModelModel45Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel45Comment(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel45Enabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel45LanAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel45LanPort(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel45Protocol(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel45WanPort(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel45Comment(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	commentDescription := `Required. Custom description text`

	var commentFlagName string
	if cmdPrefix == "" {
		commentFlagName = "comment"
	} else {
		commentFlagName = fmt.Sprintf("%v.comment", cmdPrefix)
	}

	var commentFlagDefault string

	_ = cmd.PersistentFlags().String(commentFlagName, commentFlagDefault, commentDescription)

	return nil
}

func registerModel45Enabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	enabledDescription := `Required. Set to true if rule is enabled`

	var enabledFlagName string
	if cmdPrefix == "" {
		enabledFlagName = "enabled"
	} else {
		enabledFlagName = fmt.Sprintf("%v.enabled", cmdPrefix)
	}

	var enabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(enabledFlagName, enabledFlagDefault, enabledDescription)

	return nil
}

func registerModel45LanAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	lanAddressDescription := `Required. IPv4 address of the target LAN device`

	var lanAddressFlagName string
	if cmdPrefix == "" {
		lanAddressFlagName = "lanAddress"
	} else {
		lanAddressFlagName = fmt.Sprintf("%v.lanAddress", cmdPrefix)
	}

	var lanAddressFlagDefault string

	_ = cmd.PersistentFlags().String(lanAddressFlagName, lanAddressFlagDefault, lanAddressDescription)

	return nil
}

func registerModel45LanPort(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: lanPort LanPort array type is not supported by go-swagger cli yet

	return nil
}

func registerModel45Protocol(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	protocolDescription := `Enum: ["tcpudp","tcp","udp"]. Required. Forwarded protocols`

	var protocolFlagName string
	if cmdPrefix == "" {
		protocolFlagName = "protocol"
	} else {
		protocolFlagName = fmt.Sprintf("%v.protocol", cmdPrefix)
	}

	var protocolFlagDefault string

	_ = cmd.PersistentFlags().String(protocolFlagName, protocolFlagDefault, protocolDescription)

	if err := cmd.RegisterFlagCompletionFunc(protocolFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["tcpudp","tcp","udp"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerModel45WanPort(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: wanPort WanPort array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModel45Flags(depth int, m *models.Model45, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, commentAdded := retrieveModel45CommentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || commentAdded

	err, enabledAdded := retrieveModel45EnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || enabledAdded

	err, lanAddressAdded := retrieveModel45LanAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || lanAddressAdded

	err, lanPortAdded := retrieveModel45LanPortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || lanPortAdded

	err, protocolAdded := retrieveModel45ProtocolFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || protocolAdded

	err, wanPortAdded := retrieveModel45WanPortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || wanPortAdded

	return nil, retAdded
}

func retrieveModel45CommentFlags(depth int, m *models.Model45, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	commentFlagName := fmt.Sprintf("%v.comment", cmdPrefix)
	if cmd.Flags().Changed(commentFlagName) {

		var commentFlagName string
		if cmdPrefix == "" {
			commentFlagName = "comment"
		} else {
			commentFlagName = fmt.Sprintf("%v.comment", cmdPrefix)
		}

		commentFlagValue, err := cmd.Flags().GetString(commentFlagName)
		if err != nil {
			return err, false
		}
		m.Comment = &commentFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel45EnabledFlags(depth int, m *models.Model45, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	enabledFlagName := fmt.Sprintf("%v.enabled", cmdPrefix)
	if cmd.Flags().Changed(enabledFlagName) {

		var enabledFlagName string
		if cmdPrefix == "" {
			enabledFlagName = "enabled"
		} else {
			enabledFlagName = fmt.Sprintf("%v.enabled", cmdPrefix)
		}

		enabledFlagValue, err := cmd.Flags().GetBool(enabledFlagName)
		if err != nil {
			return err, false
		}
		m.Enabled = &enabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel45LanAddressFlags(depth int, m *models.Model45, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	lanAddressFlagName := fmt.Sprintf("%v.lanAddress", cmdPrefix)
	if cmd.Flags().Changed(lanAddressFlagName) {

		var lanAddressFlagName string
		if cmdPrefix == "" {
			lanAddressFlagName = "lanAddress"
		} else {
			lanAddressFlagName = fmt.Sprintf("%v.lanAddress", cmdPrefix)
		}

		lanAddressFlagValue, err := cmd.Flags().GetString(lanAddressFlagName)
		if err != nil {
			return err, false
		}
		m.LanAddress = &lanAddressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel45LanPortFlags(depth int, m *models.Model45, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	lanPortFlagName := fmt.Sprintf("%v.lanPort", cmdPrefix)
	if cmd.Flags().Changed(lanPortFlagName) {
		// warning: lanPort array type LanPort is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveModel45ProtocolFlags(depth int, m *models.Model45, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	protocolFlagName := fmt.Sprintf("%v.protocol", cmdPrefix)
	if cmd.Flags().Changed(protocolFlagName) {

		var protocolFlagName string
		if cmdPrefix == "" {
			protocolFlagName = "protocol"
		} else {
			protocolFlagName = fmt.Sprintf("%v.protocol", cmdPrefix)
		}

		protocolFlagValue, err := cmd.Flags().GetString(protocolFlagName)
		if err != nil {
			return err, false
		}
		m.Protocol = &protocolFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel45WanPortFlags(depth int, m *models.Model45, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	wanPortFlagName := fmt.Sprintf("%v.wanPort", cmdPrefix)
	if cmd.Flags().Changed(wanPortFlagName) {
		// warning: wanPort array type WanPort is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
