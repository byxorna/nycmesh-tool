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

// Wireless wireless
//
// swagger:model wireless
type Wireless struct {

	// antenna
	Antenna *Antenna1 `json:"antenna,omitempty"`

	// Channel width in MHz radio is operating on.
	// Example: 40
	// Required: true
	// Minimum: 0
	ChannelWidth *int64 `json:"channelWidth"`

	// dfs lockouts
	// Required: true
	DfsLockouts DfsLockouts `json:"dfsLockouts"`

	// DFS wait time remaining.
	// Example: 0
	// Required: true
	DfsTimeRemaining *float64 `json:"dfsTimeRemaining"`

	// DFS expected total wait time.
	// Example: 0
	// Required: true
	DfsTimeTotal *float64 `json:"dfsTimeTotal"`

	// Downlink ratio in percentage. To calculate Uplink percentage substract DlRatio out of a hundered.
	// Example: 67
	// Required: true
	DlRatio *int64 `json:"dlRatio"`

	// frame length
	// Required: true
	// Minimum: 0
	FrameLength *float64 `json:"frameLength"`

	// Frequency in MHz, on which radio is listening.
	// Example: 5400
	// Required: true
	// Minimum: 0
	Frequency *float64 `json:"frequency"`

	// frequency band
	// Enum: [2.4GHz 3GHz 4GHz 5GHz 11GHz 24GHz 60GHz]
	FrequencyBand string `json:"frequencyBand,omitempty"`

	// Pre shared key
	// Required: true
	Key *string `json:"key"`

	// max tx power
	MaxTxPower float64 `json:"maxTxPower,omitempty"`

	// min tx power
	MinTxPower float64 `json:"minTxPower,omitempty"`

	// Wireless noise level in dBm
	// Example: -84
	// Required: true
	NoiseFloor *int64 `json:"noiseFloor"`

	// security
	// Required: true
	// Enum: [wep wpa wpa-psk wpa2 enabled none]
	Security *string `json:"security"`

	// SSID
	// Example: ubnt.com
	// Required: true
	Ssid *string `json:"ssid"`

	// transmit eirp
	TransmitEirp float64 `json:"transmitEirp,omitempty"`

	// transmit power
	TransmitPower float64 `json:"transmitPower,omitempty"`
}

