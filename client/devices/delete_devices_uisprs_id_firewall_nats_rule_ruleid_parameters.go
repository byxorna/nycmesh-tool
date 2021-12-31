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

// NewDeleteDevicesUisprsIDFirewallNatsRuleRuleidParams creates a new DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDevicesUisprsIDFirewallNatsRuleRuleidParams() *DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams {
	return &DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDevicesUisprsIDFirewallNatsRuleRuleidParamsWithTimeout creates a new DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams object
// with the ability to set a timeout on a request.
func NewDeleteDevicesUisprsIDFirewallNatsRuleRuleidParamsWithTimeout(timeout time.Duration) *DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams {
	return &DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams{
		timeout: timeout,
	}
}

// NewDeleteDevicesUisprsIDFirewallNatsRuleRuleidParamsWithContext creates a new DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams object
// with the ability to set a context for a request.
func NewDeleteDevicesUisprsIDFirewallNatsRuleRuleidParamsWithContext(ctx context.Context) *DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams {
	return &DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams{
		Context: ctx,
	}
}

// NewDeleteDevicesUisprsIDFirewallNatsRuleRuleidParamsWithHTTPClient creates a new DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDevicesUisprsIDFirewallNatsRuleRuleidParamsWithHTTPClient(client *http.Client) *DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams {
	return &DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams{
		HTTPClient: client,
	}
}

/* DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams contains all the parameters to send to the API endpoint
   for the delete devices uisprs Id firewall nats rule ruleid operation.

   Typically these are written to a http.Request.
*/
type DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams struct {

	// ID.
	ID string

	// RuleID.
	RuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete devices uisprs Id firewall nats rule ruleid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams) WithDefaults() *DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete devices uisprs Id firewall nats rule ruleid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete devices uisprs Id firewall nats rule ruleid params
func (o *DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams) WithTimeout(timeout time.Duration) *DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete devices uisprs Id firewall nats rule ruleid params
func (o *DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete devices uisprs Id firewall nats rule ruleid params
func (o *DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams) WithContext(ctx context.Context) *DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete devices uisprs Id firewall nats rule ruleid params
func (o *DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete devices uisprs Id firewall nats rule ruleid params
func (o *DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams) WithHTTPClient(client *http.Client) *DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete devices uisprs Id firewall nats rule ruleid params
func (o *DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete devices uisprs Id firewall nats rule ruleid params
func (o *DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams) WithID(id string) *DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete devices uisprs Id firewall nats rule ruleid params
func (o *DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams) SetID(id string) {
	o.ID = id
}

// WithRuleID adds the ruleID to the delete devices uisprs Id firewall nats rule ruleid params
func (o *DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams) WithRuleID(ruleID string) *DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the delete devices uisprs Id firewall nats rule ruleid params
func (o *DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDevicesUisprsIDFirewallNatsRuleRuleidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param ruleId
	if err := r.SetPathParam("ruleId", o.RuleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
