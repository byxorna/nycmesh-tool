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

// Pairs pairs
//
// swagger:model pairs
type Pairs struct {

	// pair a
	PairA *PairA `json:"pairA,omitempty"`

	// pair b
	PairB *PairB `json:"pairB,omitempty"`

	// pair c
	PairC *PairC `json:"pairC,omitempty"`

	// pair d
	PairD *PairD `json:"pairD,omitempty"`
}

// Validate validates this pairs
func (m *Pairs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePairA(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePairB(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePairC(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePairD(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Pairs) validatePairA(formats strfmt.Registry) error {
	if swag.IsZero(m.PairA) { // not required
		return nil
	}

	if m.PairA != nil {
		if err := m.PairA.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pairA")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pairA")
			}
			return err
		}
	}

	return nil
}

func (m *Pairs) validatePairB(formats strfmt.Registry) error {
	if swag.IsZero(m.PairB) { // not required
		return nil
	}

	if m.PairB != nil {
		if err := m.PairB.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pairB")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pairB")
			}
			return err
		}
	}

	return nil
}

func (m *Pairs) validatePairC(formats strfmt.Registry) error {
	if swag.IsZero(m.PairC) { // not required
		return nil
	}

	if m.PairC != nil {
		if err := m.PairC.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pairC")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pairC")
			}
			return err
		}
	}

	return nil
}

func (m *Pairs) validatePairD(formats strfmt.Registry) error {
	if swag.IsZero(m.PairD) { // not required
		return nil
	}

	if m.PairD != nil {
		if err := m.PairD.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pairD")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pairD")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this pairs based on the context it is used
func (m *Pairs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePairA(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePairB(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePairC(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePairD(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Pairs) contextValidatePairA(ctx context.Context, formats strfmt.Registry) error {

	if m.PairA != nil {
		if err := m.PairA.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pairA")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pairA")
			}
			return err
		}
	}

	return nil
}

func (m *Pairs) contextValidatePairB(ctx context.Context, formats strfmt.Registry) error {

	if m.PairB != nil {
		if err := m.PairB.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pairB")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pairB")
			}
			return err
		}
	}

	return nil
}

func (m *Pairs) contextValidatePairC(ctx context.Context, formats strfmt.Registry) error {

	if m.PairC != nil {
		if err := m.PairC.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pairC")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pairC")
			}
			return err
		}
	}

	return nil
}

func (m *Pairs) contextValidatePairD(ctx context.Context, formats strfmt.Registry) error {

	if m.PairD != nil {
		if err := m.PairD.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pairD")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pairD")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Pairs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Pairs) UnmarshalBinary(b []byte) error {
	var res Pairs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
