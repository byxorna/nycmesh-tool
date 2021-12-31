// Code generated by go-swagger; DO NOT EDIT.

package firmware

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

// NewPostFirmwaresDeleteParams creates a new PostFirmwaresDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostFirmwaresDeleteParams() *PostFirmwaresDeleteParams {
	return &PostFirmwaresDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostFirmwaresDeleteParamsWithTimeout creates a new PostFirmwaresDeleteParams object
// with the ability to set a timeout on a request.
func NewPostFirmwaresDeleteParamsWithTimeout(timeout time.Duration) *PostFirmwaresDeleteParams {
	return &PostFirmwaresDeleteParams{
		timeout: timeout,
	}
}

// NewPostFirmwaresDeleteParamsWithContext creates a new PostFirmwaresDeleteParams object
// with the ability to set a context for a request.
func NewPostFirmwaresDeleteParamsWithContext(ctx context.Context) *PostFirmwaresDeleteParams {
	return &PostFirmwaresDeleteParams{
		Context: ctx,
	}
}

// NewPostFirmwaresDeleteParamsWithHTTPClient creates a new PostFirmwaresDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostFirmwaresDeleteParamsWithHTTPClient(client *http.Client) *PostFirmwaresDeleteParams {
	return &PostFirmwaresDeleteParams{
		HTTPClient: client,
	}
}

/* PostFirmwaresDeleteParams contains all the parameters to send to the API endpoint
   for the post firmwares delete operation.

   Typically these are written to a http.Request.
*/
type PostFirmwaresDeleteParams struct {

	// Body.
	Body models.ListOfFirmwareIds

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post firmwares delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostFirmwaresDeleteParams) WithDefaults() *PostFirmwaresDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post firmwares delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostFirmwaresDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post firmwares delete params
func (o *PostFirmwaresDeleteParams) WithTimeout(timeout time.Duration) *PostFirmwaresDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post firmwares delete params
func (o *PostFirmwaresDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post firmwares delete params
func (o *PostFirmwaresDeleteParams) WithContext(ctx context.Context) *PostFirmwaresDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post firmwares delete params
func (o *PostFirmwaresDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post firmwares delete params
func (o *PostFirmwaresDeleteParams) WithHTTPClient(client *http.Client) *PostFirmwaresDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post firmwares delete params
func (o *PostFirmwaresDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post firmwares delete params
func (o *PostFirmwaresDeleteParams) WithBody(body models.ListOfFirmwareIds) *PostFirmwaresDeleteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post firmwares delete params
func (o *PostFirmwaresDeleteParams) SetBody(body models.ListOfFirmwareIds) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostFirmwaresDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
