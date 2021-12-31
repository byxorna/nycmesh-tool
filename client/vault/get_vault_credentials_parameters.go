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

// NewGetVaultCredentialsParams creates a new GetVaultCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVaultCredentialsParams() *GetVaultCredentialsParams {
	return &GetVaultCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVaultCredentialsParamsWithTimeout creates a new GetVaultCredentialsParams object
// with the ability to set a timeout on a request.
func NewGetVaultCredentialsParamsWithTimeout(timeout time.Duration) *GetVaultCredentialsParams {
	return &GetVaultCredentialsParams{
		timeout: timeout,
	}
}

// NewGetVaultCredentialsParamsWithContext creates a new GetVaultCredentialsParams object
// with the ability to set a context for a request.
func NewGetVaultCredentialsParamsWithContext(ctx context.Context) *GetVaultCredentialsParams {
	return &GetVaultCredentialsParams{
		Context: ctx,
	}
}

// NewGetVaultCredentialsParamsWithHTTPClient creates a new GetVaultCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVaultCredentialsParamsWithHTTPClient(client *http.Client) *GetVaultCredentialsParams {
	return &GetVaultCredentialsParams{
		HTTPClient: client,
	}
}

/* GetVaultCredentialsParams contains all the parameters to send to the API endpoint
   for the get vault credentials operation.

   Typically these are written to a http.Request.
*/
type GetVaultCredentialsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vault credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVaultCredentialsParams) WithDefaults() *GetVaultCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vault credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVaultCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vault credentials params
func (o *GetVaultCredentialsParams) WithTimeout(timeout time.Duration) *GetVaultCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vault credentials params
func (o *GetVaultCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vault credentials params
func (o *GetVaultCredentialsParams) WithContext(ctx context.Context) *GetVaultCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vault credentials params
func (o *GetVaultCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vault credentials params
func (o *GetVaultCredentialsParams) WithHTTPClient(client *http.Client) *GetVaultCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vault credentials params
func (o *GetVaultCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetVaultCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}