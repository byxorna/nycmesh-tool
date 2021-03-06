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

// Model57 model 57
//
// swagger:model Model 57
type Model57 struct {

	// bytes
	// Required: true
	// Minimum: 0
	Bytes *int64 `json:"bytes"`

	// chain
	// Required: true
	Chain *string `json:"chain"`

	// description
	Description string `json:"description,omitempty"`

	// destination
	Destination *Destination `json:"destination,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// id
	// Minimum: 0
	ID *int64 `json:"id,omitempty"`

	// in interface
	InInterface *InInterface `json:"inInterface,omitempty"`

	// ip version
	// Required: true
	// Enum: [both v4only v6only]
	IPVersion *string `json:"ipVersion"`

	// log
	Log bool `json:"log,omitempty"`

	// negate in interface
	NegateInInterface bool `json:"negateInInterface,omitempty"`

	// negate out interface
	NegateOutInterface bool `json:"negateOutInterface,omitempty"`

	// out interface
	OutInterface *OutInterface `json:"outInterface,omitempty"`

	// packets
	// Required: true
	// Minimum: 0
	Packets *int64 `json:"packets"`

	// protocol
	Protocol string `json:"protocol,omitempty"`

	// source
	Source *Source `json:"source,omitempty"`

	// target
	// Required: true
	Target *string `json:"target"`

	// translation
	Translation *Translation `json:"translation,omitempty"`
}

// Validate validates this model 57
func (m *Model57) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBytes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInInterface(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutInterface(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTranslation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model57) validateBytes(formats strfmt.Registry) error {

	if err := validate.Required("bytes", "body", m.Bytes); err != nil {
		return err
	}

	if err := validate.MinimumInt("bytes", "body", *m.Bytes, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Model57) validateChain(formats strfmt.Registry) error {

	if err := validate.Required("chain", "body", m.Chain); err != nil {
		return err
	}

	return nil
}

func (m *Model57) validateDestination(formats strfmt.Registry) error {
	if swag.IsZero(m.Destination) { // not required
		return nil
	}

	if m.Destination != nil {
		if err := m.Destination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destination")
			}
			return err
		}
	}

	return nil
}

func (m *Model57) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinimumInt("id", "body", *m.ID, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Model57) validateInInterface(formats strfmt.Registry) error {
	if swag.IsZero(m.InInterface) { // not required
		return nil
	}

	if m.InInterface != nil {
		if err := m.InInterface.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inInterface")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inInterface")
			}
			return err
		}
	}

	return nil
}

var model57TypeIPVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["both","v4only","v6only"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		model57TypeIPVersionPropEnum = append(model57TypeIPVersionPropEnum, v)
	}
}

const (

	// Model57IPVersionBoth captures enum value "both"
	Model57IPVersionBoth string = "both"

	// Model57IPVersionV4only captures enum value "v4only"
	Model57IPVersionV4only string = "v4only"

	// Model57IPVersionV6only captures enum value "v6only"
	Model57IPVersionV6only string = "v6only"
)

// prop value enum
func (m *Model57) validateIPVersionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, model57TypeIPVersionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Model57) validateIPVersion(formats strfmt.Registry) error {

	if err := validate.Required("ipVersion", "body", m.IPVersion); err != nil {
		return err
	}

	// value enum
	if err := m.validateIPVersionEnum("ipVersion", "body", *m.IPVersion); err != nil {
		return err
	}

	return nil
}

func (m *Model57) validateOutInterface(formats strfmt.Registry) error {
	if swag.IsZero(m.OutInterface) { // not required
		return nil
	}

	if m.OutInterface != nil {
		if err := m.OutInterface.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outInterface")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("outInterface")
			}
			return err
		}
	}

	return nil
}

func (m *Model57) validatePackets(formats strfmt.Registry) error {

	if err := validate.Required("packets", "body", m.Packets); err != nil {
		return err
	}

	if err := validate.MinimumInt("packets", "body", *m.Packets, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Model57) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

func (m *Model57) validateTarget(formats strfmt.Registry) error {

	if err := validate.Required("target", "body", m.Target); err != nil {
		return err
	}

	return nil
}

func (m *Model57) validateTranslation(formats strfmt.Registry) error {
	if swag.IsZero(m.Translation) { // not required
		return nil
	}

	if m.Translation != nil {
		if err := m.Translation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("translation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("translation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this model 57 based on the context it is used
func (m *Model57) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDestination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInInterface(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOutInterface(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTranslation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model57) contextValidateDestination(ctx context.Context, formats strfmt.Registry) error {

	if m.Destination != nil {
		if err := m.Destination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destination")
			}
			return err
		}
	}

	return nil
}

func (m *Model57) contextValidateInInterface(ctx context.Context, formats strfmt.Registry) error {

	if m.InInterface != nil {
		if err := m.InInterface.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inInterface")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inInterface")
			}
			return err
		}
	}

	return nil
}

func (m *Model57) contextValidateOutInterface(ctx context.Context, formats strfmt.Registry) error {

	if m.OutInterface != nil {
		if err := m.OutInterface.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outInterface")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("outInterface")
			}
			return err
		}
	}

	return nil
}

func (m *Model57) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {
		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

func (m *Model57) contextValidateTranslation(ctx context.Context, formats strfmt.Registry) error {

	if m.Translation != nil {
		if err := m.Translation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("translation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("translation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Model57) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model57) UnmarshalBinary(b []byte) error {
	var res Model57
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
