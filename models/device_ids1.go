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

// DeviceIds1 device ids 1
//
// swagger:model deviceIds 1
type DeviceIds1 []string

// Validate validates this device ids 1
func (m DeviceIds1) Validate(formats strfmt.Registry) error {
	var res []error

	iDeviceIds1Size := int64(len(m))

	if err := validate.MinItems("", "body", iDeviceIds1Size, 1); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this device ids 1 based on context it is used
func (m DeviceIds1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}