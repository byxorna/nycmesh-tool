// Code generated by go-swagger; DO NOT EDIT.

package c_r_m

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

// NewGetCrmRolesParams creates a new GetCrmRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCrmRolesParams() *GetCrmRolesParams {
	return &GetCrmRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCrmRolesParamsWithTimeout creates a new GetCrmRolesParams object
// with the ability to set a timeout on a request.
func NewGetCrmRolesParamsWithTimeout(timeout time.Duration) *GetCrmRolesParams {
	return &GetCrmRolesParams{
		timeout: timeout,
	}
}

// NewGetCrmRolesParamsWithContext creates a new GetCrmRolesParams object
// with the ability to set a context for a request.
func NewGetCrmRolesParamsWithContext(ctx context.Context) *GetCrmRolesParams {
	return &GetCrmRolesParams{
		Context: ctx,
	}
}

// NewGetCrmRolesParamsWithHTTPClient creates a new GetCrmRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCrmRolesParamsWithHTTPClient(client *http.Client) *GetCrmRolesParams {
	return &GetCrmRolesParams{
		HTTPClient: client,
	}
}

/* GetCrmRolesParams contains all the parameters to send to the API endpoint
   for the get crm roles operation.

   Typically these are written to a http.Request.
*/
type GetCrmRolesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get crm roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCrmRolesParams) WithDefaults() *GetCrmRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get crm roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCrmRolesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get crm roles params
func (o *GetCrmRolesParams) WithTimeout(timeout time.Duration) *GetCrmRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get crm roles params
func (o *GetCrmRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get crm roles params
func (o *GetCrmRolesParams) WithContext(ctx context.Context) *GetCrmRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get crm roles params
func (o *GetCrmRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get crm roles params
func (o *GetCrmRolesParams) WithHTTPClient(client *http.Client) *GetCrmRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get crm roles params
func (o *GetCrmRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCrmRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
