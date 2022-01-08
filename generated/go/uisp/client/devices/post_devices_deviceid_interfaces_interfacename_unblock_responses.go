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

// PostDevicesDeviceidInterfacesInterfacenameUnblockReader is a Reader for the PostDevicesDeviceidInterfacesInterfacenameUnblock structure.
type PostDevicesDeviceidInterfacesInterfacenameUnblockReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDevicesDeviceidInterfacesInterfacenameUnblockReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDevicesDeviceidInterfacesInterfacenameUnblockOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDevicesDeviceidInterfacesInterfacenameUnblockBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDevicesDeviceidInterfacesInterfacenameUnblockUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostDevicesDeviceidInterfacesInterfacenameUnblockForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostDevicesDeviceidInterfacesInterfacenameUnblockNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDevicesDeviceidInterfacesInterfacenameUnblockInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDevicesDeviceidInterfacesInterfacenameUnblockOK creates a PostDevicesDeviceidInterfacesInterfacenameUnblockOK with default headers values
func NewPostDevicesDeviceidInterfacesInterfacenameUnblockOK() *PostDevicesDeviceidInterfacesInterfacenameUnblockOK {
	return &PostDevicesDeviceidInterfacesInterfacenameUnblockOK{}
}

/* PostDevicesDeviceidInterfacesInterfacenameUnblockOK describes a response with status code 200, with default header values.

Successful
*/
type PostDevicesDeviceidInterfacesInterfacenameUnblockOK struct {
	Payload *models.Status
}

func (o *PostDevicesDeviceidInterfacesInterfacenameUnblockOK) Error() string {
	return fmt.Sprintf("[POST /devices/{deviceId}/interfaces/{interfaceName}/unblock][%d] postDevicesDeviceidInterfacesInterfacenameUnblockOK  %+v", 200, o.Payload)
}
func (o *PostDevicesDeviceidInterfacesInterfacenameUnblockOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PostDevicesDeviceidInterfacesInterfacenameUnblockOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesDeviceidInterfacesInterfacenameUnblockBadRequest creates a PostDevicesDeviceidInterfacesInterfacenameUnblockBadRequest with default headers values
func NewPostDevicesDeviceidInterfacesInterfacenameUnblockBadRequest() *PostDevicesDeviceidInterfacesInterfacenameUnblockBadRequest {
	return &PostDevicesDeviceidInterfacesInterfacenameUnblockBadRequest{}
}

/* PostDevicesDeviceidInterfacesInterfacenameUnblockBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostDevicesDeviceidInterfacesInterfacenameUnblockBadRequest struct {
	Payload *models.Error
}

func (o *PostDevicesDeviceidInterfacesInterfacenameUnblockBadRequest) Error() string {
	return fmt.Sprintf("[POST /devices/{deviceId}/interfaces/{interfaceName}/unblock][%d] postDevicesDeviceidInterfacesInterfacenameUnblockBadRequest  %+v", 400, o.Payload)
}
func (o *PostDevicesDeviceidInterfacesInterfacenameUnblockBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesDeviceidInterfacesInterfacenameUnblockBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesDeviceidInterfacesInterfacenameUnblockUnauthorized creates a PostDevicesDeviceidInterfacesInterfacenameUnblockUnauthorized with default headers values
func NewPostDevicesDeviceidInterfacesInterfacenameUnblockUnauthorized() *PostDevicesDeviceidInterfacesInterfacenameUnblockUnauthorized {
	return &PostDevicesDeviceidInterfacesInterfacenameUnblockUnauthorized{}
}

/* PostDevicesDeviceidInterfacesInterfacenameUnblockUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostDevicesDeviceidInterfacesInterfacenameUnblockUnauthorized struct {
	Payload *models.Error
}

func (o *PostDevicesDeviceidInterfacesInterfacenameUnblockUnauthorized) Error() string {
	return fmt.Sprintf("[POST /devices/{deviceId}/interfaces/{interfaceName}/unblock][%d] postDevicesDeviceidInterfacesInterfacenameUnblockUnauthorized  %+v", 401, o.Payload)
}
func (o *PostDevicesDeviceidInterfacesInterfacenameUnblockUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesDeviceidInterfacesInterfacenameUnblockUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesDeviceidInterfacesInterfacenameUnblockForbidden creates a PostDevicesDeviceidInterfacesInterfacenameUnblockForbidden with default headers values
func NewPostDevicesDeviceidInterfacesInterfacenameUnblockForbidden() *PostDevicesDeviceidInterfacesInterfacenameUnblockForbidden {
	return &PostDevicesDeviceidInterfacesInterfacenameUnblockForbidden{}
}

/* PostDevicesDeviceidInterfacesInterfacenameUnblockForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostDevicesDeviceidInterfacesInterfacenameUnblockForbidden struct {
	Payload *models.Error
}

func (o *PostDevicesDeviceidInterfacesInterfacenameUnblockForbidden) Error() string {
	return fmt.Sprintf("[POST /devices/{deviceId}/interfaces/{interfaceName}/unblock][%d] postDevicesDeviceidInterfacesInterfacenameUnblockForbidden  %+v", 403, o.Payload)
}
func (o *PostDevicesDeviceidInterfacesInterfacenameUnblockForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesDeviceidInterfacesInterfacenameUnblockForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesDeviceidInterfacesInterfacenameUnblockNotFound creates a PostDevicesDeviceidInterfacesInterfacenameUnblockNotFound with default headers values
func NewPostDevicesDeviceidInterfacesInterfacenameUnblockNotFound() *PostDevicesDeviceidInterfacesInterfacenameUnblockNotFound {
	return &PostDevicesDeviceidInterfacesInterfacenameUnblockNotFound{}
}

/* PostDevicesDeviceidInterfacesInterfacenameUnblockNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostDevicesDeviceidInterfacesInterfacenameUnblockNotFound struct {
	Payload *models.Error
}

func (o *PostDevicesDeviceidInterfacesInterfacenameUnblockNotFound) Error() string {
	return fmt.Sprintf("[POST /devices/{deviceId}/interfaces/{interfaceName}/unblock][%d] postDevicesDeviceidInterfacesInterfacenameUnblockNotFound  %+v", 404, o.Payload)
}
func (o *PostDevicesDeviceidInterfacesInterfacenameUnblockNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesDeviceidInterfacesInterfacenameUnblockNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesDeviceidInterfacesInterfacenameUnblockInternalServerError creates a PostDevicesDeviceidInterfacesInterfacenameUnblockInternalServerError with default headers values
func NewPostDevicesDeviceidInterfacesInterfacenameUnblockInternalServerError() *PostDevicesDeviceidInterfacesInterfacenameUnblockInternalServerError {
	return &PostDevicesDeviceidInterfacesInterfacenameUnblockInternalServerError{}
}

/* PostDevicesDeviceidInterfacesInterfacenameUnblockInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostDevicesDeviceidInterfacesInterfacenameUnblockInternalServerError struct {
	Payload *models.Error
}

func (o *PostDevicesDeviceidInterfacesInterfacenameUnblockInternalServerError) Error() string {
	return fmt.Sprintf("[POST /devices/{deviceId}/interfaces/{interfaceName}/unblock][%d] postDevicesDeviceidInterfacesInterfacenameUnblockInternalServerError  %+v", 500, o.Payload)
}
func (o *PostDevicesDeviceidInterfacesInterfacenameUnblockInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesDeviceidInterfacesInterfacenameUnblockInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}