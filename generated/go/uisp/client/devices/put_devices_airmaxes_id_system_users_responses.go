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

// PutDevicesAirmaxesIDSystemUsersReader is a Reader for the PutDevicesAirmaxesIDSystemUsers structure.
type PutDevicesAirmaxesIDSystemUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDevicesAirmaxesIDSystemUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutDevicesAirmaxesIDSystemUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutDevicesAirmaxesIDSystemUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutDevicesAirmaxesIDSystemUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutDevicesAirmaxesIDSystemUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutDevicesAirmaxesIDSystemUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutDevicesAirmaxesIDSystemUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutDevicesAirmaxesIDSystemUsersOK creates a PutDevicesAirmaxesIDSystemUsersOK with default headers values
func NewPutDevicesAirmaxesIDSystemUsersOK() *PutDevicesAirmaxesIDSystemUsersOK {
	return &PutDevicesAirmaxesIDSystemUsersOK{}
}

/* PutDevicesAirmaxesIDSystemUsersOK describes a response with status code 200, with default header values.

Successful
*/
type PutDevicesAirmaxesIDSystemUsersOK struct {
	Payload *models.Model41
}

func (o *PutDevicesAirmaxesIDSystemUsersOK) Error() string {
	return fmt.Sprintf("[PUT /devices/airmaxes/{id}/system/users][%d] putDevicesAirmaxesIdSystemUsersOK  %+v", 200, o.Payload)
}
func (o *PutDevicesAirmaxesIDSystemUsersOK) GetPayload() *models.Model41 {
	return o.Payload
}

func (o *PutDevicesAirmaxesIDSystemUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Model41)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesAirmaxesIDSystemUsersBadRequest creates a PutDevicesAirmaxesIDSystemUsersBadRequest with default headers values
func NewPutDevicesAirmaxesIDSystemUsersBadRequest() *PutDevicesAirmaxesIDSystemUsersBadRequest {
	return &PutDevicesAirmaxesIDSystemUsersBadRequest{}
}

/* PutDevicesAirmaxesIDSystemUsersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutDevicesAirmaxesIDSystemUsersBadRequest struct {
	Payload *models.Error
}

func (o *PutDevicesAirmaxesIDSystemUsersBadRequest) Error() string {
	return fmt.Sprintf("[PUT /devices/airmaxes/{id}/system/users][%d] putDevicesAirmaxesIdSystemUsersBadRequest  %+v", 400, o.Payload)
}
func (o *PutDevicesAirmaxesIDSystemUsersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesAirmaxesIDSystemUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesAirmaxesIDSystemUsersUnauthorized creates a PutDevicesAirmaxesIDSystemUsersUnauthorized with default headers values
func NewPutDevicesAirmaxesIDSystemUsersUnauthorized() *PutDevicesAirmaxesIDSystemUsersUnauthorized {
	return &PutDevicesAirmaxesIDSystemUsersUnauthorized{}
}

/* PutDevicesAirmaxesIDSystemUsersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutDevicesAirmaxesIDSystemUsersUnauthorized struct {
	Payload *models.Error
}

func (o *PutDevicesAirmaxesIDSystemUsersUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /devices/airmaxes/{id}/system/users][%d] putDevicesAirmaxesIdSystemUsersUnauthorized  %+v", 401, o.Payload)
}
func (o *PutDevicesAirmaxesIDSystemUsersUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesAirmaxesIDSystemUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesAirmaxesIDSystemUsersForbidden creates a PutDevicesAirmaxesIDSystemUsersForbidden with default headers values
func NewPutDevicesAirmaxesIDSystemUsersForbidden() *PutDevicesAirmaxesIDSystemUsersForbidden {
	return &PutDevicesAirmaxesIDSystemUsersForbidden{}
}

/* PutDevicesAirmaxesIDSystemUsersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutDevicesAirmaxesIDSystemUsersForbidden struct {
	Payload *models.Error
}

func (o *PutDevicesAirmaxesIDSystemUsersForbidden) Error() string {
	return fmt.Sprintf("[PUT /devices/airmaxes/{id}/system/users][%d] putDevicesAirmaxesIdSystemUsersForbidden  %+v", 403, o.Payload)
}
func (o *PutDevicesAirmaxesIDSystemUsersForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesAirmaxesIDSystemUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesAirmaxesIDSystemUsersNotFound creates a PutDevicesAirmaxesIDSystemUsersNotFound with default headers values
func NewPutDevicesAirmaxesIDSystemUsersNotFound() *PutDevicesAirmaxesIDSystemUsersNotFound {
	return &PutDevicesAirmaxesIDSystemUsersNotFound{}
}

/* PutDevicesAirmaxesIDSystemUsersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutDevicesAirmaxesIDSystemUsersNotFound struct {
	Payload *models.Error
}

func (o *PutDevicesAirmaxesIDSystemUsersNotFound) Error() string {
	return fmt.Sprintf("[PUT /devices/airmaxes/{id}/system/users][%d] putDevicesAirmaxesIdSystemUsersNotFound  %+v", 404, o.Payload)
}
func (o *PutDevicesAirmaxesIDSystemUsersNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesAirmaxesIDSystemUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesAirmaxesIDSystemUsersInternalServerError creates a PutDevicesAirmaxesIDSystemUsersInternalServerError with default headers values
func NewPutDevicesAirmaxesIDSystemUsersInternalServerError() *PutDevicesAirmaxesIDSystemUsersInternalServerError {
	return &PutDevicesAirmaxesIDSystemUsersInternalServerError{}
}

/* PutDevicesAirmaxesIDSystemUsersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutDevicesAirmaxesIDSystemUsersInternalServerError struct {
	Payload *models.Error
}

func (o *PutDevicesAirmaxesIDSystemUsersInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /devices/airmaxes/{id}/system/users][%d] putDevicesAirmaxesIdSystemUsersInternalServerError  %+v", 500, o.Payload)
}
func (o *PutDevicesAirmaxesIDSystemUsersInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesAirmaxesIDSystemUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
