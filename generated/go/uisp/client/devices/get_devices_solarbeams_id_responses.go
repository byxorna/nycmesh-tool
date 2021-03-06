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

// GetDevicesSolarbeamsIDReader is a Reader for the GetDevicesSolarbeamsID structure.
type GetDevicesSolarbeamsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesSolarbeamsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesSolarbeamsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesSolarbeamsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesSolarbeamsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesSolarbeamsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesSolarbeamsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesSolarbeamsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDevicesSolarbeamsIDOK creates a GetDevicesSolarbeamsIDOK with default headers values
func NewGetDevicesSolarbeamsIDOK() *GetDevicesSolarbeamsIDOK {
	return &GetDevicesSolarbeamsIDOK{}
}

/* GetDevicesSolarbeamsIDOK describes a response with status code 200, with default header values.

Successful
*/
type GetDevicesSolarbeamsIDOK struct {
	Payload *models.DeviceStatusOverview
}

func (o *GetDevicesSolarbeamsIDOK) Error() string {
	return fmt.Sprintf("[GET /devices/solarbeams/{id}][%d] getDevicesSolarbeamsIdOK  %+v", 200, o.Payload)
}
func (o *GetDevicesSolarbeamsIDOK) GetPayload() *models.DeviceStatusOverview {
	return o.Payload
}

func (o *GetDevicesSolarbeamsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceStatusOverview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesSolarbeamsIDBadRequest creates a GetDevicesSolarbeamsIDBadRequest with default headers values
func NewGetDevicesSolarbeamsIDBadRequest() *GetDevicesSolarbeamsIDBadRequest {
	return &GetDevicesSolarbeamsIDBadRequest{}
}

/* GetDevicesSolarbeamsIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDevicesSolarbeamsIDBadRequest struct {
	Payload *models.Error
}

func (o *GetDevicesSolarbeamsIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/solarbeams/{id}][%d] getDevicesSolarbeamsIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetDevicesSolarbeamsIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesSolarbeamsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesSolarbeamsIDUnauthorized creates a GetDevicesSolarbeamsIDUnauthorized with default headers values
func NewGetDevicesSolarbeamsIDUnauthorized() *GetDevicesSolarbeamsIDUnauthorized {
	return &GetDevicesSolarbeamsIDUnauthorized{}
}

/* GetDevicesSolarbeamsIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDevicesSolarbeamsIDUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesSolarbeamsIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/solarbeams/{id}][%d] getDevicesSolarbeamsIdUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDevicesSolarbeamsIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesSolarbeamsIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesSolarbeamsIDForbidden creates a GetDevicesSolarbeamsIDForbidden with default headers values
func NewGetDevicesSolarbeamsIDForbidden() *GetDevicesSolarbeamsIDForbidden {
	return &GetDevicesSolarbeamsIDForbidden{}
}

/* GetDevicesSolarbeamsIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDevicesSolarbeamsIDForbidden struct {
	Payload *models.Error
}

func (o *GetDevicesSolarbeamsIDForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/solarbeams/{id}][%d] getDevicesSolarbeamsIdForbidden  %+v", 403, o.Payload)
}
func (o *GetDevicesSolarbeamsIDForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesSolarbeamsIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesSolarbeamsIDNotFound creates a GetDevicesSolarbeamsIDNotFound with default headers values
func NewGetDevicesSolarbeamsIDNotFound() *GetDevicesSolarbeamsIDNotFound {
	return &GetDevicesSolarbeamsIDNotFound{}
}

/* GetDevicesSolarbeamsIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDevicesSolarbeamsIDNotFound struct {
	Payload *models.Error
}

func (o *GetDevicesSolarbeamsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/solarbeams/{id}][%d] getDevicesSolarbeamsIdNotFound  %+v", 404, o.Payload)
}
func (o *GetDevicesSolarbeamsIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesSolarbeamsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesSolarbeamsIDInternalServerError creates a GetDevicesSolarbeamsIDInternalServerError with default headers values
func NewGetDevicesSolarbeamsIDInternalServerError() *GetDevicesSolarbeamsIDInternalServerError {
	return &GetDevicesSolarbeamsIDInternalServerError{}
}

/* GetDevicesSolarbeamsIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDevicesSolarbeamsIDInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesSolarbeamsIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/solarbeams/{id}][%d] getDevicesSolarbeamsIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDevicesSolarbeamsIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesSolarbeamsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
