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

// Ports2 ports 2
//
// swagger:model ports 2
type Ports2 []int64

// Validate validates this ports 2
func (m Ports2) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		if err := validate.MinimumInt(strconv.Itoa(i), "body", m[i], 1, false); err != nil {
			return err
		}

		if err := validate.MaximumInt(strconv.Itoa(i), "body", m[i], 3, false); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this ports 2 based on context it is used
func (m Ports2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}