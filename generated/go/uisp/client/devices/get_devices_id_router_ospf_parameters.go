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

// NewGetDevicesIDRouterOspfParams creates a new GetDevicesIDRouterOspfParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicesIDRouterOspfParams() *GetDevicesIDRouterOspfParams {
	return &GetDevicesIDRouterOspfParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesIDRouterOspfParamsWithTimeout creates a new GetDevicesIDRouterOspfParams object
// with the ability to set a timeout on a request.
func NewGetDevicesIDRouterOspfParamsWithTimeout(timeout time.Duration) *GetDevicesIDRouterOspfParams {
	return &GetDevicesIDRouterOspfParams{
		timeout: timeout,
	}
}

// NewGetDevicesIDRouterOspfParamsWithContext creates a new GetDevicesIDRouterOspfParams object
// with the ability to set a context for a request.
func NewGetDevicesIDRouterOspfParamsWithContext(ctx context.Context) *GetDevicesIDRouterOspfParams {
	return &GetDevicesIDRouterOspfParams{
		Context: ctx,
	}
}

// NewGetDevicesIDRouterOspfParamsWithHTTPClient creates a new GetDevicesIDRouterOspfParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicesIDRouterOspfParamsWithHTTPClient(client *http.Client) *GetDevicesIDRouterOspfParams {
	return &GetDevicesIDRouterOspfParams{
		HTTPClient: client,
	}
}

/* GetDevicesIDRouterOspfParams contains all the parameters to send to the API endpoint
   for the get devices Id router ospf operation.

   Typically these are written to a http.Request.
*/
type GetDevicesIDRouterOspfParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get devices Id router ospf params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesIDRouterOspfParams) WithDefaults() *GetDevicesIDRouterOspfParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get devices Id router ospf params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesIDRouterOspfParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get devices Id router ospf params
func (o *GetDevicesIDRouterOspfParams) WithTimeout(timeout time.Duration) *GetDevicesIDRouterOspfParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices Id router ospf params
func (o *GetDevicesIDRouterOspfParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices Id router ospf params
func (o *GetDevicesIDRouterOspfParams) WithContext(ctx context.Context) *GetDevicesIDRouterOspfParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices Id router ospf params
func (o *GetDevicesIDRouterOspfParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices Id router ospf params
func (o *GetDevicesIDRouterOspfParams) WithHTTPClient(client *http.Client) *GetDevicesIDRouterOspfParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices Id router ospf params
func (o *GetDevicesIDRouterOspfParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get devices Id router ospf params
func (o *GetDevicesIDRouterOspfParams) WithID(id string) *GetDevicesIDRouterOspfParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get devices Id router ospf params
func (o *GetDevicesIDRouterOspfParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesIDRouterOspfParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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