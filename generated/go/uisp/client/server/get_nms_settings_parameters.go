// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewGetNmsSettingsParams creates a new GetNmsSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNmsSettingsParams() *GetNmsSettingsParams {
	return &GetNmsSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNmsSettingsParamsWithTimeout creates a new GetNmsSettingsParams object
// with the ability to set a timeout on a request.
func NewGetNmsSettingsParamsWithTimeout(timeout time.Duration) *GetNmsSettingsParams {
	return &GetNmsSettingsParams{
		timeout: timeout,
	}
}

// NewGetNmsSettingsParamsWithContext creates a new GetNmsSettingsParams object
// with the ability to set a context for a request.
func NewGetNmsSettingsParamsWithContext(ctx context.Context) *GetNmsSettingsParams {
	return &GetNmsSettingsParams{
		Context: ctx,
	}
}

// NewGetNmsSettingsParamsWithHTTPClient creates a new GetNmsSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNmsSettingsParamsWithHTTPClient(client *http.Client) *GetNmsSettingsParams {
	return &GetNmsSettingsParams{
		HTTPClient: client,
	}
}

/* GetNmsSettingsParams contains all the parameters to send to the API endpoint
   for the get nms settings operation.

   Typically these are written to a http.Request.
*/
type GetNmsSettingsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get nms settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNmsSettingsParams) WithDefaults() *GetNmsSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nms settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNmsSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get nms settings params
func (o *GetNmsSettingsParams) WithTimeout(timeout time.Duration) *GetNmsSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nms settings params
func (o *GetNmsSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nms settings params
func (o *GetNmsSettingsParams) WithContext(ctx context.Context) *GetNmsSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nms settings params
func (o *GetNmsSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nms settings params
func (o *GetNmsSettingsParams) WithHTTPClient(client *http.Client) *GetNmsSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nms settings params
func (o *GetNmsSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetNmsSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
