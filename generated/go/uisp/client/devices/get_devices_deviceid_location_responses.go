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

// GetDevicesDeviceidLocationReader is a Reader for the GetDevicesDeviceidLocation structure.
type GetDevicesDeviceidLocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesDeviceidLocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesDeviceidLocationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesDeviceidLocationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesDeviceidLocationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesDeviceidLocationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesDeviceidLocationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesDeviceidLocationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDevicesDeviceidLocationOK creates a GetDevicesDeviceidLocationOK with default headers values
func NewGetDevicesDeviceidLocationOK() *GetDevicesDeviceidLocationOK {
	return &GetDevicesDeviceidLocationOK{}
}

/* GetDevicesDeviceidLocationOK describes a response with status code 200, with default header values.

Successful
*/
type GetDevicesDeviceidLocationOK struct {
	Payload *models.DeviceLocation
}

func (o *GetDevicesDeviceidLocationOK) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/location][%d] getDevicesDeviceidLocationOK  %+v", 200, o.Payload)
}
func (o *GetDevicesDeviceidLocationOK) GetPayload() *models.DeviceLocation {
	return o.Payload
}

func (o *GetDevicesDeviceidLocationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceLocation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesDeviceidLocationBadRequest creates a GetDevicesDeviceidLocationBadRequest with default headers values
func NewGetDevicesDeviceidLocationBadRequest() *GetDevicesDeviceidLocationBadRequest {
	return &GetDevicesDeviceidLocationBadRequest{}
}

/* GetDevicesDeviceidLocationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDevicesDeviceidLocationBadRequest struct {
	Payload *models.Error
}

func (o *GetDevicesDeviceidLocationBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/location][%d] getDevicesDeviceidLocationBadRequest  %+v", 400, o.Payload)
}
func (o *GetDevicesDeviceidLocationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesDeviceidLocationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesDeviceidLocationUnauthorized creates a GetDevicesDeviceidLocationUnauthorized with default headers values
func NewGetDevicesDeviceidLocationUnauthorized() *GetDevicesDeviceidLocationUnauthorized {
	return &GetDevicesDeviceidLocationUnauthorized{}
}

/* GetDevicesDeviceidLocationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDevicesDeviceidLocationUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesDeviceidLocationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/location][%d] getDevicesDeviceidLocationUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDevicesDeviceidLocationUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesDeviceidLocationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesDeviceidLocationForbidden creates a GetDevicesDeviceidLocationForbidden with default headers values
func NewGetDevicesDeviceidLocationForbidden() *GetDevicesDeviceidLocationForbidden {
	return &GetDevicesDeviceidLocationForbidden{}
}

/* GetDevicesDeviceidLocationForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDevicesDeviceidLocationForbidden struct {
	Payload *models.Error
}

func (o *GetDevicesDeviceidLocationForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/location][%d] getDevicesDeviceidLocationForbidden  %+v", 403, o.Payload)
}
func (o *GetDevicesDeviceidLocationForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesDeviceidLocationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesDeviceidLocationNotFound creates a GetDevicesDeviceidLocationNotFound with default headers values
func NewGetDevicesDeviceidLocationNotFound() *GetDevicesDeviceidLocationNotFound {
	return &GetDevicesDeviceidLocationNotFound{}
}

/* GetDevicesDeviceidLocationNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDevicesDeviceidLocationNotFound struct {
	Payload *models.Error
}

func (o *GetDevicesDeviceidLocationNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/location][%d] getDevicesDeviceidLocationNotFound  %+v", 404, o.Payload)
}
func (o *GetDevicesDeviceidLocationNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesDeviceidLocationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesDeviceidLocationInternalServerError creates a GetDevicesDeviceidLocationInternalServerError with default headers values
func NewGetDevicesDeviceidLocationInternalServerError() *GetDevicesDeviceidLocationInternalServerError {
	return &GetDevicesDeviceidLocationInternalServerError{}
}

/* GetDevicesDeviceidLocationInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDevicesDeviceidLocationInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesDeviceidLocationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/location][%d] getDevicesDeviceidLocationInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDevicesDeviceidLocationInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesDeviceidLocationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
