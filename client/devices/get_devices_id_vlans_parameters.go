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

// NewGetDevicesIDVlansParams creates a new GetDevicesIDVlansParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicesIDVlansParams() *GetDevicesIDVlansParams {
	return &GetDevicesIDVlansParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesIDVlansParamsWithTimeout creates a new GetDevicesIDVlansParams object
// with the ability to set a timeout on a request.
func NewGetDevicesIDVlansParamsWithTimeout(timeout time.Duration) *GetDevicesIDVlansParams {
	return &GetDevicesIDVlansParams{
		timeout: timeout,
	}
}

// NewGetDevicesIDVlansParamsWithContext creates a new GetDevicesIDVlansParams object
// with the ability to set a context for a request.
func NewGetDevicesIDVlansParamsWithContext(ctx context.Context) *GetDevicesIDVlansParams {
	return &GetDevicesIDVlansParams{
		Context: ctx,
	}
}

// NewGetDevicesIDVlansParamsWithHTTPClient creates a new GetDevicesIDVlansParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicesIDVlansParamsWithHTTPClient(client *http.Client) *GetDevicesIDVlansParams {
	return &GetDevicesIDVlansParams{
		HTTPClient: client,
	}
}

/* GetDevicesIDVlansParams contains all the parameters to send to the API endpoint
   for the get devices Id vlans operation.

   Typically these are written to a http.Request.
*/
type GetDevicesIDVlansParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get devices Id vlans params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesIDVlansParams) WithDefaults() *GetDevicesIDVlansParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get devices Id vlans params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesIDVlansParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get devices Id vlans params
func (o *GetDevicesIDVlansParams) WithTimeout(timeout time.Duration) *GetDevicesIDVlansParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices Id vlans params
func (o *GetDevicesIDVlansParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices Id vlans params
func (o *GetDevicesIDVlansParams) WithContext(ctx context.Context) *GetDevicesIDVlansParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices Id vlans params
func (o *GetDevicesIDVlansParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices Id vlans params
func (o *GetDevicesIDVlansParams) WithHTTPClient(client *http.Client) *GetDevicesIDVlansParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices Id vlans params
func (o *GetDevicesIDVlansParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get devices Id vlans params
func (o *GetDevicesIDVlansParams) WithID(id string) *GetDevicesIDVlansParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get devices Id vlans params
func (o *GetDevicesIDVlansParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesIDVlansParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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