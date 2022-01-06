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

// Schema cli for Model109

// register flags to command
func registerModelModel109Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel109Bytes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel109Description(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel109Destination(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel109Enabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel109ID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel109InInterface(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel109IPVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel109Log(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel109NegateInInterface(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel109NegateOutInterface(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel109OutInterface(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel109Packets(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel109Protocol(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel109Source(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel109Target(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel109TproxyMark(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel109TproxyOnIP(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel109TproxyOnPort(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel109Bytes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	bytesDescription := `Required. `

	var bytesFlagName string
	if cmdPrefix == "" {
		bytesFlagName = "bytes"
	} else {
		bytesFlagName = fmt.Sprintf("%v.bytes", cmdPrefix)
	}

	var bytesFlagDefault int64

	_ = cmd.PersistentFlags().Int64(bytesFlagName, bytesFlagDefault, bytesDescription)

	return nil
}

func registerModel109Description(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := ``

	var descriptionFlagName string
	if cmdPrefix == "" {
		descriptionFlagName = "description"
	} else {
		descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
	}

	var descriptionFlagDefault string

	_ = cmd.PersistentFlags().String(descriptionFlagName, descriptionFlagDefault, descriptionDescription)

	return nil
}

func registerModel109Destination(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var destinationFlagName string
	if cmdPrefix == "" {
		destinationFlagName = "destination"
	} else {
		destinationFlagName = fmt.Sprintf("%v.destination", cmdPrefix)
	}

	if err := registerModelDestinationFlags(depth+1, destinationFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel109Enabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	enabledDescription := ``

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

func registerModel109ID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

	var idFlagDefault int64

	_ = cmd.PersistentFlags().Int64(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerModel109InInterface(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var inInterfaceFlagName string
	if cmdPrefix == "" {
		inInterfaceFlagName = "inInterface"
	} else {
		inInterfaceFlagName = fmt.Sprintf("%v.inInterface", cmdPrefix)
	}

	if err := registerModelInInterface1Flags(depth+1, inInterfaceFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel109IPVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ipVersionDescription := `Enum: ["both","v4only","v6only"]. Required. `

	var ipVersionFlagName string
	if cmdPrefix == "" {
		ipVersionFlagName = "ipVersion"
	} else {
		ipVersionFlagName = fmt.Sprintf("%v.ipVersion", cmdPrefix)
	}

	var ipVersionFlagDefault string

	_ = cmd.PersistentFlags().String(ipVersionFlagName, ipVersionFlagDefault, ipVersionDescription)

	if err := cmd.RegisterFlagCompletionFunc(ipVersionFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["both","v4only","v6only"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerModel109Log(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	logDescription := ``

	var logFlagName string
	if cmdPrefix == "" {
		logFlagName = "log"
	} else {
		logFlagName = fmt.Sprintf("%v.log", cmdPrefix)
	}

	var logFlagDefault bool

	_ = cmd.PersistentFlags().Bool(logFlagName, logFlagDefault, logDescription)

	return nil
}

func registerModel109NegateInInterface(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	negateInInterfaceDescription := ``

	var negateInInterfaceFlagName string
	if cmdPrefix == "" {
		negateInInterfaceFlagName = "negateInInterface"
	} else {
		negateInInterfaceFlagName = fmt.Sprintf("%v.negateInInterface", cmdPrefix)
	}

	var negateInInterfaceFlagDefault bool

	_ = cmd.PersistentFlags().Bool(negateInInterfaceFlagName, negateInInterfaceFlagDefault, negateInInterfaceDescription)

	return nil
}

func registerModel109NegateOutInterface(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	negateOutInterfaceDescription := ``

	var negateOutInterfaceFlagName string
	if cmdPrefix == "" {
		negateOutInterfaceFlagName = "negateOutInterface"
	} else {
		negateOutInterfaceFlagName = fmt.Sprintf("%v.negateOutInterface", cmdPrefix)
	}

	var negateOutInterfaceFlagDefault bool

	_ = cmd.PersistentFlags().Bool(negateOutInterfaceFlagName, negateOutInterfaceFlagDefault, negateOutInterfaceDescription)

	return nil
}

func registerModel109OutInterface(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var outInterfaceFlagName string
	if cmdPrefix == "" {
		outInterfaceFlagName = "outInterface"
	} else {
		outInterfaceFlagName = fmt.Sprintf("%v.outInterface", cmdPrefix)
	}

	if err := registerModelOutInterface1Flags(depth+1, outInterfaceFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel109Packets(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	packetsDescription := `Required. `

	var packetsFlagName string
	if cmdPrefix == "" {
		packetsFlagName = "packets"
	} else {
		packetsFlagName = fmt.Sprintf("%v.packets", cmdPrefix)
	}

	var packetsFlagDefault int64

	_ = cmd.PersistentFlags().Int64(packetsFlagName, packetsFlagDefault, packetsDescription)

	return nil
}

func registerModel109Protocol(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	protocolDescription := `Required. `

	var protocolFlagName string
	if cmdPrefix == "" {
		protocolFlagName = "protocol"
	} else {
		protocolFlagName = fmt.Sprintf("%v.protocol", cmdPrefix)
	}

	var protocolFlagDefault string

	_ = cmd.PersistentFlags().String(protocolFlagName, protocolFlagDefault, protocolDescription)

	return nil
}

func registerModel109Source(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var sourceFlagName string
	if cmdPrefix == "" {
		sourceFlagName = "source"
	} else {
		sourceFlagName = fmt.Sprintf("%v.source", cmdPrefix)
	}

	if err := registerModelSourceFlags(depth+1, sourceFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel109Target(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	targetDescription := `Required. `

	var targetFlagName string
	if cmdPrefix == "" {
		targetFlagName = "target"
	} else {
		targetFlagName = fmt.Sprintf("%v.target", cmdPrefix)
	}

	var targetFlagDefault string

	_ = cmd.PersistentFlags().String(targetFlagName, targetFlagDefault, targetDescription)

	return nil
}

func registerModel109TproxyMark(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	tproxyMarkDescription := `Enum: [1]. `

	var tproxyMarkFlagName string
	if cmdPrefix == "" {
		tproxyMarkFlagName = "tproxyMark"
	} else {
		tproxyMarkFlagName = fmt.Sprintf("%v.tproxyMark", cmdPrefix)
	}

	var tproxyMarkFlagDefault int64

	_ = cmd.PersistentFlags().Int64(tproxyMarkFlagName, tproxyMarkFlagDefault, tproxyMarkDescription)

	if err := cmd.RegisterFlagCompletionFunc(tproxyMarkFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`[1]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerModel109TproxyOnIP(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	tproxyOnIpDescription := `Enum: ["both","v4only","v6only"]. `

	var tproxyOnIpFlagName string
	if cmdPrefix == "" {
		tproxyOnIpFlagName = "tproxyOnIp"
	} else {
		tproxyOnIpFlagName = fmt.Sprintf("%v.tproxyOnIp", cmdPrefix)
	}

	var tproxyOnIpFlagDefault string

	_ = cmd.PersistentFlags().String(tproxyOnIpFlagName, tproxyOnIpFlagDefault, tproxyOnIpDescription)

	if err := cmd.RegisterFlagCompletionFunc(tproxyOnIpFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["both","v4only","v6only"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerModel109TproxyOnPort(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	tproxyOnPortDescription := ``

	var tproxyOnPortFlagName string
	if cmdPrefix == "" {
		tproxyOnPortFlagName = "tproxyOnPort"
	} else {
		tproxyOnPortFlagName = fmt.Sprintf("%v.tproxyOnPort", cmdPrefix)
	}

	var tproxyOnPortFlagDefault int64

	_ = cmd.PersistentFlags().Int64(tproxyOnPortFlagName, tproxyOnPortFlagDefault, tproxyOnPortDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModel109Flags(depth int, m *models.Model109, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, bytesAdded := retrieveModel109BytesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || bytesAdded

	err, descriptionAdded := retrieveModel109DescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, destinationAdded := retrieveModel109DestinationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || destinationAdded

	err, enabledAdded := retrieveModel109EnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || enabledAdded

	err, idAdded := retrieveModel109IDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, inInterfaceAdded := retrieveModel109InInterfaceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || inInterfaceAdded

	err, ipVersionAdded := retrieveModel109IPVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipVersionAdded

	err, logAdded := retrieveModel109LogFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || logAdded

	err, negateInInterfaceAdded := retrieveModel109NegateInInterfaceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || negateInInterfaceAdded

	err, negateOutInterfaceAdded := retrieveModel109NegateOutInterfaceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || negateOutInterfaceAdded

	err, outInterfaceAdded := retrieveModel109OutInterfaceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || outInterfaceAdded

	err, packetsAdded := retrieveModel109PacketsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || packetsAdded

	err, protocolAdded := retrieveModel109ProtocolFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || protocolAdded

	err, sourceAdded := retrieveModel109SourceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sourceAdded

	err, targetAdded := retrieveModel109TargetFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || targetAdded

	err, tproxyMarkAdded := retrieveModel109TproxyMarkFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tproxyMarkAdded

	err, tproxyOnIpAdded := retrieveModel109TproxyOnIPFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tproxyOnIpAdded

	err, tproxyOnPortAdded := retrieveModel109TproxyOnPortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tproxyOnPortAdded

	return nil, retAdded
}

func retrieveModel109BytesFlags(depth int, m *models.Model109, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	bytesFlagName := fmt.Sprintf("%v.bytes", cmdPrefix)
	if cmd.Flags().Changed(bytesFlagName) {

		var bytesFlagName string
		if cmdPrefix == "" {
			bytesFlagName = "bytes"
		} else {
			bytesFlagName = fmt.Sprintf("%v.bytes", cmdPrefix)
		}

		bytesFlagValue, err := cmd.Flags().GetInt64(bytesFlagName)
		if err != nil {
			return err, false
		}
		m.Bytes = &bytesFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel109DescriptionFlags(depth int, m *models.Model109, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	descriptionFlagName := fmt.Sprintf("%v.description", cmdPrefix)
	if cmd.Flags().Changed(descriptionFlagName) {

		var descriptionFlagName string
		if cmdPrefix == "" {
			descriptionFlagName = "description"
		} else {
			descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
		}

		descriptionFlagValue, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			return err, false
		}
		m.Description = descriptionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel109DestinationFlags(depth int, m *models.Model109, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	destinationFlagName := fmt.Sprintf("%v.destination", cmdPrefix)
	if cmd.Flags().Changed(destinationFlagName) {
		// info: complex object destination Destination is retrieved outside this Changed() block
	}
	destinationFlagValue := m.Destination
	if swag.IsZero(destinationFlagValue) {
		destinationFlagValue = &models.Destination{}
	}

	err, destinationAdded := retrieveModelDestinationFlags(depth+1, destinationFlagValue, destinationFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || destinationAdded
	if destinationAdded {
		m.Destination = destinationFlagValue
	}

	return nil, retAdded
}

func retrieveModel109EnabledFlags(depth int, m *models.Model109, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Enabled = enabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel109IDFlags(depth int, m *models.Model109, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

		idFlagValue, err := cmd.Flags().GetInt64(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = &idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel109InInterfaceFlags(depth int, m *models.Model109, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	inInterfaceFlagName := fmt.Sprintf("%v.inInterface", cmdPrefix)
	if cmd.Flags().Changed(inInterfaceFlagName) {
		// info: complex object inInterface InInterface1 is retrieved outside this Changed() block
	}
	inInterfaceFlagValue := m.InInterface
	if swag.IsZero(inInterfaceFlagValue) {
		inInterfaceFlagValue = &models.InInterface1{}
	}

	err, inInterfaceAdded := retrieveModelInInterface1Flags(depth+1, inInterfaceFlagValue, inInterfaceFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || inInterfaceAdded
	if inInterfaceAdded {
		m.InInterface = inInterfaceFlagValue
	}

	return nil, retAdded
}

func retrieveModel109IPVersionFlags(depth int, m *models.Model109, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ipVersionFlagName := fmt.Sprintf("%v.ipVersion", cmdPrefix)
	if cmd.Flags().Changed(ipVersionFlagName) {

		var ipVersionFlagName string
		if cmdPrefix == "" {
			ipVersionFlagName = "ipVersion"
		} else {
			ipVersionFlagName = fmt.Sprintf("%v.ipVersion", cmdPrefix)
		}

		ipVersionFlagValue, err := cmd.Flags().GetString(ipVersionFlagName)
		if err != nil {
			return err, false
		}
		m.IPVersion = &ipVersionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel109LogFlags(depth int, m *models.Model109, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	logFlagName := fmt.Sprintf("%v.log", cmdPrefix)
	if cmd.Flags().Changed(logFlagName) {

		var logFlagName string
		if cmdPrefix == "" {
			logFlagName = "log"
		} else {
			logFlagName = fmt.Sprintf("%v.log", cmdPrefix)
		}

		logFlagValue, err := cmd.Flags().GetBool(logFlagName)
		if err != nil {
			return err, false
		}
		m.Log = logFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel109NegateInInterfaceFlags(depth int, m *models.Model109, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	negateInInterfaceFlagName := fmt.Sprintf("%v.negateInInterface", cmdPrefix)
	if cmd.Flags().Changed(negateInInterfaceFlagName) {

		var negateInInterfaceFlagName string
		if cmdPrefix == "" {
			negateInInterfaceFlagName = "negateInInterface"
		} else {
			negateInInterfaceFlagName = fmt.Sprintf("%v.negateInInterface", cmdPrefix)
		}

		negateInInterfaceFlagValue, err := cmd.Flags().GetBool(negateInInterfaceFlagName)
		if err != nil {
			return err, false
		}
		m.NegateInInterface = negateInInterfaceFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel109NegateOutInterfaceFlags(depth int, m *models.Model109, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	negateOutInterfaceFlagName := fmt.Sprintf("%v.negateOutInterface", cmdPrefix)
	if cmd.Flags().Changed(negateOutInterfaceFlagName) {

		var negateOutInterfaceFlagName string
		if cmdPrefix == "" {
			negateOutInterfaceFlagName = "negateOutInterface"
		} else {
			negateOutInterfaceFlagName = fmt.Sprintf("%v.negateOutInterface", cmdPrefix)
		}

		negateOutInterfaceFlagValue, err := cmd.Flags().GetBool(negateOutInterfaceFlagName)
		if err != nil {
			return err, false
		}
		m.NegateOutInterface = negateOutInterfaceFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel109OutInterfaceFlags(depth int, m *models.Model109, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	outInterfaceFlagName := fmt.Sprintf("%v.outInterface", cmdPrefix)
	if cmd.Flags().Changed(outInterfaceFlagName) {
		// info: complex object outInterface OutInterface1 is retrieved outside this Changed() block
	}
	outInterfaceFlagValue := m.OutInterface
	if swag.IsZero(outInterfaceFlagValue) {
		outInterfaceFlagValue = &models.OutInterface1{}
	}

	err, outInterfaceAdded := retrieveModelOutInterface1Flags(depth+1, outInterfaceFlagValue, outInterfaceFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || outInterfaceAdded
	if outInterfaceAdded {
		m.OutInterface = outInterfaceFlagValue
	}

	return nil, retAdded
}

func retrieveModel109PacketsFlags(depth int, m *models.Model109, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	packetsFlagName := fmt.Sprintf("%v.packets", cmdPrefix)
	if cmd.Flags().Changed(packetsFlagName) {

		var packetsFlagName string
		if cmdPrefix == "" {
			packetsFlagName = "packets"
		} else {
			packetsFlagName = fmt.Sprintf("%v.packets", cmdPrefix)
		}

		packetsFlagValue, err := cmd.Flags().GetInt64(packetsFlagName)
		if err != nil {
			return err, false
		}
		m.Packets = &packetsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel109ProtocolFlags(depth int, m *models.Model109, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel109SourceFlags(depth int, m *models.Model109, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sourceFlagName := fmt.Sprintf("%v.source", cmdPrefix)
	if cmd.Flags().Changed(sourceFlagName) {
		// info: complex object source Source is retrieved outside this Changed() block
	}
	sourceFlagValue := m.Source
	if swag.IsZero(sourceFlagValue) {
		sourceFlagValue = &models.Source{}
	}

	err, sourceAdded := retrieveModelSourceFlags(depth+1, sourceFlagValue, sourceFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sourceAdded
	if sourceAdded {
		m.Source = sourceFlagValue
	}

	return nil, retAdded
}

func retrieveModel109TargetFlags(depth int, m *models.Model109, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	targetFlagName := fmt.Sprintf("%v.target", cmdPrefix)
	if cmd.Flags().Changed(targetFlagName) {

		var targetFlagName string
		if cmdPrefix == "" {
			targetFlagName = "target"
		} else {
			targetFlagName = fmt.Sprintf("%v.target", cmdPrefix)
		}

		targetFlagValue, err := cmd.Flags().GetString(targetFlagName)
		if err != nil {
			return err, false
		}
		m.Target = &targetFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel109TproxyMarkFlags(depth int, m *models.Model109, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	tproxyMarkFlagName := fmt.Sprintf("%v.tproxyMark", cmdPrefix)
	if cmd.Flags().Changed(tproxyMarkFlagName) {

		var tproxyMarkFlagName string
		if cmdPrefix == "" {
			tproxyMarkFlagName = "tproxyMark"
		} else {
			tproxyMarkFlagName = fmt.Sprintf("%v.tproxyMark", cmdPrefix)
		}

		tproxyMarkFlagValue, err := cmd.Flags().GetInt64(tproxyMarkFlagName)
		if err != nil {
			return err, false
		}
		m.TproxyMark = tproxyMarkFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel109TproxyOnIPFlags(depth int, m *models.Model109, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	tproxyOnIpFlagName := fmt.Sprintf("%v.tproxyOnIp", cmdPrefix)
	if cmd.Flags().Changed(tproxyOnIpFlagName) {

		var tproxyOnIpFlagName string
		if cmdPrefix == "" {
			tproxyOnIpFlagName = "tproxyOnIp"
		} else {
			tproxyOnIpFlagName = fmt.Sprintf("%v.tproxyOnIp", cmdPrefix)
		}

		tproxyOnIpFlagValue, err := cmd.Flags().GetString(tproxyOnIpFlagName)
		if err != nil {
			return err, false
		}
		m.TproxyOnIP = tproxyOnIpFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel109TproxyOnPortFlags(depth int, m *models.Model109, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	tproxyOnPortFlagName := fmt.Sprintf("%v.tproxyOnPort", cmdPrefix)
	if cmd.Flags().Changed(tproxyOnPortFlagName) {

		var tproxyOnPortFlagName string
		if cmdPrefix == "" {
			tproxyOnPortFlagName = "tproxyOnPort"
		} else {
			tproxyOnPortFlagName = fmt.Sprintf("%v.tproxyOnPort", cmdPrefix)
		}

		tproxyOnPortFlagValue, err := cmd.Flags().GetInt64(tproxyOnPortFlagName)
		if err != nil {
			return err, false
		}
		m.TproxyOnPort = &tproxyOnPortFlagValue

		retAdded = true
	}

	return nil, retAdded
}
