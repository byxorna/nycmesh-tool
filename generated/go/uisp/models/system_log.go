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

// SystemLog system log
//
// swagger:model systemLog
type SystemLog struct {

	// Set to true to enable sending syslog to remote logging server.
	// Example: false
	Enabled bool `json:"enabled,omitempty"`

	// Level of logs to include.
	// Enum: [emerg alert crit err warning notice info debug]
	Level string `json:"level,omitempty"`

	// Remote server port.
	// Example: 443
	// Maximum: 65535
	// Minimum: 0
	Port *int64 `json:"port,omitempty"`

	// Remote server hostname.
	// Example: 3.3.3.3
	Server string `json:"server,omitempty"`
}

// Validate validates this system log
func (m *SystemLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var systemLogTypeLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["emerg","alert","crit","err","warning","notice","info","debug"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		systemLogTypeLevelPropEnum = append(systemLogTypeLevelPropEnum, v)
	}
}

const (

	// SystemLogLevelEmerg captures enum value "emerg"
	SystemLogLevelEmerg string = "emerg"

	// SystemLogLevelAlert captures enum value "alert"
	SystemLogLevelAlert string = "alert"

	// SystemLogLevelCrit captures enum value "crit"
	SystemLogLevelCrit string = "crit"

	// SystemLogLevelErr captures enum value "err"
	SystemLogLevelErr string = "err"

	// SystemLogLevelWarning captures enum value "warning"
	SystemLogLevelWarning string = "warning"

	// SystemLogLevelNotice captures enum value "notice"
	SystemLogLevelNotice string = "notice"

	// SystemLogLevelInfo captures enum value "info"
	SystemLogLevelInfo string = "info"

	// SystemLogLevelDebug captures enum value "debug"
	SystemLogLevelDebug string = "debug"
)

// prop value enum
func (m *SystemLog) validateLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, systemLogTypeLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SystemLog) validateLevel(formats strfmt.Registry) error {
	if swag.IsZero(m.Level) { // not required
		return nil
	}

	// value enum
	if err := m.validateLevelEnum("level", "body", m.Level); err != nil {
		return err
	}

	return nil
}

func (m *SystemLog) validatePort(formats strfmt.Registry) error {
	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("port", "body", *m.Port, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", *m.Port, 65535, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this system log based on context it is used
func (m *SystemLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SystemLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemLog) UnmarshalBinary(b []byte) error {
	var res SystemLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
