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

// PutDevicesUisprsIDRouterOspfAreasAreaidReader is a Reader for the PutDevicesUisprsIDRouterOspfAreasAreaid structure.
type PutDevicesUisprsIDRouterOspfAreasAreaidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDevicesUisprsIDRouterOspfAreasAreaidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutDevicesUisprsIDRouterOspfAreasAreaidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutDevicesUisprsIDRouterOspfAreasAreaidBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutDevicesUisprsIDRouterOspfAreasAreaidUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutDevicesUisprsIDRouterOspfAreasAreaidForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutDevicesUisprsIDRouterOspfAreasAreaidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutDevicesUisprsIDRouterOspfAreasAreaidInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutDevicesUisprsIDRouterOspfAreasAreaidOK creates a PutDevicesUisprsIDRouterOspfAreasAreaidOK with default headers values
func NewPutDevicesUisprsIDRouterOspfAreasAreaidOK() *PutDevicesUisprsIDRouterOspfAreasAreaidOK {
	return &PutDevicesUisprsIDRouterOspfAreasAreaidOK{}
}

/* PutDevicesUisprsIDRouterOspfAreasAreaidOK describes a response with status code 200, with default header values.

Successful
*/
type PutDevicesUisprsIDRouterOspfAreasAreaidOK struct {
	Payload *models.Model103
}

func (o *PutDevicesUisprsIDRouterOspfAreasAreaidOK) Error() string {
	return fmt.Sprintf("[PUT /devices/uisprs/{id}/router/ospf/areas/{areaId}][%d] putDevicesUisprsIdRouterOspfAreasAreaidOK  %+v", 200, o.Payload)
}
func (o *PutDevicesUisprsIDRouterOspfAreasAreaidOK) GetPayload() *models.Model103 {
	return o.Payload
}

func (o *PutDevicesUisprsIDRouterOspfAreasAreaidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Model103)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesUisprsIDRouterOspfAreasAreaidBadRequest creates a PutDevicesUisprsIDRouterOspfAreasAreaidBadRequest with default headers values
func NewPutDevicesUisprsIDRouterOspfAreasAreaidBadRequest() *PutDevicesUisprsIDRouterOspfAreasAreaidBadRequest {
	return &PutDevicesUisprsIDRouterOspfAreasAreaidBadRequest{}
}

/* PutDevicesUisprsIDRouterOspfAreasAreaidBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutDevicesUisprsIDRouterOspfAreasAreaidBadRequest struct {
	Payload *models.Error
}

func (o *PutDevicesUisprsIDRouterOspfAreasAreaidBadRequest) Error() string {
	return fmt.Sprintf("[PUT /devices/uisprs/{id}/router/ospf/areas/{areaId}][%d] putDevicesUisprsIdRouterOspfAreasAreaidBadRequest  %+v", 400, o.Payload)
}
func (o *PutDevicesUisprsIDRouterOspfAreasAreaidBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesUisprsIDRouterOspfAreasAreaidBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesUisprsIDRouterOspfAreasAreaidUnauthorized creates a PutDevicesUisprsIDRouterOspfAreasAreaidUnauthorized with default headers values
func NewPutDevicesUisprsIDRouterOspfAreasAreaidUnauthorized() *PutDevicesUisprsIDRouterOspfAreasAreaidUnauthorized {
	return &PutDevicesUisprsIDRouterOspfAreasAreaidUnauthorized{}
}

/* PutDevicesUisprsIDRouterOspfAreasAreaidUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutDevicesUisprsIDRouterOspfAreasAreaidUnauthorized struct {
	Payload *models.Error
}

func (o *PutDevicesUisprsIDRouterOspfAreasAreaidUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /devices/uisprs/{id}/router/ospf/areas/{areaId}][%d] putDevicesUisprsIdRouterOspfAreasAreaidUnauthorized  %+v", 401, o.Payload)
}
func (o *PutDevicesUisprsIDRouterOspfAreasAreaidUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesUisprsIDRouterOspfAreasAreaidUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesUisprsIDRouterOspfAreasAreaidForbidden creates a PutDevicesUisprsIDRouterOspfAreasAreaidForbidden with default headers values
func NewPutDevicesUisprsIDRouterOspfAreasAreaidForbidden() *PutDevicesUisprsIDRouterOspfAreasAreaidForbidden {
	return &PutDevicesUisprsIDRouterOspfAreasAreaidForbidden{}
}

/* PutDevicesUisprsIDRouterOspfAreasAreaidForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutDevicesUisprsIDRouterOspfAreasAreaidForbidden struct {
	Payload *models.Error
}

func (o *PutDevicesUisprsIDRouterOspfAreasAreaidForbidden) Error() string {
	return fmt.Sprintf("[PUT /devices/uisprs/{id}/router/ospf/areas/{areaId}][%d] putDevicesUisprsIdRouterOspfAreasAreaidForbidden  %+v", 403, o.Payload)
}
func (o *PutDevicesUisprsIDRouterOspfAreasAreaidForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesUisprsIDRouterOspfAreasAreaidForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesUisprsIDRouterOspfAreasAreaidNotFound creates a PutDevicesUisprsIDRouterOspfAreasAreaidNotFound with default headers values
func NewPutDevicesUisprsIDRouterOspfAreasAreaidNotFound() *PutDevicesUisprsIDRouterOspfAreasAreaidNotFound {
	return &PutDevicesUisprsIDRouterOspfAreasAreaidNotFound{}
}

/* PutDevicesUisprsIDRouterOspfAreasAreaidNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutDevicesUisprsIDRouterOspfAreasAreaidNotFound struct {
	Payload *models.Error
}

func (o *PutDevicesUisprsIDRouterOspfAreasAreaidNotFound) Error() string {
	return fmt.Sprintf("[PUT /devices/uisprs/{id}/router/ospf/areas/{areaId}][%d] putDevicesUisprsIdRouterOspfAreasAreaidNotFound  %+v", 404, o.Payload)
}
func (o *PutDevicesUisprsIDRouterOspfAreasAreaidNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesUisprsIDRouterOspfAreasAreaidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesUisprsIDRouterOspfAreasAreaidInternalServerError creates a PutDevicesUisprsIDRouterOspfAreasAreaidInternalServerError with default headers values
func NewPutDevicesUisprsIDRouterOspfAreasAreaidInternalServerError() *PutDevicesUisprsIDRouterOspfAreasAreaidInternalServerError {
	return &PutDevicesUisprsIDRouterOspfAreasAreaidInternalServerError{}
}

/* PutDevicesUisprsIDRouterOspfAreasAreaidInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutDevicesUisprsIDRouterOspfAreasAreaidInternalServerError struct {
	Payload *models.Error
}

func (o *PutDevicesUisprsIDRouterOspfAreasAreaidInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /devices/uisprs/{id}/router/ospf/areas/{areaId}][%d] putDevicesUisprsIdRouterOspfAreasAreaidInternalServerError  %+v", 500, o.Payload)
}
func (o *PutDevicesUisprsIDRouterOspfAreasAreaidInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesUisprsIDRouterOspfAreasAreaidInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
