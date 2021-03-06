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
)

// NewPostDevicesIDLocateStartParams creates a new PostDevicesIDLocateStartParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDevicesIDLocateStartParams() *PostDevicesIDLocateStartParams {
	return &PostDevicesIDLocateStartParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesIDLocateStartParamsWithTimeout creates a new PostDevicesIDLocateStartParams object
// with the ability to set a timeout on a request.
func NewPostDevicesIDLocateStartParamsWithTimeout(timeout time.Duration) *PostDevicesIDLocateStartParams {
	return &PostDevicesIDLocateStartParams{
		timeout: timeout,
	}
}

// NewPostDevicesIDLocateStartParamsWithContext creates a new PostDevicesIDLocateStartParams object
// with the ability to set a context for a request.
func NewPostDevicesIDLocateStartParamsWithContext(ctx context.Context) *PostDevicesIDLocateStartParams {
	return &PostDevicesIDLocateStartParams{
		Context: ctx,
	}
}

// NewPostDevicesIDLocateStartParamsWithHTTPClient creates a new PostDevicesIDLocateStartParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDevicesIDLocateStartParamsWithHTTPClient(client *http.Client) *PostDevicesIDLocateStartParams {
	return &PostDevicesIDLocateStartParams{
		HTTPClient: client,
	}
}

/* PostDevicesIDLocateStartParams contains all the parameters to send to the API endpoint
   for the post devices Id locate start operation.

   Typically these are written to a http.Request.
*/
type PostDevicesIDLocateStartParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post devices Id locate start params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesIDLocateStartParams) WithDefaults() *PostDevicesIDLocateStartParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post devices Id locate start params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesIDLocateStartParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post devices Id locate start params
func (o *PostDevicesIDLocateStartParams) WithTimeout(timeout time.Duration) *PostDevicesIDLocateStartParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices Id locate start params
func (o *PostDevicesIDLocateStartParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices Id locate start params
func (o *PostDevicesIDLocateStartParams) WithContext(ctx context.Context) *PostDevicesIDLocateStartParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices Id locate start params
func (o *PostDevicesIDLocateStartParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices Id locate start params
func (o *PostDevicesIDLocateStartParams) WithHTTPClient(client *http.Client) *PostDevicesIDLocateStartParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices Id locate start params
func (o *PostDevicesIDLocateStartParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post devices Id locate start params
func (o *PostDevicesIDLocateStartParams) WithID(id string) *PostDevicesIDLocateStartParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post devices Id locate start params
func (o *PostDevicesIDLocateStartParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesIDLocateStartParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
