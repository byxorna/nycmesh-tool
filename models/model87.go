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

// Model87 model 87
//
// swagger:model Model 87
type Model87 struct {

	// note
	// Example: Everything went smoothly.
	// Max Length: 500
	Note string `json:"note,omitempty"`

	// stars
	// Example: 5
	// Required: true
	// Maximum: 5
	// Minimum: 0
	Stars *float64 `json:"stars"`
}

// Validate validates this model 87
func (m *Model87) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStars(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model87) validateNote(formats strfmt.Registry) error {
	if swag.IsZero(m.Note) { // not required
		return nil
	}

	if err := validate.MaxLength("note", "body", m.Note, 500); err != nil {
		return err
	}

	return nil
}

func (m *Model87) validateStars(formats strfmt.Registry) error {

	if err := validate.Required("stars", "body", m.Stars); err != nil {
		return err
	}

	if err := validate.Minimum("stars", "body", *m.Stars, 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("stars", "body", *m.Stars, 5, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this model 87 based on context it is used
func (m *Model87) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Model87) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model87) UnmarshalBinary(b []byte) error {
	var res Model87
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
