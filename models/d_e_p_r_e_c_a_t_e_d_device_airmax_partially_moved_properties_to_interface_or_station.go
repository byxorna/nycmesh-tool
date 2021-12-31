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

// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation d e p r e c a t e d device airmax partially moved properties to interface or station
//
// swagger:model [DEPRECATED]DeviceAirmax - partially moved properties to interface or station.
type DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation struct {

	// antenna
	Antenna string `json:"antenna,omitempty"`

	// ap device
	ApDevice *ApDevice `json:"apDevice,omitempty"`

	// ap mac
	ApMac string `json:"apMac,omitempty"`

	// authentication
	// Required: true
	// Enum: [psk psk2 ent none]
	Authentication *string `json:"authentication"`

	// available tx power range
	AvailableTxPowerRange *AvailableTxPowerRange `json:"availableTxPowerRange,omitempty"`

	// Cable Length from eth0 interface in meters. When below values is null cable is long less then 20m.
	CableLength float64 `json:"cableLength,omitempty"`

	// Average signal-to-noise ratio on eth0 interface in dB
	CableSnr float64 `json:"cableSnr,omitempty"`

	// ccq
	Ccq float64 `json:"ccq,omitempty"`

	// channel width
	// Required: true
	// Minimum: 0
	ChannelWidth *int64 `json:"channelWidth"`

	// country code
	// Required: true
	// Minimum: 0
	CountryCode *int64 `json:"countryCode"`

	// frequency
	// Example: 5400
	// Minimum: 0
	Frequency *float64 `json:"frequency,omitempty"`

	// frequency bands
	FrequencyBands FrequencyBands `json:"frequencyBands,omitempty"`

	// frequency center
	FrequencyCenter string `json:"frequencyCenter,omitempty"`

	// gps signal
	// Maximum: 1
	// Minimum: 0
	GpsSignal *float64 `json:"gpsSignal,omitempty"`

	// location
	Location *Location `json:"location,omitempty"`

	// noise floor
	NoiseFloor string `json:"noiseFloor,omitempty"`

	// polling
	Polling *Polling `json:"polling,omitempty"`

	// receive chains
	ReceiveChains float64 `json:"receiveChains,omitempty"`

	// remote signal
	RemoteSignal float64 `json:"remoteSignal,omitempty"`

	// remote signal60g
	RemoteSignal60g float64 `json:"remoteSignal60g,omitempty"`

	// Carrier to Interference-plus-Noise Ratio in dB on wireless
	RxCinr float64 `json:"rxCinr,omitempty"`

	// security
	// Required: true
	// Enum: [wep wpa wpa-psk wpa2 enabled none]
	Security *string `json:"security"`

	// series
	// Required: true
	// Enum: [AC M G60]
	Series *string `json:"series"`

	// signal chain
	SignalChain SignalChain `json:"signalChain,omitempty"`

	// signal remote chain
	SignalRemoteChain SignalRemoteChain `json:"signalRemoteChain,omitempty"`

	// SSID
	Ssid string `json:"ssid,omitempty"`

	// station name
	StationName string `json:"stationName,omitempty"`

	// stations count
	StationsCount float64 `json:"stationsCount,omitempty"`

	// tdd framing
	TddFraming string `json:"tddFraming,omitempty"`

	// transmit chains
	TransmitChains float64 `json:"transmitChains,omitempty"`

	// transmit power
	TransmitPower float64 `json:"transmitPower,omitempty"`

	// wds
	Wds bool `json:"wds,omitempty"`

	// wireless mode
	// Enum: [ap ap-ptp ap-ptmp ap-ptmp-airmax ap-ptmp-airmax-mixed ap-ptmp-airmax-ac sta sta-ptp sta-ptmp aprepeater repeater mesh]
	WirelessMode string `json:"wirelessMode,omitempty"`

	// wlan rx bytes
	WlanRxBytes float64 `json:"wlanRxBytes,omitempty"`

	// wlan tx bytes
	WlanTxBytes float64 `json:"wlanTxBytes,omitempty"`
}

// Validate validates this d e p r e c a t e d device airmax partially moved properties to interface or station
func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvailableTxPowerRange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannelWidth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountryCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequencyBands(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpsSignal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolling(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignalChain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignalRemoteChain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWirelessMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) validateApDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.ApDevice) { // not required
		return nil
	}

	if m.ApDevice != nil {
		if err := m.ApDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apDevice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apDevice")
			}
			return err
		}
	}

	return nil
}

var dEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationTypeAuthenticationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["psk","psk2","ent","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationTypeAuthenticationPropEnum = append(dEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationTypeAuthenticationPropEnum, v)
	}
}

const (

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationAuthenticationPsk captures enum value "psk"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationAuthenticationPsk string = "psk"

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationAuthenticationPsk2 captures enum value "psk2"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationAuthenticationPsk2 string = "psk2"

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationAuthenticationEnt captures enum value "ent"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationAuthenticationEnt string = "ent"

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationAuthenticationNone captures enum value "none"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationAuthenticationNone string = "none"
)

// prop value enum
func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) validateAuthenticationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationTypeAuthenticationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) validateAuthentication(formats strfmt.Registry) error {

	if err := validate.Required("authentication", "body", m.Authentication); err != nil {
		return err
	}

	// value enum
	if err := m.validateAuthenticationEnum("authentication", "body", *m.Authentication); err != nil {
		return err
	}

	return nil
}

func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) validateAvailableTxPowerRange(formats strfmt.Registry) error {
	if swag.IsZero(m.AvailableTxPowerRange) { // not required
		return nil
	}

	if m.AvailableTxPowerRange != nil {
		if err := m.AvailableTxPowerRange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("availableTxPowerRange")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("availableTxPowerRange")
			}
			return err
		}
	}

	return nil
}

func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) validateChannelWidth(formats strfmt.Registry) error {

	if err := validate.Required("channelWidth", "body", m.ChannelWidth); err != nil {
		return err
	}

	if err := validate.MinimumInt("channelWidth", "body", *m.ChannelWidth, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) validateCountryCode(formats strfmt.Registry) error {

	if err := validate.Required("countryCode", "body", m.CountryCode); err != nil {
		return err
	}

	if err := validate.MinimumInt("countryCode", "body", *m.CountryCode, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) validateFrequency(formats strfmt.Registry) error {
	if swag.IsZero(m.Frequency) { // not required
		return nil
	}

	if err := validate.Minimum("frequency", "body", *m.Frequency, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) validateFrequencyBands(formats strfmt.Registry) error {
	if swag.IsZero(m.FrequencyBands) { // not required
		return nil
	}

	if err := m.FrequencyBands.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("frequencyBands")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("frequencyBands")
		}
		return err
	}

	return nil
}

func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) validateGpsSignal(formats strfmt.Registry) error {
	if swag.IsZero(m.GpsSignal) { // not required
		return nil
	}

	if err := validate.Minimum("gpsSignal", "body", *m.GpsSignal, 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("gpsSignal", "body", *m.GpsSignal, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) validatePolling(formats strfmt.Registry) error {
	if swag.IsZero(m.Polling) { // not required
		return nil
	}

	if m.Polling != nil {
		if err := m.Polling.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("polling")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("polling")
			}
			return err
		}
	}

	return nil
}

var dEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationTypeSecurityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["wep","wpa","wpa-psk","wpa2","enabled","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationTypeSecurityPropEnum = append(dEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationTypeSecurityPropEnum, v)
	}
}

const (

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationSecurityWep captures enum value "wep"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationSecurityWep string = "wep"

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationSecurityWpa captures enum value "wpa"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationSecurityWpa string = "wpa"

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationSecurityWpaDashPsk captures enum value "wpa-psk"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationSecurityWpaDashPsk string = "wpa-psk"

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationSecurityWpa2 captures enum value "wpa2"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationSecurityWpa2 string = "wpa2"

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationSecurityEnabled captures enum value "enabled"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationSecurityEnabled string = "enabled"

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationSecurityNone captures enum value "none"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationSecurityNone string = "none"
)

// prop value enum
func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) validateSecurityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationTypeSecurityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) validateSecurity(formats strfmt.Registry) error {

	if err := validate.Required("security", "body", m.Security); err != nil {
		return err
	}

	// value enum
	if err := m.validateSecurityEnum("security", "body", *m.Security); err != nil {
		return err
	}

	return nil
}

var dEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationTypeSeriesPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AC","M","G60"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationTypeSeriesPropEnum = append(dEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationTypeSeriesPropEnum, v)
	}
}

const (

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationSeriesAC captures enum value "AC"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationSeriesAC string = "AC"

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationSeriesM captures enum value "M"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationSeriesM string = "M"

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationSeriesG60 captures enum value "G60"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationSeriesG60 string = "G60"
)

