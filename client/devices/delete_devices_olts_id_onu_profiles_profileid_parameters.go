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

// NewDeleteDevicesOltsIDOnuProfilesProfileidParams creates a new DeleteDevicesOltsIDOnuProfilesProfileidParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDevicesOltsIDOnuProfilesProfileidParams() *DeleteDevicesOltsIDOnuProfilesProfileidParams {
	return &DeleteDevicesOltsIDOnuProfilesProfileidParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDevicesOltsIDOnuProfilesProfileidParamsWithTimeout creates a new DeleteDevicesOltsIDOnuProfilesProfileidParams object
// with the ability to set a timeout on a request.
func NewDeleteDevicesOltsIDOnuProfilesProfileidParamsWithTimeout(timeout time.Duration) *DeleteDevicesOltsIDOnuProfilesProfileidParams {
	return &DeleteDevicesOltsIDOnuProfilesProfileidParams{
		timeout: timeout,
	}
}

// NewDeleteDevicesOltsIDOnuProfilesProfileidParamsWithContext creates a new DeleteDevicesOltsIDOnuProfilesProfileidParams object
// with the ability to set a context for a request.
func NewDeleteDevicesOltsIDOnuProfilesProfileidParamsWithContext(ctx context.Context) *DeleteDevicesOltsIDOnuProfilesProfileidParams {
	return &DeleteDevicesOltsIDOnuProfilesProfileidParams{
		Context: ctx,
	}
}

// NewDeleteDevicesOltsIDOnuProfilesProfileidParamsWithHTTPClient creates a new DeleteDevicesOltsIDOnuProfilesProfileidParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDevicesOltsIDOnuProfilesProfileidParamsWithHTTPClient(client *http.Client) *DeleteDevicesOltsIDOnuProfilesProfileidParams {
	return &DeleteDevicesOltsIDOnuProfilesProfileidParams{
		HTTPClient: client,
	}
}

/* DeleteDevicesOltsIDOnuProfilesProfileidParams contains all the parameters to send to the API endpoint
   for the delete devices olts Id onu profiles profileid operation.

   Typically these are written to a http.Request.
*/
type DeleteDevicesOltsIDOnuProfilesProfileidParams struct {

	// ID.
	ID string

	// ProfileID.
	ProfileID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete devices olts Id onu profiles profileid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDevicesOltsIDOnuProfilesProfileidParams) WithDefaults() *DeleteDevicesOltsIDOnuProfilesProfileidParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete devices olts Id onu profiles profileid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDevicesOltsIDOnuProfilesProfileidParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete devices olts Id onu profiles profileid params
func (o *DeleteDevicesOltsIDOnuProfilesProfileidParams) WithTimeout(timeout time.Duration) *DeleteDevicesOltsIDOnuProfilesProfileidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete devices olts Id onu profiles profileid params
func (o *DeleteDevicesOltsIDOnuProfilesProfileidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete devices olts Id onu profiles profileid params
func (o *DeleteDevicesOltsIDOnuProfilesProfileidParams) WithContext(ctx context.Context) *DeleteDevicesOltsIDOnuProfilesProfileidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete devices olts Id onu profiles profileid params
func (o *DeleteDevicesOltsIDOnuProfilesProfileidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete devices olts Id onu profiles profileid params
func (o *DeleteDevicesOltsIDOnuProfilesProfileidParams) WithHTTPClient(client *http.Client) *DeleteDevicesOltsIDOnuProfilesProfileidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete devices olts Id onu profiles profileid params
func (o *DeleteDevicesOltsIDOnuProfilesProfileidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete devices olts Id onu profiles profileid params
func (o *DeleteDevicesOltsIDOnuProfilesProfileidParams) WithID(id string) *DeleteDevicesOltsIDOnuProfilesProfileidParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete devices olts Id onu profiles profileid params
func (o *DeleteDevicesOltsIDOnuProfilesProfileidParams) SetID(id string) {
	o.ID = id
}

// WithProfileID adds the profileID to the delete devices olts Id onu profiles profileid params
func (o *DeleteDevicesOltsIDOnuProfilesProfileidParams) WithProfileID(profileID string) *DeleteDevicesOltsIDOnuProfilesProfileidParams {
	o.SetProfileID(profileID)
	return o
}

// SetProfileID adds the profileId to the delete devices olts Id onu profiles profileid params
func (o *DeleteDevicesOltsIDOnuProfilesProfileidParams) SetProfileID(profileID string) {
	o.ProfileID = profileID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDevicesOltsIDOnuProfilesProfileidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param profileId
	if err := r.SetPathParam("profileId", o.ProfileID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