// Validate validates this wireless
func (m *Wireless) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAntenna(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannelWidth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDfsLockouts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDfsTimeRemaining(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDfsTimeTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDlRatio(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrameLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequencyBand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNoiseFloor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsid(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Wireless) validateAntenna(formats strfmt.Registry) error {
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

func (m *Wireless) validateChannelWidth(formats strfmt.Registry) error {

	if err := validate.Required("channelWidth", "body", m.ChannelWidth); err != nil {
		return err
	}

	if err := validate.MinimumInt("channelWidth", "body", *m.ChannelWidth, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Wireless) validateDfsLockouts(formats strfmt.Registry) error {

	if err := validate.Required("dfsLockouts", "body", m.DfsLockouts); err != nil {
		return err
	}

	if err := m.DfsLockouts.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dfsLockouts")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("dfsLockouts")
		}
		return err
	}

	return nil
}

func (m *Wireless) validateDfsTimeRemaining(formats strfmt.Registry) error {

	if err := validate.Required("dfsTimeRemaining", "body", m.DfsTimeRemaining); err != nil {
		return err
	}

	return nil
}

func (m *Wireless) validateDfsTimeTotal(formats strfmt.Registry) error {

	if err := validate.Required("dfsTimeTotal", "body", m.DfsTimeTotal); err != nil {
		return err
	}

	return nil
}

func (m *Wireless) validateDlRatio(formats strfmt.Registry) error {

	if err := validate.Required("dlRatio", "body", m.DlRatio); err != nil {
		return err
	}

	return nil
}

func (m *Wireless) validateFrameLength(formats strfmt.Registry) error {

	if err := validate.Required("frameLength", "body", m.FrameLength); err != nil {
		return err
	}

	if err := validate.Minimum("frameLength", "body", *m.FrameLength, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Wireless) validateFrequency(formats strfmt.Registry) error {

	if err := validate.Required("frequency", "body", m.Frequency); err != nil {
		return err
	}

	if err := validate.Minimum("frequency", "body", *m.Frequency, 0, false); err != nil {
		return err
	}

	return nil
}

var wirelessTypeFrequencyBandPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["2.4GHz","3GHz","4GHz","5GHz","11GHz","24GHz","60GHz"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wirelessTypeFrequencyBandPropEnum = append(wirelessTypeFrequencyBandPropEnum, v)
	}
}

const (

	// WirelessFrequencyBandNr2Dot4GHz captures enum value "2.4GHz"
	WirelessFrequencyBandNr2Dot4GHz string = "2.4GHz"

	// WirelessFrequencyBandNr3GHz captures enum value "3GHz"
	WirelessFrequencyBandNr3GHz string = "3GHz"

	// WirelessFrequencyBandNr4GHz captures enum value "4GHz"
	WirelessFrequencyBandNr4GHz string = "4GHz"

	// WirelessFrequencyBandNr5GHz captures enum value "5GHz"
	WirelessFrequencyBandNr5GHz string = "5GHz"

	// WirelessFrequencyBandNr11GHz captures enum value "11GHz"
	WirelessFrequencyBandNr11GHz string = "11GHz"

	// WirelessFrequencyBandNr24GHz captures enum value "24GHz"
	WirelessFrequencyBandNr24GHz string = "24GHz"

	// WirelessFrequencyBandNr60GHz captures enum value "60GHz"
	WirelessFrequencyBandNr60GHz string = "60GHz"
)

// prop value enum
func (m *Wireless) validateFrequencyBandEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, wirelessTypeFrequencyBandPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Wireless) validateFrequencyBand(formats strfmt.Registry) error {
	if swag.IsZero(m.FrequencyBand) { // not required
		return nil
	}

	// value enum
	if err := m.validateFrequencyBandEnum("frequencyBand", "body", m.FrequencyBand); err != nil {
		return err
	}

	return nil
}

func (m *Wireless) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *Wireless) validateNoiseFloor(formats strfmt.Registry) error {

	if err := validate.Required("noiseFloor", "body", m.NoiseFloor); err != nil {
		return err
	}

	return nil
}

var wirelessTypeSecurityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["wep","wpa","wpa-psk","wpa2","enabled","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wirelessTypeSecurityPropEnum = append(wirelessTypeSecurityPropEnum, v)
	}
}

const (

	// WirelessSecurityWep captures enum value "wep"
	WirelessSecurityWep string = "wep"

	// WirelessSecurityWpa captures enum value "wpa"
	WirelessSecurityWpa string = "wpa"

	// WirelessSecurityWpaDashPsk captures enum value "wpa-psk"
	WirelessSecurityWpaDashPsk string = "wpa-psk"

	// WirelessSecurityWpa2 captures enum value "wpa2"
	WirelessSecurityWpa2 string = "wpa2"

	// WirelessSecurityEnabled captures enum value "enabled"
	WirelessSecurityEnabled string = "enabled"

	// WirelessSecurityNone captures enum value "none"
	WirelessSecurityNone string = "none"
)

// prop value enum
func (m *Wireless) validateSecurityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, wirelessTypeSecurityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Wireless) validateSecurity(formats strfmt.Registry) error {

	if err := validate.Required("security", "body", m.Security); err != nil {
		return err
	}

	// value enum
	if err := m.validateSecurityEnum("security", "body", *m.Security); err != nil {
		return err
	}

	return nil
}

func (m *Wireless) validateSsid(formats strfmt.Registry) error {

	if err := validate.Required("ssid", "body", m.Ssid); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this wireless based on the context it is used
func (m *Wireless) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAntenna(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDfsLockouts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Wireless) contextValidateAntenna(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Wireless) contextValidateDfsLockouts(ctx context.Context, formats strfmt.Registry) error {

	if err := m.DfsLockouts.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dfsLockouts")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("dfsLockouts")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Wireless) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Wireless) UnmarshalBinary(b []byte) error {
	var res Wireless
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
