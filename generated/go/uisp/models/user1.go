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

// User1 user 1
//
// swagger:model user 1
type User1 struct {

	// Number of users using this access group.
	// Example: 1
	// Required: true
	All *float64 `json:"all"`
}

// Validate validates this user 1
func (m *User1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAll(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User1) validateAll(formats strfmt.Registry) error {

	if err := validate.Required("all", "body", m.All); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user 1 based on context it is used
func (m *User1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *User1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User1) UnmarshalBinary(b []byte) error {
	var res User1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
