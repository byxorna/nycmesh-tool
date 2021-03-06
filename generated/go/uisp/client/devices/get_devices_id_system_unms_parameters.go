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

// NewGetDevicesIDSystemUnmsParams creates a new GetDevicesIDSystemUnmsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicesIDSystemUnmsParams() *GetDevicesIDSystemUnmsParams {
	return &GetDevicesIDSystemUnmsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesIDSystemUnmsParamsWithTimeout creates a new GetDevicesIDSystemUnmsParams object
// with the ability to set a timeout on a request.
func NewGetDevicesIDSystemUnmsParamsWithTimeout(timeout time.Duration) *GetDevicesIDSystemUnmsParams {
	return &GetDevicesIDSystemUnmsParams{
		timeout: timeout,
	}
}

// NewGetDevicesIDSystemUnmsParamsWithContext creates a new GetDevicesIDSystemUnmsParams object
// with the ability to set a context for a request.
func NewGetDevicesIDSystemUnmsParamsWithContext(ctx context.Context) *GetDevicesIDSystemUnmsParams {
	return &GetDevicesIDSystemUnmsParams{
		Context: ctx,
	}
}

// NewGetDevicesIDSystemUnmsParamsWithHTTPClient creates a new GetDevicesIDSystemUnmsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicesIDSystemUnmsParamsWithHTTPClient(client *http.Client) *GetDevicesIDSystemUnmsParams {
	return &GetDevicesIDSystemUnmsParams{
		HTTPClient: client,
	}
}

/* GetDevicesIDSystemUnmsParams contains all the parameters to send to the API endpoint
   for the get devices Id system unms operation.

   Typically these are written to a http.Request.
*/
type GetDevicesIDSystemUnmsParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get devices Id system unms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesIDSystemUnmsParams) WithDefaults() *GetDevicesIDSystemUnmsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get devices Id system unms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesIDSystemUnmsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get devices Id system unms params
func (o *GetDevicesIDSystemUnmsParams) WithTimeout(timeout time.Duration) *GetDevicesIDSystemUnmsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices Id system unms params
func (o *GetDevicesIDSystemUnmsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices Id system unms params
func (o *GetDevicesIDSystemUnmsParams) WithContext(ctx context.Context) *GetDevicesIDSystemUnmsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices Id system unms params
func (o *GetDevicesIDSystemUnmsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices Id system unms params
func (o *GetDevicesIDSystemUnmsParams) WithHTTPClient(client *http.Client) *GetDevicesIDSystemUnmsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices Id system unms params
func (o *GetDevicesIDSystemUnmsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get devices Id system unms params
func (o *GetDevicesIDSystemUnmsParams) WithID(id string) *GetDevicesIDSystemUnmsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get devices Id system unms params
func (o *GetDevicesIDSystemUnmsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesIDSystemUnmsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
