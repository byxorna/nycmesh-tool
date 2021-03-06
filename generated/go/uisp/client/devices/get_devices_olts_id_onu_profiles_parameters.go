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

// NewGetDevicesOltsIDOnuProfilesParams creates a new GetDevicesOltsIDOnuProfilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicesOltsIDOnuProfilesParams() *GetDevicesOltsIDOnuProfilesParams {
	return &GetDevicesOltsIDOnuProfilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesOltsIDOnuProfilesParamsWithTimeout creates a new GetDevicesOltsIDOnuProfilesParams object
// with the ability to set a timeout on a request.
func NewGetDevicesOltsIDOnuProfilesParamsWithTimeout(timeout time.Duration) *GetDevicesOltsIDOnuProfilesParams {
	return &GetDevicesOltsIDOnuProfilesParams{
		timeout: timeout,
	}
}

// NewGetDevicesOltsIDOnuProfilesParamsWithContext creates a new GetDevicesOltsIDOnuProfilesParams object
// with the ability to set a context for a request.
func NewGetDevicesOltsIDOnuProfilesParamsWithContext(ctx context.Context) *GetDevicesOltsIDOnuProfilesParams {
	return &GetDevicesOltsIDOnuProfilesParams{
		Context: ctx,
	}
}

// NewGetDevicesOltsIDOnuProfilesParamsWithHTTPClient creates a new GetDevicesOltsIDOnuProfilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicesOltsIDOnuProfilesParamsWithHTTPClient(client *http.Client) *GetDevicesOltsIDOnuProfilesParams {
	return &GetDevicesOltsIDOnuProfilesParams{
		HTTPClient: client,
	}
}

/* GetDevicesOltsIDOnuProfilesParams contains all the parameters to send to the API endpoint
   for the get devices olts Id onu profiles operation.

   Typically these are written to a http.Request.
*/
type GetDevicesOltsIDOnuProfilesParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get devices olts Id onu profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesOltsIDOnuProfilesParams) WithDefaults() *GetDevicesOltsIDOnuProfilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get devices olts Id onu profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesOltsIDOnuProfilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get devices olts Id onu profiles params
func (o *GetDevicesOltsIDOnuProfilesParams) WithTimeout(timeout time.Duration) *GetDevicesOltsIDOnuProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices olts Id onu profiles params
func (o *GetDevicesOltsIDOnuProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices olts Id onu profiles params
func (o *GetDevicesOltsIDOnuProfilesParams) WithContext(ctx context.Context) *GetDevicesOltsIDOnuProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices olts Id onu profiles params
func (o *GetDevicesOltsIDOnuProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices olts Id onu profiles params
func (o *GetDevicesOltsIDOnuProfilesParams) WithHTTPClient(client *http.Client) *GetDevicesOltsIDOnuProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices olts Id onu profiles params
func (o *GetDevicesOltsIDOnuProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get devices olts Id onu profiles params
func (o *GetDevicesOltsIDOnuProfilesParams) WithID(id string) *GetDevicesOltsIDOnuProfilesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get devices olts Id onu profiles params
func (o *GetDevicesOltsIDOnuProfilesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesOltsIDOnuProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
