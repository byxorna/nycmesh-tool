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

// DHCPServer1 d h c p server 1
//
// swagger:model DHCPServer 1
type DHCPServer1 struct {

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

	// Total amount of dynamic leases used.
	// Example: 5
	// Minimum: 0
	DynamicLeases *float64 `json:"dynamicLeases,omitempty"`

	// Set to TRUE to enable DHCP server.
	// Example: true
	Enabled bool `json:"enabled,omitempty"`

	// Interface id.
	// Example: eth0
	// Required: true
	Interface *string `json:"interface"`

	// DHCP lease time in seconds.
	// Example: 86400
	LeaseTime float64 `json:"leaseTime,omitempty"`

	// DHCP server custom name.
	// Example: Main DHCP server
	// Min Length: 1
	Name string `json:"name,omitempty"`

	// Total range of leases in DHCP pool.
	// Example: 100
	// Required: true
	// Minimum: 0
	PoolSize *float64 `json:"poolSize"`

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

	// Total amount of static leases used.
	// Example: 5
	// Minimum: 0
	StaticLeases *float64 `json:"staticLeases,omitempty"`

	// Unifi controller IP v4 address.
	// Example: 192.168.2.1
	UnifiController string `json:"unifiController,omitempty"`
}

// Validate validates this d h c p server 1
func (m *DHCPServer1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDynamicLeases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterface(formats); err != nil {
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

	if err := m.validateStaticLeases(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DHCPServer1) validateAvailable(formats strfmt.Registry) error {
	if swag.IsZero(m.Available) { // not required
		return nil
	}

	if err := validate.Minimum("available", "body", *m.Available, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *DHCPServer1) validateDomain(formats strfmt.Registry) error {
	if swag.IsZero(m.Domain) { // not required
		return nil
	}

	if err := validate.MinLength("domain", "body", m.Domain, 1); err != nil {
		return err
	}

	return nil
}

func (m *DHCPServer1) validateDynamicLeases(formats strfmt.Registry) error {
	if swag.IsZero(m.DynamicLeases) { // not required
		return nil
	}

	if err := validate.Minimum("dynamicLeases", "body", *m.DynamicLeases, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *DHCPServer1) validateInterface(formats strfmt.Registry) error {

	if err := validate.Required("interface", "body", m.Interface); err != nil {
		return err
	}

	return nil
}

func (m *DHCPServer1) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	return nil
}

func (m *DHCPServer1) validatePoolSize(formats strfmt.Registry) error {

	if err := validate.Required("poolSize", "body", m.PoolSize); err != nil {
		return err
	}

	if err := validate.Minimum("poolSize", "body", *m.PoolSize, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *DHCPServer1) validateRangeEnd(formats strfmt.Registry) error {

	if err := validate.Required("rangeEnd", "body", m.RangeEnd); err != nil {
		return err
	}

	return nil
}

func (m *DHCPServer1) validateRangeStart(formats strfmt.Registry) error {

	if err := validate.Required("rangeStart", "body", m.RangeStart); err != nil {
		return err
	}

	return nil
}

func (m *DHCPServer1) validateStaticLeases(formats strfmt.Registry) error {
	if swag.IsZero(m.StaticLeases) { // not required
		return nil
	}

	if err := validate.Minimum("staticLeases", "body", *m.StaticLeases, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this d h c p server 1 based on context it is used
func (m *DHCPServer1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DHCPServer1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DHCPServer1) UnmarshalBinary(b []byte) error {
	var res DHCPServer1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
