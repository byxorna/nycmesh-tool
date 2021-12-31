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

// Model103 model 103
//
// swagger:model Model 103
type Model103 struct {

	// auth
	// Enum: [md5 plaintext-password off]
	Auth string `json:"auth,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// networks
	Networks Networks `json:"networks,omitempty"`

	// type
	// Enum: [normal nssa stub]
	Type string `json:"type,omitempty"`
}

// Validate validates this model 103
func (m *Model103) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var model103TypeAuthPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["md5","plaintext-password","off"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		model103TypeAuthPropEnum = append(model103TypeAuthPropEnum, v)
	}
}

const (

	// Model103AuthMd5 captures enum value "md5"
	Model103AuthMd5 string = "md5"

	// Model103AuthPlaintextDashPassword captures enum value "plaintext-password"
	Model103AuthPlaintextDashPassword string = "plaintext-password"

	// Model103AuthOff captures enum value "off"
	Model103AuthOff string = "off"
)

// prop value enum
func (m *Model103) validateAuthEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, model103TypeAuthPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Model103) validateAuth(formats strfmt.Registry) error {
	if swag.IsZero(m.Auth) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthEnum("auth", "body", m.Auth); err != nil {
		return err
	}

	return nil
}

func (m *Model103) validateNetworks(formats strfmt.Registry) error {
	if swag.IsZero(m.Networks) { // not required
		return nil
	}

	if err := m.Networks.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("networks")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("networks")
		}
		return err
	}

	return nil
}

var model103TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["normal","nssa","stub"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		model103TypeTypePropEnum = append(model103TypeTypePropEnum, v)
	}
}

const (

	// Model103TypeNormal captures enum value "normal"
	Model103TypeNormal string = "normal"

	// Model103TypeNssa captures enum value "nssa"
	Model103TypeNssa string = "nssa"

	// Model103TypeStub captures enum value "stub"
	Model103TypeStub string = "stub"
)

// prop value enum
func (m *Model103) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, model103TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Model103) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this model 103 based on the context it is used
func (m *Model103) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model103) contextValidateNetworks(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Networks.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("networks")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("networks")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Model103) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model103) UnmarshalBinary(b []byte) error {
	var res Model103
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
