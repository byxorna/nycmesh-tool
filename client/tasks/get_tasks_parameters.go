// Code generated by go-swagger; DO NOT EDIT.

package tasks

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

// NewGetTasksParams creates a new GetTasksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTasksParams() *GetTasksParams {
	return &GetTasksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTasksParamsWithTimeout creates a new GetTasksParams object
// with the ability to set a timeout on a request.
func NewGetTasksParamsWithTimeout(timeout time.Duration) *GetTasksParams {
	return &GetTasksParams{
		timeout: timeout,
	}
}

// NewGetTasksParamsWithContext creates a new GetTasksParams object
// with the ability to set a context for a request.
func NewGetTasksParamsWithContext(ctx context.Context) *GetTasksParams {
	return &GetTasksParams{
		Context: ctx,
	}
}

// NewGetTasksParamsWithHTTPClient creates a new GetTasksParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTasksParamsWithHTTPClient(client *http.Client) *GetTasksParams {
	return &GetTasksParams{
		HTTPClient: client,
	}
}

/* GetTasksParams contains all the parameters to send to the API endpoint
   for the get tasks operation.

   Typically these are written to a http.Request.
*/
type GetTasksParams struct {

	// Count.
	Count float64

	// Page.
	Page float64

	// Period.
	Period *float64

	// Status.
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTasksParams) WithDefaults() *GetTasksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTasksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get tasks params
func (o *GetTasksParams) WithTimeout(timeout time.Duration) *GetTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tasks params
func (o *GetTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tasks params
func (o *GetTasksParams) WithContext(ctx context.Context) *GetTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tasks params
func (o *GetTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tasks params
func (o *GetTasksParams) WithHTTPClient(client *http.Client) *GetTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tasks params
func (o *GetTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get tasks params
func (o *GetTasksParams) WithCount(count float64) *GetTasksParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get tasks params
func (o *GetTasksParams) SetCount(count float64) {
	o.Count = count
}

// WithPage adds the page to the get tasks params
func (o *GetTasksParams) WithPage(page float64) *GetTasksParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get tasks params
func (o *GetTasksParams) SetPage(page float64) {
	o.Page = page
}

// WithPeriod adds the period to the get tasks params
func (o *GetTasksParams) WithPeriod(period *float64) *GetTasksParams {
	o.SetPeriod(period)
	return o
}

// SetPeriod adds the period to the get tasks params
func (o *GetTasksParams) SetPeriod(period *float64) {
	o.Period = period
}

// WithStatus adds the status to the get tasks params
func (o *GetTasksParams) WithStatus(status *string) *GetTasksParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get tasks params
func (o *GetTasksParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *GetTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
