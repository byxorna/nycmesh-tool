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
	"github.com/go-openapi/swag"
)

// NewGetDevicesAircubesIDParams creates a new GetDevicesAircubesIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicesAircubesIDParams() *GetDevicesAircubesIDParams {
	return &GetDevicesAircubesIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesAircubesIDParamsWithTimeout creates a new GetDevicesAircubesIDParams object
// with the ability to set a timeout on a request.
func NewGetDevicesAircubesIDParamsWithTimeout(timeout time.Duration) *GetDevicesAircubesIDParams {
	return &GetDevicesAircubesIDParams{
		timeout: timeout,
	}
}

// NewGetDevicesAircubesIDParamsWithContext creates a new GetDevicesAircubesIDParams object
// with the ability to set a context for a request.
func NewGetDevicesAircubesIDParamsWithContext(ctx context.Context) *GetDevicesAircubesIDParams {
	return &GetDevicesAircubesIDParams{
		Context: ctx,
	}
}

// NewGetDevicesAircubesIDParamsWithHTTPClient creates a new GetDevicesAircubesIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicesAircubesIDParamsWithHTTPClient(client *http.Client) *GetDevicesAircubesIDParams {
	return &GetDevicesAircubesIDParams{
		HTTPClient: client,
	}
}

/* GetDevicesAircubesIDParams contains all the parameters to send to the API endpoint
   for the get devices aircubes Id operation.

   Typically these are written to a http.Request.
*/
type GetDevicesAircubesIDParams struct {

	// ID.
	ID string

	// WithStations.
	WithStations *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get devices aircubes Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesAircubesIDParams) WithDefaults() *GetDevicesAircubesIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get devices aircubes Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesAircubesIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get devices aircubes Id params
func (o *GetDevicesAircubesIDParams) WithTimeout(timeout time.Duration) *GetDevicesAircubesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices aircubes Id params
func (o *GetDevicesAircubesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices aircubes Id params
func (o *GetDevicesAircubesIDParams) WithContext(ctx context.Context) *GetDevicesAircubesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices aircubes Id params
func (o *GetDevicesAircubesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices aircubes Id params
func (o *GetDevicesAircubesIDParams) WithHTTPClient(client *http.Client) *GetDevicesAircubesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices aircubes Id params
func (o *GetDevicesAircubesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get devices aircubes Id params
func (o *GetDevicesAircubesIDParams) WithID(id string) *GetDevicesAircubesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get devices aircubes Id params
func (o *GetDevicesAircubesIDParams) SetID(id string) {
	o.ID = id
}

// WithWithStations adds the withStations to the get devices aircubes Id params
func (o *GetDevicesAircubesIDParams) WithWithStations(withStations *bool) *GetDevicesAircubesIDParams {
	o.SetWithStations(withStations)
	return o
}

// SetWithStations adds the withStations to the get devices aircubes Id params
func (o *GetDevicesAircubesIDParams) SetWithStations(withStations *bool) {
	o.WithStations = withStations
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesAircubesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.WithStations != nil {

		// query param withStations
		var qrWithStations bool

		if o.WithStations != nil {
			qrWithStations = *o.WithStations
		}
		qWithStations := swag.FormatBool(qrWithStations)
		if qWithStations != "" {

			if err := r.SetQueryParam("withStations", qWithStations); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
