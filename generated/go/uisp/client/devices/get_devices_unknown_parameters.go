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
	"github.com/go-openapi/swag"
)

// NewGetDevicesUnknownParams creates a new GetDevicesUnknownParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicesUnknownParams() *GetDevicesUnknownParams {
	return &GetDevicesUnknownParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesUnknownParamsWithTimeout creates a new GetDevicesUnknownParams object
// with the ability to set a timeout on a request.
func NewGetDevicesUnknownParamsWithTimeout(timeout time.Duration) *GetDevicesUnknownParams {
	return &GetDevicesUnknownParams{
		timeout: timeout,
	}
}

// NewGetDevicesUnknownParamsWithContext creates a new GetDevicesUnknownParams object
// with the ability to set a context for a request.
func NewGetDevicesUnknownParamsWithContext(ctx context.Context) *GetDevicesUnknownParams {
	return &GetDevicesUnknownParams{
		Context: ctx,
	}
}

// NewGetDevicesUnknownParamsWithHTTPClient creates a new GetDevicesUnknownParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicesUnknownParamsWithHTTPClient(client *http.Client) *GetDevicesUnknownParams {
	return &GetDevicesUnknownParams{
		HTTPClient: client,
	}
}

/* GetDevicesUnknownParams contains all the parameters to send to the API endpoint
   for the get devices unknown operation.

   Typically these are written to a http.Request.
*/
type GetDevicesUnknownParams struct {

	/* MinTraffic.

	   Get only devices that produced at least this amount of traffic (in bytes).

	   Default: 102400
	*/
	MinTraffic *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get devices unknown params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesUnknownParams) WithDefaults() *GetDevicesUnknownParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get devices unknown params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesUnknownParams) SetDefaults() {
	var (
		minTrafficDefault = float64(102400)
	)

	val := GetDevicesUnknownParams{
		MinTraffic: &minTrafficDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get devices unknown params
func (o *GetDevicesUnknownParams) WithTimeout(timeout time.Duration) *GetDevicesUnknownParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices unknown params
func (o *GetDevicesUnknownParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices unknown params
func (o *GetDevicesUnknownParams) WithContext(ctx context.Context) *GetDevicesUnknownParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices unknown params
func (o *GetDevicesUnknownParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices unknown params
func (o *GetDevicesUnknownParams) WithHTTPClient(client *http.Client) *GetDevicesUnknownParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices unknown params
func (o *GetDevicesUnknownParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMinTraffic adds the minTraffic to the get devices unknown params
func (o *GetDevicesUnknownParams) WithMinTraffic(minTraffic *float64) *GetDevicesUnknownParams {
	o.SetMinTraffic(minTraffic)
	return o
}

// SetMinTraffic adds the minTraffic to the get devices unknown params
func (o *GetDevicesUnknownParams) SetMinTraffic(minTraffic *float64) {
	o.MinTraffic = minTraffic
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesUnknownParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MinTraffic != nil {

		// query param minTraffic
		var qrMinTraffic float64

		if o.MinTraffic != nil {
			qrMinTraffic = *o.MinTraffic
		}
		qMinTraffic := swag.FormatFloat64(qrMinTraffic)
		if qMinTraffic != "" {

			if err := r.SetQueryParam("minTraffic", qMinTraffic); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
