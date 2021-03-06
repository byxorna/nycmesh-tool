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

// NewGetDevicesIDParams creates a new GetDevicesIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicesIDParams() *GetDevicesIDParams {
	return &GetDevicesIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesIDParamsWithTimeout creates a new GetDevicesIDParams object
// with the ability to set a timeout on a request.
func NewGetDevicesIDParamsWithTimeout(timeout time.Duration) *GetDevicesIDParams {
	return &GetDevicesIDParams{
		timeout: timeout,
	}
}

// NewGetDevicesIDParamsWithContext creates a new GetDevicesIDParams object
// with the ability to set a context for a request.
func NewGetDevicesIDParamsWithContext(ctx context.Context) *GetDevicesIDParams {
	return &GetDevicesIDParams{
		Context: ctx,
	}
}

// NewGetDevicesIDParamsWithHTTPClient creates a new GetDevicesIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicesIDParamsWithHTTPClient(client *http.Client) *GetDevicesIDParams {
	return &GetDevicesIDParams{
		HTTPClient: client,
	}
}

/* GetDevicesIDParams contains all the parameters to send to the API endpoint
   for the get devices Id operation.

   Typically these are written to a http.Request.
*/
type GetDevicesIDParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get devices Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesIDParams) WithDefaults() *GetDevicesIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get devices Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get devices Id params
func (o *GetDevicesIDParams) WithTimeout(timeout time.Duration) *GetDevicesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices Id params
func (o *GetDevicesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices Id params
func (o *GetDevicesIDParams) WithContext(ctx context.Context) *GetDevicesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices Id params
func (o *GetDevicesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices Id params
func (o *GetDevicesIDParams) WithHTTPClient(client *http.Client) *GetDevicesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices Id params
func (o *GetDevicesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get devices Id params
func (o *GetDevicesIDParams) WithID(id string) *GetDevicesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get devices Id params
func (o *GetDevicesIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
