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

// NewPutDevicesUisprsIDRouterOspfParams creates a new PutDevicesUisprsIDRouterOspfParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutDevicesUisprsIDRouterOspfParams() *PutDevicesUisprsIDRouterOspfParams {
	return &PutDevicesUisprsIDRouterOspfParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutDevicesUisprsIDRouterOspfParamsWithTimeout creates a new PutDevicesUisprsIDRouterOspfParams object
// with the ability to set a timeout on a request.
func NewPutDevicesUisprsIDRouterOspfParamsWithTimeout(timeout time.Duration) *PutDevicesUisprsIDRouterOspfParams {
	return &PutDevicesUisprsIDRouterOspfParams{
		timeout: timeout,
	}
}

// NewPutDevicesUisprsIDRouterOspfParamsWithContext creates a new PutDevicesUisprsIDRouterOspfParams object
// with the ability to set a context for a request.
func NewPutDevicesUisprsIDRouterOspfParamsWithContext(ctx context.Context) *PutDevicesUisprsIDRouterOspfParams {
	return &PutDevicesUisprsIDRouterOspfParams{
		Context: ctx,
	}
}

// NewPutDevicesUisprsIDRouterOspfParamsWithHTTPClient creates a new PutDevicesUisprsIDRouterOspfParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutDevicesUisprsIDRouterOspfParamsWithHTTPClient(client *http.Client) *PutDevicesUisprsIDRouterOspfParams {
	return &PutDevicesUisprsIDRouterOspfParams{
		HTTPClient: client,
	}
}

/* PutDevicesUisprsIDRouterOspfParams contains all the parameters to send to the API endpoint
   for the put devices uisprs Id router ospf operation.

   Typically these are written to a http.Request.
*/
type PutDevicesUisprsIDRouterOspfParams struct {

	// Body.
	Body *models.Model122

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put devices uisprs Id router ospf params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDevicesUisprsIDRouterOspfParams) WithDefaults() *PutDevicesUisprsIDRouterOspfParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put devices uisprs Id router ospf params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDevicesUisprsIDRouterOspfParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put devices uisprs Id router ospf params
func (o *PutDevicesUisprsIDRouterOspfParams) WithTimeout(timeout time.Duration) *PutDevicesUisprsIDRouterOspfParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put devices uisprs Id router ospf params
func (o *PutDevicesUisprsIDRouterOspfParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put devices uisprs Id router ospf params
func (o *PutDevicesUisprsIDRouterOspfParams) WithContext(ctx context.Context) *PutDevicesUisprsIDRouterOspfParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put devices uisprs Id router ospf params
func (o *PutDevicesUisprsIDRouterOspfParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put devices uisprs Id router ospf params
func (o *PutDevicesUisprsIDRouterOspfParams) WithHTTPClient(client *http.Client) *PutDevicesUisprsIDRouterOspfParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put devices uisprs Id router ospf params
func (o *PutDevicesUisprsIDRouterOspfParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put devices uisprs Id router ospf params
func (o *PutDevicesUisprsIDRouterOspfParams) WithBody(body *models.Model122) *PutDevicesUisprsIDRouterOspfParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put devices uisprs Id router ospf params
func (o *PutDevicesUisprsIDRouterOspfParams) SetBody(body *models.Model122) {
	o.Body = body
}

// WithID adds the id to the put devices uisprs Id router ospf params
func (o *PutDevicesUisprsIDRouterOspfParams) WithID(id string) *PutDevicesUisprsIDRouterOspfParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put devices uisprs Id router ospf params
func (o *PutDevicesUisprsIDRouterOspfParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutDevicesUisprsIDRouterOspfParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
