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

// AirCubeNetworkConfig air cube network config
//
// swagger:model AirCubeNetworkConfig
type AirCubeNetworkConfig struct {

	// block management access
	// Required: true
	BlockManagementAccess *bool `json:"blockManagementAccess"`

	// data vlan
	DataVlan *DataVlan `json:"dataVlan,omitempty"`

	// lan
	Lan *Lan `json:"lan,omitempty"`

	// mgt
	Mgt *Mgt `json:"mgt,omitempty"`

	// mode
	// Required: true
	// Enum: [bridge router soho]
	Mode *string `json:"mode"`

	// nat enabled
	NatEnabled bool `json:"natEnabled,omitempty"`

	// port forward rules
	// Required: true
	PortForwardRules PortForwardRules `json:"portForwardRules"`

	// static dhcp leases
	// Required: true
	StaticDhcpLeases StaticDhcpLeases `json:"staticDhcpLeases"`

	// static routes
	// Required: true
	StaticRoutes StaticRoutes `json:"staticRoutes"`

	// upnpd
	Upnpd *Upnpd `json:"upnpd,omitempty"`

	// vlp1
	Vlp1 *Vlp1 `json:"vlp1,omitempty"`

	// vlp2
	Vlp2 *Vlp2 `json:"vlp2,omitempty"`

	// vlp3
	Vlp3 *Vlp3 `json:"vlp3,omitempty"`

	// wan
	Wan *Wan `json:"wan,omitempty"`
}

// Validate validates this air cube network config
func (m *AirCubeNetworkConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlockManagementAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataVlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMgt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortForwardRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaticDhcpLeases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaticRoutes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpnpd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlp1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlp2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlp3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AirCubeNetworkConfig) validateBlockManagementAccess(formats strfmt.Registry) error {

	if err := validate.Required("blockManagementAccess", "body", m.BlockManagementAccess); err != nil {
		return err
	}

	return nil
}

func (m *AirCubeNetworkConfig) validateDataVlan(formats strfmt.Registry) error {
	if swag.IsZero(m.DataVlan) { // not required
		return nil
	}

	if m.DataVlan != nil {
		if err := m.DataVlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataVlan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataVlan")
			}
			return err
		}
	}

	return nil
}

func (m *AirCubeNetworkConfig) validateLan(formats strfmt.Registry) error {
	if swag.IsZero(m.Lan) { // not required
		return nil
	}

	if m.Lan != nil {
		if err := m.Lan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lan")
			}
			return err
		}
	}

	return nil
}

func (m *AirCubeNetworkConfig) validateMgt(formats strfmt.Registry) error {
	if swag.IsZero(m.Mgt) { // not required
		return nil
	}

	if m.Mgt != nil {
		if err := m.Mgt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mgt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mgt")
			}
			return err
		}
	}

	return nil
}

var airCubeNetworkConfigTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bridge","router","soho"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		airCubeNetworkConfigTypeModePropEnum = append(airCubeNetworkConfigTypeModePropEnum, v)
	}
}

const (

	// AirCubeNetworkConfigModeBridge captures enum value "bridge"
	AirCubeNetworkConfigModeBridge string = "bridge"

	// AirCubeNetworkConfigModeRouter captures enum value "router"
	AirCubeNetworkConfigModeRouter string = "router"

	// AirCubeNetworkConfigModeSoho captures enum value "soho"
	AirCubeNetworkConfigModeSoho string = "soho"
)

// prop value enum
func (m *AirCubeNetworkConfig) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, airCubeNetworkConfigTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AirCubeNetworkConfig) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *AirCubeNetworkConfig) validatePortForwardRules(formats strfmt.Registry) error {

	if err := validate.Required("portForwardRules", "body", m.PortForwardRules); err != nil {
		return err
	}

	if err := m.PortForwardRules.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("portForwardRules")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("portForwardRules")
		}
		return err
	}

	return nil
}

func (m *AirCubeNetworkConfig) validateStaticDhcpLeases(formats strfmt.Registry) error {

	if err := validate.Required("staticDhcpLeases", "body", m.StaticDhcpLeases); err != nil {
		return err
	}

	if err := m.StaticDhcpLeases.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("staticDhcpLeases")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("staticDhcpLeases")
		}
		return err
	}

	return nil
}

func (m *AirCubeNetworkConfig) validateStaticRoutes(formats strfmt.Registry) error {

	if err := validate.Required("staticRoutes", "body", m.StaticRoutes); err != nil {
		return err
	}

	if err := m.StaticRoutes.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("staticRoutes")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("staticRoutes")
		}
		return err
	}

	return nil
}

