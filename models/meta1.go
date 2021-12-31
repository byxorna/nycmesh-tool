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

// Meta1 meta 1
//
// swagger:model meta 1
type Meta1 struct {

	// alias
	// Max Length: 30
	Alias string `json:"alias,omitempty"`

	// custom Ip address
	CustomIPAddress string `json:"customIpAddress,omitempty"`

	// maintenance
	// Required: true
	Maintenance *bool `json:"maintenance"`

	// note
	// Max Length: 300
	Note string `json:"note,omitempty"`
}

// Validate validates this meta 1
func (m *Meta1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlias(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaintenance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNote(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Meta1) validateAlias(formats strfmt.Registry) error {
	if swag.IsZero(m.Alias) { // not required
		return nil
	}

	if err := validate.MaxLength("alias", "body", m.Alias, 30); err != nil {
		return err
	}

	return nil
}

func (m *Meta1) validateMaintenance(formats strfmt.Registry) error {

	if err := validate.Required("maintenance", "body", m.Maintenance); err != nil {
		return err
	}

	return nil
}

func (m *Meta1) validateNote(formats strfmt.Registry) error {
	if swag.IsZero(m.Note) { // not required
		return nil
	}

	if err := validate.MaxLength("note", "body", m.Note, 300); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this meta 1 based on context it is used
func (m *Meta1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Meta1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Meta1) UnmarshalBinary(b []byte) error {
	var res Meta1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
