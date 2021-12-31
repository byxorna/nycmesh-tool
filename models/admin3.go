// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Admin3 admin 3
//
// swagger:model admin 3
type Admin3 struct {

	// login
	Login *Login5 `json:"login,omitempty"`
}

// Validate validates this admin 3
func (m *Admin3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Admin3) validateLogin(formats strfmt.Registry) error {
	if swag.IsZero(m.Login) { // not required
		return nil
	}

	if m.Login != nil {
		if err := m.Login.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("login")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("login")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this admin 3 based on the context it is used
func (m *Admin3) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Admin3) contextValidateLogin(ctx context.Context, formats strfmt.Registry) error {

	if m.Login != nil {
		if err := m.Login.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("login")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("login")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Admin3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Admin3) UnmarshalBinary(b []byte) error {
	var res Admin3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
