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

// OutInterface1 out interface 1
//
// swagger:model outInterface 1
type OutInterface1 struct {

	// id
	ID string `json:"id,omitempty"`

	// mac
	Mac string `json:"mac,omitempty"`

	// mac override
	MacOverride string `json:"macOverride,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// type
	// Enum: [eth switch port pon br pppoe vlan sfp+ wlan ath lag loopback tunnel]
	Type string `json:"type,omitempty"`
}

// Validate validates this out interface 1
func (m *OutInterface1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var outInterface1TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["eth","switch","port","pon","br","pppoe","vlan","sfp+","wlan","ath","lag","loopback","tunnel"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		outInterface1TypeTypePropEnum = append(outInterface1TypeTypePropEnum, v)
	}
}

const (

	// OutInterface1TypeEth captures enum value "eth"
	OutInterface1TypeEth string = "eth"

	// OutInterface1TypeSwitch captures enum value "switch"
	OutInterface1TypeSwitch string = "switch"

	// OutInterface1TypePort captures enum value "port"
	OutInterface1TypePort string = "port"

	// OutInterface1TypePon captures enum value "pon"
	OutInterface1TypePon string = "pon"

	// OutInterface1TypeBr captures enum value "br"
	OutInterface1TypeBr string = "br"

	// OutInterface1TypePppoe captures enum value "pppoe"
	OutInterface1TypePppoe string = "pppoe"

	// OutInterface1TypeVlan captures enum value "vlan"
	OutInterface1TypeVlan string = "vlan"

	// OutInterface1TypeSfpPlus captures enum value "sfp+"
	OutInterface1TypeSfpPlus string = "sfp+"

	// OutInterface1TypeWlan captures enum value "wlan"
	OutInterface1TypeWlan string = "wlan"

	// OutInterface1TypeAth captures enum value "ath"
	OutInterface1TypeAth string = "ath"

	// OutInterface1TypeLag captures enum value "lag"
	OutInterface1TypeLag string = "lag"

	// OutInterface1TypeLoopback captures enum value "loopback"
	OutInterface1TypeLoopback string = "loopback"

	// OutInterface1TypeTunnel captures enum value "tunnel"
	OutInterface1TypeTunnel string = "tunnel"
)

// prop value enum
func (m *OutInterface1) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, outInterface1TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OutInterface1) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this out interface 1 based on context it is used
func (m *OutInterface1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OutInterface1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OutInterface1) UnmarshalBinary(b []byte) error {
	var res OutInterface1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
