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

// Model70 model 70
//
// swagger:model Model 70
type Model70 struct {

	// Schedule upgrade over maintenance window.
	UpgradeInMaintenanceWindow bool `json:"upgradeInMaintenanceWindow,omitempty"`

	// upgrades
	// Required: true
	Upgrades Upgrades `json:"upgrades"`
}

// Validate validates this model 70
func (m *Model70) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUpgrades(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model70) validateUpgrades(formats strfmt.Registry) error {

	if err := validate.Required("upgrades", "body", m.Upgrades); err != nil {
		return err
	}

	if err := m.Upgrades.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("upgrades")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("upgrades")
		}
		return err
	}

	return nil
}

// ContextValidate validate this model 70 based on the context it is used
func (m *Model70) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUpgrades(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model70) contextValidateUpgrades(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Upgrades.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("upgrades")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("upgrades")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Model70) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model70) UnmarshalBinary(b []byte) error {
	var res Model70
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
