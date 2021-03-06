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

// DefaultIPQueue default Ip queue
//
// swagger:model defaultIpQueue
type DefaultIPQueue struct {

	// download speed
	// Minimum: 0
	DownloadSpeed *int64 `json:"downloadSpeed,omitempty"`

	// Enable configuration for QoS queue for IPs without any specific queue. The main purpose of this configuration is to block new devices without any active queue. If null, traffic not belonging to any queue is not limited.
	// Example: true
	Enabled *bool `json:"enabled,omitempty"`

	// upload speed
	// Minimum: 0
	UploadSpeed *int64 `json:"uploadSpeed,omitempty"`
}

// Validate validates this default Ip queue
func (m *DefaultIPQueue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDownloadSpeed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadSpeed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DefaultIPQueue) validateDownloadSpeed(formats strfmt.Registry) error {
	if swag.IsZero(m.DownloadSpeed) { // not required
		return nil
	}

	if err := validate.MinimumInt("downloadSpeed", "body", *m.DownloadSpeed, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *DefaultIPQueue) validateUploadSpeed(formats strfmt.Registry) error {
	if swag.IsZero(m.UploadSpeed) { // not required
		return nil
	}

	if err := validate.MinimumInt("uploadSpeed", "body", *m.UploadSpeed, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this default Ip queue based on context it is used
func (m *DefaultIPQueue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DefaultIPQueue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DefaultIPQueue) UnmarshalBinary(b []byte) error {
	var res DefaultIPQueue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
