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

// NewsSchema news schema
//
// swagger:model NewsSchema
type NewsSchema struct {

	// Timestamp in milliseconds when the news was published.
	// Example: 1557222409000
	// Required: true
	Date *float64 `json:"date"`

	// ID of the news.
	// Example: 1377eac0-2e21-4a56-b1c1-682c16f39cb9
	// Required: true
	ID *string `json:"id"`

	// URL with with more details.
	// Example: https://community.ubnt.com/
	// Required: true
	Link *string `json:"link"`

	// News description.
	// Example: Check the community.
	// Required: true
	Text *string `json:"text"`

	// News severity.
	// Example: info
	// Required: true
	Type *string `json:"type"`

	// Who issued the news.
	// Example: UBNT
	// Required: true
	Who *string `json:"who"`
}

// Validate validates this news schema
func (m *NewsSchema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWho(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewsSchema) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	return nil
}

func (m *NewsSchema) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *NewsSchema) validateLink(formats strfmt.Registry) error {

	if err := validate.Required("link", "body", m.Link); err != nil {
		return err
	}

	return nil
}

func (m *NewsSchema) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	return nil
}

func (m *NewsSchema) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *NewsSchema) validateWho(formats strfmt.Registry) error {

	if err := validate.Required("who", "body", m.Who); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this news schema based on context it is used
func (m *NewsSchema) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewsSchema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewsSchema) UnmarshalBinary(b []byte) error {
	var res NewsSchema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
