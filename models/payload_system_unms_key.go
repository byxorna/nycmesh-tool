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

// PayloadSystemUnmsKey payload system unms key
//
// swagger:model PayloadSystemUnmsKey
type PayloadSystemUnmsKey struct {

	// device ids
	DeviceIds DeviceIds `json:"deviceIds,omitempty"`

	// unms key
	// Required: true
	UnmsKey *string `json:"unmsKey"`
}

// Validate validates this payload system unms key
func (m *PayloadSystemUnmsKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnmsKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PayloadSystemUnmsKey) validateDeviceIds(formats strfmt.Registry) error {
	if swag.IsZero(m.DeviceIds) { // not required
		return nil
	}

	if err := m.DeviceIds.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deviceIds")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("deviceIds")
		}
		return err
	}

	return nil
}

func (m *PayloadSystemUnmsKey) validateUnmsKey(formats strfmt.Registry) error {

	if err := validate.Required("unmsKey", "body", m.UnmsKey); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this payload system unms key based on the context it is used
func (m *PayloadSystemUnmsKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeviceIds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PayloadSystemUnmsKey) contextValidateDeviceIds(ctx context.Context, formats strfmt.Registry) error {

	if err := m.DeviceIds.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deviceIds")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("deviceIds")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PayloadSystemUnmsKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PayloadSystemUnmsKey) UnmarshalBinary(b []byte) error {
	var res PayloadSystemUnmsKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
