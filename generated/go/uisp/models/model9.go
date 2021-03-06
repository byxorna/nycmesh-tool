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

// Model9 model 9
//
// swagger:model Model 9
type Model9 struct {

	// device
	Device *DeviceIdentification2 `json:"device,omitempty"`

	// device metadata
	DeviceMetadata *DeviceMetadata `json:"deviceMetadata,omitempty"`

	// Log item id.
	// Example: f7ac9cad-ea28-4390-93c8-7add010e8ee3
	// Required: true
	ID *string `json:"id"`

	// Log severity level
	// Example: info
	// Required: true
	// Enum: [info warning error]
	Level *string `json:"level"`

	// Log message
	// Example: lab-AF5XHD-1 MAC: 78:8a:20:5f:2a:ff, IP: 10.43.21.41 was deleted.
	// Required: true
	Message *string `json:"message"`

	// site
	Site *Site1 `json:"site,omitempty"`

	// tags
	// Required: true
	Tags Tags `json:"tags"`

	// Date and time of the log creation.
	// Example: 2020-05-12T10:19:10.056Z
	// Required: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp"`
}

// Validate validates this model 9
func (m *Model9) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model9) validateDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.Device) { // not required
		return nil
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *Model9) validateDeviceMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.DeviceMetadata) { // not required
		return nil
	}

	if m.DeviceMetadata != nil {
		if err := m.DeviceMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deviceMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deviceMetadata")
			}
			return err
		}
	}

	return nil
}

func (m *Model9) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var model9TypeLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["info","warning","error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		model9TypeLevelPropEnum = append(model9TypeLevelPropEnum, v)
	}
}

const (

	// Model9LevelInfo captures enum value "info"
	Model9LevelInfo string = "info"

	// Model9LevelWarning captures enum value "warning"
	Model9LevelWarning string = "warning"

	// Model9LevelError captures enum value "error"
	Model9LevelError string = "error"
)

// prop value enum
func (m *Model9) validateLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, model9TypeLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Model9) validateLevel(formats strfmt.Registry) error {

	if err := validate.Required("level", "body", m.Level); err != nil {
		return err
	}

	// value enum
	if err := m.validateLevelEnum("level", "body", *m.Level); err != nil {
		return err
	}

	return nil
}

func (m *Model9) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *Model9) validateSite(formats strfmt.Registry) error {
	if swag.IsZero(m.Site) { // not required
		return nil
	}

	if m.Site != nil {
		if err := m.Site.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

func (m *Model9) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("tags", "body", m.Tags); err != nil {
		return err
	}

	if err := m.Tags.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tags")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("tags")
		}
		return err
	}

	return nil
}

func (m *Model9) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this model 9 based on the context it is used
func (m *Model9) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model9) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.Device != nil {
		if err := m.Device.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *Model9) contextValidateDeviceMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.DeviceMetadata != nil {
		if err := m.DeviceMetadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deviceMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deviceMetadata")
			}
			return err
		}
	}

	return nil
}

func (m *Model9) contextValidateSite(ctx context.Context, formats strfmt.Registry) error {

	if m.Site != nil {
		if err := m.Site.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

func (m *Model9) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Tags.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tags")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("tags")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Model9) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model9) UnmarshalBinary(b []byte) error {
	var res Model9
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
