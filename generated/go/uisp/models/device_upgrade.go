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

// DeviceUpgrade device upgrade
//
// swagger:model DeviceUpgrade
type DeviceUpgrade struct {

	// firmware
	// Required: true
	Firmware *Firmware `json:"firmware"`

	// progress
	// Required: true
	Progress *float64 `json:"progress"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this device upgrade
func (m *DeviceUpgrade) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirmware(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProgress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceUpgrade) validateFirmware(formats strfmt.Registry) error {

	if err := validate.Required("firmware", "body", m.Firmware); err != nil {
		return err
	}

	if m.Firmware != nil {
		if err := m.Firmware.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firmware")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("firmware")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceUpgrade) validateProgress(formats strfmt.Registry) error {

	if err := validate.Required("progress", "body", m.Progress); err != nil {
		return err
	}

	return nil
}

func (m *DeviceUpgrade) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this device upgrade based on the context it is used
func (m *DeviceUpgrade) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFirmware(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceUpgrade) contextValidateFirmware(ctx context.Context, formats strfmt.Registry) error {

	if m.Firmware != nil {
		if err := m.Firmware.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firmware")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("firmware")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceUpgrade) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceUpgrade) UnmarshalBinary(b []byte) error {
	var res DeviceUpgrade
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}