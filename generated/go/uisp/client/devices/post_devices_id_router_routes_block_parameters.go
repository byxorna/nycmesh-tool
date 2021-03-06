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

// NewPostDevicesIDRouterRoutesBlockParams creates a new PostDevicesIDRouterRoutesBlockParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDevicesIDRouterRoutesBlockParams() *PostDevicesIDRouterRoutesBlockParams {
	return &PostDevicesIDRouterRoutesBlockParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesIDRouterRoutesBlockParamsWithTimeout creates a new PostDevicesIDRouterRoutesBlockParams object
// with the ability to set a timeout on a request.
func NewPostDevicesIDRouterRoutesBlockParamsWithTimeout(timeout time.Duration) *PostDevicesIDRouterRoutesBlockParams {
	return &PostDevicesIDRouterRoutesBlockParams{
		timeout: timeout,
	}
}

// NewPostDevicesIDRouterRoutesBlockParamsWithContext creates a new PostDevicesIDRouterRoutesBlockParams object
// with the ability to set a context for a request.
func NewPostDevicesIDRouterRoutesBlockParamsWithContext(ctx context.Context) *PostDevicesIDRouterRoutesBlockParams {
	return &PostDevicesIDRouterRoutesBlockParams{
		Context: ctx,
	}
}

// NewPostDevicesIDRouterRoutesBlockParamsWithHTTPClient creates a new PostDevicesIDRouterRoutesBlockParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDevicesIDRouterRoutesBlockParamsWithHTTPClient(client *http.Client) *PostDevicesIDRouterRoutesBlockParams {
	return &PostDevicesIDRouterRoutesBlockParams{
		HTTPClient: client,
	}
}

/* PostDevicesIDRouterRoutesBlockParams contains all the parameters to send to the API endpoint
   for the post devices Id router routes block operation.

   Typically these are written to a http.Request.
*/
type PostDevicesIDRouterRoutesBlockParams struct {

	// Body.
	Body *models.RouterRoute1

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post devices Id router routes block params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesIDRouterRoutesBlockParams) WithDefaults() *PostDevicesIDRouterRoutesBlockParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post devices Id router routes block params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesIDRouterRoutesBlockParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post devices Id router routes block params
func (o *PostDevicesIDRouterRoutesBlockParams) WithTimeout(timeout time.Duration) *PostDevicesIDRouterRoutesBlockParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices Id router routes block params
func (o *PostDevicesIDRouterRoutesBlockParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices Id router routes block params
func (o *PostDevicesIDRouterRoutesBlockParams) WithContext(ctx context.Context) *PostDevicesIDRouterRoutesBlockParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices Id router routes block params
func (o *PostDevicesIDRouterRoutesBlockParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices Id router routes block params
func (o *PostDevicesIDRouterRoutesBlockParams) WithHTTPClient(client *http.Client) *PostDevicesIDRouterRoutesBlockParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices Id router routes block params
func (o *PostDevicesIDRouterRoutesBlockParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post devices Id router routes block params
func (o *PostDevicesIDRouterRoutesBlockParams) WithBody(body *models.RouterRoute1) *PostDevicesIDRouterRoutesBlockParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post devices Id router routes block params
func (o *PostDevicesIDRouterRoutesBlockParams) SetBody(body *models.RouterRoute1) {
	o.Body = body
}

// WithID adds the id to the post devices Id router routes block params
func (o *PostDevicesIDRouterRoutesBlockParams) WithID(id string) *PostDevicesIDRouterRoutesBlockParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post devices Id router routes block params
func (o *PostDevicesIDRouterRoutesBlockParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesIDRouterRoutesBlockParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
