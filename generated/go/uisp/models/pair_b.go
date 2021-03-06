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

// PairB pair b
//
// swagger:model pairB
type PairB struct {

	// length from
	// Required: true
	LengthFrom *float64 `json:"lengthFrom"`

	// length to
	// Required: true
	LengthTo *float64 `json:"lengthTo"`

	// state
	// Required: true
	State *string `json:"state"`
}

// Validate validates this pair b
func (m *PairB) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLengthFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLengthTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PairB) validateLengthFrom(formats strfmt.Registry) error {

	if err := validate.Required("lengthFrom", "body", m.LengthFrom); err != nil {
		return err
	}

	return nil
}

func (m *PairB) validateLengthTo(formats strfmt.Registry) error {

	if err := validate.Required("lengthTo", "body", m.LengthTo); err != nil {
		return err
	}

	return nil
}

func (m *PairB) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this pair b based on context it is used
func (m *PairB) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PairB) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PairB) UnmarshalBinary(b []byte) error {
	var res PairB
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
