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

	"github.com/byxorna/nycmesh-tool/models"
)

// NewPostDevicesEroutersIDRouterRoutesDeleteParams creates a new PostDevicesEroutersIDRouterRoutesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDevicesEroutersIDRouterRoutesDeleteParams() *PostDevicesEroutersIDRouterRoutesDeleteParams {
	return &PostDevicesEroutersIDRouterRoutesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesEroutersIDRouterRoutesDeleteParamsWithTimeout creates a new PostDevicesEroutersIDRouterRoutesDeleteParams object
// with the ability to set a timeout on a request.
func NewPostDevicesEroutersIDRouterRoutesDeleteParamsWithTimeout(timeout time.Duration) *PostDevicesEroutersIDRouterRoutesDeleteParams {
	return &PostDevicesEroutersIDRouterRoutesDeleteParams{
		timeout: timeout,
	}
}

// NewPostDevicesEroutersIDRouterRoutesDeleteParamsWithContext creates a new PostDevicesEroutersIDRouterRoutesDeleteParams object
// with the ability to set a context for a request.
func NewPostDevicesEroutersIDRouterRoutesDeleteParamsWithContext(ctx context.Context) *PostDevicesEroutersIDRouterRoutesDeleteParams {
	return &PostDevicesEroutersIDRouterRoutesDeleteParams{
		Context: ctx,
	}
}

// NewPostDevicesEroutersIDRouterRoutesDeleteParamsWithHTTPClient creates a new PostDevicesEroutersIDRouterRoutesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDevicesEroutersIDRouterRoutesDeleteParamsWithHTTPClient(client *http.Client) *PostDevicesEroutersIDRouterRoutesDeleteParams {
	return &PostDevicesEroutersIDRouterRoutesDeleteParams{
		HTTPClient: client,
	}
}

/* PostDevicesEroutersIDRouterRoutesDeleteParams contains all the parameters to send to the API endpoint
   for the post devices erouters Id router routes delete operation.

   Typically these are written to a http.Request.
*/
type PostDevicesEroutersIDRouterRoutesDeleteParams struct {

	// Body.
	Body *models.RouterRoute2

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post devices erouters Id router routes delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesEroutersIDRouterRoutesDeleteParams) WithDefaults() *PostDevicesEroutersIDRouterRoutesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post devices erouters Id router routes delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesEroutersIDRouterRoutesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post devices erouters Id router routes delete params
func (o *PostDevicesEroutersIDRouterRoutesDeleteParams) WithTimeout(timeout time.Duration) *PostDevicesEroutersIDRouterRoutesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices erouters Id router routes delete params
func (o *PostDevicesEroutersIDRouterRoutesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices erouters Id router routes delete params
func (o *PostDevicesEroutersIDRouterRoutesDeleteParams) WithContext(ctx context.Context) *PostDevicesEroutersIDRouterRoutesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices erouters Id router routes delete params
func (o *PostDevicesEroutersIDRouterRoutesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices erouters Id router routes delete params
func (o *PostDevicesEroutersIDRouterRoutesDeleteParams) WithHTTPClient(client *http.Client) *PostDevicesEroutersIDRouterRoutesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices erouters Id router routes delete params
func (o *PostDevicesEroutersIDRouterRoutesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post devices erouters Id router routes delete params
func (o *PostDevicesEroutersIDRouterRoutesDeleteParams) WithBody(body *models.RouterRoute2) *PostDevicesEroutersIDRouterRoutesDeleteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post devices erouters Id router routes delete params
func (o *PostDevicesEroutersIDRouterRoutesDeleteParams) SetBody(body *models.RouterRoute2) {
	o.Body = body
}

// WithID adds the id to the post devices erouters Id router routes delete params
func (o *PostDevicesEroutersIDRouterRoutesDeleteParams) WithID(id string) *PostDevicesEroutersIDRouterRoutesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post devices erouters Id router routes delete params
func (o *PostDevicesEroutersIDRouterRoutesDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesEroutersIDRouterRoutesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
