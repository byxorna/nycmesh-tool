// Code generated by go-swagger; DO NOT EDIT.

package sites

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

// NewPostSitesUcrmBindingsParams creates a new PostSitesUcrmBindingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostSitesUcrmBindingsParams() *PostSitesUcrmBindingsParams {
	return &PostSitesUcrmBindingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostSitesUcrmBindingsParamsWithTimeout creates a new PostSitesUcrmBindingsParams object
// with the ability to set a timeout on a request.
func NewPostSitesUcrmBindingsParamsWithTimeout(timeout time.Duration) *PostSitesUcrmBindingsParams {
	return &PostSitesUcrmBindingsParams{
		timeout: timeout,
	}
}

// NewPostSitesUcrmBindingsParamsWithContext creates a new PostSitesUcrmBindingsParams object
// with the ability to set a context for a request.
func NewPostSitesUcrmBindingsParamsWithContext(ctx context.Context) *PostSitesUcrmBindingsParams {
	return &PostSitesUcrmBindingsParams{
		Context: ctx,
	}
}

// NewPostSitesUcrmBindingsParamsWithHTTPClient creates a new PostSitesUcrmBindingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostSitesUcrmBindingsParamsWithHTTPClient(client *http.Client) *PostSitesUcrmBindingsParams {
	return &PostSitesUcrmBindingsParams{
		HTTPClient: client,
	}
}

/* PostSitesUcrmBindingsParams contains all the parameters to send to the API endpoint
   for the post sites ucrm bindings operation.

   Typically these are written to a http.Request.
*/
type PostSitesUcrmBindingsParams struct {

	// Body.
	Body models.ListOfSiteBindings

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post sites ucrm bindings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSitesUcrmBindingsParams) WithDefaults() *PostSitesUcrmBindingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post sites ucrm bindings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSitesUcrmBindingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post sites ucrm bindings params
func (o *PostSitesUcrmBindingsParams) WithTimeout(timeout time.Duration) *PostSitesUcrmBindingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post sites ucrm bindings params
func (o *PostSitesUcrmBindingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post sites ucrm bindings params
func (o *PostSitesUcrmBindingsParams) WithContext(ctx context.Context) *PostSitesUcrmBindingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post sites ucrm bindings params
func (o *PostSitesUcrmBindingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post sites ucrm bindings params
func (o *PostSitesUcrmBindingsParams) WithHTTPClient(client *http.Client) *PostSitesUcrmBindingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post sites ucrm bindings params
func (o *PostSitesUcrmBindingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post sites ucrm bindings params
func (o *PostSitesUcrmBindingsParams) WithBody(body models.ListOfSiteBindings) *PostSitesUcrmBindingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post sites ucrm bindings params
func (o *PostSitesUcrmBindingsParams) SetBody(body models.ListOfSiteBindings) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostSitesUcrmBindingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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