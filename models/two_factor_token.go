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

// TwoFactorToken two factor token
//
// swagger:model TwoFactorToken
type TwoFactorToken struct {

	// exp
	// Required: true
	Exp *float64 `json:"exp"`

	// id
	// Required: true
	ID *string `json:"id"`

	// user Id
	// Required: true
	UserID *string `json:"userId"`
}

// Validate validates this two factor token
func (m *TwoFactorToken) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TwoFactorToken) validateExp(formats strfmt.Registry) error {

	if err := validate.Required("exp", "body", m.Exp); err != nil {
		return err
	}

	return nil
}

func (m *TwoFactorToken) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *TwoFactorToken) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this two factor token based on context it is used
func (m *TwoFactorToken) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TwoFactorToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TwoFactorToken) UnmarshalBinary(b []byte) error {
	var res TwoFactorToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
