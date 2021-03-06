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

// SiteIdentificationDetail site identification detail
//
// swagger:model SiteIdentificationDetail
type SiteIdentificationDetail struct {

	// id
	// Example: f7ac9cad-ea28-4390-93c8-7add010e8ee3
	// Required: true
	ID *string `json:"id"`

	// Name of the site.
	// Example: Mount Everest
	// Required: true
	Name *string `json:"name"`

	// parent
	Parent Parent4 `json:"parent,omitempty"`

	// Status of the site.
	// Example: active
	// Required: true
	// Enum: [active disconnected inactive]
	Status *string `json:"status"`

	// Whether access to internet is disabled for this site or not.
	// Example: false
	Suspended bool `json:"suspended,omitempty"`

	// Type of the site.
	// Example: site
	// Required: true
	// Enum: [site endpoint]
	Type *string `json:"type"`

	// Time when the site was last updated.
	// Example: 2018-11-14T15:20:32.004Z
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// Validate validates this site identification detail
func (m *SiteIdentificationDetail) Validate(formats strfmt.Registry) error {
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

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SiteIdentificationDetail) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SiteIdentificationDetail) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var siteIdentificationDetailTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","disconnected","inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteIdentificationDetailTypeStatusPropEnum = append(siteIdentificationDetailTypeStatusPropEnum, v)
	}
}

const (

	// SiteIdentificationDetailStatusActive captures enum value "active"
	SiteIdentificationDetailStatusActive string = "active"

	// SiteIdentificationDetailStatusDisconnected captures enum value "disconnected"
	SiteIdentificationDetailStatusDisconnected string = "disconnected"

	// SiteIdentificationDetailStatusInactive captures enum value "inactive"
	SiteIdentificationDetailStatusInactive string = "inactive"
)

// prop value enum
func (m *SiteIdentificationDetail) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, siteIdentificationDetailTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SiteIdentificationDetail) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

var siteIdentificationDetailTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["site","endpoint"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteIdentificationDetailTypeTypePropEnum = append(siteIdentificationDetailTypeTypePropEnum, v)
	}
}

const (

	// SiteIdentificationDetailTypeSite captures enum value "site"
	SiteIdentificationDetailTypeSite string = "site"

	// SiteIdentificationDetailTypeEndpoint captures enum value "endpoint"
	SiteIdentificationDetailTypeEndpoint string = "endpoint"
)

// prop value enum
func (m *SiteIdentificationDetail) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, siteIdentificationDetailTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SiteIdentificationDetail) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *SiteIdentificationDetail) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this site identification detail based on context it is used
func (m *SiteIdentificationDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SiteIdentificationDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SiteIdentificationDetail) UnmarshalBinary(b []byte) error {
	var res SiteIdentificationDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
