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

// NewPostDevicesOnusIDUnblockParams creates a new PostDevicesOnusIDUnblockParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDevicesOnusIDUnblockParams() *PostDevicesOnusIDUnblockParams {
	return &PostDevicesOnusIDUnblockParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDevicesOnusIDUnblockParamsWithTimeout creates a new PostDevicesOnusIDUnblockParams object
// with the ability to set a timeout on a request.
func NewPostDevicesOnusIDUnblockParamsWithTimeout(timeout time.Duration) *PostDevicesOnusIDUnblockParams {
	return &PostDevicesOnusIDUnblockParams{
		timeout: timeout,
	}
}

// NewPostDevicesOnusIDUnblockParamsWithContext creates a new PostDevicesOnusIDUnblockParams object
// with the ability to set a context for a request.
func NewPostDevicesOnusIDUnblockParamsWithContext(ctx context.Context) *PostDevicesOnusIDUnblockParams {
	return &PostDevicesOnusIDUnblockParams{
		Context: ctx,
	}
}

// NewPostDevicesOnusIDUnblockParamsWithHTTPClient creates a new PostDevicesOnusIDUnblockParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDevicesOnusIDUnblockParamsWithHTTPClient(client *http.Client) *PostDevicesOnusIDUnblockParams {
	return &PostDevicesOnusIDUnblockParams{
		HTTPClient: client,
	}
}

/* PostDevicesOnusIDUnblockParams contains all the parameters to send to the API endpoint
   for the post devices onus Id unblock operation.

   Typically these are written to a http.Request.
*/
type PostDevicesOnusIDUnblockParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post devices onus Id unblock params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesOnusIDUnblockParams) WithDefaults() *PostDevicesOnusIDUnblockParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post devices onus Id unblock params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDevicesOnusIDUnblockParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post devices onus Id unblock params
func (o *PostDevicesOnusIDUnblockParams) WithTimeout(timeout time.Duration) *PostDevicesOnusIDUnblockParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post devices onus Id unblock params
func (o *PostDevicesOnusIDUnblockParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post devices onus Id unblock params
func (o *PostDevicesOnusIDUnblockParams) WithContext(ctx context.Context) *PostDevicesOnusIDUnblockParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post devices onus Id unblock params
func (o *PostDevicesOnusIDUnblockParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post devices onus Id unblock params
func (o *PostDevicesOnusIDUnblockParams) WithHTTPClient(client *http.Client) *PostDevicesOnusIDUnblockParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post devices onus Id unblock params
func (o *PostDevicesOnusIDUnblockParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post devices onus Id unblock params
func (o *PostDevicesOnusIDUnblockParams) WithID(id string) *PostDevicesOnusIDUnblockParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post devices onus Id unblock params
func (o *PostDevicesOnusIDUnblockParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostDevicesOnusIDUnblockParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
