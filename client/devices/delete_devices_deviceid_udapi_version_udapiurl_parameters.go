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

// NewDeleteDevicesDeviceidUdapiVersionUdapiurlParams creates a new DeleteDevicesDeviceidUdapiVersionUdapiurlParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDevicesDeviceidUdapiVersionUdapiurlParams() *DeleteDevicesDeviceidUdapiVersionUdapiurlParams {
	return &DeleteDevicesDeviceidUdapiVersionUdapiurlParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDevicesDeviceidUdapiVersionUdapiurlParamsWithTimeout creates a new DeleteDevicesDeviceidUdapiVersionUdapiurlParams object
// with the ability to set a timeout on a request.
func NewDeleteDevicesDeviceidUdapiVersionUdapiurlParamsWithTimeout(timeout time.Duration) *DeleteDevicesDeviceidUdapiVersionUdapiurlParams {
	return &DeleteDevicesDeviceidUdapiVersionUdapiurlParams{
		timeout: timeout,
	}
}

// NewDeleteDevicesDeviceidUdapiVersionUdapiurlParamsWithContext creates a new DeleteDevicesDeviceidUdapiVersionUdapiurlParams object
// with the ability to set a context for a request.
func NewDeleteDevicesDeviceidUdapiVersionUdapiurlParamsWithContext(ctx context.Context) *DeleteDevicesDeviceidUdapiVersionUdapiurlParams {
	return &DeleteDevicesDeviceidUdapiVersionUdapiurlParams{
		Context: ctx,
	}
}

// NewDeleteDevicesDeviceidUdapiVersionUdapiurlParamsWithHTTPClient creates a new DeleteDevicesDeviceidUdapiVersionUdapiurlParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDevicesDeviceidUdapiVersionUdapiurlParamsWithHTTPClient(client *http.Client) *DeleteDevicesDeviceidUdapiVersionUdapiurlParams {
	return &DeleteDevicesDeviceidUdapiVersionUdapiurlParams{
		HTTPClient: client,
	}
}

/* DeleteDevicesDeviceidUdapiVersionUdapiurlParams contains all the parameters to send to the API endpoint
   for the delete devices deviceid udapi version udapiurl operation.

   Typically these are written to a http.Request.
*/
type DeleteDevicesDeviceidUdapiVersionUdapiurlParams struct {

	// DeviceID.
	DeviceID string

	// UdapiURL.
	UdapiURL string

	// Version.
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete devices deviceid udapi version udapiurl params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDevicesDeviceidUdapiVersionUdapiurlParams) WithDefaults() *DeleteDevicesDeviceidUdapiVersionUdapiurlParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete devices deviceid udapi version udapiurl params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDevicesDeviceidUdapiVersionUdapiurlParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete devices deviceid udapi version udapiurl params
func (o *DeleteDevicesDeviceidUdapiVersionUdapiurlParams) WithTimeout(timeout time.Duration) *DeleteDevicesDeviceidUdapiVersionUdapiurlParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete devices deviceid udapi version udapiurl params
func (o *DeleteDevicesDeviceidUdapiVersionUdapiurlParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete devices deviceid udapi version udapiurl params
func (o *DeleteDevicesDeviceidUdapiVersionUdapiurlParams) WithContext(ctx context.Context) *DeleteDevicesDeviceidUdapiVersionUdapiurlParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete devices deviceid udapi version udapiurl params
func (o *DeleteDevicesDeviceidUdapiVersionUdapiurlParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete devices deviceid udapi version udapiurl params
func (o *DeleteDevicesDeviceidUdapiVersionUdapiurlParams) WithHTTPClient(client *http.Client) *DeleteDevicesDeviceidUdapiVersionUdapiurlParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete devices deviceid udapi version udapiurl params
func (o *DeleteDevicesDeviceidUdapiVersionUdapiurlParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the delete devices deviceid udapi version udapiurl params
func (o *DeleteDevicesDeviceidUdapiVersionUdapiurlParams) WithDeviceID(deviceID string) *DeleteDevicesDeviceidUdapiVersionUdapiurlParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the delete devices deviceid udapi version udapiurl params
func (o *DeleteDevicesDeviceidUdapiVersionUdapiurlParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithUdapiURL adds the udapiURL to the delete devices deviceid udapi version udapiurl params
func (o *DeleteDevicesDeviceidUdapiVersionUdapiurlParams) WithUdapiURL(udapiURL string) *DeleteDevicesDeviceidUdapiVersionUdapiurlParams {
	o.SetUdapiURL(udapiURL)
	return o
}

// SetUdapiURL adds the udapiUrl to the delete devices deviceid udapi version udapiurl params
func (o *DeleteDevicesDeviceidUdapiVersionUdapiurlParams) SetUdapiURL(udapiURL string) {
	o.UdapiURL = udapiURL
}

// WithVersion adds the version to the delete devices deviceid udapi version udapiurl params
func (o *DeleteDevicesDeviceidUdapiVersionUdapiurlParams) WithVersion(version string) *DeleteDevicesDeviceidUdapiVersionUdapiurlParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete devices deviceid udapi version udapiurl params
func (o *DeleteDevicesDeviceidUdapiVersionUdapiurlParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDevicesDeviceidUdapiVersionUdapiurlParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

	// path param udapiUrl
	if err := r.SetPathParam("udapiUrl", o.UdapiURL); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}