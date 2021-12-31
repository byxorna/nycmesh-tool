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

// DeviceListCredentials device list credentials
//
// swagger:model Device list credentials
type DeviceListCredentials struct {

	// credentials
	// Required: true
	Credentials Credentials `json:"credentials"`

	// is passphrase missing
	IsPassphraseMissing bool `json:"isPassphraseMissing,omitempty"`
}

// Validate validates this device list credentials
func (m *DeviceListCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceListCredentials) validateCredentials(formats strfmt.Registry) error {

	if err := validate.Required("credentials", "body", m.Credentials); err != nil {
		return err
	}

	if err := m.Credentials.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("credentials")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("credentials")
		}
		return err
	}

	return nil
}

// ContextValidate validate this device list credentials based on the context it is used
func (m *DeviceListCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceListCredentials) contextValidateCredentials(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Credentials.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("credentials")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("credentials")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceListCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceListCredentials) UnmarshalBinary(b []byte) error {
	var res DeviceListCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
