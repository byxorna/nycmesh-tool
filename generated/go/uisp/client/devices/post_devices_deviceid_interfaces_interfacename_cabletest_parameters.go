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

// NewPostDevicesDeviceidInterfacesInterfacenameCabletestParams creates a new PostDevicesDeviceidInterfacesInterfacenameCabletestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDevicesDeviceidInterfacesInterfacenameCabletestParams() *PostDevicesDeviceidInterfacesInterfacenameCabletestParams {
	return &PostDevicesDeviceidInterfacesInterfacenameCabletestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesDeviceidInterfacesInterfacenameCabletestParamsWithTimeout creates a new PostDevicesDeviceidInterfacesInterfacenameCabletestParams object
// with the ability to set a timeout on a request.
func NewPostDevicesDeviceidInterfacesInterfacenameCabletestParamsWithTimeout(timeout time.Duration) *PostDevicesDeviceidInterfacesInterfacenameCabletestParams {
	return &PostDevicesDeviceidInterfacesInterfacenameCabletestParams{
		timeout: timeout,
	}
}

// NewPostDevicesDeviceidInterfacesInterfacenameCabletestParamsWithContext creates a new PostDevicesDeviceidInterfacesInterfacenameCabletestParams object
// with the ability to set a context for a request.
func NewPostDevicesDeviceidInterfacesInterfacenameCabletestParamsWithContext(ctx context.Context) *PostDevicesDeviceidInterfacesInterfacenameCabletestParams {
	return &PostDevicesDeviceidInterfacesInterfacenameCabletestParams{
		Context: ctx,
	}
}

// NewPostDevicesDeviceidInterfacesInterfacenameCabletestParamsWithHTTPClient creates a new PostDevicesDeviceidInterfacesInterfacenameCabletestParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDevicesDeviceidInterfacesInterfacenameCabletestParamsWithHTTPClient(client *http.Client) *PostDevicesDeviceidInterfacesInterfacenameCabletestParams {
	return &PostDevicesDeviceidInterfacesInterfacenameCabletestParams{
		HTTPClient: client,
	}
}

/* PostDevicesDeviceidInterfacesInterfacenameCabletestParams contains all the parameters to send to the API endpoint
   for the post devices deviceid interfaces interfacename cabletest operation.

   Typically these are written to a http.Request.
*/
type PostDevicesDeviceidInterfacesInterfacenameCabletestParams struct {

	// DeviceID.
	DeviceID string

	// InterfaceName.
	InterfaceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post devices deviceid interfaces interfacename cabletest params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesDeviceidInterfacesInterfacenameCabletestParams) WithDefaults() *PostDevicesDeviceidInterfacesInterfacenameCabletestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post devices deviceid interfaces interfacename cabletest params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesDeviceidInterfacesInterfacenameCabletestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post devices deviceid interfaces interfacename cabletest params
func (o *PostDevicesDeviceidInterfacesInterfacenameCabletestParams) WithTimeout(timeout time.Duration) *PostDevicesDeviceidInterfacesInterfacenameCabletestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices deviceid interfaces interfacename cabletest params
func (o *PostDevicesDeviceidInterfacesInterfacenameCabletestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices deviceid interfaces interfacename cabletest params
func (o *PostDevicesDeviceidInterfacesInterfacenameCabletestParams) WithContext(ctx context.Context) *PostDevicesDeviceidInterfacesInterfacenameCabletestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices deviceid interfaces interfacename cabletest params
func (o *PostDevicesDeviceidInterfacesInterfacenameCabletestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices deviceid interfaces interfacename cabletest params
func (o *PostDevicesDeviceidInterfacesInterfacenameCabletestParams) WithHTTPClient(client *http.Client) *PostDevicesDeviceidInterfacesInterfacenameCabletestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices deviceid interfaces interfacename cabletest params
func (o *PostDevicesDeviceidInterfacesInterfacenameCabletestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the post devices deviceid interfaces interfacename cabletest params
func (o *PostDevicesDeviceidInterfacesInterfacenameCabletestParams) WithDeviceID(deviceID string) *PostDevicesDeviceidInterfacesInterfacenameCabletestParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the post devices deviceid interfaces interfacename cabletest params
func (o *PostDevicesDeviceidInterfacesInterfacenameCabletestParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithInterfaceName adds the interfaceName to the post devices deviceid interfaces interfacename cabletest params
func (o *PostDevicesDeviceidInterfacesInterfacenameCabletestParams) WithInterfaceName(interfaceName string) *PostDevicesDeviceidInterfacesInterfacenameCabletestParams {
	o.SetInterfaceName(interfaceName)
	return o
}

// SetInterfaceName adds the interfaceName to the post devices deviceid interfaces interfacename cabletest params
func (o *PostDevicesDeviceidInterfacesInterfacenameCabletestParams) SetInterfaceName(interfaceName string) {
	o.InterfaceName = interfaceName
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesDeviceidInterfacesInterfacenameCabletestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
