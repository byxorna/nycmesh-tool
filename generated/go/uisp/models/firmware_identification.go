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

// FirmwareIdentification firmware identification
//
// swagger:model FirmwareIdentification
type FirmwareIdentification struct {

	// filename
	Filename string `json:"filename,omitempty"`

	// id
	// Example: f7ac9cad-ea28-4390-93c8-7add010e8ee3
	ID string `json:"id,omitempty"`

	// lite
	Lite bool `json:"lite,omitempty"`

	// models
	Models Models `json:"models,omitempty"`

	// origin
	Origin string `json:"origin,omitempty"`

	// platform Id
	// Enum: [UF_NANOG UF_LOCO UF_WIFI UF_INSTANT alpinev2 mt7621 UISPS UISPSPro UNMSS UISPLTE UISPPLite e50 e100 e200 e300 e600 e1000 u50 u100 u200 u300 u1000 eswh esgh ES esx EP EPX SW ACB WA 2WA XC 2XC XW XM TI GBE AirGW AirGWP AF AF02 AF06 AF07 AF08 AF09 af5xhd afltu aflturocket GP ltu60 SB UNKNOWN]
	PlatformID string `json:"platformId,omitempty"`

	// platform name
	// Enum: [UFiber NanoG UFiber Loco UFiber Wifi UFiber Instant UISPRPro UISPRLite UISPS UISPSPro UNMSS UISPLTE UISPPLite e50 e100 e200 e300 e600 e1000 u50 u100 u200 u300 u1000 ESWH ESGH ES ESX EP EPX SW ACB WA 2WA XC 2XC XW XM TI GBE AirGW AirGWP AF AF02 AF06 AF07 AF08 AF09 af5xhd afltu lturocket GP ltu60 SunMax UNKNOWN]
	PlatformName string `json:"platformName,omitempty"`

	// stable
	Stable bool `json:"stable,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this firmware identification
func (m *FirmwareIdentification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FirmwareIdentification) validateModels(formats strfmt.Registry) error {
	if swag.IsZero(m.Models) { // not required
		return nil
	}

	if err := m.Models.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("models")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("models")
		}
		return err
	}

	return nil
}

var firmwareIdentificationTypePlatformIDPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UF_NANOG","UF_LOCO","UF_WIFI","UF_INSTANT","alpinev2","mt7621","UISPS","UISPSPro","UNMSS","UISPLTE","UISPPLite","e50","e100","e200","e300","e600","e1000","u50","u100","u200","u300","u1000","eswh","esgh","ES","esx","EP","EPX","SW","ACB","WA","2WA","XC","2XC","XW","XM","TI","GBE","AirGW","AirGWP","AF","AF02","AF06","AF07","AF08","AF09","af5xhd","afltu","aflturocket","GP","ltu60","SB","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		firmwareIdentificationTypePlatformIDPropEnum = append(firmwareIdentificationTypePlatformIDPropEnum, v)
	}
}

const (

	// FirmwareIdentificationPlatformIDUFNANOG captures enum value "UF_NANOG"
	FirmwareIdentificationPlatformIDUFNANOG string = "UF_NANOG"

	// FirmwareIdentificationPlatformIDUFLOCO captures enum value "UF_LOCO"
	FirmwareIdentificationPlatformIDUFLOCO string = "UF_LOCO"

	// FirmwareIdentificationPlatformIDUFWIFI captures enum value "UF_WIFI"
	FirmwareIdentificationPlatformIDUFWIFI string = "UF_WIFI"

	// FirmwareIdentificationPlatformIDUFINSTANT captures enum value "UF_INSTANT"
	FirmwareIdentificationPlatformIDUFINSTANT string = "UF_INSTANT"

	// FirmwareIdentificationPlatformIDAlpinev2 captures enum value "alpinev2"
	FirmwareIdentificationPlatformIDAlpinev2 string = "alpinev2"

	// FirmwareIdentificationPlatformIDMt7621 captures enum value "mt7621"
	FirmwareIdentificationPlatformIDMt7621 string = "mt7621"

	// FirmwareIdentificationPlatformIDUISPS captures enum value "UISPS"
	FirmwareIdentificationPlatformIDUISPS string = "UISPS"

	// FirmwareIdentificationPlatformIDUISPSPro captures enum value "UISPSPro"
	FirmwareIdentificationPlatformIDUISPSPro string = "UISPSPro"

	// FirmwareIdentificationPlatformIDUNMSS captures enum value "UNMSS"
	FirmwareIdentificationPlatformIDUNMSS string = "UNMSS"

	// FirmwareIdentificationPlatformIDUISPLTE captures enum value "UISPLTE"
	FirmwareIdentificationPlatformIDUISPLTE string = "UISPLTE"

	// FirmwareIdentificationPlatformIDUISPPLite captures enum value "UISPPLite"
	FirmwareIdentificationPlatformIDUISPPLite string = "UISPPLite"

	// FirmwareIdentificationPlatformIDE50 captures enum value "e50"
	FirmwareIdentificationPlatformIDE50 string = "e50"

	// FirmwareIdentificationPlatformIDE100 captures enum value "e100"
	FirmwareIdentificationPlatformIDE100 string = "e100"

	// FirmwareIdentificationPlatformIDE200 captures enum value "e200"
	FirmwareIdentificationPlatformIDE200 string = "e200"

	// FirmwareIdentificationPlatformIDE300 captures enum value "e300"
	FirmwareIdentificationPlatformIDE300 string = "e300"

	// FirmwareIdentificationPlatformIDE600 captures enum value "e600"
	FirmwareIdentificationPlatformIDE600 string = "e600"

	// FirmwareIdentificationPlatformIDE1000 captures enum value "e1000"
	FirmwareIdentificationPlatformIDE1000 string = "e1000"

	// FirmwareIdentificationPlatformIDU50 captures enum value "u50"
	FirmwareIdentificationPlatformIDU50 string = "u50"

	// FirmwareIdentificationPlatformIDU100 captures enum value "u100"
	FirmwareIdentificationPlatformIDU100 string = "u100"

	// FirmwareIdentificationPlatformIDU200 captures enum value "u200"
	FirmwareIdentificationPlatformIDU200 string = "u200"

	// FirmwareIdentificationPlatformIDU300 captures enum value "u300"
	FirmwareIdentificationPlatformIDU300 string = "u300"

	// FirmwareIdentificationPlatformIDU1000 captures enum value "u1000"
	FirmwareIdentificationPlatformIDU1000 string = "u1000"

	// FirmwareIdentificationPlatformIDEswh captures enum value "eswh"
	FirmwareIdentificationPlatformIDEswh string = "eswh"

	// FirmwareIdentificationPlatformIDEsgh captures enum value "esgh"
	FirmwareIdentificationPlatformIDEsgh string = "esgh"

	// FirmwareIdentificationPlatformIDES captures enum value "ES"
	FirmwareIdentificationPlatformIDES string = "ES"

	// FirmwareIdentificationPlatformIDEsx captures enum value "esx"
	FirmwareIdentificationPlatformIDEsx string = "esx"

	// FirmwareIdentificationPlatformIDEP captures enum value "EP"
	FirmwareIdentificationPlatformIDEP string = "EP"

	// FirmwareIdentificationPlatformIDEPX captures enum value "EPX"
	FirmwareIdentificationPlatformIDEPX string = "EPX"

	// FirmwareIdentificationPlatformIDSW captures enum value "SW"
	FirmwareIdentificationPlatformIDSW string = "SW"

	// FirmwareIdentificationPlatformIDACB captures enum value "ACB"
	FirmwareIdentificationPlatformIDACB string = "ACB"

	// FirmwareIdentificationPlatformIDWA captures enum value "WA"
	FirmwareIdentificationPlatformIDWA string = "WA"

	// FirmwareIdentificationPlatformIDNr2WA captures enum value "2WA"
	FirmwareIdentificationPlatformIDNr2WA string = "2WA"

	// FirmwareIdentificationPlatformIDXC captures enum value "XC"
	FirmwareIdentificationPlatformIDXC string = "XC"

	// FirmwareIdentificationPlatformIDNr2XC captures enum value "2XC"
	FirmwareIdentificationPlatformIDNr2XC string = "2XC"

	// FirmwareIdentificationPlatformIDXW captures enum value "XW"
	FirmwareIdentificationPlatformIDXW string = "XW"

	// FirmwareIdentificationPlatformIDXM captures enum value "XM"
	FirmwareIdentificationPlatformIDXM string = "XM"

	// FirmwareIdentificationPlatformIDTI captures enum value "TI"
	FirmwareIdentificationPlatformIDTI string = "TI"

	// FirmwareIdentificationPlatformIDGBE captures enum value "GBE"
	FirmwareIdentificationPlatformIDGBE string = "GBE"

	// FirmwareIdentificationPlatformIDAirGW captures enum value "AirGW"
	FirmwareIdentificationPlatformIDAirGW string = "AirGW"

	// FirmwareIdentificationPlatformIDAirGWP captures enum value "AirGWP"
	FirmwareIdentificationPlatformIDAirGWP string = "AirGWP"

	// FirmwareIdentificationPlatformIDAF captures enum value "AF"
	FirmwareIdentificationPlatformIDAF string = "AF"

	// FirmwareIdentificationPlatformIDAF02 captures enum value "AF02"
	FirmwareIdentificationPlatformIDAF02 string = "AF02"

	// FirmwareIdentificationPlatformIDAF06 captures enum value "AF06"
	FirmwareIdentificationPlatformIDAF06 string = "AF06"

	// FirmwareIdentificationPlatformIDAF07 captures enum value "AF07"
	FirmwareIdentificationPlatformIDAF07 string = "AF07"

	// FirmwareIdentificationPlatformIDAF08 captures enum value "AF08"
	FirmwareIdentificationPlatformIDAF08 string = "AF08"

	// FirmwareIdentificationPlatformIDAF09 captures enum value "AF09"
	FirmwareIdentificationPlatformIDAF09 string = "AF09"

	// FirmwareIdentificationPlatformIDAf5xhd captures enum value "af5xhd"
	FirmwareIdentificationPlatformIDAf5xhd string = "af5xhd"

	// FirmwareIdentificationPlatformIDAfltu captures enum value "afltu"
	FirmwareIdentificationPlatformIDAfltu string = "afltu"

	// FirmwareIdentificationPlatformIDAflturocket captures enum value "aflturocket"
	FirmwareIdentificationPlatformIDAflturocket string = "aflturocket"

	// FirmwareIdentificationPlatformIDGP captures enum value "GP"
	FirmwareIdentificationPlatformIDGP string = "GP"

	// FirmwareIdentificationPlatformIDLtu60 captures enum value "ltu60"
	FirmwareIdentificationPlatformIDLtu60 string = "ltu60"

	// FirmwareIdentificationPlatformIDSB captures enum value "SB"
	FirmwareIdentificationPlatformIDSB string = "SB"

	// FirmwareIdentificationPlatformIDUNKNOWN captures enum value "UNKNOWN"
	FirmwareIdentificationPlatformIDUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (m *FirmwareIdentification) validatePlatformIDEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, firmwareIdentificationTypePlatformIDPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FirmwareIdentification) validatePlatformID(formats strfmt.Registry) error {
	if swag.IsZero(m.PlatformID) { // not required
		return nil
	}

	// value enum
	if err := m.validatePlatformIDEnum("platformId", "body", m.PlatformID); err != nil {
		return err
	}

	return nil
}

var firmwareIdentificationTypePlatformNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UFiber NanoG","UFiber Loco","UFiber Wifi","UFiber Instant","UISPRPro","UISPRLite","UISPS","UISPSPro","UNMSS","UISPLTE","UISPPLite","e50","e100","e200","e300","e600","e1000","u50","u100","u200","u300","u1000","ESWH","ESGH","ES","ESX","EP","EPX","SW","ACB","WA","2WA","XC","2XC","XW","XM","TI","GBE","AirGW","AirGWP","AF","AF02","AF06","AF07","AF08","AF09","af5xhd","afltu","lturocket","GP","ltu60","SunMax","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		firmwareIdentificationTypePlatformNamePropEnum = append(firmwareIdentificationTypePlatformNamePropEnum, v)
	}
}

const (

	// FirmwareIdentificationPlatformNameUFiberNanoG captures enum value "UFiber NanoG"
	FirmwareIdentificationPlatformNameUFiberNanoG string = "UFiber NanoG"

	// FirmwareIdentificationPlatformNameUFiberLoco captures enum value "UFiber Loco"
	FirmwareIdentificationPlatformNameUFiberLoco string = "UFiber Loco"

	// FirmwareIdentificationPlatformNameUFiberWifi captures enum value "UFiber Wifi"
	FirmwareIdentificationPlatformNameUFiberWifi string = "UFiber Wifi"

	// FirmwareIdentificationPlatformNameUFiberInstant captures enum value "UFiber Instant"
	FirmwareIdentificationPlatformNameUFiberInstant string = "UFiber Instant"

	// FirmwareIdentificationPlatformNameUISPRPro captures enum value "UISPRPro"
	FirmwareIdentificationPlatformNameUISPRPro string = "UISPRPro"

	// FirmwareIdentificationPlatformNameUISPRLite captures enum value "UISPRLite"
	FirmwareIdentificationPlatformNameUISPRLite string = "UISPRLite"

	// FirmwareIdentificationPlatformNameUISPS captures enum value "UISPS"
	FirmwareIdentificationPlatformNameUISPS string = "UISPS"

	// FirmwareIdentificationPlatformNameUISPSPro captures enum value "UISPSPro"
	FirmwareIdentificationPlatformNameUISPSPro string = "UISPSPro"

	// FirmwareIdentificationPlatformNameUNMSS captures enum value "UNMSS"
	FirmwareIdentificationPlatformNameUNMSS string = "UNMSS"

	// FirmwareIdentificationPlatformNameUISPLTE captures enum value "UISPLTE"
	FirmwareIdentificationPlatformNameUISPLTE string = "UISPLTE"

	// FirmwareIdentificationPlatformNameUISPPLite captures enum value "UISPPLite"
	FirmwareIdentificationPlatformNameUISPPLite string = "UISPPLite"

	// FirmwareIdentificationPlatformNameE50 captures enum value "e50"
	FirmwareIdentificationPlatformNameE50 string = "e50"

	// FirmwareIdentificationPlatformNameE100 captures enum value "e100"
	FirmwareIdentificationPlatformNameE100 string = "e100"

	// FirmwareIdentificationPlatformNameE200 captures enum value "e200"
	FirmwareIdentificationPlatformNameE200 string = "e200"

	// FirmwareIdentificationPlatformNameE300 captures enum value "e300"
	FirmwareIdentificationPlatformNameE300 string = "e300"

	// FirmwareIdentificationPlatformNameE600 captures enum value "e600"
	FirmwareIdentificationPlatformNameE600 string = "e600"

	// FirmwareIdentificationPlatformNameE1000 captures enum value "e1000"
	FirmwareIdentificationPlatformNameE1000 string = "e1000"

	// FirmwareIdentificationPlatformNameU50 captures enum value "u50"
	FirmwareIdentificationPlatformNameU50 string = "u50"

	// FirmwareIdentificationPlatformNameU100 captures enum value "u100"
	FirmwareIdentificationPlatformNameU100 string = "u100"

	// FirmwareIdentificationPlatformNameU200 captures enum value "u200"
	FirmwareIdentificationPlatformNameU200 string = "u200"

	// FirmwareIdentificationPlatformNameU300 captures enum value "u300"
	FirmwareIdentificationPlatformNameU300 string = "u300"

	// FirmwareIdentificationPlatformNameU1000 captures enum value "u1000"
	FirmwareIdentificationPlatformNameU1000 string = "u1000"

	// FirmwareIdentificationPlatformNameESWH captures enum value "ESWH"
	FirmwareIdentificationPlatformNameESWH string = "ESWH"

	// FirmwareIdentificationPlatformNameESGH captures enum value "ESGH"
	FirmwareIdentificationPlatformNameESGH string = "ESGH"

	// FirmwareIdentificationPlatformNameES captures enum value "ES"
	FirmwareIdentificationPlatformNameES string = "ES"

	// FirmwareIdentificationPlatformNameESX captures enum value "ESX"
	FirmwareIdentificationPlatformNameESX string = "ESX"

	// FirmwareIdentificationPlatformNameEP captures enum value "EP"
	FirmwareIdentificationPlatformNameEP string = "EP"

	// FirmwareIdentificationPlatformNameEPX captures enum value "EPX"
	FirmwareIdentificationPlatformNameEPX string = "EPX"

	// FirmwareIdentificationPlatformNameSW captures enum value "SW"
	FirmwareIdentificationPlatformNameSW string = "SW"

	// FirmwareIdentificationPlatformNameACB captures enum value "ACB"
	FirmwareIdentificationPlatformNameACB string = "ACB"

	// FirmwareIdentificationPlatformNameWA captures enum value "WA"
	FirmwareIdentificationPlatformNameWA string = "WA"

	// FirmwareIdentificationPlatformNameNr2WA captures enum value "2WA"
	FirmwareIdentificationPlatformNameNr2WA string = "2WA"

	// FirmwareIdentificationPlatformNameXC captures enum value "XC"
	FirmwareIdentificationPlatformNameXC string = "XC"

	// FirmwareIdentificationPlatformNameNr2XC captures enum value "2XC"
	FirmwareIdentificationPlatformNameNr2XC string = "2XC"

	// FirmwareIdentificationPlatformNameXW captures enum value "XW"
	FirmwareIdentificationPlatformNameXW string = "XW"

	// FirmwareIdentificationPlatformNameXM captures enum value "XM"
	FirmwareIdentificationPlatformNameXM string = "XM"

	// FirmwareIdentificationPlatformNameTI captures enum value "TI"
	FirmwareIdentificationPlatformNameTI string = "TI"

	// FirmwareIdentificationPlatformNameGBE captures enum value "GBE"
	FirmwareIdentificationPlatformNameGBE string = "GBE"

	// FirmwareIdentificationPlatformNameAirGW captures enum value "AirGW"
	FirmwareIdentificationPlatformNameAirGW string = "AirGW"

	// FirmwareIdentificationPlatformNameAirGWP captures enum value "AirGWP"
	FirmwareIdentificationPlatformNameAirGWP string = "AirGWP"

	// FirmwareIdentificationPlatformNameAF captures enum value "AF"
	FirmwareIdentificationPlatformNameAF string = "AF"

	// FirmwareIdentificationPlatformNameAF02 captures enum value "AF02"
	FirmwareIdentificationPlatformNameAF02 string = "AF02"

	// FirmwareIdentificationPlatformNameAF06 captures enum value "AF06"
	FirmwareIdentificationPlatformNameAF06 string = "AF06"

	// FirmwareIdentificationPlatformNameAF07 captures enum value "AF07"
	FirmwareIdentificationPlatformNameAF07 string = "AF07"

	// FirmwareIdentificationPlatformNameAF08 captures enum value "AF08"
	FirmwareIdentificationPlatformNameAF08 string = "AF08"

	// FirmwareIdentificationPlatformNameAF09 captures enum value "AF09"
	FirmwareIdentificationPlatformNameAF09 string = "AF09"

	// FirmwareIdentificationPlatformNameAf5xhd captures enum value "af5xhd"
	FirmwareIdentificationPlatformNameAf5xhd string = "af5xhd"

	// FirmwareIdentificationPlatformNameAfltu captures enum value "afltu"
	FirmwareIdentificationPlatformNameAfltu string = "afltu"

	// FirmwareIdentificationPlatformNameLturocket captures enum value "lturocket"
	FirmwareIdentificationPlatformNameLturocket string = "lturocket"

	// FirmwareIdentificationPlatformNameGP captures enum value "GP"
	FirmwareIdentificationPlatformNameGP string = "GP"

	// FirmwareIdentificationPlatformNameLtu60 captures enum value "ltu60"
	FirmwareIdentificationPlatformNameLtu60 string = "ltu60"

	// FirmwareIdentificationPlatformNameSunMax captures enum value "SunMax"
	FirmwareIdentificationPlatformNameSunMax string = "SunMax"

	// FirmwareIdentificationPlatformNameUNKNOWN captures enum value "UNKNOWN"
	FirmwareIdentificationPlatformNameUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (m *FirmwareIdentification) validatePlatformNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, firmwareIdentificationTypePlatformNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FirmwareIdentification) validatePlatformName(formats strfmt.Registry) error {
	if swag.IsZero(m.PlatformName) { // not required
		return nil
	}

	// value enum
	if err := m.validatePlatformNameEnum("platformName", "body", m.PlatformName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this firmware identification based on the context it is used
func (m *FirmwareIdentification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateModels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FirmwareIdentification) contextValidateModels(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Models.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("models")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("models")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FirmwareIdentification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FirmwareIdentification) UnmarshalBinary(b []byte) error {
	var res FirmwareIdentification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
