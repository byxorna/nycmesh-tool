// Code generated by go-swagger; DO NOT EDIT.

package discovery

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

// NewPostDiscoveryConnectUbntParams creates a new PostDiscoveryConnectUbntParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDiscoveryConnectUbntParams() *PostDiscoveryConnectUbntParams {
	return &PostDiscoveryConnectUbntParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDiscoveryConnectUbntParamsWithTimeout creates a new PostDiscoveryConnectUbntParams object
// with the ability to set a timeout on a request.
func NewPostDiscoveryConnectUbntParamsWithTimeout(timeout time.Duration) *PostDiscoveryConnectUbntParams {
	return &PostDiscoveryConnectUbntParams{
		timeout: timeout,
	}
}

// NewPostDiscoveryConnectUbntParamsWithContext creates a new PostDiscoveryConnectUbntParams object
// with the ability to set a context for a request.
func NewPostDiscoveryConnectUbntParamsWithContext(ctx context.Context) *PostDiscoveryConnectUbntParams {
	return &PostDiscoveryConnectUbntParams{
		Context: ctx,
	}
}

// NewPostDiscoveryConnectUbntParamsWithHTTPClient creates a new PostDiscoveryConnectUbntParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDiscoveryConnectUbntParamsWithHTTPClient(client *http.Client) *PostDiscoveryConnectUbntParams {
	return &PostDiscoveryConnectUbntParams{
		HTTPClient: client,
	}
}

/* PostDiscoveryConnectUbntParams contains all the parameters to send to the API endpoint
   for the post discovery connect ubnt operation.

   Typically these are written to a http.Request.
*/
type PostDiscoveryConnectUbntParams struct {

	// Body.
	Body *models.Model84

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post discovery connect ubnt params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDiscoveryConnectUbntParams) WithDefaults() *PostDiscoveryConnectUbntParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post discovery connect ubnt params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDiscoveryConnectUbntParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post discovery connect ubnt params
func (o *PostDiscoveryConnectUbntParams) WithTimeout(timeout time.Duration) *PostDiscoveryConnectUbntParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post discovery connect ubnt params
func (o *PostDiscoveryConnectUbntParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post discovery connect ubnt params
func (o *PostDiscoveryConnectUbntParams) WithContext(ctx context.Context) *PostDiscoveryConnectUbntParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post discovery connect ubnt params
func (o *PostDiscoveryConnectUbntParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post discovery connect ubnt params
func (o *PostDiscoveryConnectUbntParams) WithHTTPClient(client *http.Client) *PostDiscoveryConnectUbntParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post discovery connect ubnt params
func (o *PostDiscoveryConnectUbntParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post discovery connect ubnt params
func (o *PostDiscoveryConnectUbntParams) WithBody(body *models.Model84) *PostDiscoveryConnectUbntParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post discovery connect ubnt params
func (o *PostDiscoveryConnectUbntParams) SetBody(body *models.Model84) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostDiscoveryConnectUbntParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
