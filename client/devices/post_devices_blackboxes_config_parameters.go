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

// NewPostDevicesBlackboxesConfigParams creates a new PostDevicesBlackboxesConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDevicesBlackboxesConfigParams() *PostDevicesBlackboxesConfigParams {
	return &PostDevicesBlackboxesConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesBlackboxesConfigParamsWithTimeout creates a new PostDevicesBlackboxesConfigParams object
// with the ability to set a timeout on a request.
func NewPostDevicesBlackboxesConfigParamsWithTimeout(timeout time.Duration) *PostDevicesBlackboxesConfigParams {
	return &PostDevicesBlackboxesConfigParams{
		timeout: timeout,
	}
}

// NewPostDevicesBlackboxesConfigParamsWithContext creates a new PostDevicesBlackboxesConfigParams object
// with the ability to set a context for a request.
func NewPostDevicesBlackboxesConfigParamsWithContext(ctx context.Context) *PostDevicesBlackboxesConfigParams {
	return &PostDevicesBlackboxesConfigParams{
		Context: ctx,
	}
}

// NewPostDevicesBlackboxesConfigParamsWithHTTPClient creates a new PostDevicesBlackboxesConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDevicesBlackboxesConfigParamsWithHTTPClient(client *http.Client) *PostDevicesBlackboxesConfigParams {
	return &PostDevicesBlackboxesConfigParams{
		HTTPClient: client,
	}
}

/* PostDevicesBlackboxesConfigParams contains all the parameters to send to the API endpoint
   for the post devices blackboxes config operation.

   Typically these are written to a http.Request.
*/
type PostDevicesBlackboxesConfigParams struct {

	// Body.
	Body *models.DeviceBlackBoxConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post devices blackboxes config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesBlackboxesConfigParams) WithDefaults() *PostDevicesBlackboxesConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post devices blackboxes config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesBlackboxesConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post devices blackboxes config params
func (o *PostDevicesBlackboxesConfigParams) WithTimeout(timeout time.Duration) *PostDevicesBlackboxesConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices blackboxes config params
func (o *PostDevicesBlackboxesConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices blackboxes config params
func (o *PostDevicesBlackboxesConfigParams) WithContext(ctx context.Context) *PostDevicesBlackboxesConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices blackboxes config params
func (o *PostDevicesBlackboxesConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices blackboxes config params
func (o *PostDevicesBlackboxesConfigParams) WithHTTPClient(client *http.Client) *PostDevicesBlackboxesConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices blackboxes config params
func (o *PostDevicesBlackboxesConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post devices blackboxes config params
func (o *PostDevicesBlackboxesConfigParams) WithBody(body *models.DeviceBlackBoxConfig) *PostDevicesBlackboxesConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post devices blackboxes config params
func (o *PostDevicesBlackboxesConfigParams) SetBody(body *models.DeviceBlackBoxConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesBlackboxesConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
