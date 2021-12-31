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

// NewGetDevicesToughswitchesIDSystemParams creates a new GetDevicesToughswitchesIDSystemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicesToughswitchesIDSystemParams() *GetDevicesToughswitchesIDSystemParams {
	return &GetDevicesToughswitchesIDSystemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesToughswitchesIDSystemParamsWithTimeout creates a new GetDevicesToughswitchesIDSystemParams object
// with the ability to set a timeout on a request.
func NewGetDevicesToughswitchesIDSystemParamsWithTimeout(timeout time.Duration) *GetDevicesToughswitchesIDSystemParams {
	return &GetDevicesToughswitchesIDSystemParams{
		timeout: timeout,
	}
}

// NewGetDevicesToughswitchesIDSystemParamsWithContext creates a new GetDevicesToughswitchesIDSystemParams object
// with the ability to set a context for a request.
func NewGetDevicesToughswitchesIDSystemParamsWithContext(ctx context.Context) *GetDevicesToughswitchesIDSystemParams {
	return &GetDevicesToughswitchesIDSystemParams{
		Context: ctx,
	}
}

// NewGetDevicesToughswitchesIDSystemParamsWithHTTPClient creates a new GetDevicesToughswitchesIDSystemParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicesToughswitchesIDSystemParamsWithHTTPClient(client *http.Client) *GetDevicesToughswitchesIDSystemParams {
	return &GetDevicesToughswitchesIDSystemParams{
		HTTPClient: client,
	}
}

/* GetDevicesToughswitchesIDSystemParams contains all the parameters to send to the API endpoint
   for the get devices toughswitches Id system operation.

   Typically these are written to a http.Request.
*/
type GetDevicesToughswitchesIDSystemParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get devices toughswitches Id system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesToughswitchesIDSystemParams) WithDefaults() *GetDevicesToughswitchesIDSystemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get devices toughswitches Id system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesToughswitchesIDSystemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get devices toughswitches Id system params
func (o *GetDevicesToughswitchesIDSystemParams) WithTimeout(timeout time.Duration) *GetDevicesToughswitchesIDSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices toughswitches Id system params
func (o *GetDevicesToughswitchesIDSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices toughswitches Id system params
func (o *GetDevicesToughswitchesIDSystemParams) WithContext(ctx context.Context) *GetDevicesToughswitchesIDSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices toughswitches Id system params
func (o *GetDevicesToughswitchesIDSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices toughswitches Id system params
func (o *GetDevicesToughswitchesIDSystemParams) WithHTTPClient(client *http.Client) *GetDevicesToughswitchesIDSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices toughswitches Id system params
func (o *GetDevicesToughswitchesIDSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get devices toughswitches Id system params
func (o *GetDevicesToughswitchesIDSystemParams) WithID(id string) *GetDevicesToughswitchesIDSystemParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get devices toughswitches Id system params
func (o *GetDevicesToughswitchesIDSystemParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesToughswitchesIDSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
