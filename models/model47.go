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

// Model47 model 47
//
// swagger:model Model 47
type Model47 struct {

	// HTTP port number
	// Example: 80
	// Required: true
	// Maximum: 65535
	// Minimum: 0
	HTTPPort *int64 `json:"httpPort"`

	// Set to true if SSH is enabled
	// Example: true
	// Required: true
	SSHEnabled *bool `json:"sshEnabled"`

	// SSH port number
	// Example: 22
	// Required: true
	// Maximum: 65535
	// Minimum: 0
	SSHPort *int64 `json:"sshPort"`

	// Set to true if Telnet is enabled
	// Example: true
	// Required: true
	TelnetEnabled *bool `json:"telnetEnabled"`

	// Telnet port number
	// Example: 23
	// Required: true
	// Maximum: 65535
	// Minimum: 0
	TelnetPort *int64 `json:"telnetPort"`

	// Set to true if UNMS discovery is enabled
	// Example: true
	// Required: true
	UbntDiscoveryEnabled *bool `json:"ubntDiscoveryEnabled"`
}

// Validate validates this model 47
func (m *Model47) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHTTPPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSHEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSHPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTelnetEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTelnetPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUbntDiscoveryEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model47) validateHTTPPort(formats strfmt.Registry) error {

	if err := validate.Required("httpPort", "body", m.HTTPPort); err != nil {
		return err
	}

	if err := validate.MinimumInt("httpPort", "body", *m.HTTPPort, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("httpPort", "body", *m.HTTPPort, 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *Model47) validateSSHEnabled(formats strfmt.Registry) error {

	if err := validate.Required("sshEnabled", "body", m.SSHEnabled); err != nil {
		return err
	}

	return nil
}

func (m *Model47) validateSSHPort(formats strfmt.Registry) error {

	if err := validate.Required("sshPort", "body", m.SSHPort); err != nil {
		return err
	}

	if err := validate.MinimumInt("sshPort", "body", *m.SSHPort, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("sshPort", "body", *m.SSHPort, 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *Model47) validateTelnetEnabled(formats strfmt.Registry) error {

	if err := validate.Required("telnetEnabled", "body", m.TelnetEnabled); err != nil {
		return err
	}

	return nil
}

func (m *Model47) validateTelnetPort(formats strfmt.Registry) error {

	if err := validate.Required("telnetPort", "body", m.TelnetPort); err != nil {
		return err
	}

	if err := validate.MinimumInt("telnetPort", "body", *m.TelnetPort, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("telnetPort", "body", *m.TelnetPort, 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *Model47) validateUbntDiscoveryEnabled(formats strfmt.Registry) error {

	if err := validate.Required("ubntDiscoveryEnabled", "body", m.UbntDiscoveryEnabled); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this model 47 based on context it is used
func (m *Model47) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Model47) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model47) UnmarshalBinary(b []byte) error {
	var res Model47
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
