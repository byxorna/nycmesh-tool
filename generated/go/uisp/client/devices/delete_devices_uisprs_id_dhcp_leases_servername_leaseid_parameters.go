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

// NewDeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams creates a new DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams() *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams {
	return &DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParamsWithTimeout creates a new DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams object
// with the ability to set a timeout on a request.
func NewDeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParamsWithTimeout(timeout time.Duration) *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams {
	return &DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams{
		timeout: timeout,
	}
}

// NewDeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParamsWithContext creates a new DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams object
// with the ability to set a context for a request.
func NewDeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParamsWithContext(ctx context.Context) *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams {
	return &DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams{
		Context: ctx,
	}
}

// NewDeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParamsWithHTTPClient creates a new DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParamsWithHTTPClient(client *http.Client) *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams {
	return &DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams{
		HTTPClient: client,
	}
}

/* DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams contains all the parameters to send to the API endpoint
   for the delete devices uisprs Id dhcp leases servername leaseid operation.

   Typically these are written to a http.Request.
*/
type DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams struct {

	// ID.
	ID string

	// LeaseID.
	LeaseID string

	// ServerName.
	ServerName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete devices uisprs Id dhcp leases servername leaseid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams) WithDefaults() *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete devices uisprs Id dhcp leases servername leaseid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete devices uisprs Id dhcp leases servername leaseid params
func (o *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams) WithTimeout(timeout time.Duration) *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete devices uisprs Id dhcp leases servername leaseid params
func (o *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete devices uisprs Id dhcp leases servername leaseid params
func (o *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams) WithContext(ctx context.Context) *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete devices uisprs Id dhcp leases servername leaseid params
func (o *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete devices uisprs Id dhcp leases servername leaseid params
func (o *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams) WithHTTPClient(client *http.Client) *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete devices uisprs Id dhcp leases servername leaseid params
func (o *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete devices uisprs Id dhcp leases servername leaseid params
func (o *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams) WithID(id string) *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete devices uisprs Id dhcp leases servername leaseid params
func (o *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams) SetID(id string) {
	o.ID = id
}

// WithLeaseID adds the leaseID to the delete devices uisprs Id dhcp leases servername leaseid params
func (o *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams) WithLeaseID(leaseID string) *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams {
	o.SetLeaseID(leaseID)
	return o
}

// SetLeaseID adds the leaseId to the delete devices uisprs Id dhcp leases servername leaseid params
func (o *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams) SetLeaseID(leaseID string) {
	o.LeaseID = leaseID
}

// WithServerName adds the serverName to the delete devices uisprs Id dhcp leases servername leaseid params
func (o *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams) WithServerName(serverName string) *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams {
	o.SetServerName(serverName)
	return o
}

// SetServerName adds the serverName to the delete devices uisprs Id dhcp leases servername leaseid params
func (o *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams) SetServerName(serverName string) {
	o.ServerName = serverName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDevicesUisprsIDDhcpLeasesServernameLeaseidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param leaseId
	if err := r.SetPathParam("leaseId", o.LeaseID); err != nil {
		return err
	}

	// path param serverName
	if err := r.SetPathParam("serverName", o.ServerName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}