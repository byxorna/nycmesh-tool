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

// SimulationDevicePayload simulation device payload
//
// swagger:model SimulationDevicePayload
type SimulationDevicePayload struct {

	// altitude
	// Required: true
	// Maximum: 10000
	// Minimum: -10000
	Altitude *float64 `json:"altitude"`

	// antenna
	Antenna *Antenna2 `json:"antenna,omitempty"`

	// ap type
	// Required: true
	// Enum: [ptp ptmp]
	ApType *string `json:"apType"`

	// channel width
	// Example: 40
	// Required: true
	// Enum: [0]
	ChannelWidth *float64 `json:"channelWidth"`

	// color
	// Example: #ff00ff
	// Required: true
	// Max Length: 20
	Color *string `json:"color"`

	// coverage cpe height
	// Required: true
	// Enum: [0]
	CoverageCpeHeight *float64 `json:"coverageCpeHeight"`

	// device Id
	// Example: f7ac9cad-ea28-4390-93c8-7add010e8ee3
	// Required: true
	DeviceID *string `json:"deviceId"`

	// device name
	// Example: NanoStation 5AC Loco
	// Required: true
	// Max Length: 200
	DeviceName *string `json:"deviceName"`

	// eirp
	// Required: true
	// Enum: [0]
	Eirp *int64 `json:"eirp"`

	// frequency
	// Required: true
	// Enum: [0]
	Frequency *int64 `json:"frequency"`

	// heading
	// Required: true
	Heading *float64 `json:"heading"`

	// height
	// Required: true
	// Enum: [0]
	Height *float64 `json:"height"`

	// latitude
	// Example: 49.83455844211215
	// Required: true
	// Maximum: 90
	// Minimum: -90
	Latitude *float64 `json:"latitude"`

	// link
	Link *Link `json:"link,omitempty"`

	// longitude
	// Example: 13.463579999999956
	// Required: true
	// Maximum: 180
	// Minimum: -180
	Longitude *float64 `json:"longitude"`

	// name
	// Required: true
	// Max Length: 200
	Name *string `json:"name"`

	// radius
	// Required: true
	// Enum: [0]
	Radius *float64 `json:"radius"`

	// role
	// Required: true
	// Enum: [ap cpe]
	Role *string `json:"role"`
}

