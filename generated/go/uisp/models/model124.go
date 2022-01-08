// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Model124 model 124
//
// swagger:model Model 124
type Model124 struct {

	// id
	ID float64 `json:"id,omitempty"`

	// key
	Key string `json:"key,omitempty"`
}

// Validate validates this model 124
func (m *Model124) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this model 124 based on context it is used
func (m *Model124) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Model124) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model124) UnmarshalBinary(b []byte) error {
	var res Model124
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}