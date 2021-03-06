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

// NewGetDevicesUisprsIDFirewallNatsParams creates a new GetDevicesUisprsIDFirewallNatsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicesUisprsIDFirewallNatsParams() *GetDevicesUisprsIDFirewallNatsParams {
	return &GetDevicesUisprsIDFirewallNatsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesUisprsIDFirewallNatsParamsWithTimeout creates a new GetDevicesUisprsIDFirewallNatsParams object
// with the ability to set a timeout on a request.
func NewGetDevicesUisprsIDFirewallNatsParamsWithTimeout(timeout time.Duration) *GetDevicesUisprsIDFirewallNatsParams {
	return &GetDevicesUisprsIDFirewallNatsParams{
		timeout: timeout,
	}
}

// NewGetDevicesUisprsIDFirewallNatsParamsWithContext creates a new GetDevicesUisprsIDFirewallNatsParams object
// with the ability to set a context for a request.
func NewGetDevicesUisprsIDFirewallNatsParamsWithContext(ctx context.Context) *GetDevicesUisprsIDFirewallNatsParams {
	return &GetDevicesUisprsIDFirewallNatsParams{
		Context: ctx,
	}
}

// NewGetDevicesUisprsIDFirewallNatsParamsWithHTTPClient creates a new GetDevicesUisprsIDFirewallNatsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicesUisprsIDFirewallNatsParamsWithHTTPClient(client *http.Client) *GetDevicesUisprsIDFirewallNatsParams {
	return &GetDevicesUisprsIDFirewallNatsParams{
		HTTPClient: client,
	}
}

/* GetDevicesUisprsIDFirewallNatsParams contains all the parameters to send to the API endpoint
   for the get devices uisprs Id firewall nats operation.

   Typically these are written to a http.Request.
*/
type GetDevicesUisprsIDFirewallNatsParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get devices uisprs Id firewall nats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesUisprsIDFirewallNatsParams) WithDefaults() *GetDevicesUisprsIDFirewallNatsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get devices uisprs Id firewall nats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesUisprsIDFirewallNatsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get devices uisprs Id firewall nats params
func (o *GetDevicesUisprsIDFirewallNatsParams) WithTimeout(timeout time.Duration) *GetDevicesUisprsIDFirewallNatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices uisprs Id firewall nats params
func (o *GetDevicesUisprsIDFirewallNatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices uisprs Id firewall nats params
func (o *GetDevicesUisprsIDFirewallNatsParams) WithContext(ctx context.Context) *GetDevicesUisprsIDFirewallNatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices uisprs Id firewall nats params
func (o *GetDevicesUisprsIDFirewallNatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices uisprs Id firewall nats params
func (o *GetDevicesUisprsIDFirewallNatsParams) WithHTTPClient(client *http.Client) *GetDevicesUisprsIDFirewallNatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices uisprs Id firewall nats params
func (o *GetDevicesUisprsIDFirewallNatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get devices uisprs Id firewall nats params
func (o *GetDevicesUisprsIDFirewallNatsParams) WithID(id string) *GetDevicesUisprsIDFirewallNatsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get devices uisprs Id firewall nats params
func (o *GetDevicesUisprsIDFirewallNatsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesUisprsIDFirewallNatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
