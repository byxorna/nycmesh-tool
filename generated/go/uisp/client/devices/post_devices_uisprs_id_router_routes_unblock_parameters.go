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

// NewPostDevicesUisprsIDRouterRoutesUnblockParams creates a new PostDevicesUisprsIDRouterRoutesUnblockParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDevicesUisprsIDRouterRoutesUnblockParams() *PostDevicesUisprsIDRouterRoutesUnblockParams {
	return &PostDevicesUisprsIDRouterRoutesUnblockParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesUisprsIDRouterRoutesUnblockParamsWithTimeout creates a new PostDevicesUisprsIDRouterRoutesUnblockParams object
// with the ability to set a timeout on a request.
func NewPostDevicesUisprsIDRouterRoutesUnblockParamsWithTimeout(timeout time.Duration) *PostDevicesUisprsIDRouterRoutesUnblockParams {
	return &PostDevicesUisprsIDRouterRoutesUnblockParams{
		timeout: timeout,
	}
}

// NewPostDevicesUisprsIDRouterRoutesUnblockParamsWithContext creates a new PostDevicesUisprsIDRouterRoutesUnblockParams object
// with the ability to set a context for a request.
func NewPostDevicesUisprsIDRouterRoutesUnblockParamsWithContext(ctx context.Context) *PostDevicesUisprsIDRouterRoutesUnblockParams {
	return &PostDevicesUisprsIDRouterRoutesUnblockParams{
		Context: ctx,
	}
}

// NewPostDevicesUisprsIDRouterRoutesUnblockParamsWithHTTPClient creates a new PostDevicesUisprsIDRouterRoutesUnblockParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDevicesUisprsIDRouterRoutesUnblockParamsWithHTTPClient(client *http.Client) *PostDevicesUisprsIDRouterRoutesUnblockParams {
	return &PostDevicesUisprsIDRouterRoutesUnblockParams{
		HTTPClient: client,
	}
}

/* PostDevicesUisprsIDRouterRoutesUnblockParams contains all the parameters to send to the API endpoint
   for the post devices uisprs Id router routes unblock operation.

   Typically these are written to a http.Request.
*/
type PostDevicesUisprsIDRouterRoutesUnblockParams struct {

	// Body.
	Body *models.RouterRoute3

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post devices uisprs Id router routes unblock params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesUisprsIDRouterRoutesUnblockParams) WithDefaults() *PostDevicesUisprsIDRouterRoutesUnblockParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post devices uisprs Id router routes unblock params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesUisprsIDRouterRoutesUnblockParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post devices uisprs Id router routes unblock params
func (o *PostDevicesUisprsIDRouterRoutesUnblockParams) WithTimeout(timeout time.Duration) *PostDevicesUisprsIDRouterRoutesUnblockParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices uisprs Id router routes unblock params
func (o *PostDevicesUisprsIDRouterRoutesUnblockParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices uisprs Id router routes unblock params
func (o *PostDevicesUisprsIDRouterRoutesUnblockParams) WithContext(ctx context.Context) *PostDevicesUisprsIDRouterRoutesUnblockParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices uisprs Id router routes unblock params
func (o *PostDevicesUisprsIDRouterRoutesUnblockParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices uisprs Id router routes unblock params
func (o *PostDevicesUisprsIDRouterRoutesUnblockParams) WithHTTPClient(client *http.Client) *PostDevicesUisprsIDRouterRoutesUnblockParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices uisprs Id router routes unblock params
func (o *PostDevicesUisprsIDRouterRoutesUnblockParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post devices uisprs Id router routes unblock params
func (o *PostDevicesUisprsIDRouterRoutesUnblockParams) WithBody(body *models.RouterRoute3) *PostDevicesUisprsIDRouterRoutesUnblockParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post devices uisprs Id router routes unblock params
func (o *PostDevicesUisprsIDRouterRoutesUnblockParams) SetBody(body *models.RouterRoute3) {
	o.Body = body
}

// WithID adds the id to the post devices uisprs Id router routes unblock params
func (o *PostDevicesUisprsIDRouterRoutesUnblockParams) WithID(id string) *PostDevicesUisprsIDRouterRoutesUnblockParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post devices uisprs Id router routes unblock params
func (o *PostDevicesUisprsIDRouterRoutesUnblockParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesUisprsIDRouterRoutesUnblockParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
