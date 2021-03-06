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

// PostDevicesUisprsIDDhcpLeasesReader is a Reader for the PostDevicesUisprsIDDhcpLeases structure.
type PostDevicesUisprsIDDhcpLeasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDevicesUisprsIDDhcpLeasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDevicesUisprsIDDhcpLeasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDevicesUisprsIDDhcpLeasesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDevicesUisprsIDDhcpLeasesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostDevicesUisprsIDDhcpLeasesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostDevicesUisprsIDDhcpLeasesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDevicesUisprsIDDhcpLeasesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDevicesUisprsIDDhcpLeasesOK creates a PostDevicesUisprsIDDhcpLeasesOK with default headers values
func NewPostDevicesUisprsIDDhcpLeasesOK() *PostDevicesUisprsIDDhcpLeasesOK {
	return &PostDevicesUisprsIDDhcpLeasesOK{}
}

/* PostDevicesUisprsIDDhcpLeasesOK describes a response with status code 200, with default header values.

Successful
*/
type PostDevicesUisprsIDDhcpLeasesOK struct {
	Payload *models.DHCPLease1
}

func (o *PostDevicesUisprsIDDhcpLeasesOK) Error() string {
	return fmt.Sprintf("[POST /devices/uisprs/{id}/dhcp/leases][%d] postDevicesUisprsIdDhcpLeasesOK  %+v", 200, o.Payload)
}
func (o *PostDevicesUisprsIDDhcpLeasesOK) GetPayload() *models.DHCPLease1 {
	return o.Payload
}

func (o *PostDevicesUisprsIDDhcpLeasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DHCPLease1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesUisprsIDDhcpLeasesBadRequest creates a PostDevicesUisprsIDDhcpLeasesBadRequest with default headers values
func NewPostDevicesUisprsIDDhcpLeasesBadRequest() *PostDevicesUisprsIDDhcpLeasesBadRequest {
	return &PostDevicesUisprsIDDhcpLeasesBadRequest{}
}

/* PostDevicesUisprsIDDhcpLeasesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostDevicesUisprsIDDhcpLeasesBadRequest struct {
	Payload *models.Error
}

func (o *PostDevicesUisprsIDDhcpLeasesBadRequest) Error() string {
	return fmt.Sprintf("[POST /devices/uisprs/{id}/dhcp/leases][%d] postDevicesUisprsIdDhcpLeasesBadRequest  %+v", 400, o.Payload)
}
func (o *PostDevicesUisprsIDDhcpLeasesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesUisprsIDDhcpLeasesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesUisprsIDDhcpLeasesUnauthorized creates a PostDevicesUisprsIDDhcpLeasesUnauthorized with default headers values
func NewPostDevicesUisprsIDDhcpLeasesUnauthorized() *PostDevicesUisprsIDDhcpLeasesUnauthorized {
	return &PostDevicesUisprsIDDhcpLeasesUnauthorized{}
}

/* PostDevicesUisprsIDDhcpLeasesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostDevicesUisprsIDDhcpLeasesUnauthorized struct {
	Payload *models.Error
}

func (o *PostDevicesUisprsIDDhcpLeasesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /devices/uisprs/{id}/dhcp/leases][%d] postDevicesUisprsIdDhcpLeasesUnauthorized  %+v", 401, o.Payload)
}
func (o *PostDevicesUisprsIDDhcpLeasesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesUisprsIDDhcpLeasesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesUisprsIDDhcpLeasesForbidden creates a PostDevicesUisprsIDDhcpLeasesForbidden with default headers values
func NewPostDevicesUisprsIDDhcpLeasesForbidden() *PostDevicesUisprsIDDhcpLeasesForbidden {
	return &PostDevicesUisprsIDDhcpLeasesForbidden{}
}

/* PostDevicesUisprsIDDhcpLeasesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostDevicesUisprsIDDhcpLeasesForbidden struct {
	Payload *models.Error
}

func (o *PostDevicesUisprsIDDhcpLeasesForbidden) Error() string {
	return fmt.Sprintf("[POST /devices/uisprs/{id}/dhcp/leases][%d] postDevicesUisprsIdDhcpLeasesForbidden  %+v", 403, o.Payload)
}
func (o *PostDevicesUisprsIDDhcpLeasesForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesUisprsIDDhcpLeasesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesUisprsIDDhcpLeasesNotFound creates a PostDevicesUisprsIDDhcpLeasesNotFound with default headers values
func NewPostDevicesUisprsIDDhcpLeasesNotFound() *PostDevicesUisprsIDDhcpLeasesNotFound {
	return &PostDevicesUisprsIDDhcpLeasesNotFound{}
}

/* PostDevicesUisprsIDDhcpLeasesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostDevicesUisprsIDDhcpLeasesNotFound struct {
	Payload *models.Error
}

func (o *PostDevicesUisprsIDDhcpLeasesNotFound) Error() string {
	return fmt.Sprintf("[POST /devices/uisprs/{id}/dhcp/leases][%d] postDevicesUisprsIdDhcpLeasesNotFound  %+v", 404, o.Payload)
}
func (o *PostDevicesUisprsIDDhcpLeasesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesUisprsIDDhcpLeasesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesUisprsIDDhcpLeasesInternalServerError creates a PostDevicesUisprsIDDhcpLeasesInternalServerError with default headers values
func NewPostDevicesUisprsIDDhcpLeasesInternalServerError() *PostDevicesUisprsIDDhcpLeasesInternalServerError {
	return &PostDevicesUisprsIDDhcpLeasesInternalServerError{}
}

/* PostDevicesUisprsIDDhcpLeasesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostDevicesUisprsIDDhcpLeasesInternalServerError struct {
	Payload *models.Error
}

func (o *PostDevicesUisprsIDDhcpLeasesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /devices/uisprs/{id}/dhcp/leases][%d] postDevicesUisprsIdDhcpLeasesInternalServerError  %+v", 500, o.Payload)
}
func (o *PostDevicesUisprsIDDhcpLeasesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesUisprsIDDhcpLeasesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
