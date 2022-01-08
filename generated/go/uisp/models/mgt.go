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

// Mgt mgt
//
// swagger:model mgt
type Mgt struct {

	// cidr
	// Required: true
	Cidr *string `json:"cidr"`

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// password
	// Required: true
	Password *string `json:"password"`

	// proto
	// Required: true
	// Enum: [static dhcp pppoe]
	Proto *string `json:"proto"`

	// service
	// Required: true
	Service *string `json:"service"`

	// username
	// Required: true
	Username *string `json:"username"`

	// vlan Id
	// Required: true
	// Maximum: 4096
	// Minimum: 1
	VlanID *int64 `json:"vlanId"`
}

// Validate validates this mgt
func (m *Mgt) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProto(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Mgt) validateCidr(formats strfmt.Registry) error {

	if err := validate.Required("cidr", "body", m.Cidr); err != nil {
		return err
	}

	return nil
}

func (m *Mgt) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *Mgt) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

var mgtTypeProtoPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["static","dhcp","pppoe"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mgtTypeProtoPropEnum = append(mgtTypeProtoPropEnum, v)
	}
}

const (

	// MgtProtoStatic captures enum value "static"
	MgtProtoStatic string = "static"

	// MgtProtoDhcp captures enum value "dhcp"
	MgtProtoDhcp string = "dhcp"

	// MgtProtoPppoe captures enum value "pppoe"
	MgtProtoPppoe string = "pppoe"
)

// prop value enum
func (m *Mgt) validateProtoEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mgtTypeProtoPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Mgt) validateProto(formats strfmt.Registry) error {

	if err := validate.Required("proto", "body", m.Proto); err != nil {
		return err
	}

	// value enum
	if err := m.validateProtoEnum("proto", "body", *m.Proto); err != nil {
		return err
	}

	return nil
}

func (m *Mgt) validateService(formats strfmt.Registry) error {

	if err := validate.Required("service", "body", m.Service); err != nil {
		return err
	}

	return nil
}

func (m *Mgt) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

func (m *Mgt) validateVlanID(formats strfmt.Registry) error {

	if err := validate.Required("vlanId", "body", m.VlanID); err != nil {
		return err
	}

	if err := validate.MinimumInt("vlanId", "body", *m.VlanID, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vlanId", "body", *m.VlanID, 4096, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this mgt based on context it is used
func (m *Mgt) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Mgt) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Mgt) UnmarshalBinary(b []byte) error {
	var res Mgt
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}