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
)

// NewGetDevicesIDNetflowParams creates a new GetDevicesIDNetflowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicesIDNetflowParams() *GetDevicesIDNetflowParams {
	return &GetDevicesIDNetflowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesIDNetflowParamsWithTimeout creates a new GetDevicesIDNetflowParams object
// with the ability to set a timeout on a request.
func NewGetDevicesIDNetflowParamsWithTimeout(timeout time.Duration) *GetDevicesIDNetflowParams {
	return &GetDevicesIDNetflowParams{
		timeout: timeout,
	}
}

// NewGetDevicesIDNetflowParamsWithContext creates a new GetDevicesIDNetflowParams object
// with the ability to set a context for a request.
func NewGetDevicesIDNetflowParamsWithContext(ctx context.Context) *GetDevicesIDNetflowParams {
	return &GetDevicesIDNetflowParams{
		Context: ctx,
	}
}

// NewGetDevicesIDNetflowParamsWithHTTPClient creates a new GetDevicesIDNetflowParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicesIDNetflowParamsWithHTTPClient(client *http.Client) *GetDevicesIDNetflowParams {
	return &GetDevicesIDNetflowParams{
		HTTPClient: client,
	}
}

/* GetDevicesIDNetflowParams contains all the parameters to send to the API endpoint
   for the get devices Id netflow operation.

   Typically these are written to a http.Request.
*/
type GetDevicesIDNetflowParams struct {

	// ID.
	ID string

	// InterfaceID.
	InterfaceID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get devices Id netflow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesIDNetflowParams) WithDefaults() *GetDevicesIDNetflowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get devices Id netflow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesIDNetflowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get devices Id netflow params
func (o *GetDevicesIDNetflowParams) WithTimeout(timeout time.Duration) *GetDevicesIDNetflowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices Id netflow params
func (o *GetDevicesIDNetflowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices Id netflow params
func (o *GetDevicesIDNetflowParams) WithContext(ctx context.Context) *GetDevicesIDNetflowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices Id netflow params
func (o *GetDevicesIDNetflowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices Id netflow params
func (o *GetDevicesIDNetflowParams) WithHTTPClient(client *http.Client) *GetDevicesIDNetflowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices Id netflow params
func (o *GetDevicesIDNetflowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get devices Id netflow params
func (o *GetDevicesIDNetflowParams) WithID(id string) *GetDevicesIDNetflowParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get devices Id netflow params
func (o *GetDevicesIDNetflowParams) SetID(id string) {
	o.ID = id
}

// WithInterfaceID adds the interfaceID to the get devices Id netflow params
func (o *GetDevicesIDNetflowParams) WithInterfaceID(interfaceID *string) *GetDevicesIDNetflowParams {
	o.SetInterfaceID(interfaceID)
	return o
}

// SetInterfaceID adds the interfaceId to the get devices Id netflow params
func (o *GetDevicesIDNetflowParams) SetInterfaceID(interfaceID *string) {
	o.InterfaceID = interfaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesIDNetflowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.InterfaceID != nil {

		// query param interfaceId
		var qrInterfaceID string

		if o.InterfaceID != nil {
			qrInterfaceID = *o.InterfaceID
		}
		qInterfaceID := qrInterfaceID
		if qInterfaceID != "" {

			if err := r.SetQueryParam("interfaceId", qInterfaceID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
