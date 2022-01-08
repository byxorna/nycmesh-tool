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

// NewPutDevicesIDSystemUsersParams creates a new PutDevicesIDSystemUsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutDevicesIDSystemUsersParams() *PutDevicesIDSystemUsersParams {
	return &PutDevicesIDSystemUsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutDevicesIDSystemUsersParamsWithTimeout creates a new PutDevicesIDSystemUsersParams object
// with the ability to set a timeout on a request.
func NewPutDevicesIDSystemUsersParamsWithTimeout(timeout time.Duration) *PutDevicesIDSystemUsersParams {
	return &PutDevicesIDSystemUsersParams{
		timeout: timeout,
	}
}

// NewPutDevicesIDSystemUsersParamsWithContext creates a new PutDevicesIDSystemUsersParams object
// with the ability to set a context for a request.
func NewPutDevicesIDSystemUsersParamsWithContext(ctx context.Context) *PutDevicesIDSystemUsersParams {
	return &PutDevicesIDSystemUsersParams{
		Context: ctx,
	}
}

// NewPutDevicesIDSystemUsersParamsWithHTTPClient creates a new PutDevicesIDSystemUsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutDevicesIDSystemUsersParamsWithHTTPClient(client *http.Client) *PutDevicesIDSystemUsersParams {
	return &PutDevicesIDSystemUsersParams{
		HTTPClient: client,
	}
}

/* PutDevicesIDSystemUsersParams contains all the parameters to send to the API endpoint
   for the put devices Id system users operation.

   Typically these are written to a http.Request.
*/
type PutDevicesIDSystemUsersParams struct {

	// Body.
	Body *models.Model116

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put devices Id system users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDevicesIDSystemUsersParams) WithDefaults() *PutDevicesIDSystemUsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put devices Id system users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDevicesIDSystemUsersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put devices Id system users params
func (o *PutDevicesIDSystemUsersParams) WithTimeout(timeout time.Duration) *PutDevicesIDSystemUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put devices Id system users params
func (o *PutDevicesIDSystemUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put devices Id system users params
func (o *PutDevicesIDSystemUsersParams) WithContext(ctx context.Context) *PutDevicesIDSystemUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put devices Id system users params
func (o *PutDevicesIDSystemUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put devices Id system users params
func (o *PutDevicesIDSystemUsersParams) WithHTTPClient(client *http.Client) *PutDevicesIDSystemUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put devices Id system users params
func (o *PutDevicesIDSystemUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put devices Id system users params
func (o *PutDevicesIDSystemUsersParams) WithBody(body *models.Model116) *PutDevicesIDSystemUsersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put devices Id system users params
func (o *PutDevicesIDSystemUsersParams) SetBody(body *models.Model116) {
	o.Body = body
}

// WithID adds the id to the put devices Id system users params
func (o *PutDevicesIDSystemUsersParams) WithID(id string) *PutDevicesIDSystemUsersParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put devices Id system users params
func (o *PutDevicesIDSystemUsersParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutDevicesIDSystemUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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