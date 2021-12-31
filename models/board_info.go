// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BoardInfo board info
//
// swagger:model boardInfo
type BoardInfo struct {

	// auto channel width list
	AutoChannelWidthList AutoChannelWidthList `json:"autoChannelWidthList,omitempty"`

	// is external reset supported
	// Required: true
	IsExternalResetSupported *bool `json:"isExternalResetSupported"`

	// led count
	// Required: true
	LedCount *float64 `json:"ledCount"`

	// model
	// Required: true
	Model *string `json:"model"`

	// physical interface count
	// Required: true
	PhysicalInterfaceCount *float64 `json:"physicalInterfaceCount"`

	// physical interface m a c addresses
	PhysicalInterfaceMACAddresses PhysicalInterfaceMACAddresses `json:"physicalInterfaceMACAddresses,omitempty"`

	// product
	// Required: true
	Product *string `json:"product"`

	// radio1
	Radio1 *Radio1 `json:"radio1,omitempty"`

	// reboot timeout
	// Required: true
	RebootTimeout *float64 `json:"rebootTimeout"`

	// upgrade timeout
	// Required: true
	UpgradeTimeout *float64 `json:"upgradeTimeout"`
}

// Validate validates this board info
func (m *BoardInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoChannelWidthList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsExternalResetSupported(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLedCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhysicalInterfaceCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhysicalInterfaceMACAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProduct(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRadio1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRebootTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpgradeTimeout(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BoardInfo) validateAutoChannelWidthList(formats strfmt.Registry) error {
	if swag.IsZero(m.AutoChannelWidthList) { // not required
		return nil
	}

	if err := m.AutoChannelWidthList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("autoChannelWidthList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("autoChannelWidthList")
		}
		return err
	}

	return nil
}

func (m *BoardInfo) validateIsExternalResetSupported(formats strfmt.Registry) error {

	if err := validate.Required("isExternalResetSupported", "body", m.IsExternalResetSupported); err != nil {
		return err
	}

	return nil
}

func (m *BoardInfo) validateLedCount(formats strfmt.Registry) error {

	if err := validate.Required("ledCount", "body", m.LedCount); err != nil {
		return err
	}

	return nil
}

func (m *BoardInfo) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	return nil
}

func (m *BoardInfo) validatePhysicalInterfaceCount(formats strfmt.Registry) error {

	if err := validate.Required("physicalInterfaceCount", "body", m.PhysicalInterfaceCount); err != nil {
		return err
	}

	return nil
}

func (m *BoardInfo) validatePhysicalInterfaceMACAddresses(formats strfmt.Registry) error {
	if swag.IsZero(m.PhysicalInterfaceMACAddresses) { // not required
		return nil
	}

	if err := m.PhysicalInterfaceMACAddresses.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("physicalInterfaceMACAddresses")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("physicalInterfaceMACAddresses")
		}
		return err
	}

	return nil
}

func (m *BoardInfo) validateProduct(formats strfmt.Registry) error {

	if err := validate.Required("product", "body", m.Product); err != nil {
		return err
	}

	return nil
}

func (m *BoardInfo) validateRadio1(formats strfmt.Registry) error {
	if swag.IsZero(m.Radio1) { // not required
		return nil
	}

	if m.Radio1 != nil {
		if err := m.Radio1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("radio1")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("radio1")
			}
			return err
		}
	}

	return nil
}

func (m *BoardInfo) validateRebootTimeout(formats strfmt.Registry) error {

	if err := validate.Required("rebootTimeout", "body", m.RebootTimeout); err != nil {
		return err
	}

	return nil
}

func (m *BoardInfo) validateUpgradeTimeout(formats strfmt.Registry) error {

	if err := validate.Required("upgradeTimeout", "body", m.UpgradeTimeout); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this board info based on the context it is used
func (m *BoardInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutoChannelWidthList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhysicalInterfaceMACAddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRadio1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BoardInfo) contextValidateAutoChannelWidthList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AutoChannelWidthList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("autoChannelWidthList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("autoChannelWidthList")
		}
		return err
	}

	return nil
}

func (m *BoardInfo) contextValidatePhysicalInterfaceMACAddresses(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PhysicalInterfaceMACAddresses.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("physicalInterfaceMACAddresses")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("physicalInterfaceMACAddresses")
		}
		return err
	}

	return nil
}

func (m *BoardInfo) contextValidateRadio1(ctx context.Context, formats strfmt.Registry) error {

	if m.Radio1 != nil {
		if err := m.Radio1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("radio1")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("radio1")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BoardInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BoardInfo) UnmarshalBinary(b []byte) error {
	var res BoardInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
