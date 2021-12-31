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

// ConntrackModules conntrack modules
//
// swagger:model conntrackModules
type ConntrackModules struct {

	// ftp
	// Required: true
	Ftp *bool `json:"ftp"`

	// gre
	// Required: true
	Gre *bool `json:"gre"`

	// h323
	// Required: true
	H323 *bool `json:"h323"`

	// pptp
	// Required: true
	Pptp *bool `json:"pptp"`

	// sip
	// Required: true
	Sip *bool `json:"sip"`

	// tftp
	// Required: true
	Tftp *bool `json:"tftp"`
}

// Validate validates this conntrack modules
func (m *ConntrackModules) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFtp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGre(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateH323(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePptp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSip(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTftp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConntrackModules) validateFtp(formats strfmt.Registry) error {

	if err := validate.Required("ftp", "body", m.Ftp); err != nil {
		return err
	}

	return nil
}

func (m *ConntrackModules) validateGre(formats strfmt.Registry) error {

	if err := validate.Required("gre", "body", m.Gre); err != nil {
		return err
	}

	return nil
}

func (m *ConntrackModules) validateH323(formats strfmt.Registry) error {

	if err := validate.Required("h323", "body", m.H323); err != nil {
		return err
	}

	return nil
}

func (m *ConntrackModules) validatePptp(formats strfmt.Registry) error {

	if err := validate.Required("pptp", "body", m.Pptp); err != nil {
		return err
	}

	return nil
}

func (m *ConntrackModules) validateSip(formats strfmt.Registry) error {

	if err := validate.Required("sip", "body", m.Sip); err != nil {
		return err
	}

	return nil
}

func (m *ConntrackModules) validateTftp(formats strfmt.Registry) error {

	if err := validate.Required("tftp", "body", m.Tftp); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this conntrack modules based on context it is used
func (m *ConntrackModules) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConntrackModules) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConntrackModules) UnmarshalBinary(b []byte) error {
	var res ConntrackModules
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
