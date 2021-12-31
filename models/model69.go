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

// Model69 model 69
//
// swagger:model Model 69
type Model69 struct {

	// device Id
	// Example: f7ac9cad-ea28-4390-93c8-7add010e8ee3
	// Required: true
	DeviceID *string `json:"deviceId"`

	// For accessPoint (AC series) with M stations.
	FirmwareForMStationsVersion string `json:"firmwareForMStationsVersion,omitempty"`

	// firmware version
	// Required: true
	FirmwareVersion *string `json:"firmwareVersion"`
}

// Validate validates this model 69
func (m *Model69) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirmwareVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model69) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("deviceId", "body", m.DeviceID); err != nil {
		return err
	}

	return nil
}

func (m *Model69) validateFirmwareVersion(formats strfmt.Registry) error {

	if err := validate.Required("firmwareVersion", "body", m.FirmwareVersion); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this model 69 based on context it is used
func (m *Model69) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Model69) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model69) UnmarshalBinary(b []byte) error {
	var res Model69
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
