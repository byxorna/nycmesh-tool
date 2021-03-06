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

// Outages outages
//
// swagger:model Outages
type Outages struct {

	// aggregation
	Aggregation *Aggregation1 `json:"aggregation,omitempty"`

	// items
	Items OutageList `json:"items,omitempty"`

	// pagination
	Pagination *Pagination `json:"pagination,omitempty"`
}

// Validate validates this outages
func (m *Outages) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Outages) validateAggregation(formats strfmt.Registry) error {
	if swag.IsZero(m.Aggregation) { // not required
		return nil
	}

	if m.Aggregation != nil {
		if err := m.Aggregation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aggregation")
			}
			return err
		}
	}

	return nil
}

func (m *Outages) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(m.Items) { // not required
		return nil
	}

	if err := m.Items.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("items")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("items")
		}
		return err
	}

	return nil
}

func (m *Outages) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this outages based on the context it is used
func (m *Outages) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAggregation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Outages) contextValidateAggregation(ctx context.Context, formats strfmt.Registry) error {

	if m.Aggregation != nil {
		if err := m.Aggregation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aggregation")
			}
			return err
		}
	}

	return nil
}

func (m *Outages) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Items.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("items")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("items")
		}
		return err
	}

	return nil
}

func (m *Outages) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if m.Pagination != nil {
		if err := m.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Outages) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Outages) UnmarshalBinary(b []byte) error {
	var res Outages
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
