// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// ChannelWidthList channel width list
//
// swagger:model channelWidthList
type ChannelWidthList []float64

// Validate validates this channel width list
func (m ChannelWidthList) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this channel width list based on context it is used
func (m ChannelWidthList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
