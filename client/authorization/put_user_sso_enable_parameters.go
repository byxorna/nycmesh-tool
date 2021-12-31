// Code generated by go-swagger; DO NOT EDIT.

package authorization

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

// NewPutUserSsoEnableParams creates a new PutUserSsoEnableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutUserSsoEnableParams() *PutUserSsoEnableParams {
	return &PutUserSsoEnableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutUserSsoEnableParamsWithTimeout creates a new PutUserSsoEnableParams object
// with the ability to set a timeout on a request.
func NewPutUserSsoEnableParamsWithTimeout(timeout time.Duration) *PutUserSsoEnableParams {
	return &PutUserSsoEnableParams{
		timeout: timeout,
	}
}

// NewPutUserSsoEnableParamsWithContext creates a new PutUserSsoEnableParams object
// with the ability to set a context for a request.
func NewPutUserSsoEnableParamsWithContext(ctx context.Context) *PutUserSsoEnableParams {
	return &PutUserSsoEnableParams{
		Context: ctx,
	}
}

// NewPutUserSsoEnableParamsWithHTTPClient creates a new PutUserSsoEnableParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutUserSsoEnableParamsWithHTTPClient(client *http.Client) *PutUserSsoEnableParams {
	return &PutUserSsoEnableParams{
		HTTPClient: client,
	}
}

/* PutUserSsoEnableParams contains all the parameters to send to the API endpoint
   for the put user sso enable operation.

   Typically these are written to a http.Request.
*/
type PutUserSsoEnableParams struct {

	// Body.
	Body *models.EnableSso

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put user sso enable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutUserSsoEnableParams) WithDefaults() *PutUserSsoEnableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put user sso enable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutUserSsoEnableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put user sso enable params
func (o *PutUserSsoEnableParams) WithTimeout(timeout time.Duration) *PutUserSsoEnableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put user sso enable params
func (o *PutUserSsoEnableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put user sso enable params
func (o *PutUserSsoEnableParams) WithContext(ctx context.Context) *PutUserSsoEnableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put user sso enable params
func (o *PutUserSsoEnableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put user sso enable params
func (o *PutUserSsoEnableParams) WithHTTPClient(client *http.Client) *PutUserSsoEnableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put user sso enable params
func (o *PutUserSsoEnableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put user sso enable params
func (o *PutUserSsoEnableParams) WithBody(body *models.EnableSso) *PutUserSsoEnableParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put user sso enable params
func (o *PutUserSsoEnableParams) SetBody(body *models.EnableSso) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutUserSsoEnableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
