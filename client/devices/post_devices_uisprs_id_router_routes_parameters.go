// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// NewPostDevicesUisprsIDRouterRoutesParams creates a new PostDevicesUisprsIDRouterRoutesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDevicesUisprsIDRouterRoutesParams() *PostDevicesUisprsIDRouterRoutesParams {
	return &PostDevicesUisprsIDRouterRoutesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesUisprsIDRouterRoutesParamsWithTimeout creates a new PostDevicesUisprsIDRouterRoutesParams object
// with the ability to set a timeout on a request.
func NewPostDevicesUisprsIDRouterRoutesParamsWithTimeout(timeout time.Duration) *PostDevicesUisprsIDRouterRoutesParams {
	return &PostDevicesUisprsIDRouterRoutesParams{
		timeout: timeout,
	}
}

// NewPostDevicesUisprsIDRouterRoutesParamsWithContext creates a new PostDevicesUisprsIDRouterRoutesParams object
// with the ability to set a context for a request.
func NewPostDevicesUisprsIDRouterRoutesParamsWithContext(ctx context.Context) *PostDevicesUisprsIDRouterRoutesParams {
	return &PostDevicesUisprsIDRouterRoutesParams{
		Context: ctx,
	}
}

// NewPostDevicesUisprsIDRouterRoutesParamsWithHTTPClient creates a new PostDevicesUisprsIDRouterRoutesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDevicesUisprsIDRouterRoutesParamsWithHTTPClient(client *http.Client) *PostDevicesUisprsIDRouterRoutesParams {
	return &PostDevicesUisprsIDRouterRoutesParams{
		HTTPClient: client,
	}
}

/* PostDevicesUisprsIDRouterRoutesParams contains all the parameters to send to the API endpoint
   for the post devices uisprs Id router routes operation.

   Typically these are written to a http.Request.
*/
type PostDevicesUisprsIDRouterRoutesParams struct {

	// Body.
	Body *models.RouterRouteUpsert1

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post devices uisprs Id router routes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesUisprsIDRouterRoutesParams) WithDefaults() *PostDevicesUisprsIDRouterRoutesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post devices uisprs Id router routes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesUisprsIDRouterRoutesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post devices uisprs Id router routes params
func (o *PostDevicesUisprsIDRouterRoutesParams) WithTimeout(timeout time.Duration) *PostDevicesUisprsIDRouterRoutesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices uisprs Id router routes params
func (o *PostDevicesUisprsIDRouterRoutesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices uisprs Id router routes params
func (o *PostDevicesUisprsIDRouterRoutesParams) WithContext(ctx context.Context) *PostDevicesUisprsIDRouterRoutesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices uisprs Id router routes params
func (o *PostDevicesUisprsIDRouterRoutesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices uisprs Id router routes params
func (o *PostDevicesUisprsIDRouterRoutesParams) WithHTTPClient(client *http.Client) *PostDevicesUisprsIDRouterRoutesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices uisprs Id router routes params
func (o *PostDevicesUisprsIDRouterRoutesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post devices uisprs Id router routes params
func (o *PostDevicesUisprsIDRouterRoutesParams) WithBody(body *models.RouterRouteUpsert1) *PostDevicesUisprsIDRouterRoutesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post devices uisprs Id router routes params
func (o *PostDevicesUisprsIDRouterRoutesParams) SetBody(body *models.RouterRouteUpsert1) {
	o.Body = body
}

// WithID adds the id to the post devices uisprs Id router routes params
func (o *PostDevicesUisprsIDRouterRoutesParams) WithID(id string) *PostDevicesUisprsIDRouterRoutesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post devices uisprs Id router routes params
func (o *PostDevicesUisprsIDRouterRoutesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesUisprsIDRouterRoutesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
