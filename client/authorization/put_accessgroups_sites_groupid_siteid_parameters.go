// Code generated by go-swagger; DO NOT EDIT.

package authorization

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

// NewPutAccessgroupsSitesGroupidSiteidParams creates a new PutAccessgroupsSitesGroupidSiteidParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutAccessgroupsSitesGroupidSiteidParams() *PutAccessgroupsSitesGroupidSiteidParams {
	return &PutAccessgroupsSitesGroupidSiteidParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutAccessgroupsSitesGroupidSiteidParamsWithTimeout creates a new PutAccessgroupsSitesGroupidSiteidParams object
// with the ability to set a timeout on a request.
func NewPutAccessgroupsSitesGroupidSiteidParamsWithTimeout(timeout time.Duration) *PutAccessgroupsSitesGroupidSiteidParams {
	return &PutAccessgroupsSitesGroupidSiteidParams{
		timeout: timeout,
	}
}

// NewPutAccessgroupsSitesGroupidSiteidParamsWithContext creates a new PutAccessgroupsSitesGroupidSiteidParams object
// with the ability to set a context for a request.
func NewPutAccessgroupsSitesGroupidSiteidParamsWithContext(ctx context.Context) *PutAccessgroupsSitesGroupidSiteidParams {
	return &PutAccessgroupsSitesGroupidSiteidParams{
		Context: ctx,
	}
}

// NewPutAccessgroupsSitesGroupidSiteidParamsWithHTTPClient creates a new PutAccessgroupsSitesGroupidSiteidParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutAccessgroupsSitesGroupidSiteidParamsWithHTTPClient(client *http.Client) *PutAccessgroupsSitesGroupidSiteidParams {
	return &PutAccessgroupsSitesGroupidSiteidParams{
		HTTPClient: client,
	}
}

/* PutAccessgroupsSitesGroupidSiteidParams contains all the parameters to send to the API endpoint
   for the put accessgroups sites groupid siteid operation.

   Typically these are written to a http.Request.
*/
type PutAccessgroupsSitesGroupidSiteidParams struct {

	// Body.
	Body *models.RequestSiteAccessGroupSingleSite

	// GroupID.
	GroupID string

	// SiteID.
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put accessgroups sites groupid siteid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutAccessgroupsSitesGroupidSiteidParams) WithDefaults() *PutAccessgroupsSitesGroupidSiteidParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put accessgroups sites groupid siteid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutAccessgroupsSitesGroupidSiteidParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put accessgroups sites groupid siteid params
func (o *PutAccessgroupsSitesGroupidSiteidParams) WithTimeout(timeout time.Duration) *PutAccessgroupsSitesGroupidSiteidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put accessgroups sites groupid siteid params
func (o *PutAccessgroupsSitesGroupidSiteidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put accessgroups sites groupid siteid params
func (o *PutAccessgroupsSitesGroupidSiteidParams) WithContext(ctx context.Context) *PutAccessgroupsSitesGroupidSiteidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put accessgroups sites groupid siteid params
func (o *PutAccessgroupsSitesGroupidSiteidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put accessgroups sites groupid siteid params
func (o *PutAccessgroupsSitesGroupidSiteidParams) WithHTTPClient(client *http.Client) *PutAccessgroupsSitesGroupidSiteidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put accessgroups sites groupid siteid params
func (o *PutAccessgroupsSitesGroupidSiteidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put accessgroups sites groupid siteid params
func (o *PutAccessgroupsSitesGroupidSiteidParams) WithBody(body *models.RequestSiteAccessGroupSingleSite) *PutAccessgroupsSitesGroupidSiteidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put accessgroups sites groupid siteid params
func (o *PutAccessgroupsSitesGroupidSiteidParams) SetBody(body *models.RequestSiteAccessGroupSingleSite) {
	o.Body = body
}

// WithGroupID adds the groupID to the put accessgroups sites groupid siteid params
func (o *PutAccessgroupsSitesGroupidSiteidParams) WithGroupID(groupID string) *PutAccessgroupsSitesGroupidSiteidParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the put accessgroups sites groupid siteid params
func (o *PutAccessgroupsSitesGroupidSiteidParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithSiteID adds the siteID to the put accessgroups sites groupid siteid params
func (o *PutAccessgroupsSitesGroupidSiteidParams) WithSiteID(siteID string) *PutAccessgroupsSitesGroupidSiteidParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the put accessgroups sites groupid siteid params
func (o *PutAccessgroupsSitesGroupidSiteidParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *PutAccessgroupsSitesGroupidSiteidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
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
