// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Alerts alerts
//
// swagger:model alerts
type Alerts []string

var alertsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HighOutage","MediumOutage","OutdatedFirmwareNetwork","OutdatedFirmwareDevice","HighUtilization","CriticalUtilization","SignalLow","SignalMedium"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		alertsItemsEnum = append(alertsItemsEnum, v)
	}
}

func (m *Alerts) validateAlertsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, alertsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this alerts
func (m Alerts) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		// value enum
		if err := m.validateAlertsItemsEnum(strconv.Itoa(i), "body", m[i]); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this alerts based on context it is used
func (m Alerts) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
