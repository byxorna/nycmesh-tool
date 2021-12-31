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

// NewPostDevicesDeviceidInterfacesVlanParams creates a new PostDevicesDeviceidInterfacesVlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDevicesDeviceidInterfacesVlanParams() *PostDevicesDeviceidInterfacesVlanParams {
	return &PostDevicesDeviceidInterfacesVlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesDeviceidInterfacesVlanParamsWithTimeout creates a new PostDevicesDeviceidInterfacesVlanParams object
// with the ability to set a timeout on a request.
func NewPostDevicesDeviceidInterfacesVlanParamsWithTimeout(timeout time.Duration) *PostDevicesDeviceidInterfacesVlanParams {
	return &PostDevicesDeviceidInterfacesVlanParams{
		timeout: timeout,
	}
}

// NewPostDevicesDeviceidInterfacesVlanParamsWithContext creates a new PostDevicesDeviceidInterfacesVlanParams object
// with the ability to set a context for a request.
func NewPostDevicesDeviceidInterfacesVlanParamsWithContext(ctx context.Context) *PostDevicesDeviceidInterfacesVlanParams {
	return &PostDevicesDeviceidInterfacesVlanParams{
		Context: ctx,
	}
}

// NewPostDevicesDeviceidInterfacesVlanParamsWithHTTPClient creates a new PostDevicesDeviceidInterfacesVlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDevicesDeviceidInterfacesVlanParamsWithHTTPClient(client *http.Client) *PostDevicesDeviceidInterfacesVlanParams {
	return &PostDevicesDeviceidInterfacesVlanParams{
		HTTPClient: client,
	}
}

/* PostDevicesDeviceidInterfacesVlanParams contains all the parameters to send to the API endpoint
   for the post devices deviceid interfaces vlan operation.

   Typically these are written to a http.Request.
*/
type PostDevicesDeviceidInterfacesVlanParams struct {

	// Body.
	Body *models.Model97

	// DeviceID.
	DeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post devices deviceid interfaces vlan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesDeviceidInterfacesVlanParams) WithDefaults() *PostDevicesDeviceidInterfacesVlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post devices deviceid interfaces vlan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesDeviceidInterfacesVlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post devices deviceid interfaces vlan params
func (o *PostDevicesDeviceidInterfacesVlanParams) WithTimeout(timeout time.Duration) *PostDevicesDeviceidInterfacesVlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices deviceid interfaces vlan params
func (o *PostDevicesDeviceidInterfacesVlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices deviceid interfaces vlan params
func (o *PostDevicesDeviceidInterfacesVlanParams) WithContext(ctx context.Context) *PostDevicesDeviceidInterfacesVlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices deviceid interfaces vlan params
func (o *PostDevicesDeviceidInterfacesVlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices deviceid interfaces vlan params
func (o *PostDevicesDeviceidInterfacesVlanParams) WithHTTPClient(client *http.Client) *PostDevicesDeviceidInterfacesVlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices deviceid interfaces vlan params
func (o *PostDevicesDeviceidInterfacesVlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post devices deviceid interfaces vlan params
func (o *PostDevicesDeviceidInterfacesVlanParams) WithBody(body *models.Model97) *PostDevicesDeviceidInterfacesVlanParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post devices deviceid interfaces vlan params
func (o *PostDevicesDeviceidInterfacesVlanParams) SetBody(body *models.Model97) {
	o.Body = body
}

// WithDeviceID adds the deviceID to the post devices deviceid interfaces vlan params
func (o *PostDevicesDeviceidInterfacesVlanParams) WithDeviceID(deviceID string) *PostDevicesDeviceidInterfacesVlanParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the post devices deviceid interfaces vlan params
func (o *PostDevicesDeviceidInterfacesVlanParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesDeviceidInterfacesVlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
