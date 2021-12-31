// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PayloadUnmsHostnameReachable payload unms hostname reachable
//
// swagger:model PayloadUnmsHostnameReachable
type PayloadUnmsHostnameReachable struct {

	// hostname
	Hostname string `json:"hostname,omitempty"`

	// port
	// Example: 443
	// Maximum: 65535
	// Minimum: 0
	Port *int64 `json:"port,omitempty"`
}

// Validate validates this payload unms hostname reachable
func (m *PayloadUnmsHostnameReachable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PayloadUnmsHostnameReachable) validatePort(formats strfmt.Registry) error {
	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("port", "body", *m.Port, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", *m.Port, 65535, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this payload unms hostname reachable based on context it is used
func (m *PayloadUnmsHostnameReachable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PayloadUnmsHostnameReachable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PayloadUnmsHostnameReachable) UnmarshalBinary(b []byte) error {
	var res PayloadUnmsHostnameReachable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}