// Code generated by go-swagger; DO NOT EDIT.

package token

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

// NewDeleteTokenTokenidParams creates a new DeleteTokenTokenidParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteTokenTokenidParams() *DeleteTokenTokenidParams {
	return &DeleteTokenTokenidParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTokenTokenidParamsWithTimeout creates a new DeleteTokenTokenidParams object
// with the ability to set a timeout on a request.
func NewDeleteTokenTokenidParamsWithTimeout(timeout time.Duration) *DeleteTokenTokenidParams {
	return &DeleteTokenTokenidParams{
		timeout: timeout,
	}
}

// NewDeleteTokenTokenidParamsWithContext creates a new DeleteTokenTokenidParams object
// with the ability to set a context for a request.
func NewDeleteTokenTokenidParamsWithContext(ctx context.Context) *DeleteTokenTokenidParams {
	return &DeleteTokenTokenidParams{
		Context: ctx,
	}
}

// NewDeleteTokenTokenidParamsWithHTTPClient creates a new DeleteTokenTokenidParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteTokenTokenidParamsWithHTTPClient(client *http.Client) *DeleteTokenTokenidParams {
	return &DeleteTokenTokenidParams{
		HTTPClient: client,
	}
}

/* DeleteTokenTokenidParams contains all the parameters to send to the API endpoint
   for the delete token tokenid operation.

   Typically these are written to a http.Request.
*/
type DeleteTokenTokenidParams struct {

	// TokenID.
	TokenID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete token tokenid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTokenTokenidParams) WithDefaults() *DeleteTokenTokenidParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete token tokenid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTokenTokenidParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete token tokenid params
func (o *DeleteTokenTokenidParams) WithTimeout(timeout time.Duration) *DeleteTokenTokenidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete token tokenid params
func (o *DeleteTokenTokenidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete token tokenid params
func (o *DeleteTokenTokenidParams) WithContext(ctx context.Context) *DeleteTokenTokenidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete token tokenid params
func (o *DeleteTokenTokenidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete token tokenid params
func (o *DeleteTokenTokenidParams) WithHTTPClient(client *http.Client) *DeleteTokenTokenidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete token tokenid params
func (o *DeleteTokenTokenidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTokenID adds the tokenID to the delete token tokenid params
func (o *DeleteTokenTokenidParams) WithTokenID(tokenID string) *DeleteTokenTokenidParams {
	o.SetTokenID(tokenID)
	return o
}

// SetTokenID adds the tokenId to the delete token tokenid params
func (o *DeleteTokenTokenidParams) SetTokenID(tokenID string) {
	o.TokenID = tokenID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTokenTokenidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tokenId
	if err := r.SetPathParam("tokenId", o.TokenID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
