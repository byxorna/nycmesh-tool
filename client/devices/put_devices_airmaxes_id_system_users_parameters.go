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

// NewPutDevicesAirmaxesIDSystemUsersParams creates a new PutDevicesAirmaxesIDSystemUsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutDevicesAirmaxesIDSystemUsersParams() *PutDevicesAirmaxesIDSystemUsersParams {
	return &PutDevicesAirmaxesIDSystemUsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutDevicesAirmaxesIDSystemUsersParamsWithTimeout creates a new PutDevicesAirmaxesIDSystemUsersParams object
// with the ability to set a timeout on a request.
func NewPutDevicesAirmaxesIDSystemUsersParamsWithTimeout(timeout time.Duration) *PutDevicesAirmaxesIDSystemUsersParams {
	return &PutDevicesAirmaxesIDSystemUsersParams{
		timeout: timeout,
	}
}

// NewPutDevicesAirmaxesIDSystemUsersParamsWithContext creates a new PutDevicesAirmaxesIDSystemUsersParams object
// with the ability to set a context for a request.
func NewPutDevicesAirmaxesIDSystemUsersParamsWithContext(ctx context.Context) *PutDevicesAirmaxesIDSystemUsersParams {
	return &PutDevicesAirmaxesIDSystemUsersParams{
		Context: ctx,
	}
}

// NewPutDevicesAirmaxesIDSystemUsersParamsWithHTTPClient creates a new PutDevicesAirmaxesIDSystemUsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutDevicesAirmaxesIDSystemUsersParamsWithHTTPClient(client *http.Client) *PutDevicesAirmaxesIDSystemUsersParams {
	return &PutDevicesAirmaxesIDSystemUsersParams{
		HTTPClient: client,
	}
}

/* PutDevicesAirmaxesIDSystemUsersParams contains all the parameters to send to the API endpoint
   for the put devices airmaxes Id system users operation.

   Typically these are written to a http.Request.
*/
type PutDevicesAirmaxesIDSystemUsersParams struct {

	// Body.
	Body *models.Model118

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put devices airmaxes Id system users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDevicesAirmaxesIDSystemUsersParams) WithDefaults() *PutDevicesAirmaxesIDSystemUsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put devices airmaxes Id system users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDevicesAirmaxesIDSystemUsersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put devices airmaxes Id system users params
func (o *PutDevicesAirmaxesIDSystemUsersParams) WithTimeout(timeout time.Duration) *PutDevicesAirmaxesIDSystemUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put devices airmaxes Id system users params
func (o *PutDevicesAirmaxesIDSystemUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put devices airmaxes Id system users params
func (o *PutDevicesAirmaxesIDSystemUsersParams) WithContext(ctx context.Context) *PutDevicesAirmaxesIDSystemUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put devices airmaxes Id system users params
func (o *PutDevicesAirmaxesIDSystemUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put devices airmaxes Id system users params
func (o *PutDevicesAirmaxesIDSystemUsersParams) WithHTTPClient(client *http.Client) *PutDevicesAirmaxesIDSystemUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put devices airmaxes Id system users params
func (o *PutDevicesAirmaxesIDSystemUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put devices airmaxes Id system users params
func (o *PutDevicesAirmaxesIDSystemUsersParams) WithBody(body *models.Model118) *PutDevicesAirmaxesIDSystemUsersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put devices airmaxes Id system users params
func (o *PutDevicesAirmaxesIDSystemUsersParams) SetBody(body *models.Model118) {
	o.Body = body
}

// WithID adds the id to the put devices airmaxes Id system users params
func (o *PutDevicesAirmaxesIDSystemUsersParams) WithID(id string) *PutDevicesAirmaxesIDSystemUsersParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put devices airmaxes Id system users params
func (o *PutDevicesAirmaxesIDSystemUsersParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutDevicesAirmaxesIDSystemUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
