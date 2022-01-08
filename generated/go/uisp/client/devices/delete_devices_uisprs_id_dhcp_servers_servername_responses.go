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

// DeleteDevicesUisprsIDDhcpServersServernameReader is a Reader for the DeleteDevicesUisprsIDDhcpServersServername structure.
type DeleteDevicesUisprsIDDhcpServersServernameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDevicesUisprsIDDhcpServersServernameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDevicesUisprsIDDhcpServersServernameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteDevicesUisprsIDDhcpServersServernameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteDevicesUisprsIDDhcpServersServernameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteDevicesUisprsIDDhcpServersServernameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDevicesUisprsIDDhcpServersServernameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDevicesUisprsIDDhcpServersServernameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDevicesUisprsIDDhcpServersServernameOK creates a DeleteDevicesUisprsIDDhcpServersServernameOK with default headers values
func NewDeleteDevicesUisprsIDDhcpServersServernameOK() *DeleteDevicesUisprsIDDhcpServersServernameOK {
	return &DeleteDevicesUisprsIDDhcpServersServernameOK{}
}

/* DeleteDevicesUisprsIDDhcpServersServernameOK describes a response with status code 200, with default header values.

Successful
*/
type DeleteDevicesUisprsIDDhcpServersServernameOK struct {
	Payload *models.Status
}

func (o *DeleteDevicesUisprsIDDhcpServersServernameOK) Error() string {
	return fmt.Sprintf("[DELETE /devices/uisprs/{id}/dhcp/servers/{serverName}][%d] deleteDevicesUisprsIdDhcpServersServernameOK  %+v", 200, o.Payload)
}
func (o *DeleteDevicesUisprsIDDhcpServersServernameOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *DeleteDevicesUisprsIDDhcpServersServernameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesUisprsIDDhcpServersServernameBadRequest creates a DeleteDevicesUisprsIDDhcpServersServernameBadRequest with default headers values
func NewDeleteDevicesUisprsIDDhcpServersServernameBadRequest() *DeleteDevicesUisprsIDDhcpServersServernameBadRequest {
	return &DeleteDevicesUisprsIDDhcpServersServernameBadRequest{}
}

/* DeleteDevicesUisprsIDDhcpServersServernameBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteDevicesUisprsIDDhcpServersServernameBadRequest struct {
	Payload *models.Error
}

func (o *DeleteDevicesUisprsIDDhcpServersServernameBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /devices/uisprs/{id}/dhcp/servers/{serverName}][%d] deleteDevicesUisprsIdDhcpServersServernameBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteDevicesUisprsIDDhcpServersServernameBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesUisprsIDDhcpServersServernameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesUisprsIDDhcpServersServernameUnauthorized creates a DeleteDevicesUisprsIDDhcpServersServernameUnauthorized with default headers values
func NewDeleteDevicesUisprsIDDhcpServersServernameUnauthorized() *DeleteDevicesUisprsIDDhcpServersServernameUnauthorized {
	return &DeleteDevicesUisprsIDDhcpServersServernameUnauthorized{}
}

/* DeleteDevicesUisprsIDDhcpServersServernameUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteDevicesUisprsIDDhcpServersServernameUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteDevicesUisprsIDDhcpServersServernameUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /devices/uisprs/{id}/dhcp/servers/{serverName}][%d] deleteDevicesUisprsIdDhcpServersServernameUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteDevicesUisprsIDDhcpServersServernameUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesUisprsIDDhcpServersServernameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesUisprsIDDhcpServersServernameForbidden creates a DeleteDevicesUisprsIDDhcpServersServernameForbidden with default headers values
func NewDeleteDevicesUisprsIDDhcpServersServernameForbidden() *DeleteDevicesUisprsIDDhcpServersServernameForbidden {
	return &DeleteDevicesUisprsIDDhcpServersServernameForbidden{}
}

/* DeleteDevicesUisprsIDDhcpServersServernameForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteDevicesUisprsIDDhcpServersServernameForbidden struct {
	Payload *models.Error
}

func (o *DeleteDevicesUisprsIDDhcpServersServernameForbidden) Error() string {
	return fmt.Sprintf("[DELETE /devices/uisprs/{id}/dhcp/servers/{serverName}][%d] deleteDevicesUisprsIdDhcpServersServernameForbidden  %+v", 403, o.Payload)
}
func (o *DeleteDevicesUisprsIDDhcpServersServernameForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesUisprsIDDhcpServersServernameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesUisprsIDDhcpServersServernameNotFound creates a DeleteDevicesUisprsIDDhcpServersServernameNotFound with default headers values
func NewDeleteDevicesUisprsIDDhcpServersServernameNotFound() *DeleteDevicesUisprsIDDhcpServersServernameNotFound {
	return &DeleteDevicesUisprsIDDhcpServersServernameNotFound{}
}

/* DeleteDevicesUisprsIDDhcpServersServernameNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteDevicesUisprsIDDhcpServersServernameNotFound struct {
	Payload *models.Error
}

func (o *DeleteDevicesUisprsIDDhcpServersServernameNotFound) Error() string {
	return fmt.Sprintf("[DELETE /devices/uisprs/{id}/dhcp/servers/{serverName}][%d] deleteDevicesUisprsIdDhcpServersServernameNotFound  %+v", 404, o.Payload)
}
func (o *DeleteDevicesUisprsIDDhcpServersServernameNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesUisprsIDDhcpServersServernameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesUisprsIDDhcpServersServernameInternalServerError creates a DeleteDevicesUisprsIDDhcpServersServernameInternalServerError with default headers values
func NewDeleteDevicesUisprsIDDhcpServersServernameInternalServerError() *DeleteDevicesUisprsIDDhcpServersServernameInternalServerError {
	return &DeleteDevicesUisprsIDDhcpServersServernameInternalServerError{}
}

/* DeleteDevicesUisprsIDDhcpServersServernameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteDevicesUisprsIDDhcpServersServernameInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteDevicesUisprsIDDhcpServersServernameInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /devices/uisprs/{id}/dhcp/servers/{serverName}][%d] deleteDevicesUisprsIdDhcpServersServernameInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteDevicesUisprsIDDhcpServersServernameInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesUisprsIDDhcpServersServernameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}