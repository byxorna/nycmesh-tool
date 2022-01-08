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

// PutDevicesOnusIDNetworkReader is a Reader for the PutDevicesOnusIDNetwork structure.
type PutDevicesOnusIDNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDevicesOnusIDNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutDevicesOnusIDNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutDevicesOnusIDNetworkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutDevicesOnusIDNetworkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutDevicesOnusIDNetworkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutDevicesOnusIDNetworkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutDevicesOnusIDNetworkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutDevicesOnusIDNetworkOK creates a PutDevicesOnusIDNetworkOK with default headers values
func NewPutDevicesOnusIDNetworkOK() *PutDevicesOnusIDNetworkOK {
	return &PutDevicesOnusIDNetworkOK{}
}

/* PutDevicesOnusIDNetworkOK describes a response with status code 200, with default header values.

Successful
*/
type PutDevicesOnusIDNetworkOK struct {
	Payload *models.OnuNetwork
}

func (o *PutDevicesOnusIDNetworkOK) Error() string {
	return fmt.Sprintf("[PUT /devices/onus/{id}/network][%d] putDevicesOnusIdNetworkOK  %+v", 200, o.Payload)
}
func (o *PutDevicesOnusIDNetworkOK) GetPayload() *models.OnuNetwork {
	return o.Payload
}

func (o *PutDevicesOnusIDNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OnuNetwork)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesOnusIDNetworkBadRequest creates a PutDevicesOnusIDNetworkBadRequest with default headers values
func NewPutDevicesOnusIDNetworkBadRequest() *PutDevicesOnusIDNetworkBadRequest {
	return &PutDevicesOnusIDNetworkBadRequest{}
}

/* PutDevicesOnusIDNetworkBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutDevicesOnusIDNetworkBadRequest struct {
	Payload *models.Error
}

func (o *PutDevicesOnusIDNetworkBadRequest) Error() string {
	return fmt.Sprintf("[PUT /devices/onus/{id}/network][%d] putDevicesOnusIdNetworkBadRequest  %+v", 400, o.Payload)
}
func (o *PutDevicesOnusIDNetworkBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesOnusIDNetworkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesOnusIDNetworkUnauthorized creates a PutDevicesOnusIDNetworkUnauthorized with default headers values
func NewPutDevicesOnusIDNetworkUnauthorized() *PutDevicesOnusIDNetworkUnauthorized {
	return &PutDevicesOnusIDNetworkUnauthorized{}
}

/* PutDevicesOnusIDNetworkUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutDevicesOnusIDNetworkUnauthorized struct {
	Payload *models.Error
}

func (o *PutDevicesOnusIDNetworkUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /devices/onus/{id}/network][%d] putDevicesOnusIdNetworkUnauthorized  %+v", 401, o.Payload)
}
func (o *PutDevicesOnusIDNetworkUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesOnusIDNetworkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesOnusIDNetworkForbidden creates a PutDevicesOnusIDNetworkForbidden with default headers values
func NewPutDevicesOnusIDNetworkForbidden() *PutDevicesOnusIDNetworkForbidden {
	return &PutDevicesOnusIDNetworkForbidden{}
}

/* PutDevicesOnusIDNetworkForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutDevicesOnusIDNetworkForbidden struct {
	Payload *models.Error
}

func (o *PutDevicesOnusIDNetworkForbidden) Error() string {
	return fmt.Sprintf("[PUT /devices/onus/{id}/network][%d] putDevicesOnusIdNetworkForbidden  %+v", 403, o.Payload)
}
func (o *PutDevicesOnusIDNetworkForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesOnusIDNetworkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesOnusIDNetworkNotFound creates a PutDevicesOnusIDNetworkNotFound with default headers values
func NewPutDevicesOnusIDNetworkNotFound() *PutDevicesOnusIDNetworkNotFound {
	return &PutDevicesOnusIDNetworkNotFound{}
}

/* PutDevicesOnusIDNetworkNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutDevicesOnusIDNetworkNotFound struct {
	Payload *models.Error
}

func (o *PutDevicesOnusIDNetworkNotFound) Error() string {
	return fmt.Sprintf("[PUT /devices/onus/{id}/network][%d] putDevicesOnusIdNetworkNotFound  %+v", 404, o.Payload)
}
func (o *PutDevicesOnusIDNetworkNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesOnusIDNetworkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesOnusIDNetworkInternalServerError creates a PutDevicesOnusIDNetworkInternalServerError with default headers values
func NewPutDevicesOnusIDNetworkInternalServerError() *PutDevicesOnusIDNetworkInternalServerError {
	return &PutDevicesOnusIDNetworkInternalServerError{}
}

/* PutDevicesOnusIDNetworkInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutDevicesOnusIDNetworkInternalServerError struct {
	Payload *models.Error
}

func (o *PutDevicesOnusIDNetworkInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /devices/onus/{id}/network][%d] putDevicesOnusIdNetworkInternalServerError  %+v", 500, o.Payload)
}
func (o *PutDevicesOnusIDNetworkInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesOnusIDNetworkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}