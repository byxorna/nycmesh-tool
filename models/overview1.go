// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Overview1 overview 1
//
// swagger:model overview 1
type Overview1 struct {

	// Read-only value generated by UISP.
	Status string `json:"status,omitempty"`
}

// Validate validates this overview 1
func (m *Overview1) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this overview 1 based on context it is used
func (m *Overview1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Overview1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Overview1) UnmarshalBinary(b []byte) error {
	var res Overview1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
