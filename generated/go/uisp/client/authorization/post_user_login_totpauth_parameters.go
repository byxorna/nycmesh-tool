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

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// NewPostUserLoginTotpauthParams creates a new PostUserLoginTotpauthParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostUserLoginTotpauthParams() *PostUserLoginTotpauthParams {
	return &PostUserLoginTotpauthParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostUserLoginTotpauthParamsWithTimeout creates a new PostUserLoginTotpauthParams object
// with the ability to set a timeout on a request.
func NewPostUserLoginTotpauthParamsWithTimeout(timeout time.Duration) *PostUserLoginTotpauthParams {
	return &PostUserLoginTotpauthParams{
		timeout: timeout,
	}
}

// NewPostUserLoginTotpauthParamsWithContext creates a new PostUserLoginTotpauthParams object
// with the ability to set a context for a request.
func NewPostUserLoginTotpauthParamsWithContext(ctx context.Context) *PostUserLoginTotpauthParams {
	return &PostUserLoginTotpauthParams{
		Context: ctx,
	}
}

// NewPostUserLoginTotpauthParamsWithHTTPClient creates a new PostUserLoginTotpauthParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostUserLoginTotpauthParamsWithHTTPClient(client *http.Client) *PostUserLoginTotpauthParams {
	return &PostUserLoginTotpauthParams{
		HTTPClient: client,
	}
}

/* PostUserLoginTotpauthParams contains all the parameters to send to the API endpoint
   for the post user login totpauth operation.

   Typically these are written to a http.Request.
*/
type PostUserLoginTotpauthParams struct {

	// Body.
	Body *models.TwoFactorLogin

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post user login totpauth params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUserLoginTotpauthParams) WithDefaults() *PostUserLoginTotpauthParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post user login totpauth params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUserLoginTotpauthParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post user login totpauth params
func (o *PostUserLoginTotpauthParams) WithTimeout(timeout time.Duration) *PostUserLoginTotpauthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post user login totpauth params
func (o *PostUserLoginTotpauthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post user login totpauth params
func (o *PostUserLoginTotpauthParams) WithContext(ctx context.Context) *PostUserLoginTotpauthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post user login totpauth params
func (o *PostUserLoginTotpauthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post user login totpauth params
func (o *PostUserLoginTotpauthParams) WithHTTPClient(client *http.Client) *PostUserLoginTotpauthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post user login totpauth params
func (o *PostUserLoginTotpauthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post user login totpauth params
func (o *PostUserLoginTotpauthParams) WithBody(body *models.TwoFactorLogin) *PostUserLoginTotpauthParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post user login totpauth params
func (o *PostUserLoginTotpauthParams) SetBody(body *models.TwoFactorLogin) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostUserLoginTotpauthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
