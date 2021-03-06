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

// NewPostUserPasswordResetParams creates a new PostUserPasswordResetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostUserPasswordResetParams() *PostUserPasswordResetParams {
	return &PostUserPasswordResetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostUserPasswordResetParamsWithTimeout creates a new PostUserPasswordResetParams object
// with the ability to set a timeout on a request.
func NewPostUserPasswordResetParamsWithTimeout(timeout time.Duration) *PostUserPasswordResetParams {
	return &PostUserPasswordResetParams{
		timeout: timeout,
	}
}

// NewPostUserPasswordResetParamsWithContext creates a new PostUserPasswordResetParams object
// with the ability to set a context for a request.
func NewPostUserPasswordResetParamsWithContext(ctx context.Context) *PostUserPasswordResetParams {
	return &PostUserPasswordResetParams{
		Context: ctx,
	}
}

// NewPostUserPasswordResetParamsWithHTTPClient creates a new PostUserPasswordResetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostUserPasswordResetParamsWithHTTPClient(client *http.Client) *PostUserPasswordResetParams {
	return &PostUserPasswordResetParams{
		HTTPClient: client,
	}
}

/* PostUserPasswordResetParams contains all the parameters to send to the API endpoint
   for the post user password reset operation.

   Typically these are written to a http.Request.
*/
type PostUserPasswordResetParams struct {

	// Body.
	Body *models.PasswordResetAction

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post user password reset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUserPasswordResetParams) WithDefaults() *PostUserPasswordResetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post user password reset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUserPasswordResetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post user password reset params
func (o *PostUserPasswordResetParams) WithTimeout(timeout time.Duration) *PostUserPasswordResetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post user password reset params
func (o *PostUserPasswordResetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post user password reset params
func (o *PostUserPasswordResetParams) WithContext(ctx context.Context) *PostUserPasswordResetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post user password reset params
func (o *PostUserPasswordResetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post user password reset params
func (o *PostUserPasswordResetParams) WithHTTPClient(client *http.Client) *PostUserPasswordResetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post user password reset params
func (o *PostUserPasswordResetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post user password reset params
func (o *PostUserPasswordResetParams) WithBody(body *models.PasswordResetAction) *PostUserPasswordResetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post user password reset params
func (o *PostUserPasswordResetParams) SetBody(body *models.PasswordResetAction) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostUserPasswordResetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
