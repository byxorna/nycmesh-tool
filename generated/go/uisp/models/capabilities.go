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

// Capabilities capabilities
//
// swagger:model capabilities
type Capabilities struct {

	// load balance values
	LoadBalanceValues LoadBalanceValues `json:"loadBalanceValues,omitempty"`

	// support auto edge
	SupportAutoEdge bool `json:"supportAutoEdge,omitempty"`

	// support cable test
	SupportCableTest bool `json:"supportCableTest,omitempty"`

	// support reset
	SupportReset bool `json:"supportReset,omitempty"`
}

// Validate validates this capabilities
func (m *Capabilities) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLoadBalanceValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Capabilities) validateLoadBalanceValues(formats strfmt.Registry) error {
	if swag.IsZero(m.LoadBalanceValues) { // not required
		return nil
	}

	if err := m.LoadBalanceValues.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("loadBalanceValues")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("loadBalanceValues")
		}
		return err
	}

	return nil
}

// ContextValidate validate this capabilities based on the context it is used
func (m *Capabilities) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLoadBalanceValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Capabilities) contextValidateLoadBalanceValues(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LoadBalanceValues.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("loadBalanceValues")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("loadBalanceValues")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Capabilities) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Capabilities) UnmarshalBinary(b []byte) error {
	var res Capabilities
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
