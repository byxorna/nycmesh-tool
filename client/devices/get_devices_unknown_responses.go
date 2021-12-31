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

// GetDevicesUnknownReader is a Reader for the GetDevicesUnknown structure.
type GetDevicesUnknownReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesUnknownReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesUnknownOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesUnknownBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesUnknownUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesUnknownForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesUnknownInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDevicesUnknownOK creates a GetDevicesUnknownOK with default headers values
func NewGetDevicesUnknownOK() *GetDevicesUnknownOK {
	return &GetDevicesUnknownOK{}
}

/* GetDevicesUnknownOK describes a response with status code 200, with default header values.

List of unknown devices.
*/
type GetDevicesUnknownOK struct {
	Payload models.UnknownDevicesList
}

func (o *GetDevicesUnknownOK) Error() string {
	return fmt.Sprintf("[GET /devices/unknown][%d] getDevicesUnknownOK  %+v", 200, o.Payload)
}
func (o *GetDevicesUnknownOK) GetPayload() models.UnknownDevicesList {
	return o.Payload
}

func (o *GetDevicesUnknownOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesUnknownBadRequest creates a GetDevicesUnknownBadRequest with default headers values
func NewGetDevicesUnknownBadRequest() *GetDevicesUnknownBadRequest {
	return &GetDevicesUnknownBadRequest{}
}

/* GetDevicesUnknownBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDevicesUnknownBadRequest struct {
	Payload *models.Error
}

func (o *GetDevicesUnknownBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/unknown][%d] getDevicesUnknownBadRequest  %+v", 400, o.Payload)
}
func (o *GetDevicesUnknownBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesUnknownBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesUnknownUnauthorized creates a GetDevicesUnknownUnauthorized with default headers values
func NewGetDevicesUnknownUnauthorized() *GetDevicesUnknownUnauthorized {
	return &GetDevicesUnknownUnauthorized{}
}

/* GetDevicesUnknownUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDevicesUnknownUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesUnknownUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/unknown][%d] getDevicesUnknownUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDevicesUnknownUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesUnknownUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesUnknownForbidden creates a GetDevicesUnknownForbidden with default headers values
func NewGetDevicesUnknownForbidden() *GetDevicesUnknownForbidden {
	return &GetDevicesUnknownForbidden{}
}

/* GetDevicesUnknownForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDevicesUnknownForbidden struct {
	Payload *models.Error
}

func (o *GetDevicesUnknownForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/unknown][%d] getDevicesUnknownForbidden  %+v", 403, o.Payload)
}
func (o *GetDevicesUnknownForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesUnknownForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesUnknownInternalServerError creates a GetDevicesUnknownInternalServerError with default headers values
func NewGetDevicesUnknownInternalServerError() *GetDevicesUnknownInternalServerError {
	return &GetDevicesUnknownInternalServerError{}
}

/* GetDevicesUnknownInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDevicesUnknownInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesUnknownInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/unknown][%d] getDevicesUnknownInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDevicesUnknownInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesUnknownInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
