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
)

// NewPostUserLogoutParams creates a new PostUserLogoutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostUserLogoutParams() *PostUserLogoutParams {
	return &PostUserLogoutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostUserLogoutParamsWithTimeout creates a new PostUserLogoutParams object
// with the ability to set a timeout on a request.
func NewPostUserLogoutParamsWithTimeout(timeout time.Duration) *PostUserLogoutParams {
	return &PostUserLogoutParams{
		timeout: timeout,
	}
}

// NewPostUserLogoutParamsWithContext creates a new PostUserLogoutParams object
// with the ability to set a context for a request.
func NewPostUserLogoutParamsWithContext(ctx context.Context) *PostUserLogoutParams {
	return &PostUserLogoutParams{
		Context: ctx,
	}
}

// NewPostUserLogoutParamsWithHTTPClient creates a new PostUserLogoutParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostUserLogoutParamsWithHTTPClient(client *http.Client) *PostUserLogoutParams {
	return &PostUserLogoutParams{
		HTTPClient: client,
	}
}

/* PostUserLogoutParams contains all the parameters to send to the API endpoint
   for the post user logout operation.

   Typically these are written to a http.Request.
*/
type PostUserLogoutParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post user logout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUserLogoutParams) WithDefaults() *PostUserLogoutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post user logout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUserLogoutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post user logout params
func (o *PostUserLogoutParams) WithTimeout(timeout time.Duration) *PostUserLogoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post user logout params
func (o *PostUserLogoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post user logout params
func (o *PostUserLogoutParams) WithContext(ctx context.Context) *PostUserLogoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post user logout params
func (o *PostUserLogoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post user logout params
func (o *PostUserLogoutParams) WithHTTPClient(client *http.Client) *PostUserLogoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post user logout params
func (o *PostUserLogoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostUserLogoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
