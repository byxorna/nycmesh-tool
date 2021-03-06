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

// DeviceOnu device onu
//
// swagger:model DeviceOnu
type DeviceOnu struct {

	// id
	// Example: f7ac9cad-ea28-4390-93c8-7add010e8ee3
	ID string `json:"id,omitempty"`

	// port
	// Maximum: 65535
	// Minimum: 1
	Port int64 `json:"port,omitempty"`

	// port forwards
	PortForwards string `json:"portForwards,omitempty"`

	// pppoe user
	PppoeUser string `json:"pppoeUser,omitempty"`

	// profile
	Profile string `json:"profile,omitempty"`

	// profile name
	ProfileName string `json:"profileName,omitempty"`

	// receive power
	ReceivePower float64 `json:"receivePower,omitempty"`

	// rx rate
	RxRate float64 `json:"rxRate,omitempty"`

	// transmit power
	TransmitPower float64 `json:"transmitPower,omitempty"`

	// tx rate
	TxRate float64 `json:"txRate,omitempty"`

	// wan address
	WanAddress string `json:"wanAddress,omitempty"`
}

// Validate validates this device onu
func (m *DeviceOnu) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceOnu) validatePort(formats strfmt.Registry) error {
	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("port", "body", m.Port, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", m.Port, 65535, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this device onu based on context it is used
func (m *DeviceOnu) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceOnu) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceOnu) UnmarshalBinary(b []byte) error {
	var res DeviceOnu
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
