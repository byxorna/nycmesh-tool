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

// NewGetDevicesAircubesIDFrequencylistsParams creates a new GetDevicesAircubesIDFrequencylistsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicesAircubesIDFrequencylistsParams() *GetDevicesAircubesIDFrequencylistsParams {
	return &GetDevicesAircubesIDFrequencylistsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesAircubesIDFrequencylistsParamsWithTimeout creates a new GetDevicesAircubesIDFrequencylistsParams object
// with the ability to set a timeout on a request.
func NewGetDevicesAircubesIDFrequencylistsParamsWithTimeout(timeout time.Duration) *GetDevicesAircubesIDFrequencylistsParams {
	return &GetDevicesAircubesIDFrequencylistsParams{
		timeout: timeout,
	}
}

// NewGetDevicesAircubesIDFrequencylistsParamsWithContext creates a new GetDevicesAircubesIDFrequencylistsParams object
// with the ability to set a context for a request.
func NewGetDevicesAircubesIDFrequencylistsParamsWithContext(ctx context.Context) *GetDevicesAircubesIDFrequencylistsParams {
	return &GetDevicesAircubesIDFrequencylistsParams{
		Context: ctx,
	}
}

// NewGetDevicesAircubesIDFrequencylistsParamsWithHTTPClient creates a new GetDevicesAircubesIDFrequencylistsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicesAircubesIDFrequencylistsParamsWithHTTPClient(client *http.Client) *GetDevicesAircubesIDFrequencylistsParams {
	return &GetDevicesAircubesIDFrequencylistsParams{
		HTTPClient: client,
	}
}

/* GetDevicesAircubesIDFrequencylistsParams contains all the parameters to send to the API endpoint
   for the get devices aircubes Id frequencylists operation.

   Typically these are written to a http.Request.
*/
type GetDevicesAircubesIDFrequencylistsParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get devices aircubes Id frequencylists params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesAircubesIDFrequencylistsParams) WithDefaults() *GetDevicesAircubesIDFrequencylistsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get devices aircubes Id frequencylists params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesAircubesIDFrequencylistsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get devices aircubes Id frequencylists params
func (o *GetDevicesAircubesIDFrequencylistsParams) WithTimeout(timeout time.Duration) *GetDevicesAircubesIDFrequencylistsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices aircubes Id frequencylists params
func (o *GetDevicesAircubesIDFrequencylistsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices aircubes Id frequencylists params
func (o *GetDevicesAircubesIDFrequencylistsParams) WithContext(ctx context.Context) *GetDevicesAircubesIDFrequencylistsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices aircubes Id frequencylists params
func (o *GetDevicesAircubesIDFrequencylistsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices aircubes Id frequencylists params
func (o *GetDevicesAircubesIDFrequencylistsParams) WithHTTPClient(client *http.Client) *GetDevicesAircubesIDFrequencylistsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices aircubes Id frequencylists params
func (o *GetDevicesAircubesIDFrequencylistsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get devices aircubes Id frequencylists params
func (o *GetDevicesAircubesIDFrequencylistsParams) WithID(id string) *GetDevicesAircubesIDFrequencylistsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get devices aircubes Id frequencylists params
func (o *GetDevicesAircubesIDFrequencylistsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesAircubesIDFrequencylistsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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