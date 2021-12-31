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

// NewGetAirlinkProxyAirlinkbeParams creates a new GetAirlinkProxyAirlinkbeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAirlinkProxyAirlinkbeParams() *GetAirlinkProxyAirlinkbeParams {
	return &GetAirlinkProxyAirlinkbeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAirlinkProxyAirlinkbeParamsWithTimeout creates a new GetAirlinkProxyAirlinkbeParams object
// with the ability to set a timeout on a request.
func NewGetAirlinkProxyAirlinkbeParamsWithTimeout(timeout time.Duration) *GetAirlinkProxyAirlinkbeParams {
	return &GetAirlinkProxyAirlinkbeParams{
		timeout: timeout,
	}
}

// NewGetAirlinkProxyAirlinkbeParamsWithContext creates a new GetAirlinkProxyAirlinkbeParams object
// with the ability to set a context for a request.
func NewGetAirlinkProxyAirlinkbeParamsWithContext(ctx context.Context) *GetAirlinkProxyAirlinkbeParams {
	return &GetAirlinkProxyAirlinkbeParams{
		Context: ctx,
	}
}

// NewGetAirlinkProxyAirlinkbeParamsWithHTTPClient creates a new GetAirlinkProxyAirlinkbeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAirlinkProxyAirlinkbeParamsWithHTTPClient(client *http.Client) *GetAirlinkProxyAirlinkbeParams {
	return &GetAirlinkProxyAirlinkbeParams{
		HTTPClient: client,
	}
}

/* GetAirlinkProxyAirlinkbeParams contains all the parameters to send to the API endpoint
   for the get airlink proxy airlinkbe operation.

   Typically these are written to a http.Request.
*/
type GetAirlinkProxyAirlinkbeParams struct {

	/* Lat1.

	   Latitude of the first point.
	*/
	Lat1 float64

	/* Lat2.

	   Latitude of the second point.
	*/
	Lat2 float64

	/* Lon1.

	   Longitude of the first point.
	*/
	Lon1 float64

	/* Lon2.

	   Longitude of the second point.
	*/
	Lon2 float64

	/* Samples.

	   Number of elevations to get.

	   Default: 256
	*/
	Samples *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get airlink proxy airlinkbe params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAirlinkProxyAirlinkbeParams) WithDefaults() *GetAirlinkProxyAirlinkbeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get airlink proxy airlinkbe params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAirlinkProxyAirlinkbeParams) SetDefaults() {
	var (
		samplesDefault = float64(256)
	)

	val := GetAirlinkProxyAirlinkbeParams{
		Samples: &samplesDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get airlink proxy airlinkbe params
func (o *GetAirlinkProxyAirlinkbeParams) WithTimeout(timeout time.Duration) *GetAirlinkProxyAirlinkbeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get airlink proxy airlinkbe params
func (o *GetAirlinkProxyAirlinkbeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get airlink proxy airlinkbe params
func (o *GetAirlinkProxyAirlinkbeParams) WithContext(ctx context.Context) *GetAirlinkProxyAirlinkbeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get airlink proxy airlinkbe params
func (o *GetAirlinkProxyAirlinkbeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get airlink proxy airlinkbe params
func (o *GetAirlinkProxyAirlinkbeParams) WithHTTPClient(client *http.Client) *GetAirlinkProxyAirlinkbeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get airlink proxy airlinkbe params
func (o *GetAirlinkProxyAirlinkbeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLat1 adds the lat1 to the get airlink proxy airlinkbe params
func (o *GetAirlinkProxyAirlinkbeParams) WithLat1(lat1 float64) *GetAirlinkProxyAirlinkbeParams {
	o.SetLat1(lat1)
	return o
}

// SetLat1 adds the lat1 to the get airlink proxy airlinkbe params
func (o *GetAirlinkProxyAirlinkbeParams) SetLat1(lat1 float64) {
	o.Lat1 = lat1
}

// WithLat2 adds the lat2 to the get airlink proxy airlinkbe params
func (o *GetAirlinkProxyAirlinkbeParams) WithLat2(lat2 float64) *GetAirlinkProxyAirlinkbeParams {
	o.SetLat2(lat2)
	return o
}

// SetLat2 adds the lat2 to the get airlink proxy airlinkbe params
func (o *GetAirlinkProxyAirlinkbeParams) SetLat2(lat2 float64) {
	o.Lat2 = lat2
}

// WithLon1 adds the lon1 to the get airlink proxy airlinkbe params
func (o *GetAirlinkProxyAirlinkbeParams) WithLon1(lon1 float64) *GetAirlinkProxyAirlinkbeParams {
	o.SetLon1(lon1)
	return o
}

// SetLon1 adds the lon1 to the get airlink proxy airlinkbe params
func (o *GetAirlinkProxyAirlinkbeParams) SetLon1(lon1 float64) {
	o.Lon1 = lon1
}

// WithLon2 adds the lon2 to the get airlink proxy airlinkbe params
func (o *GetAirlinkProxyAirlinkbeParams) WithLon2(lon2 float64) *GetAirlinkProxyAirlinkbeParams {
	o.SetLon2(lon2)
	return o
}

// SetLon2 adds the lon2 to the get airlink proxy airlinkbe params
func (o *GetAirlinkProxyAirlinkbeParams) SetLon2(lon2 float64) {
	o.Lon2 = lon2
}

// WithSamples adds the samples to the get airlink proxy airlinkbe params
func (o *GetAirlinkProxyAirlinkbeParams) WithSamples(samples *float64) *GetAirlinkProxyAirlinkbeParams {
	o.SetSamples(samples)
	return o
}

// SetSamples adds the samples to the get airlink proxy airlinkbe params
func (o *GetAirlinkProxyAirlinkbeParams) SetSamples(samples *float64) {
	o.Samples = samples
}

// WriteToRequest writes these params to a swagger request
func (o *GetAirlinkProxyAirlinkbeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param lat1
	qrLat1 := o.Lat1
	qLat1 := swag.FormatFloat64(qrLat1)
	if qLat1 != "" {

		if err := r.SetQueryParam("lat1", qLat1); err != nil {
			return err
		}
	}

	// query param lat2
	qrLat2 := o.Lat2
	qLat2 := swag.FormatFloat64(qrLat2)
	if qLat2 != "" {

		if err := r.SetQueryParam("lat2", qLat2); err != nil {
			return err
		}
	}

	// query param lon1
	qrLon1 := o.Lon1
	qLon1 := swag.FormatFloat64(qrLon1)
	if qLon1 != "" {

		if err := r.SetQueryParam("lon1", qLon1); err != nil {
			return err
		}
	}

	// query param lon2
	qrLon2 := o.Lon2
	qLon2 := swag.FormatFloat64(qrLon2)
	if qLon2 != "" {

		if err := r.SetQueryParam("lon2", qLon2); err != nil {
			return err
		}
	}

	if o.Samples != nil {

		// query param samples
		var qrSamples float64

		if o.Samples != nil {
			qrSamples = *o.Samples
		}
		qSamples := swag.FormatFloat64(qrSamples)
		if qSamples != "" {

			if err := r.SetQueryParam("samples", qSamples); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}