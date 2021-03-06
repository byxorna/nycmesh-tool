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

// PostDevicesEroutersIDDhcpServersReader is a Reader for the PostDevicesEroutersIDDhcpServers structure.
type PostDevicesEroutersIDDhcpServersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDevicesEroutersIDDhcpServersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDevicesEroutersIDDhcpServersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDevicesEroutersIDDhcpServersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDevicesEroutersIDDhcpServersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostDevicesEroutersIDDhcpServersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostDevicesEroutersIDDhcpServersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDevicesEroutersIDDhcpServersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDevicesEroutersIDDhcpServersOK creates a PostDevicesEroutersIDDhcpServersOK with default headers values
func NewPostDevicesEroutersIDDhcpServersOK() *PostDevicesEroutersIDDhcpServersOK {
	return &PostDevicesEroutersIDDhcpServersOK{}
}

/* PostDevicesEroutersIDDhcpServersOK describes a response with status code 200, with default header values.

Successful
*/
type PostDevicesEroutersIDDhcpServersOK struct {
	Payload *models.DHCPServer
}

func (o *PostDevicesEroutersIDDhcpServersOK) Error() string {
	return fmt.Sprintf("[POST /devices/erouters/{id}/dhcp/servers][%d] postDevicesEroutersIdDhcpServersOK  %+v", 200, o.Payload)
}
func (o *PostDevicesEroutersIDDhcpServersOK) GetPayload() *models.DHCPServer {
	return o.Payload
}

func (o *PostDevicesEroutersIDDhcpServersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DHCPServer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesEroutersIDDhcpServersBadRequest creates a PostDevicesEroutersIDDhcpServersBadRequest with default headers values
func NewPostDevicesEroutersIDDhcpServersBadRequest() *PostDevicesEroutersIDDhcpServersBadRequest {
	return &PostDevicesEroutersIDDhcpServersBadRequest{}
}

/* PostDevicesEroutersIDDhcpServersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostDevicesEroutersIDDhcpServersBadRequest struct {
	Payload *models.Error
}

func (o *PostDevicesEroutersIDDhcpServersBadRequest) Error() string {
	return fmt.Sprintf("[POST /devices/erouters/{id}/dhcp/servers][%d] postDevicesEroutersIdDhcpServersBadRequest  %+v", 400, o.Payload)
}
func (o *PostDevicesEroutersIDDhcpServersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesEroutersIDDhcpServersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesEroutersIDDhcpServersUnauthorized creates a PostDevicesEroutersIDDhcpServersUnauthorized with default headers values
func NewPostDevicesEroutersIDDhcpServersUnauthorized() *PostDevicesEroutersIDDhcpServersUnauthorized {
	return &PostDevicesEroutersIDDhcpServersUnauthorized{}
}

/* PostDevicesEroutersIDDhcpServersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostDevicesEroutersIDDhcpServersUnauthorized struct {
	Payload *models.Error
}

func (o *PostDevicesEroutersIDDhcpServersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /devices/erouters/{id}/dhcp/servers][%d] postDevicesEroutersIdDhcpServersUnauthorized  %+v", 401, o.Payload)
}
func (o *PostDevicesEroutersIDDhcpServersUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesEroutersIDDhcpServersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesEroutersIDDhcpServersForbidden creates a PostDevicesEroutersIDDhcpServersForbidden with default headers values
func NewPostDevicesEroutersIDDhcpServersForbidden() *PostDevicesEroutersIDDhcpServersForbidden {
	return &PostDevicesEroutersIDDhcpServersForbidden{}
}

/* PostDevicesEroutersIDDhcpServersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostDevicesEroutersIDDhcpServersForbidden struct {
	Payload *models.Error
}

func (o *PostDevicesEroutersIDDhcpServersForbidden) Error() string {
	return fmt.Sprintf("[POST /devices/erouters/{id}/dhcp/servers][%d] postDevicesEroutersIdDhcpServersForbidden  %+v", 403, o.Payload)
}
func (o *PostDevicesEroutersIDDhcpServersForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesEroutersIDDhcpServersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesEroutersIDDhcpServersNotFound creates a PostDevicesEroutersIDDhcpServersNotFound with default headers values
func NewPostDevicesEroutersIDDhcpServersNotFound() *PostDevicesEroutersIDDhcpServersNotFound {
	return &PostDevicesEroutersIDDhcpServersNotFound{}
}

/* PostDevicesEroutersIDDhcpServersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostDevicesEroutersIDDhcpServersNotFound struct {
	Payload *models.Error
}

func (o *PostDevicesEroutersIDDhcpServersNotFound) Error() string {
	return fmt.Sprintf("[POST /devices/erouters/{id}/dhcp/servers][%d] postDevicesEroutersIdDhcpServersNotFound  %+v", 404, o.Payload)
}
func (o *PostDevicesEroutersIDDhcpServersNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesEroutersIDDhcpServersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesEroutersIDDhcpServersInternalServerError creates a PostDevicesEroutersIDDhcpServersInternalServerError with default headers values
func NewPostDevicesEroutersIDDhcpServersInternalServerError() *PostDevicesEroutersIDDhcpServersInternalServerError {
	return &PostDevicesEroutersIDDhcpServersInternalServerError{}
}

/* PostDevicesEroutersIDDhcpServersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostDevicesEroutersIDDhcpServersInternalServerError struct {
	Payload *models.Error
}

func (o *PostDevicesEroutersIDDhcpServersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /devices/erouters/{id}/dhcp/servers][%d] postDevicesEroutersIdDhcpServersInternalServerError  %+v", 500, o.Payload)
}
func (o *PostDevicesEroutersIDDhcpServersInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesEroutersIDDhcpServersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
