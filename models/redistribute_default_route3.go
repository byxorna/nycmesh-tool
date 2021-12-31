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

// RedistributeDefaultRoute3 redistribute default route 3
//
// swagger:model redistributeDefaultRoute 3
type RedistributeDefaultRoute3 struct {

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`
}

// Validate validates this redistribute default route 3
func (m *RedistributeDefaultRoute3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RedistributeDefaultRoute3) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this redistribute default route 3 based on context it is used
func (m *RedistributeDefaultRoute3) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RedistributeDefaultRoute3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RedistributeDefaultRoute3) UnmarshalBinary(b []byte) error {
	var res RedistributeDefaultRoute3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
