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

// PostUserLogoutReader is a Reader for the PostUserLogout structure.
type PostUserLogoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUserLogoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUserLogoutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostUserLogoutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUserLogoutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUserLogoutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostUserLogoutOK creates a PostUserLogoutOK with default headers values
func NewPostUserLogoutOK() *PostUserLogoutOK {
	return &PostUserLogoutOK{}
}

/* PostUserLogoutOK describes a response with status code 200, with default header values.

Successful
*/
type PostUserLogoutOK struct {
	Payload *models.Status
}

func (o *PostUserLogoutOK) Error() string {
	return fmt.Sprintf("[POST /user/logout][%d] postUserLogoutOK  %+v", 200, o.Payload)
}
func (o *PostUserLogoutOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PostUserLogoutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserLogoutUnauthorized creates a PostUserLogoutUnauthorized with default headers values
func NewPostUserLogoutUnauthorized() *PostUserLogoutUnauthorized {
	return &PostUserLogoutUnauthorized{}
}

/* PostUserLogoutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostUserLogoutUnauthorized struct {
	Payload *models.Error
}

func (o *PostUserLogoutUnauthorized) Error() string {
	return fmt.Sprintf("[POST /user/logout][%d] postUserLogoutUnauthorized  %+v", 401, o.Payload)
}
func (o *PostUserLogoutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUserLogoutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserLogoutForbidden creates a PostUserLogoutForbidden with default headers values
func NewPostUserLogoutForbidden() *PostUserLogoutForbidden {
	return &PostUserLogoutForbidden{}
}

/* PostUserLogoutForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostUserLogoutForbidden struct {
	Payload *models.Error
}

func (o *PostUserLogoutForbidden) Error() string {
	return fmt.Sprintf("[POST /user/logout][%d] postUserLogoutForbidden  %+v", 403, o.Payload)
}
func (o *PostUserLogoutForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUserLogoutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserLogoutInternalServerError creates a PostUserLogoutInternalServerError with default headers values
func NewPostUserLogoutInternalServerError() *PostUserLogoutInternalServerError {
	return &PostUserLogoutInternalServerError{}
}

/* PostUserLogoutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostUserLogoutInternalServerError struct {
	Payload *models.Error
}

func (o *PostUserLogoutInternalServerError) Error() string {
	return fmt.Sprintf("[POST /user/logout][%d] postUserLogoutInternalServerError  %+v", 500, o.Payload)
}
func (o *PostUserLogoutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUserLogoutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}