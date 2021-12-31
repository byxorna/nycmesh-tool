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

// SimulationLink1 simulation link 1
//
// swagger:model SimulationLink 1
type SimulationLink1 struct {

	// ap device Id
	// Example: f7ac9cad-ea28-4390-93c8-7add010e8ee3
	// Required: true
	ApDeviceID *string `json:"apDeviceId"`

	// cpe device Id
	// Example: f7ac9cad-ea28-4390-93c8-7add010e8ee3
	// Required: true
	CpeDeviceID *string `json:"cpeDeviceId"`

	// created at
	// Required: true
	// Format: date
	CreatedAt *strfmt.Date `json:"createdAt"`

	// id
	// Example: f7ac9cad-ea28-4390-93c8-7add010e8ee3
	// Required: true
	ID *string `json:"id"`

	// quality
	// Required: true
	// Max Length: 200
	Quality *string `json:"quality"`

	// terrain
	Terrain *Terrain `json:"terrain,omitempty"`

	// updated at
	// Required: true
	// Format: date
	UpdatedAt *strfmt.Date `json:"updatedAt"`
}

// Validate validates this simulation link 1
func (m *SimulationLink1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCpeDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuality(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerrain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SimulationLink1) validateApDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("apDeviceId", "body", m.ApDeviceID); err != nil {
		return err
	}

	return nil
}

func (m *SimulationLink1) validateCpeDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("cpeDeviceId", "body", m.CpeDeviceID); err != nil {
		return err
	}

	return nil
}

func (m *SimulationLink1) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SimulationLink1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SimulationLink1) validateQuality(formats strfmt.Registry) error {

	if err := validate.Required("quality", "body", m.Quality); err != nil {
		return err
	}

	if err := validate.MaxLength("quality", "body", *m.Quality, 200); err != nil {
		return err
	}

	return nil
}

func (m *SimulationLink1) validateTerrain(formats strfmt.Registry) error {
	if swag.IsZero(m.Terrain) { // not required
		return nil
	}

	if m.Terrain != nil {
		if err := m.Terrain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terrain")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("terrain")
			}
			return err
		}
	}

	return nil
}

func (m *SimulationLink1) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this simulation link 1 based on the context it is used
func (m *SimulationLink1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTerrain(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SimulationLink1) contextValidateTerrain(ctx context.Context, formats strfmt.Registry) error {

	if m.Terrain != nil {
		if err := m.Terrain.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terrain")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("terrain")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SimulationLink1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SimulationLink1) UnmarshalBinary(b []byte) error {
	var res SimulationLink1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
