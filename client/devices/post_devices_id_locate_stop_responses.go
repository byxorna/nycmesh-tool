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

// PostDevicesIDLocateStopReader is a Reader for the PostDevicesIDLocateStop structure.
type PostDevicesIDLocateStopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDevicesIDLocateStopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDevicesIDLocateStopOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDevicesIDLocateStopBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDevicesIDLocateStopUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostDevicesIDLocateStopForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostDevicesIDLocateStopNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDevicesIDLocateStopInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDevicesIDLocateStopOK creates a PostDevicesIDLocateStopOK with default headers values
func NewPostDevicesIDLocateStopOK() *PostDevicesIDLocateStopOK {
	return &PostDevicesIDLocateStopOK{}
}

/* PostDevicesIDLocateStopOK describes a response with status code 200, with default header values.

Successful
*/
type PostDevicesIDLocateStopOK struct {
	Payload *models.Status
}

func (o *PostDevicesIDLocateStopOK) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/locate/stop][%d] postDevicesIdLocateStopOK  %+v", 200, o.Payload)
}
func (o *PostDevicesIDLocateStopOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PostDevicesIDLocateStopOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDLocateStopBadRequest creates a PostDevicesIDLocateStopBadRequest with default headers values
func NewPostDevicesIDLocateStopBadRequest() *PostDevicesIDLocateStopBadRequest {
	return &PostDevicesIDLocateStopBadRequest{}
}

/* PostDevicesIDLocateStopBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostDevicesIDLocateStopBadRequest struct {
	Payload *models.Error
}

func (o *PostDevicesIDLocateStopBadRequest) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/locate/stop][%d] postDevicesIdLocateStopBadRequest  %+v", 400, o.Payload)
}
func (o *PostDevicesIDLocateStopBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDLocateStopBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDLocateStopUnauthorized creates a PostDevicesIDLocateStopUnauthorized with default headers values
func NewPostDevicesIDLocateStopUnauthorized() *PostDevicesIDLocateStopUnauthorized {
	return &PostDevicesIDLocateStopUnauthorized{}
}

/* PostDevicesIDLocateStopUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostDevicesIDLocateStopUnauthorized struct {
	Payload *models.Error
}

func (o *PostDevicesIDLocateStopUnauthorized) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/locate/stop][%d] postDevicesIdLocateStopUnauthorized  %+v", 401, o.Payload)
}
func (o *PostDevicesIDLocateStopUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDLocateStopUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDLocateStopForbidden creates a PostDevicesIDLocateStopForbidden with default headers values
func NewPostDevicesIDLocateStopForbidden() *PostDevicesIDLocateStopForbidden {
	return &PostDevicesIDLocateStopForbidden{}
}

/* PostDevicesIDLocateStopForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostDevicesIDLocateStopForbidden struct {
	Payload *models.Error
}

func (o *PostDevicesIDLocateStopForbidden) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/locate/stop][%d] postDevicesIdLocateStopForbidden  %+v", 403, o.Payload)
}
func (o *PostDevicesIDLocateStopForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDLocateStopForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDLocateStopNotFound creates a PostDevicesIDLocateStopNotFound with default headers values
func NewPostDevicesIDLocateStopNotFound() *PostDevicesIDLocateStopNotFound {
	return &PostDevicesIDLocateStopNotFound{}
}

/* PostDevicesIDLocateStopNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostDevicesIDLocateStopNotFound struct {
	Payload *models.Error
}

func (o *PostDevicesIDLocateStopNotFound) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/locate/stop][%d] postDevicesIdLocateStopNotFound  %+v", 404, o.Payload)
}
func (o *PostDevicesIDLocateStopNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDLocateStopNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDLocateStopInternalServerError creates a PostDevicesIDLocateStopInternalServerError with default headers values
func NewPostDevicesIDLocateStopInternalServerError() *PostDevicesIDLocateStopInternalServerError {
	return &PostDevicesIDLocateStopInternalServerError{}
}

/* PostDevicesIDLocateStopInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostDevicesIDLocateStopInternalServerError struct {
	Payload *models.Error
}

func (o *PostDevicesIDLocateStopInternalServerError) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/locate/stop][%d] postDevicesIdLocateStopInternalServerError  %+v", 500, o.Payload)
}
func (o *PostDevicesIDLocateStopInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDLocateStopInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
