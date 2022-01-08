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

// GetDevicesToughswitchesIDReader is a Reader for the GetDevicesToughswitchesID structure.
type GetDevicesToughswitchesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesToughswitchesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesToughswitchesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesToughswitchesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesToughswitchesIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesToughswitchesIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesToughswitchesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesToughswitchesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDevicesToughswitchesIDOK creates a GetDevicesToughswitchesIDOK with default headers values
func NewGetDevicesToughswitchesIDOK() *GetDevicesToughswitchesIDOK {
	return &GetDevicesToughswitchesIDOK{}
}

/* GetDevicesToughswitchesIDOK describes a response with status code 200, with default header values.

Successful
*/
type GetDevicesToughswitchesIDOK struct {
	Payload *models.DeviceStatusOverview
}

func (o *GetDevicesToughswitchesIDOK) Error() string {
	return fmt.Sprintf("[GET /devices/toughswitches/{id}][%d] getDevicesToughswitchesIdOK  %+v", 200, o.Payload)
}
func (o *GetDevicesToughswitchesIDOK) GetPayload() *models.DeviceStatusOverview {
	return o.Payload
}

func (o *GetDevicesToughswitchesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceStatusOverview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesToughswitchesIDBadRequest creates a GetDevicesToughswitchesIDBadRequest with default headers values
func NewGetDevicesToughswitchesIDBadRequest() *GetDevicesToughswitchesIDBadRequest {
	return &GetDevicesToughswitchesIDBadRequest{}
}

/* GetDevicesToughswitchesIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDevicesToughswitchesIDBadRequest struct {
	Payload *models.Error
}

func (o *GetDevicesToughswitchesIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/toughswitches/{id}][%d] getDevicesToughswitchesIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetDevicesToughswitchesIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesToughswitchesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesToughswitchesIDUnauthorized creates a GetDevicesToughswitchesIDUnauthorized with default headers values
func NewGetDevicesToughswitchesIDUnauthorized() *GetDevicesToughswitchesIDUnauthorized {
	return &GetDevicesToughswitchesIDUnauthorized{}
}

/* GetDevicesToughswitchesIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDevicesToughswitchesIDUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesToughswitchesIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/toughswitches/{id}][%d] getDevicesToughswitchesIdUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDevicesToughswitchesIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesToughswitchesIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesToughswitchesIDForbidden creates a GetDevicesToughswitchesIDForbidden with default headers values
func NewGetDevicesToughswitchesIDForbidden() *GetDevicesToughswitchesIDForbidden {
	return &GetDevicesToughswitchesIDForbidden{}
}

/* GetDevicesToughswitchesIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDevicesToughswitchesIDForbidden struct {
	Payload *models.Error
}

func (o *GetDevicesToughswitchesIDForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/toughswitches/{id}][%d] getDevicesToughswitchesIdForbidden  %+v", 403, o.Payload)
}
func (o *GetDevicesToughswitchesIDForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesToughswitchesIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesToughswitchesIDNotFound creates a GetDevicesToughswitchesIDNotFound with default headers values
func NewGetDevicesToughswitchesIDNotFound() *GetDevicesToughswitchesIDNotFound {
	return &GetDevicesToughswitchesIDNotFound{}
}

/* GetDevicesToughswitchesIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDevicesToughswitchesIDNotFound struct {
	Payload *models.Error
}

func (o *GetDevicesToughswitchesIDNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/toughswitches/{id}][%d] getDevicesToughswitchesIdNotFound  %+v", 404, o.Payload)
}
func (o *GetDevicesToughswitchesIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesToughswitchesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesToughswitchesIDInternalServerError creates a GetDevicesToughswitchesIDInternalServerError with default headers values
func NewGetDevicesToughswitchesIDInternalServerError() *GetDevicesToughswitchesIDInternalServerError {
	return &GetDevicesToughswitchesIDInternalServerError{}
}

/* GetDevicesToughswitchesIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDevicesToughswitchesIDInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesToughswitchesIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/toughswitches/{id}][%d] getDevicesToughswitchesIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDevicesToughswitchesIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesToughswitchesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}