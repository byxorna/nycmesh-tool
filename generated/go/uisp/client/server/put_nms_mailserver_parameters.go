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

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// NewPutNmsMailserverParams creates a new PutNmsMailserverParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutNmsMailserverParams() *PutNmsMailserverParams {
	return &PutNmsMailserverParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutNmsMailserverParamsWithTimeout creates a new PutNmsMailserverParams object
// with the ability to set a timeout on a request.
func NewPutNmsMailserverParamsWithTimeout(timeout time.Duration) *PutNmsMailserverParams {
	return &PutNmsMailserverParams{
		timeout: timeout,
	}
}

// NewPutNmsMailserverParamsWithContext creates a new PutNmsMailserverParams object
// with the ability to set a context for a request.
func NewPutNmsMailserverParamsWithContext(ctx context.Context) *PutNmsMailserverParams {
	return &PutNmsMailserverParams{
		Context: ctx,
	}
}

// NewPutNmsMailserverParamsWithHTTPClient creates a new PutNmsMailserverParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutNmsMailserverParamsWithHTTPClient(client *http.Client) *PutNmsMailserverParams {
	return &PutNmsMailserverParams{
		HTTPClient: client,
	}
}

/* PutNmsMailserverParams contains all the parameters to send to the API endpoint
   for the put nms mailserver operation.

   Typically these are written to a http.Request.
*/
type PutNmsMailserverParams struct {

	// Body.
	Body *models.SMTP

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put nms mailserver params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNmsMailserverParams) WithDefaults() *PutNmsMailserverParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put nms mailserver params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNmsMailserverParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put nms mailserver params
func (o *PutNmsMailserverParams) WithTimeout(timeout time.Duration) *PutNmsMailserverParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put nms mailserver params
func (o *PutNmsMailserverParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put nms mailserver params
func (o *PutNmsMailserverParams) WithContext(ctx context.Context) *PutNmsMailserverParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put nms mailserver params
func (o *PutNmsMailserverParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put nms mailserver params
func (o *PutNmsMailserverParams) WithHTTPClient(client *http.Client) *PutNmsMailserverParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put nms mailserver params
func (o *PutNmsMailserverParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put nms mailserver params
func (o *PutNmsMailserverParams) WithBody(body *models.SMTP) *PutNmsMailserverParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put nms mailserver params
func (o *PutNmsMailserverParams) SetBody(body *models.SMTP) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutNmsMailserverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}