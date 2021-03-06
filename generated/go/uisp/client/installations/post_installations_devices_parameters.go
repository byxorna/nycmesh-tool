// Code generated by go-swagger; DO NOT EDIT.

package installations

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

// NewPostInstallationsDevicesParams creates a new PostInstallationsDevicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostInstallationsDevicesParams() *PostInstallationsDevicesParams {
	return &PostInstallationsDevicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostInstallationsDevicesParamsWithTimeout creates a new PostInstallationsDevicesParams object
// with the ability to set a timeout on a request.
func NewPostInstallationsDevicesParamsWithTimeout(timeout time.Duration) *PostInstallationsDevicesParams {
	return &PostInstallationsDevicesParams{
		timeout: timeout,
	}
}

// NewPostInstallationsDevicesParamsWithContext creates a new PostInstallationsDevicesParams object
// with the ability to set a context for a request.
func NewPostInstallationsDevicesParamsWithContext(ctx context.Context) *PostInstallationsDevicesParams {
	return &PostInstallationsDevicesParams{
		Context: ctx,
	}
}

// NewPostInstallationsDevicesParamsWithHTTPClient creates a new PostInstallationsDevicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostInstallationsDevicesParamsWithHTTPClient(client *http.Client) *PostInstallationsDevicesParams {
	return &PostInstallationsDevicesParams{
		HTTPClient: client,
	}
}

/* PostInstallationsDevicesParams contains all the parameters to send to the API endpoint
   for the post installations devices operation.

   Typically these are written to a http.Request.
*/
type PostInstallationsDevicesParams struct {

	// Body.
	Body *models.NewInstallationDevice

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post installations devices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostInstallationsDevicesParams) WithDefaults() *PostInstallationsDevicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post installations devices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostInstallationsDevicesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post installations devices params
func (o *PostInstallationsDevicesParams) WithTimeout(timeout time.Duration) *PostInstallationsDevicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post installations devices params
func (o *PostInstallationsDevicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post installations devices params
func (o *PostInstallationsDevicesParams) WithContext(ctx context.Context) *PostInstallationsDevicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post installations devices params
func (o *PostInstallationsDevicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post installations devices params
func (o *PostInstallationsDevicesParams) WithHTTPClient(client *http.Client) *PostInstallationsDevicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post installations devices params
func (o *PostInstallationsDevicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post installations devices params
func (o *PostInstallationsDevicesParams) WithBody(body *models.NewInstallationDevice) *PostInstallationsDevicesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post installations devices params
func (o *PostInstallationsDevicesParams) SetBody(body *models.NewInstallationDevice) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostInstallationsDevicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
