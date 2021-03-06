// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Model67 model 67
//
// swagger:model Model 67
type Model67 struct {

	// display name
	// Example: eth0
	DisplayName string `json:"displayName,omitempty"`

	// Interface name.
	// Example: eth0
	ID string `json:"id,omitempty"`

	// usable for gateway
	// Example: true
	UsableForGateway bool `json:"usableForGateway,omitempty"`

	// used auto
	// Example: true
	UsedAuto bool `json:"usedAuto,omitempty"`
}

// Validate validates this model 67
func (m *Model67) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this model 67 based on context it is used
func (m *Model67) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Model67) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model67) UnmarshalBinary(b []byte) error {
	var res Model67
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
