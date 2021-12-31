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

// NewPostUserLoginInviteUbiquitiParams creates a new PostUserLoginInviteUbiquitiParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostUserLoginInviteUbiquitiParams() *PostUserLoginInviteUbiquitiParams {
	return &PostUserLoginInviteUbiquitiParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostUserLoginInviteUbiquitiParamsWithTimeout creates a new PostUserLoginInviteUbiquitiParams object
// with the ability to set a timeout on a request.
func NewPostUserLoginInviteUbiquitiParamsWithTimeout(timeout time.Duration) *PostUserLoginInviteUbiquitiParams {
	return &PostUserLoginInviteUbiquitiParams{
		timeout: timeout,
	}
}

// NewPostUserLoginInviteUbiquitiParamsWithContext creates a new PostUserLoginInviteUbiquitiParams object
// with the ability to set a context for a request.
func NewPostUserLoginInviteUbiquitiParamsWithContext(ctx context.Context) *PostUserLoginInviteUbiquitiParams {
	return &PostUserLoginInviteUbiquitiParams{
		Context: ctx,
	}
}

// NewPostUserLoginInviteUbiquitiParamsWithHTTPClient creates a new PostUserLoginInviteUbiquitiParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostUserLoginInviteUbiquitiParamsWithHTTPClient(client *http.Client) *PostUserLoginInviteUbiquitiParams {
	return &PostUserLoginInviteUbiquitiParams{
		HTTPClient: client,
	}
}

/* PostUserLoginInviteUbiquitiParams contains all the parameters to send to the API endpoint
   for the post user login invite ubiquiti operation.

   Typically these are written to a http.Request.
*/
type PostUserLoginInviteUbiquitiParams struct {

	// Body.
	Body *models.SsoInvitationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post user login invite ubiquiti params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUserLoginInviteUbiquitiParams) WithDefaults() *PostUserLoginInviteUbiquitiParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post user login invite ubiquiti params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUserLoginInviteUbiquitiParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post user login invite ubiquiti params
func (o *PostUserLoginInviteUbiquitiParams) WithTimeout(timeout time.Duration) *PostUserLoginInviteUbiquitiParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post user login invite ubiquiti params
func (o *PostUserLoginInviteUbiquitiParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post user login invite ubiquiti params
func (o *PostUserLoginInviteUbiquitiParams) WithContext(ctx context.Context) *PostUserLoginInviteUbiquitiParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post user login invite ubiquiti params
func (o *PostUserLoginInviteUbiquitiParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post user login invite ubiquiti params
func (o *PostUserLoginInviteUbiquitiParams) WithHTTPClient(client *http.Client) *PostUserLoginInviteUbiquitiParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post user login invite ubiquiti params
func (o *PostUserLoginInviteUbiquitiParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post user login invite ubiquiti params
func (o *PostUserLoginInviteUbiquitiParams) WithBody(body *models.SsoInvitationRequest) *PostUserLoginInviteUbiquitiParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post user login invite ubiquiti params
func (o *PostUserLoginInviteUbiquitiParams) SetBody(body *models.SsoInvitationRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostUserLoginInviteUbiquitiParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
