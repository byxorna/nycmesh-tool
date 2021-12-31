// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Coverage List of modulation values in Mbps.
// Example: [[100,110,111],[100,111,112],[110,120,122]]
//
// swagger:model coverage
type Coverage []Model76

// Validate validates this coverage
func (m Coverage) Validate(formats strfmt.Registry) error {
	var res []error

	iCoverageSize := int64(len(m))

	if err := validate.MinItems("", "body", iCoverageSize, 0); err != nil {
		return err
	}

	for i := 0; i < len(m); i++ {

		if err := m[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName(strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName(strconv.Itoa(i))
			}
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this coverage based on the context it is used
func (m Coverage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		if err := m[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName(strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName(strconv.Itoa(i))
			}
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}