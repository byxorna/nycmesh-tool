// Code generated by go-swagger; DO NOT EDIT.

package web_socket

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

// NewPostWsapiSubscriptionParams creates a new PostWsapiSubscriptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWsapiSubscriptionParams() *PostWsapiSubscriptionParams {
	return &PostWsapiSubscriptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWsapiSubscriptionParamsWithTimeout creates a new PostWsapiSubscriptionParams object
// with the ability to set a timeout on a request.
func NewPostWsapiSubscriptionParamsWithTimeout(timeout time.Duration) *PostWsapiSubscriptionParams {
	return &PostWsapiSubscriptionParams{
		timeout: timeout,
	}
}

// NewPostWsapiSubscriptionParamsWithContext creates a new PostWsapiSubscriptionParams object
// with the ability to set a context for a request.
func NewPostWsapiSubscriptionParamsWithContext(ctx context.Context) *PostWsapiSubscriptionParams {
	return &PostWsapiSubscriptionParams{
		Context: ctx,
	}
}

// NewPostWsapiSubscriptionParamsWithHTTPClient creates a new PostWsapiSubscriptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWsapiSubscriptionParamsWithHTTPClient(client *http.Client) *PostWsapiSubscriptionParams {
	return &PostWsapiSubscriptionParams{
		HTTPClient: client,
	}
}

/* PostWsapiSubscriptionParams contains all the parameters to send to the API endpoint
   for the post wsapi subscription operation.

   Typically these are written to a http.Request.
*/
type PostWsapiSubscriptionParams struct {

	// Body.
	Body *models.Model80

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post wsapi subscription params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWsapiSubscriptionParams) WithDefaults() *PostWsapiSubscriptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post wsapi subscription params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWsapiSubscriptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post wsapi subscription params
func (o *PostWsapiSubscriptionParams) WithTimeout(timeout time.Duration) *PostWsapiSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post wsapi subscription params
func (o *PostWsapiSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post wsapi subscription params
func (o *PostWsapiSubscriptionParams) WithContext(ctx context.Context) *PostWsapiSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post wsapi subscription params
func (o *PostWsapiSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post wsapi subscription params
func (o *PostWsapiSubscriptionParams) WithHTTPClient(client *http.Client) *PostWsapiSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post wsapi subscription params
func (o *PostWsapiSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post wsapi subscription params
func (o *PostWsapiSubscriptionParams) WithBody(body *models.Model80) *PostWsapiSubscriptionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post wsapi subscription params
func (o *PostWsapiSubscriptionParams) SetBody(body *models.Model80) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostWsapiSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
