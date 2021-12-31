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

// NewGetDevicesSolarbeamsIDConfigSystemParams creates a new GetDevicesSolarbeamsIDConfigSystemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicesSolarbeamsIDConfigSystemParams() *GetDevicesSolarbeamsIDConfigSystemParams {
	return &GetDevicesSolarbeamsIDConfigSystemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesSolarbeamsIDConfigSystemParamsWithTimeout creates a new GetDevicesSolarbeamsIDConfigSystemParams object
// with the ability to set a timeout on a request.
func NewGetDevicesSolarbeamsIDConfigSystemParamsWithTimeout(timeout time.Duration) *GetDevicesSolarbeamsIDConfigSystemParams {
	return &GetDevicesSolarbeamsIDConfigSystemParams{
		timeout: timeout,
	}
}

// NewGetDevicesSolarbeamsIDConfigSystemParamsWithContext creates a new GetDevicesSolarbeamsIDConfigSystemParams object
// with the ability to set a context for a request.
func NewGetDevicesSolarbeamsIDConfigSystemParamsWithContext(ctx context.Context) *GetDevicesSolarbeamsIDConfigSystemParams {
	return &GetDevicesSolarbeamsIDConfigSystemParams{
		Context: ctx,
	}
}

// NewGetDevicesSolarbeamsIDConfigSystemParamsWithHTTPClient creates a new GetDevicesSolarbeamsIDConfigSystemParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicesSolarbeamsIDConfigSystemParamsWithHTTPClient(client *http.Client) *GetDevicesSolarbeamsIDConfigSystemParams {
	return &GetDevicesSolarbeamsIDConfigSystemParams{
		HTTPClient: client,
	}
}

/* GetDevicesSolarbeamsIDConfigSystemParams contains all the parameters to send to the API endpoint
   for the get devices solarbeams Id config system operation.

   Typically these are written to a http.Request.
*/
type GetDevicesSolarbeamsIDConfigSystemParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get devices solarbeams Id config system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesSolarbeamsIDConfigSystemParams) WithDefaults() *GetDevicesSolarbeamsIDConfigSystemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get devices solarbeams Id config system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesSolarbeamsIDConfigSystemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get devices solarbeams Id config system params
func (o *GetDevicesSolarbeamsIDConfigSystemParams) WithTimeout(timeout time.Duration) *GetDevicesSolarbeamsIDConfigSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices solarbeams Id config system params
func (o *GetDevicesSolarbeamsIDConfigSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices solarbeams Id config system params
func (o *GetDevicesSolarbeamsIDConfigSystemParams) WithContext(ctx context.Context) *GetDevicesSolarbeamsIDConfigSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices solarbeams Id config system params
func (o *GetDevicesSolarbeamsIDConfigSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices solarbeams Id config system params
func (o *GetDevicesSolarbeamsIDConfigSystemParams) WithHTTPClient(client *http.Client) *GetDevicesSolarbeamsIDConfigSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices solarbeams Id config system params
func (o *GetDevicesSolarbeamsIDConfigSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get devices solarbeams Id config system params
func (o *GetDevicesSolarbeamsIDConfigSystemParams) WithID(id string) *GetDevicesSolarbeamsIDConfigSystemParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get devices solarbeams Id config system params
func (o *GetDevicesSolarbeamsIDConfigSystemParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesSolarbeamsIDConfigSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
