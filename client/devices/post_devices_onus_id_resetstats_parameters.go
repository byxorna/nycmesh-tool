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

// NewPostDevicesOnusIDResetstatsParams creates a new PostDevicesOnusIDResetstatsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDevicesOnusIDResetstatsParams() *PostDevicesOnusIDResetstatsParams {
	return &PostDevicesOnusIDResetstatsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesOnusIDResetstatsParamsWithTimeout creates a new PostDevicesOnusIDResetstatsParams object
// with the ability to set a timeout on a request.
func NewPostDevicesOnusIDResetstatsParamsWithTimeout(timeout time.Duration) *PostDevicesOnusIDResetstatsParams {
	return &PostDevicesOnusIDResetstatsParams{
		timeout: timeout,
	}
}

// NewPostDevicesOnusIDResetstatsParamsWithContext creates a new PostDevicesOnusIDResetstatsParams object
// with the ability to set a context for a request.
func NewPostDevicesOnusIDResetstatsParamsWithContext(ctx context.Context) *PostDevicesOnusIDResetstatsParams {
	return &PostDevicesOnusIDResetstatsParams{
		Context: ctx,
	}
}

// NewPostDevicesOnusIDResetstatsParamsWithHTTPClient creates a new PostDevicesOnusIDResetstatsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDevicesOnusIDResetstatsParamsWithHTTPClient(client *http.Client) *PostDevicesOnusIDResetstatsParams {
	return &PostDevicesOnusIDResetstatsParams{
		HTTPClient: client,
	}
}

/* PostDevicesOnusIDResetstatsParams contains all the parameters to send to the API endpoint
   for the post devices onus Id resetstats operation.

   Typically these are written to a http.Request.
*/
type PostDevicesOnusIDResetstatsParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post devices onus Id resetstats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesOnusIDResetstatsParams) WithDefaults() *PostDevicesOnusIDResetstatsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post devices onus Id resetstats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesOnusIDResetstatsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post devices onus Id resetstats params
func (o *PostDevicesOnusIDResetstatsParams) WithTimeout(timeout time.Duration) *PostDevicesOnusIDResetstatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices onus Id resetstats params
func (o *PostDevicesOnusIDResetstatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices onus Id resetstats params
func (o *PostDevicesOnusIDResetstatsParams) WithContext(ctx context.Context) *PostDevicesOnusIDResetstatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices onus Id resetstats params
func (o *PostDevicesOnusIDResetstatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices onus Id resetstats params
func (o *PostDevicesOnusIDResetstatsParams) WithHTTPClient(client *http.Client) *PostDevicesOnusIDResetstatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices onus Id resetstats params
func (o *PostDevicesOnusIDResetstatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post devices onus Id resetstats params
func (o *PostDevicesOnusIDResetstatsParams) WithID(id string) *PostDevicesOnusIDResetstatsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post devices onus Id resetstats params
func (o *PostDevicesOnusIDResetstatsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesOnusIDResetstatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
