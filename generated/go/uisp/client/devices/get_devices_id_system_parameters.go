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

// NewGetDevicesIDSystemParams creates a new GetDevicesIDSystemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicesIDSystemParams() *GetDevicesIDSystemParams {
	return &GetDevicesIDSystemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesIDSystemParamsWithTimeout creates a new GetDevicesIDSystemParams object
// with the ability to set a timeout on a request.
func NewGetDevicesIDSystemParamsWithTimeout(timeout time.Duration) *GetDevicesIDSystemParams {
	return &GetDevicesIDSystemParams{
		timeout: timeout,
	}
}

// NewGetDevicesIDSystemParamsWithContext creates a new GetDevicesIDSystemParams object
// with the ability to set a context for a request.
func NewGetDevicesIDSystemParamsWithContext(ctx context.Context) *GetDevicesIDSystemParams {
	return &GetDevicesIDSystemParams{
		Context: ctx,
	}
}

// NewGetDevicesIDSystemParamsWithHTTPClient creates a new GetDevicesIDSystemParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicesIDSystemParamsWithHTTPClient(client *http.Client) *GetDevicesIDSystemParams {
	return &GetDevicesIDSystemParams{
		HTTPClient: client,
	}
}

/* GetDevicesIDSystemParams contains all the parameters to send to the API endpoint
   for the get devices Id system operation.

   Typically these are written to a http.Request.
*/
type GetDevicesIDSystemParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get devices Id system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesIDSystemParams) WithDefaults() *GetDevicesIDSystemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get devices Id system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesIDSystemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get devices Id system params
func (o *GetDevicesIDSystemParams) WithTimeout(timeout time.Duration) *GetDevicesIDSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices Id system params
func (o *GetDevicesIDSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices Id system params
func (o *GetDevicesIDSystemParams) WithContext(ctx context.Context) *GetDevicesIDSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices Id system params
func (o *GetDevicesIDSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices Id system params
func (o *GetDevicesIDSystemParams) WithHTTPClient(client *http.Client) *GetDevicesIDSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices Id system params
func (o *GetDevicesIDSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get devices Id system params
func (o *GetDevicesIDSystemParams) WithID(id string) *GetDevicesIDSystemParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get devices Id system params
func (o *GetDevicesIDSystemParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesIDSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
