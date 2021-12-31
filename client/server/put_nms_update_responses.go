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

// PutNmsUpdateReader is a Reader for the PutNmsUpdate structure.
type PutNmsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNmsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutNmsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutNmsUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutNmsUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutNmsUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutNmsUpdateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutNmsUpdateOK creates a PutNmsUpdateOK with default headers values
func NewPutNmsUpdateOK() *PutNmsUpdateOK {
	return &PutNmsUpdateOK{}
}

/* PutNmsUpdateOK describes a response with status code 200, with default header values.

Successful
*/
type PutNmsUpdateOK struct {
	Payload *models.Status
}

func (o *PutNmsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /nms/update][%d] putNmsUpdateOK  %+v", 200, o.Payload)
}
func (o *PutNmsUpdateOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PutNmsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNmsUpdateUnauthorized creates a PutNmsUpdateUnauthorized with default headers values
func NewPutNmsUpdateUnauthorized() *PutNmsUpdateUnauthorized {
	return &PutNmsUpdateUnauthorized{}
}

/* PutNmsUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutNmsUpdateUnauthorized struct {
	Payload *models.Error
}

func (o *PutNmsUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /nms/update][%d] putNmsUpdateUnauthorized  %+v", 401, o.Payload)
}
func (o *PutNmsUpdateUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNmsUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNmsUpdateForbidden creates a PutNmsUpdateForbidden with default headers values
func NewPutNmsUpdateForbidden() *PutNmsUpdateForbidden {
	return &PutNmsUpdateForbidden{}
}

/* PutNmsUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutNmsUpdateForbidden struct {
	Payload *models.Error
}

func (o *PutNmsUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /nms/update][%d] putNmsUpdateForbidden  %+v", 403, o.Payload)
}
func (o *PutNmsUpdateForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNmsUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNmsUpdateInternalServerError creates a PutNmsUpdateInternalServerError with default headers values
func NewPutNmsUpdateInternalServerError() *PutNmsUpdateInternalServerError {
	return &PutNmsUpdateInternalServerError{}
}

/* PutNmsUpdateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutNmsUpdateInternalServerError struct {
	Payload *models.Error
}

func (o *PutNmsUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /nms/update][%d] putNmsUpdateInternalServerError  %+v", 500, o.Payload)
}
func (o *PutNmsUpdateInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNmsUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNmsUpdateServiceUnavailable creates a PutNmsUpdateServiceUnavailable with default headers values
func NewPutNmsUpdateServiceUnavailable() *PutNmsUpdateServiceUnavailable {
	return &PutNmsUpdateServiceUnavailable{}
}

/* PutNmsUpdateServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable
*/
type PutNmsUpdateServiceUnavailable struct {
	Payload *models.Error
}

func (o *PutNmsUpdateServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /nms/update][%d] putNmsUpdateServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PutNmsUpdateServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNmsUpdateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
