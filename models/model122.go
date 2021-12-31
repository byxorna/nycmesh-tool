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

// Model122 model 122
//
// swagger:model Model 122
type Model122 struct {

	// redistribute connected
	// Required: true
	RedistributeConnected *RedistributeConnected `json:"redistributeConnected"`

	// redistribute default route
	// Required: true
	RedistributeDefaultRoute *RedistributeDefaultRoute3 `json:"redistributeDefaultRoute"`

	// redistribute static
	// Required: true
	RedistributeStatic *RedistributeStatic `json:"redistributeStatic"`

	// router
	// Required: true
	Router *string `json:"router"`
}

// Validate validates this model 122
func (m *Model122) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRedistributeConnected(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedistributeDefaultRoute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedistributeStatic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model122) validateRedistributeConnected(formats strfmt.Registry) error {

	if err := validate.Required("redistributeConnected", "body", m.RedistributeConnected); err != nil {
		return err
	}

	if m.RedistributeConnected != nil {
		if err := m.RedistributeConnected.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redistributeConnected")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("redistributeConnected")
			}
			return err
		}
	}

	return nil
}

func (m *Model122) validateRedistributeDefaultRoute(formats strfmt.Registry) error {

	if err := validate.Required("redistributeDefaultRoute", "body", m.RedistributeDefaultRoute); err != nil {
		return err
	}

	if m.RedistributeDefaultRoute != nil {
		if err := m.RedistributeDefaultRoute.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redistributeDefaultRoute")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("redistributeDefaultRoute")
			}
			return err
		}
	}

	return nil
}

func (m *Model122) validateRedistributeStatic(formats strfmt.Registry) error {

	if err := validate.Required("redistributeStatic", "body", m.RedistributeStatic); err != nil {
		return err
	}

	if m.RedistributeStatic != nil {
		if err := m.RedistributeStatic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redistributeStatic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("redistributeStatic")
			}
			return err
		}
	}

	return nil
}

func (m *Model122) validateRouter(formats strfmt.Registry) error {

	if err := validate.Required("router", "body", m.Router); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this model 122 based on the context it is used
func (m *Model122) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRedistributeConnected(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRedistributeDefaultRoute(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRedistributeStatic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model122) contextValidateRedistributeConnected(ctx context.Context, formats strfmt.Registry) error {

	if m.RedistributeConnected != nil {
		if err := m.RedistributeConnected.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redistributeConnected")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("redistributeConnected")
			}
			return err
		}
	}

	return nil
}

func (m *Model122) contextValidateRedistributeDefaultRoute(ctx context.Context, formats strfmt.Registry) error {

	if m.RedistributeDefaultRoute != nil {
		if err := m.RedistributeDefaultRoute.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redistributeDefaultRoute")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("redistributeDefaultRoute")
			}
			return err
		}
	}

	return nil
}

func (m *Model122) contextValidateRedistributeStatic(ctx context.Context, formats strfmt.Registry) error {

	if m.RedistributeStatic != nil {
		if err := m.RedistributeStatic.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redistributeStatic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("redistributeStatic")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Model122) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model122) UnmarshalBinary(b []byte) error {
	var res Model122
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
