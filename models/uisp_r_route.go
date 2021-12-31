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

// UispRRoute uisp r route
//
// swagger:model UispRRoute
type UispRRoute struct {

	// description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// destination
	// Required: true
	Destination *string `json:"destination"`

	// distance
	// Maximum: 256
	// Minimum: 0
	Distance *float64 `json:"distance,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// fib
	Fib bool `json:"fib,omitempty"`

	// gateway
	Gateway string `json:"gateway,omitempty"`

	// gateway status
	GatewayStatus string `json:"gatewayStatus,omitempty"`

	// interface
	Interface string `json:"interface,omitempty"`

	// next hop
	NextHop string `json:"nextHop,omitempty"`

	// selected
	Selected bool `json:"selected,omitempty"`

	// static type
	StaticType string `json:"staticType,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this uisp r route
func (m *UispRRoute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDistance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UispRRoute) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *UispRRoute) validateDestination(formats strfmt.Registry) error {

	if err := validate.Required("destination", "body", m.Destination); err != nil {
		return err
	}

	return nil
}

func (m *UispRRoute) validateDistance(formats strfmt.Registry) error {
	if swag.IsZero(m.Distance) { // not required
		return nil
	}

	if err := validate.Minimum("distance", "body", *m.Distance, 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("distance", "body", *m.Distance, 256, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this uisp r route based on context it is used
func (m *UispRRoute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UispRRoute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UispRRoute) UnmarshalBinary(b []byte) error {
	var res UispRRoute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}