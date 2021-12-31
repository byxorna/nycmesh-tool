// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InterfaceStatistics interface statistics
//
// swagger:model InterfaceStatistics
type InterfaceStatistics struct {

	// dropped
	// Example: 0
	Dropped float64 `json:"dropped,omitempty"`

	// errors
	// Example: 0
	Errors float64 `json:"errors,omitempty"`

	// poe power
	// Example: 736
	PoePower float64 `json:"poePower,omitempty"`

	// rxbytes
	// Example: 7487858302
	Rxbytes float64 `json:"rxbytes,omitempty"`

	// rxrate
	// Example: 3440
	Rxrate float64 `json:"rxrate,omitempty"`

	// txbytes
	// Example: 368737600
	Txbytes float64 `json:"txbytes,omitempty"`

	// txrate
	// Example: 736
	Txrate float64 `json:"txrate,omitempty"`
}

// Validate validates this interface statistics
func (m *InterfaceStatistics) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this interface statistics based on context it is used
func (m *InterfaceStatistics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InterfaceStatistics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceStatistics) UnmarshalBinary(b []byte) error {
	var res InterfaceStatistics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
