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

// AirLinkBeSchema air link be schema
//
// swagger:model AirLinkBeSchema
type AirLinkBeSchema struct {

	// distance
	// Required: true
	Distance *float64 `json:"distance"`

	// elevations
	// Required: true
	Elevations *Elevations1 `json:"elevations"`
}

// Validate validates this air link be schema
func (m *AirLinkBeSchema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDistance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElevations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AirLinkBeSchema) validateDistance(formats strfmt.Registry) error {

	if err := validate.Required("distance", "body", m.Distance); err != nil {
		return err
	}

	return nil
}

func (m *AirLinkBeSchema) validateElevations(formats strfmt.Registry) error {

	if err := validate.Required("elevations", "body", m.Elevations); err != nil {
		return err
	}

	if m.Elevations != nil {
		if err := m.Elevations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("elevations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("elevations")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this air link be schema based on the context it is used
func (m *AirLinkBeSchema) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateElevations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AirLinkBeSchema) contextValidateElevations(ctx context.Context, formats strfmt.Registry) error {

	if m.Elevations != nil {
		if err := m.Elevations.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("elevations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("elevations")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AirLinkBeSchema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AirLinkBeSchema) UnmarshalBinary(b []byte) error {
	var res AirLinkBeSchema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
