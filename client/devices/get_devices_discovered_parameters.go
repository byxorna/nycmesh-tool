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

// NewGetDevicesDiscoveredParams creates a new GetDevicesDiscoveredParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicesDiscoveredParams() *GetDevicesDiscoveredParams {
	return &GetDevicesDiscoveredParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesDiscoveredParamsWithTimeout creates a new GetDevicesDiscoveredParams object
// with the ability to set a timeout on a request.
func NewGetDevicesDiscoveredParamsWithTimeout(timeout time.Duration) *GetDevicesDiscoveredParams {
	return &GetDevicesDiscoveredParams{
		timeout: timeout,
	}
}

// NewGetDevicesDiscoveredParamsWithContext creates a new GetDevicesDiscoveredParams object
// with the ability to set a context for a request.
func NewGetDevicesDiscoveredParamsWithContext(ctx context.Context) *GetDevicesDiscoveredParams {
	return &GetDevicesDiscoveredParams{
		Context: ctx,
	}
}

// NewGetDevicesDiscoveredParamsWithHTTPClient creates a new GetDevicesDiscoveredParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicesDiscoveredParamsWithHTTPClient(client *http.Client) *GetDevicesDiscoveredParams {
	return &GetDevicesDiscoveredParams{
		HTTPClient: client,
	}
}

/* GetDevicesDiscoveredParams contains all the parameters to send to the API endpoint
   for the get devices discovered operation.

   Typically these are written to a http.Request.
*/
type GetDevicesDiscoveredParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get devices discovered params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesDiscoveredParams) WithDefaults() *GetDevicesDiscoveredParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get devices discovered params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesDiscoveredParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get devices discovered params
func (o *GetDevicesDiscoveredParams) WithTimeout(timeout time.Duration) *GetDevicesDiscoveredParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices discovered params
func (o *GetDevicesDiscoveredParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices discovered params
func (o *GetDevicesDiscoveredParams) WithContext(ctx context.Context) *GetDevicesDiscoveredParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices discovered params
func (o *GetDevicesDiscoveredParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices discovered params
func (o *GetDevicesDiscoveredParams) WithHTTPClient(client *http.Client) *GetDevicesDiscoveredParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices discovered params
func (o *GetDevicesDiscoveredParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesDiscoveredParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}