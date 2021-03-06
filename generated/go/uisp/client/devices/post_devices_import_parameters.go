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

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// NewPostDevicesImportParams creates a new PostDevicesImportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDevicesImportParams() *PostDevicesImportParams {
	return &PostDevicesImportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesImportParamsWithTimeout creates a new PostDevicesImportParams object
// with the ability to set a timeout on a request.
func NewPostDevicesImportParamsWithTimeout(timeout time.Duration) *PostDevicesImportParams {
	return &PostDevicesImportParams{
		timeout: timeout,
	}
}

// NewPostDevicesImportParamsWithContext creates a new PostDevicesImportParams object
// with the ability to set a context for a request.
func NewPostDevicesImportParamsWithContext(ctx context.Context) *PostDevicesImportParams {
	return &PostDevicesImportParams{
		Context: ctx,
	}
}

// NewPostDevicesImportParamsWithHTTPClient creates a new PostDevicesImportParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDevicesImportParamsWithHTTPClient(client *http.Client) *PostDevicesImportParams {
	return &PostDevicesImportParams{
		HTTPClient: client,
	}
}

/* PostDevicesImportParams contains all the parameters to send to the API endpoint
   for the post devices import operation.

   Typically these are written to a http.Request.
*/
type PostDevicesImportParams struct {

	// Body.
	Body models.DevicesImport

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post devices import params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesImportParams) WithDefaults() *PostDevicesImportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post devices import params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesImportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post devices import params
func (o *PostDevicesImportParams) WithTimeout(timeout time.Duration) *PostDevicesImportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices import params
func (o *PostDevicesImportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices import params
func (o *PostDevicesImportParams) WithContext(ctx context.Context) *PostDevicesImportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices import params
func (o *PostDevicesImportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices import params
func (o *PostDevicesImportParams) WithHTTPClient(client *http.Client) *PostDevicesImportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices import params
func (o *PostDevicesImportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post devices import params
func (o *PostDevicesImportParams) WithBody(body models.DevicesImport) *PostDevicesImportParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post devices import params
func (o *PostDevicesImportParams) SetBody(body models.DevicesImport) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesImportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
