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

// NewPostDevicesDeviceidBackupsBackupidApplyParams creates a new PostDevicesDeviceidBackupsBackupidApplyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDevicesDeviceidBackupsBackupidApplyParams() *PostDevicesDeviceidBackupsBackupidApplyParams {
	return &PostDevicesDeviceidBackupsBackupidApplyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesDeviceidBackupsBackupidApplyParamsWithTimeout creates a new PostDevicesDeviceidBackupsBackupidApplyParams object
// with the ability to set a timeout on a request.
func NewPostDevicesDeviceidBackupsBackupidApplyParamsWithTimeout(timeout time.Duration) *PostDevicesDeviceidBackupsBackupidApplyParams {
	return &PostDevicesDeviceidBackupsBackupidApplyParams{
		timeout: timeout,
	}
}

// NewPostDevicesDeviceidBackupsBackupidApplyParamsWithContext creates a new PostDevicesDeviceidBackupsBackupidApplyParams object
// with the ability to set a context for a request.
func NewPostDevicesDeviceidBackupsBackupidApplyParamsWithContext(ctx context.Context) *PostDevicesDeviceidBackupsBackupidApplyParams {
	return &PostDevicesDeviceidBackupsBackupidApplyParams{
		Context: ctx,
	}
}

// NewPostDevicesDeviceidBackupsBackupidApplyParamsWithHTTPClient creates a new PostDevicesDeviceidBackupsBackupidApplyParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDevicesDeviceidBackupsBackupidApplyParamsWithHTTPClient(client *http.Client) *PostDevicesDeviceidBackupsBackupidApplyParams {
	return &PostDevicesDeviceidBackupsBackupidApplyParams{
		HTTPClient: client,
	}
}

/* PostDevicesDeviceidBackupsBackupidApplyParams contains all the parameters to send to the API endpoint
   for the post devices deviceid backups backupid apply operation.

   Typically these are written to a http.Request.
*/
type PostDevicesDeviceidBackupsBackupidApplyParams struct {

	// BackupID.
	BackupID string

	// DeviceID.
	DeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post devices deviceid backups backupid apply params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesDeviceidBackupsBackupidApplyParams) WithDefaults() *PostDevicesDeviceidBackupsBackupidApplyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post devices deviceid backups backupid apply params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesDeviceidBackupsBackupidApplyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post devices deviceid backups backupid apply params
func (o *PostDevicesDeviceidBackupsBackupidApplyParams) WithTimeout(timeout time.Duration) *PostDevicesDeviceidBackupsBackupidApplyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices deviceid backups backupid apply params
func (o *PostDevicesDeviceidBackupsBackupidApplyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices deviceid backups backupid apply params
func (o *PostDevicesDeviceidBackupsBackupidApplyParams) WithContext(ctx context.Context) *PostDevicesDeviceidBackupsBackupidApplyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices deviceid backups backupid apply params
func (o *PostDevicesDeviceidBackupsBackupidApplyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices deviceid backups backupid apply params
func (o *PostDevicesDeviceidBackupsBackupidApplyParams) WithHTTPClient(client *http.Client) *PostDevicesDeviceidBackupsBackupidApplyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices deviceid backups backupid apply params
func (o *PostDevicesDeviceidBackupsBackupidApplyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackupID adds the backupID to the post devices deviceid backups backupid apply params
func (o *PostDevicesDeviceidBackupsBackupidApplyParams) WithBackupID(backupID string) *PostDevicesDeviceidBackupsBackupidApplyParams {
	o.SetBackupID(backupID)
	return o
}

// SetBackupID adds the backupId to the post devices deviceid backups backupid apply params
func (o *PostDevicesDeviceidBackupsBackupidApplyParams) SetBackupID(backupID string) {
	o.BackupID = backupID
}

// WithDeviceID adds the deviceID to the post devices deviceid backups backupid apply params
func (o *PostDevicesDeviceidBackupsBackupidApplyParams) WithDeviceID(deviceID string) *PostDevicesDeviceidBackupsBackupidApplyParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the post devices deviceid backups backupid apply params
func (o *PostDevicesDeviceidBackupsBackupidApplyParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesDeviceidBackupsBackupidApplyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param backupId
	if err := r.SetPathParam("backupId", o.BackupID); err != nil {
		return err
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
