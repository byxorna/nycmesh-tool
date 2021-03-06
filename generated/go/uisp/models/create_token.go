// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateToken create token
//
// swagger:model CreateToken
type CreateToken struct {

	// Any token metadata.
	// Max Length: 256
	Meta string `json:"meta,omitempty"`

	// Human readable name of the token.
	// Example: Mobile App
	// Max Length: 20
	Name string `json:"name,omitempty"`
}

// Validate validates this create token
func (m *CreateToken) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateToken) validateMeta(formats strfmt.Registry) error {
	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if err := validate.MaxLength("meta", "body", m.Meta, 256); err != nil {
		return err
	}

	return nil
}

func (m *CreateToken) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", m.Name, 20); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create token based on context it is used
func (m *CreateToken) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateToken) UnmarshalBinary(b []byte) error {
	var res CreateToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
