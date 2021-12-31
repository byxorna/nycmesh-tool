// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RouterOspf router ospf
//
// swagger:model RouterOspf
type RouterOspf struct {

	// redistribute connected
	RedistributeConnected *RedistributeConnected `json:"redistributeConnected,omitempty"`

	// redistribute default route
	RedistributeDefaultRoute *RedistributeDefaultRoute `json:"redistributeDefaultRoute,omitempty"`

	// redistribute static
	RedistributeStatic *RedistributeStatic `json:"redistributeStatic,omitempty"`

	// router
	Router string `json:"router,omitempty"`
}

// Validate validates this router ospf
func (m *RouterOspf) Validate(formats strfmt.Registry) error {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RouterOspf) validateRedistributeConnected(formats strfmt.Registry) error {
	if swag.IsZero(m.RedistributeConnected) { // not required
		return nil
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

func (m *RouterOspf) validateRedistributeDefaultRoute(formats strfmt.Registry) error {
	if swag.IsZero(m.RedistributeDefaultRoute) { // not required
		return nil
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

func (m *RouterOspf) validateRedistributeStatic(formats strfmt.Registry) error {
	if swag.IsZero(m.RedistributeStatic) { // not required
		return nil
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

// ContextValidate validate this router ospf based on the context it is used
func (m *RouterOspf) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *RouterOspf) contextValidateRedistributeConnected(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RouterOspf) contextValidateRedistributeDefaultRoute(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RouterOspf) contextValidateRedistributeStatic(ctx context.Context, formats strfmt.Registry) error {

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
func (m *RouterOspf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RouterOspf) UnmarshalBinary(b []byte) error {
	var res RouterOspf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
