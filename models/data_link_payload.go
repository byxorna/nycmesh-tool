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

// DataLinkPayload data link payload
//
// swagger:model DataLinkPayload
type DataLinkPayload struct {

	// device Id from
	// Required: true
	DeviceIDFrom *string `json:"deviceIdFrom"`

	// device Id to
	// Required: true
	DeviceIDTo *string `json:"deviceIdTo"`

	// id
	ID string `json:"id,omitempty"`

	// interface Id from
	// Required: true
	InterfaceIDFrom *string `json:"interfaceIdFrom"`

	// interface Id to
	// Required: true
	InterfaceIDTo *string `json:"interfaceIdTo"`
}

// Validate validates this data link payload
func (m *DataLinkPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceIDFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceIDTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterfaceIDFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterfaceIDTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataLinkPayload) validateDeviceIDFrom(formats strfmt.Registry) error {

	if err := validate.Required("deviceIdFrom", "body", m.DeviceIDFrom); err != nil {
		return err
	}

	return nil
}

func (m *DataLinkPayload) validateDeviceIDTo(formats strfmt.Registry) error {

	if err := validate.Required("deviceIdTo", "body", m.DeviceIDTo); err != nil {
		return err
	}

	return nil
}

func (m *DataLinkPayload) validateInterfaceIDFrom(formats strfmt.Registry) error {

	if err := validate.Required("interfaceIdFrom", "body", m.InterfaceIDFrom); err != nil {
		return err
	}

	return nil
}

func (m *DataLinkPayload) validateInterfaceIDTo(formats strfmt.Registry) error {

	if err := validate.Required("interfaceIdTo", "body", m.InterfaceIDTo); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this data link payload based on context it is used
func (m *DataLinkPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataLinkPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataLinkPayload) UnmarshalBinary(b []byte) error {
	var res DataLinkPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}