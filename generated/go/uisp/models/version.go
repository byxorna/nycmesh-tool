// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Version version
//
// swagger:model Version
type Version struct {

	// Build identification.
	// Example: release.860d2c9.20181218093614
	// Required: true
	Build *string `json:"build"`

	// How was the UISP deployed. Content of UISP_DEPLOYMENT env variable.
	// Example: jenkins-unms-test-deploy-1234
	// Required: true
	Deployment *string `json:"deployment"`

	// Time of the build.
	// Example: 2018-11-20T11:01:32+01:00
	// Required: true
	Time *string `json:"time"`

	// Version of UISP.
	// Example: 1.3.11
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this version
func (m *Version) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuild(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeployment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Version) validateBuild(formats strfmt.Registry) error {

	if err := validate.Required("build", "body", m.Build); err != nil {
		return err
	}

	return nil
}

func (m *Version) validateDeployment(formats strfmt.Registry) error {

	if err := validate.Required("deployment", "body", m.Deployment); err != nil {
		return err
	}

	return nil
}

func (m *Version) validateTime(formats strfmt.Registry) error {

	if err := validate.Required("time", "body", m.Time); err != nil {
		return err
	}

	return nil
}

func (m *Version) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this version based on context it is used
func (m *Version) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Version) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Version) UnmarshalBinary(b []byte) error {
	var res Version
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
