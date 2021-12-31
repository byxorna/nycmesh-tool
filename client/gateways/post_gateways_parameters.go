// Code generated by go-swagger; DO NOT EDIT.

package gateways

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

// NewPostGatewaysParams creates a new PostGatewaysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostGatewaysParams() *PostGatewaysParams {
	return &PostGatewaysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostGatewaysParamsWithTimeout creates a new PostGatewaysParams object
// with the ability to set a timeout on a request.
func NewPostGatewaysParamsWithTimeout(timeout time.Duration) *PostGatewaysParams {
	return &PostGatewaysParams{
		timeout: timeout,
	}
}

// NewPostGatewaysParamsWithContext creates a new PostGatewaysParams object
// with the ability to set a context for a request.
func NewPostGatewaysParamsWithContext(ctx context.Context) *PostGatewaysParams {
	return &PostGatewaysParams{
		Context: ctx,
	}
}

// NewPostGatewaysParamsWithHTTPClient creates a new PostGatewaysParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostGatewaysParamsWithHTTPClient(client *http.Client) *PostGatewaysParams {
	return &PostGatewaysParams{
		HTTPClient: client,
	}
}

/* PostGatewaysParams contains all the parameters to send to the API endpoint
   for the post gateways operation.

   Typically these are written to a http.Request.
*/
type PostGatewaysParams struct {

	// Body.
	Body *models.NewGateway

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post gateways params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostGatewaysParams) WithDefaults() *PostGatewaysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post gateways params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostGatewaysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post gateways params
func (o *PostGatewaysParams) WithTimeout(timeout time.Duration) *PostGatewaysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post gateways params
func (o *PostGatewaysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post gateways params
func (o *PostGatewaysParams) WithContext(ctx context.Context) *PostGatewaysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post gateways params
func (o *PostGatewaysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post gateways params
func (o *PostGatewaysParams) WithHTTPClient(client *http.Client) *PostGatewaysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post gateways params
func (o *PostGatewaysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post gateways params
func (o *PostGatewaysParams) WithBody(body *models.NewGateway) *PostGatewaysParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post gateways params
func (o *PostGatewaysParams) SetBody(body *models.NewGateway) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostGatewaysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
