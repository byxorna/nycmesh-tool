// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewPostNmsQuestionnairesIDTouchParams creates a new PostNmsQuestionnairesIDTouchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostNmsQuestionnairesIDTouchParams() *PostNmsQuestionnairesIDTouchParams {
	return &PostNmsQuestionnairesIDTouchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostNmsQuestionnairesIDTouchParamsWithTimeout creates a new PostNmsQuestionnairesIDTouchParams object
// with the ability to set a timeout on a request.
func NewPostNmsQuestionnairesIDTouchParamsWithTimeout(timeout time.Duration) *PostNmsQuestionnairesIDTouchParams {
	return &PostNmsQuestionnairesIDTouchParams{
		timeout: timeout,
	}
}

// NewPostNmsQuestionnairesIDTouchParamsWithContext creates a new PostNmsQuestionnairesIDTouchParams object
// with the ability to set a context for a request.
func NewPostNmsQuestionnairesIDTouchParamsWithContext(ctx context.Context) *PostNmsQuestionnairesIDTouchParams {
	return &PostNmsQuestionnairesIDTouchParams{
		Context: ctx,
	}
}

// NewPostNmsQuestionnairesIDTouchParamsWithHTTPClient creates a new PostNmsQuestionnairesIDTouchParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostNmsQuestionnairesIDTouchParamsWithHTTPClient(client *http.Client) *PostNmsQuestionnairesIDTouchParams {
	return &PostNmsQuestionnairesIDTouchParams{
		HTTPClient: client,
	}
}

/* PostNmsQuestionnairesIDTouchParams contains all the parameters to send to the API endpoint
   for the post nms questionnaires Id touch operation.

   Typically these are written to a http.Request.
*/
type PostNmsQuestionnairesIDTouchParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post nms questionnaires Id touch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostNmsQuestionnairesIDTouchParams) WithDefaults() *PostNmsQuestionnairesIDTouchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post nms questionnaires Id touch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostNmsQuestionnairesIDTouchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post nms questionnaires Id touch params
func (o *PostNmsQuestionnairesIDTouchParams) WithTimeout(timeout time.Duration) *PostNmsQuestionnairesIDTouchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post nms questionnaires Id touch params
func (o *PostNmsQuestionnairesIDTouchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post nms questionnaires Id touch params
func (o *PostNmsQuestionnairesIDTouchParams) WithContext(ctx context.Context) *PostNmsQuestionnairesIDTouchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post nms questionnaires Id touch params
func (o *PostNmsQuestionnairesIDTouchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post nms questionnaires Id touch params
func (o *PostNmsQuestionnairesIDTouchParams) WithHTTPClient(client *http.Client) *PostNmsQuestionnairesIDTouchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post nms questionnaires Id touch params
func (o *PostNmsQuestionnairesIDTouchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post nms questionnaires Id touch params
func (o *PostNmsQuestionnairesIDTouchParams) WithID(id string) *PostNmsQuestionnairesIDTouchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post nms questionnaires Id touch params
func (o *PostNmsQuestionnairesIDTouchParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostNmsQuestionnairesIDTouchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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