// prop value enum
func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) validateSeriesEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationTypeSeriesPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) validateSeries(formats strfmt.Registry) error {

	if err := validate.Required("series", "body", m.Series); err != nil {
		return err
	}

	// value enum
	if err := m.validateSeriesEnum("series", "body", *m.Series); err != nil {
		return err
	}

	return nil
}

func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) validateSignalChain(formats strfmt.Registry) error {
	if swag.IsZero(m.SignalChain) { // not required
		return nil
	}

	if err := m.SignalChain.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("signalChain")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("signalChain")
		}
		return err
	}

	return nil
}

func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) validateSignalRemoteChain(formats strfmt.Registry) error {
	if swag.IsZero(m.SignalRemoteChain) { // not required
		return nil
	}

	if err := m.SignalRemoteChain.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("signalRemoteChain")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("signalRemoteChain")
		}
		return err
	}

	return nil
}

var dEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationTypeWirelessModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ap","ap-ptp","ap-ptmp","ap-ptmp-airmax","ap-ptmp-airmax-mixed","ap-ptmp-airmax-ac","sta","sta-ptp","sta-ptmp","aprepeater","repeater","mesh"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationTypeWirelessModePropEnum = append(dEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationTypeWirelessModePropEnum, v)
	}
}

const (

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeAp captures enum value "ap"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeAp string = "ap"

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeApDashPtp captures enum value "ap-ptp"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeApDashPtp string = "ap-ptp"

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeApDashPtmp captures enum value "ap-ptmp"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeApDashPtmp string = "ap-ptmp"

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeApDashPtmpDashAirmax captures enum value "ap-ptmp-airmax"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeApDashPtmpDashAirmax string = "ap-ptmp-airmax"

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeApDashPtmpDashAirmaxDashMixed captures enum value "ap-ptmp-airmax-mixed"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeApDashPtmpDashAirmaxDashMixed string = "ap-ptmp-airmax-mixed"

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeApDashPtmpDashAirmaxDashAc captures enum value "ap-ptmp-airmax-ac"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeApDashPtmpDashAirmaxDashAc string = "ap-ptmp-airmax-ac"

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeSta captures enum value "sta"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeSta string = "sta"

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeStaDashPtp captures enum value "sta-ptp"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeStaDashPtp string = "sta-ptp"

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeStaDashPtmp captures enum value "sta-ptmp"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeStaDashPtmp string = "sta-ptmp"

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeAprepeater captures enum value "aprepeater"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeAprepeater string = "aprepeater"

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeRepeater captures enum value "repeater"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeRepeater string = "repeater"

	// DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeMesh captures enum value "mesh"
	DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationWirelessModeMesh string = "mesh"
)

// prop value enum
func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) validateWirelessModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStationTypeWirelessModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) validateWirelessMode(formats strfmt.Registry) error {
	if swag.IsZero(m.WirelessMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateWirelessModeEnum("wirelessMode", "body", m.WirelessMode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this d e p r e c a t e d device airmax partially moved properties to interface or station based on the context it is used
func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAvailableTxPowerRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFrequencyBands(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolling(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSignalChain(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSignalRemoteChain(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) contextValidateApDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.ApDevice != nil {
		if err := m.ApDevice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apDevice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apDevice")
			}
			return err
		}
	}

	return nil
}

func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) contextValidateAvailableTxPowerRange(ctx context.Context, formats strfmt.Registry) error {

	if m.AvailableTxPowerRange != nil {
		if err := m.AvailableTxPowerRange.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("availableTxPowerRange")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("availableTxPowerRange")
			}
			return err
		}
	}

	return nil
}

func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) contextValidateFrequencyBands(ctx context.Context, formats strfmt.Registry) error {

	if err := m.FrequencyBands.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("frequencyBands")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("frequencyBands")
		}
		return err
	}

	return nil
}

func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {
		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) contextValidatePolling(ctx context.Context, formats strfmt.Registry) error {

	if m.Polling != nil {
		if err := m.Polling.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("polling")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("polling")
			}
			return err
		}
	}

	return nil
}

func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) contextValidateSignalChain(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SignalChain.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("signalChain")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("signalChain")
		}
		return err
	}

	return nil
}

func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) contextValidateSignalRemoteChain(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SignalRemoteChain.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("signalRemoteChain")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("signalRemoteChain")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation) UnmarshalBinary(b []byte) error {
	var res DEPRECATEDDeviceAirmaxPartiallyMovedPropertiesToInterfaceOrStation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
