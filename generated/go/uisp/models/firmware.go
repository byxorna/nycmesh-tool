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

// Firmware firmware
//
// swagger:model firmware
type Firmware struct {

	// major
	// Example: 1
	// Required: true
	Major *float64 `json:"major"`

	// minor
	// Example: 10
	// Required: true
	Minor *float64 `json:"minor"`

	// order
	// Example: 65546.8.0
	Order string `json:"order,omitempty"`

	// patch
	// Example: 8
	// Required: true
	Patch *float64 `json:"patch"`

	// upgrade recommended to version
	UpgradeRecommendedToVersion string `json:"upgradeRecommendedToVersion,omitempty"`
}

// Validate validates this firmware
func (m *Firmware) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMajor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePatch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Firmware) validateMajor(formats strfmt.Registry) error {

	if err := validate.Required("major", "body", m.Major); err != nil {
		return err
	}

	return nil
}

func (m *Firmware) validateMinor(formats strfmt.Registry) error {

	if err := validate.Required("minor", "body", m.Minor); err != nil {
		return err
	}

	return nil
}

func (m *Firmware) validatePatch(formats strfmt.Registry) error {

	if err := validate.Required("patch", "body", m.Patch); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this firmware based on context it is used
func (m *Firmware) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Firmware) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Firmware) UnmarshalBinary(b []byte) error {
	var res Firmware
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