// Validate validates this simulation device payload
func (m *SimulationDevicePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAltitude(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAntenna(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannelWidth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCoverageCpeHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEirp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeading(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatitude(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLongitude(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRadius(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SimulationDevicePayload) validateAltitude(formats strfmt.Registry) error {

	if err := validate.Required("altitude", "body", m.Altitude); err != nil {
		return err
	}

	if err := validate.Minimum("altitude", "body", *m.Altitude, -10000, false); err != nil {
		return err
	}

	if err := validate.Maximum("altitude", "body", *m.Altitude, 10000, false); err != nil {
		return err
	}

	return nil
}

func (m *SimulationDevicePayload) validateAntenna(formats strfmt.Registry) error {
	if swag.IsZero(m.Antenna) { // not required
		return nil
	}

	if m.Antenna != nil {
		if err := m.Antenna.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("antenna")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("antenna")
			}
			return err
		}
	}

	return nil
}

var simulationDevicePayloadTypeApTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ptp","ptmp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simulationDevicePayloadTypeApTypePropEnum = append(simulationDevicePayloadTypeApTypePropEnum, v)
	}
}

const (

	// SimulationDevicePayloadApTypePtp captures enum value "ptp"
	SimulationDevicePayloadApTypePtp string = "ptp"

	// SimulationDevicePayloadApTypePtmp captures enum value "ptmp"
	SimulationDevicePayloadApTypePtmp string = "ptmp"
)

// prop value enum
func (m *SimulationDevicePayload) validateApTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, simulationDevicePayloadTypeApTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SimulationDevicePayload) validateApType(formats strfmt.Registry) error {

	if err := validate.Required("apType", "body", m.ApType); err != nil {
		return err
	}

	// value enum
	if err := m.validateApTypeEnum("apType", "body", *m.ApType); err != nil {
		return err
	}

	return nil
}

var simulationDevicePayloadTypeChannelWidthPropEnum []interface{}

func init() {
	var res []float64
	if err := json.Unmarshal([]byte(`[0]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simulationDevicePayloadTypeChannelWidthPropEnum = append(simulationDevicePayloadTypeChannelWidthPropEnum, v)
	}
}

// prop value enum
func (m *SimulationDevicePayload) validateChannelWidthEnum(path, location string, value float64) error {
	if err := validate.EnumCase(path, location, value, simulationDevicePayloadTypeChannelWidthPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SimulationDevicePayload) validateChannelWidth(formats strfmt.Registry) error {

	if err := validate.Required("channelWidth", "body", m.ChannelWidth); err != nil {
		return err
	}

	// value enum
	if err := m.validateChannelWidthEnum("channelWidth", "body", *m.ChannelWidth); err != nil {
		return err
	}

	return nil
}

func (m *SimulationDevicePayload) validateColor(formats strfmt.Registry) error {

	if err := validate.Required("color", "body", m.Color); err != nil {
		return err
	}

	if err := validate.MaxLength("color", "body", *m.Color, 20); err != nil {
		return err
	}

	return nil
}

var simulationDevicePayloadTypeCoverageCpeHeightPropEnum []interface{}

func init() {
	var res []float64
	if err := json.Unmarshal([]byte(`[0]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simulationDevicePayloadTypeCoverageCpeHeightPropEnum = append(simulationDevicePayloadTypeCoverageCpeHeightPropEnum, v)
	}
}

// prop value enum
func (m *SimulationDevicePayload) validateCoverageCpeHeightEnum(path, location string, value float64) error {
	if err := validate.EnumCase(path, location, value, simulationDevicePayloadTypeCoverageCpeHeightPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SimulationDevicePayload) validateCoverageCpeHeight(formats strfmt.Registry) error {

	if err := validate.Required("coverageCpeHeight", "body", m.CoverageCpeHeight); err != nil {
		return err
	}

	// value enum
	if err := m.validateCoverageCpeHeightEnum("coverageCpeHeight", "body", *m.CoverageCpeHeight); err != nil {
		return err
	}

	return nil
}

func (m *SimulationDevicePayload) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("deviceId", "body", m.DeviceID); err != nil {
		return err
	}

	return nil
}

func (m *SimulationDevicePayload) validateDeviceName(formats strfmt.Registry) error {

	if err := validate.Required("deviceName", "body", m.DeviceName); err != nil {
		return err
	}

	if err := validate.MaxLength("deviceName", "body", *m.DeviceName, 200); err != nil {
		return err
	}

	return nil
}

var simulationDevicePayloadTypeEirpPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simulationDevicePayloadTypeEirpPropEnum = append(simulationDevicePayloadTypeEirpPropEnum, v)
	}
}

// prop value enum
func (m *SimulationDevicePayload) validateEirpEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, simulationDevicePayloadTypeEirpPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SimulationDevicePayload) validateEirp(formats strfmt.Registry) error {

	if err := validate.Required("eirp", "body", m.Eirp); err != nil {
		return err
	}

	// value enum
	if err := m.validateEirpEnum("eirp", "body", *m.Eirp); err != nil {
		return err
	}

	return nil
}

var simulationDevicePayloadTypeFrequencyPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simulationDevicePayloadTypeFrequencyPropEnum = append(simulationDevicePayloadTypeFrequencyPropEnum, v)
	}
}

