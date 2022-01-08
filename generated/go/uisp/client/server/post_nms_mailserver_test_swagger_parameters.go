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

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// NewPostNmsMailserverTestParams creates a new PostNmsMailserverTestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostNmsMailserverTestParams() *PostNmsMailserverTestParams {
	return &PostNmsMailserverTestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostNmsMailserverTestParamsWithTimeout creates a new PostNmsMailserverTestParams object
// with the ability to set a timeout on a request.
func NewPostNmsMailserverTestParamsWithTimeout(timeout time.Duration) *PostNmsMailserverTestParams {
	return &PostNmsMailserverTestParams{
		timeout: timeout,
	}
}

// NewPostNmsMailserverTestParamsWithContext creates a new PostNmsMailserverTestParams object
// with the ability to set a context for a request.
func NewPostNmsMailserverTestParamsWithContext(ctx context.Context) *PostNmsMailserverTestParams {
	return &PostNmsMailserverTestParams{
		Context: ctx,
	}
}

// NewPostNmsMailserverTestParamsWithHTTPClient creates a new PostNmsMailserverTestParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostNmsMailserverTestParamsWithHTTPClient(client *http.Client) *PostNmsMailserverTestParams {
	return &PostNmsMailserverTestParams{
		HTTPClient: client,
	}
}

/* PostNmsMailserverTestParams contains all the parameters to send to the API endpoint
   for the post nms mailserver test operation.

   Typically these are written to a http.Request.
*/
type PostNmsMailserverTestParams struct {

	// Body.
	Body *models.Model85

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post nms mailserver test params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostNmsMailserverTestParams) WithDefaults() *PostNmsMailserverTestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post nms mailserver test params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostNmsMailserverTestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post nms mailserver test params
func (o *PostNmsMailserverTestParams) WithTimeout(timeout time.Duration) *PostNmsMailserverTestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post nms mailserver test params
func (o *PostNmsMailserverTestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post nms mailserver test params
func (o *PostNmsMailserverTestParams) WithContext(ctx context.Context) *PostNmsMailserverTestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post nms mailserver test params
func (o *PostNmsMailserverTestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post nms mailserver test params
func (o *PostNmsMailserverTestParams) WithHTTPClient(client *http.Client) *PostNmsMailserverTestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post nms mailserver test params
func (o *PostNmsMailserverTestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post nms mailserver test params
func (o *PostNmsMailserverTestParams) WithBody(body *models.Model85) *PostNmsMailserverTestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post nms mailserver test params
func (o *PostNmsMailserverTestParams) SetBody(body *models.Model85) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostNmsMailserverTestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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