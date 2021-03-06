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

// NewGetAirlinkProxyElevationParams creates a new GetAirlinkProxyElevationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAirlinkProxyElevationParams() *GetAirlinkProxyElevationParams {
	return &GetAirlinkProxyElevationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAirlinkProxyElevationParamsWithTimeout creates a new GetAirlinkProxyElevationParams object
// with the ability to set a timeout on a request.
func NewGetAirlinkProxyElevationParamsWithTimeout(timeout time.Duration) *GetAirlinkProxyElevationParams {
	return &GetAirlinkProxyElevationParams{
		timeout: timeout,
	}
}

// NewGetAirlinkProxyElevationParamsWithContext creates a new GetAirlinkProxyElevationParams object
// with the ability to set a context for a request.
func NewGetAirlinkProxyElevationParamsWithContext(ctx context.Context) *GetAirlinkProxyElevationParams {
	return &GetAirlinkProxyElevationParams{
		Context: ctx,
	}
}

// NewGetAirlinkProxyElevationParamsWithHTTPClient creates a new GetAirlinkProxyElevationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAirlinkProxyElevationParamsWithHTTPClient(client *http.Client) *GetAirlinkProxyElevationParams {
	return &GetAirlinkProxyElevationParams{
		HTTPClient: client,
	}
}

/* GetAirlinkProxyElevationParams contains all the parameters to send to the API endpoint
   for the get airlink proxy elevation operation.

   Typically these are written to a http.Request.
*/
type GetAirlinkProxyElevationParams struct {

	/* File.

	   Elevation file name.
	*/
	File string

	/* Version.

	   Elevation file version.

	   Default: 1
	*/
	Version *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get airlink proxy elevation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAirlinkProxyElevationParams) WithDefaults() *GetAirlinkProxyElevationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get airlink proxy elevation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAirlinkProxyElevationParams) SetDefaults() {
	var (
		versionDefault = float64(1)
	)

	val := GetAirlinkProxyElevationParams{
		Version: &versionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get airlink proxy elevation params
func (o *GetAirlinkProxyElevationParams) WithTimeout(timeout time.Duration) *GetAirlinkProxyElevationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get airlink proxy elevation params
func (o *GetAirlinkProxyElevationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get airlink proxy elevation params
func (o *GetAirlinkProxyElevationParams) WithContext(ctx context.Context) *GetAirlinkProxyElevationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get airlink proxy elevation params
func (o *GetAirlinkProxyElevationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get airlink proxy elevation params
func (o *GetAirlinkProxyElevationParams) WithHTTPClient(client *http.Client) *GetAirlinkProxyElevationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get airlink proxy elevation params
func (o *GetAirlinkProxyElevationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFile adds the file to the get airlink proxy elevation params
func (o *GetAirlinkProxyElevationParams) WithFile(file string) *GetAirlinkProxyElevationParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the get airlink proxy elevation params
func (o *GetAirlinkProxyElevationParams) SetFile(file string) {
	o.File = file
}

// WithVersion adds the version to the get airlink proxy elevation params
func (o *GetAirlinkProxyElevationParams) WithVersion(version *float64) *GetAirlinkProxyElevationParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get airlink proxy elevation params
func (o *GetAirlinkProxyElevationParams) SetVersion(version *float64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetAirlinkProxyElevationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param file
	qrFile := o.File
	qFile := qrFile
	if qFile != "" {

		if err := r.SetQueryParam("file", qFile); err != nil {
			return err
		}
	}

	if o.Version != nil {

		// query param version
		var qrVersion float64

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatFloat64(qrVersion)
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
