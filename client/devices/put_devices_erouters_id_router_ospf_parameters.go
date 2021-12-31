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

// NewPutDevicesEroutersIDRouterOspfParams creates a new PutDevicesEroutersIDRouterOspfParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutDevicesEroutersIDRouterOspfParams() *PutDevicesEroutersIDRouterOspfParams {
	return &PutDevicesEroutersIDRouterOspfParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutDevicesEroutersIDRouterOspfParamsWithTimeout creates a new PutDevicesEroutersIDRouterOspfParams object
// with the ability to set a timeout on a request.
func NewPutDevicesEroutersIDRouterOspfParamsWithTimeout(timeout time.Duration) *PutDevicesEroutersIDRouterOspfParams {
	return &PutDevicesEroutersIDRouterOspfParams{
		timeout: timeout,
	}
}

// NewPutDevicesEroutersIDRouterOspfParamsWithContext creates a new PutDevicesEroutersIDRouterOspfParams object
// with the ability to set a context for a request.
func NewPutDevicesEroutersIDRouterOspfParamsWithContext(ctx context.Context) *PutDevicesEroutersIDRouterOspfParams {
	return &PutDevicesEroutersIDRouterOspfParams{
		Context: ctx,
	}
}

// NewPutDevicesEroutersIDRouterOspfParamsWithHTTPClient creates a new PutDevicesEroutersIDRouterOspfParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutDevicesEroutersIDRouterOspfParamsWithHTTPClient(client *http.Client) *PutDevicesEroutersIDRouterOspfParams {
	return &PutDevicesEroutersIDRouterOspfParams{
		HTTPClient: client,
	}
}

/* PutDevicesEroutersIDRouterOspfParams contains all the parameters to send to the API endpoint
   for the put devices erouters Id router ospf operation.

   Typically these are written to a http.Request.
*/
type PutDevicesEroutersIDRouterOspfParams struct {

	// Body.
	Body *models.Model119

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put devices erouters Id router ospf params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDevicesEroutersIDRouterOspfParams) WithDefaults() *PutDevicesEroutersIDRouterOspfParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put devices erouters Id router ospf params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDevicesEroutersIDRouterOspfParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put devices erouters Id router ospf params
func (o *PutDevicesEroutersIDRouterOspfParams) WithTimeout(timeout time.Duration) *PutDevicesEroutersIDRouterOspfParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put devices erouters Id router ospf params
func (o *PutDevicesEroutersIDRouterOspfParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put devices erouters Id router ospf params
func (o *PutDevicesEroutersIDRouterOspfParams) WithContext(ctx context.Context) *PutDevicesEroutersIDRouterOspfParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put devices erouters Id router ospf params
func (o *PutDevicesEroutersIDRouterOspfParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put devices erouters Id router ospf params
func (o *PutDevicesEroutersIDRouterOspfParams) WithHTTPClient(client *http.Client) *PutDevicesEroutersIDRouterOspfParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put devices erouters Id router ospf params
func (o *PutDevicesEroutersIDRouterOspfParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put devices erouters Id router ospf params
func (o *PutDevicesEroutersIDRouterOspfParams) WithBody(body *models.Model119) *PutDevicesEroutersIDRouterOspfParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put devices erouters Id router ospf params
func (o *PutDevicesEroutersIDRouterOspfParams) SetBody(body *models.Model119) {
	o.Body = body
}

// WithID adds the id to the put devices erouters Id router ospf params
func (o *PutDevicesEroutersIDRouterOspfParams) WithID(id string) *PutDevicesEroutersIDRouterOspfParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put devices erouters Id router ospf params
func (o *PutDevicesEroutersIDRouterOspfParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutDevicesEroutersIDRouterOspfParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
