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

// Latest latest
//
// swagger:model latest
type Latest struct {

	// major
	// Required: true
	Major *float64 `json:"major"`

	// minor
	// Required: true
	Minor *float64 `json:"minor"`

	// order
	// Required: true
	Order *string `json:"order"`

	// patch
	// Required: true
	Patch *float64 `json:"patch"`

	// prerelease
	Prerelease Prerelease `json:"prerelease,omitempty"`
}

// Validate validates this latest
func (m *Latest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMajor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePatch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrerelease(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Latest) validateMajor(formats strfmt.Registry) error {

	if err := validate.Required("major", "body", m.Major); err != nil {
		return err
	}

	return nil
}

func (m *Latest) validateMinor(formats strfmt.Registry) error {

	if err := validate.Required("minor", "body", m.Minor); err != nil {
		return err
	}

	return nil
}

func (m *Latest) validateOrder(formats strfmt.Registry) error {

	if err := validate.Required("order", "body", m.Order); err != nil {
		return err
	}

	return nil
}

func (m *Latest) validatePatch(formats strfmt.Registry) error {

	if err := validate.Required("patch", "body", m.Patch); err != nil {
		return err
	}

	return nil
}

func (m *Latest) validatePrerelease(formats strfmt.Registry) error {
	if swag.IsZero(m.Prerelease) { // not required
		return nil
	}

	if err := m.Prerelease.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("prerelease")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("prerelease")
		}
		return err
	}

	return nil
}

// ContextValidate validate this latest based on the context it is used
func (m *Latest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrerelease(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Latest) contextValidatePrerelease(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Prerelease.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("prerelease")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("prerelease")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Latest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Latest) UnmarshalBinary(b []byte) error {
	var res Latest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
