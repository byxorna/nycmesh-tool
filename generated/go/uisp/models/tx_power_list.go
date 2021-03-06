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

// TxPowerList tx power list
//
// swagger:model TxPowerList
type TxPowerList struct {

	// radio2 ghz tx power list
	// Required: true
	Radio2GhzTxPowerList Radio2GhzTxPowerList `json:"radio2GhzTxPowerList"`

	// radio5 ghz tx power list
	// Required: true
	Radio5GhzTxPowerList Radio5GhzTxPowerList `json:"radio5GhzTxPowerList"`
}

// Validate validates this tx power list
func (m *TxPowerList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRadio2GhzTxPowerList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRadio5GhzTxPowerList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TxPowerList) validateRadio2GhzTxPowerList(formats strfmt.Registry) error {

	if err := validate.Required("radio2GhzTxPowerList", "body", m.Radio2GhzTxPowerList); err != nil {
		return err
	}

	if err := m.Radio2GhzTxPowerList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("radio2GhzTxPowerList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("radio2GhzTxPowerList")
		}
		return err
	}

	return nil
}

func (m *TxPowerList) validateRadio5GhzTxPowerList(formats strfmt.Registry) error {

	if err := validate.Required("radio5GhzTxPowerList", "body", m.Radio5GhzTxPowerList); err != nil {
		return err
	}

	if err := m.Radio5GhzTxPowerList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("radio5GhzTxPowerList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("radio5GhzTxPowerList")
		}
		return err
	}

	return nil
}

// ContextValidate validate this tx power list based on the context it is used
func (m *TxPowerList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRadio2GhzTxPowerList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRadio5GhzTxPowerList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TxPowerList) contextValidateRadio2GhzTxPowerList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Radio2GhzTxPowerList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("radio2GhzTxPowerList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("radio2GhzTxPowerList")
		}
		return err
	}

	return nil
}

func (m *TxPowerList) contextValidateRadio5GhzTxPowerList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Radio5GhzTxPowerList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("radio5GhzTxPowerList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("radio5GhzTxPowerList")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TxPowerList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TxPowerList) UnmarshalBinary(b []byte) error {
	var res TxPowerList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
