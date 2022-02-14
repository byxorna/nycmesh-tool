// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SfpPlus1 sfp plus1
//
// swagger:model sfpPlus1
type SfpPlus1 struct {

	// vlan native
	VlanNative float64 `json:"vlanNative,omitempty"`
}

// Validate validates this sfp plus1
func (m *SfpPlus1) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sfp plus1 based on context it is used
func (m *SfpPlus1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SfpPlus1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SfpPlus1) UnmarshalBinary(b []byte) error {
	var res SfpPlus1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
