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

	"github.com/byxorna/nycmesh-tool/models"
)

// NewPostVaultCredentialsRegenerateParams creates a new PostVaultCredentialsRegenerateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostVaultCredentialsRegenerateParams() *PostVaultCredentialsRegenerateParams {
	return &PostVaultCredentialsRegenerateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostVaultCredentialsRegenerateParamsWithTimeout creates a new PostVaultCredentialsRegenerateParams object
// with the ability to set a timeout on a request.
func NewPostVaultCredentialsRegenerateParamsWithTimeout(timeout time.Duration) *PostVaultCredentialsRegenerateParams {
	return &PostVaultCredentialsRegenerateParams{
		timeout: timeout,
	}
}

// NewPostVaultCredentialsRegenerateParamsWithContext creates a new PostVaultCredentialsRegenerateParams object
// with the ability to set a context for a request.
func NewPostVaultCredentialsRegenerateParamsWithContext(ctx context.Context) *PostVaultCredentialsRegenerateParams {
	return &PostVaultCredentialsRegenerateParams{
		Context: ctx,
	}
}

// NewPostVaultCredentialsRegenerateParamsWithHTTPClient creates a new PostVaultCredentialsRegenerateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostVaultCredentialsRegenerateParamsWithHTTPClient(client *http.Client) *PostVaultCredentialsRegenerateParams {
	return &PostVaultCredentialsRegenerateParams{
		HTTPClient: client,
	}
}

/* PostVaultCredentialsRegenerateParams contains all the parameters to send to the API endpoint
   for the post vault credentials regenerate operation.

   Typically these are written to a http.Request.
*/
type PostVaultCredentialsRegenerateParams struct {

	// Body.
	Body *models.Model73

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post vault credentials regenerate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostVaultCredentialsRegenerateParams) WithDefaults() *PostVaultCredentialsRegenerateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post vault credentials regenerate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostVaultCredentialsRegenerateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post vault credentials regenerate params
func (o *PostVaultCredentialsRegenerateParams) WithTimeout(timeout time.Duration) *PostVaultCredentialsRegenerateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post vault credentials regenerate params
func (o *PostVaultCredentialsRegenerateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post vault credentials regenerate params
func (o *PostVaultCredentialsRegenerateParams) WithContext(ctx context.Context) *PostVaultCredentialsRegenerateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post vault credentials regenerate params
func (o *PostVaultCredentialsRegenerateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post vault credentials regenerate params
func (o *PostVaultCredentialsRegenerateParams) WithHTTPClient(client *http.Client) *PostVaultCredentialsRegenerateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post vault credentials regenerate params
func (o *PostVaultCredentialsRegenerateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post vault credentials regenerate params
func (o *PostVaultCredentialsRegenerateParams) WithBody(body *models.Model73) *PostVaultCredentialsRegenerateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post vault credentials regenerate params
func (o *PostVaultCredentialsRegenerateParams) SetBody(body *models.Model73) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostVaultCredentialsRegenerateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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