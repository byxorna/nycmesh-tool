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

// NewPutDevicesEroutersIDDhcpServersServernameParams creates a new PutDevicesEroutersIDDhcpServersServernameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutDevicesEroutersIDDhcpServersServernameParams() *PutDevicesEroutersIDDhcpServersServernameParams {
	return &PutDevicesEroutersIDDhcpServersServernameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutDevicesEroutersIDDhcpServersServernameParamsWithTimeout creates a new PutDevicesEroutersIDDhcpServersServernameParams object
// with the ability to set a timeout on a request.
func NewPutDevicesEroutersIDDhcpServersServernameParamsWithTimeout(timeout time.Duration) *PutDevicesEroutersIDDhcpServersServernameParams {
	return &PutDevicesEroutersIDDhcpServersServernameParams{
		timeout: timeout,
	}
}

// NewPutDevicesEroutersIDDhcpServersServernameParamsWithContext creates a new PutDevicesEroutersIDDhcpServersServernameParams object
// with the ability to set a context for a request.
func NewPutDevicesEroutersIDDhcpServersServernameParamsWithContext(ctx context.Context) *PutDevicesEroutersIDDhcpServersServernameParams {
	return &PutDevicesEroutersIDDhcpServersServernameParams{
		Context: ctx,
	}
}

// NewPutDevicesEroutersIDDhcpServersServernameParamsWithHTTPClient creates a new PutDevicesEroutersIDDhcpServersServernameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutDevicesEroutersIDDhcpServersServernameParamsWithHTTPClient(client *http.Client) *PutDevicesEroutersIDDhcpServersServernameParams {
	return &PutDevicesEroutersIDDhcpServersServernameParams{
		HTTPClient: client,
	}
}

/* PutDevicesEroutersIDDhcpServersServernameParams contains all the parameters to send to the API endpoint
   for the put devices erouters Id dhcp servers servername operation.

   Typically these are written to a http.Request.
*/
type PutDevicesEroutersIDDhcpServersServernameParams struct {

	// Body.
	Body *models.UpdateDHCPServer

	// ID.
	ID string

	// ServerName.
	ServerName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put devices erouters Id dhcp servers servername params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDevicesEroutersIDDhcpServersServernameParams) WithDefaults() *PutDevicesEroutersIDDhcpServersServernameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put devices erouters Id dhcp servers servername params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDevicesEroutersIDDhcpServersServernameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put devices erouters Id dhcp servers servername params
func (o *PutDevicesEroutersIDDhcpServersServernameParams) WithTimeout(timeout time.Duration) *PutDevicesEroutersIDDhcpServersServernameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put devices erouters Id dhcp servers servername params
func (o *PutDevicesEroutersIDDhcpServersServernameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put devices erouters Id dhcp servers servername params
func (o *PutDevicesEroutersIDDhcpServersServernameParams) WithContext(ctx context.Context) *PutDevicesEroutersIDDhcpServersServernameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put devices erouters Id dhcp servers servername params
func (o *PutDevicesEroutersIDDhcpServersServernameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put devices erouters Id dhcp servers servername params
func (o *PutDevicesEroutersIDDhcpServersServernameParams) WithHTTPClient(client *http.Client) *PutDevicesEroutersIDDhcpServersServernameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put devices erouters Id dhcp servers servername params
func (o *PutDevicesEroutersIDDhcpServersServernameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put devices erouters Id dhcp servers servername params
func (o *PutDevicesEroutersIDDhcpServersServernameParams) WithBody(body *models.UpdateDHCPServer) *PutDevicesEroutersIDDhcpServersServernameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put devices erouters Id dhcp servers servername params
func (o *PutDevicesEroutersIDDhcpServersServernameParams) SetBody(body *models.UpdateDHCPServer) {
	o.Body = body
}

// WithID adds the id to the put devices erouters Id dhcp servers servername params
func (o *PutDevicesEroutersIDDhcpServersServernameParams) WithID(id string) *PutDevicesEroutersIDDhcpServersServernameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put devices erouters Id dhcp servers servername params
func (o *PutDevicesEroutersIDDhcpServersServernameParams) SetID(id string) {
	o.ID = id
}

// WithServerName adds the serverName to the put devices erouters Id dhcp servers servername params
func (o *PutDevicesEroutersIDDhcpServersServernameParams) WithServerName(serverName string) *PutDevicesEroutersIDDhcpServersServernameParams {
	o.SetServerName(serverName)
	return o
}

// SetServerName adds the serverName to the put devices erouters Id dhcp servers servername params
func (o *PutDevicesEroutersIDDhcpServersServernameParams) SetServerName(serverName string) {
	o.ServerName = serverName
}

// WriteToRequest writes these params to a swagger request
func (o *PutDevicesEroutersIDDhcpServersServernameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
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
