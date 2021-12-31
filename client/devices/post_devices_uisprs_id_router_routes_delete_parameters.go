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

// NewPostDevicesUisprsIDRouterRoutesDeleteParams creates a new PostDevicesUisprsIDRouterRoutesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDevicesUisprsIDRouterRoutesDeleteParams() *PostDevicesUisprsIDRouterRoutesDeleteParams {
	return &PostDevicesUisprsIDRouterRoutesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesUisprsIDRouterRoutesDeleteParamsWithTimeout creates a new PostDevicesUisprsIDRouterRoutesDeleteParams object
// with the ability to set a timeout on a request.
func NewPostDevicesUisprsIDRouterRoutesDeleteParamsWithTimeout(timeout time.Duration) *PostDevicesUisprsIDRouterRoutesDeleteParams {
	return &PostDevicesUisprsIDRouterRoutesDeleteParams{
		timeout: timeout,
	}
}

// NewPostDevicesUisprsIDRouterRoutesDeleteParamsWithContext creates a new PostDevicesUisprsIDRouterRoutesDeleteParams object
// with the ability to set a context for a request.
func NewPostDevicesUisprsIDRouterRoutesDeleteParamsWithContext(ctx context.Context) *PostDevicesUisprsIDRouterRoutesDeleteParams {
	return &PostDevicesUisprsIDRouterRoutesDeleteParams{
		Context: ctx,
	}
}

// NewPostDevicesUisprsIDRouterRoutesDeleteParamsWithHTTPClient creates a new PostDevicesUisprsIDRouterRoutesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDevicesUisprsIDRouterRoutesDeleteParamsWithHTTPClient(client *http.Client) *PostDevicesUisprsIDRouterRoutesDeleteParams {
	return &PostDevicesUisprsIDRouterRoutesDeleteParams{
		HTTPClient: client,
	}
}

/* PostDevicesUisprsIDRouterRoutesDeleteParams contains all the parameters to send to the API endpoint
   for the post devices uisprs Id router routes delete operation.

   Typically these are written to a http.Request.
*/
type PostDevicesUisprsIDRouterRoutesDeleteParams struct {

	// Body.
	Body *models.RouterRoute3

	// ID.
	ID string

	// IsForceApply.
	IsForceApply *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post devices uisprs Id router routes delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesUisprsIDRouterRoutesDeleteParams) WithDefaults() *PostDevicesUisprsIDRouterRoutesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post devices uisprs Id router routes delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesUisprsIDRouterRoutesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post devices uisprs Id router routes delete params
func (o *PostDevicesUisprsIDRouterRoutesDeleteParams) WithTimeout(timeout time.Duration) *PostDevicesUisprsIDRouterRoutesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices uisprs Id router routes delete params
func (o *PostDevicesUisprsIDRouterRoutesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices uisprs Id router routes delete params
func (o *PostDevicesUisprsIDRouterRoutesDeleteParams) WithContext(ctx context.Context) *PostDevicesUisprsIDRouterRoutesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices uisprs Id router routes delete params
func (o *PostDevicesUisprsIDRouterRoutesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices uisprs Id router routes delete params
func (o *PostDevicesUisprsIDRouterRoutesDeleteParams) WithHTTPClient(client *http.Client) *PostDevicesUisprsIDRouterRoutesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices uisprs Id router routes delete params
func (o *PostDevicesUisprsIDRouterRoutesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post devices uisprs Id router routes delete params
func (o *PostDevicesUisprsIDRouterRoutesDeleteParams) WithBody(body *models.RouterRoute3) *PostDevicesUisprsIDRouterRoutesDeleteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post devices uisprs Id router routes delete params
func (o *PostDevicesUisprsIDRouterRoutesDeleteParams) SetBody(body *models.RouterRoute3) {
	o.Body = body
}

// WithID adds the id to the post devices uisprs Id router routes delete params
func (o *PostDevicesUisprsIDRouterRoutesDeleteParams) WithID(id string) *PostDevicesUisprsIDRouterRoutesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post devices uisprs Id router routes delete params
func (o *PostDevicesUisprsIDRouterRoutesDeleteParams) SetID(id string) {
	o.ID = id
}

// WithIsForceApply adds the isForceApply to the post devices uisprs Id router routes delete params
func (o *PostDevicesUisprsIDRouterRoutesDeleteParams) WithIsForceApply(isForceApply *string) *PostDevicesUisprsIDRouterRoutesDeleteParams {
	o.SetIsForceApply(isForceApply)
	return o
}

// SetIsForceApply adds the isForceApply to the post devices uisprs Id router routes delete params
func (o *PostDevicesUisprsIDRouterRoutesDeleteParams) SetIsForceApply(isForceApply *string) {
	o.IsForceApply = isForceApply
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesUisprsIDRouterRoutesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IsForceApply != nil {

		// query param isForceApply
		var qrIsForceApply string

		if o.IsForceApply != nil {
			qrIsForceApply = *o.IsForceApply
		}
		qIsForceApply := qrIsForceApply
		if qIsForceApply != "" {

			if err := r.SetQueryParam("isForceApply", qIsForceApply); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
