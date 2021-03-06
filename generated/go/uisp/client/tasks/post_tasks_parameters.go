// Code generated by go-swagger; DO NOT EDIT.

package tasks

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

// NewPostTasksParams creates a new PostTasksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostTasksParams() *PostTasksParams {
	return &PostTasksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostTasksParamsWithTimeout creates a new PostTasksParams object
// with the ability to set a timeout on a request.
func NewPostTasksParamsWithTimeout(timeout time.Duration) *PostTasksParams {
	return &PostTasksParams{
		timeout: timeout,
	}
}

// NewPostTasksParamsWithContext creates a new PostTasksParams object
// with the ability to set a context for a request.
func NewPostTasksParamsWithContext(ctx context.Context) *PostTasksParams {
	return &PostTasksParams{
		Context: ctx,
	}
}

// NewPostTasksParamsWithHTTPClient creates a new PostTasksParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostTasksParamsWithHTTPClient(client *http.Client) *PostTasksParams {
	return &PostTasksParams{
		HTTPClient: client,
	}
}

/* PostTasksParams contains all the parameters to send to the API endpoint
   for the post tasks operation.

   Typically these are written to a http.Request.
*/
type PostTasksParams struct {

	// Body.
	Body *models.Model70

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostTasksParams) WithDefaults() *PostTasksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostTasksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post tasks params
func (o *PostTasksParams) WithTimeout(timeout time.Duration) *PostTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post tasks params
func (o *PostTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post tasks params
func (o *PostTasksParams) WithContext(ctx context.Context) *PostTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post tasks params
func (o *PostTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post tasks params
func (o *PostTasksParams) WithHTTPClient(client *http.Client) *PostTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post tasks params
func (o *PostTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post tasks params
func (o *PostTasksParams) WithBody(body *models.Model70) *PostTasksParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post tasks params
func (o *PostTasksParams) SetBody(body *models.Model70) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
