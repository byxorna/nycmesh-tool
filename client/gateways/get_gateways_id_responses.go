// Code generated by go-swagger; DO NOT EDIT.

package gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/models"
)

// GetGatewaysIDReader is a Reader for the GetGatewaysID structure.
type GetGatewaysIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGatewaysIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGatewaysIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetGatewaysIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGatewaysIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGatewaysIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGatewaysIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGatewaysIDOK creates a GetGatewaysIDOK with default headers values
func NewGetGatewaysIDOK() *GetGatewaysIDOK {
	return &GetGatewaysIDOK{}
}

/* GetGatewaysIDOK describes a response with status code 200, with default header values.

Successful
*/
type GetGatewaysIDOK struct {
	Payload *models.Model8
}

func (o *GetGatewaysIDOK) Error() string {
	return fmt.Sprintf("[GET /gateways/{id}][%d] getGatewaysIdOK  %+v", 200, o.Payload)
}
func (o *GetGatewaysIDOK) GetPayload() *models.Model8 {
	return o.Payload
}

func (o *GetGatewaysIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Model8)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGatewaysIDUnauthorized creates a GetGatewaysIDUnauthorized with default headers values
func NewGetGatewaysIDUnauthorized() *GetGatewaysIDUnauthorized {
	return &GetGatewaysIDUnauthorized{}
}

/* GetGatewaysIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetGatewaysIDUnauthorized struct {
	Payload *models.Error
}

func (o *GetGatewaysIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /gateways/{id}][%d] getGatewaysIdUnauthorized  %+v", 401, o.Payload)
}
func (o *GetGatewaysIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGatewaysIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGatewaysIDForbidden creates a GetGatewaysIDForbidden with default headers values
func NewGetGatewaysIDForbidden() *GetGatewaysIDForbidden {
	return &GetGatewaysIDForbidden{}
}

/* GetGatewaysIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetGatewaysIDForbidden struct {
	Payload *models.Error
}

func (o *GetGatewaysIDForbidden) Error() string {
	return fmt.Sprintf("[GET /gateways/{id}][%d] getGatewaysIdForbidden  %+v", 403, o.Payload)
}
func (o *GetGatewaysIDForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGatewaysIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGatewaysIDNotFound creates a GetGatewaysIDNotFound with default headers values
func NewGetGatewaysIDNotFound() *GetGatewaysIDNotFound {
	return &GetGatewaysIDNotFound{}
}

/* GetGatewaysIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetGatewaysIDNotFound struct {
	Payload *models.Error
}

func (o *GetGatewaysIDNotFound) Error() string {
	return fmt.Sprintf("[GET /gateways/{id}][%d] getGatewaysIdNotFound  %+v", 404, o.Payload)
}
func (o *GetGatewaysIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGatewaysIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGatewaysIDInternalServerError creates a GetGatewaysIDInternalServerError with default headers values
func NewGetGatewaysIDInternalServerError() *GetGatewaysIDInternalServerError {
	return &GetGatewaysIDInternalServerError{}
}

/* GetGatewaysIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetGatewaysIDInternalServerError struct {
	Payload *models.Error
}

func (o *GetGatewaysIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /gateways/{id}][%d] getGatewaysIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetGatewaysIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGatewaysIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
