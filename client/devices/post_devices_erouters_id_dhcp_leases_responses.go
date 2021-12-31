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

// PostDevicesEroutersIDDhcpLeasesReader is a Reader for the PostDevicesEroutersIDDhcpLeases structure.
type PostDevicesEroutersIDDhcpLeasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDevicesEroutersIDDhcpLeasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDevicesEroutersIDDhcpLeasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDevicesEroutersIDDhcpLeasesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDevicesEroutersIDDhcpLeasesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostDevicesEroutersIDDhcpLeasesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostDevicesEroutersIDDhcpLeasesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDevicesEroutersIDDhcpLeasesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDevicesEroutersIDDhcpLeasesOK creates a PostDevicesEroutersIDDhcpLeasesOK with default headers values
func NewPostDevicesEroutersIDDhcpLeasesOK() *PostDevicesEroutersIDDhcpLeasesOK {
	return &PostDevicesEroutersIDDhcpLeasesOK{}
}

/* PostDevicesEroutersIDDhcpLeasesOK describes a response with status code 200, with default header values.

Successful
*/
type PostDevicesEroutersIDDhcpLeasesOK struct {
	Payload *models.DHCPLease
}

func (o *PostDevicesEroutersIDDhcpLeasesOK) Error() string {
	return fmt.Sprintf("[POST /devices/erouters/{id}/dhcp/leases][%d] postDevicesEroutersIdDhcpLeasesOK  %+v", 200, o.Payload)
}
func (o *PostDevicesEroutersIDDhcpLeasesOK) GetPayload() *models.DHCPLease {
	return o.Payload
}

func (o *PostDevicesEroutersIDDhcpLeasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DHCPLease)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesEroutersIDDhcpLeasesBadRequest creates a PostDevicesEroutersIDDhcpLeasesBadRequest with default headers values
func NewPostDevicesEroutersIDDhcpLeasesBadRequest() *PostDevicesEroutersIDDhcpLeasesBadRequest {
	return &PostDevicesEroutersIDDhcpLeasesBadRequest{}
}

/* PostDevicesEroutersIDDhcpLeasesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostDevicesEroutersIDDhcpLeasesBadRequest struct {
	Payload *models.Error
}

func (o *PostDevicesEroutersIDDhcpLeasesBadRequest) Error() string {
	return fmt.Sprintf("[POST /devices/erouters/{id}/dhcp/leases][%d] postDevicesEroutersIdDhcpLeasesBadRequest  %+v", 400, o.Payload)
}
func (o *PostDevicesEroutersIDDhcpLeasesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesEroutersIDDhcpLeasesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesEroutersIDDhcpLeasesUnauthorized creates a PostDevicesEroutersIDDhcpLeasesUnauthorized with default headers values
func NewPostDevicesEroutersIDDhcpLeasesUnauthorized() *PostDevicesEroutersIDDhcpLeasesUnauthorized {
	return &PostDevicesEroutersIDDhcpLeasesUnauthorized{}
}

/* PostDevicesEroutersIDDhcpLeasesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostDevicesEroutersIDDhcpLeasesUnauthorized struct {
	Payload *models.Error
}

func (o *PostDevicesEroutersIDDhcpLeasesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /devices/erouters/{id}/dhcp/leases][%d] postDevicesEroutersIdDhcpLeasesUnauthorized  %+v", 401, o.Payload)
}
func (o *PostDevicesEroutersIDDhcpLeasesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesEroutersIDDhcpLeasesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesEroutersIDDhcpLeasesForbidden creates a PostDevicesEroutersIDDhcpLeasesForbidden with default headers values
func NewPostDevicesEroutersIDDhcpLeasesForbidden() *PostDevicesEroutersIDDhcpLeasesForbidden {
	return &PostDevicesEroutersIDDhcpLeasesForbidden{}
}

/* PostDevicesEroutersIDDhcpLeasesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostDevicesEroutersIDDhcpLeasesForbidden struct {
	Payload *models.Error
}

func (o *PostDevicesEroutersIDDhcpLeasesForbidden) Error() string {
	return fmt.Sprintf("[POST /devices/erouters/{id}/dhcp/leases][%d] postDevicesEroutersIdDhcpLeasesForbidden  %+v", 403, o.Payload)
}
func (o *PostDevicesEroutersIDDhcpLeasesForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesEroutersIDDhcpLeasesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesEroutersIDDhcpLeasesNotFound creates a PostDevicesEroutersIDDhcpLeasesNotFound with default headers values
func NewPostDevicesEroutersIDDhcpLeasesNotFound() *PostDevicesEroutersIDDhcpLeasesNotFound {
	return &PostDevicesEroutersIDDhcpLeasesNotFound{}
}

/* PostDevicesEroutersIDDhcpLeasesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostDevicesEroutersIDDhcpLeasesNotFound struct {
	Payload *models.Error
}

func (o *PostDevicesEroutersIDDhcpLeasesNotFound) Error() string {
	return fmt.Sprintf("[POST /devices/erouters/{id}/dhcp/leases][%d] postDevicesEroutersIdDhcpLeasesNotFound  %+v", 404, o.Payload)
}
func (o *PostDevicesEroutersIDDhcpLeasesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesEroutersIDDhcpLeasesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesEroutersIDDhcpLeasesInternalServerError creates a PostDevicesEroutersIDDhcpLeasesInternalServerError with default headers values
func NewPostDevicesEroutersIDDhcpLeasesInternalServerError() *PostDevicesEroutersIDDhcpLeasesInternalServerError {
	return &PostDevicesEroutersIDDhcpLeasesInternalServerError{}
}

/* PostDevicesEroutersIDDhcpLeasesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostDevicesEroutersIDDhcpLeasesInternalServerError struct {
	Payload *models.Error
}

func (o *PostDevicesEroutersIDDhcpLeasesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /devices/erouters/{id}/dhcp/leases][%d] postDevicesEroutersIdDhcpLeasesInternalServerError  %+v", 500, o.Payload)
}
func (o *PostDevicesEroutersIDDhcpLeasesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesEroutersIDDhcpLeasesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
