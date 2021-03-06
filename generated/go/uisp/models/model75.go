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

// Model75 model 75
//
// swagger:model Model 75
type Model75 struct {

	// Channel frequency in MHz
	// Example: 66400
	// Minimum: 0
	Channel *int64 `json:"channel,omitempty"`

	// Channel width in MHz
	// Example: 100
	// Minimum: 0
	ChannelWidth *int64 `json:"channelWidth,omitempty"`

	// Antenna's elevation in meters, defaults to 10m
	// Example: 570
	// Minimum: -1000
	Elevation int64 `json:"elevation,omitempty"`

	// Antenna's azimuth rotation in rad, max 6.28318. North: 0, East: PI/2, South: PI, West: 3*PI/2.
	// Example: 1.2566370614359172
	// Maximum: 6.28319
	// Minimum: 0
	Heading *float64 `json:"heading,omitempty"`

	// latitude
	// Example: 49.680859
	// Required: true
	// Maximum: 90
	// Minimum: -90
	Latitude *float64 `json:"latitude"`

	// longitude
	// Example: 13.465609
	// Required: true
	// Maximum: 180
	// Minimum: -180
	Longitude *float64 `json:"longitude"`

	// Antenna's azimuth rotation in rad with (magnetic declination is included), max 6.28318. North: 0, East: PI/2, South: PI, West: 3*PI/2.
	// Example: 1.1814056081614541
	// Maximum: 6.28319
	// Minimum: 0
	MagneticHeading *float64 `json:"magneticHeading,omitempty"`

	// model
	// Example: WaveAP
	// Required: true
	// Enum: [WaveAP]
	Model *string `json:"model"`

	// Resolution of calculated coverage data
	// Example: maxSize257
	// Enum: [maxSize81 maxSize257 maxSize385 maxSize513 maxSize769 maxSize1025]
	Resolution string `json:"resolution,omitempty"`

	// Antenna's axial rotation in rad. Horizontal: 0, increased by right (clockwise) rotation.
	// Example: 6.265732014659643
	// Maximum: 6.28319
	// Minimum: 0
	Roll *float64 `json:"roll,omitempty"`

	// Antenna's vertical rotation in rad. Straight: 0, Down: -PI/2, Up: PI/2.
	// Example: -0.05235987755982988
	// Maximum: 1.5708
	// Minimum: -1.5708
	Tilt *float64 `json:"tilt,omitempty"`
}

// Validate validates this model 75
func (m *Model75) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannelWidth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElevation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeading(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatitude(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLongitude(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMagneticHeading(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResolution(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoll(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTilt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model75) validateChannel(formats strfmt.Registry) error {
	if swag.IsZero(m.Channel) { // not required
		return nil
	}

	if err := validate.MinimumInt("channel", "body", *m.Channel, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Model75) validateChannelWidth(formats strfmt.Registry) error {
	if swag.IsZero(m.ChannelWidth) { // not required
		return nil
	}

	if err := validate.MinimumInt("channelWidth", "body", *m.ChannelWidth, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Model75) validateElevation(formats strfmt.Registry) error {
	if swag.IsZero(m.Elevation) { // not required
		return nil
	}

	if err := validate.MinimumInt("elevation", "body", m.Elevation, -1000, false); err != nil {
		return err
	}

	return nil
}

func (m *Model75) validateHeading(formats strfmt.Registry) error {
	if swag.IsZero(m.Heading) { // not required
		return nil
	}

	if err := validate.Minimum("heading", "body", *m.Heading, 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("heading", "body", *m.Heading, 6.28319, false); err != nil {
		return err
	}

	return nil
}

func (m *Model75) validateLatitude(formats strfmt.Registry) error {

	if err := validate.Required("latitude", "body", m.Latitude); err != nil {
		return err
	}

	if err := validate.Minimum("latitude", "body", *m.Latitude, -90, false); err != nil {
		return err
	}

	if err := validate.Maximum("latitude", "body", *m.Latitude, 90, false); err != nil {
		return err
	}

	return nil
}

func (m *Model75) validateLongitude(formats strfmt.Registry) error {

	if err := validate.Required("longitude", "body", m.Longitude); err != nil {
		return err
	}

	if err := validate.Minimum("longitude", "body", *m.Longitude, -180, false); err != nil {
		return err
	}

	if err := validate.Maximum("longitude", "body", *m.Longitude, 180, false); err != nil {
		return err
	}

	return nil
}

func (m *Model75) validateMagneticHeading(formats strfmt.Registry) error {
	if swag.IsZero(m.MagneticHeading) { // not required
		return nil
	}

	if err := validate.Minimum("magneticHeading", "body", *m.MagneticHeading, 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("magneticHeading", "body", *m.MagneticHeading, 6.28319, false); err != nil {
		return err
	}

	return nil
}

var model75TypeModelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["WaveAP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		model75TypeModelPropEnum = append(model75TypeModelPropEnum, v)
	}
}

const (

	// Model75ModelWaveAP captures enum value "WaveAP"
	Model75ModelWaveAP string = "WaveAP"
)

// prop value enum
func (m *Model75) validateModelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, model75TypeModelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Model75) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	// value enum
	if err := m.validateModelEnum("model", "body", *m.Model); err != nil {
		return err
	}

	return nil
}

var model75TypeResolutionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["maxSize81","maxSize257","maxSize385","maxSize513","maxSize769","maxSize1025"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		model75TypeResolutionPropEnum = append(model75TypeResolutionPropEnum, v)
	}
}

const (

	// Model75ResolutionMaxSize81 captures enum value "maxSize81"
	Model75ResolutionMaxSize81 string = "maxSize81"

	// Model75ResolutionMaxSize257 captures enum value "maxSize257"
	Model75ResolutionMaxSize257 string = "maxSize257"

	// Model75ResolutionMaxSize385 captures enum value "maxSize385"
	Model75ResolutionMaxSize385 string = "maxSize385"

	// Model75ResolutionMaxSize513 captures enum value "maxSize513"
	Model75ResolutionMaxSize513 string = "maxSize513"

	// Model75ResolutionMaxSize769 captures enum value "maxSize769"
	Model75ResolutionMaxSize769 string = "maxSize769"

	// Model75ResolutionMaxSize1025 captures enum value "maxSize1025"
	Model75ResolutionMaxSize1025 string = "maxSize1025"
)

// prop value enum
func (m *Model75) validateResolutionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, model75TypeResolutionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Model75) validateResolution(formats strfmt.Registry) error {
	if swag.IsZero(m.Resolution) { // not required
		return nil
	}

	// value enum
	if err := m.validateResolutionEnum("resolution", "body", m.Resolution); err != nil {
		return err
	}

	return nil
}

func (m *Model75) validateRoll(formats strfmt.Registry) error {
	if swag.IsZero(m.Roll) { // not required
		return nil
	}

	if err := validate.Minimum("roll", "body", *m.Roll, 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("roll", "body", *m.Roll, 6.28319, false); err != nil {
		return err
	}

	return nil
}

func (m *Model75) validateTilt(formats strfmt.Registry) error {
	if swag.IsZero(m.Tilt) { // not required
		return nil
	}

	if err := validate.Minimum("tilt", "body", *m.Tilt, -1.5708, false); err != nil {
		return err
	}

	if err := validate.Maximum("tilt", "body", *m.Tilt, 1.5708, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this model 75 based on context it is used
func (m *Model75) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Model75) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model75) UnmarshalBinary(b []byte) error {
	var res Model75
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
