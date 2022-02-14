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

// NewGetDevicesIDServicesParams creates a new GetDevicesIDServicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicesIDServicesParams() *GetDevicesIDServicesParams {
	return &GetDevicesIDServicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesIDServicesParamsWithTimeout creates a new GetDevicesIDServicesParams object
// with the ability to set a timeout on a request.
func NewGetDevicesIDServicesParamsWithTimeout(timeout time.Duration) *GetDevicesIDServicesParams {
	return &GetDevicesIDServicesParams{
		timeout: timeout,
	}
}

// NewGetDevicesIDServicesParamsWithContext creates a new GetDevicesIDServicesParams object
// with the ability to set a context for a request.
func NewGetDevicesIDServicesParamsWithContext(ctx context.Context) *GetDevicesIDServicesParams {
	return &GetDevicesIDServicesParams{
		Context: ctx,
	}
}

// NewGetDevicesIDServicesParamsWithHTTPClient creates a new GetDevicesIDServicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicesIDServicesParamsWithHTTPClient(client *http.Client) *GetDevicesIDServicesParams {
	return &GetDevicesIDServicesParams{
		HTTPClient: client,
	}
}

/* GetDevicesIDServicesParams contains all the parameters to send to the API endpoint
   for the get devices Id services operation.

   Typically these are written to a http.Request.
*/
type GetDevicesIDServicesParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get devices Id services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesIDServicesParams) WithDefaults() *GetDevicesIDServicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get devices Id services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesIDServicesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get devices Id services params
func (o *GetDevicesIDServicesParams) WithTimeout(timeout time.Duration) *GetDevicesIDServicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices Id services params
func (o *GetDevicesIDServicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices Id services params
func (o *GetDevicesIDServicesParams) WithContext(ctx context.Context) *GetDevicesIDServicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices Id services params
func (o *GetDevicesIDServicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices Id services params
func (o *GetDevicesIDServicesParams) WithHTTPClient(client *http.Client) *GetDevicesIDServicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices Id services params
func (o *GetDevicesIDServicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get devices Id services params
func (o *GetDevicesIDServicesParams) WithID(id string) *GetDevicesIDServicesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get devices Id services params
func (o *GetDevicesIDServicesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesIDServicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
