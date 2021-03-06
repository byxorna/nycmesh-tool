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

// NewGetDevicesDeviceidInterfacesDatalinkAvailableParams creates a new GetDevicesDeviceidInterfacesDatalinkAvailableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicesDeviceidInterfacesDatalinkAvailableParams() *GetDevicesDeviceidInterfacesDatalinkAvailableParams {
	return &GetDevicesDeviceidInterfacesDatalinkAvailableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesDeviceidInterfacesDatalinkAvailableParamsWithTimeout creates a new GetDevicesDeviceidInterfacesDatalinkAvailableParams object
// with the ability to set a timeout on a request.
func NewGetDevicesDeviceidInterfacesDatalinkAvailableParamsWithTimeout(timeout time.Duration) *GetDevicesDeviceidInterfacesDatalinkAvailableParams {
	return &GetDevicesDeviceidInterfacesDatalinkAvailableParams{
		timeout: timeout,
	}
}

// NewGetDevicesDeviceidInterfacesDatalinkAvailableParamsWithContext creates a new GetDevicesDeviceidInterfacesDatalinkAvailableParams object
// with the ability to set a context for a request.
func NewGetDevicesDeviceidInterfacesDatalinkAvailableParamsWithContext(ctx context.Context) *GetDevicesDeviceidInterfacesDatalinkAvailableParams {
	return &GetDevicesDeviceidInterfacesDatalinkAvailableParams{
		Context: ctx,
	}
}

// NewGetDevicesDeviceidInterfacesDatalinkAvailableParamsWithHTTPClient creates a new GetDevicesDeviceidInterfacesDatalinkAvailableParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicesDeviceidInterfacesDatalinkAvailableParamsWithHTTPClient(client *http.Client) *GetDevicesDeviceidInterfacesDatalinkAvailableParams {
	return &GetDevicesDeviceidInterfacesDatalinkAvailableParams{
		HTTPClient: client,
	}
}

/* GetDevicesDeviceidInterfacesDatalinkAvailableParams contains all the parameters to send to the API endpoint
   for the get devices deviceid interfaces datalink available operation.

   Typically these are written to a http.Request.
*/
type GetDevicesDeviceidInterfacesDatalinkAvailableParams struct {

	// DataLinkID.
	DataLinkID *string

	// DeviceID.
	DeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get devices deviceid interfaces datalink available params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesDeviceidInterfacesDatalinkAvailableParams) WithDefaults() *GetDevicesDeviceidInterfacesDatalinkAvailableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get devices deviceid interfaces datalink available params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesDeviceidInterfacesDatalinkAvailableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get devices deviceid interfaces datalink available params
func (o *GetDevicesDeviceidInterfacesDatalinkAvailableParams) WithTimeout(timeout time.Duration) *GetDevicesDeviceidInterfacesDatalinkAvailableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices deviceid interfaces datalink available params
func (o *GetDevicesDeviceidInterfacesDatalinkAvailableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices deviceid interfaces datalink available params
func (o *GetDevicesDeviceidInterfacesDatalinkAvailableParams) WithContext(ctx context.Context) *GetDevicesDeviceidInterfacesDatalinkAvailableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices deviceid interfaces datalink available params
func (o *GetDevicesDeviceidInterfacesDatalinkAvailableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices deviceid interfaces datalink available params
func (o *GetDevicesDeviceidInterfacesDatalinkAvailableParams) WithHTTPClient(client *http.Client) *GetDevicesDeviceidInterfacesDatalinkAvailableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices deviceid interfaces datalink available params
func (o *GetDevicesDeviceidInterfacesDatalinkAvailableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDataLinkID adds the dataLinkID to the get devices deviceid interfaces datalink available params
func (o *GetDevicesDeviceidInterfacesDatalinkAvailableParams) WithDataLinkID(dataLinkID *string) *GetDevicesDeviceidInterfacesDatalinkAvailableParams {
	o.SetDataLinkID(dataLinkID)
	return o
}

// SetDataLinkID adds the dataLinkId to the get devices deviceid interfaces datalink available params
func (o *GetDevicesDeviceidInterfacesDatalinkAvailableParams) SetDataLinkID(dataLinkID *string) {
	o.DataLinkID = dataLinkID
}

// WithDeviceID adds the deviceID to the get devices deviceid interfaces datalink available params
func (o *GetDevicesDeviceidInterfacesDatalinkAvailableParams) WithDeviceID(deviceID string) *GetDevicesDeviceidInterfacesDatalinkAvailableParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get devices deviceid interfaces datalink available params
func (o *GetDevicesDeviceidInterfacesDatalinkAvailableParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesDeviceidInterfacesDatalinkAvailableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DataLinkID != nil {

		// query param dataLinkId
		var qrDataLinkID string

		if o.DataLinkID != nil {
			qrDataLinkID = *o.DataLinkID
		}
		qDataLinkID := qrDataLinkID
		if qDataLinkID != "" {

			if err := r.SetQueryParam("dataLinkId", qDataLinkID); err != nil {
				return err
			}
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
