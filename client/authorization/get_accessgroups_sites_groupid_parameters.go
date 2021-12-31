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
)

// NewGetAccessgroupsSitesGroupidParams creates a new GetAccessgroupsSitesGroupidParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccessgroupsSitesGroupidParams() *GetAccessgroupsSitesGroupidParams {
	return &GetAccessgroupsSitesGroupidParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccessgroupsSitesGroupidParamsWithTimeout creates a new GetAccessgroupsSitesGroupidParams object
// with the ability to set a timeout on a request.
func NewGetAccessgroupsSitesGroupidParamsWithTimeout(timeout time.Duration) *GetAccessgroupsSitesGroupidParams {
	return &GetAccessgroupsSitesGroupidParams{
		timeout: timeout,
	}
}

// NewGetAccessgroupsSitesGroupidParamsWithContext creates a new GetAccessgroupsSitesGroupidParams object
// with the ability to set a context for a request.
func NewGetAccessgroupsSitesGroupidParamsWithContext(ctx context.Context) *GetAccessgroupsSitesGroupidParams {
	return &GetAccessgroupsSitesGroupidParams{
		Context: ctx,
	}
}

// NewGetAccessgroupsSitesGroupidParamsWithHTTPClient creates a new GetAccessgroupsSitesGroupidParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccessgroupsSitesGroupidParamsWithHTTPClient(client *http.Client) *GetAccessgroupsSitesGroupidParams {
	return &GetAccessgroupsSitesGroupidParams{
		HTTPClient: client,
	}
}

/* GetAccessgroupsSitesGroupidParams contains all the parameters to send to the API endpoint
   for the get accessgroups sites groupid operation.

   Typically these are written to a http.Request.
*/
type GetAccessgroupsSitesGroupidParams struct {

	// GroupID.
	GroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get accessgroups sites groupid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccessgroupsSitesGroupidParams) WithDefaults() *GetAccessgroupsSitesGroupidParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get accessgroups sites groupid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccessgroupsSitesGroupidParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get accessgroups sites groupid params
func (o *GetAccessgroupsSitesGroupidParams) WithTimeout(timeout time.Duration) *GetAccessgroupsSitesGroupidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get accessgroups sites groupid params
func (o *GetAccessgroupsSitesGroupidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get accessgroups sites groupid params
func (o *GetAccessgroupsSitesGroupidParams) WithContext(ctx context.Context) *GetAccessgroupsSitesGroupidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get accessgroups sites groupid params
func (o *GetAccessgroupsSitesGroupidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get accessgroups sites groupid params
func (o *GetAccessgroupsSitesGroupidParams) WithHTTPClient(client *http.Client) *GetAccessgroupsSitesGroupidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get accessgroups sites groupid params
func (o *GetAccessgroupsSitesGroupidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the get accessgroups sites groupid params
func (o *GetAccessgroupsSitesGroupidParams) WithGroupID(groupID string) *GetAccessgroupsSitesGroupidParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get accessgroups sites groupid params
func (o *GetAccessgroupsSitesGroupidParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccessgroupsSitesGroupidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
