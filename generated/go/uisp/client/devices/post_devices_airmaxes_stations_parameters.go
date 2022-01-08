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

// NewPostDevicesAirmaxesStationsParams creates a new PostDevicesAirmaxesStationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDevicesAirmaxesStationsParams() *PostDevicesAirmaxesStationsParams {
	return &PostDevicesAirmaxesStationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesAirmaxesStationsParamsWithTimeout creates a new PostDevicesAirmaxesStationsParams object
// with the ability to set a timeout on a request.
func NewPostDevicesAirmaxesStationsParamsWithTimeout(timeout time.Duration) *PostDevicesAirmaxesStationsParams {
	return &PostDevicesAirmaxesStationsParams{
		timeout: timeout,
	}
}

// NewPostDevicesAirmaxesStationsParamsWithContext creates a new PostDevicesAirmaxesStationsParams object
// with the ability to set a context for a request.
func NewPostDevicesAirmaxesStationsParamsWithContext(ctx context.Context) *PostDevicesAirmaxesStationsParams {
	return &PostDevicesAirmaxesStationsParams{
		Context: ctx,
	}
}

// NewPostDevicesAirmaxesStationsParamsWithHTTPClient creates a new PostDevicesAirmaxesStationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDevicesAirmaxesStationsParamsWithHTTPClient(client *http.Client) *PostDevicesAirmaxesStationsParams {
	return &PostDevicesAirmaxesStationsParams{
		HTTPClient: client,
	}
}

/* PostDevicesAirmaxesStationsParams contains all the parameters to send to the API endpoint
   for the post devices airmaxes stations operation.

   Typically these are written to a http.Request.
*/
type PostDevicesAirmaxesStationsParams struct {

	// Body.
	Body *models.PayloadStationsCount

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post devices airmaxes stations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesAirmaxesStationsParams) WithDefaults() *PostDevicesAirmaxesStationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post devices airmaxes stations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesAirmaxesStationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post devices airmaxes stations params
func (o *PostDevicesAirmaxesStationsParams) WithTimeout(timeout time.Duration) *PostDevicesAirmaxesStationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices airmaxes stations params
func (o *PostDevicesAirmaxesStationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices airmaxes stations params
func (o *PostDevicesAirmaxesStationsParams) WithContext(ctx context.Context) *PostDevicesAirmaxesStationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices airmaxes stations params
func (o *PostDevicesAirmaxesStationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices airmaxes stations params
func (o *PostDevicesAirmaxesStationsParams) WithHTTPClient(client *http.Client) *PostDevicesAirmaxesStationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices airmaxes stations params
func (o *PostDevicesAirmaxesStationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post devices airmaxes stations params
func (o *PostDevicesAirmaxesStationsParams) WithBody(body *models.PayloadStationsCount) *PostDevicesAirmaxesStationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post devices airmaxes stations params
func (o *PostDevicesAirmaxesStationsParams) SetBody(body *models.PayloadStationsCount) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesAirmaxesStationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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