// prop value enum
func (m *SimulationDevicePayload) validateFrequencyEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, simulationDevicePayloadTypeFrequencyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SimulationDevicePayload) validateFrequency(formats strfmt.Registry) error {

	if err := validate.Required("frequency", "body", m.Frequency); err != nil {
		return err
	}

	// value enum
	if err := m.validateFrequencyEnum("frequency", "body", *m.Frequency); err != nil {
		return err
	}

	return nil
}

func (m *SimulationDevicePayload) validateHeading(formats strfmt.Registry) error {

	if err := validate.Required("heading", "body", m.Heading); err != nil {
		return err
	}

	return nil
}

var simulationDevicePayloadTypeHeightPropEnum []interface{}

func init() {
	var res []float64
	if err := json.Unmarshal([]byte(`[0]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simulationDevicePayloadTypeHeightPropEnum = append(simulationDevicePayloadTypeHeightPropEnum, v)
	}
}

// prop value enum
func (m *SimulationDevicePayload) validateHeightEnum(path, location string, value float64) error {
	if err := validate.EnumCase(path, location, value, simulationDevicePayloadTypeHeightPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SimulationDevicePayload) validateHeight(formats strfmt.Registry) error {

	if err := validate.Required("height", "body", m.Height); err != nil {
		return err
	}

	// value enum
	if err := m.validateHeightEnum("height", "body", *m.Height); err != nil {
		return err
	}

	return nil
}

func (m *SimulationDevicePayload) validateLatitude(formats strfmt.Registry) error {

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

func (m *SimulationDevicePayload) validateLink(formats strfmt.Registry) error {
	if swag.IsZero(m.Link) { // not required
		return nil
	}

	if m.Link != nil {
		if err := m.Link.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("link")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("link")
			}
			return err
		}
	}

	return nil
}

func (m *SimulationDevicePayload) validateLongitude(formats strfmt.Registry) error {

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

func (m *SimulationDevicePayload) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 200); err != nil {
		return err
	}

	return nil
}

var simulationDevicePayloadTypeRadiusPropEnum []interface{}

func init() {
	var res []float64
	if err := json.Unmarshal([]byte(`[0]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simulationDevicePayloadTypeRadiusPropEnum = append(simulationDevicePayloadTypeRadiusPropEnum, v)
	}
}

// prop value enum
func (m *SimulationDevicePayload) validateRadiusEnum(path, location string, value float64) error {
	if err := validate.EnumCase(path, location, value, simulationDevicePayloadTypeRadiusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SimulationDevicePayload) validateRadius(formats strfmt.Registry) error {

	if err := validate.Required("radius", "body", m.Radius); err != nil {
		return err
	}

	// value enum
	if err := m.validateRadiusEnum("radius", "body", *m.Radius); err != nil {
		return err
	}

	return nil
}

var simulationDevicePayloadTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ap","cpe"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simulationDevicePayloadTypeRolePropEnum = append(simulationDevicePayloadTypeRolePropEnum, v)
	}
}

const (

	// SimulationDevicePayloadRoleAp captures enum value "ap"
	SimulationDevicePayloadRoleAp string = "ap"

	// SimulationDevicePayloadRoleCpe captures enum value "cpe"
	SimulationDevicePayloadRoleCpe string = "cpe"
)

// prop value enum
func (m *SimulationDevicePayload) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, simulationDevicePayloadTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SimulationDevicePayload) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", *m.Role); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this simulation device payload based on the context it is used
func (m *SimulationDevicePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAntenna(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLink(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SimulationDevicePayload) contextValidateAntenna(ctx context.Context, formats strfmt.Registry) error {

	if m.Antenna != nil {
		if err := m.Antenna.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("antenna")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("antenna")
			}
			return err
		}
	}

	return nil
}

func (m *SimulationDevicePayload) contextValidateLink(ctx context.Context, formats strfmt.Registry) error {

	if m.Link != nil {
		if err := m.Link.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("link")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("link")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SimulationDevicePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SimulationDevicePayload) UnmarshalBinary(b []byte) error {
	var res SimulationDevicePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
