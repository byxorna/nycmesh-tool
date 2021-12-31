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

// NewPostDevicesIDRouterRoutesDeleteParams creates a new PostDevicesIDRouterRoutesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDevicesIDRouterRoutesDeleteParams() *PostDevicesIDRouterRoutesDeleteParams {
	return &PostDevicesIDRouterRoutesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesIDRouterRoutesDeleteParamsWithTimeout creates a new PostDevicesIDRouterRoutesDeleteParams object
// with the ability to set a timeout on a request.
func NewPostDevicesIDRouterRoutesDeleteParamsWithTimeout(timeout time.Duration) *PostDevicesIDRouterRoutesDeleteParams {
	return &PostDevicesIDRouterRoutesDeleteParams{
		timeout: timeout,
	}
}

// NewPostDevicesIDRouterRoutesDeleteParamsWithContext creates a new PostDevicesIDRouterRoutesDeleteParams object
// with the ability to set a context for a request.
func NewPostDevicesIDRouterRoutesDeleteParamsWithContext(ctx context.Context) *PostDevicesIDRouterRoutesDeleteParams {
	return &PostDevicesIDRouterRoutesDeleteParams{
		Context: ctx,
	}
}

// NewPostDevicesIDRouterRoutesDeleteParamsWithHTTPClient creates a new PostDevicesIDRouterRoutesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDevicesIDRouterRoutesDeleteParamsWithHTTPClient(client *http.Client) *PostDevicesIDRouterRoutesDeleteParams {
	return &PostDevicesIDRouterRoutesDeleteParams{
		HTTPClient: client,
	}
}

/* PostDevicesIDRouterRoutesDeleteParams contains all the parameters to send to the API endpoint
   for the post devices Id router routes delete operation.

   Typically these are written to a http.Request.
*/
type PostDevicesIDRouterRoutesDeleteParams struct {

	// Body.
	Body *models.RouterRoute1

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post devices Id router routes delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesIDRouterRoutesDeleteParams) WithDefaults() *PostDevicesIDRouterRoutesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post devices Id router routes delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesIDRouterRoutesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post devices Id router routes delete params
func (o *PostDevicesIDRouterRoutesDeleteParams) WithTimeout(timeout time.Duration) *PostDevicesIDRouterRoutesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices Id router routes delete params
func (o *PostDevicesIDRouterRoutesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices Id router routes delete params
func (o *PostDevicesIDRouterRoutesDeleteParams) WithContext(ctx context.Context) *PostDevicesIDRouterRoutesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices Id router routes delete params
func (o *PostDevicesIDRouterRoutesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices Id router routes delete params
func (o *PostDevicesIDRouterRoutesDeleteParams) WithHTTPClient(client *http.Client) *PostDevicesIDRouterRoutesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices Id router routes delete params
func (o *PostDevicesIDRouterRoutesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post devices Id router routes delete params
func (o *PostDevicesIDRouterRoutesDeleteParams) WithBody(body *models.RouterRoute1) *PostDevicesIDRouterRoutesDeleteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post devices Id router routes delete params
func (o *PostDevicesIDRouterRoutesDeleteParams) SetBody(body *models.RouterRoute1) {
	o.Body = body
}

// WithID adds the id to the post devices Id router routes delete params
func (o *PostDevicesIDRouterRoutesDeleteParams) WithID(id string) *PostDevicesIDRouterRoutesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post devices Id router routes delete params
func (o *PostDevicesIDRouterRoutesDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesIDRouterRoutesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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