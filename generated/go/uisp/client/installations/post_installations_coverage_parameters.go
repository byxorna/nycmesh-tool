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

// NewPostInstallationsCoverageParams creates a new PostInstallationsCoverageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostInstallationsCoverageParams() *PostInstallationsCoverageParams {
	return &PostInstallationsCoverageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostInstallationsCoverageParamsWithTimeout creates a new PostInstallationsCoverageParams object
// with the ability to set a timeout on a request.
func NewPostInstallationsCoverageParamsWithTimeout(timeout time.Duration) *PostInstallationsCoverageParams {
	return &PostInstallationsCoverageParams{
		timeout: timeout,
	}
}

// NewPostInstallationsCoverageParamsWithContext creates a new PostInstallationsCoverageParams object
// with the ability to set a context for a request.
func NewPostInstallationsCoverageParamsWithContext(ctx context.Context) *PostInstallationsCoverageParams {
	return &PostInstallationsCoverageParams{
		Context: ctx,
	}
}

// NewPostInstallationsCoverageParamsWithHTTPClient creates a new PostInstallationsCoverageParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostInstallationsCoverageParamsWithHTTPClient(client *http.Client) *PostInstallationsCoverageParams {
	return &PostInstallationsCoverageParams{
		HTTPClient: client,
	}
}

/* PostInstallationsCoverageParams contains all the parameters to send to the API endpoint
   for the post installations coverage operation.

   Typically these are written to a http.Request.
*/
type PostInstallationsCoverageParams struct {

	// Body.
	Body *models.Model75

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post installations coverage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostInstallationsCoverageParams) WithDefaults() *PostInstallationsCoverageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post installations coverage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostInstallationsCoverageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post installations coverage params
func (o *PostInstallationsCoverageParams) WithTimeout(timeout time.Duration) *PostInstallationsCoverageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post installations coverage params
func (o *PostInstallationsCoverageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post installations coverage params
func (o *PostInstallationsCoverageParams) WithContext(ctx context.Context) *PostInstallationsCoverageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post installations coverage params
func (o *PostInstallationsCoverageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post installations coverage params
func (o *PostInstallationsCoverageParams) WithHTTPClient(client *http.Client) *PostInstallationsCoverageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post installations coverage params
func (o *PostInstallationsCoverageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post installations coverage params
func (o *PostInstallationsCoverageParams) WithBody(body *models.Model75) *PostInstallationsCoverageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post installations coverage params
func (o *PostInstallationsCoverageParams) SetBody(body *models.Model75) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostInstallationsCoverageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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