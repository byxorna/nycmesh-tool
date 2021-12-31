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

// NewPostUserPasswordRequestresetParams creates a new PostUserPasswordRequestresetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostUserPasswordRequestresetParams() *PostUserPasswordRequestresetParams {
	return &PostUserPasswordRequestresetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostUserPasswordRequestresetParamsWithTimeout creates a new PostUserPasswordRequestresetParams object
// with the ability to set a timeout on a request.
func NewPostUserPasswordRequestresetParamsWithTimeout(timeout time.Duration) *PostUserPasswordRequestresetParams {
	return &PostUserPasswordRequestresetParams{
		timeout: timeout,
	}
}

// NewPostUserPasswordRequestresetParamsWithContext creates a new PostUserPasswordRequestresetParams object
// with the ability to set a context for a request.
func NewPostUserPasswordRequestresetParamsWithContext(ctx context.Context) *PostUserPasswordRequestresetParams {
	return &PostUserPasswordRequestresetParams{
		Context: ctx,
	}
}

// NewPostUserPasswordRequestresetParamsWithHTTPClient creates a new PostUserPasswordRequestresetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostUserPasswordRequestresetParamsWithHTTPClient(client *http.Client) *PostUserPasswordRequestresetParams {
	return &PostUserPasswordRequestresetParams{
		HTTPClient: client,
	}
}

/* PostUserPasswordRequestresetParams contains all the parameters to send to the API endpoint
   for the post user password requestreset operation.

   Typically these are written to a http.Request.
*/
type PostUserPasswordRequestresetParams struct {

	// Body.
	Body *models.PasswordResetRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post user password requestreset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUserPasswordRequestresetParams) WithDefaults() *PostUserPasswordRequestresetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post user password requestreset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUserPasswordRequestresetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post user password requestreset params
func (o *PostUserPasswordRequestresetParams) WithTimeout(timeout time.Duration) *PostUserPasswordRequestresetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post user password requestreset params
func (o *PostUserPasswordRequestresetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post user password requestreset params
func (o *PostUserPasswordRequestresetParams) WithContext(ctx context.Context) *PostUserPasswordRequestresetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post user password requestreset params
func (o *PostUserPasswordRequestresetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post user password requestreset params
func (o *PostUserPasswordRequestresetParams) WithHTTPClient(client *http.Client) *PostUserPasswordRequestresetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post user password requestreset params
func (o *PostUserPasswordRequestresetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post user password requestreset params
func (o *PostUserPasswordRequestresetParams) WithBody(body *models.PasswordResetRequest) *PostUserPasswordRequestresetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post user password requestreset params
func (o *PostUserPasswordRequestresetParams) SetBody(body *models.PasswordResetRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostUserPasswordRequestresetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
