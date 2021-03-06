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

// Model42 model 42
//
// swagger:model Model 42
type Model42 struct {

	// addresses
	Addresses Addresses2 `json:"addresses,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// mac
	Mac string `json:"mac,omitempty"`

	// name
	// Required: true
	// Max Length: 256
	// Min Length: 1
	Name *string `json:"name"`

	// position
	// Required: true
	Position *int64 `json:"position"`

	// type
	// Enum: [eth sfp+ wlan]
	Type string `json:"type,omitempty"`
}

// Validate validates this model 42
func (m *Model42) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
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

func (m *Model42) validateAddresses(formats strfmt.Registry) error {
	if swag.IsZero(m.Addresses) { // not required
		return nil
	}

	if err := m.Addresses.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("addresses")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("addresses")
		}
		return err
	}

	return nil
}

func (m *Model42) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Model42) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 256); err != nil {
		return err
	}

	return nil
}

func (m *Model42) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("position", "body", m.Position); err != nil {
		return err
	}

	return nil
}

var model42TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["eth","sfp+","wlan"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		model42TypeTypePropEnum = append(model42TypeTypePropEnum, v)
	}
}

const (

	// Model42TypeEth captures enum value "eth"
	Model42TypeEth string = "eth"

	// Model42TypeSfpPlus captures enum value "sfp+"
	Model42TypeSfpPlus string = "sfp+"

	// Model42TypeWlan captures enum value "wlan"
	Model42TypeWlan string = "wlan"
)

// prop value enum
func (m *Model42) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, model42TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Model42) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this model 42 based on the context it is used
func (m *Model42) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model42) contextValidateAddresses(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Addresses.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("addresses")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("addresses")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Model42) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model42) UnmarshalBinary(b []byte) error {
	var res Model42
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
