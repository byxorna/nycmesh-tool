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

// Vlans1 vlans 1
//
// swagger:model vlans 1
type Vlans1 struct {

	// interface
	Interface *Interface1 `json:"interface,omitempty"`

	// vlan Id
	VlanID VlanID `json:"vlanId,omitempty"`
}

// Validate validates this vlans 1
func (m *Vlans1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInterface(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Vlans1) validateInterface(formats strfmt.Registry) error {
	if swag.IsZero(m.Interface) { // not required
		return nil
	}

	if m.Interface != nil {
		if err := m.Interface.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("interface")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vlans 1 based on the context it is used
func (m *Vlans1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInterface(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Vlans1) contextValidateInterface(ctx context.Context, formats strfmt.Registry) error {

	if m.Interface != nil {
		if err := m.Interface.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("interface")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Vlans1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Vlans1) UnmarshalBinary(b []byte) error {
	var res Vlans1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
