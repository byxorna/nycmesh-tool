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

// Tags Log tags.
// Example: ["device","device-state"]
//
// swagger:model tags
type Tags []string

var tagsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["login","device","email-dispatch","nms-backup","nms-update","nms-error","device-state","device-backup","device-upgrade","device-interface","site"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tagsItemsEnum = append(tagsItemsEnum, v)
	}
}

func (m *Tags) validateTagsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tagsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this tags
func (m Tags) Validate(formats strfmt.Registry) error {
	var res []error

	iTagsSize := int64(len(m))

	if err := validate.MinItems("", "body", iTagsSize, 0); err != nil {
		return err
	}

	for i := 0; i < len(m); i++ {

		// value enum
		if err := m.validateTagsItemsEnum(strconv.Itoa(i), "body", m[i]); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this tags based on context it is used
func (m Tags) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}