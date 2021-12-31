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

// NewPutDevicesUisprsIDFirewallFilterChainParams creates a new PutDevicesUisprsIDFirewallFilterChainParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutDevicesUisprsIDFirewallFilterChainParams() *PutDevicesUisprsIDFirewallFilterChainParams {
	return &PutDevicesUisprsIDFirewallFilterChainParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutDevicesUisprsIDFirewallFilterChainParamsWithTimeout creates a new PutDevicesUisprsIDFirewallFilterChainParams object
// with the ability to set a timeout on a request.
func NewPutDevicesUisprsIDFirewallFilterChainParamsWithTimeout(timeout time.Duration) *PutDevicesUisprsIDFirewallFilterChainParams {
	return &PutDevicesUisprsIDFirewallFilterChainParams{
		timeout: timeout,
	}
}

// NewPutDevicesUisprsIDFirewallFilterChainParamsWithContext creates a new PutDevicesUisprsIDFirewallFilterChainParams object
// with the ability to set a context for a request.
func NewPutDevicesUisprsIDFirewallFilterChainParamsWithContext(ctx context.Context) *PutDevicesUisprsIDFirewallFilterChainParams {
	return &PutDevicesUisprsIDFirewallFilterChainParams{
		Context: ctx,
	}
}

// NewPutDevicesUisprsIDFirewallFilterChainParamsWithHTTPClient creates a new PutDevicesUisprsIDFirewallFilterChainParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutDevicesUisprsIDFirewallFilterChainParamsWithHTTPClient(client *http.Client) *PutDevicesUisprsIDFirewallFilterChainParams {
	return &PutDevicesUisprsIDFirewallFilterChainParams{
		HTTPClient: client,
	}
}

/* PutDevicesUisprsIDFirewallFilterChainParams contains all the parameters to send to the API endpoint
   for the put devices uisprs Id firewall filter chain operation.

   Typically these are written to a http.Request.
*/
type PutDevicesUisprsIDFirewallFilterChainParams struct {

	// Body.
	Body *models.Model60

	// ID.
	ID string

	// IsForceApply.
	IsForceApply *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put devices uisprs Id firewall filter chain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDevicesUisprsIDFirewallFilterChainParams) WithDefaults() *PutDevicesUisprsIDFirewallFilterChainParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put devices uisprs Id firewall filter chain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDevicesUisprsIDFirewallFilterChainParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put devices uisprs Id firewall filter chain params
func (o *PutDevicesUisprsIDFirewallFilterChainParams) WithTimeout(timeout time.Duration) *PutDevicesUisprsIDFirewallFilterChainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put devices uisprs Id firewall filter chain params
func (o *PutDevicesUisprsIDFirewallFilterChainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put devices uisprs Id firewall filter chain params
func (o *PutDevicesUisprsIDFirewallFilterChainParams) WithContext(ctx context.Context) *PutDevicesUisprsIDFirewallFilterChainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put devices uisprs Id firewall filter chain params
func (o *PutDevicesUisprsIDFirewallFilterChainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put devices uisprs Id firewall filter chain params
func (o *PutDevicesUisprsIDFirewallFilterChainParams) WithHTTPClient(client *http.Client) *PutDevicesUisprsIDFirewallFilterChainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put devices uisprs Id firewall filter chain params
func (o *PutDevicesUisprsIDFirewallFilterChainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put devices uisprs Id firewall filter chain params
func (o *PutDevicesUisprsIDFirewallFilterChainParams) WithBody(body *models.Model60) *PutDevicesUisprsIDFirewallFilterChainParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put devices uisprs Id firewall filter chain params
func (o *PutDevicesUisprsIDFirewallFilterChainParams) SetBody(body *models.Model60) {
	o.Body = body
}

// WithID adds the id to the put devices uisprs Id firewall filter chain params
func (o *PutDevicesUisprsIDFirewallFilterChainParams) WithID(id string) *PutDevicesUisprsIDFirewallFilterChainParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put devices uisprs Id firewall filter chain params
func (o *PutDevicesUisprsIDFirewallFilterChainParams) SetID(id string) {
	o.ID = id
}

// WithIsForceApply adds the isForceApply to the put devices uisprs Id firewall filter chain params
func (o *PutDevicesUisprsIDFirewallFilterChainParams) WithIsForceApply(isForceApply *string) *PutDevicesUisprsIDFirewallFilterChainParams {
	o.SetIsForceApply(isForceApply)
	return o
}

// SetIsForceApply adds the isForceApply to the put devices uisprs Id firewall filter chain params
func (o *PutDevicesUisprsIDFirewallFilterChainParams) SetIsForceApply(isForceApply *string) {
	o.IsForceApply = isForceApply
}

// WriteToRequest writes these params to a swagger request
func (o *PutDevicesUisprsIDFirewallFilterChainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IsForceApply != nil {

		// query param isForceApply
		var qrIsForceApply string

		if o.IsForceApply != nil {
			qrIsForceApply = *o.IsForceApply
		}
		qIsForceApply := qrIsForceApply
		if qIsForceApply != "" {

			if err := r.SetQueryParam("isForceApply", qIsForceApply); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
