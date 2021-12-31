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

// Model125 model 125
//
// swagger:model Model 125
type Model125 struct {

	// auth
	// Required: true
	// Enum: [md5 plaintext-password off]
	Auth *string `json:"auth"`

	// auth key
	AuthKey string `json:"authKey,omitempty"`

	// auth keys m d5
	AuthKeysMD5 AuthKeysMD52 `json:"authKeysMD5,omitempty"`

	// cost
	// Required: true
	// Maximum: 255
	// Minimum: 1
	Cost *int64 `json:"cost"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this model 125
func (m *Model125) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthKeysMD5(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCost(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var model125TypeAuthPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["md5","plaintext-password","off"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		model125TypeAuthPropEnum = append(model125TypeAuthPropEnum, v)
	}
}

const (

	// Model125AuthMd5 captures enum value "md5"
	Model125AuthMd5 string = "md5"

	// Model125AuthPlaintextDashPassword captures enum value "plaintext-password"
	Model125AuthPlaintextDashPassword string = "plaintext-password"

	// Model125AuthOff captures enum value "off"
	Model125AuthOff string = "off"
)

// prop value enum
func (m *Model125) validateAuthEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, model125TypeAuthPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Model125) validateAuth(formats strfmt.Registry) error {

	if err := validate.Required("auth", "body", m.Auth); err != nil {
		return err
	}

	// value enum
	if err := m.validateAuthEnum("auth", "body", *m.Auth); err != nil {
		return err
	}

	return nil
}

func (m *Model125) validateAuthKeysMD5(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthKeysMD5) { // not required
		return nil
	}

	if err := m.AuthKeysMD5.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("authKeysMD5")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("authKeysMD5")
		}
		return err
	}

	return nil
}

func (m *Model125) validateCost(formats strfmt.Registry) error {

	if err := validate.Required("cost", "body", m.Cost); err != nil {
		return err
	}

	if err := validate.MinimumInt("cost", "body", *m.Cost, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("cost", "body", *m.Cost, 255, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this model 125 based on the context it is used
func (m *Model125) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthKeysMD5(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model125) contextValidateAuthKeysMD5(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AuthKeysMD5.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("authKeysMD5")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("authKeysMD5")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Model125) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model125) UnmarshalBinary(b []byte) error {
	var res Model125
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
