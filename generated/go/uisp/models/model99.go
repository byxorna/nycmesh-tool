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

// Model99 model 99
//
// swagger:model Model 99
type Model99 struct {

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

// Validate validates this model 99
func (m *Model99) Validate(formats strfmt.Registry) error {
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

var model99TypeAuthPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["md5","plaintext-password","off"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		model99TypeAuthPropEnum = append(model99TypeAuthPropEnum, v)
	}
}

const (

	// Model99AuthMd5 captures enum value "md5"
	Model99AuthMd5 string = "md5"

	// Model99AuthPlaintextDashPassword captures enum value "plaintext-password"
	Model99AuthPlaintextDashPassword string = "plaintext-password"

	// Model99AuthOff captures enum value "off"
	Model99AuthOff string = "off"
)

// prop value enum
func (m *Model99) validateAuthEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, model99TypeAuthPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Model99) validateAuth(formats strfmt.Registry) error {
	if swag.IsZero(m.Auth) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthEnum("auth", "body", m.Auth); err != nil {
		return err
	}

	return nil
}

func (m *Model99) validateNetworks(formats strfmt.Registry) error {
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

var model99TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["normal","nssa","stub"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		model99TypeTypePropEnum = append(model99TypeTypePropEnum, v)
	}
}

const (

	// Model99TypeNormal captures enum value "normal"
	Model99TypeNormal string = "normal"

	// Model99TypeNssa captures enum value "nssa"
	Model99TypeNssa string = "nssa"

	// Model99TypeStub captures enum value "stub"
	Model99TypeStub string = "stub"
)

// prop value enum
func (m *Model99) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, model99TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Model99) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this model 99 based on the context it is used
func (m *Model99) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model99) contextValidateNetworks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Model99) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model99) UnmarshalBinary(b []byte) error {
	var res Model99
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
