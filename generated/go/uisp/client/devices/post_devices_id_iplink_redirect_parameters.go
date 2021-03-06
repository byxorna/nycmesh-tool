// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPostDevicesIDIplinkRedirectParams creates a new PostDevicesIDIplinkRedirectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDevicesIDIplinkRedirectParams() *PostDevicesIDIplinkRedirectParams {
	return &PostDevicesIDIplinkRedirectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesIDIplinkRedirectParamsWithTimeout creates a new PostDevicesIDIplinkRedirectParams object
// with the ability to set a timeout on a request.
func NewPostDevicesIDIplinkRedirectParamsWithTimeout(timeout time.Duration) *PostDevicesIDIplinkRedirectParams {
	return &PostDevicesIDIplinkRedirectParams{
		timeout: timeout,
	}
}

// NewPostDevicesIDIplinkRedirectParamsWithContext creates a new PostDevicesIDIplinkRedirectParams object
// with the ability to set a context for a request.
func NewPostDevicesIDIplinkRedirectParamsWithContext(ctx context.Context) *PostDevicesIDIplinkRedirectParams {
	return &PostDevicesIDIplinkRedirectParams{
		Context: ctx,
	}
}

// NewPostDevicesIDIplinkRedirectParamsWithHTTPClient creates a new PostDevicesIDIplinkRedirectParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDevicesIDIplinkRedirectParamsWithHTTPClient(client *http.Client) *PostDevicesIDIplinkRedirectParams {
	return &PostDevicesIDIplinkRedirectParams{
		HTTPClient: client,
	}
}

/* PostDevicesIDIplinkRedirectParams contains all the parameters to send to the API endpoint
   for the post devices Id iplink redirect operation.

   Typically these are written to a http.Request.
*/
type PostDevicesIDIplinkRedirectParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post devices Id iplink redirect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesIDIplinkRedirectParams) WithDefaults() *PostDevicesIDIplinkRedirectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post devices Id iplink redirect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesIDIplinkRedirectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post devices Id iplink redirect params
func (o *PostDevicesIDIplinkRedirectParams) WithTimeout(timeout time.Duration) *PostDevicesIDIplinkRedirectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices Id iplink redirect params
func (o *PostDevicesIDIplinkRedirectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices Id iplink redirect params
func (o *PostDevicesIDIplinkRedirectParams) WithContext(ctx context.Context) *PostDevicesIDIplinkRedirectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices Id iplink redirect params
func (o *PostDevicesIDIplinkRedirectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices Id iplink redirect params
func (o *PostDevicesIDIplinkRedirectParams) WithHTTPClient(client *http.Client) *PostDevicesIDIplinkRedirectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices Id iplink redirect params
func (o *PostDevicesIDIplinkRedirectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post devices Id iplink redirect params
func (o *PostDevicesIDIplinkRedirectParams) WithID(id string) *PostDevicesIDIplinkRedirectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post devices Id iplink redirect params
func (o *PostDevicesIDIplinkRedirectParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesIDIplinkRedirectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
