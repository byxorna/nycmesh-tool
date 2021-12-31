// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// WirelessActiveInterfaceIds Main active radios currently used to send/receive traffic.
// Example: ["wlan0"]
//
// swagger:model wirelessActiveInterfaceIds
type WirelessActiveInterfaceIds []string

// Validate validates this wireless active interface ids
func (m WirelessActiveInterfaceIds) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this wireless active interface ids based on context it is used
func (m WirelessActiveInterfaceIds) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
