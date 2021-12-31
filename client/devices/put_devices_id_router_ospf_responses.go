// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/models"
)

// PutDevicesIDRouterOspfReader is a Reader for the PutDevicesIDRouterOspf structure.
type PutDevicesIDRouterOspfReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDevicesIDRouterOspfReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutDevicesIDRouterOspfOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutDevicesIDRouterOspfBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutDevicesIDRouterOspfUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutDevicesIDRouterOspfForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutDevicesIDRouterOspfNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutDevicesIDRouterOspfInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutDevicesIDRouterOspfOK creates a PutDevicesIDRouterOspfOK with default headers values
func NewPutDevicesIDRouterOspfOK() *PutDevicesIDRouterOspfOK {
	return &PutDevicesIDRouterOspfOK{}
}

/* PutDevicesIDRouterOspfOK describes a response with status code 200, with default header values.

Successful
*/
type PutDevicesIDRouterOspfOK struct {
	Payload *models.RouterOspf
}

func (o *PutDevicesIDRouterOspfOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/router/ospf][%d] putDevicesIdRouterOspfOK  %+v", 200, o.Payload)
}
func (o *PutDevicesIDRouterOspfOK) GetPayload() *models.RouterOspf {
	return o.Payload
}

func (o *PutDevicesIDRouterOspfOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RouterOspf)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDRouterOspfBadRequest creates a PutDevicesIDRouterOspfBadRequest with default headers values
func NewPutDevicesIDRouterOspfBadRequest() *PutDevicesIDRouterOspfBadRequest {
	return &PutDevicesIDRouterOspfBadRequest{}
}

/* PutDevicesIDRouterOspfBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutDevicesIDRouterOspfBadRequest struct {
	Payload *models.Error
}

func (o *PutDevicesIDRouterOspfBadRequest) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/router/ospf][%d] putDevicesIdRouterOspfBadRequest  %+v", 400, o.Payload)
}
func (o *PutDevicesIDRouterOspfBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDRouterOspfBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDRouterOspfUnauthorized creates a PutDevicesIDRouterOspfUnauthorized with default headers values
func NewPutDevicesIDRouterOspfUnauthorized() *PutDevicesIDRouterOspfUnauthorized {
	return &PutDevicesIDRouterOspfUnauthorized{}
}

/* PutDevicesIDRouterOspfUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutDevicesIDRouterOspfUnauthorized struct {
	Payload *models.Error
}

func (o *PutDevicesIDRouterOspfUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/router/ospf][%d] putDevicesIdRouterOspfUnauthorized  %+v", 401, o.Payload)
}
func (o *PutDevicesIDRouterOspfUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDRouterOspfUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDRouterOspfForbidden creates a PutDevicesIDRouterOspfForbidden with default headers values
func NewPutDevicesIDRouterOspfForbidden() *PutDevicesIDRouterOspfForbidden {
	return &PutDevicesIDRouterOspfForbidden{}
}

/* PutDevicesIDRouterOspfForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutDevicesIDRouterOspfForbidden struct {
	Payload *models.Error
}

func (o *PutDevicesIDRouterOspfForbidden) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/router/ospf][%d] putDevicesIdRouterOspfForbidden  %+v", 403, o.Payload)
}
func (o *PutDevicesIDRouterOspfForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDRouterOspfForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDRouterOspfNotFound creates a PutDevicesIDRouterOspfNotFound with default headers values
func NewPutDevicesIDRouterOspfNotFound() *PutDevicesIDRouterOspfNotFound {
	return &PutDevicesIDRouterOspfNotFound{}
}

/* PutDevicesIDRouterOspfNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutDevicesIDRouterOspfNotFound struct {
	Payload *models.Error
}

func (o *PutDevicesIDRouterOspfNotFound) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/router/ospf][%d] putDevicesIdRouterOspfNotFound  %+v", 404, o.Payload)
}
func (o *PutDevicesIDRouterOspfNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDRouterOspfNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDRouterOspfInternalServerError creates a PutDevicesIDRouterOspfInternalServerError with default headers values
func NewPutDevicesIDRouterOspfInternalServerError() *PutDevicesIDRouterOspfInternalServerError {
	return &PutDevicesIDRouterOspfInternalServerError{}
}

/* PutDevicesIDRouterOspfInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutDevicesIDRouterOspfInternalServerError struct {
	Payload *models.Error
}

func (o *PutDevicesIDRouterOspfInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/router/ospf][%d] putDevicesIdRouterOspfInternalServerError  %+v", 500, o.Payload)
}
func (o *PutDevicesIDRouterOspfInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDRouterOspfInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
