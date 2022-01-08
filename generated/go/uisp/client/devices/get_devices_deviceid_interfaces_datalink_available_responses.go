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

// GetDevicesDeviceidInterfacesDatalinkAvailableReader is a Reader for the GetDevicesDeviceidInterfacesDatalinkAvailable structure.
type GetDevicesDeviceidInterfacesDatalinkAvailableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesDeviceidInterfacesDatalinkAvailableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesDeviceidInterfacesDatalinkAvailableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesDeviceidInterfacesDatalinkAvailableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesDeviceidInterfacesDatalinkAvailableUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesDeviceidInterfacesDatalinkAvailableForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesDeviceidInterfacesDatalinkAvailableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesDeviceidInterfacesDatalinkAvailableInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDevicesDeviceidInterfacesDatalinkAvailableOK creates a GetDevicesDeviceidInterfacesDatalinkAvailableOK with default headers values
func NewGetDevicesDeviceidInterfacesDatalinkAvailableOK() *GetDevicesDeviceidInterfacesDatalinkAvailableOK {
	return &GetDevicesDeviceidInterfacesDatalinkAvailableOK{}
}

/* GetDevicesDeviceidInterfacesDatalinkAvailableOK describes a response with status code 200, with default header values.

Successful
*/
type GetDevicesDeviceidInterfacesDatalinkAvailableOK struct {
	Payload *models.ListOfAvailableInterfaces
}

func (o *GetDevicesDeviceidInterfacesDatalinkAvailableOK) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/interfaces/data-link/available][%d] getDevicesDeviceidInterfacesDatalinkAvailableOK  %+v", 200, o.Payload)
}
func (o *GetDevicesDeviceidInterfacesDatalinkAvailableOK) GetPayload() *models.ListOfAvailableInterfaces {
	return o.Payload
}

func (o *GetDevicesDeviceidInterfacesDatalinkAvailableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListOfAvailableInterfaces)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesDeviceidInterfacesDatalinkAvailableBadRequest creates a GetDevicesDeviceidInterfacesDatalinkAvailableBadRequest with default headers values
func NewGetDevicesDeviceidInterfacesDatalinkAvailableBadRequest() *GetDevicesDeviceidInterfacesDatalinkAvailableBadRequest {
	return &GetDevicesDeviceidInterfacesDatalinkAvailableBadRequest{}
}

/* GetDevicesDeviceidInterfacesDatalinkAvailableBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDevicesDeviceidInterfacesDatalinkAvailableBadRequest struct {
	Payload *models.Error
}

func (o *GetDevicesDeviceidInterfacesDatalinkAvailableBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/interfaces/data-link/available][%d] getDevicesDeviceidInterfacesDatalinkAvailableBadRequest  %+v", 400, o.Payload)
}
func (o *GetDevicesDeviceidInterfacesDatalinkAvailableBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesDeviceidInterfacesDatalinkAvailableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesDeviceidInterfacesDatalinkAvailableUnauthorized creates a GetDevicesDeviceidInterfacesDatalinkAvailableUnauthorized with default headers values
func NewGetDevicesDeviceidInterfacesDatalinkAvailableUnauthorized() *GetDevicesDeviceidInterfacesDatalinkAvailableUnauthorized {
	return &GetDevicesDeviceidInterfacesDatalinkAvailableUnauthorized{}
}

/* GetDevicesDeviceidInterfacesDatalinkAvailableUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDevicesDeviceidInterfacesDatalinkAvailableUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesDeviceidInterfacesDatalinkAvailableUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/interfaces/data-link/available][%d] getDevicesDeviceidInterfacesDatalinkAvailableUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDevicesDeviceidInterfacesDatalinkAvailableUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesDeviceidInterfacesDatalinkAvailableUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesDeviceidInterfacesDatalinkAvailableForbidden creates a GetDevicesDeviceidInterfacesDatalinkAvailableForbidden with default headers values
func NewGetDevicesDeviceidInterfacesDatalinkAvailableForbidden() *GetDevicesDeviceidInterfacesDatalinkAvailableForbidden {
	return &GetDevicesDeviceidInterfacesDatalinkAvailableForbidden{}
}

/* GetDevicesDeviceidInterfacesDatalinkAvailableForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDevicesDeviceidInterfacesDatalinkAvailableForbidden struct {
	Payload *models.Error
}

func (o *GetDevicesDeviceidInterfacesDatalinkAvailableForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/interfaces/data-link/available][%d] getDevicesDeviceidInterfacesDatalinkAvailableForbidden  %+v", 403, o.Payload)
}
func (o *GetDevicesDeviceidInterfacesDatalinkAvailableForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesDeviceidInterfacesDatalinkAvailableForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesDeviceidInterfacesDatalinkAvailableNotFound creates a GetDevicesDeviceidInterfacesDatalinkAvailableNotFound with default headers values
func NewGetDevicesDeviceidInterfacesDatalinkAvailableNotFound() *GetDevicesDeviceidInterfacesDatalinkAvailableNotFound {
	return &GetDevicesDeviceidInterfacesDatalinkAvailableNotFound{}
}

/* GetDevicesDeviceidInterfacesDatalinkAvailableNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDevicesDeviceidInterfacesDatalinkAvailableNotFound struct {
	Payload *models.Error
}

func (o *GetDevicesDeviceidInterfacesDatalinkAvailableNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/interfaces/data-link/available][%d] getDevicesDeviceidInterfacesDatalinkAvailableNotFound  %+v", 404, o.Payload)
}
func (o *GetDevicesDeviceidInterfacesDatalinkAvailableNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesDeviceidInterfacesDatalinkAvailableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesDeviceidInterfacesDatalinkAvailableInternalServerError creates a GetDevicesDeviceidInterfacesDatalinkAvailableInternalServerError with default headers values
func NewGetDevicesDeviceidInterfacesDatalinkAvailableInternalServerError() *GetDevicesDeviceidInterfacesDatalinkAvailableInternalServerError {
	return &GetDevicesDeviceidInterfacesDatalinkAvailableInternalServerError{}
}

/* GetDevicesDeviceidInterfacesDatalinkAvailableInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDevicesDeviceidInterfacesDatalinkAvailableInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesDeviceidInterfacesDatalinkAvailableInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/interfaces/data-link/available][%d] getDevicesDeviceidInterfacesDatalinkAvailableInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDevicesDeviceidInterfacesDatalinkAvailableInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesDeviceidInterfacesDatalinkAvailableInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}