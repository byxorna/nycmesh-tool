// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/models"
)

// PostUsersInviteReader is a Reader for the PostUsersInvite structure.
type PostUsersInviteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUsersInviteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUsersInviteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUsersInviteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUsersInviteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUsersInviteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUsersInviteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostUsersInviteOK creates a PostUsersInviteOK with default headers values
func NewPostUsersInviteOK() *PostUsersInviteOK {
	return &PostUsersInviteOK{}
}

/* PostUsersInviteOK describes a response with status code 200, with default header values.

Successful
*/
type PostUsersInviteOK struct {
	Payload *models.InviteUserResponse
}

func (o *PostUsersInviteOK) Error() string {
	return fmt.Sprintf("[POST /users/invite][%d] postUsersInviteOK  %+v", 200, o.Payload)
}
func (o *PostUsersInviteOK) GetPayload() *models.InviteUserResponse {
	return o.Payload
}

func (o *PostUsersInviteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InviteUserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersInviteBadRequest creates a PostUsersInviteBadRequest with default headers values
func NewPostUsersInviteBadRequest() *PostUsersInviteBadRequest {
	return &PostUsersInviteBadRequest{}
}

/* PostUsersInviteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostUsersInviteBadRequest struct {
	Payload *models.Error
}

func (o *PostUsersInviteBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/invite][%d] postUsersInviteBadRequest  %+v", 400, o.Payload)
}
func (o *PostUsersInviteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUsersInviteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersInviteUnauthorized creates a PostUsersInviteUnauthorized with default headers values
func NewPostUsersInviteUnauthorized() *PostUsersInviteUnauthorized {
	return &PostUsersInviteUnauthorized{}
}

/* PostUsersInviteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostUsersInviteUnauthorized struct {
	Payload *models.Error
}

func (o *PostUsersInviteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /users/invite][%d] postUsersInviteUnauthorized  %+v", 401, o.Payload)
}
func (o *PostUsersInviteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUsersInviteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersInviteForbidden creates a PostUsersInviteForbidden with default headers values
func NewPostUsersInviteForbidden() *PostUsersInviteForbidden {
	return &PostUsersInviteForbidden{}
}

/* PostUsersInviteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostUsersInviteForbidden struct {
	Payload *models.Error
}

func (o *PostUsersInviteForbidden) Error() string {
	return fmt.Sprintf("[POST /users/invite][%d] postUsersInviteForbidden  %+v", 403, o.Payload)
}
func (o *PostUsersInviteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUsersInviteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersInviteInternalServerError creates a PostUsersInviteInternalServerError with default headers values
func NewPostUsersInviteInternalServerError() *PostUsersInviteInternalServerError {
	return &PostUsersInviteInternalServerError{}
}

/* PostUsersInviteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostUsersInviteInternalServerError struct {
	Payload *models.Error
}

func (o *PostUsersInviteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users/invite][%d] postUsersInviteInternalServerError  %+v", 500, o.Payload)
}
func (o *PostUsersInviteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUsersInviteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
