// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// DeletedIds deleted ids
//
// swagger:model deletedIds
type DeletedIds []string

// Validate validates this deleted ids
func (m DeletedIds) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this deleted ids based on context it is used
func (m DeletedIds) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}