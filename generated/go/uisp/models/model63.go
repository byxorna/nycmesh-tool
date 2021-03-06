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

// Model63 model 63
//
// swagger:model Model 63
type Model63 struct {

	// config
	// Required: true
	Config *Config1 `json:"config"`

	// rules
	// Required: true
	Rules Rules1 `json:"rules"`

	// unmatched bytes
	// Required: true
	// Minimum: 0
	UnmatchedBytes *int64 `json:"unmatchedBytes"`

	// unmatched packets
	// Required: true
	// Minimum: 0
	UnmatchedPackets *int64 `json:"unmatchedPackets"`
}

// Validate validates this model 63
func (m *Model63) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnmatchedBytes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnmatchedPackets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model63) validateConfig(formats strfmt.Registry) error {

	if err := validate.Required("config", "body", m.Config); err != nil {
		return err
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *Model63) validateRules(formats strfmt.Registry) error {

	if err := validate.Required("rules", "body", m.Rules); err != nil {
		return err
	}

	if err := m.Rules.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rules")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("rules")
		}
		return err
	}

	return nil
}

func (m *Model63) validateUnmatchedBytes(formats strfmt.Registry) error {

	if err := validate.Required("unmatchedBytes", "body", m.UnmatchedBytes); err != nil {
		return err
	}

	if err := validate.MinimumInt("unmatchedBytes", "body", *m.UnmatchedBytes, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Model63) validateUnmatchedPackets(formats strfmt.Registry) error {

	if err := validate.Required("unmatchedPackets", "body", m.UnmatchedPackets); err != nil {
		return err
	}

	if err := validate.MinimumInt("unmatchedPackets", "body", *m.UnmatchedPackets, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this model 63 based on the context it is used
func (m *Model63) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model63) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {
		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *Model63) contextValidateRules(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Rules.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rules")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("rules")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Model63) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model63) UnmarshalBinary(b []byte) error {
	var res Model63
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
