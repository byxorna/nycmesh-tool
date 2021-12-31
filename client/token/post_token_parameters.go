// Code generated by go-swagger; DO NOT EDIT.

package token

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

// NewPostTokenParams creates a new PostTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostTokenParams() *PostTokenParams {
	return &PostTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostTokenParamsWithTimeout creates a new PostTokenParams object
// with the ability to set a timeout on a request.
func NewPostTokenParamsWithTimeout(timeout time.Duration) *PostTokenParams {
	return &PostTokenParams{
		timeout: timeout,
	}
}

// NewPostTokenParamsWithContext creates a new PostTokenParams object
// with the ability to set a context for a request.
func NewPostTokenParamsWithContext(ctx context.Context) *PostTokenParams {
	return &PostTokenParams{
		Context: ctx,
	}
}

// NewPostTokenParamsWithHTTPClient creates a new PostTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostTokenParamsWithHTTPClient(client *http.Client) *PostTokenParams {
	return &PostTokenParams{
		HTTPClient: client,
	}
}

/* PostTokenParams contains all the parameters to send to the API endpoint
   for the post token operation.

   Typically these are written to a http.Request.
*/
type PostTokenParams struct {

	// Body.
	Body *models.CreateToken

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostTokenParams) WithDefaults() *PostTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post token params
func (o *PostTokenParams) WithTimeout(timeout time.Duration) *PostTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post token params
func (o *PostTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post token params
func (o *PostTokenParams) WithContext(ctx context.Context) *PostTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post token params
func (o *PostTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post token params
func (o *PostTokenParams) WithHTTPClient(client *http.Client) *PostTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post token params
func (o *PostTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post token params
func (o *PostTokenParams) WithBody(body *models.CreateToken) *PostTokenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post token params
func (o *PostTokenParams) SetBody(body *models.CreateToken) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
