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

// OutInterface out interface
//
// swagger:model outInterface
type OutInterface struct {

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

// Validate validates this out interface
func (m *OutInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var outInterfaceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["eth","switch","port","pon","br","pppoe","vlan","sfp+","wlan","ath","lag","loopback","tunnel"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		outInterfaceTypeTypePropEnum = append(outInterfaceTypeTypePropEnum, v)
	}
}

const (

	// OutInterfaceTypeEth captures enum value "eth"
	OutInterfaceTypeEth string = "eth"

	// OutInterfaceTypeSwitch captures enum value "switch"
	OutInterfaceTypeSwitch string = "switch"

	// OutInterfaceTypePort captures enum value "port"
	OutInterfaceTypePort string = "port"

	// OutInterfaceTypePon captures enum value "pon"
	OutInterfaceTypePon string = "pon"

	// OutInterfaceTypeBr captures enum value "br"
	OutInterfaceTypeBr string = "br"

	// OutInterfaceTypePppoe captures enum value "pppoe"
	OutInterfaceTypePppoe string = "pppoe"

	// OutInterfaceTypeVlan captures enum value "vlan"
	OutInterfaceTypeVlan string = "vlan"

	// OutInterfaceTypeSfpPlus captures enum value "sfp+"
	OutInterfaceTypeSfpPlus string = "sfp+"

	// OutInterfaceTypeWlan captures enum value "wlan"
	OutInterfaceTypeWlan string = "wlan"

	// OutInterfaceTypeAth captures enum value "ath"
	OutInterfaceTypeAth string = "ath"

	// OutInterfaceTypeLag captures enum value "lag"
	OutInterfaceTypeLag string = "lag"

	// OutInterfaceTypeLoopback captures enum value "loopback"
	OutInterfaceTypeLoopback string = "loopback"

	// OutInterfaceTypeTunnel captures enum value "tunnel"
	OutInterfaceTypeTunnel string = "tunnel"
)

// prop value enum
func (m *OutInterface) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, outInterfaceTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OutInterface) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this out interface based on context it is used
func (m *OutInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OutInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OutInterface) UnmarshalBinary(b []byte) error {
	var res OutInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
