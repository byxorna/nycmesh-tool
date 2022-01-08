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

// DeviceStatistics device statistics
//
// swagger:model DeviceStatistics
type DeviceStatistics struct {

	// air time
	AirTime AirTime `json:"airTime,omitempty"`

	// ccq
	Ccq ListOfCoordinates `json:"ccq,omitempty"`

	// clients
	Clients string `json:"clients,omitempty"`

	// cpu
	CPU ListOfCoordinates `json:"cpu,omitempty"`

	// downlink capacity
	DownlinkCapacity string `json:"downlinkCapacity,omitempty"`

	// errors
	Errors ListOfCoordinates `json:"errors,omitempty"`

	// interfaces
	Interfaces Interfaces `json:"interfaces,omitempty"`

	// interval
	Interval *Interval `json:"interval,omitempty"`

	// local chain0
	LocalChain0 LocalChain0 `json:"localChain0,omitempty"`

	// local chain1
	LocalChain1 LocalChain1 `json:"localChain1,omitempty"`

	// output
	Output string `json:"output,omitempty"`

	// period
	Period float64 `json:"period,omitempty"`

	// ping
	Ping ListOfCoordinates `json:"ping,omitempty"`

	// psu
	Psu string `json:"psu,omitempty"`

	// pv
	Pv string `json:"pv,omitempty"`

	// ram
	RAM ListOfCoordinates `json:"ram,omitempty"`

	// remote chain0
	RemoteChain0 RemoteChain0 `json:"remoteChain0,omitempty"`

	// remote chain1
	RemoteChain1 RemoteChain1 `json:"remoteChain1,omitempty"`

	// remote signal
	RemoteSignal ListOfCoordinates `json:"remoteSignal,omitempty"`

	// remote signal60g
	RemoteSignal60g ListOfCoordinates `json:"remoteSignal60g,omitempty"`

	// signal
	Signal ListOfCoordinates `json:"signal,omitempty"`

	// signal60g
	Signal60g ListOfCoordinates `json:"signal60g,omitempty"`

	// stations
	Stations Stations1 `json:"stations,omitempty"`

	// temperature
	Temperature Temperature `json:"temperature,omitempty"`

	// uplink capacity
	UplinkCapacity string `json:"uplinkCapacity,omitempty"`
}

// Validate validates this device statistics
func (m *DeviceStatistics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAirTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCcq(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterfaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalChain0(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalChain1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePing(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRAM(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteChain0(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteChain1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteSignal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteSignal60g(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignal60g(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemperature(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceStatistics) validateAirTime(formats strfmt.Registry) error {
	if swag.IsZero(m.AirTime) { // not required
		return nil
	}

	if err := m.AirTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("airTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("airTime")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) validateCcq(formats strfmt.Registry) error {
	if swag.IsZero(m.Ccq) { // not required
		return nil
	}

	if err := m.Ccq.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ccq")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ccq")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) validateCPU(formats strfmt.Registry) error {
	if swag.IsZero(m.CPU) { // not required
		return nil
	}

	if err := m.CPU.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cpu")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("cpu")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) validateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	if err := m.Errors.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("errors")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("errors")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) validateInterfaces(formats strfmt.Registry) error {
	if swag.IsZero(m.Interfaces) { // not required
		return nil
	}

	if err := m.Interfaces.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("interfaces")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("interfaces")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) validateInterval(formats strfmt.Registry) error {
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

func (m *DeviceStatistics) validateLocalChain0(formats strfmt.Registry) error {
	if swag.IsZero(m.LocalChain0) { // not required
		return nil
	}

	if err := m.LocalChain0.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("localChain0")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("localChain0")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) validateLocalChain1(formats strfmt.Registry) error {
	if swag.IsZero(m.LocalChain1) { // not required
		return nil
	}

	if err := m.LocalChain1.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("localChain1")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("localChain1")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) validatePing(formats strfmt.Registry) error {
	if swag.IsZero(m.Ping) { // not required
		return nil
	}

	if err := m.Ping.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ping")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ping")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) validateRAM(formats strfmt.Registry) error {
	if swag.IsZero(m.RAM) { // not required
		return nil
	}

	if err := m.RAM.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ram")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ram")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) validateRemoteChain0(formats strfmt.Registry) error {
	if swag.IsZero(m.RemoteChain0) { // not required
		return nil
	}

	if err := m.RemoteChain0.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("remoteChain0")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("remoteChain0")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) validateRemoteChain1(formats strfmt.Registry) error {
	if swag.IsZero(m.RemoteChain1) { // not required
		return nil
	}

	if err := m.RemoteChain1.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("remoteChain1")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("remoteChain1")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) validateRemoteSignal(formats strfmt.Registry) error {
	if swag.IsZero(m.RemoteSignal) { // not required
		return nil
	}

	if err := m.RemoteSignal.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("remoteSignal")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("remoteSignal")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) validateRemoteSignal60g(formats strfmt.Registry) error {
	if swag.IsZero(m.RemoteSignal60g) { // not required
		return nil
	}

	if err := m.RemoteSignal60g.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("remoteSignal60g")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("remoteSignal60g")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) validateSignal(formats strfmt.Registry) error {
	if swag.IsZero(m.Signal) { // not required
		return nil
	}

	if err := m.Signal.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("signal")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("signal")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) validateSignal60g(formats strfmt.Registry) error {
	if swag.IsZero(m.Signal60g) { // not required
		return nil
	}

	if err := m.Signal60g.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("signal60g")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("signal60g")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) validateTemperature(formats strfmt.Registry) error {
	if swag.IsZero(m.Temperature) { // not required
		return nil
	}

	if err := m.Temperature.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("temperature")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("temperature")
		}
		return err
	}

	return nil
}

// ContextValidate validate this device statistics based on the context it is used
func (m *DeviceStatistics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAirTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCcq(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCPU(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterfaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterval(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocalChain0(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocalChain1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePing(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRAM(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemoteChain0(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemoteChain1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemoteSignal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemoteSignal60g(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSignal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSignal60g(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTemperature(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceStatistics) contextValidateAirTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AirTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("airTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("airTime")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) contextValidateCcq(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Ccq.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ccq")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ccq")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) contextValidateCPU(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CPU.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cpu")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("cpu")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Errors.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("errors")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("errors")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) contextValidateInterfaces(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Interfaces.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("interfaces")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("interfaces")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) contextValidateInterval(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DeviceStatistics) contextValidateLocalChain0(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LocalChain0.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("localChain0")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("localChain0")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) contextValidateLocalChain1(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LocalChain1.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("localChain1")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("localChain1")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) contextValidatePing(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Ping.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ping")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ping")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) contextValidateRAM(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RAM.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ram")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ram")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) contextValidateRemoteChain0(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RemoteChain0.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("remoteChain0")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("remoteChain0")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) contextValidateRemoteChain1(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RemoteChain1.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("remoteChain1")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("remoteChain1")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) contextValidateRemoteSignal(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RemoteSignal.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("remoteSignal")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("remoteSignal")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) contextValidateRemoteSignal60g(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RemoteSignal60g.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("remoteSignal60g")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("remoteSignal60g")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) contextValidateSignal(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Signal.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("signal")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("signal")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) contextValidateSignal60g(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Signal60g.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("signal60g")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("signal60g")
		}
		return err
	}

	return nil
}

func (m *DeviceStatistics) contextValidateTemperature(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Temperature.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("temperature")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("temperature")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceStatistics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceStatistics) UnmarshalBinary(b []byte) error {
	var res DeviceStatistics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}