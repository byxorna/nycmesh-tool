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

// SiteAccessGroupSingleSite site access group single site
//
// swagger:model SiteAccessGroupSingleSite
type SiteAccessGroupSingleSite struct {

	// id
	// Example: f7ac9cad-ea28-4390-93c8-7add010e8ee3
	// Required: true
	ID *string `json:"id"`

	// If true, access to site is read-only. If false, access to site is read-write.
	// Example: true
	// Required: true
	IsReadOnly *bool `json:"isReadOnly"`

	// Name of the site.
	// Example: Building B
	// Required: true
	Name *string `json:"name"`

	// Type of the site.
	// Example: endpoint
	// Required: true
	// Enum: [site endpoint client]
	Type *string `json:"type"`
}

// Validate validates this site access group single site
func (m *SiteAccessGroupSingleSite) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsReadOnly(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *SiteAccessGroupSingleSite) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SiteAccessGroupSingleSite) validateIsReadOnly(formats strfmt.Registry) error {

	if err := validate.Required("isReadOnly", "body", m.IsReadOnly); err != nil {
		return err
	}

	return nil
}

func (m *SiteAccessGroupSingleSite) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var siteAccessGroupSingleSiteTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["site","endpoint","client"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteAccessGroupSingleSiteTypeTypePropEnum = append(siteAccessGroupSingleSiteTypeTypePropEnum, v)
	}
}

const (

	// SiteAccessGroupSingleSiteTypeSite captures enum value "site"
	SiteAccessGroupSingleSiteTypeSite string = "site"

	// SiteAccessGroupSingleSiteTypeEndpoint captures enum value "endpoint"
	SiteAccessGroupSingleSiteTypeEndpoint string = "endpoint"

	// SiteAccessGroupSingleSiteTypeClient captures enum value "client"
	SiteAccessGroupSingleSiteTypeClient string = "client"
)

// prop value enum
func (m *SiteAccessGroupSingleSite) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, siteAccessGroupSingleSiteTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SiteAccessGroupSingleSite) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this site access group single site based on context it is used
func (m *SiteAccessGroupSingleSite) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SiteAccessGroupSingleSite) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SiteAccessGroupSingleSite) UnmarshalBinary(b []byte) error {
	var res SiteAccessGroupSingleSite
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
