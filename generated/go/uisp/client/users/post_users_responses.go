// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// PostUsersReader is a Reader for the PostUsers structure.
type PostUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostUsersOK creates a PostUsersOK with default headers values
func NewPostUsersOK() *PostUsersOK {
	return &PostUsersOK{}
}

/* PostUsersOK describes a response with status code 200, with default header values.

Successful
*/
type PostUsersOK struct {
	Payload *models.User
}

func (o *PostUsersOK) Error() string {
	return fmt.Sprintf("[POST /users][%d] postUsersOK  %+v", 200, o.Payload)
}
func (o *PostUsersOK) GetPayload() *models.User {
	return o.Payload
}

func (o *PostUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersBadRequest creates a PostUsersBadRequest with default headers values
func NewPostUsersBadRequest() *PostUsersBadRequest {
	return &PostUsersBadRequest{}
}

/* PostUsersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostUsersBadRequest struct {
	Payload *models.Error
}

func (o *PostUsersBadRequest) Error() string {
	return fmt.Sprintf("[POST /users][%d] postUsersBadRequest  %+v", 400, o.Payload)
}
func (o *PostUsersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersUnauthorized creates a PostUsersUnauthorized with default headers values
func NewPostUsersUnauthorized() *PostUsersUnauthorized {
	return &PostUsersUnauthorized{}
}

/* PostUsersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostUsersUnauthorized struct {
	Payload *models.Error
}

func (o *PostUsersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /users][%d] postUsersUnauthorized  %+v", 401, o.Payload)
}
func (o *PostUsersUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersForbidden creates a PostUsersForbidden with default headers values
func NewPostUsersForbidden() *PostUsersForbidden {
	return &PostUsersForbidden{}
}

/* PostUsersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostUsersForbidden struct {
	Payload *models.Error
}

func (o *PostUsersForbidden) Error() string {
	return fmt.Sprintf("[POST /users][%d] postUsersForbidden  %+v", 403, o.Payload)
}
func (o *PostUsersForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersInternalServerError creates a PostUsersInternalServerError with default headers values
func NewPostUsersInternalServerError() *PostUsersInternalServerError {
	return &PostUsersInternalServerError{}
}

/* PostUsersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostUsersInternalServerError struct {
	Payload *models.Error
}

func (o *PostUsersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users][%d] postUsersInternalServerError  %+v", 500, o.Payload)
}
func (o *PostUsersInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
