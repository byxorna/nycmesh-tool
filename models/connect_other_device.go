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

// ConnectOtherDevice connect other device
//
// swagger:model ConnectOtherDevice
type ConnectOtherDevice struct {

	// enable ping
	// Required: true
	EnablePing bool `json:"enablePing"`

	// hostname
	// Required: true
	Hostname *string `json:"hostname"`

	// ip
	// Required: true
	IP *string `json:"ip"`

	// role
	// Required: true
	// Enum: [router switch gpon ap station other ups server wireless convertor gateway]
	Role *string `json:"role"`

	// snmp community
	// Required: true
	SnmpCommunity *string `json:"snmpCommunity"`
}

// Validate validates this connect other device
func (m *ConnectOtherDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnablePing(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnmpCommunity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConnectOtherDevice) validateEnablePing(formats strfmt.Registry) error {

	if err := validate.Required("enablePing", "body", bool(m.EnablePing)); err != nil {
		return err
	}

	return nil
}

func (m *ConnectOtherDevice) validateHostname(formats strfmt.Registry) error {

	if err := validate.Required("hostname", "body", m.Hostname); err != nil {
		return err
	}

	return nil
}

func (m *ConnectOtherDevice) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

var connectOtherDeviceTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["router","switch","gpon","ap","station","other","ups","server","wireless","convertor","gateway"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		connectOtherDeviceTypeRolePropEnum = append(connectOtherDeviceTypeRolePropEnum, v)
	}
}

const (

	// ConnectOtherDeviceRoleRouter captures enum value "router"
	ConnectOtherDeviceRoleRouter string = "router"

	// ConnectOtherDeviceRoleSwitch captures enum value "switch"
	ConnectOtherDeviceRoleSwitch string = "switch"

	// ConnectOtherDeviceRoleGpon captures enum value "gpon"
	ConnectOtherDeviceRoleGpon string = "gpon"

	// ConnectOtherDeviceRoleAp captures enum value "ap"
	ConnectOtherDeviceRoleAp string = "ap"

	// ConnectOtherDeviceRoleStation captures enum value "station"
	ConnectOtherDeviceRoleStation string = "station"

	// ConnectOtherDeviceRoleOther captures enum value "other"
	ConnectOtherDeviceRoleOther string = "other"

	// ConnectOtherDeviceRoleUps captures enum value "ups"
	ConnectOtherDeviceRoleUps string = "ups"

	// ConnectOtherDeviceRoleServer captures enum value "server"
	ConnectOtherDeviceRoleServer string = "server"

	// ConnectOtherDeviceRoleWireless captures enum value "wireless"
	ConnectOtherDeviceRoleWireless string = "wireless"

	// ConnectOtherDeviceRoleConvertor captures enum value "convertor"
	ConnectOtherDeviceRoleConvertor string = "convertor"

	// ConnectOtherDeviceRoleGateway captures enum value "gateway"
	ConnectOtherDeviceRoleGateway string = "gateway"
)

// prop value enum
func (m *ConnectOtherDevice) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, connectOtherDeviceTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConnectOtherDevice) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", *m.Role); err != nil {
		return err
	}

	return nil
}

func (m *ConnectOtherDevice) validateSnmpCommunity(formats strfmt.Registry) error {

	if err := validate.Required("snmpCommunity", "body", m.SnmpCommunity); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this connect other device based on context it is used
func (m *ConnectOtherDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConnectOtherDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectOtherDevice) UnmarshalBinary(b []byte) error {
	var res ConnectOtherDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}