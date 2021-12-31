// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DeviceIds5 device ids 5
//
// swagger:model deviceIds 5
type DeviceIds5 []string

// Validate validates this device ids 5
func (m DeviceIds5) Validate(formats strfmt.Registry) error {
	var res []error

	iDeviceIds5Size := int64(len(m))

	if err := validate.MinItems("", "body", iDeviceIds5Size, 1); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this device ids 5 based on context it is used
func (m DeviceIds5) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}