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

// Terrain terrain
//
// swagger:model terrain
type Terrain struct {

	// altitude ap
	// Required: true
	// Maximum: 10000
	// Minimum: -10000
	AltitudeAp *float64 `json:"altitudeAp"`

	// altitude cpe
	// Required: true
	// Maximum: 10000
	// Minimum: -10000
	AltitudeCpe *float64 `json:"altitudeCpe"`
}

// Validate validates this terrain
func (m *Terrain) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAltitudeAp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAltitudeCpe(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Terrain) validateAltitudeAp(formats strfmt.Registry) error {

	if err := validate.Required("altitudeAp", "body", m.AltitudeAp); err != nil {
		return err
	}

	if err := validate.Minimum("altitudeAp", "body", *m.AltitudeAp, -10000, false); err != nil {
		return err
	}

	if err := validate.Maximum("altitudeAp", "body", *m.AltitudeAp, 10000, false); err != nil {
		return err
	}

	return nil
}

func (m *Terrain) validateAltitudeCpe(formats strfmt.Registry) error {

	if err := validate.Required("altitudeCpe", "body", m.AltitudeCpe); err != nil {
		return err
	}

	if err := validate.Minimum("altitudeCpe", "body", *m.AltitudeCpe, -10000, false); err != nil {
		return err
	}

	if err := validate.Maximum("altitudeCpe", "body", *m.AltitudeCpe, 10000, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this terrain based on context it is used
func (m *Terrain) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Terrain) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Terrain) UnmarshalBinary(b []byte) error {
	var res Terrain
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
