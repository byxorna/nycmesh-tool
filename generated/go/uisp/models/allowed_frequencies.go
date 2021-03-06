// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// AllowedFrequencies allowed frequencies
//
// swagger:model allowedFrequencies
type AllowedFrequencies []float64

// Validate validates this allowed frequencies
func (m AllowedFrequencies) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this allowed frequencies based on context it is used
func (m AllowedFrequencies) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
