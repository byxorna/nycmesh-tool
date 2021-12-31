// Code generated by go-swagger; DO NOT EDIT.

package logs

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

// NewGetLogsParams creates a new GetLogsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLogsParams() *GetLogsParams {
	return &GetLogsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLogsParamsWithTimeout creates a new GetLogsParams object
// with the ability to set a timeout on a request.
func NewGetLogsParamsWithTimeout(timeout time.Duration) *GetLogsParams {
	return &GetLogsParams{
		timeout: timeout,
	}
}

// NewGetLogsParamsWithContext creates a new GetLogsParams object
// with the ability to set a context for a request.
func NewGetLogsParamsWithContext(ctx context.Context) *GetLogsParams {
	return &GetLogsParams{
		Context: ctx,
	}
}

// NewGetLogsParamsWithHTTPClient creates a new GetLogsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLogsParamsWithHTTPClient(client *http.Client) *GetLogsParams {
	return &GetLogsParams{
		HTTPClient: client,
	}
}

/* GetLogsParams contains all the parameters to send to the API endpoint
   for the get logs operation.

   Typically these are written to a http.Request.
*/
type GetLogsParams struct {

	// Count.
	Count float64

	// DeviceID.
	DeviceID []string

	// Level.
	Level *string

	// Page.
	Page float64

	/* Period.

	   Unix timestamp in milliseconds.
	*/
	Period *float64

	/* Query.

	   Search pattern.
	*/
	Query *string

	// SiteID.
	SiteID *string

	// Tag.
	Tag *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogsParams) WithDefaults() *GetLogsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get logs params
func (o *GetLogsParams) WithTimeout(timeout time.Duration) *GetLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get logs params
func (o *GetLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get logs params
func (o *GetLogsParams) WithContext(ctx context.Context) *GetLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get logs params
func (o *GetLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get logs params
func (o *GetLogsParams) WithHTTPClient(client *http.Client) *GetLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get logs params
func (o *GetLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get logs params
func (o *GetLogsParams) WithCount(count float64) *GetLogsParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get logs params
func (o *GetLogsParams) SetCount(count float64) {
	o.Count = count
}

// WithDeviceID adds the deviceID to the get logs params
func (o *GetLogsParams) WithDeviceID(deviceID []string) *GetLogsParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get logs params
func (o *GetLogsParams) SetDeviceID(deviceID []string) {
	o.DeviceID = deviceID
}

// WithLevel adds the level to the get logs params
func (o *GetLogsParams) WithLevel(level *string) *GetLogsParams {
	o.SetLevel(level)
	return o
}

// SetLevel adds the level to the get logs params
func (o *GetLogsParams) SetLevel(level *string) {
	o.Level = level
}

// WithPage adds the page to the get logs params
func (o *GetLogsParams) WithPage(page float64) *GetLogsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get logs params
func (o *GetLogsParams) SetPage(page float64) {
	o.Page = page
}

// WithPeriod adds the period to the get logs params
func (o *GetLogsParams) WithPeriod(period *float64) *GetLogsParams {
	o.SetPeriod(period)
	return o
}

// SetPeriod adds the period to the get logs params
func (o *GetLogsParams) SetPeriod(period *float64) {
	o.Period = period
}

// WithQuery adds the query to the get logs params
func (o *GetLogsParams) WithQuery(query *string) *GetLogsParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get logs params
func (o *GetLogsParams) SetQuery(query *string) {
	o.Query = query
}

// WithSiteID adds the siteID to the get logs params
func (o *GetLogsParams) WithSiteID(siteID *string) *GetLogsParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get logs params
func (o *GetLogsParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithTag adds the tag to the get logs params
func (o *GetLogsParams) WithTag(tag *string) *GetLogsParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the get logs params
func (o *GetLogsParams) SetTag(tag *string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *GetLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param count
	qrCount := o.Count
	qCount := swag.FormatFloat64(qrCount)
	if qCount != "" {

		if err := r.SetQueryParam("count", qCount); err != nil {
			return err
		}
	}

	if o.DeviceID != nil {

		// binding items for deviceId
		joinedDeviceID := o.bindParamDeviceID(reg)

		// query array param deviceId
		if err := r.SetQueryParam("deviceId", joinedDeviceID...); err != nil {
			return err
		}
	}

	if o.Level != nil {

		// query param level
		var qrLevel string

		if o.Level != nil {
			qrLevel = *o.Level
		}
		qLevel := qrLevel
		if qLevel != "" {

			if err := r.SetQueryParam("level", qLevel); err != nil {
				return err
			}
		}
	}

	// query param page
	qrPage := o.Page
	qPage := swag.FormatFloat64(qrPage)
	if qPage != "" {

		if err := r.SetQueryParam("page", qPage); err != nil {
			return err
		}
	}

	if o.Period != nil {

		// query param period
		var qrPeriod float64

		if o.Period != nil {
			qrPeriod = *o.Period
		}
		qPeriod := swag.FormatFloat64(qrPeriod)
		if qPeriod != "" {

			if err := r.SetQueryParam("period", qPeriod); err != nil {
				return err
			}
		}
	}

	if o.Query != nil {

		// query param query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}
	}

	if o.SiteID != nil {

		// query param siteId
		var qrSiteID string

		if o.SiteID != nil {
			qrSiteID = *o.SiteID
		}
		qSiteID := qrSiteID
		if qSiteID != "" {

			if err := r.SetQueryParam("siteId", qSiteID); err != nil {
				return err
			}
		}
	}

	if o.Tag != nil {

		// query param tag
		var qrTag string

		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {

			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetLogs binds the parameter deviceId
func (o *GetLogsParams) bindParamDeviceID(formats strfmt.Registry) []string {
	deviceIDIR := o.DeviceID

	var deviceIDIC []string
	for _, deviceIDIIR := range deviceIDIR { // explode []string

		deviceIDIIV := deviceIDIIR // string as string
		deviceIDIC = append(deviceIDIC, deviceIDIIV)
	}

	// items.CollectionFormat: "multi"
	deviceIDIS := swag.JoinByFormat(deviceIDIC, "multi")

	return deviceIDIS
}
