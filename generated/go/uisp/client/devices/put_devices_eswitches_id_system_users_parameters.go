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

// NewPutDevicesEswitchesIDSystemUsersParams creates a new PutDevicesEswitchesIDSystemUsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutDevicesEswitchesIDSystemUsersParams() *PutDevicesEswitchesIDSystemUsersParams {
	return &PutDevicesEswitchesIDSystemUsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutDevicesEswitchesIDSystemUsersParamsWithTimeout creates a new PutDevicesEswitchesIDSystemUsersParams object
// with the ability to set a timeout on a request.
func NewPutDevicesEswitchesIDSystemUsersParamsWithTimeout(timeout time.Duration) *PutDevicesEswitchesIDSystemUsersParams {
	return &PutDevicesEswitchesIDSystemUsersParams{
		timeout: timeout,
	}
}

// NewPutDevicesEswitchesIDSystemUsersParamsWithContext creates a new PutDevicesEswitchesIDSystemUsersParams object
// with the ability to set a context for a request.
func NewPutDevicesEswitchesIDSystemUsersParamsWithContext(ctx context.Context) *PutDevicesEswitchesIDSystemUsersParams {
	return &PutDevicesEswitchesIDSystemUsersParams{
		Context: ctx,
	}
}

// NewPutDevicesEswitchesIDSystemUsersParamsWithHTTPClient creates a new PutDevicesEswitchesIDSystemUsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutDevicesEswitchesIDSystemUsersParamsWithHTTPClient(client *http.Client) *PutDevicesEswitchesIDSystemUsersParams {
	return &PutDevicesEswitchesIDSystemUsersParams{
		HTTPClient: client,
	}
}

/* PutDevicesEswitchesIDSystemUsersParams contains all the parameters to send to the API endpoint
   for the put devices eswitches Id system users operation.

   Typically these are written to a http.Request.
*/
type PutDevicesEswitchesIDSystemUsersParams struct {

	// Body.
	Body *models.Model120

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put devices eswitches Id system users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDevicesEswitchesIDSystemUsersParams) WithDefaults() *PutDevicesEswitchesIDSystemUsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put devices eswitches Id system users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDevicesEswitchesIDSystemUsersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put devices eswitches Id system users params
func (o *PutDevicesEswitchesIDSystemUsersParams) WithTimeout(timeout time.Duration) *PutDevicesEswitchesIDSystemUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put devices eswitches Id system users params
func (o *PutDevicesEswitchesIDSystemUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put devices eswitches Id system users params
func (o *PutDevicesEswitchesIDSystemUsersParams) WithContext(ctx context.Context) *PutDevicesEswitchesIDSystemUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put devices eswitches Id system users params
func (o *PutDevicesEswitchesIDSystemUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put devices eswitches Id system users params
func (o *PutDevicesEswitchesIDSystemUsersParams) WithHTTPClient(client *http.Client) *PutDevicesEswitchesIDSystemUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put devices eswitches Id system users params
func (o *PutDevicesEswitchesIDSystemUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put devices eswitches Id system users params
func (o *PutDevicesEswitchesIDSystemUsersParams) WithBody(body *models.Model120) *PutDevicesEswitchesIDSystemUsersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put devices eswitches Id system users params
func (o *PutDevicesEswitchesIDSystemUsersParams) SetBody(body *models.Model120) {
	o.Body = body
}

// WithID adds the id to the put devices eswitches Id system users params
func (o *PutDevicesEswitchesIDSystemUsersParams) WithID(id string) *PutDevicesEswitchesIDSystemUsersParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put devices eswitches Id system users params
func (o *PutDevicesEswitchesIDSystemUsersParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutDevicesEswitchesIDSystemUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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