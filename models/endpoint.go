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

// Endpoint endpoint
//
// swagger:model endpoint
type Endpoint struct {

	// Number of client sites with read access.
	// Example: 3
	// Required: true
	All *float64 `json:"all"`

	// Number of client sites with read-only access.
	// Example: 1
	// Required: true
	ReadOnly *float64 `json:"readOnly"`

	// Number of client sites with read-write access.
	// Example: 2
	// Required: true
	ReadWrite *float64 `json:"readWrite"`
}

// Validate validates this endpoint
func (m *Endpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAll(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReadOnly(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReadWrite(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Endpoint) validateAll(formats strfmt.Registry) error {

	if err := validate.Required("all", "body", m.All); err != nil {
		return err
	}

	return nil
}

func (m *Endpoint) validateReadOnly(formats strfmt.Registry) error {

	if err := validate.Required("readOnly", "body", m.ReadOnly); err != nil {
		return err
	}

	return nil
}

func (m *Endpoint) validateReadWrite(formats strfmt.Registry) error {

	if err := validate.Required("readWrite", "body", m.ReadWrite); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this endpoint based on context it is used
func (m *Endpoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Endpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Endpoint) UnmarshalBinary(b []byte) error {
	var res Endpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
