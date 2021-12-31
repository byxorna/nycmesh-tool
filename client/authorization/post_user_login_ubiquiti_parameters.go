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

// NewPostUserLoginUbiquitiParams creates a new PostUserLoginUbiquitiParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostUserLoginUbiquitiParams() *PostUserLoginUbiquitiParams {
	return &PostUserLoginUbiquitiParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostUserLoginUbiquitiParamsWithTimeout creates a new PostUserLoginUbiquitiParams object
// with the ability to set a timeout on a request.
func NewPostUserLoginUbiquitiParamsWithTimeout(timeout time.Duration) *PostUserLoginUbiquitiParams {
	return &PostUserLoginUbiquitiParams{
		timeout: timeout,
	}
}

// NewPostUserLoginUbiquitiParamsWithContext creates a new PostUserLoginUbiquitiParams object
// with the ability to set a context for a request.
func NewPostUserLoginUbiquitiParamsWithContext(ctx context.Context) *PostUserLoginUbiquitiParams {
	return &PostUserLoginUbiquitiParams{
		Context: ctx,
	}
}

// NewPostUserLoginUbiquitiParamsWithHTTPClient creates a new PostUserLoginUbiquitiParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostUserLoginUbiquitiParamsWithHTTPClient(client *http.Client) *PostUserLoginUbiquitiParams {
	return &PostUserLoginUbiquitiParams{
		HTTPClient: client,
	}
}

/* PostUserLoginUbiquitiParams contains all the parameters to send to the API endpoint
   for the post user login ubiquiti operation.

   Typically these are written to a http.Request.
*/
type PostUserLoginUbiquitiParams struct {

	// Body.
	Body *models.SsoRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post user login ubiquiti params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUserLoginUbiquitiParams) WithDefaults() *PostUserLoginUbiquitiParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post user login ubiquiti params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUserLoginUbiquitiParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post user login ubiquiti params
func (o *PostUserLoginUbiquitiParams) WithTimeout(timeout time.Duration) *PostUserLoginUbiquitiParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post user login ubiquiti params
func (o *PostUserLoginUbiquitiParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post user login ubiquiti params
func (o *PostUserLoginUbiquitiParams) WithContext(ctx context.Context) *PostUserLoginUbiquitiParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post user login ubiquiti params
func (o *PostUserLoginUbiquitiParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post user login ubiquiti params
func (o *PostUserLoginUbiquitiParams) WithHTTPClient(client *http.Client) *PostUserLoginUbiquitiParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post user login ubiquiti params
func (o *PostUserLoginUbiquitiParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post user login ubiquiti params
func (o *PostUserLoginUbiquitiParams) WithBody(body *models.SsoRequest) *PostUserLoginUbiquitiParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post user login ubiquiti params
func (o *PostUserLoginUbiquitiParams) SetBody(body *models.SsoRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostUserLoginUbiquitiParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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