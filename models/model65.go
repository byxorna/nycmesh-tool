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

// Model65 model 65
//
// swagger:model Model 65
type Model65 struct {

	// description
	Description string `json:"description,omitempty"`

	// entries
	// Required: true
	Entries Entries `json:"entries"`

	// identification
	// Required: true
	Identification *Identification4 `json:"identification"`
}

// Validate validates this model 65
func (m *Model65) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model65) validateEntries(formats strfmt.Registry) error {

	if err := validate.Required("entries", "body", m.Entries); err != nil {
		return err
	}

	if err := m.Entries.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("entries")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("entries")
		}
		return err
	}

	return nil
}

func (m *Model65) validateIdentification(formats strfmt.Registry) error {

	if err := validate.Required("identification", "body", m.Identification); err != nil {
		return err
	}

	if m.Identification != nil {
		if err := m.Identification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identification")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this model 65 based on the context it is used
func (m *Model65) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentification(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model65) contextValidateEntries(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Entries.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("entries")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("entries")
		}
		return err
	}

	return nil
}

func (m *Model65) contextValidateIdentification(ctx context.Context, formats strfmt.Registry) error {

	if m.Identification != nil {
		if err := m.Identification.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identification")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Model65) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model65) UnmarshalBinary(b []byte) error {
	var res Model65
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}