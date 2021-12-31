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

// NewPutDevicesAirmaxesIDConfigWirelessParams creates a new PutDevicesAirmaxesIDConfigWirelessParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutDevicesAirmaxesIDConfigWirelessParams() *PutDevicesAirmaxesIDConfigWirelessParams {
	return &PutDevicesAirmaxesIDConfigWirelessParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutDevicesAirmaxesIDConfigWirelessParamsWithTimeout creates a new PutDevicesAirmaxesIDConfigWirelessParams object
// with the ability to set a timeout on a request.
func NewPutDevicesAirmaxesIDConfigWirelessParamsWithTimeout(timeout time.Duration) *PutDevicesAirmaxesIDConfigWirelessParams {
	return &PutDevicesAirmaxesIDConfigWirelessParams{
		timeout: timeout,
	}
}

// NewPutDevicesAirmaxesIDConfigWirelessParamsWithContext creates a new PutDevicesAirmaxesIDConfigWirelessParams object
// with the ability to set a context for a request.
func NewPutDevicesAirmaxesIDConfigWirelessParamsWithContext(ctx context.Context) *PutDevicesAirmaxesIDConfigWirelessParams {
	return &PutDevicesAirmaxesIDConfigWirelessParams{
		Context: ctx,
	}
}

// NewPutDevicesAirmaxesIDConfigWirelessParamsWithHTTPClient creates a new PutDevicesAirmaxesIDConfigWirelessParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutDevicesAirmaxesIDConfigWirelessParamsWithHTTPClient(client *http.Client) *PutDevicesAirmaxesIDConfigWirelessParams {
	return &PutDevicesAirmaxesIDConfigWirelessParams{
		HTTPClient: client,
	}
}

/* PutDevicesAirmaxesIDConfigWirelessParams contains all the parameters to send to the API endpoint
   for the put devices airmaxes Id config wireless operation.

   Typically these are written to a http.Request.
*/
type PutDevicesAirmaxesIDConfigWirelessParams struct {

	// Body.
	Body *models.AirMaxWifiConfig

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put devices airmaxes Id config wireless params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDevicesAirmaxesIDConfigWirelessParams) WithDefaults() *PutDevicesAirmaxesIDConfigWirelessParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put devices airmaxes Id config wireless params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDevicesAirmaxesIDConfigWirelessParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put devices airmaxes Id config wireless params
func (o *PutDevicesAirmaxesIDConfigWirelessParams) WithTimeout(timeout time.Duration) *PutDevicesAirmaxesIDConfigWirelessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put devices airmaxes Id config wireless params
func (o *PutDevicesAirmaxesIDConfigWirelessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put devices airmaxes Id config wireless params
func (o *PutDevicesAirmaxesIDConfigWirelessParams) WithContext(ctx context.Context) *PutDevicesAirmaxesIDConfigWirelessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put devices airmaxes Id config wireless params
func (o *PutDevicesAirmaxesIDConfigWirelessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put devices airmaxes Id config wireless params
func (o *PutDevicesAirmaxesIDConfigWirelessParams) WithHTTPClient(client *http.Client) *PutDevicesAirmaxesIDConfigWirelessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put devices airmaxes Id config wireless params
func (o *PutDevicesAirmaxesIDConfigWirelessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put devices airmaxes Id config wireless params
func (o *PutDevicesAirmaxesIDConfigWirelessParams) WithBody(body *models.AirMaxWifiConfig) *PutDevicesAirmaxesIDConfigWirelessParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put devices airmaxes Id config wireless params
func (o *PutDevicesAirmaxesIDConfigWirelessParams) SetBody(body *models.AirMaxWifiConfig) {
	o.Body = body
}

// WithID adds the id to the put devices airmaxes Id config wireless params
func (o *PutDevicesAirmaxesIDConfigWirelessParams) WithID(id string) *PutDevicesAirmaxesIDConfigWirelessParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put devices airmaxes Id config wireless params
func (o *PutDevicesAirmaxesIDConfigWirelessParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutDevicesAirmaxesIDConfigWirelessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
