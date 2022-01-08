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

// UserLogin user login
//
// swagger:model UserLogin
type UserLogin struct {

	// alerts
	// Required: true
	Alerts *bool `json:"alerts"`

	// email
	// Example: email@example.com
	// Required: true
	Email *string `json:"email"`

	// Whether the user can log in or not.
	// Example: true
	// Required: true
	Enabled *bool `json:"enabled"`

	// first name
	// Example: John
	FirstName string `json:"firstName,omitempty"`

	// force change password
	// Required: true
	ForceChangePassword *bool `json:"forceChangePassword"`

	// Where should the user be redirected after login.
	// Example: /nms
	// Required: true
	// Enum: [/nms /crm]
	HomeScreen *string `json:"homeScreen"`

	// id
	// Required: true
	ID *string `json:"id"`

	// image Url
	ImageURL string `json:"imageUrl,omitempty"`

	// is local login enabled
	// Required: true
	IsLocalLoginEnabled *bool `json:"isLocalLoginEnabled"`

	// is sso login enabled
	// Required: true
	IsSsoLoginEnabled *bool `json:"isSsoLoginEnabled"`

	// Whether session should be cleared when browser window is closed.
	// Example: true
	// Required: true
	KeepMeSignedIn *bool `json:"keepMeSignedIn"`

	// last log item Id
	LastLogItemID string `json:"lastLogItemId,omitempty"`

	// last name
	// Example: Doe
	LastName string `json:"lastName,omitempty"`

	// Last news check time in ISO format.
	// Example: 2018-11-14T15:20:32.004Z
	// Required: true
	// Format: date-time
	LastNewsSeenDate *strfmt.DateTime `json:"lastNewsSeenDate"`

	// Last seen version of release notes.
	// Example: 1.0.0-alpha
	// Required: true
	LastReleaseNotesSeen *string `json:"lastReleaseNotesSeen"`

	// location
	Location *UserLocation `json:"location,omitempty"`

	// map config
	MapConfig *MapConfig `json:"mapConfig,omitempty"`

	// preferences
	Preferences Preferences `json:"preferences,omitempty"`

	// Obsolete.
	// Example: false
	PresentationMode bool `json:"presentationMode,omitempty"`

	// Role of this user in UISP. 'anonymous' if user is only allowed to access CRM.
	// Example: admin
	// Enum: [superadmin admin guest anonymous installer]
	Role string `json:"role,omitempty"`

	// Session expiration in milliseconds, max 30 days.
	// Example: 3600000
	// Minimum: 1
	SessionTimeout int64 `json:"sessionTimeout,omitempty"`

	// table config
	TableConfig TableConfig `json:"tableConfig,omitempty"`

	// totp auth enabled
	TotpAuthEnabled bool `json:"totpAuthEnabled,omitempty"`

	// ID of UCRM user that is bound with this user.
	// Example: 123
	// Required: true
	UcrmID *string `json:"ucrmId"`

	// Role of this user in UCRM. 'null' if user is not allowed to access UCRM.
	// Example: 1
	// Required: true
	UcrmRole *string `json:"ucrmRole"`

	// Last time the user was updated.
	// Example: 2018-11-14T15:20:32.004Z
	// Required: true
	// Format: date-time
	Updated *strfmt.DateTime `json:"updated"`

	// username
	// Example: admin
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this user login
func (m *UserLogin) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlerts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForceChangePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHomeScreen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsLocalLoginEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsSsoLoginEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeepMeSignedIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastNewsSeenDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastReleaseNotesSeen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMapConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUcrmID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUcrmRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
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

func (m *UserLogin) validateAlerts(formats strfmt.Registry) error {

	if err := validate.Required("alerts", "body", m.Alerts); err != nil {
		return err
	}

	return nil
}

func (m *UserLogin) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *UserLogin) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *UserLogin) validateForceChangePassword(formats strfmt.Registry) error {

	if err := validate.Required("forceChangePassword", "body", m.ForceChangePassword); err != nil {
		return err
	}

	return nil
}

var userLoginTypeHomeScreenPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["/nms","/crm"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userLoginTypeHomeScreenPropEnum = append(userLoginTypeHomeScreenPropEnum, v)
	}
}

