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

// NewPostDevicesUisprsIDDhcpServersParams creates a new PostDevicesUisprsIDDhcpServersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDevicesUisprsIDDhcpServersParams() *PostDevicesUisprsIDDhcpServersParams {
	return &PostDevicesUisprsIDDhcpServersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesUisprsIDDhcpServersParamsWithTimeout creates a new PostDevicesUisprsIDDhcpServersParams object
// with the ability to set a timeout on a request.
func NewPostDevicesUisprsIDDhcpServersParamsWithTimeout(timeout time.Duration) *PostDevicesUisprsIDDhcpServersParams {
	return &PostDevicesUisprsIDDhcpServersParams{
		timeout: timeout,
	}
}

// NewPostDevicesUisprsIDDhcpServersParamsWithContext creates a new PostDevicesUisprsIDDhcpServersParams object
// with the ability to set a context for a request.
func NewPostDevicesUisprsIDDhcpServersParamsWithContext(ctx context.Context) *PostDevicesUisprsIDDhcpServersParams {
	return &PostDevicesUisprsIDDhcpServersParams{
		Context: ctx,
	}
}

// NewPostDevicesUisprsIDDhcpServersParamsWithHTTPClient creates a new PostDevicesUisprsIDDhcpServersParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDevicesUisprsIDDhcpServersParamsWithHTTPClient(client *http.Client) *PostDevicesUisprsIDDhcpServersParams {
	return &PostDevicesUisprsIDDhcpServersParams{
		HTTPClient: client,
	}
}

/* PostDevicesUisprsIDDhcpServersParams contains all the parameters to send to the API endpoint
   for the post devices uisprs Id dhcp servers operation.

   Typically these are written to a http.Request.
*/
type PostDevicesUisprsIDDhcpServersParams struct {

	// Body.
	Body *models.CreateDHCPServer1

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post devices uisprs Id dhcp servers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesUisprsIDDhcpServersParams) WithDefaults() *PostDevicesUisprsIDDhcpServersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post devices uisprs Id dhcp servers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesUisprsIDDhcpServersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post devices uisprs Id dhcp servers params
func (o *PostDevicesUisprsIDDhcpServersParams) WithTimeout(timeout time.Duration) *PostDevicesUisprsIDDhcpServersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices uisprs Id dhcp servers params
func (o *PostDevicesUisprsIDDhcpServersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices uisprs Id dhcp servers params
func (o *PostDevicesUisprsIDDhcpServersParams) WithContext(ctx context.Context) *PostDevicesUisprsIDDhcpServersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices uisprs Id dhcp servers params
func (o *PostDevicesUisprsIDDhcpServersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices uisprs Id dhcp servers params
func (o *PostDevicesUisprsIDDhcpServersParams) WithHTTPClient(client *http.Client) *PostDevicesUisprsIDDhcpServersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices uisprs Id dhcp servers params
func (o *PostDevicesUisprsIDDhcpServersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post devices uisprs Id dhcp servers params
func (o *PostDevicesUisprsIDDhcpServersParams) WithBody(body *models.CreateDHCPServer1) *PostDevicesUisprsIDDhcpServersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post devices uisprs Id dhcp servers params
func (o *PostDevicesUisprsIDDhcpServersParams) SetBody(body *models.CreateDHCPServer1) {
	o.Body = body
}

// WithID adds the id to the post devices uisprs Id dhcp servers params
func (o *PostDevicesUisprsIDDhcpServersParams) WithID(id string) *PostDevicesUisprsIDDhcpServersParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post devices uisprs Id dhcp servers params
func (o *PostDevicesUisprsIDDhcpServersParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesUisprsIDDhcpServersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}