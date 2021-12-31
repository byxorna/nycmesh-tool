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

// GetDevicesOnusIDServicesReader is a Reader for the GetDevicesOnusIDServices structure.
type GetDevicesOnusIDServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesOnusIDServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesOnusIDServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesOnusIDServicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesOnusIDServicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesOnusIDServicesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesOnusIDServicesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesOnusIDServicesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDevicesOnusIDServicesOK creates a GetDevicesOnusIDServicesOK with default headers values
func NewGetDevicesOnusIDServicesOK() *GetDevicesOnusIDServicesOK {
	return &GetDevicesOnusIDServicesOK{}
}

/* GetDevicesOnusIDServicesOK describes a response with status code 200, with default header values.

Successful
*/
type GetDevicesOnusIDServicesOK struct {
	Payload *models.Model47
}

func (o *GetDevicesOnusIDServicesOK) Error() string {
	return fmt.Sprintf("[GET /devices/onus/{id}/services][%d] getDevicesOnusIdServicesOK  %+v", 200, o.Payload)
}
func (o *GetDevicesOnusIDServicesOK) GetPayload() *models.Model47 {
	return o.Payload
}

func (o *GetDevicesOnusIDServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Model47)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesOnusIDServicesBadRequest creates a GetDevicesOnusIDServicesBadRequest with default headers values
func NewGetDevicesOnusIDServicesBadRequest() *GetDevicesOnusIDServicesBadRequest {
	return &GetDevicesOnusIDServicesBadRequest{}
}

/* GetDevicesOnusIDServicesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDevicesOnusIDServicesBadRequest struct {
	Payload *models.Error
}

func (o *GetDevicesOnusIDServicesBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/onus/{id}/services][%d] getDevicesOnusIdServicesBadRequest  %+v", 400, o.Payload)
}
func (o *GetDevicesOnusIDServicesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesOnusIDServicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesOnusIDServicesUnauthorized creates a GetDevicesOnusIDServicesUnauthorized with default headers values
func NewGetDevicesOnusIDServicesUnauthorized() *GetDevicesOnusIDServicesUnauthorized {
	return &GetDevicesOnusIDServicesUnauthorized{}
}

/* GetDevicesOnusIDServicesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDevicesOnusIDServicesUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesOnusIDServicesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/onus/{id}/services][%d] getDevicesOnusIdServicesUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDevicesOnusIDServicesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesOnusIDServicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesOnusIDServicesForbidden creates a GetDevicesOnusIDServicesForbidden with default headers values
func NewGetDevicesOnusIDServicesForbidden() *GetDevicesOnusIDServicesForbidden {
	return &GetDevicesOnusIDServicesForbidden{}
}

/* GetDevicesOnusIDServicesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDevicesOnusIDServicesForbidden struct {
	Payload *models.Error
}

func (o *GetDevicesOnusIDServicesForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/onus/{id}/services][%d] getDevicesOnusIdServicesForbidden  %+v", 403, o.Payload)
}
func (o *GetDevicesOnusIDServicesForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesOnusIDServicesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesOnusIDServicesNotFound creates a GetDevicesOnusIDServicesNotFound with default headers values
func NewGetDevicesOnusIDServicesNotFound() *GetDevicesOnusIDServicesNotFound {
	return &GetDevicesOnusIDServicesNotFound{}
}

/* GetDevicesOnusIDServicesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDevicesOnusIDServicesNotFound struct {
	Payload *models.Error
}

func (o *GetDevicesOnusIDServicesNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/onus/{id}/services][%d] getDevicesOnusIdServicesNotFound  %+v", 404, o.Payload)
}
func (o *GetDevicesOnusIDServicesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesOnusIDServicesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesOnusIDServicesInternalServerError creates a GetDevicesOnusIDServicesInternalServerError with default headers values
func NewGetDevicesOnusIDServicesInternalServerError() *GetDevicesOnusIDServicesInternalServerError {
	return &GetDevicesOnusIDServicesInternalServerError{}
}

/* GetDevicesOnusIDServicesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDevicesOnusIDServicesInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesOnusIDServicesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/onus/{id}/services][%d] getDevicesOnusIdServicesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDevicesOnusIDServicesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesOnusIDServicesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
