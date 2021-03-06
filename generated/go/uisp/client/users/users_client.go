// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteUsersID(params *DeleteUsersIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUsersIDOK, error)

	GetNmsKeepalive(params *GetNmsKeepaliveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetNmsKeepaliveOK, error)

	GetUsers(params *GetUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUsersOK, error)

	PostUsers(params *PostUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostUsersOK, error)

	PostUsersIDReinvite(params *PostUsersIDReinviteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostUsersIDReinviteOK, error)

	PostUsersIDReinviteLink(params *PostUsersIDReinviteLinkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostUsersIDReinviteLinkOK, error)

	PostUsersImport(params *PostUsersImportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostUsersImportOK, error)

	PostUsersInvite(params *PostUsersInviteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostUsersInviteOK, error)

	PutUsersID(params *PutUsersIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutUsersIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteUsersID deletes user
*/
func (a *Client) DeleteUsersID(params *DeleteUsersIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUsersIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUsersIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteUsersId",
		Method:             "DELETE",
		PathPattern:        "/users/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteUsersIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteUsersIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteUsersId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNmsKeepalive extends validity of user token
*/
func (a *Client) GetNmsKeepalive(params *GetNmsKeepaliveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetNmsKeepaliveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNmsKeepaliveParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getNmsKeepalive",
		Method:             "GET",
		PathPattern:        "/nms/keep-alive",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetNmsKeepaliveReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNmsKeepaliveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNmsKeepalive: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUsers returns list of all users
*/
func (a *Client) GetUsers(params *GetUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUsers",
		Method:             "GET",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostUsers creates new user
*/
func (a *Client) PostUsers(params *PostUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postUsers",
		Method:             "POST",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostUsersIDReinvite reinvites user by email
*/
func (a *Client) PostUsersIDReinvite(params *PostUsersIDReinviteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostUsersIDReinviteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostUsersIDReinviteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postUsersIdReinvite",
		Method:             "POST",
		PathPattern:        "/users/{id}/reinvite",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostUsersIDReinviteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostUsersIDReinviteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postUsersIdReinvite: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostUsersIDReinviteLink generates new invitation link
*/
func (a *Client) PostUsersIDReinviteLink(params *PostUsersIDReinviteLinkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostUsersIDReinviteLinkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostUsersIDReinviteLinkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postUsersIdReinviteLink",
		Method:             "POST",
		PathPattern:        "/users/{id}/reinvite/link",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostUsersIDReinviteLinkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostUsersIDReinviteLinkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postUsersIdReinviteLink: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostUsersImport imports users
*/
func (a *Client) PostUsersImport(params *PostUsersImportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostUsersImportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostUsersImportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postUsersImport",
		Method:             "POST",
		PathPattern:        "/users/import",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostUsersImportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostUsersImportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postUsersImport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostUsersInvite invites new user
*/
func (a *Client) PostUsersInvite(params *PostUsersInviteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostUsersInviteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostUsersInviteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postUsersInvite",
		Method:             "POST",
		PathPattern:        "/users/invite",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostUsersInviteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostUsersInviteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postUsersInvite: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutUsersID updates user
*/
func (a *Client) PutUsersID(params *PutUsersIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutUsersIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutUsersIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "putUsersId",
		Method:             "PUT",
		PathPattern:        "/users/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutUsersIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutUsersIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for putUsersId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
