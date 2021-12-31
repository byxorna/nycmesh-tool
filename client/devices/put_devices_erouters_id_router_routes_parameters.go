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

// NewPutDevicesEroutersIDRouterRoutesParams creates a new PutDevicesEroutersIDRouterRoutesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutDevicesEroutersIDRouterRoutesParams() *PutDevicesEroutersIDRouterRoutesParams {
	return &PutDevicesEroutersIDRouterRoutesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutDevicesEroutersIDRouterRoutesParamsWithTimeout creates a new PutDevicesEroutersIDRouterRoutesParams object
// with the ability to set a timeout on a request.
func NewPutDevicesEroutersIDRouterRoutesParamsWithTimeout(timeout time.Duration) *PutDevicesEroutersIDRouterRoutesParams {
	return &PutDevicesEroutersIDRouterRoutesParams{
		timeout: timeout,
	}
}

// NewPutDevicesEroutersIDRouterRoutesParamsWithContext creates a new PutDevicesEroutersIDRouterRoutesParams object
// with the ability to set a context for a request.
func NewPutDevicesEroutersIDRouterRoutesParamsWithContext(ctx context.Context) *PutDevicesEroutersIDRouterRoutesParams {
	return &PutDevicesEroutersIDRouterRoutesParams{
		Context: ctx,
	}
}

// NewPutDevicesEroutersIDRouterRoutesParamsWithHTTPClient creates a new PutDevicesEroutersIDRouterRoutesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutDevicesEroutersIDRouterRoutesParamsWithHTTPClient(client *http.Client) *PutDevicesEroutersIDRouterRoutesParams {
	return &PutDevicesEroutersIDRouterRoutesParams{
		HTTPClient: client,
	}
}

/* PutDevicesEroutersIDRouterRoutesParams contains all the parameters to send to the API endpoint
   for the put devices erouters Id router routes operation.

   Typically these are written to a http.Request.
*/
type PutDevicesEroutersIDRouterRoutesParams struct {

	// Body.
	Body *models.RouterRouteUpsert

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put devices erouters Id router routes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDevicesEroutersIDRouterRoutesParams) WithDefaults() *PutDevicesEroutersIDRouterRoutesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put devices erouters Id router routes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDevicesEroutersIDRouterRoutesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put devices erouters Id router routes params
func (o *PutDevicesEroutersIDRouterRoutesParams) WithTimeout(timeout time.Duration) *PutDevicesEroutersIDRouterRoutesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put devices erouters Id router routes params
func (o *PutDevicesEroutersIDRouterRoutesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put devices erouters Id router routes params
func (o *PutDevicesEroutersIDRouterRoutesParams) WithContext(ctx context.Context) *PutDevicesEroutersIDRouterRoutesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put devices erouters Id router routes params
func (o *PutDevicesEroutersIDRouterRoutesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put devices erouters Id router routes params
func (o *PutDevicesEroutersIDRouterRoutesParams) WithHTTPClient(client *http.Client) *PutDevicesEroutersIDRouterRoutesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put devices erouters Id router routes params
func (o *PutDevicesEroutersIDRouterRoutesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put devices erouters Id router routes params
func (o *PutDevicesEroutersIDRouterRoutesParams) WithBody(body *models.RouterRouteUpsert) *PutDevicesEroutersIDRouterRoutesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put devices erouters Id router routes params
func (o *PutDevicesEroutersIDRouterRoutesParams) SetBody(body *models.RouterRouteUpsert) {
	o.Body = body
}

// WithID adds the id to the put devices erouters Id router routes params
func (o *PutDevicesEroutersIDRouterRoutesParams) WithID(id string) *PutDevicesEroutersIDRouterRoutesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put devices erouters Id router routes params
func (o *PutDevicesEroutersIDRouterRoutesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutDevicesEroutersIDRouterRoutesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
