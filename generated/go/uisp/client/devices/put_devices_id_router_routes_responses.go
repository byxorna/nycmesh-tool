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

// PutDevicesIDRouterRoutesReader is a Reader for the PutDevicesIDRouterRoutes structure.
type PutDevicesIDRouterRoutesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDevicesIDRouterRoutesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutDevicesIDRouterRoutesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutDevicesIDRouterRoutesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutDevicesIDRouterRoutesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutDevicesIDRouterRoutesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutDevicesIDRouterRoutesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutDevicesIDRouterRoutesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutDevicesIDRouterRoutesOK creates a PutDevicesIDRouterRoutesOK with default headers values
func NewPutDevicesIDRouterRoutesOK() *PutDevicesIDRouterRoutesOK {
	return &PutDevicesIDRouterRoutesOK{}
}

/* PutDevicesIDRouterRoutesOK describes a response with status code 200, with default header values.

Successful
*/
type PutDevicesIDRouterRoutesOK struct {
	Payload *models.RouterRoute
}

func (o *PutDevicesIDRouterRoutesOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/router/routes][%d] putDevicesIdRouterRoutesOK  %+v", 200, o.Payload)
}
func (o *PutDevicesIDRouterRoutesOK) GetPayload() *models.RouterRoute {
	return o.Payload
}

func (o *PutDevicesIDRouterRoutesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RouterRoute)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDRouterRoutesBadRequest creates a PutDevicesIDRouterRoutesBadRequest with default headers values
func NewPutDevicesIDRouterRoutesBadRequest() *PutDevicesIDRouterRoutesBadRequest {
	return &PutDevicesIDRouterRoutesBadRequest{}
}

/* PutDevicesIDRouterRoutesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutDevicesIDRouterRoutesBadRequest struct {
	Payload *models.Error
}

func (o *PutDevicesIDRouterRoutesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/router/routes][%d] putDevicesIdRouterRoutesBadRequest  %+v", 400, o.Payload)
}
func (o *PutDevicesIDRouterRoutesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDRouterRoutesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDRouterRoutesUnauthorized creates a PutDevicesIDRouterRoutesUnauthorized with default headers values
func NewPutDevicesIDRouterRoutesUnauthorized() *PutDevicesIDRouterRoutesUnauthorized {
	return &PutDevicesIDRouterRoutesUnauthorized{}
}

/* PutDevicesIDRouterRoutesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutDevicesIDRouterRoutesUnauthorized struct {
	Payload *models.Error
}

func (o *PutDevicesIDRouterRoutesUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/router/routes][%d] putDevicesIdRouterRoutesUnauthorized  %+v", 401, o.Payload)
}
func (o *PutDevicesIDRouterRoutesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDRouterRoutesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDRouterRoutesForbidden creates a PutDevicesIDRouterRoutesForbidden with default headers values
func NewPutDevicesIDRouterRoutesForbidden() *PutDevicesIDRouterRoutesForbidden {
	return &PutDevicesIDRouterRoutesForbidden{}
}

/* PutDevicesIDRouterRoutesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutDevicesIDRouterRoutesForbidden struct {
	Payload *models.Error
}

func (o *PutDevicesIDRouterRoutesForbidden) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/router/routes][%d] putDevicesIdRouterRoutesForbidden  %+v", 403, o.Payload)
}
func (o *PutDevicesIDRouterRoutesForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDRouterRoutesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDRouterRoutesNotFound creates a PutDevicesIDRouterRoutesNotFound with default headers values
func NewPutDevicesIDRouterRoutesNotFound() *PutDevicesIDRouterRoutesNotFound {
	return &PutDevicesIDRouterRoutesNotFound{}
}

/* PutDevicesIDRouterRoutesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutDevicesIDRouterRoutesNotFound struct {
	Payload *models.Error
}

func (o *PutDevicesIDRouterRoutesNotFound) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/router/routes][%d] putDevicesIdRouterRoutesNotFound  %+v", 404, o.Payload)
}
func (o *PutDevicesIDRouterRoutesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDRouterRoutesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDRouterRoutesInternalServerError creates a PutDevicesIDRouterRoutesInternalServerError with default headers values
func NewPutDevicesIDRouterRoutesInternalServerError() *PutDevicesIDRouterRoutesInternalServerError {
	return &PutDevicesIDRouterRoutesInternalServerError{}
}

/* PutDevicesIDRouterRoutesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutDevicesIDRouterRoutesInternalServerError struct {
	Payload *models.Error
}

func (o *PutDevicesIDRouterRoutesInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/router/routes][%d] putDevicesIdRouterRoutesInternalServerError  %+v", 500, o.Payload)
}
func (o *PutDevicesIDRouterRoutesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDRouterRoutesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
