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

// NewPostDevicesIDUpgradetolatestParams creates a new PostDevicesIDUpgradetolatestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDevicesIDUpgradetolatestParams() *PostDevicesIDUpgradetolatestParams {
	return &PostDevicesIDUpgradetolatestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesIDUpgradetolatestParamsWithTimeout creates a new PostDevicesIDUpgradetolatestParams object
// with the ability to set a timeout on a request.
func NewPostDevicesIDUpgradetolatestParamsWithTimeout(timeout time.Duration) *PostDevicesIDUpgradetolatestParams {
	return &PostDevicesIDUpgradetolatestParams{
		timeout: timeout,
	}
}

// NewPostDevicesIDUpgradetolatestParamsWithContext creates a new PostDevicesIDUpgradetolatestParams object
// with the ability to set a context for a request.
func NewPostDevicesIDUpgradetolatestParamsWithContext(ctx context.Context) *PostDevicesIDUpgradetolatestParams {
	return &PostDevicesIDUpgradetolatestParams{
		Context: ctx,
	}
}

// NewPostDevicesIDUpgradetolatestParamsWithHTTPClient creates a new PostDevicesIDUpgradetolatestParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDevicesIDUpgradetolatestParamsWithHTTPClient(client *http.Client) *PostDevicesIDUpgradetolatestParams {
	return &PostDevicesIDUpgradetolatestParams{
		HTTPClient: client,
	}
}

/* PostDevicesIDUpgradetolatestParams contains all the parameters to send to the API endpoint
   for the post devices Id upgradetolatest operation.

   Typically these are written to a http.Request.
*/
type PostDevicesIDUpgradetolatestParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post devices Id upgradetolatest params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesIDUpgradetolatestParams) WithDefaults() *PostDevicesIDUpgradetolatestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post devices Id upgradetolatest params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesIDUpgradetolatestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post devices Id upgradetolatest params
func (o *PostDevicesIDUpgradetolatestParams) WithTimeout(timeout time.Duration) *PostDevicesIDUpgradetolatestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices Id upgradetolatest params
func (o *PostDevicesIDUpgradetolatestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices Id upgradetolatest params
func (o *PostDevicesIDUpgradetolatestParams) WithContext(ctx context.Context) *PostDevicesIDUpgradetolatestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices Id upgradetolatest params
func (o *PostDevicesIDUpgradetolatestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices Id upgradetolatest params
func (o *PostDevicesIDUpgradetolatestParams) WithHTTPClient(client *http.Client) *PostDevicesIDUpgradetolatestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices Id upgradetolatest params
func (o *PostDevicesIDUpgradetolatestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post devices Id upgradetolatest params
func (o *PostDevicesIDUpgradetolatestParams) WithID(id string) *PostDevicesIDUpgradetolatestParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post devices Id upgradetolatest params
func (o *PostDevicesIDUpgradetolatestParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesIDUpgradetolatestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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