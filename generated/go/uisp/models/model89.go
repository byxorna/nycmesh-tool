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

// Model89 model 89
//
// swagger:model Model 89
type Model89 struct {

	// guesses
	// Required: true
	Guesses *float64 `json:"guesses"`

	// guesses log10
	// Required: true
	GuessesLog10 *float64 `json:"guesses_log10"`

	// i
	// Required: true
	I *float64 `json:"i"`

	// j
	// Required: true
	J *float64 `json:"j"`

	// pattern
	// Required: true
	Pattern *string `json:"pattern"`

	// token
	// Required: true
	Token *string `json:"token"`
}

// Validate validates this model 89
func (m *Model89) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGuesses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGuessesLog10(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJ(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePattern(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Model89) validateGuesses(formats strfmt.Registry) error {

	if err := validate.Required("guesses", "body", m.Guesses); err != nil {
		return err
	}

	return nil
}

func (m *Model89) validateGuessesLog10(formats strfmt.Registry) error {

	if err := validate.Required("guesses_log10", "body", m.GuessesLog10); err != nil {
		return err
	}

	return nil
}

func (m *Model89) validateI(formats strfmt.Registry) error {

	if err := validate.Required("i", "body", m.I); err != nil {
		return err
	}

	return nil
}

func (m *Model89) validateJ(formats strfmt.Registry) error {

	if err := validate.Required("j", "body", m.J); err != nil {
		return err
	}

	return nil
}

func (m *Model89) validatePattern(formats strfmt.Registry) error {

	if err := validate.Required("pattern", "body", m.Pattern); err != nil {
		return err
	}

	return nil
}

func (m *Model89) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this model 89 based on context it is used
func (m *Model89) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Model89) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Model89) UnmarshalBinary(b []byte) error {
	var res Model89
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
