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

// InterfaceIdentification interface identification
//
// swagger:model InterfaceIdentification
type InterfaceIdentification struct {

	// Nullable string.
	// Example: Uplink
	Description string `json:"description,omitempty"`

	// Computed display name from name and description
	// Example: eth0
	DisplayName string `json:"displayName,omitempty"`

	// mac
	// Example: fc:ec:da:03:bb:a8
	// Pattern: ^([0-9a-fA-F][0-9a-fA-F]:){5}([0-9a-fA-F][0-9a-fA-F])$|^([0-9a-fA-F]){12}$
	Mac string `json:"mac,omitempty"`

	// Interface name.
	// Example: eth0
	Name string `json:"name,omitempty"`

	// Physical port position.
	// Example: 0
	Position float64 `json:"position,omitempty"`

	// type
	// Example: eth
	Type string `json:"type,omitempty"`
}

// Validate validates this interface identification
func (m *InterfaceIdentification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMac(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InterfaceIdentification) validateMac(formats strfmt.Registry) error {
	if swag.IsZero(m.Mac) { // not required
		return nil
	}

	if err := validate.Pattern("mac", "body", m.Mac, `^([0-9a-fA-F][0-9a-fA-F]:){5}([0-9a-fA-F][0-9a-fA-F])$|^([0-9a-fA-F]){12}$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this interface identification based on context it is used
func (m *InterfaceIdentification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InterfaceIdentification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceIdentification) UnmarshalBinary(b []byte) error {
	var res InterfaceIdentification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}