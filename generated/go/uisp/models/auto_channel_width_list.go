// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// AutoChannelWidthList auto channel width list
//
// swagger:model autoChannelWidthList
type AutoChannelWidthList []int64

// Validate validates this auto channel width list
func (m AutoChannelWidthList) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this auto channel width list based on context it is used
func (m AutoChannelWidthList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}