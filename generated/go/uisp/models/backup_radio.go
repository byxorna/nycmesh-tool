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

// BackupRadio Backup radio basic info
//
// swagger:model backupRadio
type BackupRadio struct {

	// frequency
	// Example: 5400
	// Required: true
	// Minimum: 0
	Frequency *float64 `json:"frequency"`
}

// Validate validates this backup radio
func (m *BackupRadio) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrequency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupRadio) validateFrequency(formats strfmt.Registry) error {

	if err := validate.Required("frequency", "body", m.Frequency); err != nil {
		return err
	}

	if err := validate.Minimum("frequency", "body", *m.Frequency, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this backup radio based on context it is used
func (m *BackupRadio) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BackupRadio) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupRadio) UnmarshalBinary(b []byte) error {
	var res BackupRadio
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
