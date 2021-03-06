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

// SpeedTestObject speed test object
//
// swagger:model SpeedTestObject
type SpeedTestObject struct {

	// direction
	// Required: true
	// Enum: [uplink downlink bidirectional]
	Direction *string `json:"direction"`

	// result index
	// Required: true
	ResultIndex *float64 `json:"resultIndex"`

	// speed
	// Required: true
	Speed *float64 `json:"speed"`

	// speed test Id
	// Required: true
	SpeedTestID *string `json:"speedTestId"`

	// timestamp
	// Required: true
	Timestamp *float64 `json:"timestamp"`
}

// Validate validates this speed test object
func (m *SpeedTestObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResultIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpeed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpeedTestID(formats); err != nil {
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

var speedTestObjectTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["uplink","downlink","bidirectional"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		speedTestObjectTypeDirectionPropEnum = append(speedTestObjectTypeDirectionPropEnum, v)
	}
}

const (

	// SpeedTestObjectDirectionUplink captures enum value "uplink"
	SpeedTestObjectDirectionUplink string = "uplink"

	// SpeedTestObjectDirectionDownlink captures enum value "downlink"
	SpeedTestObjectDirectionDownlink string = "downlink"

	// SpeedTestObjectDirectionBidirectional captures enum value "bidirectional"
	SpeedTestObjectDirectionBidirectional string = "bidirectional"
)

// prop value enum
func (m *SpeedTestObject) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, speedTestObjectTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SpeedTestObject) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("direction", "body", m.Direction); err != nil {
		return err
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", *m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *SpeedTestObject) validateResultIndex(formats strfmt.Registry) error {

	if err := validate.Required("resultIndex", "body", m.ResultIndex); err != nil {
		return err
	}

	return nil
}

func (m *SpeedTestObject) validateSpeed(formats strfmt.Registry) error {

	if err := validate.Required("speed", "body", m.Speed); err != nil {
		return err
	}

	return nil
}

func (m *SpeedTestObject) validateSpeedTestID(formats strfmt.Registry) error {

	if err := validate.Required("speedTestId", "body", m.SpeedTestID); err != nil {
		return err
	}

	return nil
}

func (m *SpeedTestObject) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this speed test object based on context it is used
func (m *SpeedTestObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SpeedTestObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SpeedTestObject) UnmarshalBinary(b []byte) error {
	var res SpeedTestObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
