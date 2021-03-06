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

// Model86 model 86
//
// swagger:model Model 86
type Model86 struct {

	// Timestamp in milliseconds when the setup was initiated.
	// Example: 1557222409000
	// Minimum: 0
	StartedAt *int64 `json:"startedAt,omitempty"`
}

// Validate validates this model 86
func (m *Model86) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model86) validateStartedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if err := validate.MinimumInt("startedAt", "body", *m.StartedAt, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this model 86 based on context it is used
func (m *Model86) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Model86) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model86) UnmarshalBinary(b []byte) error {
	var res Model86
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
