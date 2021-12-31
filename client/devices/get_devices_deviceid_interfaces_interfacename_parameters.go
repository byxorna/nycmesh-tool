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

// NewGetDevicesDeviceidInterfacesInterfacenameParams creates a new GetDevicesDeviceidInterfacesInterfacenameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicesDeviceidInterfacesInterfacenameParams() *GetDevicesDeviceidInterfacesInterfacenameParams {
	return &GetDevicesDeviceidInterfacesInterfacenameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesDeviceidInterfacesInterfacenameParamsWithTimeout creates a new GetDevicesDeviceidInterfacesInterfacenameParams object
// with the ability to set a timeout on a request.
func NewGetDevicesDeviceidInterfacesInterfacenameParamsWithTimeout(timeout time.Duration) *GetDevicesDeviceidInterfacesInterfacenameParams {
	return &GetDevicesDeviceidInterfacesInterfacenameParams{
		timeout: timeout,
	}
}

// NewGetDevicesDeviceidInterfacesInterfacenameParamsWithContext creates a new GetDevicesDeviceidInterfacesInterfacenameParams object
// with the ability to set a context for a request.
func NewGetDevicesDeviceidInterfacesInterfacenameParamsWithContext(ctx context.Context) *GetDevicesDeviceidInterfacesInterfacenameParams {
	return &GetDevicesDeviceidInterfacesInterfacenameParams{
		Context: ctx,
	}
}

// NewGetDevicesDeviceidInterfacesInterfacenameParamsWithHTTPClient creates a new GetDevicesDeviceidInterfacesInterfacenameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicesDeviceidInterfacesInterfacenameParamsWithHTTPClient(client *http.Client) *GetDevicesDeviceidInterfacesInterfacenameParams {
	return &GetDevicesDeviceidInterfacesInterfacenameParams{
		HTTPClient: client,
	}
}

/* GetDevicesDeviceidInterfacesInterfacenameParams contains all the parameters to send to the API endpoint
   for the get devices deviceid interfaces interfacename operation.

   Typically these are written to a http.Request.
*/
type GetDevicesDeviceidInterfacesInterfacenameParams struct {

	// DeviceID.
	DeviceID string

	// InterfaceName.
	InterfaceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get devices deviceid interfaces interfacename params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesDeviceidInterfacesInterfacenameParams) WithDefaults() *GetDevicesDeviceidInterfacesInterfacenameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get devices deviceid interfaces interfacename params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesDeviceidInterfacesInterfacenameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get devices deviceid interfaces interfacename params
func (o *GetDevicesDeviceidInterfacesInterfacenameParams) WithTimeout(timeout time.Duration) *GetDevicesDeviceidInterfacesInterfacenameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices deviceid interfaces interfacename params
func (o *GetDevicesDeviceidInterfacesInterfacenameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices deviceid interfaces interfacename params
func (o *GetDevicesDeviceidInterfacesInterfacenameParams) WithContext(ctx context.Context) *GetDevicesDeviceidInterfacesInterfacenameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices deviceid interfaces interfacename params
func (o *GetDevicesDeviceidInterfacesInterfacenameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices deviceid interfaces interfacename params
func (o *GetDevicesDeviceidInterfacesInterfacenameParams) WithHTTPClient(client *http.Client) *GetDevicesDeviceidInterfacesInterfacenameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices deviceid interfaces interfacename params
func (o *GetDevicesDeviceidInterfacesInterfacenameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get devices deviceid interfaces interfacename params
func (o *GetDevicesDeviceidInterfacesInterfacenameParams) WithDeviceID(deviceID string) *GetDevicesDeviceidInterfacesInterfacenameParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get devices deviceid interfaces interfacename params
func (o *GetDevicesDeviceidInterfacesInterfacenameParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithInterfaceName adds the interfaceName to the get devices deviceid interfaces interfacename params
func (o *GetDevicesDeviceidInterfacesInterfacenameParams) WithInterfaceName(interfaceName string) *GetDevicesDeviceidInterfacesInterfacenameParams {
	o.SetInterfaceName(interfaceName)
	return o
}

// SetInterfaceName adds the interfaceName to the get devices deviceid interfaces interfacename params
func (o *GetDevicesDeviceidInterfacesInterfacenameParams) SetInterfaceName(interfaceName string) {
	o.InterfaceName = interfaceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesDeviceidInterfacesInterfacenameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

	// path param interfaceName
	if err := r.SetPathParam("interfaceName", o.InterfaceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
