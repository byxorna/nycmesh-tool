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

// DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidReader is a Reader for the DeleteDevicesEroutersIDDhcpLeasesServernameLeaseid structure.
type DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDevicesEroutersIDDhcpLeasesServernameLeaseidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteDevicesEroutersIDDhcpLeasesServernameLeaseidBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteDevicesEroutersIDDhcpLeasesServernameLeaseidUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteDevicesEroutersIDDhcpLeasesServernameLeaseidForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDevicesEroutersIDDhcpLeasesServernameLeaseidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDevicesEroutersIDDhcpLeasesServernameLeaseidInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDevicesEroutersIDDhcpLeasesServernameLeaseidOK creates a DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidOK with default headers values
func NewDeleteDevicesEroutersIDDhcpLeasesServernameLeaseidOK() *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidOK {
	return &DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidOK{}
}

/* DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidOK describes a response with status code 200, with default header values.

Successful
*/
type DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidOK struct {
	Payload *models.Status
}

func (o *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidOK) Error() string {
	return fmt.Sprintf("[DELETE /devices/erouters/{id}/dhcp/leases/{serverName}/{leaseId}][%d] deleteDevicesEroutersIdDhcpLeasesServernameLeaseidOK  %+v", 200, o.Payload)
}
func (o *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesEroutersIDDhcpLeasesServernameLeaseidBadRequest creates a DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidBadRequest with default headers values
func NewDeleteDevicesEroutersIDDhcpLeasesServernameLeaseidBadRequest() *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidBadRequest {
	return &DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidBadRequest{}
}

/* DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidBadRequest struct {
	Payload *models.Error
}

func (o *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /devices/erouters/{id}/dhcp/leases/{serverName}/{leaseId}][%d] deleteDevicesEroutersIdDhcpLeasesServernameLeaseidBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesEroutersIDDhcpLeasesServernameLeaseidUnauthorized creates a DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidUnauthorized with default headers values
func NewDeleteDevicesEroutersIDDhcpLeasesServernameLeaseidUnauthorized() *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidUnauthorized {
	return &DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidUnauthorized{}
}

/* DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /devices/erouters/{id}/dhcp/leases/{serverName}/{leaseId}][%d] deleteDevicesEroutersIdDhcpLeasesServernameLeaseidUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesEroutersIDDhcpLeasesServernameLeaseidForbidden creates a DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidForbidden with default headers values
func NewDeleteDevicesEroutersIDDhcpLeasesServernameLeaseidForbidden() *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidForbidden {
	return &DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidForbidden{}
}

/* DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidForbidden struct {
	Payload *models.Error
}

func (o *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidForbidden) Error() string {
	return fmt.Sprintf("[DELETE /devices/erouters/{id}/dhcp/leases/{serverName}/{leaseId}][%d] deleteDevicesEroutersIdDhcpLeasesServernameLeaseidForbidden  %+v", 403, o.Payload)
}
func (o *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesEroutersIDDhcpLeasesServernameLeaseidNotFound creates a DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidNotFound with default headers values
func NewDeleteDevicesEroutersIDDhcpLeasesServernameLeaseidNotFound() *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidNotFound {
	return &DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidNotFound{}
}

/* DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidNotFound struct {
	Payload *models.Error
}

func (o *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /devices/erouters/{id}/dhcp/leases/{serverName}/{leaseId}][%d] deleteDevicesEroutersIdDhcpLeasesServernameLeaseidNotFound  %+v", 404, o.Payload)
}
func (o *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesEroutersIDDhcpLeasesServernameLeaseidInternalServerError creates a DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidInternalServerError with default headers values
func NewDeleteDevicesEroutersIDDhcpLeasesServernameLeaseidInternalServerError() *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidInternalServerError {
	return &DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidInternalServerError{}
}

/* DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /devices/erouters/{id}/dhcp/leases/{serverName}/{leaseId}][%d] deleteDevicesEroutersIdDhcpLeasesServernameLeaseidInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesEroutersIDDhcpLeasesServernameLeaseidInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
