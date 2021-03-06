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

// PutDevicesIDRouterOspfAreasAreaidReader is a Reader for the PutDevicesIDRouterOspfAreasAreaid structure.
type PutDevicesIDRouterOspfAreasAreaidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDevicesIDRouterOspfAreasAreaidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutDevicesIDRouterOspfAreasAreaidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutDevicesIDRouterOspfAreasAreaidBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutDevicesIDRouterOspfAreasAreaidUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutDevicesIDRouterOspfAreasAreaidForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutDevicesIDRouterOspfAreasAreaidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutDevicesIDRouterOspfAreasAreaidInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutDevicesIDRouterOspfAreasAreaidOK creates a PutDevicesIDRouterOspfAreasAreaidOK with default headers values
func NewPutDevicesIDRouterOspfAreasAreaidOK() *PutDevicesIDRouterOspfAreasAreaidOK {
	return &PutDevicesIDRouterOspfAreasAreaidOK{}
}

/* PutDevicesIDRouterOspfAreasAreaidOK describes a response with status code 200, with default header values.

Successful
*/
type PutDevicesIDRouterOspfAreasAreaidOK struct {
	Payload *models.Model99
}

func (o *PutDevicesIDRouterOspfAreasAreaidOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/router/ospf/areas/{areaId}][%d] putDevicesIdRouterOspfAreasAreaidOK  %+v", 200, o.Payload)
}
func (o *PutDevicesIDRouterOspfAreasAreaidOK) GetPayload() *models.Model99 {
	return o.Payload
}

func (o *PutDevicesIDRouterOspfAreasAreaidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Model99)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDRouterOspfAreasAreaidBadRequest creates a PutDevicesIDRouterOspfAreasAreaidBadRequest with default headers values
func NewPutDevicesIDRouterOspfAreasAreaidBadRequest() *PutDevicesIDRouterOspfAreasAreaidBadRequest {
	return &PutDevicesIDRouterOspfAreasAreaidBadRequest{}
}

/* PutDevicesIDRouterOspfAreasAreaidBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutDevicesIDRouterOspfAreasAreaidBadRequest struct {
	Payload *models.Error
}

func (o *PutDevicesIDRouterOspfAreasAreaidBadRequest) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/router/ospf/areas/{areaId}][%d] putDevicesIdRouterOspfAreasAreaidBadRequest  %+v", 400, o.Payload)
}
func (o *PutDevicesIDRouterOspfAreasAreaidBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDRouterOspfAreasAreaidBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDRouterOspfAreasAreaidUnauthorized creates a PutDevicesIDRouterOspfAreasAreaidUnauthorized with default headers values
func NewPutDevicesIDRouterOspfAreasAreaidUnauthorized() *PutDevicesIDRouterOspfAreasAreaidUnauthorized {
	return &PutDevicesIDRouterOspfAreasAreaidUnauthorized{}
}

/* PutDevicesIDRouterOspfAreasAreaidUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutDevicesIDRouterOspfAreasAreaidUnauthorized struct {
	Payload *models.Error
}

func (o *PutDevicesIDRouterOspfAreasAreaidUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/router/ospf/areas/{areaId}][%d] putDevicesIdRouterOspfAreasAreaidUnauthorized  %+v", 401, o.Payload)
}
func (o *PutDevicesIDRouterOspfAreasAreaidUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDRouterOspfAreasAreaidUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDRouterOspfAreasAreaidForbidden creates a PutDevicesIDRouterOspfAreasAreaidForbidden with default headers values
func NewPutDevicesIDRouterOspfAreasAreaidForbidden() *PutDevicesIDRouterOspfAreasAreaidForbidden {
	return &PutDevicesIDRouterOspfAreasAreaidForbidden{}
}

/* PutDevicesIDRouterOspfAreasAreaidForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutDevicesIDRouterOspfAreasAreaidForbidden struct {
	Payload *models.Error
}

func (o *PutDevicesIDRouterOspfAreasAreaidForbidden) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/router/ospf/areas/{areaId}][%d] putDevicesIdRouterOspfAreasAreaidForbidden  %+v", 403, o.Payload)
}
func (o *PutDevicesIDRouterOspfAreasAreaidForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDRouterOspfAreasAreaidForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDRouterOspfAreasAreaidNotFound creates a PutDevicesIDRouterOspfAreasAreaidNotFound with default headers values
func NewPutDevicesIDRouterOspfAreasAreaidNotFound() *PutDevicesIDRouterOspfAreasAreaidNotFound {
	return &PutDevicesIDRouterOspfAreasAreaidNotFound{}
}

/* PutDevicesIDRouterOspfAreasAreaidNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutDevicesIDRouterOspfAreasAreaidNotFound struct {
	Payload *models.Error
}

func (o *PutDevicesIDRouterOspfAreasAreaidNotFound) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/router/ospf/areas/{areaId}][%d] putDevicesIdRouterOspfAreasAreaidNotFound  %+v", 404, o.Payload)
}
func (o *PutDevicesIDRouterOspfAreasAreaidNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDRouterOspfAreasAreaidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDRouterOspfAreasAreaidInternalServerError creates a PutDevicesIDRouterOspfAreasAreaidInternalServerError with default headers values
func NewPutDevicesIDRouterOspfAreasAreaidInternalServerError() *PutDevicesIDRouterOspfAreasAreaidInternalServerError {
	return &PutDevicesIDRouterOspfAreasAreaidInternalServerError{}
}

/* PutDevicesIDRouterOspfAreasAreaidInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutDevicesIDRouterOspfAreasAreaidInternalServerError struct {
	Payload *models.Error
}

func (o *PutDevicesIDRouterOspfAreasAreaidInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/router/ospf/areas/{areaId}][%d] putDevicesIdRouterOspfAreasAreaidInternalServerError  %+v", 500, o.Payload)
}
func (o *PutDevicesIDRouterOspfAreasAreaidInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDRouterOspfAreasAreaidInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
