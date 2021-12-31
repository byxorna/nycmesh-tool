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

// GetDevicesEroutersIDReader is a Reader for the GetDevicesEroutersID structure.
type GetDevicesEroutersIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesEroutersIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesEroutersIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesEroutersIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesEroutersIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesEroutersIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesEroutersIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesEroutersIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDevicesEroutersIDOK creates a GetDevicesEroutersIDOK with default headers values
func NewGetDevicesEroutersIDOK() *GetDevicesEroutersIDOK {
	return &GetDevicesEroutersIDOK{}
}

/* GetDevicesEroutersIDOK describes a response with status code 200, with default header values.

Successful
*/
type GetDevicesEroutersIDOK struct {
	Payload *models.DeviceStatusOverview
}

func (o *GetDevicesEroutersIDOK) Error() string {
	return fmt.Sprintf("[GET /devices/erouters/{id}][%d] getDevicesEroutersIdOK  %+v", 200, o.Payload)
}
func (o *GetDevicesEroutersIDOK) GetPayload() *models.DeviceStatusOverview {
	return o.Payload
}

func (o *GetDevicesEroutersIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceStatusOverview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesEroutersIDBadRequest creates a GetDevicesEroutersIDBadRequest with default headers values
func NewGetDevicesEroutersIDBadRequest() *GetDevicesEroutersIDBadRequest {
	return &GetDevicesEroutersIDBadRequest{}
}

/* GetDevicesEroutersIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDevicesEroutersIDBadRequest struct {
	Payload *models.Error
}

func (o *GetDevicesEroutersIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/erouters/{id}][%d] getDevicesEroutersIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetDevicesEroutersIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesEroutersIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesEroutersIDUnauthorized creates a GetDevicesEroutersIDUnauthorized with default headers values
func NewGetDevicesEroutersIDUnauthorized() *GetDevicesEroutersIDUnauthorized {
	return &GetDevicesEroutersIDUnauthorized{}
}

/* GetDevicesEroutersIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDevicesEroutersIDUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesEroutersIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/erouters/{id}][%d] getDevicesEroutersIdUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDevicesEroutersIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesEroutersIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesEroutersIDForbidden creates a GetDevicesEroutersIDForbidden with default headers values
func NewGetDevicesEroutersIDForbidden() *GetDevicesEroutersIDForbidden {
	return &GetDevicesEroutersIDForbidden{}
}

/* GetDevicesEroutersIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDevicesEroutersIDForbidden struct {
	Payload *models.Error
}

func (o *GetDevicesEroutersIDForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/erouters/{id}][%d] getDevicesEroutersIdForbidden  %+v", 403, o.Payload)
}
func (o *GetDevicesEroutersIDForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesEroutersIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesEroutersIDNotFound creates a GetDevicesEroutersIDNotFound with default headers values
func NewGetDevicesEroutersIDNotFound() *GetDevicesEroutersIDNotFound {
	return &GetDevicesEroutersIDNotFound{}
}

/* GetDevicesEroutersIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDevicesEroutersIDNotFound struct {
	Payload *models.Error
}

func (o *GetDevicesEroutersIDNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/erouters/{id}][%d] getDevicesEroutersIdNotFound  %+v", 404, o.Payload)
}
func (o *GetDevicesEroutersIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesEroutersIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesEroutersIDInternalServerError creates a GetDevicesEroutersIDInternalServerError with default headers values
func NewGetDevicesEroutersIDInternalServerError() *GetDevicesEroutersIDInternalServerError {
	return &GetDevicesEroutersIDInternalServerError{}
}

/* GetDevicesEroutersIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDevicesEroutersIDInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesEroutersIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/erouters/{id}][%d] getDevicesEroutersIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDevicesEroutersIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesEroutersIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
