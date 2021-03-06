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

// ConnectionState connection state
//
// swagger:model connectionState
type ConnectionState []string

var connectionStateItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["established"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		connectionStateItemsEnum = append(connectionStateItemsEnum, v)
	}
}

func (m *ConnectionState) validateConnectionStateItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, connectionStateItemsEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this connection state
func (m ConnectionState) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		// value enum
		if err := m.validateConnectionStateItemsEnum(strconv.Itoa(i), "body", m[i]); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this connection state based on context it is used
func (m ConnectionState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
