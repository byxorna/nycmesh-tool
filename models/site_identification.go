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

// SiteIdentification site identification
//
// swagger:model SiteIdentification
type SiteIdentification struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// parent
	Parent Parent2 `json:"parent,omitempty"`

	// status
	// Required: true
	// Enum: [active disconnected inactive]
	Status *string `json:"status"`

	// type
	// Required: true
	// Enum: [site endpoint client]
	Type *string `json:"type"`
}

// Validate validates this site identification
func (m *SiteIdentification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *SiteIdentification) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SiteIdentification) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var siteIdentificationTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","disconnected","inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteIdentificationTypeStatusPropEnum = append(siteIdentificationTypeStatusPropEnum, v)
	}
}

const (

	// SiteIdentificationStatusActive captures enum value "active"
	SiteIdentificationStatusActive string = "active"

	// SiteIdentificationStatusDisconnected captures enum value "disconnected"
	SiteIdentificationStatusDisconnected string = "disconnected"

	// SiteIdentificationStatusInactive captures enum value "inactive"
	SiteIdentificationStatusInactive string = "inactive"
)

// prop value enum
func (m *SiteIdentification) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, siteIdentificationTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SiteIdentification) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

var siteIdentificationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["site","endpoint","client"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteIdentificationTypeTypePropEnum = append(siteIdentificationTypeTypePropEnum, v)
	}
}

const (

	// SiteIdentificationTypeSite captures enum value "site"
	SiteIdentificationTypeSite string = "site"

	// SiteIdentificationTypeEndpoint captures enum value "endpoint"
	SiteIdentificationTypeEndpoint string = "endpoint"

	// SiteIdentificationTypeClient captures enum value "client"
	SiteIdentificationTypeClient string = "client"
)

// prop value enum
func (m *SiteIdentification) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, siteIdentificationTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SiteIdentification) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this site identification based on context it is used
func (m *SiteIdentification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SiteIdentification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SiteIdentification) UnmarshalBinary(b []byte) error {
	var res SiteIdentification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}