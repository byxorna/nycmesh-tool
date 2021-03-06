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

// RouterRouteUpsert1 router route upsert 1
//
// swagger:model RouterRouteUpsert 1
type RouterRouteUpsert1 struct {

	// description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// destination
	// Required: true
	Destination *string `json:"destination"`

	// distance
	// Required: true
	// Maximum: 256
	// Minimum: 0
	Distance *float64 `json:"distance"`

	// static type
	// Required: true
	// Enum: [interface blackhole gateway]
	StaticType *string `json:"staticType"`
}

// Validate validates this router route upsert 1
func (m *RouterRouteUpsert1) Validate(formats strfmt.Registry) error {
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

func (m *RouterRouteUpsert1) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *RouterRouteUpsert1) validateDestination(formats strfmt.Registry) error {

	if err := validate.Required("destination", "body", m.Destination); err != nil {
		return err
	}

	return nil
}

func (m *RouterRouteUpsert1) validateDistance(formats strfmt.Registry) error {

	if err := validate.Required("distance", "body", m.Distance); err != nil {
		return err
	}

	if err := validate.Minimum("distance", "body", *m.Distance, 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("distance", "body", *m.Distance, 256, false); err != nil {
		return err
	}

	return nil
}

var routerRouteUpsert1TypeStaticTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["interface","blackhole","gateway"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routerRouteUpsert1TypeStaticTypePropEnum = append(routerRouteUpsert1TypeStaticTypePropEnum, v)
	}
}

const (

	// RouterRouteUpsert1StaticTypeInterface captures enum value "interface"
	RouterRouteUpsert1StaticTypeInterface string = "interface"

	// RouterRouteUpsert1StaticTypeBlackhole captures enum value "blackhole"
	RouterRouteUpsert1StaticTypeBlackhole string = "blackhole"

	// RouterRouteUpsert1StaticTypeGateway captures enum value "gateway"
	RouterRouteUpsert1StaticTypeGateway string = "gateway"
)

// prop value enum
func (m *RouterRouteUpsert1) validateStaticTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, routerRouteUpsert1TypeStaticTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RouterRouteUpsert1) validateStaticType(formats strfmt.Registry) error {

	if err := validate.Required("staticType", "body", m.StaticType); err != nil {
		return err
	}

	// value enum
	if err := m.validateStaticTypeEnum("staticType", "body", *m.StaticType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this router route upsert 1 based on context it is used
func (m *RouterRouteUpsert1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RouterRouteUpsert1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RouterRouteUpsert1) UnmarshalBinary(b []byte) error {
	var res RouterRouteUpsert1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
