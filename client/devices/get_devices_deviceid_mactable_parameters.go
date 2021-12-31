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

// NewGetDevicesDeviceidMactableParams creates a new GetDevicesDeviceidMactableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicesDeviceidMactableParams() *GetDevicesDeviceidMactableParams {
	return &GetDevicesDeviceidMactableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesDeviceidMactableParamsWithTimeout creates a new GetDevicesDeviceidMactableParams object
// with the ability to set a timeout on a request.
func NewGetDevicesDeviceidMactableParamsWithTimeout(timeout time.Duration) *GetDevicesDeviceidMactableParams {
	return &GetDevicesDeviceidMactableParams{
		timeout: timeout,
	}
}

// NewGetDevicesDeviceidMactableParamsWithContext creates a new GetDevicesDeviceidMactableParams object
// with the ability to set a context for a request.
func NewGetDevicesDeviceidMactableParamsWithContext(ctx context.Context) *GetDevicesDeviceidMactableParams {
	return &GetDevicesDeviceidMactableParams{
		Context: ctx,
	}
}

// NewGetDevicesDeviceidMactableParamsWithHTTPClient creates a new GetDevicesDeviceidMactableParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicesDeviceidMactableParamsWithHTTPClient(client *http.Client) *GetDevicesDeviceidMactableParams {
	return &GetDevicesDeviceidMactableParams{
		HTTPClient: client,
	}
}

/* GetDevicesDeviceidMactableParams contains all the parameters to send to the API endpoint
   for the get devices deviceid mactable operation.

   Typically these are written to a http.Request.
*/
type GetDevicesDeviceidMactableParams struct {

	/* Count.

	   items count
	*/
	Count *int64

	// DeviceID.
	DeviceID string

	/* Page.

	   number page
	*/
	Page *int64

	// Search.
	Search *string

	/* SortBy.

	   column name
	*/
	SortBy *string

	/* SortDesc.

	   descending sort rows [z-a] or [9-0]
	*/
	SortDesc *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get devices deviceid mactable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesDeviceidMactableParams) WithDefaults() *GetDevicesDeviceidMactableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get devices deviceid mactable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesDeviceidMactableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get devices deviceid mactable params
func (o *GetDevicesDeviceidMactableParams) WithTimeout(timeout time.Duration) *GetDevicesDeviceidMactableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices deviceid mactable params
func (o *GetDevicesDeviceidMactableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices deviceid mactable params
func (o *GetDevicesDeviceidMactableParams) WithContext(ctx context.Context) *GetDevicesDeviceidMactableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices deviceid mactable params
func (o *GetDevicesDeviceidMactableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices deviceid mactable params
func (o *GetDevicesDeviceidMactableParams) WithHTTPClient(client *http.Client) *GetDevicesDeviceidMactableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices deviceid mactable params
func (o *GetDevicesDeviceidMactableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get devices deviceid mactable params
func (o *GetDevicesDeviceidMactableParams) WithCount(count *int64) *GetDevicesDeviceidMactableParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get devices deviceid mactable params
func (o *GetDevicesDeviceidMactableParams) SetCount(count *int64) {
	o.Count = count
}

// WithDeviceID adds the deviceID to the get devices deviceid mactable params
func (o *GetDevicesDeviceidMactableParams) WithDeviceID(deviceID string) *GetDevicesDeviceidMactableParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get devices deviceid mactable params
func (o *GetDevicesDeviceidMactableParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithPage adds the page to the get devices deviceid mactable params
func (o *GetDevicesDeviceidMactableParams) WithPage(page *int64) *GetDevicesDeviceidMactableParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get devices deviceid mactable params
func (o *GetDevicesDeviceidMactableParams) SetPage(page *int64) {
	o.Page = page
}

// WithSearch adds the search to the get devices deviceid mactable params
func (o *GetDevicesDeviceidMactableParams) WithSearch(search *string) *GetDevicesDeviceidMactableParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get devices deviceid mactable params
func (o *GetDevicesDeviceidMactableParams) SetSearch(search *string) {
	o.Search = search
}

// WithSortBy adds the sortBy to the get devices deviceid mactable params
func (o *GetDevicesDeviceidMactableParams) WithSortBy(sortBy *string) *GetDevicesDeviceidMactableParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get devices deviceid mactable params
func (o *GetDevicesDeviceidMactableParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortDesc adds the sortDesc to the get devices deviceid mactable params
func (o *GetDevicesDeviceidMactableParams) WithSortDesc(sortDesc *bool) *GetDevicesDeviceidMactableParams {
	o.SetSortDesc(sortDesc)
	return o
}

// SetSortDesc adds the sortDesc to the get devices deviceid mactable params
func (o *GetDevicesDeviceidMactableParams) SetSortDesc(sortDesc *bool) {
	o.SortDesc = sortDesc
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesDeviceidMactableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int64

		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {

			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}
	}

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.Search != nil {

		// query param search
		var qrSearch string

		if o.Search != nil {
			qrSearch = *o.Search
		}
		qSearch := qrSearch
		if qSearch != "" {

			if err := r.SetQueryParam("search", qSearch); err != nil {
				return err
			}
		}
	}

	if o.SortBy != nil {

		// query param sortBy
		var qrSortBy string

		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {

			if err := r.SetQueryParam("sortBy", qSortBy); err != nil {
				return err
			}
		}
	}

	if o.SortDesc != nil {

		// query param sortDesc
		var qrSortDesc bool

		if o.SortDesc != nil {
			qrSortDesc = *o.SortDesc
		}
		qSortDesc := swag.FormatBool(qrSortDesc)
		if qSortDesc != "" {

			if err := r.SetQueryParam("sortDesc", qSortDesc); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
