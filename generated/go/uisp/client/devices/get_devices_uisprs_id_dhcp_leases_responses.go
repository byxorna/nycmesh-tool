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

// GetDevicesUisprsIDDhcpLeasesReader is a Reader for the GetDevicesUisprsIDDhcpLeases structure.
type GetDevicesUisprsIDDhcpLeasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesUisprsIDDhcpLeasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesUisprsIDDhcpLeasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesUisprsIDDhcpLeasesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesUisprsIDDhcpLeasesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesUisprsIDDhcpLeasesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesUisprsIDDhcpLeasesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesUisprsIDDhcpLeasesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDevicesUisprsIDDhcpLeasesOK creates a GetDevicesUisprsIDDhcpLeasesOK with default headers values
func NewGetDevicesUisprsIDDhcpLeasesOK() *GetDevicesUisprsIDDhcpLeasesOK {
	return &GetDevicesUisprsIDDhcpLeasesOK{}
}

/* GetDevicesUisprsIDDhcpLeasesOK describes a response with status code 200, with default header values.

Successful
*/
type GetDevicesUisprsIDDhcpLeasesOK struct {
	Payload models.DHCPLeaseList1
}

func (o *GetDevicesUisprsIDDhcpLeasesOK) Error() string {
	return fmt.Sprintf("[GET /devices/uisprs/{id}/dhcp/leases][%d] getDevicesUisprsIdDhcpLeasesOK  %+v", 200, o.Payload)
}
func (o *GetDevicesUisprsIDDhcpLeasesOK) GetPayload() models.DHCPLeaseList1 {
	return o.Payload
}

func (o *GetDevicesUisprsIDDhcpLeasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesUisprsIDDhcpLeasesBadRequest creates a GetDevicesUisprsIDDhcpLeasesBadRequest with default headers values
func NewGetDevicesUisprsIDDhcpLeasesBadRequest() *GetDevicesUisprsIDDhcpLeasesBadRequest {
	return &GetDevicesUisprsIDDhcpLeasesBadRequest{}
}

/* GetDevicesUisprsIDDhcpLeasesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDevicesUisprsIDDhcpLeasesBadRequest struct {
	Payload *models.Error
}

func (o *GetDevicesUisprsIDDhcpLeasesBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/uisprs/{id}/dhcp/leases][%d] getDevicesUisprsIdDhcpLeasesBadRequest  %+v", 400, o.Payload)
}
func (o *GetDevicesUisprsIDDhcpLeasesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesUisprsIDDhcpLeasesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesUisprsIDDhcpLeasesUnauthorized creates a GetDevicesUisprsIDDhcpLeasesUnauthorized with default headers values
func NewGetDevicesUisprsIDDhcpLeasesUnauthorized() *GetDevicesUisprsIDDhcpLeasesUnauthorized {
	return &GetDevicesUisprsIDDhcpLeasesUnauthorized{}
}

/* GetDevicesUisprsIDDhcpLeasesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDevicesUisprsIDDhcpLeasesUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesUisprsIDDhcpLeasesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/uisprs/{id}/dhcp/leases][%d] getDevicesUisprsIdDhcpLeasesUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDevicesUisprsIDDhcpLeasesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesUisprsIDDhcpLeasesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesUisprsIDDhcpLeasesForbidden creates a GetDevicesUisprsIDDhcpLeasesForbidden with default headers values
func NewGetDevicesUisprsIDDhcpLeasesForbidden() *GetDevicesUisprsIDDhcpLeasesForbidden {
	return &GetDevicesUisprsIDDhcpLeasesForbidden{}
}

/* GetDevicesUisprsIDDhcpLeasesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDevicesUisprsIDDhcpLeasesForbidden struct {
	Payload *models.Error
}

func (o *GetDevicesUisprsIDDhcpLeasesForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/uisprs/{id}/dhcp/leases][%d] getDevicesUisprsIdDhcpLeasesForbidden  %+v", 403, o.Payload)
}
func (o *GetDevicesUisprsIDDhcpLeasesForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesUisprsIDDhcpLeasesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesUisprsIDDhcpLeasesNotFound creates a GetDevicesUisprsIDDhcpLeasesNotFound with default headers values
func NewGetDevicesUisprsIDDhcpLeasesNotFound() *GetDevicesUisprsIDDhcpLeasesNotFound {
	return &GetDevicesUisprsIDDhcpLeasesNotFound{}
}

/* GetDevicesUisprsIDDhcpLeasesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDevicesUisprsIDDhcpLeasesNotFound struct {
	Payload *models.Error
}

func (o *GetDevicesUisprsIDDhcpLeasesNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/uisprs/{id}/dhcp/leases][%d] getDevicesUisprsIdDhcpLeasesNotFound  %+v", 404, o.Payload)
}
func (o *GetDevicesUisprsIDDhcpLeasesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesUisprsIDDhcpLeasesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesUisprsIDDhcpLeasesInternalServerError creates a GetDevicesUisprsIDDhcpLeasesInternalServerError with default headers values
func NewGetDevicesUisprsIDDhcpLeasesInternalServerError() *GetDevicesUisprsIDDhcpLeasesInternalServerError {
	return &GetDevicesUisprsIDDhcpLeasesInternalServerError{}
}

/* GetDevicesUisprsIDDhcpLeasesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDevicesUisprsIDDhcpLeasesInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesUisprsIDDhcpLeasesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/uisprs/{id}/dhcp/leases][%d] getDevicesUisprsIdDhcpLeasesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDevicesUisprsIDDhcpLeasesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesUisprsIDDhcpLeasesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
