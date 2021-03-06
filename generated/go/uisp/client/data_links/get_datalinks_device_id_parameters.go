// Code generated by go-swagger; DO NOT EDIT.

package data_links

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

// NewGetDatalinksDeviceIDParams creates a new GetDatalinksDeviceIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDatalinksDeviceIDParams() *GetDatalinksDeviceIDParams {
	return &GetDatalinksDeviceIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDatalinksDeviceIDParamsWithTimeout creates a new GetDatalinksDeviceIDParams object
// with the ability to set a timeout on a request.
func NewGetDatalinksDeviceIDParamsWithTimeout(timeout time.Duration) *GetDatalinksDeviceIDParams {
	return &GetDatalinksDeviceIDParams{
		timeout: timeout,
	}
}

// NewGetDatalinksDeviceIDParamsWithContext creates a new GetDatalinksDeviceIDParams object
// with the ability to set a context for a request.
func NewGetDatalinksDeviceIDParamsWithContext(ctx context.Context) *GetDatalinksDeviceIDParams {
	return &GetDatalinksDeviceIDParams{
		Context: ctx,
	}
}

// NewGetDatalinksDeviceIDParamsWithHTTPClient creates a new GetDatalinksDeviceIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDatalinksDeviceIDParamsWithHTTPClient(client *http.Client) *GetDatalinksDeviceIDParams {
	return &GetDatalinksDeviceIDParams{
		HTTPClient: client,
	}
}

/* GetDatalinksDeviceIDParams contains all the parameters to send to the API endpoint
   for the get datalinks device Id operation.

   Typically these are written to a http.Request.
*/
type GetDatalinksDeviceIDParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get datalinks device Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDatalinksDeviceIDParams) WithDefaults() *GetDatalinksDeviceIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get datalinks device Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDatalinksDeviceIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get datalinks device Id params
func (o *GetDatalinksDeviceIDParams) WithTimeout(timeout time.Duration) *GetDatalinksDeviceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get datalinks device Id params
func (o *GetDatalinksDeviceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get datalinks device Id params
func (o *GetDatalinksDeviceIDParams) WithContext(ctx context.Context) *GetDatalinksDeviceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get datalinks device Id params
func (o *GetDatalinksDeviceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get datalinks device Id params
func (o *GetDatalinksDeviceIDParams) WithHTTPClient(client *http.Client) *GetDatalinksDeviceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get datalinks device Id params
func (o *GetDatalinksDeviceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get datalinks device Id params
func (o *GetDatalinksDeviceIDParams) WithID(id string) *GetDatalinksDeviceIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get datalinks device Id params
func (o *GetDatalinksDeviceIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDatalinksDeviceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
