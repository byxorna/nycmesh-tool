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

// NewGetDevicesAirmaxesIDFrequencybandsParams creates a new GetDevicesAirmaxesIDFrequencybandsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicesAirmaxesIDFrequencybandsParams() *GetDevicesAirmaxesIDFrequencybandsParams {
	return &GetDevicesAirmaxesIDFrequencybandsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesAirmaxesIDFrequencybandsParamsWithTimeout creates a new GetDevicesAirmaxesIDFrequencybandsParams object
// with the ability to set a timeout on a request.
func NewGetDevicesAirmaxesIDFrequencybandsParamsWithTimeout(timeout time.Duration) *GetDevicesAirmaxesIDFrequencybandsParams {
	return &GetDevicesAirmaxesIDFrequencybandsParams{
		timeout: timeout,
	}
}

// NewGetDevicesAirmaxesIDFrequencybandsParamsWithContext creates a new GetDevicesAirmaxesIDFrequencybandsParams object
// with the ability to set a context for a request.
func NewGetDevicesAirmaxesIDFrequencybandsParamsWithContext(ctx context.Context) *GetDevicesAirmaxesIDFrequencybandsParams {
	return &GetDevicesAirmaxesIDFrequencybandsParams{
		Context: ctx,
	}
}

// NewGetDevicesAirmaxesIDFrequencybandsParamsWithHTTPClient creates a new GetDevicesAirmaxesIDFrequencybandsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicesAirmaxesIDFrequencybandsParamsWithHTTPClient(client *http.Client) *GetDevicesAirmaxesIDFrequencybandsParams {
	return &GetDevicesAirmaxesIDFrequencybandsParams{
		HTTPClient: client,
	}
}

/* GetDevicesAirmaxesIDFrequencybandsParams contains all the parameters to send to the API endpoint
   for the get devices airmaxes Id frequencybands operation.

   Typically these are written to a http.Request.
*/
type GetDevicesAirmaxesIDFrequencybandsParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get devices airmaxes Id frequencybands params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesAirmaxesIDFrequencybandsParams) WithDefaults() *GetDevicesAirmaxesIDFrequencybandsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get devices airmaxes Id frequencybands params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesAirmaxesIDFrequencybandsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get devices airmaxes Id frequencybands params
func (o *GetDevicesAirmaxesIDFrequencybandsParams) WithTimeout(timeout time.Duration) *GetDevicesAirmaxesIDFrequencybandsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices airmaxes Id frequencybands params
func (o *GetDevicesAirmaxesIDFrequencybandsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices airmaxes Id frequencybands params
func (o *GetDevicesAirmaxesIDFrequencybandsParams) WithContext(ctx context.Context) *GetDevicesAirmaxesIDFrequencybandsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices airmaxes Id frequencybands params
func (o *GetDevicesAirmaxesIDFrequencybandsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices airmaxes Id frequencybands params
func (o *GetDevicesAirmaxesIDFrequencybandsParams) WithHTTPClient(client *http.Client) *GetDevicesAirmaxesIDFrequencybandsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices airmaxes Id frequencybands params
func (o *GetDevicesAirmaxesIDFrequencybandsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get devices airmaxes Id frequencybands params
func (o *GetDevicesAirmaxesIDFrequencybandsParams) WithID(id string) *GetDevicesAirmaxesIDFrequencybandsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get devices airmaxes Id frequencybands params
func (o *GetDevicesAirmaxesIDFrequencybandsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesAirmaxesIDFrequencybandsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
