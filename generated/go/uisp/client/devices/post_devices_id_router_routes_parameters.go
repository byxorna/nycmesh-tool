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

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// NewPostDevicesIDRouterRoutesParams creates a new PostDevicesIDRouterRoutesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDevicesIDRouterRoutesParams() *PostDevicesIDRouterRoutesParams {
	return &PostDevicesIDRouterRoutesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesIDRouterRoutesParamsWithTimeout creates a new PostDevicesIDRouterRoutesParams object
// with the ability to set a timeout on a request.
func NewPostDevicesIDRouterRoutesParamsWithTimeout(timeout time.Duration) *PostDevicesIDRouterRoutesParams {
	return &PostDevicesIDRouterRoutesParams{
		timeout: timeout,
	}
}

// NewPostDevicesIDRouterRoutesParamsWithContext creates a new PostDevicesIDRouterRoutesParams object
// with the ability to set a context for a request.
func NewPostDevicesIDRouterRoutesParamsWithContext(ctx context.Context) *PostDevicesIDRouterRoutesParams {
	return &PostDevicesIDRouterRoutesParams{
		Context: ctx,
	}
}

// NewPostDevicesIDRouterRoutesParamsWithHTTPClient creates a new PostDevicesIDRouterRoutesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDevicesIDRouterRoutesParamsWithHTTPClient(client *http.Client) *PostDevicesIDRouterRoutesParams {
	return &PostDevicesIDRouterRoutesParams{
		HTTPClient: client,
	}
}

/* PostDevicesIDRouterRoutesParams contains all the parameters to send to the API endpoint
   for the post devices Id router routes operation.

   Typically these are written to a http.Request.
*/
type PostDevicesIDRouterRoutesParams struct {

	// Body.
	Body *models.RouterRouteUpsert

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post devices Id router routes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesIDRouterRoutesParams) WithDefaults() *PostDevicesIDRouterRoutesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post devices Id router routes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesIDRouterRoutesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post devices Id router routes params
func (o *PostDevicesIDRouterRoutesParams) WithTimeout(timeout time.Duration) *PostDevicesIDRouterRoutesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices Id router routes params
func (o *PostDevicesIDRouterRoutesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices Id router routes params
func (o *PostDevicesIDRouterRoutesParams) WithContext(ctx context.Context) *PostDevicesIDRouterRoutesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices Id router routes params
func (o *PostDevicesIDRouterRoutesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices Id router routes params
func (o *PostDevicesIDRouterRoutesParams) WithHTTPClient(client *http.Client) *PostDevicesIDRouterRoutesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices Id router routes params
func (o *PostDevicesIDRouterRoutesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post devices Id router routes params
func (o *PostDevicesIDRouterRoutesParams) WithBody(body *models.RouterRouteUpsert) *PostDevicesIDRouterRoutesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post devices Id router routes params
func (o *PostDevicesIDRouterRoutesParams) SetBody(body *models.RouterRouteUpsert) {
	o.Body = body
}

// WithID adds the id to the post devices Id router routes params
func (o *PostDevicesIDRouterRoutesParams) WithID(id string) *PostDevicesIDRouterRoutesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post devices Id router routes params
func (o *PostDevicesIDRouterRoutesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesIDRouterRoutesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
