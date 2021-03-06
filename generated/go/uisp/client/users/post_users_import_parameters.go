// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewPostUsersImportParams creates a new PostUsersImportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostUsersImportParams() *PostUsersImportParams {
	return &PostUsersImportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostUsersImportParamsWithTimeout creates a new PostUsersImportParams object
// with the ability to set a timeout on a request.
func NewPostUsersImportParamsWithTimeout(timeout time.Duration) *PostUsersImportParams {
	return &PostUsersImportParams{
		timeout: timeout,
	}
}

// NewPostUsersImportParamsWithContext creates a new PostUsersImportParams object
// with the ability to set a context for a request.
func NewPostUsersImportParamsWithContext(ctx context.Context) *PostUsersImportParams {
	return &PostUsersImportParams{
		Context: ctx,
	}
}

// NewPostUsersImportParamsWithHTTPClient creates a new PostUsersImportParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostUsersImportParamsWithHTTPClient(client *http.Client) *PostUsersImportParams {
	return &PostUsersImportParams{
		HTTPClient: client,
	}
}

/* PostUsersImportParams contains all the parameters to send to the API endpoint
   for the post users import operation.

   Typically these are written to a http.Request.
*/
type PostUsersImportParams struct {

	// Body.
	Body models.ImportUsersListSchema

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post users import params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUsersImportParams) WithDefaults() *PostUsersImportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post users import params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUsersImportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post users import params
func (o *PostUsersImportParams) WithTimeout(timeout time.Duration) *PostUsersImportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post users import params
func (o *PostUsersImportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post users import params
func (o *PostUsersImportParams) WithContext(ctx context.Context) *PostUsersImportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post users import params
func (o *PostUsersImportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post users import params
func (o *PostUsersImportParams) WithHTTPClient(client *http.Client) *PostUsersImportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post users import params
func (o *PostUsersImportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post users import params
func (o *PostUsersImportParams) WithBody(body models.ImportUsersListSchema) *PostUsersImportParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post users import params
func (o *PostUsersImportParams) SetBody(body models.ImportUsersListSchema) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostUsersImportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
