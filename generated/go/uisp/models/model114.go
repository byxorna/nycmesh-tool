// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Model114 model 114
//
// swagger:model Model 114
type Model114 struct {

	// file
	// Required: true
	// Format: binary
	File io.ReadCloser `json:"file"`
}

// Validate validates this model 114
func (m *Model114) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model114) validateFile(formats strfmt.Registry) error {

	if err := validate.Required("file", "body", io.ReadCloser(m.File)); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this model 114 based on context it is used
func (m *Model114) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Model114) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model114) UnmarshalBinary(b []byte) error {
	var res Model114
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}