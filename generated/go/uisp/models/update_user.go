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

// UpdateUser update user
//
// swagger:model UpdateUser
type UpdateUser struct {

	// Whether to send email alerts or not.
	// Example: true
	// Required: true
	Alerts *bool `json:"alerts"`

	// User's email.
	// Example: admin@example.com
	// Required: true
	Email *string `json:"email"`

	// Whether user is allowed to login or not.
	// Example: true
	Enabled bool `json:"enabled,omitempty"`

	// User's first name.
	// Example: John
	FirstName string `json:"firstName,omitempty"`

	// home screen
	// Example: /nms
	// Enum: [/nms /crm]
	HomeScreen string `json:"homeScreen,omitempty"`

	// id
	// Example: f7ac9cad-ea28-4390-93c8-7add010e8ee3
	// Required: true
	ID *string `json:"id"`

	// User's last name.
	// Example: Doe
	LastName string `json:"lastName,omitempty"`

	// Users's role in Network.
	// Example: admin
	// Required: true
	// Enum: [superadmin admin guest anonymous installer]
	Role *string `json:"role"`

	// Whether two-factor authentication is enabled or not.
	// Example: false
	// Required: true
	TotpAuthEnabled *bool `json:"totpAuthEnabled"`

	// Ignored, kept for backward compatibility.
	UcrmID string `json:"ucrmId,omitempty"`

	// Users's role ID in CRM.
	UcrmRole string `json:"ucrmRole,omitempty"`

	// Username used for login.
	// Example: admin
	// Required: true
	// Max Length: 320
	// Min Length: 1
	// Pattern: ^[a-zA-Z0-9_]*$
	Username *string `json:"username"`
}

// Validate validates this update user
func (m *UpdateUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlerts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHomeScreen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotpAuthEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateUser) validateAlerts(formats strfmt.Registry) error {

	if err := validate.Required("alerts", "body", m.Alerts); err != nil {
		return err
	}

	return nil
}

func (m *UpdateUser) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

var updateUserTypeHomeScreenPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["/nms","/crm"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateUserTypeHomeScreenPropEnum = append(updateUserTypeHomeScreenPropEnum, v)
	}
}

const (

	// UpdateUserHomeScreenNms captures enum value "/nms"
	UpdateUserHomeScreenNms string = "/nms"

	// UpdateUserHomeScreenCrm captures enum value "/crm"
	UpdateUserHomeScreenCrm string = "/crm"
)

// prop value enum
func (m *UpdateUser) validateHomeScreenEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateUserTypeHomeScreenPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateUser) validateHomeScreen(formats strfmt.Registry) error {
	if swag.IsZero(m.HomeScreen) { // not required
		return nil
	}

	// value enum
	if err := m.validateHomeScreenEnum("homeScreen", "body", m.HomeScreen); err != nil {
		return err
	}

	return nil
}

func (m *UpdateUser) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var updateUserTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["superadmin","admin","guest","anonymous","installer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateUserTypeRolePropEnum = append(updateUserTypeRolePropEnum, v)
	}
}

const (

	// UpdateUserRoleSuperadmin captures enum value "superadmin"
	UpdateUserRoleSuperadmin string = "superadmin"

	// UpdateUserRoleAdmin captures enum value "admin"
	UpdateUserRoleAdmin string = "admin"

	// UpdateUserRoleGuest captures enum value "guest"
	UpdateUserRoleGuest string = "guest"

	// UpdateUserRoleAnonymous captures enum value "anonymous"
	UpdateUserRoleAnonymous string = "anonymous"

	// UpdateUserRoleInstaller captures enum value "installer"
	UpdateUserRoleInstaller string = "installer"
)

// prop value enum
func (m *UpdateUser) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateUserTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateUser) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", *m.Role); err != nil {
		return err
	}

	return nil
}

func (m *UpdateUser) validateTotpAuthEnabled(formats strfmt.Registry) error {

	if err := validate.Required("totpAuthEnabled", "body", m.TotpAuthEnabled); err != nil {
		return err
	}

	return nil
}

func (m *UpdateUser) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	if err := validate.MinLength("username", "body", *m.Username, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("username", "body", *m.Username, 320); err != nil {
		return err
	}

	if err := validate.Pattern("username", "body", *m.Username, `^[a-zA-Z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update user based on context it is used
func (m *UpdateUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateUser) UnmarshalBinary(b []byte) error {
	var res UpdateUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
