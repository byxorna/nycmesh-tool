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

// Identification3 identification 3
//
// swagger:model identification 3
type Identification3 struct {

	// id
	// Example: f7ac9cad-ea28-4390-93c8-7add010e8ee3
	// Required: true
	ID *string `json:"id"`

	// type
	// Enum: [onu olt uispp uispr uisps uispLte erouter eswitch epower airCube airMax airFiber toughSwitch solarBeam wave blackBox]
	Type string `json:"type,omitempty"`
}

// Validate validates this identification 3
func (m *Identification3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Identification3) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var identification3TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["onu","olt","uispp","uispr","uisps","uispLte","erouter","eswitch","epower","airCube","airMax","airFiber","toughSwitch","solarBeam","wave","blackBox"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		identification3TypeTypePropEnum = append(identification3TypeTypePropEnum, v)
	}
}

const (

	// Identification3TypeOnu captures enum value "onu"
	Identification3TypeOnu string = "onu"

	// Identification3TypeOlt captures enum value "olt"
	Identification3TypeOlt string = "olt"

	// Identification3TypeUispp captures enum value "uispp"
	Identification3TypeUispp string = "uispp"

	// Identification3TypeUispr captures enum value "uispr"
	Identification3TypeUispr string = "uispr"

	// Identification3TypeUisps captures enum value "uisps"
	Identification3TypeUisps string = "uisps"

	// Identification3TypeUispLte captures enum value "uispLte"
	Identification3TypeUispLte string = "uispLte"

	// Identification3TypeErouter captures enum value "erouter"
	Identification3TypeErouter string = "erouter"

	// Identification3TypeEswitch captures enum value "eswitch"
	Identification3TypeEswitch string = "eswitch"

	// Identification3TypeEpower captures enum value "epower"
	Identification3TypeEpower string = "epower"

	// Identification3TypeAirCube captures enum value "airCube"
	Identification3TypeAirCube string = "airCube"

	// Identification3TypeAirMax captures enum value "airMax"
	Identification3TypeAirMax string = "airMax"

	// Identification3TypeAirFiber captures enum value "airFiber"
	Identification3TypeAirFiber string = "airFiber"

	// Identification3TypeToughSwitch captures enum value "toughSwitch"
	Identification3TypeToughSwitch string = "toughSwitch"

	// Identification3TypeSolarBeam captures enum value "solarBeam"
	Identification3TypeSolarBeam string = "solarBeam"

	// Identification3TypeWave captures enum value "wave"
	Identification3TypeWave string = "wave"

	// Identification3TypeBlackBox captures enum value "blackBox"
	Identification3TypeBlackBox string = "blackBox"
)

// prop value enum
func (m *Identification3) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, identification3TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Identification3) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this identification 3 based on context it is used
func (m *Identification3) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Identification3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Identification3) UnmarshalBinary(b []byte) error {
	var res Identification3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
