// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AddressData address data
//
// swagger:model addressData
type AddressData struct {

	// address components
	AddressComponents Interface `json:"addressComponents,omitempty"`

	// Map provider.
	// Example: GoogleMaps
	// Required: true
	// Enum: [GoogleMaps OpenStreetMap]
	Type *string `json:"type"`
}

// Validate validates this address data
func (m *AddressData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addressDataTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GoogleMaps","OpenStreetMap"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addressDataTypeTypePropEnum = append(addressDataTypeTypePropEnum, v)
	}
}

const (

	// AddressDataTypeGoogleMaps captures enum value "GoogleMaps"
	AddressDataTypeGoogleMaps string = "GoogleMaps"

	// AddressDataTypeOpenStreetMap captures enum value "OpenStreetMap"
	AddressDataTypeOpenStreetMap string = "OpenStreetMap"
)

// prop value enum
func (m *AddressData) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addressDataTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AddressData) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this address data based on context it is used
func (m *AddressData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AddressData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddressData) UnmarshalBinary(b []byte) error {
	var res AddressData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
