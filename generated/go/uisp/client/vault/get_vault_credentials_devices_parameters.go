// Code generated by go-swagger; DO NOT EDIT.

package vault

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

// NewGetVaultCredentialsDevicesParams creates a new GetVaultCredentialsDevicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVaultCredentialsDevicesParams() *GetVaultCredentialsDevicesParams {
	return &GetVaultCredentialsDevicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVaultCredentialsDevicesParamsWithTimeout creates a new GetVaultCredentialsDevicesParams object
// with the ability to set a timeout on a request.
func NewGetVaultCredentialsDevicesParamsWithTimeout(timeout time.Duration) *GetVaultCredentialsDevicesParams {
	return &GetVaultCredentialsDevicesParams{
		timeout: timeout,
	}
}

// NewGetVaultCredentialsDevicesParamsWithContext creates a new GetVaultCredentialsDevicesParams object
// with the ability to set a context for a request.
func NewGetVaultCredentialsDevicesParamsWithContext(ctx context.Context) *GetVaultCredentialsDevicesParams {
	return &GetVaultCredentialsDevicesParams{
		Context: ctx,
	}
}

// NewGetVaultCredentialsDevicesParamsWithHTTPClient creates a new GetVaultCredentialsDevicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVaultCredentialsDevicesParamsWithHTTPClient(client *http.Client) *GetVaultCredentialsDevicesParams {
	return &GetVaultCredentialsDevicesParams{
		HTTPClient: client,
	}
}

/* GetVaultCredentialsDevicesParams contains all the parameters to send to the API endpoint
   for the get vault credentials devices operation.

   Typically these are written to a http.Request.
*/
type GetVaultCredentialsDevicesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vault credentials devices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVaultCredentialsDevicesParams) WithDefaults() *GetVaultCredentialsDevicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vault credentials devices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVaultCredentialsDevicesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vault credentials devices params
func (o *GetVaultCredentialsDevicesParams) WithTimeout(timeout time.Duration) *GetVaultCredentialsDevicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vault credentials devices params
func (o *GetVaultCredentialsDevicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vault credentials devices params
func (o *GetVaultCredentialsDevicesParams) WithContext(ctx context.Context) *GetVaultCredentialsDevicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vault credentials devices params
func (o *GetVaultCredentialsDevicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vault credentials devices params
func (o *GetVaultCredentialsDevicesParams) WithHTTPClient(client *http.Client) *GetVaultCredentialsDevicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vault credentials devices params
func (o *GetVaultCredentialsDevicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetVaultCredentialsDevicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}