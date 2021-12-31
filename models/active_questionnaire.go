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

// ActiveQuestionnaire active questionnaire
//
// swagger:model ActiveQuestionnaire
type ActiveQuestionnaire struct {

	// ID of the currently active questionnaire. Null if no questionnaire is active.
	// Example: Q1
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this active questionnaire
func (m *ActiveQuestionnaire) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActiveQuestionnaire) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this active questionnaire based on context it is used
func (m *ActiveQuestionnaire) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ActiveQuestionnaire) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActiveQuestionnaire) UnmarshalBinary(b []byte) error {
	var res ActiveQuestionnaire
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
