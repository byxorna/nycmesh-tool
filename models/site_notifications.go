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

// SiteNotifications site notifications
//
// swagger:model SiteNotifications
type SiteNotifications struct {

	// What notification system is used by this site.
	// Example: system
	// Required: true
	// Enum: [system custom none]
	Type *string `json:"type"`

	// users
	// Required: true
	Users Users `json:"users"`
}

// Validate validates this site notifications
func (m *SiteNotifications) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var siteNotificationsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["system","custom","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteNotificationsTypeTypePropEnum = append(siteNotificationsTypeTypePropEnum, v)
	}
}

const (

	// SiteNotificationsTypeSystem captures enum value "system"
	SiteNotificationsTypeSystem string = "system"

	// SiteNotificationsTypeCustom captures enum value "custom"
	SiteNotificationsTypeCustom string = "custom"

	// SiteNotificationsTypeNone captures enum value "none"
	SiteNotificationsTypeNone string = "none"
)

// prop value enum
func (m *SiteNotifications) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, siteNotificationsTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SiteNotifications) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *SiteNotifications) validateUsers(formats strfmt.Registry) error {

	if err := validate.Required("users", "body", m.Users); err != nil {
		return err
	}

	if err := m.Users.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("users")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("users")
		}
		return err
	}

	return nil
}

// ContextValidate validate this site notifications based on the context it is used
func (m *SiteNotifications) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SiteNotifications) contextValidateUsers(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Users.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("users")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("users")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SiteNotifications) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SiteNotifications) UnmarshalBinary(b []byte) error {
	var res SiteNotifications
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
