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

// Login1 login 1
//
// swagger:model login 1
type Login1 struct {

	// Admin's password
	// Example: secredPa$$word
	// Max Length: 64
	// Min Length: 4
	Password string `json:"password,omitempty"`

	// Admin's username
	// Example: admin
	// Required: true
	// Max Length: 320
	// Min Length: 1
	// Pattern: ^[a-zA-Z0-9_]*$
	Username *string `json:"username"`
}

// Validate validates this login 1
func (m *Login1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Login1) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := validate.MinLength("password", "body", m.Password, 4); err != nil {
		return err
	}

	if err := validate.MaxLength("password", "body", m.Password, 64); err != nil {
		return err
	}

	return nil
}

func (m *Login1) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	if err := validate.MinLength("username", "body", *m.Username, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("username", "body", *m.Username, 320); err != nil {
		return err
	}

	if err := validate.Pattern("username", "body", *m.Username, `^[a-zA-Z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this login 1 based on context it is used
func (m *Login1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Login1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Login1) UnmarshalBinary(b []byte) error {
	var res Login1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
