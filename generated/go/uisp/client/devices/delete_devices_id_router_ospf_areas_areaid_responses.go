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

// DeleteDevicesIDRouterOspfAreasAreaidReader is a Reader for the DeleteDevicesIDRouterOspfAreasAreaid structure.
type DeleteDevicesIDRouterOspfAreasAreaidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDevicesIDRouterOspfAreasAreaidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDevicesIDRouterOspfAreasAreaidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteDevicesIDRouterOspfAreasAreaidBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteDevicesIDRouterOspfAreasAreaidUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteDevicesIDRouterOspfAreasAreaidForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDevicesIDRouterOspfAreasAreaidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDevicesIDRouterOspfAreasAreaidInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDevicesIDRouterOspfAreasAreaidOK creates a DeleteDevicesIDRouterOspfAreasAreaidOK with default headers values
func NewDeleteDevicesIDRouterOspfAreasAreaidOK() *DeleteDevicesIDRouterOspfAreasAreaidOK {
	return &DeleteDevicesIDRouterOspfAreasAreaidOK{}
}

/* DeleteDevicesIDRouterOspfAreasAreaidOK describes a response with status code 200, with default header values.

Successful
*/
type DeleteDevicesIDRouterOspfAreasAreaidOK struct {
	Payload *models.Status
}

func (o *DeleteDevicesIDRouterOspfAreasAreaidOK) Error() string {
	return fmt.Sprintf("[DELETE /devices/{id}/router/ospf/areas/{areaId}][%d] deleteDevicesIdRouterOspfAreasAreaidOK  %+v", 200, o.Payload)
}
func (o *DeleteDevicesIDRouterOspfAreasAreaidOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *DeleteDevicesIDRouterOspfAreasAreaidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesIDRouterOspfAreasAreaidBadRequest creates a DeleteDevicesIDRouterOspfAreasAreaidBadRequest with default headers values
func NewDeleteDevicesIDRouterOspfAreasAreaidBadRequest() *DeleteDevicesIDRouterOspfAreasAreaidBadRequest {
	return &DeleteDevicesIDRouterOspfAreasAreaidBadRequest{}
}

/* DeleteDevicesIDRouterOspfAreasAreaidBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteDevicesIDRouterOspfAreasAreaidBadRequest struct {
	Payload *models.Error
}

func (o *DeleteDevicesIDRouterOspfAreasAreaidBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /devices/{id}/router/ospf/areas/{areaId}][%d] deleteDevicesIdRouterOspfAreasAreaidBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteDevicesIDRouterOspfAreasAreaidBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesIDRouterOspfAreasAreaidBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesIDRouterOspfAreasAreaidUnauthorized creates a DeleteDevicesIDRouterOspfAreasAreaidUnauthorized with default headers values
func NewDeleteDevicesIDRouterOspfAreasAreaidUnauthorized() *DeleteDevicesIDRouterOspfAreasAreaidUnauthorized {
	return &DeleteDevicesIDRouterOspfAreasAreaidUnauthorized{}
}

/* DeleteDevicesIDRouterOspfAreasAreaidUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteDevicesIDRouterOspfAreasAreaidUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteDevicesIDRouterOspfAreasAreaidUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /devices/{id}/router/ospf/areas/{areaId}][%d] deleteDevicesIdRouterOspfAreasAreaidUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteDevicesIDRouterOspfAreasAreaidUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesIDRouterOspfAreasAreaidUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesIDRouterOspfAreasAreaidForbidden creates a DeleteDevicesIDRouterOspfAreasAreaidForbidden with default headers values
func NewDeleteDevicesIDRouterOspfAreasAreaidForbidden() *DeleteDevicesIDRouterOspfAreasAreaidForbidden {
	return &DeleteDevicesIDRouterOspfAreasAreaidForbidden{}
}

/* DeleteDevicesIDRouterOspfAreasAreaidForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteDevicesIDRouterOspfAreasAreaidForbidden struct {
	Payload *models.Error
}

func (o *DeleteDevicesIDRouterOspfAreasAreaidForbidden) Error() string {
	return fmt.Sprintf("[DELETE /devices/{id}/router/ospf/areas/{areaId}][%d] deleteDevicesIdRouterOspfAreasAreaidForbidden  %+v", 403, o.Payload)
}
func (o *DeleteDevicesIDRouterOspfAreasAreaidForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesIDRouterOspfAreasAreaidForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesIDRouterOspfAreasAreaidNotFound creates a DeleteDevicesIDRouterOspfAreasAreaidNotFound with default headers values
func NewDeleteDevicesIDRouterOspfAreasAreaidNotFound() *DeleteDevicesIDRouterOspfAreasAreaidNotFound {
	return &DeleteDevicesIDRouterOspfAreasAreaidNotFound{}
}

/* DeleteDevicesIDRouterOspfAreasAreaidNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteDevicesIDRouterOspfAreasAreaidNotFound struct {
	Payload *models.Error
}

func (o *DeleteDevicesIDRouterOspfAreasAreaidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /devices/{id}/router/ospf/areas/{areaId}][%d] deleteDevicesIdRouterOspfAreasAreaidNotFound  %+v", 404, o.Payload)
}
func (o *DeleteDevicesIDRouterOspfAreasAreaidNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesIDRouterOspfAreasAreaidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesIDRouterOspfAreasAreaidInternalServerError creates a DeleteDevicesIDRouterOspfAreasAreaidInternalServerError with default headers values
func NewDeleteDevicesIDRouterOspfAreasAreaidInternalServerError() *DeleteDevicesIDRouterOspfAreasAreaidInternalServerError {
	return &DeleteDevicesIDRouterOspfAreasAreaidInternalServerError{}
}

/* DeleteDevicesIDRouterOspfAreasAreaidInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteDevicesIDRouterOspfAreasAreaidInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteDevicesIDRouterOspfAreasAreaidInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /devices/{id}/router/ospf/areas/{areaId}][%d] deleteDevicesIdRouterOspfAreasAreaidInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteDevicesIDRouterOspfAreasAreaidInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesIDRouterOspfAreasAreaidInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}