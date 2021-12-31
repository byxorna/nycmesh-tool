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

// Uisps uisps
//
// swagger:model uisps
type Uisps struct {

	// vlans
	Vlans *Vlans `json:"vlans,omitempty"`
}

// Validate validates this uisps
func (m *Uisps) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVlans(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Uisps) validateVlans(formats strfmt.Registry) error {
	if swag.IsZero(m.Vlans) { // not required
		return nil
	}

	if m.Vlans != nil {
		if err := m.Vlans.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlans")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlans")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this uisps based on the context it is used
func (m *Uisps) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVlans(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Uisps) contextValidateVlans(ctx context.Context, formats strfmt.Registry) error {

	if m.Vlans != nil {
		if err := m.Vlans.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlans")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlans")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Uisps) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Uisps) UnmarshalBinary(b []byte) error {
	var res Uisps
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}