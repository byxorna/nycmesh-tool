// Code generated by go-swagger; DO NOT EDIT.

package export

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new export API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for export API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetGdprClientsID(params *GetGdprClientsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) error

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetGdprClientsID downloads g d p r client report
*/
func (a *Client) GetGdprClientsID(params *GetGdprClientsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGdprClientsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGdprClientsId",
		Method:             "GET",
		PathPattern:        "/gdpr/clients/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetGdprClientsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
