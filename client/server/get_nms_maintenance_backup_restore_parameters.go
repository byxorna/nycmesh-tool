// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewGetNmsMaintenanceBackupRestoreParams creates a new GetNmsMaintenanceBackupRestoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNmsMaintenanceBackupRestoreParams() *GetNmsMaintenanceBackupRestoreParams {
	return &GetNmsMaintenanceBackupRestoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNmsMaintenanceBackupRestoreParamsWithTimeout creates a new GetNmsMaintenanceBackupRestoreParams object
// with the ability to set a timeout on a request.
func NewGetNmsMaintenanceBackupRestoreParamsWithTimeout(timeout time.Duration) *GetNmsMaintenanceBackupRestoreParams {
	return &GetNmsMaintenanceBackupRestoreParams{
		timeout: timeout,
	}
}

// NewGetNmsMaintenanceBackupRestoreParamsWithContext creates a new GetNmsMaintenanceBackupRestoreParams object
// with the ability to set a context for a request.
func NewGetNmsMaintenanceBackupRestoreParamsWithContext(ctx context.Context) *GetNmsMaintenanceBackupRestoreParams {
	return &GetNmsMaintenanceBackupRestoreParams{
		Context: ctx,
	}
}

// NewGetNmsMaintenanceBackupRestoreParamsWithHTTPClient creates a new GetNmsMaintenanceBackupRestoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNmsMaintenanceBackupRestoreParamsWithHTTPClient(client *http.Client) *GetNmsMaintenanceBackupRestoreParams {
	return &GetNmsMaintenanceBackupRestoreParams{
		HTTPClient: client,
	}
}

/* GetNmsMaintenanceBackupRestoreParams contains all the parameters to send to the API endpoint
   for the get nms maintenance backup restore operation.

   Typically these are written to a http.Request.
*/
type GetNmsMaintenanceBackupRestoreParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get nms maintenance backup restore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNmsMaintenanceBackupRestoreParams) WithDefaults() *GetNmsMaintenanceBackupRestoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nms maintenance backup restore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNmsMaintenanceBackupRestoreParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get nms maintenance backup restore params
func (o *GetNmsMaintenanceBackupRestoreParams) WithTimeout(timeout time.Duration) *GetNmsMaintenanceBackupRestoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nms maintenance backup restore params
func (o *GetNmsMaintenanceBackupRestoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nms maintenance backup restore params
func (o *GetNmsMaintenanceBackupRestoreParams) WithContext(ctx context.Context) *GetNmsMaintenanceBackupRestoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nms maintenance backup restore params
func (o *GetNmsMaintenanceBackupRestoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nms maintenance backup restore params
func (o *GetNmsMaintenanceBackupRestoreParams) WithHTTPClient(client *http.Client) *GetNmsMaintenanceBackupRestoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nms maintenance backup restore params
func (o *GetNmsMaintenanceBackupRestoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetNmsMaintenanceBackupRestoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