const (

	// UserLoginHomeScreenNms captures enum value "/nms"
	UserLoginHomeScreenNms string = "/nms"

	// UserLoginHomeScreenCrm captures enum value "/crm"
	UserLoginHomeScreenCrm string = "/crm"
)

// prop value enum
func (m *UserLogin) validateHomeScreenEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userLoginTypeHomeScreenPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserLogin) validateHomeScreen(formats strfmt.Registry) error {

	if err := validate.Required("homeScreen", "body", m.HomeScreen); err != nil {
		return err
	}

	// value enum
	if err := m.validateHomeScreenEnum("homeScreen", "body", *m.HomeScreen); err != nil {
		return err
	}

	return nil
}

func (m *UserLogin) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *UserLogin) validateIsLocalLoginEnabled(formats strfmt.Registry) error {

	if err := validate.Required("isLocalLoginEnabled", "body", m.IsLocalLoginEnabled); err != nil {
		return err
	}

	return nil
}

func (m *UserLogin) validateIsSsoLoginEnabled(formats strfmt.Registry) error {

	if err := validate.Required("isSsoLoginEnabled", "body", m.IsSsoLoginEnabled); err != nil {
		return err
	}

	return nil
}

func (m *UserLogin) validateKeepMeSignedIn(formats strfmt.Registry) error {

	if err := validate.Required("keepMeSignedIn", "body", m.KeepMeSignedIn); err != nil {
		return err
	}

	return nil
}

func (m *UserLogin) validateLastNewsSeenDate(formats strfmt.Registry) error {

	if err := validate.Required("lastNewsSeenDate", "body", m.LastNewsSeenDate); err != nil {
		return err
	}

	if err := validate.FormatOf("lastNewsSeenDate", "body", "date-time", m.LastNewsSeenDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserLogin) validateLastReleaseNotesSeen(formats strfmt.Registry) error {

	if err := validate.Required("lastReleaseNotesSeen", "body", m.LastReleaseNotesSeen); err != nil {
		return err
	}

	return nil
}

func (m *UserLogin) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *UserLogin) validateMapConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.MapConfig) { // not required
		return nil
	}

	if m.MapConfig != nil {
		if err := m.MapConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mapConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mapConfig")
			}
			return err
		}
	}

	return nil
}

var userLoginTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["superadmin","admin","guest","anonymous","installer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userLoginTypeRolePropEnum = append(userLoginTypeRolePropEnum, v)
	}
}

const (

	// UserLoginRoleSuperadmin captures enum value "superadmin"
	UserLoginRoleSuperadmin string = "superadmin"

	// UserLoginRoleAdmin captures enum value "admin"
	UserLoginRoleAdmin string = "admin"

	// UserLoginRoleGuest captures enum value "guest"
	UserLoginRoleGuest string = "guest"

	// UserLoginRoleAnonymous captures enum value "anonymous"
	UserLoginRoleAnonymous string = "anonymous"

	// UserLoginRoleInstaller captures enum value "installer"
	UserLoginRoleInstaller string = "installer"
)

// prop value enum
func (m *UserLogin) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userLoginTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserLogin) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

func (m *UserLogin) validateSessionTimeout(formats strfmt.Registry) error {
	if swag.IsZero(m.SessionTimeout) { // not required
		return nil
	}

	if err := validate.MinimumInt("sessionTimeout", "body", m.SessionTimeout, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *UserLogin) validateUcrmID(formats strfmt.Registry) error {

	if err := validate.Required("ucrmId", "body", m.UcrmID); err != nil {
		return err
	}

	return nil
}

func (m *UserLogin) validateUcrmRole(formats strfmt.Registry) error {

	if err := validate.Required("ucrmRole", "body", m.UcrmRole); err != nil {
		return err
	}

	return nil
}

func (m *UserLogin) validateUpdated(formats strfmt.Registry) error {

	if err := validate.Required("updated", "body", m.Updated); err != nil {
		return err
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserLogin) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this user login based on the context it is used
func (m *UserLogin) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMapConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserLogin) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {
		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *UserLogin) contextValidateMapConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.MapConfig != nil {
		if err := m.MapConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mapConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mapConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserLogin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserLogin) UnmarshalBinary(b []byte) error {
	var res UserLogin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}