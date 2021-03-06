// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// SignalChain signal chain
//
// swagger:model signalChain
type SignalChain []float64

// Validate validates this signal chain
func (m SignalChain) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this signal chain based on context it is used
func (m SignalChain) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
