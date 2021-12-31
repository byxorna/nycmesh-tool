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

// DeviceFirmware1 device firmware 1
//
// swagger:model DeviceFirmware 1
type DeviceFirmware1 struct {

	// If firmware is compatible with UISP
	Compatible bool `json:"compatible,omitempty"`

	// Current firmware version
	// Required: true
	Current *string `json:"current"`

	// Latest known firmware version
	// Required: true
	Latest *string `json:"latest"`

	// Latest known firmware version on current major version
	// Required: true
	LatestOnCurrentMajorVersion *string `json:"latestOnCurrentMajorVersion"`

	// Upgrade this firmware version as first if target
	//   firmware is higher than latestOver. (AirCube must be updated over 2.8.0 fw)
	// Required: true
	LatestOver *string `json:"latestOver"`

	// semver
	Semver *Semver `json:"semver,omitempty"`

	// Recommended firmware version to upgrade to.
	// Required: true
	UpgradeRecommendedToVersion *string `json:"upgradeRecommendedToVersion"`
}

// Validate validates this device firmware 1
func (m *DeviceFirmware1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestOnCurrentMajorVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestOver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSemver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpgradeRecommendedToVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceFirmware1) validateCurrent(formats strfmt.Registry) error {

	if err := validate.Required("current", "body", m.Current); err != nil {
		return err
	}

	return nil
}

func (m *DeviceFirmware1) validateLatest(formats strfmt.Registry) error {

	if err := validate.Required("latest", "body", m.Latest); err != nil {
		return err
	}

	return nil
}

func (m *DeviceFirmware1) validateLatestOnCurrentMajorVersion(formats strfmt.Registry) error {

	if err := validate.Required("latestOnCurrentMajorVersion", "body", m.LatestOnCurrentMajorVersion); err != nil {
		return err
	}

	return nil
}

func (m *DeviceFirmware1) validateLatestOver(formats strfmt.Registry) error {

	if err := validate.Required("latestOver", "body", m.LatestOver); err != nil {
		return err
	}

	return nil
}

func (m *DeviceFirmware1) validateSemver(formats strfmt.Registry) error {
	if swag.IsZero(m.Semver) { // not required
		return nil
	}

	if m.Semver != nil {
		if err := m.Semver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("semver")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("semver")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceFirmware1) validateUpgradeRecommendedToVersion(formats strfmt.Registry) error {

	if err := validate.Required("upgradeRecommendedToVersion", "body", m.UpgradeRecommendedToVersion); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this device firmware 1 based on the context it is used
func (m *DeviceFirmware1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSemver(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceFirmware1) contextValidateSemver(ctx context.Context, formats strfmt.Registry) error {

	if m.Semver != nil {
		if err := m.Semver.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("semver")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("semver")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceFirmware1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceFirmware1) UnmarshalBinary(b []byte) error {
	var res DeviceFirmware1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
