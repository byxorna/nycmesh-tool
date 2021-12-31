// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/models"
)

// GetNmsConnectionReader is a Reader for the GetNmsConnection structure.
type GetNmsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNmsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNmsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetNmsConnectionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNmsConnectionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNmsConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNmsConnectionOK creates a GetNmsConnectionOK with default headers values
func NewGetNmsConnectionOK() *GetNmsConnectionOK {
	return &GetNmsConnectionOK{}
}

/* GetNmsConnectionOK describes a response with status code 200, with default header values.

Successful
*/
type GetNmsConnectionOK struct {
	Payload string
}

func (o *GetNmsConnectionOK) Error() string {
	return fmt.Sprintf("[GET /nms/connection][%d] getNmsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetNmsConnectionOK) GetPayload() string {
	return o.Payload
}

func (o *GetNmsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNmsConnectionUnauthorized creates a GetNmsConnectionUnauthorized with default headers values
func NewGetNmsConnectionUnauthorized() *GetNmsConnectionUnauthorized {
	return &GetNmsConnectionUnauthorized{}
}

/* GetNmsConnectionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetNmsConnectionUnauthorized struct {
	Payload *models.Error
}

func (o *GetNmsConnectionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /nms/connection][%d] getNmsConnectionUnauthorized  %+v", 401, o.Payload)
}
func (o *GetNmsConnectionUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNmsConnectionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNmsConnectionForbidden creates a GetNmsConnectionForbidden with default headers values
func NewGetNmsConnectionForbidden() *GetNmsConnectionForbidden {
	return &GetNmsConnectionForbidden{}
}

/* GetNmsConnectionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetNmsConnectionForbidden struct {
	Payload *models.Error
}

func (o *GetNmsConnectionForbidden) Error() string {
	return fmt.Sprintf("[GET /nms/connection][%d] getNmsConnectionForbidden  %+v", 403, o.Payload)
}
func (o *GetNmsConnectionForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNmsConnectionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNmsConnectionInternalServerError creates a GetNmsConnectionInternalServerError with default headers values
func NewGetNmsConnectionInternalServerError() *GetNmsConnectionInternalServerError {
	return &GetNmsConnectionInternalServerError{}
}

/* GetNmsConnectionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetNmsConnectionInternalServerError struct {
	Payload *models.Error
}

func (o *GetNmsConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /nms/connection][%d] getNmsConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetNmsConnectionInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNmsConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
