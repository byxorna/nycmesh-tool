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

// LocalInvitationRequest local invitation request
//
// swagger:model LocalInvitationRequest
type LocalInvitationRequest struct {

	// User invitation token.
	// Required: true
	InvitationToken *string `json:"invitationToken"`

	// New user password.
	// Required: true
	Password *string `json:"password"`
}

// Validate validates this local invitation request
func (m *LocalInvitationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInvitationToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalInvitationRequest) validateInvitationToken(formats strfmt.Registry) error {

	if err := validate.Required("invitationToken", "body", m.InvitationToken); err != nil {
		return err
	}

	return nil
}

func (m *LocalInvitationRequest) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this local invitation request based on context it is used
func (m *LocalInvitationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LocalInvitationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocalInvitationRequest) UnmarshalBinary(b []byte) error {
	var res LocalInvitationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
