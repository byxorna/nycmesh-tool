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

// NewPostSitesBulkdeleteParams creates a new PostSitesBulkdeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostSitesBulkdeleteParams() *PostSitesBulkdeleteParams {
	return &PostSitesBulkdeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostSitesBulkdeleteParamsWithTimeout creates a new PostSitesBulkdeleteParams object
// with the ability to set a timeout on a request.
func NewPostSitesBulkdeleteParamsWithTimeout(timeout time.Duration) *PostSitesBulkdeleteParams {
	return &PostSitesBulkdeleteParams{
		timeout: timeout,
	}
}

// NewPostSitesBulkdeleteParamsWithContext creates a new PostSitesBulkdeleteParams object
// with the ability to set a context for a request.
func NewPostSitesBulkdeleteParamsWithContext(ctx context.Context) *PostSitesBulkdeleteParams {
	return &PostSitesBulkdeleteParams{
		Context: ctx,
	}
}

// NewPostSitesBulkdeleteParamsWithHTTPClient creates a new PostSitesBulkdeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostSitesBulkdeleteParamsWithHTTPClient(client *http.Client) *PostSitesBulkdeleteParams {
	return &PostSitesBulkdeleteParams{
		HTTPClient: client,
	}
}

/* PostSitesBulkdeleteParams contains all the parameters to send to the API endpoint
   for the post sites bulkdelete operation.

   Typically these are written to a http.Request.
*/
type PostSitesBulkdeleteParams struct {

	// Body.
	Body *models.ListOfSiteIds

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post sites bulkdelete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSitesBulkdeleteParams) WithDefaults() *PostSitesBulkdeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post sites bulkdelete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSitesBulkdeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post sites bulkdelete params
func (o *PostSitesBulkdeleteParams) WithTimeout(timeout time.Duration) *PostSitesBulkdeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post sites bulkdelete params
func (o *PostSitesBulkdeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post sites bulkdelete params
func (o *PostSitesBulkdeleteParams) WithContext(ctx context.Context) *PostSitesBulkdeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post sites bulkdelete params
func (o *PostSitesBulkdeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post sites bulkdelete params
func (o *PostSitesBulkdeleteParams) WithHTTPClient(client *http.Client) *PostSitesBulkdeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post sites bulkdelete params
func (o *PostSitesBulkdeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post sites bulkdelete params
func (o *PostSitesBulkdeleteParams) WithBody(body *models.ListOfSiteIds) *PostSitesBulkdeleteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post sites bulkdelete params
func (o *PostSitesBulkdeleteParams) SetBody(body *models.ListOfSiteIds) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostSitesBulkdeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
