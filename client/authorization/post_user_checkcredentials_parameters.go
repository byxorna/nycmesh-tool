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

// NewPostUserCheckcredentialsParams creates a new PostUserCheckcredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostUserCheckcredentialsParams() *PostUserCheckcredentialsParams {
	return &PostUserCheckcredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostUserCheckcredentialsParamsWithTimeout creates a new PostUserCheckcredentialsParams object
// with the ability to set a timeout on a request.
func NewPostUserCheckcredentialsParamsWithTimeout(timeout time.Duration) *PostUserCheckcredentialsParams {
	return &PostUserCheckcredentialsParams{
		timeout: timeout,
	}
}

// NewPostUserCheckcredentialsParamsWithContext creates a new PostUserCheckcredentialsParams object
// with the ability to set a context for a request.
func NewPostUserCheckcredentialsParamsWithContext(ctx context.Context) *PostUserCheckcredentialsParams {
	return &PostUserCheckcredentialsParams{
		Context: ctx,
	}
}

// NewPostUserCheckcredentialsParamsWithHTTPClient creates a new PostUserCheckcredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostUserCheckcredentialsParamsWithHTTPClient(client *http.Client) *PostUserCheckcredentialsParams {
	return &PostUserCheckcredentialsParams{
		HTTPClient: client,
	}
}

/* PostUserCheckcredentialsParams contains all the parameters to send to the API endpoint
   for the post user checkcredentials operation.

   Typically these are written to a http.Request.
*/
type PostUserCheckcredentialsParams struct {

	// Body.
	Body *models.UserCredentials

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post user checkcredentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUserCheckcredentialsParams) WithDefaults() *PostUserCheckcredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post user checkcredentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUserCheckcredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post user checkcredentials params
func (o *PostUserCheckcredentialsParams) WithTimeout(timeout time.Duration) *PostUserCheckcredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post user checkcredentials params
func (o *PostUserCheckcredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post user checkcredentials params
func (o *PostUserCheckcredentialsParams) WithContext(ctx context.Context) *PostUserCheckcredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post user checkcredentials params
func (o *PostUserCheckcredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post user checkcredentials params
func (o *PostUserCheckcredentialsParams) WithHTTPClient(client *http.Client) *PostUserCheckcredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post user checkcredentials params
func (o *PostUserCheckcredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post user checkcredentials params
func (o *PostUserCheckcredentialsParams) WithBody(body *models.UserCredentials) *PostUserCheckcredentialsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post user checkcredentials params
func (o *PostUserCheckcredentialsParams) SetBody(body *models.UserCredentials) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostUserCheckcredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
