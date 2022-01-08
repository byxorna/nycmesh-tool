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

// UpdateDHCPServer update d h c p server
//
// swagger:model UpdateDHCPServer
type UpdateDHCPServer struct {

	// Available leases in DHCP pool.
	// Example: 90
	// Minimum: 0
	Available *float64 `json:"available,omitempty"`

	// Primary DNS server address.
	// Example: 8.8.8.8
	Dns1 string `json:"dns1,omitempty"`

	// Secondary DNS server address.
	// Example: 8.8.4.4
	Dns2 string `json:"dns2,omitempty"`

	// Domain name.
	// Example: uisp.com
	// Min Length: 1
	Domain string `json:"domain,omitempty"`

	// Set to TRUE to enable DHCP server.
	// Example: true
	Enabled bool `json:"enabled,omitempty"`

	// Interface IP v4 address in CIDR format.
	// Example: 192.168.2.0/24
	// Required: true
	Interface *string `json:"interface"`

	// DHCP lease time in seconds.
	// Example: 86400
	// Required: true
	LeaseTime *float64 `json:"leaseTime"`

	// DHCP server custom name.
	// Example: Main DHCP server
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`

	// Total range of leases in DHCP pool.
	// Example: 100
	// Minimum: 0
	PoolSize *float64 `json:"poolSize,omitempty"`

	// DHCP addresses pool end in CIDR format.
	// Example: 192.168.2.250/24
	// Required: true
	RangeEnd *string `json:"rangeEnd"`

	// DHCP addresses pool start in CIDR format.
	// Example: 192.168.2.51/24
	// Required: true
	RangeStart *string `json:"rangeStart"`

	// Router IP v4 address.
	// Example: 192.168.2.1
	Router string `json:"router,omitempty"`

	// Unifi controller IP v4 address.
	// Example: 192.168.2.1
	UnifiController string `json:"unifiController,omitempty"`
}

// Validate validates this update d h c p server
func (m *UpdateDHCPServer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterface(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLeaseTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePoolSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRangeEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRangeStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateDHCPServer) validateAvailable(formats strfmt.Registry) error {
	if swag.IsZero(m.Available) { // not required
		return nil
	}

	if err := validate.Minimum("available", "body", *m.Available, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *UpdateDHCPServer) validateDomain(formats strfmt.Registry) error {
	if swag.IsZero(m.Domain) { // not required
		return nil
	}

	if err := validate.MinLength("domain", "body", m.Domain, 1); err != nil {
		return err
	}

	return nil
}

func (m *UpdateDHCPServer) validateInterface(formats strfmt.Registry) error {

	if err := validate.Required("interface", "body", m.Interface); err != nil {
		return err
	}

	return nil
}

func (m *UpdateDHCPServer) validateLeaseTime(formats strfmt.Registry) error {

	if err := validate.Required("leaseTime", "body", m.LeaseTime); err != nil {
		return err
	}

	return nil
}

func (m *UpdateDHCPServer) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	return nil
}

func (m *UpdateDHCPServer) validatePoolSize(formats strfmt.Registry) error {
	if swag.IsZero(m.PoolSize) { // not required
		return nil
	}

	if err := validate.Minimum("poolSize", "body", *m.PoolSize, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *UpdateDHCPServer) validateRangeEnd(formats strfmt.Registry) error {

	if err := validate.Required("rangeEnd", "body", m.RangeEnd); err != nil {
		return err
	}

	return nil
}

func (m *UpdateDHCPServer) validateRangeStart(formats strfmt.Registry) error {

	if err := validate.Required("rangeStart", "body", m.RangeStart); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update d h c p server based on context it is used
func (m *UpdateDHCPServer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateDHCPServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateDHCPServer) UnmarshalBinary(b []byte) error {
	var res UpdateDHCPServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}