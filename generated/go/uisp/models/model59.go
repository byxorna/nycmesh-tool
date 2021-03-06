// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Model59 model 59
//
// swagger:model Model 59
type Model59 struct {

	// application Id
	// Minimum: 1
	ApplicationID int64 `json:"applicationId,omitempty"`

	// bytes
	// Required: true
	// Minimum: 0
	Bytes *int64 `json:"bytes"`

	// category Id
	// Minimum: 0
	CategoryID *int64 `json:"categoryId,omitempty"`

	// connection state
	// Required: true
	ConnectionState ConnectionState `json:"connectionState"`

	// description
	Description string `json:"description,omitempty"`

	// destination
	Destination *Destination `json:"destination,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// icmp type
	// Enum: [address-mask-reply address-mask-request any communication-prohibited destination-unreachable echo-reply echo-request fragmentation-needed host-precedence-violation host-prohibited host-redirect host-unknown host-unreachable ip-header-bad network-prohibited network-redirect network-unknown network-unreachable parameter-problem port-unreachable precedence-cutoff protocol-unreachable redirect required-option-missing router-advertisement router-solicitation source-quench source-route-failed time-exceeded timestamp-reply timestamp-request TOS-host-redirect TOS-host-unreachable TOS-network-redirect TOS-network-unreachable ttl-zero-during-reassembly ttl-zero-during-transit]
	IcmpType string `json:"icmpType,omitempty"`

	// icmpv6 type
	// Enum: [address-unreachable bad-header beyond-scope communication-prohibited destination-unreachable echo-reply echo-request failed-policy neighbor-advertisement neighbor-solicitation no-route packet-too-big parameter-problem port-unreachable redirect reject-route router-advertisement router-solicitation time-exceeded ttl-zero-during-reassembly ttl-zero-during-transit unknown-header-type unknown-option]
	Icmpv6Type string `json:"icmpv6Type,omitempty"`

	// id
	// Minimum: 0
	ID *int64 `json:"id,omitempty"`

	// in interface
	InInterface *InInterface1 `json:"inInterface,omitempty"`

	// ip version
	// Required: true
	// Enum: [both v4only v6only]
	IPVersion *string `json:"ipVersion"`

	// ipsec policy
	// Enum: [match_inbound_ipsec match_inbound_nonipsec]
	IpsecPolicy string `json:"ipsecPolicy,omitempty"`

	// log
	Log bool `json:"log,omitempty"`

	// negate in interface
	NegateInInterface bool `json:"negateInInterface,omitempty"`

	// negate out interface
	NegateOutInterface bool `json:"negateOutInterface,omitempty"`

	// out interface
	OutInterface *OutInterface1 `json:"outInterface,omitempty"`

	// packets
	// Required: true
	// Minimum: 0
	Packets *int64 `json:"packets"`

	// protocol
	// Required: true
	Protocol *string `json:"protocol"`

	// reject with
	// Enum: [icmp_net_unreachable icmp_host_unreachable icmp_port_unreachable icmp_proto_unreachable icmp_net_prohibited icmp_host_prohibited icmp_admin_prohibited tcp_reset,]
	RejectWith string `json:"rejectWith,omitempty"`

	// source
	Source *Source `json:"source,omitempty"`

	// target
	// Required: true
	Target *string `json:"target"`
}

// Validate validates this model 59
func (m *Model59) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBytes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCategoryID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIcmpType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIcmpv6Type(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInInterface(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIpsecPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutInterface(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRejectWith(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model59) validateApplicationID(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplicationID) { // not required
		return nil
	}

	if err := validate.MinimumInt("applicationId", "body", m.ApplicationID, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Model59) validateBytes(formats strfmt.Registry) error {

	if err := validate.Required("bytes", "body", m.Bytes); err != nil {
		return err
	}

	if err := validate.MinimumInt("bytes", "body", *m.Bytes, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Model59) validateCategoryID(formats strfmt.Registry) error {
	if swag.IsZero(m.CategoryID) { // not required
		return nil
	}

	if err := validate.MinimumInt("categoryId", "body", *m.CategoryID, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Model59) validateConnectionState(formats strfmt.Registry) error {

	if err := validate.Required("connectionState", "body", m.ConnectionState); err != nil {
		return err
	}

	if err := m.ConnectionState.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("connectionState")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("connectionState")
		}
		return err
	}

	return nil
}

func (m *Model59) validateDestination(formats strfmt.Registry) error {
	if swag.IsZero(m.Destination) { // not required
		return nil
	}

	if m.Destination != nil {
		if err := m.Destination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destination")
			}
			return err
		}
	}

	return nil
}

var model59TypeIcmpTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["address-mask-reply","address-mask-request","any","communication-prohibited","destination-unreachable","echo-reply","echo-request","fragmentation-needed","host-precedence-violation","host-prohibited","host-redirect","host-unknown","host-unreachable","ip-header-bad","network-prohibited","network-redirect","network-unknown","network-unreachable","parameter-problem","port-unreachable","precedence-cutoff","protocol-unreachable","redirect","required-option-missing","router-advertisement","router-solicitation","source-quench","source-route-failed","time-exceeded","timestamp-reply","timestamp-request","TOS-host-redirect","TOS-host-unreachable","TOS-network-redirect","TOS-network-unreachable","ttl-zero-during-reassembly","ttl-zero-during-transit"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		model59TypeIcmpTypePropEnum = append(model59TypeIcmpTypePropEnum, v)
	}
}

const (

	// Model59IcmpTypeAddressDashMaskDashReply captures enum value "address-mask-reply"
	Model59IcmpTypeAddressDashMaskDashReply string = "address-mask-reply"

	// Model59IcmpTypeAddressDashMaskDashRequest captures enum value "address-mask-request"
	Model59IcmpTypeAddressDashMaskDashRequest string = "address-mask-request"

	// Model59IcmpTypeAny captures enum value "any"
	Model59IcmpTypeAny string = "any"

	// Model59IcmpTypeCommunicationDashProhibited captures enum value "communication-prohibited"
	Model59IcmpTypeCommunicationDashProhibited string = "communication-prohibited"

	// Model59IcmpTypeDestinationDashUnreachable captures enum value "destination-unreachable"
	Model59IcmpTypeDestinationDashUnreachable string = "destination-unreachable"

	// Model59IcmpTypeEchoDashReply captures enum value "echo-reply"
	Model59IcmpTypeEchoDashReply string = "echo-reply"

	// Model59IcmpTypeEchoDashRequest captures enum value "echo-request"
	Model59IcmpTypeEchoDashRequest string = "echo-request"

	// Model59IcmpTypeFragmentationDashNeeded captures enum value "fragmentation-needed"
	Model59IcmpTypeFragmentationDashNeeded string = "fragmentation-needed"

	// Model59IcmpTypeHostDashPrecedenceDashViolation captures enum value "host-precedence-violation"
	Model59IcmpTypeHostDashPrecedenceDashViolation string = "host-precedence-violation"

	// Model59IcmpTypeHostDashProhibited captures enum value "host-prohibited"
	Model59IcmpTypeHostDashProhibited string = "host-prohibited"

	// Model59IcmpTypeHostDashRedirect captures enum value "host-redirect"
	Model59IcmpTypeHostDashRedirect string = "host-redirect"

	// Model59IcmpTypeHostDashUnknown captures enum value "host-unknown"
	Model59IcmpTypeHostDashUnknown string = "host-unknown"

	// Model59IcmpTypeHostDashUnreachable captures enum value "host-unreachable"
	Model59IcmpTypeHostDashUnreachable string = "host-unreachable"

	// Model59IcmpTypeIPDashHeaderDashBad captures enum value "ip-header-bad"
	Model59IcmpTypeIPDashHeaderDashBad string = "ip-header-bad"

	// Model59IcmpTypeNetworkDashProhibited captures enum value "network-prohibited"
	Model59IcmpTypeNetworkDashProhibited string = "network-prohibited"

	// Model59IcmpTypeNetworkDashRedirect captures enum value "network-redirect"
	Model59IcmpTypeNetworkDashRedirect string = "network-redirect"

	// Model59IcmpTypeNetworkDashUnknown captures enum value "network-unknown"
	Model59IcmpTypeNetworkDashUnknown string = "network-unknown"

	// Model59IcmpTypeNetworkDashUnreachable captures enum value "network-unreachable"
	Model59IcmpTypeNetworkDashUnreachable string = "network-unreachable"

	// Model59IcmpTypeParameterDashProblem captures enum value "parameter-problem"
	Model59IcmpTypeParameterDashProblem string = "parameter-problem"

	// Model59IcmpTypePortDashUnreachable captures enum value "port-unreachable"
	Model59IcmpTypePortDashUnreachable string = "port-unreachable"

	// Model59IcmpTypePrecedenceDashCutoff captures enum value "precedence-cutoff"
	Model59IcmpTypePrecedenceDashCutoff string = "precedence-cutoff"

	// Model59IcmpTypeProtocolDashUnreachable captures enum value "protocol-unreachable"
	Model59IcmpTypeProtocolDashUnreachable string = "protocol-unreachable"

	// Model59IcmpTypeRedirect captures enum value "redirect"
	Model59IcmpTypeRedirect string = "redirect"

	// Model59IcmpTypeRequiredDashOptionDashMissing captures enum value "required-option-missing"
	Model59IcmpTypeRequiredDashOptionDashMissing string = "required-option-missing"

	// Model59IcmpTypeRouterDashAdvertisement captures enum value "router-advertisement"
	Model59IcmpTypeRouterDashAdvertisement string = "router-advertisement"

	// Model59IcmpTypeRouterDashSolicitation captures enum value "router-solicitation"
	Model59IcmpTypeRouterDashSolicitation string = "router-solicitation"

	// Model59IcmpTypeSourceDashQuench captures enum value "source-quench"
	Model59IcmpTypeSourceDashQuench string = "source-quench"

	// Model59IcmpTypeSourceDashRouteDashFailed captures enum value "source-route-failed"
	Model59IcmpTypeSourceDashRouteDashFailed string = "source-route-failed"

	// Model59IcmpTypeTimeDashExceeded captures enum value "time-exceeded"
	Model59IcmpTypeTimeDashExceeded string = "time-exceeded"

	// Model59IcmpTypeTimestampDashReply captures enum value "timestamp-reply"
	Model59IcmpTypeTimestampDashReply string = "timestamp-reply"

	// Model59IcmpTypeTimestampDashRequest captures enum value "timestamp-request"
	Model59IcmpTypeTimestampDashRequest string = "timestamp-request"

	// Model59IcmpTypeTOSDashHostDashRedirect captures enum value "TOS-host-redirect"
	Model59IcmpTypeTOSDashHostDashRedirect string = "TOS-host-redirect"

	// Model59IcmpTypeTOSDashHostDashUnreachable captures enum value "TOS-host-unreachable"
	Model59IcmpTypeTOSDashHostDashUnreachable string = "TOS-host-unreachable"

	// Model59IcmpTypeTOSDashNetworkDashRedirect captures enum value "TOS-network-redirect"
	Model59IcmpTypeTOSDashNetworkDashRedirect string = "TOS-network-redirect"

	// Model59IcmpTypeTOSDashNetworkDashUnreachable captures enum value "TOS-network-unreachable"
	Model59IcmpTypeTOSDashNetworkDashUnreachable string = "TOS-network-unreachable"

	// Model59IcmpTypeTTLDashZeroDashDuringDashReassembly captures enum value "ttl-zero-during-reassembly"
	Model59IcmpTypeTTLDashZeroDashDuringDashReassembly string = "ttl-zero-during-reassembly"

	// Model59IcmpTypeTTLDashZeroDashDuringDashTransit captures enum value "ttl-zero-during-transit"
	Model59IcmpTypeTTLDashZeroDashDuringDashTransit string = "ttl-zero-during-transit"
)

// prop value enum
func (m *Model59) validateIcmpTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, model59TypeIcmpTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Model59) validateIcmpType(formats strfmt.Registry) error {
	if swag.IsZero(m.IcmpType) { // not required
		return nil
	}

	// value enum
	if err := m.validateIcmpTypeEnum("icmpType", "body", m.IcmpType); err != nil {
		return err
	}

	return nil
}

var model59TypeIcmpv6TypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["address-unreachable","bad-header","beyond-scope","communication-prohibited","destination-unreachable","echo-reply","echo-request","failed-policy","neighbor-advertisement","neighbor-solicitation","no-route","packet-too-big","parameter-problem","port-unreachable","redirect","reject-route","router-advertisement","router-solicitation","time-exceeded","ttl-zero-during-reassembly","ttl-zero-during-transit","unknown-header-type","unknown-option"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		model59TypeIcmpv6TypePropEnum = append(model59TypeIcmpv6TypePropEnum, v)
	}
}

const (

	// Model59Icmpv6TypeAddressDashUnreachable captures enum value "address-unreachable"
	Model59Icmpv6TypeAddressDashUnreachable string = "address-unreachable"

	// Model59Icmpv6TypeBadDashHeader captures enum value "bad-header"
	Model59Icmpv6TypeBadDashHeader string = "bad-header"

	// Model59Icmpv6TypeBeyondDashScope captures enum value "beyond-scope"
	Model59Icmpv6TypeBeyondDashScope string = "beyond-scope"

	// Model59Icmpv6TypeCommunicationDashProhibited captures enum value "communication-prohibited"
	Model59Icmpv6TypeCommunicationDashProhibited string = "communication-prohibited"

	// Model59Icmpv6TypeDestinationDashUnreachable captures enum value "destination-unreachable"
	Model59Icmpv6TypeDestinationDashUnreachable string = "destination-unreachable"

	// Model59Icmpv6TypeEchoDashReply captures enum value "echo-reply"
	Model59Icmpv6TypeEchoDashReply string = "echo-reply"

	// Model59Icmpv6TypeEchoDashRequest captures enum value "echo-request"
	Model59Icmpv6TypeEchoDashRequest string = "echo-request"

	// Model59Icmpv6TypeFailedDashPolicy captures enum value "failed-policy"
	Model59Icmpv6TypeFailedDashPolicy string = "failed-policy"

	// Model59Icmpv6TypeNeighborDashAdvertisement captures enum value "neighbor-advertisement"
	Model59Icmpv6TypeNeighborDashAdvertisement string = "neighbor-advertisement"

	// Model59Icmpv6TypeNeighborDashSolicitation captures enum value "neighbor-solicitation"
	Model59Icmpv6TypeNeighborDashSolicitation string = "neighbor-solicitation"

	// Model59Icmpv6TypeNoDashRoute captures enum value "no-route"
	Model59Icmpv6TypeNoDashRoute string = "no-route"

	// Model59Icmpv6TypePacketDashTooDashBig captures enum value "packet-too-big"
	Model59Icmpv6TypePacketDashTooDashBig string = "packet-too-big"

	// Model59Icmpv6TypeParameterDashProblem captures enum value "parameter-problem"
	Model59Icmpv6TypeParameterDashProblem string = "parameter-problem"

	// Model59Icmpv6TypePortDashUnreachable captures enum value "port-unreachable"
	Model59Icmpv6TypePortDashUnreachable string = "port-unreachable"

	// Model59Icmpv6TypeRedirect captures enum value "redirect"
	Model59Icmpv6TypeRedirect string = "redirect"

	// Model59Icmpv6TypeRejectDashRoute captures enum value "reject-route"
	Model59Icmpv6TypeRejectDashRoute string = "reject-route"

	// Model59Icmpv6TypeRouterDashAdvertisement captures enum value "router-advertisement"
	Model59Icmpv6TypeRouterDashAdvertisement string = "router-advertisement"

	// Model59Icmpv6TypeRouterDashSolicitation captures enum value "router-solicitation"
	Model59Icmpv6TypeRouterDashSolicitation string = "router-solicitation"

	// Model59Icmpv6TypeTimeDashExceeded captures enum value "time-exceeded"
	Model59Icmpv6TypeTimeDashExceeded string = "time-exceeded"

	// Model59Icmpv6TypeTTLDashZeroDashDuringDashReassembly captures enum value "ttl-zero-during-reassembly"
	Model59Icmpv6TypeTTLDashZeroDashDuringDashReassembly string = "ttl-zero-during-reassembly"

	// Model59Icmpv6TypeTTLDashZeroDashDuringDashTransit captures enum value "ttl-zero-during-transit"
	Model59Icmpv6TypeTTLDashZeroDashDuringDashTransit string = "ttl-zero-during-transit"

	// Model59Icmpv6TypeUnknownDashHeaderDashType captures enum value "unknown-header-type"
	Model59Icmpv6TypeUnknownDashHeaderDashType string = "unknown-header-type"

	// Model59Icmpv6TypeUnknownDashOption captures enum value "unknown-option"
	Model59Icmpv6TypeUnknownDashOption string = "unknown-option"
)

// prop value enum
func (m *Model59) validateIcmpv6TypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, model59TypeIcmpv6TypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Model59) validateIcmpv6Type(formats strfmt.Registry) error {
	if swag.IsZero(m.Icmpv6Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateIcmpv6TypeEnum("icmpv6Type", "body", m.Icmpv6Type); err != nil {
		return err
	}

	return nil
}

func (m *Model59) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinimumInt("id", "body", *m.ID, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Model59) validateInInterface(formats strfmt.Registry) error {
	if swag.IsZero(m.InInterface) { // not required
		return nil
	}

	if m.InInterface != nil {
		if err := m.InInterface.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inInterface")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inInterface")
			}
			return err
		}
	}

	return nil
}

var model59TypeIPVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["both","v4only","v6only"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		model59TypeIPVersionPropEnum = append(model59TypeIPVersionPropEnum, v)
	}
}

const (

	// Model59IPVersionBoth captures enum value "both"
	Model59IPVersionBoth string = "both"

	// Model59IPVersionV4only captures enum value "v4only"
	Model59IPVersionV4only string = "v4only"

	// Model59IPVersionV6only captures enum value "v6only"
	Model59IPVersionV6only string = "v6only"
)

// prop value enum
func (m *Model59) validateIPVersionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, model59TypeIPVersionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Model59) validateIPVersion(formats strfmt.Registry) error {

	if err := validate.Required("ipVersion", "body", m.IPVersion); err != nil {
		return err
	}

	// value enum
	if err := m.validateIPVersionEnum("ipVersion", "body", *m.IPVersion); err != nil {
		return err
	}

	return nil
}

var model59TypeIpsecPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["match_inbound_ipsec","match_inbound_nonipsec"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		model59TypeIpsecPolicyPropEnum = append(model59TypeIpsecPolicyPropEnum, v)
	}
}

const (

	// Model59IpsecPolicyMatchInboundIpsec captures enum value "match_inbound_ipsec"
	Model59IpsecPolicyMatchInboundIpsec string = "match_inbound_ipsec"

	// Model59IpsecPolicyMatchInboundNonipsec captures enum value "match_inbound_nonipsec"
	Model59IpsecPolicyMatchInboundNonipsec string = "match_inbound_nonipsec"
)

// prop value enum
func (m *Model59) validateIpsecPolicyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, model59TypeIpsecPolicyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Model59) validateIpsecPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.IpsecPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateIpsecPolicyEnum("ipsecPolicy", "body", m.IpsecPolicy); err != nil {
		return err
	}

	return nil
}

func (m *Model59) validateOutInterface(formats strfmt.Registry) error {
	if swag.IsZero(m.OutInterface) { // not required
		return nil
	}

	if m.OutInterface != nil {
		if err := m.OutInterface.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outInterface")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("outInterface")
			}
			return err
		}
	}

	return nil
}

func (m *Model59) validatePackets(formats strfmt.Registry) error {

	if err := validate.Required("packets", "body", m.Packets); err != nil {
		return err
	}

	if err := validate.MinimumInt("packets", "body", *m.Packets, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Model59) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", m.Protocol); err != nil {
		return err
	}

	return nil
}

var model59TypeRejectWithPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["icmp_net_unreachable","icmp_host_unreachable","icmp_port_unreachable","icmp_proto_unreachable","icmp_net_prohibited","icmp_host_prohibited","icmp_admin_prohibited","tcp_reset,"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		model59TypeRejectWithPropEnum = append(model59TypeRejectWithPropEnum, v)
	}
}

const (

	// Model59RejectWithIcmpNetUnreachable captures enum value "icmp_net_unreachable"
	Model59RejectWithIcmpNetUnreachable string = "icmp_net_unreachable"

	// Model59RejectWithIcmpHostUnreachable captures enum value "icmp_host_unreachable"
	Model59RejectWithIcmpHostUnreachable string = "icmp_host_unreachable"

	// Model59RejectWithIcmpPortUnreachable captures enum value "icmp_port_unreachable"
	Model59RejectWithIcmpPortUnreachable string = "icmp_port_unreachable"

	// Model59RejectWithIcmpProtoUnreachable captures enum value "icmp_proto_unreachable"
	Model59RejectWithIcmpProtoUnreachable string = "icmp_proto_unreachable"

	// Model59RejectWithIcmpNetProhibited captures enum value "icmp_net_prohibited"
	Model59RejectWithIcmpNetProhibited string = "icmp_net_prohibited"

	// Model59RejectWithIcmpHostProhibited captures enum value "icmp_host_prohibited"
	Model59RejectWithIcmpHostProhibited string = "icmp_host_prohibited"

	// Model59RejectWithIcmpAdminProhibited captures enum value "icmp_admin_prohibited"
	Model59RejectWithIcmpAdminProhibited string = "icmp_admin_prohibited"

	// Model59RejectWithTCPReset captures enum value "tcp_reset,"
	Model59RejectWithTCPReset string = "tcp_reset,"
)

// prop value enum
func (m *Model59) validateRejectWithEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, model59TypeRejectWithPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Model59) validateRejectWith(formats strfmt.Registry) error {
	if swag.IsZero(m.RejectWith) { // not required
		return nil
	}

	// value enum
	if err := m.validateRejectWithEnum("rejectWith", "body", m.RejectWith); err != nil {
		return err
	}

	return nil
}

func (m *Model59) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

func (m *Model59) validateTarget(formats strfmt.Registry) error {

	if err := validate.Required("target", "body", m.Target); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this model 59 based on the context it is used
func (m *Model59) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnectionState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDestination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInInterface(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOutInterface(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model59) contextValidateConnectionState(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ConnectionState.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("connectionState")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("connectionState")
		}
		return err
	}

	return nil
}

func (m *Model59) contextValidateDestination(ctx context.Context, formats strfmt.Registry) error {

	if m.Destination != nil {
		if err := m.Destination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destination")
			}
			return err
		}
	}

	return nil
}

func (m *Model59) contextValidateInInterface(ctx context.Context, formats strfmt.Registry) error {

	if m.InInterface != nil {
		if err := m.InInterface.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inInterface")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inInterface")
			}
			return err
		}
	}

	return nil
}

func (m *Model59) contextValidateOutInterface(ctx context.Context, formats strfmt.Registry) error {

	if m.OutInterface != nil {
		if err := m.OutInterface.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outInterface")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("outInterface")
			}
			return err
		}
	}

	return nil
}

func (m *Model59) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {
		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Model59) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model59) UnmarshalBinary(b []byte) error {
	var res Model59
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
