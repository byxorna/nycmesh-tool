// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for Model110

// register flags to command
func registerModelModel110Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel110ApplicationID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel110CategoryID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel110ConnectionState(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel110Description(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel110Destination(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel110Enabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel110IcmpType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel110Icmpv6Type(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel110ID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel110InInterface(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel110IPVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel110IpsecPolicy(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel110Log(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel110NegateInInterface(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel110NegateOutInterface(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel110OutInterface(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel110Protocol(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel110RejectWith(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel110Source(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel110Target(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel110ApplicationID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	applicationIdDescription := ``

	var applicationIdFlagName string
	if cmdPrefix == "" {
		applicationIdFlagName = "applicationId"
	} else {
		applicationIdFlagName = fmt.Sprintf("%v.applicationId", cmdPrefix)
	}

	var applicationIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(applicationIdFlagName, applicationIdFlagDefault, applicationIdDescription)

	return nil
}

func registerModel110CategoryID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	categoryIdDescription := ``

	var categoryIdFlagName string
	if cmdPrefix == "" {
		categoryIdFlagName = "categoryId"
	} else {
		categoryIdFlagName = fmt.Sprintf("%v.categoryId", cmdPrefix)
	}

	var categoryIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(categoryIdFlagName, categoryIdFlagDefault, categoryIdDescription)

	return nil
}

func registerModel110ConnectionState(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: connectionState ConnectionState array type is not supported by go-swagger cli yet

	return nil
}

func registerModel110Description(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModel110Destination(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModel110Enabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModel110IcmpType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	icmpTypeDescription := `Enum: ["address-mask-reply","address-mask-request","any","communication-prohibited","destination-unreachable","echo-reply","echo-request","fragmentation-needed","host-precedence-violation","host-prohibited","host-redirect","host-unknown","host-unreachable","ip-header-bad","network-prohibited","network-redirect","network-unknown","network-unreachable","parameter-problem","port-unreachable","precedence-cutoff","protocol-unreachable","redirect","required-option-missing","router-advertisement","router-solicitation","source-quench","source-route-failed","time-exceeded","timestamp-reply","timestamp-request","TOS-host-redirect","TOS-host-unreachable","TOS-network-redirect","TOS-network-unreachable","ttl-zero-during-reassembly","ttl-zero-during-transit"]. `

	var icmpTypeFlagName string
	if cmdPrefix == "" {
		icmpTypeFlagName = "icmpType"
	} else {
		icmpTypeFlagName = fmt.Sprintf("%v.icmpType", cmdPrefix)
	}

	var icmpTypeFlagDefault string

	_ = cmd.PersistentFlags().String(icmpTypeFlagName, icmpTypeFlagDefault, icmpTypeDescription)

	if err := cmd.RegisterFlagCompletionFunc(icmpTypeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["address-mask-reply","address-mask-request","any","communication-prohibited","destination-unreachable","echo-reply","echo-request","fragmentation-needed","host-precedence-violation","host-prohibited","host-redirect","host-unknown","host-unreachable","ip-header-bad","network-prohibited","network-redirect","network-unknown","network-unreachable","parameter-problem","port-unreachable","precedence-cutoff","protocol-unreachable","redirect","required-option-missing","router-advertisement","router-solicitation","source-quench","source-route-failed","time-exceeded","timestamp-reply","timestamp-request","TOS-host-redirect","TOS-host-unreachable","TOS-network-redirect","TOS-network-unreachable","ttl-zero-during-reassembly","ttl-zero-during-transit"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerModel110Icmpv6Type(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	icmpv6TypeDescription := `Enum: ["address-unreachable","bad-header","beyond-scope","communication-prohibited","destination-unreachable","echo-reply","echo-request","failed-policy","neighbor-advertisement","neighbor-solicitation","no-route","packet-too-big","parameter-problem","port-unreachable","redirect","reject-route","router-advertisement","router-solicitation","time-exceeded","ttl-zero-during-reassembly","ttl-zero-during-transit","unknown-header-type","unknown-option"]. `

	var icmpv6TypeFlagName string
	if cmdPrefix == "" {
		icmpv6TypeFlagName = "icmpv6Type"
	} else {
		icmpv6TypeFlagName = fmt.Sprintf("%v.icmpv6Type", cmdPrefix)
	}

	var icmpv6TypeFlagDefault string

	_ = cmd.PersistentFlags().String(icmpv6TypeFlagName, icmpv6TypeFlagDefault, icmpv6TypeDescription)

	if err := cmd.RegisterFlagCompletionFunc(icmpv6TypeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["address-unreachable","bad-header","beyond-scope","communication-prohibited","destination-unreachable","echo-reply","echo-request","failed-policy","neighbor-advertisement","neighbor-solicitation","no-route","packet-too-big","parameter-problem","port-unreachable","redirect","reject-route","router-advertisement","router-solicitation","time-exceeded","ttl-zero-during-reassembly","ttl-zero-during-transit","unknown-header-type","unknown-option"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerModel110ID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModel110InInterface(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModel110IPVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModel110IpsecPolicy(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ipsecPolicyDescription := `Enum: ["match_inbound_ipsec","match_inbound_nonipsec"]. `

	var ipsecPolicyFlagName string
	if cmdPrefix == "" {
		ipsecPolicyFlagName = "ipsecPolicy"
	} else {
		ipsecPolicyFlagName = fmt.Sprintf("%v.ipsecPolicy", cmdPrefix)
	}

	var ipsecPolicyFlagDefault string

	_ = cmd.PersistentFlags().String(ipsecPolicyFlagName, ipsecPolicyFlagDefault, ipsecPolicyDescription)

	if err := cmd.RegisterFlagCompletionFunc(ipsecPolicyFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["match_inbound_ipsec","match_inbound_nonipsec"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerModel110Log(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModel110NegateInInterface(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModel110NegateOutInterface(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModel110OutInterface(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModel110Protocol(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModel110RejectWith(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	rejectWithDescription := `Enum: ["icmp_net_unreachable","icmp_host_unreachable","icmp_port_unreachable","icmp_proto_unreachable","icmp_net_prohibited","icmp_host_prohibited","icmp_admin_prohibited","tcp_reset,"]. `

	var rejectWithFlagName string
	if cmdPrefix == "" {
		rejectWithFlagName = "rejectWith"
	} else {
		rejectWithFlagName = fmt.Sprintf("%v.rejectWith", cmdPrefix)
	}

	var rejectWithFlagDefault string

	_ = cmd.PersistentFlags().String(rejectWithFlagName, rejectWithFlagDefault, rejectWithDescription)

	if err := cmd.RegisterFlagCompletionFunc(rejectWithFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["icmp_net_unreachable","icmp_host_unreachable","icmp_port_unreachable","icmp_proto_unreachable","icmp_net_prohibited","icmp_host_prohibited","icmp_admin_prohibited","tcp_reset,"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerModel110Source(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModel110Target(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModel110Flags(depth int, m *models.Model110, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, applicationIdAdded := retrieveModel110ApplicationIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || applicationIdAdded

	err, categoryIdAdded := retrieveModel110CategoryIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || categoryIdAdded

	err, connectionStateAdded := retrieveModel110ConnectionStateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || connectionStateAdded

	err, descriptionAdded := retrieveModel110DescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, destinationAdded := retrieveModel110DestinationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || destinationAdded

	err, enabledAdded := retrieveModel110EnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || enabledAdded

	err, icmpTypeAdded := retrieveModel110IcmpTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || icmpTypeAdded

	err, icmpv6TypeAdded := retrieveModel110Icmpv6TypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || icmpv6TypeAdded

	err, idAdded := retrieveModel110IDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, inInterfaceAdded := retrieveModel110InInterfaceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || inInterfaceAdded

	err, ipVersionAdded := retrieveModel110IPVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipVersionAdded

	err, ipsecPolicyAdded := retrieveModel110IpsecPolicyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipsecPolicyAdded

	err, logAdded := retrieveModel110LogFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || logAdded

	err, negateInInterfaceAdded := retrieveModel110NegateInInterfaceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || negateInInterfaceAdded

	err, negateOutInterfaceAdded := retrieveModel110NegateOutInterfaceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || negateOutInterfaceAdded

	err, outInterfaceAdded := retrieveModel110OutInterfaceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || outInterfaceAdded

	err, protocolAdded := retrieveModel110ProtocolFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || protocolAdded

	err, rejectWithAdded := retrieveModel110RejectWithFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || rejectWithAdded

	err, sourceAdded := retrieveModel110SourceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sourceAdded

	err, targetAdded := retrieveModel110TargetFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || targetAdded

	return nil, retAdded
}

func retrieveModel110ApplicationIDFlags(depth int, m *models.Model110, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	applicationIdFlagName := fmt.Sprintf("%v.applicationId", cmdPrefix)
	if cmd.Flags().Changed(applicationIdFlagName) {

		var applicationIdFlagName string
		if cmdPrefix == "" {
			applicationIdFlagName = "applicationId"
		} else {
			applicationIdFlagName = fmt.Sprintf("%v.applicationId", cmdPrefix)
		}

		applicationIdFlagValue, err := cmd.Flags().GetInt64(applicationIdFlagName)
		if err != nil {
			return err, false
		}
		m.ApplicationID = applicationIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel110CategoryIDFlags(depth int, m *models.Model110, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	categoryIdFlagName := fmt.Sprintf("%v.categoryId", cmdPrefix)
	if cmd.Flags().Changed(categoryIdFlagName) {

		var categoryIdFlagName string
		if cmdPrefix == "" {
			categoryIdFlagName = "categoryId"
		} else {
			categoryIdFlagName = fmt.Sprintf("%v.categoryId", cmdPrefix)
		}

		categoryIdFlagValue, err := cmd.Flags().GetInt64(categoryIdFlagName)
		if err != nil {
			return err, false
		}
		m.CategoryID = &categoryIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel110ConnectionStateFlags(depth int, m *models.Model110, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	connectionStateFlagName := fmt.Sprintf("%v.connectionState", cmdPrefix)
	if cmd.Flags().Changed(connectionStateFlagName) {
		// warning: connectionState array type ConnectionState is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveModel110DescriptionFlags(depth int, m *models.Model110, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel110DestinationFlags(depth int, m *models.Model110, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel110EnabledFlags(depth int, m *models.Model110, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel110IcmpTypeFlags(depth int, m *models.Model110, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	icmpTypeFlagName := fmt.Sprintf("%v.icmpType", cmdPrefix)
	if cmd.Flags().Changed(icmpTypeFlagName) {

		var icmpTypeFlagName string
		if cmdPrefix == "" {
			icmpTypeFlagName = "icmpType"
		} else {
			icmpTypeFlagName = fmt.Sprintf("%v.icmpType", cmdPrefix)
		}

		icmpTypeFlagValue, err := cmd.Flags().GetString(icmpTypeFlagName)
		if err != nil {
			return err, false
		}
		m.IcmpType = icmpTypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel110Icmpv6TypeFlags(depth int, m *models.Model110, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	icmpv6TypeFlagName := fmt.Sprintf("%v.icmpv6Type", cmdPrefix)
	if cmd.Flags().Changed(icmpv6TypeFlagName) {

		var icmpv6TypeFlagName string
		if cmdPrefix == "" {
			icmpv6TypeFlagName = "icmpv6Type"
		} else {
			icmpv6TypeFlagName = fmt.Sprintf("%v.icmpv6Type", cmdPrefix)
		}

		icmpv6TypeFlagValue, err := cmd.Flags().GetString(icmpv6TypeFlagName)
		if err != nil {
			return err, false
		}
		m.Icmpv6Type = icmpv6TypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel110IDFlags(depth int, m *models.Model110, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel110InInterfaceFlags(depth int, m *models.Model110, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel110IPVersionFlags(depth int, m *models.Model110, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel110IpsecPolicyFlags(depth int, m *models.Model110, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ipsecPolicyFlagName := fmt.Sprintf("%v.ipsecPolicy", cmdPrefix)
	if cmd.Flags().Changed(ipsecPolicyFlagName) {

		var ipsecPolicyFlagName string
		if cmdPrefix == "" {
			ipsecPolicyFlagName = "ipsecPolicy"
		} else {
			ipsecPolicyFlagName = fmt.Sprintf("%v.ipsecPolicy", cmdPrefix)
		}

		ipsecPolicyFlagValue, err := cmd.Flags().GetString(ipsecPolicyFlagName)
		if err != nil {
			return err, false
		}
		m.IpsecPolicy = ipsecPolicyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel110LogFlags(depth int, m *models.Model110, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel110NegateInInterfaceFlags(depth int, m *models.Model110, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel110NegateOutInterfaceFlags(depth int, m *models.Model110, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel110OutInterfaceFlags(depth int, m *models.Model110, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel110ProtocolFlags(depth int, m *models.Model110, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel110RejectWithFlags(depth int, m *models.Model110, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	rejectWithFlagName := fmt.Sprintf("%v.rejectWith", cmdPrefix)
	if cmd.Flags().Changed(rejectWithFlagName) {

		var rejectWithFlagName string
		if cmdPrefix == "" {
			rejectWithFlagName = "rejectWith"
		} else {
			rejectWithFlagName = fmt.Sprintf("%v.rejectWith", cmdPrefix)
		}

		rejectWithFlagValue, err := cmd.Flags().GetString(rejectWithFlagName)
		if err != nil {
			return err, false
		}
		m.RejectWith = rejectWithFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel110SourceFlags(depth int, m *models.Model110, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel110TargetFlags(depth int, m *models.Model110, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
