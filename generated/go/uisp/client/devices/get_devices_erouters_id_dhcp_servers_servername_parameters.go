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

// NewGetDevicesEroutersIDDhcpServersServernameParams creates a new GetDevicesEroutersIDDhcpServersServernameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicesEroutersIDDhcpServersServernameParams() *GetDevicesEroutersIDDhcpServersServernameParams {
	return &GetDevicesEroutersIDDhcpServersServernameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesEroutersIDDhcpServersServernameParamsWithTimeout creates a new GetDevicesEroutersIDDhcpServersServernameParams object
// with the ability to set a timeout on a request.
func NewGetDevicesEroutersIDDhcpServersServernameParamsWithTimeout(timeout time.Duration) *GetDevicesEroutersIDDhcpServersServernameParams {
	return &GetDevicesEroutersIDDhcpServersServernameParams{
		timeout: timeout,
	}
}

// NewGetDevicesEroutersIDDhcpServersServernameParamsWithContext creates a new GetDevicesEroutersIDDhcpServersServernameParams object
// with the ability to set a context for a request.
func NewGetDevicesEroutersIDDhcpServersServernameParamsWithContext(ctx context.Context) *GetDevicesEroutersIDDhcpServersServernameParams {
	return &GetDevicesEroutersIDDhcpServersServernameParams{
		Context: ctx,
	}
}

// NewGetDevicesEroutersIDDhcpServersServernameParamsWithHTTPClient creates a new GetDevicesEroutersIDDhcpServersServernameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicesEroutersIDDhcpServersServernameParamsWithHTTPClient(client *http.Client) *GetDevicesEroutersIDDhcpServersServernameParams {
	return &GetDevicesEroutersIDDhcpServersServernameParams{
		HTTPClient: client,
	}
}

/* GetDevicesEroutersIDDhcpServersServernameParams contains all the parameters to send to the API endpoint
   for the get devices erouters Id dhcp servers servername operation.

   Typically these are written to a http.Request.
*/
type GetDevicesEroutersIDDhcpServersServernameParams struct {

	// ID.
	ID string

	// ServerName.
	ServerName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get devices erouters Id dhcp servers servername params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesEroutersIDDhcpServersServernameParams) WithDefaults() *GetDevicesEroutersIDDhcpServersServernameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get devices erouters Id dhcp servers servername params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesEroutersIDDhcpServersServernameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get devices erouters Id dhcp servers servername params
func (o *GetDevicesEroutersIDDhcpServersServernameParams) WithTimeout(timeout time.Duration) *GetDevicesEroutersIDDhcpServersServernameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices erouters Id dhcp servers servername params
func (o *GetDevicesEroutersIDDhcpServersServernameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices erouters Id dhcp servers servername params
func (o *GetDevicesEroutersIDDhcpServersServernameParams) WithContext(ctx context.Context) *GetDevicesEroutersIDDhcpServersServernameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices erouters Id dhcp servers servername params
func (o *GetDevicesEroutersIDDhcpServersServernameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices erouters Id dhcp servers servername params
func (o *GetDevicesEroutersIDDhcpServersServernameParams) WithHTTPClient(client *http.Client) *GetDevicesEroutersIDDhcpServersServernameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices erouters Id dhcp servers servername params
func (o *GetDevicesEroutersIDDhcpServersServernameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get devices erouters Id dhcp servers servername params
func (o *GetDevicesEroutersIDDhcpServersServernameParams) WithID(id string) *GetDevicesEroutersIDDhcpServersServernameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get devices erouters Id dhcp servers servername params
func (o *GetDevicesEroutersIDDhcpServersServernameParams) SetID(id string) {
	o.ID = id
}

// WithServerName adds the serverName to the get devices erouters Id dhcp servers servername params
func (o *GetDevicesEroutersIDDhcpServersServernameParams) WithServerName(serverName string) *GetDevicesEroutersIDDhcpServersServernameParams {
	o.SetServerName(serverName)
	return o
}

// SetServerName adds the serverName to the get devices erouters Id dhcp servers servername params
func (o *GetDevicesEroutersIDDhcpServersServernameParams) SetServerName(serverName string) {
	o.ServerName = serverName
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesEroutersIDDhcpServersServernameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
