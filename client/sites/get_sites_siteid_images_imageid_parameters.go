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
)

// NewGetSitesSiteidImagesImageidParams creates a new GetSitesSiteidImagesImageidParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSitesSiteidImagesImageidParams() *GetSitesSiteidImagesImageidParams {
	return &GetSitesSiteidImagesImageidParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSitesSiteidImagesImageidParamsWithTimeout creates a new GetSitesSiteidImagesImageidParams object
// with the ability to set a timeout on a request.
func NewGetSitesSiteidImagesImageidParamsWithTimeout(timeout time.Duration) *GetSitesSiteidImagesImageidParams {
	return &GetSitesSiteidImagesImageidParams{
		timeout: timeout,
	}
}

// NewGetSitesSiteidImagesImageidParamsWithContext creates a new GetSitesSiteidImagesImageidParams object
// with the ability to set a context for a request.
func NewGetSitesSiteidImagesImageidParamsWithContext(ctx context.Context) *GetSitesSiteidImagesImageidParams {
	return &GetSitesSiteidImagesImageidParams{
		Context: ctx,
	}
}

// NewGetSitesSiteidImagesImageidParamsWithHTTPClient creates a new GetSitesSiteidImagesImageidParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSitesSiteidImagesImageidParamsWithHTTPClient(client *http.Client) *GetSitesSiteidImagesImageidParams {
	return &GetSitesSiteidImagesImageidParams{
		HTTPClient: client,
	}
}

/* GetSitesSiteidImagesImageidParams contains all the parameters to send to the API endpoint
   for the get sites siteid images imageid operation.

   Typically these are written to a http.Request.
*/
type GetSitesSiteidImagesImageidParams struct {

	// ImageID.
	ImageID string

	// SiteID.
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get sites siteid images imageid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSitesSiteidImagesImageidParams) WithDefaults() *GetSitesSiteidImagesImageidParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get sites siteid images imageid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSitesSiteidImagesImageidParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get sites siteid images imageid params
func (o *GetSitesSiteidImagesImageidParams) WithTimeout(timeout time.Duration) *GetSitesSiteidImagesImageidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sites siteid images imageid params
func (o *GetSitesSiteidImagesImageidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sites siteid images imageid params
func (o *GetSitesSiteidImagesImageidParams) WithContext(ctx context.Context) *GetSitesSiteidImagesImageidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sites siteid images imageid params
func (o *GetSitesSiteidImagesImageidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sites siteid images imageid params
func (o *GetSitesSiteidImagesImageidParams) WithHTTPClient(client *http.Client) *GetSitesSiteidImagesImageidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sites siteid images imageid params
func (o *GetSitesSiteidImagesImageidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImageID adds the imageID to the get sites siteid images imageid params
func (o *GetSitesSiteidImagesImageidParams) WithImageID(imageID string) *GetSitesSiteidImagesImageidParams {
	o.SetImageID(imageID)
	return o
}

// SetImageID adds the imageId to the get sites siteid images imageid params
func (o *GetSitesSiteidImagesImageidParams) SetImageID(imageID string) {
	o.ImageID = imageID
}

// WithSiteID adds the siteID to the get sites siteid images imageid params
func (o *GetSitesSiteidImagesImageidParams) WithSiteID(siteID string) *GetSitesSiteidImagesImageidParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get sites siteid images imageid params
func (o *GetSitesSiteidImagesImageidParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSitesSiteidImagesImageidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param imageId
	if err := r.SetPathParam("imageId", o.ImageID); err != nil {
		return err
	}

	// path param siteId
	if err := r.SetPathParam("siteId", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
