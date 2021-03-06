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

// Model41 model 41
//
// swagger:model Model 41
type Model41 struct {

	// admin
	Admin *Admin `json:"admin,omitempty"`
}

// Validate validates this model 41
func (m *Model41) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdmin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model41) validateAdmin(formats strfmt.Registry) error {
	if swag.IsZero(m.Admin) { // not required
		return nil
	}

	if m.Admin != nil {
		if err := m.Admin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("admin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("admin")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this model 41 based on the context it is used
func (m *Model41) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdmin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model41) contextValidateAdmin(ctx context.Context, formats strfmt.Registry) error {

	if m.Admin != nil {
		if err := m.Admin.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("admin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("admin")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Model41) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model41) UnmarshalBinary(b []byte) error {
	var res Model41
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
