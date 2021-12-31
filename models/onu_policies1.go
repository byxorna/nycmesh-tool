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

// OnuPolicies1 onu policies 1
//
// swagger:model OnuPolicies 1
type OnuPolicies1 struct {

	// default state
	// Required: true
	// Enum: [enabled disabled]
	DefaultState *string `json:"defaultState"`

	// dhcp option82
	DhcpOption82 bool `json:"dhcpOption82,omitempty"`
}

// Validate validates this onu policies 1
func (m *OnuPolicies1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var onuPolicies1TypeDefaultStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		onuPolicies1TypeDefaultStatePropEnum = append(onuPolicies1TypeDefaultStatePropEnum, v)
	}
}

const (

	// OnuPolicies1DefaultStateEnabled captures enum value "enabled"
	OnuPolicies1DefaultStateEnabled string = "enabled"

	// OnuPolicies1DefaultStateDisabled captures enum value "disabled"
	OnuPolicies1DefaultStateDisabled string = "disabled"
)

// prop value enum
func (m *OnuPolicies1) validateDefaultStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, onuPolicies1TypeDefaultStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OnuPolicies1) validateDefaultState(formats strfmt.Registry) error {

	if err := validate.Required("defaultState", "body", m.DefaultState); err != nil {
		return err
	}

	// value enum
	if err := m.validateDefaultStateEnum("defaultState", "body", *m.DefaultState); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this onu policies 1 based on context it is used
func (m *OnuPolicies1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OnuPolicies1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OnuPolicies1) UnmarshalBinary(b []byte) error {
	var res OnuPolicies1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
