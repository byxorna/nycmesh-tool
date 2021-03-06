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

// DataLinkDevice data link device
//
// swagger:model DataLinkDevice
type DataLinkDevice struct {

	// attributes
	// Required: true
	Attributes *Attributes `json:"attributes"`

	// frequency
	Frequency float64 `json:"frequency,omitempty"`

	// identification
	Identification *Identification1 `json:"identification,omitempty"`

	// meta
	Meta Meta `json:"meta,omitempty"`

	// overview
	Overview *Overview `json:"overview,omitempty"`

	// signal
	Signal float64 `json:"signal,omitempty"`

	// ssid
	Ssid string `json:"ssid,omitempty"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this data link device
func (m *DataLinkDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverview(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataLinkDevice) validateAttributes(formats strfmt.Registry) error {

	if err := validate.Required("attributes", "body", m.Attributes); err != nil {
		return err
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

func (m *DataLinkDevice) validateIdentification(formats strfmt.Registry) error {
	if swag.IsZero(m.Identification) { // not required
		return nil
	}

	if m.Identification != nil {
		if err := m.Identification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identification")
			}
			return err
		}
	}

	return nil
}

func (m *DataLinkDevice) validateOverview(formats strfmt.Registry) error {
	if swag.IsZero(m.Overview) { // not required
		return nil
	}

	if m.Overview != nil {
		if err := m.Overview.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("overview")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("overview")
			}
			return err
		}
	}

	return nil
}

func (m *DataLinkDevice) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this data link device based on the context it is used
func (m *DataLinkDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentification(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOverview(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataLinkDevice) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	if m.Attributes != nil {
		if err := m.Attributes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

func (m *DataLinkDevice) contextValidateIdentification(ctx context.Context, formats strfmt.Registry) error {

	if m.Identification != nil {
		if err := m.Identification.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identification")
			}
			return err
		}
	}

	return nil
}

func (m *DataLinkDevice) contextValidateOverview(ctx context.Context, formats strfmt.Registry) error {

	if m.Overview != nil {
		if err := m.Overview.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("overview")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("overview")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataLinkDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataLinkDevice) UnmarshalBinary(b []byte) error {
	var res DataLinkDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
