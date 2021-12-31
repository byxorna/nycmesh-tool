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

// NewPostDevicesIDRouterOspfAreasParams creates a new PostDevicesIDRouterOspfAreasParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDevicesIDRouterOspfAreasParams() *PostDevicesIDRouterOspfAreasParams {
	return &PostDevicesIDRouterOspfAreasParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesIDRouterOspfAreasParamsWithTimeout creates a new PostDevicesIDRouterOspfAreasParams object
// with the ability to set a timeout on a request.
func NewPostDevicesIDRouterOspfAreasParamsWithTimeout(timeout time.Duration) *PostDevicesIDRouterOspfAreasParams {
	return &PostDevicesIDRouterOspfAreasParams{
		timeout: timeout,
	}
}

// NewPostDevicesIDRouterOspfAreasParamsWithContext creates a new PostDevicesIDRouterOspfAreasParams object
// with the ability to set a context for a request.
func NewPostDevicesIDRouterOspfAreasParamsWithContext(ctx context.Context) *PostDevicesIDRouterOspfAreasParams {
	return &PostDevicesIDRouterOspfAreasParams{
		Context: ctx,
	}
}

// NewPostDevicesIDRouterOspfAreasParamsWithHTTPClient creates a new PostDevicesIDRouterOspfAreasParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDevicesIDRouterOspfAreasParamsWithHTTPClient(client *http.Client) *PostDevicesIDRouterOspfAreasParams {
	return &PostDevicesIDRouterOspfAreasParams{
		HTTPClient: client,
	}
}

/* PostDevicesIDRouterOspfAreasParams contains all the parameters to send to the API endpoint
   for the post devices Id router ospf areas operation.

   Typically these are written to a http.Request.
*/
type PostDevicesIDRouterOspfAreasParams struct {

	// Body.
	Body *models.OspfArea

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post devices Id router ospf areas params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesIDRouterOspfAreasParams) WithDefaults() *PostDevicesIDRouterOspfAreasParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post devices Id router ospf areas params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesIDRouterOspfAreasParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post devices Id router ospf areas params
func (o *PostDevicesIDRouterOspfAreasParams) WithTimeout(timeout time.Duration) *PostDevicesIDRouterOspfAreasParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices Id router ospf areas params
func (o *PostDevicesIDRouterOspfAreasParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices Id router ospf areas params
func (o *PostDevicesIDRouterOspfAreasParams) WithContext(ctx context.Context) *PostDevicesIDRouterOspfAreasParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices Id router ospf areas params
func (o *PostDevicesIDRouterOspfAreasParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices Id router ospf areas params
func (o *PostDevicesIDRouterOspfAreasParams) WithHTTPClient(client *http.Client) *PostDevicesIDRouterOspfAreasParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices Id router ospf areas params
func (o *PostDevicesIDRouterOspfAreasParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post devices Id router ospf areas params
func (o *PostDevicesIDRouterOspfAreasParams) WithBody(body *models.OspfArea) *PostDevicesIDRouterOspfAreasParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post devices Id router ospf areas params
func (o *PostDevicesIDRouterOspfAreasParams) SetBody(body *models.OspfArea) {
	o.Body = body
}

// WithID adds the id to the post devices Id router ospf areas params
func (o *PostDevicesIDRouterOspfAreasParams) WithID(id string) *PostDevicesIDRouterOspfAreasParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post devices Id router ospf areas params
func (o *PostDevicesIDRouterOspfAreasParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesIDRouterOspfAreasParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
