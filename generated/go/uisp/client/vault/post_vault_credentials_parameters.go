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

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// NewPostVaultCredentialsParams creates a new PostVaultCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostVaultCredentialsParams() *PostVaultCredentialsParams {
	return &PostVaultCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostVaultCredentialsParamsWithTimeout creates a new PostVaultCredentialsParams object
// with the ability to set a timeout on a request.
func NewPostVaultCredentialsParamsWithTimeout(timeout time.Duration) *PostVaultCredentialsParams {
	return &PostVaultCredentialsParams{
		timeout: timeout,
	}
}

// NewPostVaultCredentialsParamsWithContext creates a new PostVaultCredentialsParams object
// with the ability to set a context for a request.
func NewPostVaultCredentialsParamsWithContext(ctx context.Context) *PostVaultCredentialsParams {
	return &PostVaultCredentialsParams{
		Context: ctx,
	}
}

// NewPostVaultCredentialsParamsWithHTTPClient creates a new PostVaultCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostVaultCredentialsParamsWithHTTPClient(client *http.Client) *PostVaultCredentialsParams {
	return &PostVaultCredentialsParams{
		HTTPClient: client,
	}
}

/* PostVaultCredentialsParams contains all the parameters to send to the API endpoint
   for the post vault credentials operation.

   Typically these are written to a http.Request.
*/
type PostVaultCredentialsParams struct {

	// Body.
	Body *models.VaultCredentials

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post vault credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostVaultCredentialsParams) WithDefaults() *PostVaultCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post vault credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostVaultCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post vault credentials params
func (o *PostVaultCredentialsParams) WithTimeout(timeout time.Duration) *PostVaultCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post vault credentials params
func (o *PostVaultCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post vault credentials params
func (o *PostVaultCredentialsParams) WithContext(ctx context.Context) *PostVaultCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post vault credentials params
func (o *PostVaultCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post vault credentials params
func (o *PostVaultCredentialsParams) WithHTTPClient(client *http.Client) *PostVaultCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post vault credentials params
func (o *PostVaultCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post vault credentials params
func (o *PostVaultCredentialsParams) WithBody(body *models.VaultCredentials) *PostVaultCredentialsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post vault credentials params
func (o *PostVaultCredentialsParams) SetBody(body *models.VaultCredentials) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostVaultCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}