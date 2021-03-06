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

// NewPostDevicesIDVlansParams creates a new PostDevicesIDVlansParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDevicesIDVlansParams() *PostDevicesIDVlansParams {
	return &PostDevicesIDVlansParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesIDVlansParamsWithTimeout creates a new PostDevicesIDVlansParams object
// with the ability to set a timeout on a request.
func NewPostDevicesIDVlansParamsWithTimeout(timeout time.Duration) *PostDevicesIDVlansParams {
	return &PostDevicesIDVlansParams{
		timeout: timeout,
	}
}

// NewPostDevicesIDVlansParamsWithContext creates a new PostDevicesIDVlansParams object
// with the ability to set a context for a request.
func NewPostDevicesIDVlansParamsWithContext(ctx context.Context) *PostDevicesIDVlansParams {
	return &PostDevicesIDVlansParams{
		Context: ctx,
	}
}

// NewPostDevicesIDVlansParamsWithHTTPClient creates a new PostDevicesIDVlansParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDevicesIDVlansParamsWithHTTPClient(client *http.Client) *PostDevicesIDVlansParams {
	return &PostDevicesIDVlansParams{
		HTTPClient: client,
	}
}

/* PostDevicesIDVlansParams contains all the parameters to send to the API endpoint
   for the post devices Id vlans operation.

   Typically these are written to a http.Request.
*/
type PostDevicesIDVlansParams struct {

	// Body.
	Body *models.Model30

	// ID.
	ID string

	// IsForceApply.
	IsForceApply *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post devices Id vlans params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesIDVlansParams) WithDefaults() *PostDevicesIDVlansParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post devices Id vlans params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesIDVlansParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post devices Id vlans params
func (o *PostDevicesIDVlansParams) WithTimeout(timeout time.Duration) *PostDevicesIDVlansParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices Id vlans params
func (o *PostDevicesIDVlansParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices Id vlans params
func (o *PostDevicesIDVlansParams) WithContext(ctx context.Context) *PostDevicesIDVlansParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices Id vlans params
func (o *PostDevicesIDVlansParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices Id vlans params
func (o *PostDevicesIDVlansParams) WithHTTPClient(client *http.Client) *PostDevicesIDVlansParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices Id vlans params
func (o *PostDevicesIDVlansParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post devices Id vlans params
func (o *PostDevicesIDVlansParams) WithBody(body *models.Model30) *PostDevicesIDVlansParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post devices Id vlans params
func (o *PostDevicesIDVlansParams) SetBody(body *models.Model30) {
	o.Body = body
}

// WithID adds the id to the post devices Id vlans params
func (o *PostDevicesIDVlansParams) WithID(id string) *PostDevicesIDVlansParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post devices Id vlans params
func (o *PostDevicesIDVlansParams) SetID(id string) {
	o.ID = id
}

// WithIsForceApply adds the isForceApply to the post devices Id vlans params
func (o *PostDevicesIDVlansParams) WithIsForceApply(isForceApply *string) *PostDevicesIDVlansParams {
	o.SetIsForceApply(isForceApply)
	return o
}

// SetIsForceApply adds the isForceApply to the post devices Id vlans params
func (o *PostDevicesIDVlansParams) SetIsForceApply(isForceApply *string) {
	o.IsForceApply = isForceApply
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesIDVlansParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IsForceApply != nil {

		// query param isForceApply
		var qrIsForceApply string

		if o.IsForceApply != nil {
			qrIsForceApply = *o.IsForceApply
		}
		qIsForceApply := qrIsForceApply
		if qIsForceApply != "" {

			if err := r.SetQueryParam("isForceApply", qIsForceApply); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
