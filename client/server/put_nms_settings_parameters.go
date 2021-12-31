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

	"github.com/byxorna/nycmesh-tool/models"
)

// NewPutNmsSettingsParams creates a new PutNmsSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutNmsSettingsParams() *PutNmsSettingsParams {
	return &PutNmsSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutNmsSettingsParamsWithTimeout creates a new PutNmsSettingsParams object
// with the ability to set a timeout on a request.
func NewPutNmsSettingsParamsWithTimeout(timeout time.Duration) *PutNmsSettingsParams {
	return &PutNmsSettingsParams{
		timeout: timeout,
	}
}

// NewPutNmsSettingsParamsWithContext creates a new PutNmsSettingsParams object
// with the ability to set a context for a request.
func NewPutNmsSettingsParamsWithContext(ctx context.Context) *PutNmsSettingsParams {
	return &PutNmsSettingsParams{
		Context: ctx,
	}
}

// NewPutNmsSettingsParamsWithHTTPClient creates a new PutNmsSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutNmsSettingsParamsWithHTTPClient(client *http.Client) *PutNmsSettingsParams {
	return &PutNmsSettingsParams{
		HTTPClient: client,
	}
}

/* PutNmsSettingsParams contains all the parameters to send to the API endpoint
   for the put nms settings operation.

   Typically these are written to a http.Request.
*/
type PutNmsSettingsParams struct {

	// Body.
	Body *models.NmsUpdateSettings

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put nms settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNmsSettingsParams) WithDefaults() *PutNmsSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put nms settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNmsSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put nms settings params
func (o *PutNmsSettingsParams) WithTimeout(timeout time.Duration) *PutNmsSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put nms settings params
func (o *PutNmsSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put nms settings params
func (o *PutNmsSettingsParams) WithContext(ctx context.Context) *PutNmsSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put nms settings params
func (o *PutNmsSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put nms settings params
func (o *PutNmsSettingsParams) WithHTTPClient(client *http.Client) *PutNmsSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put nms settings params
func (o *PutNmsSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put nms settings params
func (o *PutNmsSettingsParams) WithBody(body *models.NmsUpdateSettings) *PutNmsSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put nms settings params
func (o *PutNmsSettingsParams) SetBody(body *models.NmsUpdateSettings) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutNmsSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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