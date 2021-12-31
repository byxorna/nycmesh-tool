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

// SitesStatistics sites statistics
//
// swagger:model SitesStatistics
type SitesStatistics struct {

	// download
	Download GraphDataSet `json:"download,omitempty"`

	// interval
	Interval *Interval1 `json:"interval,omitempty"`

	// Number of milliseconds between data points
	// Example: 15000
	Period float64 `json:"period,omitempty"`

	// upload
	Upload GraphDataSet `json:"upload,omitempty"`
}

// Validate validates this sites statistics
func (m *SitesStatistics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDownload(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpload(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SitesStatistics) validateDownload(formats strfmt.Registry) error {
	if swag.IsZero(m.Download) { // not required
		return nil
	}

	if err := m.Download.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("download")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("download")
		}
		return err
	}

	return nil
}

func (m *SitesStatistics) validateInterval(formats strfmt.Registry) error {
	if swag.IsZero(m.Interval) { // not required
		return nil
	}

	if m.Interval != nil {
		if err := m.Interval.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interval")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("interval")
			}
			return err
		}
	}

	return nil
}

func (m *SitesStatistics) validateUpload(formats strfmt.Registry) error {
	if swag.IsZero(m.Upload) { // not required
		return nil
	}

	if err := m.Upload.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("upload")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("upload")
		}
		return err
	}

	return nil
}

// ContextValidate validate this sites statistics based on the context it is used
func (m *SitesStatistics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDownload(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterval(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpload(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SitesStatistics) contextValidateDownload(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Download.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("download")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("download")
		}
		return err
	}

	return nil
}

func (m *SitesStatistics) contextValidateInterval(ctx context.Context, formats strfmt.Registry) error {

	if m.Interval != nil {
		if err := m.Interval.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interval")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("interval")
			}
			return err
		}
	}

	return nil
}

func (m *SitesStatistics) contextValidateUpload(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Upload.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("upload")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("upload")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SitesStatistics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SitesStatistics) UnmarshalBinary(b []byte) error {
	var res SitesStatistics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
