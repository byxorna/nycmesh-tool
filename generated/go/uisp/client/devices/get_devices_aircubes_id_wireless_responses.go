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

// GetDevicesAircubesIDWirelessReader is a Reader for the GetDevicesAircubesIDWireless structure.
type GetDevicesAircubesIDWirelessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesAircubesIDWirelessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesAircubesIDWirelessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesAircubesIDWirelessBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesAircubesIDWirelessUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesAircubesIDWirelessForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesAircubesIDWirelessNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesAircubesIDWirelessInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDevicesAircubesIDWirelessOK creates a GetDevicesAircubesIDWirelessOK with default headers values
func NewGetDevicesAircubesIDWirelessOK() *GetDevicesAircubesIDWirelessOK {
	return &GetDevicesAircubesIDWirelessOK{}
}

/* GetDevicesAircubesIDWirelessOK describes a response with status code 200, with default header values.

Successful
*/
type GetDevicesAircubesIDWirelessOK struct {
	Payload *models.AirCubeWifiConfig
}

func (o *GetDevicesAircubesIDWirelessOK) Error() string {
	return fmt.Sprintf("[GET /devices/aircubes/{id}/wireless][%d] getDevicesAircubesIdWirelessOK  %+v", 200, o.Payload)
}
func (o *GetDevicesAircubesIDWirelessOK) GetPayload() *models.AirCubeWifiConfig {
	return o.Payload
}

func (o *GetDevicesAircubesIDWirelessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AirCubeWifiConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAircubesIDWirelessBadRequest creates a GetDevicesAircubesIDWirelessBadRequest with default headers values
func NewGetDevicesAircubesIDWirelessBadRequest() *GetDevicesAircubesIDWirelessBadRequest {
	return &GetDevicesAircubesIDWirelessBadRequest{}
}

/* GetDevicesAircubesIDWirelessBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDevicesAircubesIDWirelessBadRequest struct {
	Payload *models.Error
}

func (o *GetDevicesAircubesIDWirelessBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/aircubes/{id}/wireless][%d] getDevicesAircubesIdWirelessBadRequest  %+v", 400, o.Payload)
}
func (o *GetDevicesAircubesIDWirelessBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAircubesIDWirelessBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAircubesIDWirelessUnauthorized creates a GetDevicesAircubesIDWirelessUnauthorized with default headers values
func NewGetDevicesAircubesIDWirelessUnauthorized() *GetDevicesAircubesIDWirelessUnauthorized {
	return &GetDevicesAircubesIDWirelessUnauthorized{}
}

/* GetDevicesAircubesIDWirelessUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDevicesAircubesIDWirelessUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesAircubesIDWirelessUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/aircubes/{id}/wireless][%d] getDevicesAircubesIdWirelessUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDevicesAircubesIDWirelessUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAircubesIDWirelessUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAircubesIDWirelessForbidden creates a GetDevicesAircubesIDWirelessForbidden with default headers values
func NewGetDevicesAircubesIDWirelessForbidden() *GetDevicesAircubesIDWirelessForbidden {
	return &GetDevicesAircubesIDWirelessForbidden{}
}

/* GetDevicesAircubesIDWirelessForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDevicesAircubesIDWirelessForbidden struct {
	Payload *models.Error
}

func (o *GetDevicesAircubesIDWirelessForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/aircubes/{id}/wireless][%d] getDevicesAircubesIdWirelessForbidden  %+v", 403, o.Payload)
}
func (o *GetDevicesAircubesIDWirelessForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAircubesIDWirelessForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAircubesIDWirelessNotFound creates a GetDevicesAircubesIDWirelessNotFound with default headers values
func NewGetDevicesAircubesIDWirelessNotFound() *GetDevicesAircubesIDWirelessNotFound {
	return &GetDevicesAircubesIDWirelessNotFound{}
}

/* GetDevicesAircubesIDWirelessNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDevicesAircubesIDWirelessNotFound struct {
	Payload *models.Error
}

func (o *GetDevicesAircubesIDWirelessNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/aircubes/{id}/wireless][%d] getDevicesAircubesIdWirelessNotFound  %+v", 404, o.Payload)
}
func (o *GetDevicesAircubesIDWirelessNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAircubesIDWirelessNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAircubesIDWirelessInternalServerError creates a GetDevicesAircubesIDWirelessInternalServerError with default headers values
func NewGetDevicesAircubesIDWirelessInternalServerError() *GetDevicesAircubesIDWirelessInternalServerError {
	return &GetDevicesAircubesIDWirelessInternalServerError{}
}

/* GetDevicesAircubesIDWirelessInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDevicesAircubesIDWirelessInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesAircubesIDWirelessInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/aircubes/{id}/wireless][%d] getDevicesAircubesIdWirelessInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDevicesAircubesIDWirelessInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAircubesIDWirelessInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
