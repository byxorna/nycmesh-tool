// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// PutDevicesEroutersIDRouterOspfReader is a Reader for the PutDevicesEroutersIDRouterOspf structure.
type PutDevicesEroutersIDRouterOspfReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDevicesEroutersIDRouterOspfReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutDevicesEroutersIDRouterOspfOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutDevicesEroutersIDRouterOspfBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutDevicesEroutersIDRouterOspfUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutDevicesEroutersIDRouterOspfForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutDevicesEroutersIDRouterOspfNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutDevicesEroutersIDRouterOspfInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutDevicesEroutersIDRouterOspfOK creates a PutDevicesEroutersIDRouterOspfOK with default headers values
func NewPutDevicesEroutersIDRouterOspfOK() *PutDevicesEroutersIDRouterOspfOK {
	return &PutDevicesEroutersIDRouterOspfOK{}
}

/* PutDevicesEroutersIDRouterOspfOK describes a response with status code 200, with default header values.

Successful
*/
type PutDevicesEroutersIDRouterOspfOK struct {
	Payload *models.EdgeRouterOspf
}

func (o *PutDevicesEroutersIDRouterOspfOK) Error() string {
	return fmt.Sprintf("[PUT /devices/erouters/{id}/router/ospf][%d] putDevicesEroutersIdRouterOspfOK  %+v", 200, o.Payload)
}
func (o *PutDevicesEroutersIDRouterOspfOK) GetPayload() *models.EdgeRouterOspf {
	return o.Payload
}

func (o *PutDevicesEroutersIDRouterOspfOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgeRouterOspf)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesEroutersIDRouterOspfBadRequest creates a PutDevicesEroutersIDRouterOspfBadRequest with default headers values
func NewPutDevicesEroutersIDRouterOspfBadRequest() *PutDevicesEroutersIDRouterOspfBadRequest {
	return &PutDevicesEroutersIDRouterOspfBadRequest{}
}

/* PutDevicesEroutersIDRouterOspfBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutDevicesEroutersIDRouterOspfBadRequest struct {
	Payload *models.Error
}

func (o *PutDevicesEroutersIDRouterOspfBadRequest) Error() string {
	return fmt.Sprintf("[PUT /devices/erouters/{id}/router/ospf][%d] putDevicesEroutersIdRouterOspfBadRequest  %+v", 400, o.Payload)
}
func (o *PutDevicesEroutersIDRouterOspfBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesEroutersIDRouterOspfBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesEroutersIDRouterOspfUnauthorized creates a PutDevicesEroutersIDRouterOspfUnauthorized with default headers values
func NewPutDevicesEroutersIDRouterOspfUnauthorized() *PutDevicesEroutersIDRouterOspfUnauthorized {
	return &PutDevicesEroutersIDRouterOspfUnauthorized{}
}

/* PutDevicesEroutersIDRouterOspfUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutDevicesEroutersIDRouterOspfUnauthorized struct {
	Payload *models.Error
}

func (o *PutDevicesEroutersIDRouterOspfUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /devices/erouters/{id}/router/ospf][%d] putDevicesEroutersIdRouterOspfUnauthorized  %+v", 401, o.Payload)
}
func (o *PutDevicesEroutersIDRouterOspfUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesEroutersIDRouterOspfUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesEroutersIDRouterOspfForbidden creates a PutDevicesEroutersIDRouterOspfForbidden with default headers values
func NewPutDevicesEroutersIDRouterOspfForbidden() *PutDevicesEroutersIDRouterOspfForbidden {
	return &PutDevicesEroutersIDRouterOspfForbidden{}
}

/* PutDevicesEroutersIDRouterOspfForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutDevicesEroutersIDRouterOspfForbidden struct {
	Payload *models.Error
}

func (o *PutDevicesEroutersIDRouterOspfForbidden) Error() string {
	return fmt.Sprintf("[PUT /devices/erouters/{id}/router/ospf][%d] putDevicesEroutersIdRouterOspfForbidden  %+v", 403, o.Payload)
}
func (o *PutDevicesEroutersIDRouterOspfForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesEroutersIDRouterOspfForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesEroutersIDRouterOspfNotFound creates a PutDevicesEroutersIDRouterOspfNotFound with default headers values
func NewPutDevicesEroutersIDRouterOspfNotFound() *PutDevicesEroutersIDRouterOspfNotFound {
	return &PutDevicesEroutersIDRouterOspfNotFound{}
}

/* PutDevicesEroutersIDRouterOspfNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutDevicesEroutersIDRouterOspfNotFound struct {
	Payload *models.Error
}

func (o *PutDevicesEroutersIDRouterOspfNotFound) Error() string {
	return fmt.Sprintf("[PUT /devices/erouters/{id}/router/ospf][%d] putDevicesEroutersIdRouterOspfNotFound  %+v", 404, o.Payload)
}
func (o *PutDevicesEroutersIDRouterOspfNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesEroutersIDRouterOspfNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesEroutersIDRouterOspfInternalServerError creates a PutDevicesEroutersIDRouterOspfInternalServerError with default headers values
func NewPutDevicesEroutersIDRouterOspfInternalServerError() *PutDevicesEroutersIDRouterOspfInternalServerError {
	return &PutDevicesEroutersIDRouterOspfInternalServerError{}
}

/* PutDevicesEroutersIDRouterOspfInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutDevicesEroutersIDRouterOspfInternalServerError struct {
	Payload *models.Error
}

func (o *PutDevicesEroutersIDRouterOspfInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /devices/erouters/{id}/router/ospf][%d] putDevicesEroutersIdRouterOspfInternalServerError  %+v", 500, o.Payload)
}
func (o *PutDevicesEroutersIDRouterOspfInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesEroutersIDRouterOspfInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
