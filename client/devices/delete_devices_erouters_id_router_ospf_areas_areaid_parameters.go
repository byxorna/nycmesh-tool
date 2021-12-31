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

// NewDeleteDevicesEroutersIDRouterOspfAreasAreaidParams creates a new DeleteDevicesEroutersIDRouterOspfAreasAreaidParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDevicesEroutersIDRouterOspfAreasAreaidParams() *DeleteDevicesEroutersIDRouterOspfAreasAreaidParams {
	return &DeleteDevicesEroutersIDRouterOspfAreasAreaidParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDevicesEroutersIDRouterOspfAreasAreaidParamsWithTimeout creates a new DeleteDevicesEroutersIDRouterOspfAreasAreaidParams object
// with the ability to set a timeout on a request.
func NewDeleteDevicesEroutersIDRouterOspfAreasAreaidParamsWithTimeout(timeout time.Duration) *DeleteDevicesEroutersIDRouterOspfAreasAreaidParams {
	return &DeleteDevicesEroutersIDRouterOspfAreasAreaidParams{
		timeout: timeout,
	}
}

// NewDeleteDevicesEroutersIDRouterOspfAreasAreaidParamsWithContext creates a new DeleteDevicesEroutersIDRouterOspfAreasAreaidParams object
// with the ability to set a context for a request.
func NewDeleteDevicesEroutersIDRouterOspfAreasAreaidParamsWithContext(ctx context.Context) *DeleteDevicesEroutersIDRouterOspfAreasAreaidParams {
	return &DeleteDevicesEroutersIDRouterOspfAreasAreaidParams{
		Context: ctx,
	}
}

// NewDeleteDevicesEroutersIDRouterOspfAreasAreaidParamsWithHTTPClient creates a new DeleteDevicesEroutersIDRouterOspfAreasAreaidParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDevicesEroutersIDRouterOspfAreasAreaidParamsWithHTTPClient(client *http.Client) *DeleteDevicesEroutersIDRouterOspfAreasAreaidParams {
	return &DeleteDevicesEroutersIDRouterOspfAreasAreaidParams{
		HTTPClient: client,
	}
}

/* DeleteDevicesEroutersIDRouterOspfAreasAreaidParams contains all the parameters to send to the API endpoint
   for the delete devices erouters Id router ospf areas areaid operation.

   Typically these are written to a http.Request.
*/
type DeleteDevicesEroutersIDRouterOspfAreasAreaidParams struct {

	// AreaID.
	AreaID string

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete devices erouters Id router ospf areas areaid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDevicesEroutersIDRouterOspfAreasAreaidParams) WithDefaults() *DeleteDevicesEroutersIDRouterOspfAreasAreaidParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete devices erouters Id router ospf areas areaid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDevicesEroutersIDRouterOspfAreasAreaidParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete devices erouters Id router ospf areas areaid params
func (o *DeleteDevicesEroutersIDRouterOspfAreasAreaidParams) WithTimeout(timeout time.Duration) *DeleteDevicesEroutersIDRouterOspfAreasAreaidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete devices erouters Id router ospf areas areaid params
func (o *DeleteDevicesEroutersIDRouterOspfAreasAreaidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete devices erouters Id router ospf areas areaid params
func (o *DeleteDevicesEroutersIDRouterOspfAreasAreaidParams) WithContext(ctx context.Context) *DeleteDevicesEroutersIDRouterOspfAreasAreaidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete devices erouters Id router ospf areas areaid params
func (o *DeleteDevicesEroutersIDRouterOspfAreasAreaidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete devices erouters Id router ospf areas areaid params
func (o *DeleteDevicesEroutersIDRouterOspfAreasAreaidParams) WithHTTPClient(client *http.Client) *DeleteDevicesEroutersIDRouterOspfAreasAreaidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete devices erouters Id router ospf areas areaid params
func (o *DeleteDevicesEroutersIDRouterOspfAreasAreaidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAreaID adds the areaID to the delete devices erouters Id router ospf areas areaid params
func (o *DeleteDevicesEroutersIDRouterOspfAreasAreaidParams) WithAreaID(areaID string) *DeleteDevicesEroutersIDRouterOspfAreasAreaidParams {
	o.SetAreaID(areaID)
	return o
}

// SetAreaID adds the areaId to the delete devices erouters Id router ospf areas areaid params
func (o *DeleteDevicesEroutersIDRouterOspfAreasAreaidParams) SetAreaID(areaID string) {
	o.AreaID = areaID
}

// WithID adds the id to the delete devices erouters Id router ospf areas areaid params
func (o *DeleteDevicesEroutersIDRouterOspfAreasAreaidParams) WithID(id string) *DeleteDevicesEroutersIDRouterOspfAreasAreaidParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete devices erouters Id router ospf areas areaid params
func (o *DeleteDevicesEroutersIDRouterOspfAreasAreaidParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDevicesEroutersIDRouterOspfAreasAreaidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param areaId
	if err := r.SetPathParam("areaId", o.AreaID); err != nil {
		return err
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
