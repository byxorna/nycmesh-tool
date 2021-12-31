// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Semver semver
//
// swagger:model semver
type Semver struct {

	// current
	Current *Current `json:"current,omitempty"`

	// latest
	Latest *Latest `json:"latest,omitempty"`

	// latest on current major version
	LatestOnCurrentMajorVersion *LatestOnCurrentMajorVersion `json:"latestOnCurrentMajorVersion,omitempty"`

	// latest over
	LatestOver *LatestOver `json:"latestOver,omitempty"`
}

// Validate validates this semver
func (m *Semver) Validate(formats strfmt.Registry) error {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Semver) validateCurrent(formats strfmt.Registry) error {
	if swag.IsZero(m.Current) { // not required
		return nil
	}

	if m.Current != nil {
		if err := m.Current.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("current")
			}
			return err
		}
	}

	return nil
}

func (m *Semver) validateLatest(formats strfmt.Registry) error {
	if swag.IsZero(m.Latest) { // not required
		return nil
	}

	if m.Latest != nil {
		if err := m.Latest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latest")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latest")
			}
			return err
		}
	}

	return nil
}

func (m *Semver) validateLatestOnCurrentMajorVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.LatestOnCurrentMajorVersion) { // not required
		return nil
	}

	if m.LatestOnCurrentMajorVersion != nil {
		if err := m.LatestOnCurrentMajorVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latestOnCurrentMajorVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latestOnCurrentMajorVersion")
			}
			return err
		}
	}

	return nil
}

func (m *Semver) validateLatestOver(formats strfmt.Registry) error {
	if swag.IsZero(m.LatestOver) { // not required
		return nil
	}

	if m.LatestOver != nil {
		if err := m.LatestOver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latestOver")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latestOver")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this semver based on the context it is used
func (m *Semver) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCurrent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatestOnCurrentMajorVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatestOver(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Semver) contextValidateCurrent(ctx context.Context, formats strfmt.Registry) error {

	if m.Current != nil {
		if err := m.Current.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("current")
			}
			return err
		}
	}

	return nil
}

func (m *Semver) contextValidateLatest(ctx context.Context, formats strfmt.Registry) error {

	if m.Latest != nil {
		if err := m.Latest.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latest")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latest")
			}
			return err
		}
	}

	return nil
}

func (m *Semver) contextValidateLatestOnCurrentMajorVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.LatestOnCurrentMajorVersion != nil {
		if err := m.LatestOnCurrentMajorVersion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latestOnCurrentMajorVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latestOnCurrentMajorVersion")
			}
			return err
		}
	}

	return nil
}

func (m *Semver) contextValidateLatestOver(ctx context.Context, formats strfmt.Registry) error {

	if m.LatestOver != nil {
		if err := m.LatestOver.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latestOver")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latestOver")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Semver) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Semver) UnmarshalBinary(b []byte) error {
	var res Semver
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
