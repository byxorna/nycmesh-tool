// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DataLink data link
//
// swagger:model DataLink
type DataLink struct {

	// can delete
	// Required: true
	CanDelete *bool `json:"canDelete"`

	// Nullable property in meters.
	DevicesDistance float64 `json:"devicesDistance,omitempty"`

	// frequency
	// Required: true
	Frequency *float64 `json:"frequency"`

	// from
	From *From `json:"from,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// Generated by UISP or manually added
	// Required: true
	// Enum: [unms manual]
	Origin *string `json:"origin"`

	// signal
	// Required: true
	Signal *float64 `json:"signal"`

	// Nullable property in meters.
	SitesDistance float64 `json:"sitesDistance,omitempty"`

	// ssid
	// Required: true
	Ssid *string `json:"ssid"`

	// state
	// Enum: [active disabled disconnected proposed]
	State string `json:"state,omitempty"`

	// to
	To *To `json:"to,omitempty"`

	// type
	// Enum: [ethernet pon wireless]
	Type string `json:"type,omitempty"`
}

// Validate validates this data link
func (m *DataLink) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanDelete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataLink) validateCanDelete(formats strfmt.Registry) error {

	if err := validate.Required("canDelete", "body", m.CanDelete); err != nil {
		return err
	}

	return nil
}

func (m *DataLink) validateFrequency(formats strfmt.Registry) error {

	if err := validate.Required("frequency", "body", m.Frequency); err != nil {
		return err
	}

	return nil
}

func (m *DataLink) validateFrom(formats strfmt.Registry) error {
	if swag.IsZero(m.From) { // not required
		return nil
	}

	if m.From != nil {
		if err := m.From.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("from")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("from")
			}
			return err
		}
	}

	return nil
}

func (m *DataLink) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var dataLinkTypeOriginPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unms","manual"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dataLinkTypeOriginPropEnum = append(dataLinkTypeOriginPropEnum, v)
	}
}

const (

	// DataLinkOriginUnms captures enum value "unms"
	DataLinkOriginUnms string = "unms"

	// DataLinkOriginManual captures enum value "manual"
	DataLinkOriginManual string = "manual"
)

// prop value enum
func (m *DataLink) validateOriginEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dataLinkTypeOriginPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DataLink) validateOrigin(formats strfmt.Registry) error {

	if err := validate.Required("origin", "body", m.Origin); err != nil {
		return err
	}

	// value enum
	if err := m.validateOriginEnum("origin", "body", *m.Origin); err != nil {
		return err
	}

	return nil
}

func (m *DataLink) validateSignal(formats strfmt.Registry) error {

	if err := validate.Required("signal", "body", m.Signal); err != nil {
		return err
	}

	return nil
}

func (m *DataLink) validateSsid(formats strfmt.Registry) error {

	if err := validate.Required("ssid", "body", m.Ssid); err != nil {
		return err
	}

	return nil
}

var dataLinkTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","disabled","disconnected","proposed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dataLinkTypeStatePropEnum = append(dataLinkTypeStatePropEnum, v)
	}
}

const (

	// DataLinkStateActive captures enum value "active"
	DataLinkStateActive string = "active"

	// DataLinkStateDisabled captures enum value "disabled"
	DataLinkStateDisabled string = "disabled"

	// DataLinkStateDisconnected captures enum value "disconnected"
	DataLinkStateDisconnected string = "disconnected"

	// DataLinkStateProposed captures enum value "proposed"
	DataLinkStateProposed string = "proposed"
)

// prop value enum
func (m *DataLink) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dataLinkTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DataLink) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *DataLink) validateTo(formats strfmt.Registry) error {
	if swag.IsZero(m.To) { // not required
		return nil
	}

	if m.To != nil {
		if err := m.To.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("to")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("to")
			}
			return err
		}
	}

	return nil
}

var dataLinkTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ethernet","pon","wireless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dataLinkTypeTypePropEnum = append(dataLinkTypeTypePropEnum, v)
	}
}

const (

	// DataLinkTypeEthernet captures enum value "ethernet"
	DataLinkTypeEthernet string = "ethernet"

	// DataLinkTypePon captures enum value "pon"
	DataLinkTypePon string = "pon"

	// DataLinkTypeWireless captures enum value "wireless"
	DataLinkTypeWireless string = "wireless"
)

// prop value enum
func (m *DataLink) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dataLinkTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DataLink) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this data link based on the context it is used
func (m *DataLink) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFrom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataLink) contextValidateFrom(ctx context.Context, formats strfmt.Registry) error {

	if m.From != nil {
		if err := m.From.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("from")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("from")
			}
			return err
		}
	}

	return nil
}

func (m *DataLink) contextValidateTo(ctx context.Context, formats strfmt.Registry) error {

	if m.To != nil {
		if err := m.To.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("to")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("to")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataLink) UnmarshalBinary(b []byte) error {
	var res DataLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}