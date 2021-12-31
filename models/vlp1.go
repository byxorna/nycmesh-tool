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

// Vlp1 vlp1
//
// swagger:model vlp1
type Vlp1 struct {

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// ports
	Ports Ports2 `json:"ports,omitempty"`

	// tagged
	// Required: true
	Tagged *bool `json:"tagged"`

	// vlan Id
	// Required: true
	// Maximum: 4096
	// Minimum: 1
	VlanID *int64 `json:"vlanId"`
}

// Validate validates this vlp1
func (m *Vlp1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTagged(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Vlp1) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *Vlp1) validatePorts(formats strfmt.Registry) error {
	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	if err := m.Ports.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ports")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ports")
		}
		return err
	}

	return nil
}

func (m *Vlp1) validateTagged(formats strfmt.Registry) error {

	if err := validate.Required("tagged", "body", m.Tagged); err != nil {
		return err
	}

	return nil
}

func (m *Vlp1) validateVlanID(formats strfmt.Registry) error {

	if err := validate.Required("vlanId", "body", m.VlanID); err != nil {
		return err
	}

	if err := validate.MinimumInt("vlanId", "body", *m.VlanID, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vlanId", "body", *m.VlanID, 4096, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this vlp1 based on the context it is used
func (m *Vlp1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Vlp1) contextValidatePorts(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Ports.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ports")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ports")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Vlp1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Vlp1) UnmarshalBinary(b []byte) error {
	var res Vlp1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
