// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewPostNmsBlogPostsReadallParams creates a new PostNmsBlogPostsReadallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostNmsBlogPostsReadallParams() *PostNmsBlogPostsReadallParams {
	return &PostNmsBlogPostsReadallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostNmsBlogPostsReadallParamsWithTimeout creates a new PostNmsBlogPostsReadallParams object
// with the ability to set a timeout on a request.
func NewPostNmsBlogPostsReadallParamsWithTimeout(timeout time.Duration) *PostNmsBlogPostsReadallParams {
	return &PostNmsBlogPostsReadallParams{
		timeout: timeout,
	}
}

// NewPostNmsBlogPostsReadallParamsWithContext creates a new PostNmsBlogPostsReadallParams object
// with the ability to set a context for a request.
func NewPostNmsBlogPostsReadallParamsWithContext(ctx context.Context) *PostNmsBlogPostsReadallParams {
	return &PostNmsBlogPostsReadallParams{
		Context: ctx,
	}
}

// NewPostNmsBlogPostsReadallParamsWithHTTPClient creates a new PostNmsBlogPostsReadallParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostNmsBlogPostsReadallParamsWithHTTPClient(client *http.Client) *PostNmsBlogPostsReadallParams {
	return &PostNmsBlogPostsReadallParams{
		HTTPClient: client,
	}
}

/* PostNmsBlogPostsReadallParams contains all the parameters to send to the API endpoint
   for the post nms blog posts readall operation.

   Typically these are written to a http.Request.
*/
type PostNmsBlogPostsReadallParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post nms blog posts readall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostNmsBlogPostsReadallParams) WithDefaults() *PostNmsBlogPostsReadallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post nms blog posts readall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostNmsBlogPostsReadallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post nms blog posts readall params
func (o *PostNmsBlogPostsReadallParams) WithTimeout(timeout time.Duration) *PostNmsBlogPostsReadallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post nms blog posts readall params
func (o *PostNmsBlogPostsReadallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post nms blog posts readall params
func (o *PostNmsBlogPostsReadallParams) WithContext(ctx context.Context) *PostNmsBlogPostsReadallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post nms blog posts readall params
func (o *PostNmsBlogPostsReadallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post nms blog posts readall params
func (o *PostNmsBlogPostsReadallParams) WithHTTPClient(client *http.Client) *PostNmsBlogPostsReadallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post nms blog posts readall params
func (o *PostNmsBlogPostsReadallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostNmsBlogPostsReadallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
