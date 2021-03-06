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

// GetDevicesIDReader is a Reader for the GetDevicesID structure.
type GetDevicesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDevicesIDOK creates a GetDevicesIDOK with default headers values
func NewGetDevicesIDOK() *GetDevicesIDOK {
	return &GetDevicesIDOK{}
}

/* GetDevicesIDOK describes a response with status code 200, with default header values.

Successful
*/
type GetDevicesIDOK struct {
	Payload *models.DeviceSchema
}

func (o *GetDevicesIDOK) Error() string {
	return fmt.Sprintf("[GET /devices/{id}][%d] getDevicesIdOK  %+v", 200, o.Payload)
}
func (o *GetDevicesIDOK) GetPayload() *models.DeviceSchema {
	return o.Payload
}

func (o *GetDevicesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceSchema)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesIDBadRequest creates a GetDevicesIDBadRequest with default headers values
func NewGetDevicesIDBadRequest() *GetDevicesIDBadRequest {
	return &GetDevicesIDBadRequest{}
}

/* GetDevicesIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDevicesIDBadRequest struct {
	Payload *models.Error
}

func (o *GetDevicesIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/{id}][%d] getDevicesIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetDevicesIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesIDUnauthorized creates a GetDevicesIDUnauthorized with default headers values
func NewGetDevicesIDUnauthorized() *GetDevicesIDUnauthorized {
	return &GetDevicesIDUnauthorized{}
}

/* GetDevicesIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDevicesIDUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/{id}][%d] getDevicesIdUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDevicesIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesIDForbidden creates a GetDevicesIDForbidden with default headers values
func NewGetDevicesIDForbidden() *GetDevicesIDForbidden {
	return &GetDevicesIDForbidden{}
}

/* GetDevicesIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDevicesIDForbidden struct {
	Payload *models.Error
}

func (o *GetDevicesIDForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/{id}][%d] getDevicesIdForbidden  %+v", 403, o.Payload)
}
func (o *GetDevicesIDForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesIDInternalServerError creates a GetDevicesIDInternalServerError with default headers values
func NewGetDevicesIDInternalServerError() *GetDevicesIDInternalServerError {
	return &GetDevicesIDInternalServerError{}
}

/* GetDevicesIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDevicesIDInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/{id}][%d] getDevicesIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDevicesIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
