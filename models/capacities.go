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

// Capacities capacities
//
// swagger:model capacities
type Capacities []string

var capacitiesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["off"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		capacitiesItemsEnum = append(capacitiesItemsEnum, v)
	}
}

func (m *Capacities) validateCapacitiesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, capacitiesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this capacities
func (m Capacities) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		// value enum
		if err := m.validateCapacitiesItemsEnum(strconv.Itoa(i), "body", m[i]); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this capacities based on context it is used
func (m Capacities) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
