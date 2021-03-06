// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Aggregation1 aggregation 1
//
// swagger:model aggregation 1
type Aggregation1 struct {

	// all count
	AllCount float64 `json:"allCount,omitempty"`

	// outage count
	OutageCount float64 `json:"outageCount,omitempty"`

	// unreachable count
	UnreachableCount float64 `json:"unreachableCount,omitempty"`
}

// Validate validates this aggregation 1
func (m *Aggregation1) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this aggregation 1 based on context it is used
func (m *Aggregation1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Aggregation1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Aggregation1) UnmarshalBinary(b []byte) error {
	var res Aggregation1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
