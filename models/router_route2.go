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

// RouterRoute2 router route 2
//
// swagger:model RouterRoute 2
type RouterRoute2 struct {

	// description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// destination
	// Required: true
	Destination *string `json:"destination"`

	// distance
	// Maximum: 255
	// Minimum: 1
	Distance float64 `json:"distance,omitempty"`

	// static type
	// Required: true
	// Enum: [interface blackhole gateway]
	StaticType *string `json:"staticType"`
}

// Validate validates this router route 2
func (m *RouterRoute2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDistance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaticType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RouterRoute2) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *RouterRoute2) validateDestination(formats strfmt.Registry) error {

	if err := validate.Required("destination", "body", m.Destination); err != nil {
		return err
	}

	return nil
}

func (m *RouterRoute2) validateDistance(formats strfmt.Registry) error {
	if swag.IsZero(m.Distance) { // not required
		return nil
	}

	if err := validate.Minimum("distance", "body", m.Distance, 1, false); err != nil {
		return err
	}

	if err := validate.Maximum("distance", "body", m.Distance, 255, false); err != nil {
		return err
	}

	return nil
}

var routerRoute2TypeStaticTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["interface","blackhole","gateway"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routerRoute2TypeStaticTypePropEnum = append(routerRoute2TypeStaticTypePropEnum, v)
	}
}

const (

	// RouterRoute2StaticTypeInterface captures enum value "interface"
	RouterRoute2StaticTypeInterface string = "interface"

	// RouterRoute2StaticTypeBlackhole captures enum value "blackhole"
	RouterRoute2StaticTypeBlackhole string = "blackhole"

	// RouterRoute2StaticTypeGateway captures enum value "gateway"
	RouterRoute2StaticTypeGateway string = "gateway"
)

// prop value enum
func (m *RouterRoute2) validateStaticTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, routerRoute2TypeStaticTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RouterRoute2) validateStaticType(formats strfmt.Registry) error {

	if err := validate.Required("staticType", "body", m.StaticType); err != nil {
		return err
	}

	// value enum
	if err := m.validateStaticTypeEnum("staticType", "body", *m.StaticType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this router route 2 based on context it is used
func (m *RouterRoute2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RouterRoute2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RouterRoute2) UnmarshalBinary(b []byte) error {
	var res RouterRoute2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
