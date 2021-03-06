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

// NewPostDevicesUisprsIDRouterOspfAreasParams creates a new PostDevicesUisprsIDRouterOspfAreasParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDevicesUisprsIDRouterOspfAreasParams() *PostDevicesUisprsIDRouterOspfAreasParams {
	return &PostDevicesUisprsIDRouterOspfAreasParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesUisprsIDRouterOspfAreasParamsWithTimeout creates a new PostDevicesUisprsIDRouterOspfAreasParams object
// with the ability to set a timeout on a request.
func NewPostDevicesUisprsIDRouterOspfAreasParamsWithTimeout(timeout time.Duration) *PostDevicesUisprsIDRouterOspfAreasParams {
	return &PostDevicesUisprsIDRouterOspfAreasParams{
		timeout: timeout,
	}
}

// NewPostDevicesUisprsIDRouterOspfAreasParamsWithContext creates a new PostDevicesUisprsIDRouterOspfAreasParams object
// with the ability to set a context for a request.
func NewPostDevicesUisprsIDRouterOspfAreasParamsWithContext(ctx context.Context) *PostDevicesUisprsIDRouterOspfAreasParams {
	return &PostDevicesUisprsIDRouterOspfAreasParams{
		Context: ctx,
	}
}

// NewPostDevicesUisprsIDRouterOspfAreasParamsWithHTTPClient creates a new PostDevicesUisprsIDRouterOspfAreasParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDevicesUisprsIDRouterOspfAreasParamsWithHTTPClient(client *http.Client) *PostDevicesUisprsIDRouterOspfAreasParams {
	return &PostDevicesUisprsIDRouterOspfAreasParams{
		HTTPClient: client,
	}
}

/* PostDevicesUisprsIDRouterOspfAreasParams contains all the parameters to send to the API endpoint
   for the post devices uisprs Id router ospf areas operation.

   Typically these are written to a http.Request.
*/
type PostDevicesUisprsIDRouterOspfAreasParams struct {

	// Body.
	Body *models.OspfArea

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post devices uisprs Id router ospf areas params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesUisprsIDRouterOspfAreasParams) WithDefaults() *PostDevicesUisprsIDRouterOspfAreasParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post devices uisprs Id router ospf areas params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesUisprsIDRouterOspfAreasParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post devices uisprs Id router ospf areas params
func (o *PostDevicesUisprsIDRouterOspfAreasParams) WithTimeout(timeout time.Duration) *PostDevicesUisprsIDRouterOspfAreasParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices uisprs Id router ospf areas params
func (o *PostDevicesUisprsIDRouterOspfAreasParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices uisprs Id router ospf areas params
func (o *PostDevicesUisprsIDRouterOspfAreasParams) WithContext(ctx context.Context) *PostDevicesUisprsIDRouterOspfAreasParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices uisprs Id router ospf areas params
func (o *PostDevicesUisprsIDRouterOspfAreasParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices uisprs Id router ospf areas params
func (o *PostDevicesUisprsIDRouterOspfAreasParams) WithHTTPClient(client *http.Client) *PostDevicesUisprsIDRouterOspfAreasParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices uisprs Id router ospf areas params
func (o *PostDevicesUisprsIDRouterOspfAreasParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post devices uisprs Id router ospf areas params
func (o *PostDevicesUisprsIDRouterOspfAreasParams) WithBody(body *models.OspfArea) *PostDevicesUisprsIDRouterOspfAreasParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post devices uisprs Id router ospf areas params
func (o *PostDevicesUisprsIDRouterOspfAreasParams) SetBody(body *models.OspfArea) {
	o.Body = body
}

// WithID adds the id to the post devices uisprs Id router ospf areas params
func (o *PostDevicesUisprsIDRouterOspfAreasParams) WithID(id string) *PostDevicesUisprsIDRouterOspfAreasParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post devices uisprs Id router ospf areas params
func (o *PostDevicesUisprsIDRouterOspfAreasParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesUisprsIDRouterOspfAreasParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
