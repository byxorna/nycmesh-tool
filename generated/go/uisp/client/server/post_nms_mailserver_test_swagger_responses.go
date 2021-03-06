// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// PostNmsMailserverTestReader is a Reader for the PostNmsMailserverTest structure.
type PostNmsMailserverTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNmsMailserverTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostNmsMailserverTestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostNmsMailserverTestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostNmsMailserverTestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostNmsMailserverTestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostNmsMailserverTestOK creates a PostNmsMailserverTestOK with default headers values
func NewPostNmsMailserverTestOK() *PostNmsMailserverTestOK {
	return &PostNmsMailserverTestOK{}
}

/* PostNmsMailserverTestOK describes a response with status code 200, with default header values.

Successful
*/
type PostNmsMailserverTestOK struct {
	Payload *models.Status
}

func (o *PostNmsMailserverTestOK) Error() string {
	return fmt.Sprintf("[POST /nms/mailserver/test][%d] postNmsMailserverTestOK  %+v", 200, o.Payload)
}
func (o *PostNmsMailserverTestOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PostNmsMailserverTestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsMailserverTestUnauthorized creates a PostNmsMailserverTestUnauthorized with default headers values
func NewPostNmsMailserverTestUnauthorized() *PostNmsMailserverTestUnauthorized {
	return &PostNmsMailserverTestUnauthorized{}
}

/* PostNmsMailserverTestUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostNmsMailserverTestUnauthorized struct {
	Payload *models.Error
}

func (o *PostNmsMailserverTestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /nms/mailserver/test][%d] postNmsMailserverTestUnauthorized  %+v", 401, o.Payload)
}
func (o *PostNmsMailserverTestUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsMailserverTestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsMailserverTestForbidden creates a PostNmsMailserverTestForbidden with default headers values
func NewPostNmsMailserverTestForbidden() *PostNmsMailserverTestForbidden {
	return &PostNmsMailserverTestForbidden{}
}

/* PostNmsMailserverTestForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostNmsMailserverTestForbidden struct {
	Payload *models.Error
}

func (o *PostNmsMailserverTestForbidden) Error() string {
	return fmt.Sprintf("[POST /nms/mailserver/test][%d] postNmsMailserverTestForbidden  %+v", 403, o.Payload)
}
func (o *PostNmsMailserverTestForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsMailserverTestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsMailserverTestInternalServerError creates a PostNmsMailserverTestInternalServerError with default headers values
func NewPostNmsMailserverTestInternalServerError() *PostNmsMailserverTestInternalServerError {
	return &PostNmsMailserverTestInternalServerError{}
}

/* PostNmsMailserverTestInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostNmsMailserverTestInternalServerError struct {
	Payload *models.Error
}

func (o *PostNmsMailserverTestInternalServerError) Error() string {
	return fmt.Sprintf("[POST /nms/mailserver/test][%d] postNmsMailserverTestInternalServerError  %+v", 500, o.Payload)
}
func (o *PostNmsMailserverTestInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsMailserverTestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
