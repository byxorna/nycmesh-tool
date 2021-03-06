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

// NewGetDevicesAirmaxesIDSystemParams creates a new GetDevicesAirmaxesIDSystemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicesAirmaxesIDSystemParams() *GetDevicesAirmaxesIDSystemParams {
	return &GetDevicesAirmaxesIDSystemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesAirmaxesIDSystemParamsWithTimeout creates a new GetDevicesAirmaxesIDSystemParams object
// with the ability to set a timeout on a request.
func NewGetDevicesAirmaxesIDSystemParamsWithTimeout(timeout time.Duration) *GetDevicesAirmaxesIDSystemParams {
	return &GetDevicesAirmaxesIDSystemParams{
		timeout: timeout,
	}
}

// NewGetDevicesAirmaxesIDSystemParamsWithContext creates a new GetDevicesAirmaxesIDSystemParams object
// with the ability to set a context for a request.
func NewGetDevicesAirmaxesIDSystemParamsWithContext(ctx context.Context) *GetDevicesAirmaxesIDSystemParams {
	return &GetDevicesAirmaxesIDSystemParams{
		Context: ctx,
	}
}

// NewGetDevicesAirmaxesIDSystemParamsWithHTTPClient creates a new GetDevicesAirmaxesIDSystemParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicesAirmaxesIDSystemParamsWithHTTPClient(client *http.Client) *GetDevicesAirmaxesIDSystemParams {
	return &GetDevicesAirmaxesIDSystemParams{
		HTTPClient: client,
	}
}

/* GetDevicesAirmaxesIDSystemParams contains all the parameters to send to the API endpoint
   for the get devices airmaxes Id system operation.

   Typically these are written to a http.Request.
*/
type GetDevicesAirmaxesIDSystemParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get devices airmaxes Id system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesAirmaxesIDSystemParams) WithDefaults() *GetDevicesAirmaxesIDSystemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get devices airmaxes Id system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesAirmaxesIDSystemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get devices airmaxes Id system params
func (o *GetDevicesAirmaxesIDSystemParams) WithTimeout(timeout time.Duration) *GetDevicesAirmaxesIDSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices airmaxes Id system params
func (o *GetDevicesAirmaxesIDSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices airmaxes Id system params
func (o *GetDevicesAirmaxesIDSystemParams) WithContext(ctx context.Context) *GetDevicesAirmaxesIDSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices airmaxes Id system params
func (o *GetDevicesAirmaxesIDSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices airmaxes Id system params
func (o *GetDevicesAirmaxesIDSystemParams) WithHTTPClient(client *http.Client) *GetDevicesAirmaxesIDSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices airmaxes Id system params
func (o *GetDevicesAirmaxesIDSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get devices airmaxes Id system params
func (o *GetDevicesAirmaxesIDSystemParams) WithID(id string) *GetDevicesAirmaxesIDSystemParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get devices airmaxes Id system params
func (o *GetDevicesAirmaxesIDSystemParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesAirmaxesIDSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
