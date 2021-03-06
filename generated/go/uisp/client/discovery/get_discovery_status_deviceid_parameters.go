// Code generated by go-swagger; DO NOT EDIT.

package discovery

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

// NewGetDiscoveryStatusDeviceidParams creates a new GetDiscoveryStatusDeviceidParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDiscoveryStatusDeviceidParams() *GetDiscoveryStatusDeviceidParams {
	return &GetDiscoveryStatusDeviceidParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDiscoveryStatusDeviceidParamsWithTimeout creates a new GetDiscoveryStatusDeviceidParams object
// with the ability to set a timeout on a request.
func NewGetDiscoveryStatusDeviceidParamsWithTimeout(timeout time.Duration) *GetDiscoveryStatusDeviceidParams {
	return &GetDiscoveryStatusDeviceidParams{
		timeout: timeout,
	}
}

// NewGetDiscoveryStatusDeviceidParamsWithContext creates a new GetDiscoveryStatusDeviceidParams object
// with the ability to set a context for a request.
func NewGetDiscoveryStatusDeviceidParamsWithContext(ctx context.Context) *GetDiscoveryStatusDeviceidParams {
	return &GetDiscoveryStatusDeviceidParams{
		Context: ctx,
	}
}

// NewGetDiscoveryStatusDeviceidParamsWithHTTPClient creates a new GetDiscoveryStatusDeviceidParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDiscoveryStatusDeviceidParamsWithHTTPClient(client *http.Client) *GetDiscoveryStatusDeviceidParams {
	return &GetDiscoveryStatusDeviceidParams{
		HTTPClient: client,
	}
}

/* GetDiscoveryStatusDeviceidParams contains all the parameters to send to the API endpoint
   for the get discovery status deviceid operation.

   Typically these are written to a http.Request.
*/
type GetDiscoveryStatusDeviceidParams struct {

	// DeviceID.
	DeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get discovery status deviceid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDiscoveryStatusDeviceidParams) WithDefaults() *GetDiscoveryStatusDeviceidParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get discovery status deviceid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDiscoveryStatusDeviceidParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get discovery status deviceid params
func (o *GetDiscoveryStatusDeviceidParams) WithTimeout(timeout time.Duration) *GetDiscoveryStatusDeviceidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get discovery status deviceid params
func (o *GetDiscoveryStatusDeviceidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get discovery status deviceid params
func (o *GetDiscoveryStatusDeviceidParams) WithContext(ctx context.Context) *GetDiscoveryStatusDeviceidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get discovery status deviceid params
func (o *GetDiscoveryStatusDeviceidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get discovery status deviceid params
func (o *GetDiscoveryStatusDeviceidParams) WithHTTPClient(client *http.Client) *GetDiscoveryStatusDeviceidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get discovery status deviceid params
func (o *GetDiscoveryStatusDeviceidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get discovery status deviceid params
func (o *GetDiscoveryStatusDeviceidParams) WithDeviceID(deviceID string) *GetDiscoveryStatusDeviceidParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get discovery status deviceid params
func (o *GetDiscoveryStatusDeviceidParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDiscoveryStatusDeviceidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
