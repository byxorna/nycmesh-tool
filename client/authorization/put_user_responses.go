// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/models"
)

// PutUserReader is a Reader for the PutUser structure.
type PutUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutUserOK creates a PutUserOK with default headers values
func NewPutUserOK() *PutUserOK {
	return &PutUserOK{}
}

/* PutUserOK describes a response with status code 200, with default header values.

Successful
*/
type PutUserOK struct {
	Payload *models.User
}

func (o *PutUserOK) Error() string {
	return fmt.Sprintf("[PUT /user][%d] putUserOK  %+v", 200, o.Payload)
}
func (o *PutUserOK) GetPayload() *models.User {
	return o.Payload
}

func (o *PutUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserBadRequest creates a PutUserBadRequest with default headers values
func NewPutUserBadRequest() *PutUserBadRequest {
	return &PutUserBadRequest{}
}

/* PutUserBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutUserBadRequest struct {
	Payload *models.Error
}

func (o *PutUserBadRequest) Error() string {
	return fmt.Sprintf("[PUT /user][%d] putUserBadRequest  %+v", 400, o.Payload)
}
func (o *PutUserBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserUnauthorized creates a PutUserUnauthorized with default headers values
func NewPutUserUnauthorized() *PutUserUnauthorized {
	return &PutUserUnauthorized{}
}

/* PutUserUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutUserUnauthorized struct {
	Payload *models.Error
}

func (o *PutUserUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /user][%d] putUserUnauthorized  %+v", 401, o.Payload)
}
func (o *PutUserUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserForbidden creates a PutUserForbidden with default headers values
func NewPutUserForbidden() *PutUserForbidden {
	return &PutUserForbidden{}
}

/* PutUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutUserForbidden struct {
	Payload *models.Error
}

func (o *PutUserForbidden) Error() string {
	return fmt.Sprintf("[PUT /user][%d] putUserForbidden  %+v", 403, o.Payload)
}
func (o *PutUserForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserInternalServerError creates a PutUserInternalServerError with default headers values
func NewPutUserInternalServerError() *PutUserInternalServerError {
	return &PutUserInternalServerError{}
}

/* PutUserInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutUserInternalServerError struct {
	Payload *models.Error
}

func (o *PutUserInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /user][%d] putUserInternalServerError  %+v", 500, o.Payload)
}
func (o *PutUserInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
