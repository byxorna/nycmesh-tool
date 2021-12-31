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

// OnuProfile1 onu profile 1
//
// swagger:model OnuProfile 1
type OnuProfile1 struct {

	// admin password
	// Required: true
	// Max Length: 12
	// Min Length: 1
	AdminPassword *string `json:"adminPassword"`

	// bandwidth limit down
	// Maximum: 1000
	// Minimum: 1
	BandwidthLimitDown int64 `json:"bandwidthLimitDown,omitempty"`

	// bandwidth limit enabled
	BandwidthLimitEnabled bool `json:"bandwidthLimitEnabled,omitempty"`

	// bandwidth limit up
	// Maximum: 1000
	// Minimum: 1
	BandwidthLimitUp int64 `json:"bandwidthLimitUp,omitempty"`

	// bridge
	Bridge *Bridge2 `json:"bridge,omitempty"`

	// http port
	// Required: true
	// Maximum: 65535
	// Minimum: 1
	HTTPPort *int64 `json:"httpPort"`

	// lan address
	// Required: true
	LanAddress *string `json:"lanAddress"`

	// lan provisioned
	// Required: true
	LanProvisioned *bool `json:"lanProvisioned"`

	// link speed
	LinkSpeed LinkSpeed `json:"linkSpeed,omitempty"`

	// mode
	// Required: true
	// Enum: [bridge router soho]
	Mode *string `json:"mode"`

	// name
	// Required: true
	// Max Length: 80
	// Min Length: 1
	Name *string `json:"name"`

	// router
	Router *Router1 `json:"router,omitempty"`

	// ssh enabled
	SSHEnabled bool `json:"sshEnabled,omitempty"`

	// ssh port
	// Maximum: 65535
	// Minimum: 1
	SSHPort int64 `json:"sshPort,omitempty"`

	// telnet enabled
	TelnetEnabled bool `json:"telnetEnabled,omitempty"`

	// telnet port
	// Maximum: 65535
	// Minimum: 1
	TelnetPort int64 `json:"telnetPort,omitempty"`

	// ubnt discovery enabled
	UbntDiscoveryEnabled bool `json:"ubntDiscoveryEnabled,omitempty"`
}

// Validate validates this onu profile 1
func (m *OnuProfile1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdminPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBandwidthLimitDown(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBandwidthLimitUp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBridge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanProvisioned(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSHPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTelnetPort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OnuProfile1) validateAdminPassword(formats strfmt.Registry) error {

	if err := validate.Required("adminPassword", "body", m.AdminPassword); err != nil {
		return err
	}

	if err := validate.MinLength("adminPassword", "body", *m.AdminPassword, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("adminPassword", "body", *m.AdminPassword, 12); err != nil {
		return err
	}

	return nil
}

func (m *OnuProfile1) validateBandwidthLimitDown(formats strfmt.Registry) error {
	if swag.IsZero(m.BandwidthLimitDown) { // not required
		return nil
	}

	if err := validate.MinimumInt("bandwidthLimitDown", "body", m.BandwidthLimitDown, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("bandwidthLimitDown", "body", m.BandwidthLimitDown, 1000, false); err != nil {
		return err
	}

	return nil
}

func (m *OnuProfile1) validateBandwidthLimitUp(formats strfmt.Registry) error {
	if swag.IsZero(m.BandwidthLimitUp) { // not required
		return nil
	}

	if err := validate.MinimumInt("bandwidthLimitUp", "body", m.BandwidthLimitUp, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("bandwidthLimitUp", "body", m.BandwidthLimitUp, 1000, false); err != nil {
		return err
	}

	return nil
}

func (m *OnuProfile1) validateBridge(formats strfmt.Registry) error {
	if swag.IsZero(m.Bridge) { // not required
		return nil
	}

	if m.Bridge != nil {
		if err := m.Bridge.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bridge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bridge")
			}
			return err
		}
	}

	return nil
}

func (m *OnuProfile1) validateHTTPPort(formats strfmt.Registry) error {

	if err := validate.Required("httpPort", "body", m.HTTPPort); err != nil {
		return err
	}

	if err := validate.MinimumInt("httpPort", "body", *m.HTTPPort, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("httpPort", "body", *m.HTTPPort, 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *OnuProfile1) validateLanAddress(formats strfmt.Registry) error {

	if err := validate.Required("lanAddress", "body", m.LanAddress); err != nil {
		return err
	}

	return nil
}

func (m *OnuProfile1) validateLanProvisioned(formats strfmt.Registry) error {

	if err := validate.Required("lanProvisioned", "body", m.LanProvisioned); err != nil {
		return err
	}

	return nil
}

var onuProfile1TypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bridge","router","soho"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		onuProfile1TypeModePropEnum = append(onuProfile1TypeModePropEnum, v)
	}
}

const (

	// OnuProfile1ModeBridge captures enum value "bridge"
	OnuProfile1ModeBridge string = "bridge"

	// OnuProfile1ModeRouter captures enum value "router"
	OnuProfile1ModeRouter string = "router"

	// OnuProfile1ModeSoho captures enum value "soho"
	OnuProfile1ModeSoho string = "soho"
)

// prop value enum
func (m *OnuProfile1) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, onuProfile1TypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OnuProfile1) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *OnuProfile1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 80); err != nil {
		return err
	}

	return nil
}

func (m *OnuProfile1) validateRouter(formats strfmt.Registry) error {
	if swag.IsZero(m.Router) { // not required
		return nil
	}

	if m.Router != nil {
		if err := m.Router.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("router")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("router")
			}
			return err
		}
	}

	return nil
}

func (m *OnuProfile1) validateSSHPort(formats strfmt.Registry) error {
	if swag.IsZero(m.SSHPort) { // not required
		return nil
	}

	if err := validate.MinimumInt("sshPort", "body", m.SSHPort, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("sshPort", "body", m.SSHPort, 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *OnuProfile1) validateTelnetPort(formats strfmt.Registry) error {
	if swag.IsZero(m.TelnetPort) { // not required
		return nil
	}

	if err := validate.MinimumInt("telnetPort", "body", m.TelnetPort, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("telnetPort", "body", m.TelnetPort, 65535, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this onu profile 1 based on the context it is used
func (m *OnuProfile1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBridge(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRouter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OnuProfile1) contextValidateBridge(ctx context.Context, formats strfmt.Registry) error {

	if m.Bridge != nil {
		if err := m.Bridge.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bridge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bridge")
			}
			return err
		}
	}

	return nil
}

func (m *OnuProfile1) contextValidateRouter(ctx context.Context, formats strfmt.Registry) error {

	if m.Router != nil {
		if err := m.Router.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("router")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("router")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OnuProfile1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OnuProfile1) UnmarshalBinary(b []byte) error {
	var res OnuProfile1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}