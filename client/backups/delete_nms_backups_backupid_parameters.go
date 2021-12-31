// Code generated by go-swagger; DO NOT EDIT.

package backups

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

// NewDeleteNmsBackupsBackupidParams creates a new DeleteNmsBackupsBackupidParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNmsBackupsBackupidParams() *DeleteNmsBackupsBackupidParams {
	return &DeleteNmsBackupsBackupidParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNmsBackupsBackupidParamsWithTimeout creates a new DeleteNmsBackupsBackupidParams object
// with the ability to set a timeout on a request.
func NewDeleteNmsBackupsBackupidParamsWithTimeout(timeout time.Duration) *DeleteNmsBackupsBackupidParams {
	return &DeleteNmsBackupsBackupidParams{
		timeout: timeout,
	}
}

// NewDeleteNmsBackupsBackupidParamsWithContext creates a new DeleteNmsBackupsBackupidParams object
// with the ability to set a context for a request.
func NewDeleteNmsBackupsBackupidParamsWithContext(ctx context.Context) *DeleteNmsBackupsBackupidParams {
	return &DeleteNmsBackupsBackupidParams{
		Context: ctx,
	}
}

// NewDeleteNmsBackupsBackupidParamsWithHTTPClient creates a new DeleteNmsBackupsBackupidParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNmsBackupsBackupidParamsWithHTTPClient(client *http.Client) *DeleteNmsBackupsBackupidParams {
	return &DeleteNmsBackupsBackupidParams{
		HTTPClient: client,
	}
}

/* DeleteNmsBackupsBackupidParams contains all the parameters to send to the API endpoint
   for the delete nms backups backupid operation.

   Typically these are written to a http.Request.
*/
type DeleteNmsBackupsBackupidParams struct {

	// BackupID.
	BackupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete nms backups backupid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNmsBackupsBackupidParams) WithDefaults() *DeleteNmsBackupsBackupidParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete nms backups backupid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNmsBackupsBackupidParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete nms backups backupid params
func (o *DeleteNmsBackupsBackupidParams) WithTimeout(timeout time.Duration) *DeleteNmsBackupsBackupidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete nms backups backupid params
func (o *DeleteNmsBackupsBackupidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete nms backups backupid params
func (o *DeleteNmsBackupsBackupidParams) WithContext(ctx context.Context) *DeleteNmsBackupsBackupidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete nms backups backupid params
func (o *DeleteNmsBackupsBackupidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete nms backups backupid params
func (o *DeleteNmsBackupsBackupidParams) WithHTTPClient(client *http.Client) *DeleteNmsBackupsBackupidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete nms backups backupid params
func (o *DeleteNmsBackupsBackupidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackupID adds the backupID to the delete nms backups backupid params
func (o *DeleteNmsBackupsBackupidParams) WithBackupID(backupID string) *DeleteNmsBackupsBackupidParams {
	o.SetBackupID(backupID)
	return o
}

// SetBackupID adds the backupId to the delete nms backups backupid params
func (o *DeleteNmsBackupsBackupidParams) SetBackupID(backupID string) {
	o.BackupID = backupID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNmsBackupsBackupidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param backupId
	if err := r.SetPathParam("backupId", o.BackupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