func (m *AirCubeNetworkConfig) validateUpnpd(formats strfmt.Registry) error {
	if swag.IsZero(m.Upnpd) { // not required
		return nil
	}

	if m.Upnpd != nil {
		if err := m.Upnpd.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upnpd")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("upnpd")
			}
			return err
		}
	}

	return nil
}

func (m *AirCubeNetworkConfig) validateVlp1(formats strfmt.Registry) error {
	if swag.IsZero(m.Vlp1) { // not required
		return nil
	}

	if m.Vlp1 != nil {
		if err := m.Vlp1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlp1")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlp1")
			}
			return err
		}
	}

	return nil
}

func (m *AirCubeNetworkConfig) validateVlp2(formats strfmt.Registry) error {
	if swag.IsZero(m.Vlp2) { // not required
		return nil
	}

	if m.Vlp2 != nil {
		if err := m.Vlp2.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlp2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlp2")
			}
			return err
		}
	}

	return nil
}

func (m *AirCubeNetworkConfig) validateVlp3(formats strfmt.Registry) error {
	if swag.IsZero(m.Vlp3) { // not required
		return nil
	}

	if m.Vlp3 != nil {
		if err := m.Vlp3.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlp3")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlp3")
			}
			return err
		}
	}

	return nil
}

func (m *AirCubeNetworkConfig) validateWan(formats strfmt.Registry) error {
	if swag.IsZero(m.Wan) { // not required
		return nil
	}

	if m.Wan != nil {
		if err := m.Wan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("wan")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this air cube network config based on the context it is used
func (m *AirCubeNetworkConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDataVlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMgt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePortForwardRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStaticDhcpLeases(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStaticRoutes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpnpd(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVlp1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVlp2(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVlp3(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AirCubeNetworkConfig) contextValidateDataVlan(ctx context.Context, formats strfmt.Registry) error {

	if m.DataVlan != nil {
		if err := m.DataVlan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataVlan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataVlan")
			}
			return err
		}
	}

	return nil
}

func (m *AirCubeNetworkConfig) contextValidateLan(ctx context.Context, formats strfmt.Registry) error {

	if m.Lan != nil {
		if err := m.Lan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lan")
			}
			return err
		}
	}

	return nil
}

func (m *AirCubeNetworkConfig) contextValidateMgt(ctx context.Context, formats strfmt.Registry) error {

	if m.Mgt != nil {
		if err := m.Mgt.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mgt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mgt")
			}
			return err
		}
	}

	return nil
}

func (m *AirCubeNetworkConfig) contextValidatePortForwardRules(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PortForwardRules.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("portForwardRules")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("portForwardRules")
		}
		return err
	}

	return nil
}

func (m *AirCubeNetworkConfig) contextValidateStaticDhcpLeases(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StaticDhcpLeases.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("staticDhcpLeases")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("staticDhcpLeases")
		}
		return err
	}

	return nil
}

func (m *AirCubeNetworkConfig) contextValidateStaticRoutes(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StaticRoutes.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("staticRoutes")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("staticRoutes")
		}
		return err
	}

	return nil
}

func (m *AirCubeNetworkConfig) contextValidateUpnpd(ctx context.Context, formats strfmt.Registry) error {

	if m.Upnpd != nil {
		if err := m.Upnpd.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upnpd")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("upnpd")
			}
			return err
		}
	}

	return nil
}

func (m *AirCubeNetworkConfig) contextValidateVlp1(ctx context.Context, formats strfmt.Registry) error {

	if m.Vlp1 != nil {
		if err := m.Vlp1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlp1")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlp1")
			}
			return err
		}
	}

	return nil
}

func (m *AirCubeNetworkConfig) contextValidateVlp2(ctx context.Context, formats strfmt.Registry) error {

	if m.Vlp2 != nil {
		if err := m.Vlp2.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlp2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlp2")
			}
			return err
		}
	}

	return nil
}

func (m *AirCubeNetworkConfig) contextValidateVlp3(ctx context.Context, formats strfmt.Registry) error {

	if m.Vlp3 != nil {
		if err := m.Vlp3.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlp3")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlp3")
			}
			return err
		}
	}

	return nil
}

func (m *AirCubeNetworkConfig) contextValidateWan(ctx context.Context, formats strfmt.Registry) error {

	if m.Wan != nil {
		if err := m.Wan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("wan")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AirCubeNetworkConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AirCubeNetworkConfig) UnmarshalBinary(b []byte) error {
	var res AirCubeNetworkConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
