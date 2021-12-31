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

// NewGetDevicesOnusIDSystemParams creates a new GetDevicesOnusIDSystemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicesOnusIDSystemParams() *GetDevicesOnusIDSystemParams {
	return &GetDevicesOnusIDSystemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesOnusIDSystemParamsWithTimeout creates a new GetDevicesOnusIDSystemParams object
// with the ability to set a timeout on a request.
func NewGetDevicesOnusIDSystemParamsWithTimeout(timeout time.Duration) *GetDevicesOnusIDSystemParams {
	return &GetDevicesOnusIDSystemParams{
		timeout: timeout,
	}
}

// NewGetDevicesOnusIDSystemParamsWithContext creates a new GetDevicesOnusIDSystemParams object
// with the ability to set a context for a request.
func NewGetDevicesOnusIDSystemParamsWithContext(ctx context.Context) *GetDevicesOnusIDSystemParams {
	return &GetDevicesOnusIDSystemParams{
		Context: ctx,
	}
}

// NewGetDevicesOnusIDSystemParamsWithHTTPClient creates a new GetDevicesOnusIDSystemParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicesOnusIDSystemParamsWithHTTPClient(client *http.Client) *GetDevicesOnusIDSystemParams {
	return &GetDevicesOnusIDSystemParams{
		HTTPClient: client,
	}
}

/* GetDevicesOnusIDSystemParams contains all the parameters to send to the API endpoint
   for the get devices onus Id system operation.

   Typically these are written to a http.Request.
*/
type GetDevicesOnusIDSystemParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get devices onus Id system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesOnusIDSystemParams) WithDefaults() *GetDevicesOnusIDSystemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get devices onus Id system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesOnusIDSystemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get devices onus Id system params
func (o *GetDevicesOnusIDSystemParams) WithTimeout(timeout time.Duration) *GetDevicesOnusIDSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices onus Id system params
func (o *GetDevicesOnusIDSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices onus Id system params
func (o *GetDevicesOnusIDSystemParams) WithContext(ctx context.Context) *GetDevicesOnusIDSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices onus Id system params
func (o *GetDevicesOnusIDSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices onus Id system params
func (o *GetDevicesOnusIDSystemParams) WithHTTPClient(client *http.Client) *GetDevicesOnusIDSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices onus Id system params
func (o *GetDevicesOnusIDSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get devices onus Id system params
func (o *GetDevicesOnusIDSystemParams) WithID(id string) *GetDevicesOnusIDSystemParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get devices onus Id system params
func (o *GetDevicesOnusIDSystemParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesOnusIDSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
