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

// NewGetNmsServerconfigParams creates a new GetNmsServerconfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNmsServerconfigParams() *GetNmsServerconfigParams {
	return &GetNmsServerconfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNmsServerconfigParamsWithTimeout creates a new GetNmsServerconfigParams object
// with the ability to set a timeout on a request.
func NewGetNmsServerconfigParamsWithTimeout(timeout time.Duration) *GetNmsServerconfigParams {
	return &GetNmsServerconfigParams{
		timeout: timeout,
	}
}

// NewGetNmsServerconfigParamsWithContext creates a new GetNmsServerconfigParams object
// with the ability to set a context for a request.
func NewGetNmsServerconfigParamsWithContext(ctx context.Context) *GetNmsServerconfigParams {
	return &GetNmsServerconfigParams{
		Context: ctx,
	}
}

// NewGetNmsServerconfigParamsWithHTTPClient creates a new GetNmsServerconfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNmsServerconfigParamsWithHTTPClient(client *http.Client) *GetNmsServerconfigParams {
	return &GetNmsServerconfigParams{
		HTTPClient: client,
	}
}

/* GetNmsServerconfigParams contains all the parameters to send to the API endpoint
   for the get nms serverconfig operation.

   Typically these are written to a http.Request.
*/
type GetNmsServerconfigParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get nms serverconfig params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNmsServerconfigParams) WithDefaults() *GetNmsServerconfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nms serverconfig params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNmsServerconfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get nms serverconfig params
func (o *GetNmsServerconfigParams) WithTimeout(timeout time.Duration) *GetNmsServerconfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nms serverconfig params
func (o *GetNmsServerconfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nms serverconfig params
func (o *GetNmsServerconfigParams) WithContext(ctx context.Context) *GetNmsServerconfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nms serverconfig params
func (o *GetNmsServerconfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nms serverconfig params
func (o *GetNmsServerconfigParams) WithHTTPClient(client *http.Client) *GetNmsServerconfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nms serverconfig params
func (o *GetNmsServerconfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetNmsServerconfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}