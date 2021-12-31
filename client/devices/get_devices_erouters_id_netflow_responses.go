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

// GetDevicesEroutersIDNetflowReader is a Reader for the GetDevicesEroutersIDNetflow structure.
type GetDevicesEroutersIDNetflowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesEroutersIDNetflowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesEroutersIDNetflowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesEroutersIDNetflowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesEroutersIDNetflowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesEroutersIDNetflowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesEroutersIDNetflowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDevicesEroutersIDNetflowOK creates a GetDevicesEroutersIDNetflowOK with default headers values
func NewGetDevicesEroutersIDNetflowOK() *GetDevicesEroutersIDNetflowOK {
	return &GetDevicesEroutersIDNetflowOK{}
}

/* GetDevicesEroutersIDNetflowOK describes a response with status code 200, with default header values.

Successful
*/
type GetDevicesEroutersIDNetflowOK struct {
	Payload *models.Model26
}

func (o *GetDevicesEroutersIDNetflowOK) Error() string {
	return fmt.Sprintf("[GET /devices/erouters/{id}/netflow][%d] getDevicesEroutersIdNetflowOK  %+v", 200, o.Payload)
}
func (o *GetDevicesEroutersIDNetflowOK) GetPayload() *models.Model26 {
	return o.Payload
}

func (o *GetDevicesEroutersIDNetflowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Model26)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesEroutersIDNetflowBadRequest creates a GetDevicesEroutersIDNetflowBadRequest with default headers values
func NewGetDevicesEroutersIDNetflowBadRequest() *GetDevicesEroutersIDNetflowBadRequest {
	return &GetDevicesEroutersIDNetflowBadRequest{}
}

/* GetDevicesEroutersIDNetflowBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDevicesEroutersIDNetflowBadRequest struct {
	Payload *models.Error
}

func (o *GetDevicesEroutersIDNetflowBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/erouters/{id}/netflow][%d] getDevicesEroutersIdNetflowBadRequest  %+v", 400, o.Payload)
}
func (o *GetDevicesEroutersIDNetflowBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesEroutersIDNetflowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesEroutersIDNetflowUnauthorized creates a GetDevicesEroutersIDNetflowUnauthorized with default headers values
func NewGetDevicesEroutersIDNetflowUnauthorized() *GetDevicesEroutersIDNetflowUnauthorized {
	return &GetDevicesEroutersIDNetflowUnauthorized{}
}

/* GetDevicesEroutersIDNetflowUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDevicesEroutersIDNetflowUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesEroutersIDNetflowUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/erouters/{id}/netflow][%d] getDevicesEroutersIdNetflowUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDevicesEroutersIDNetflowUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesEroutersIDNetflowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesEroutersIDNetflowNotFound creates a GetDevicesEroutersIDNetflowNotFound with default headers values
func NewGetDevicesEroutersIDNetflowNotFound() *GetDevicesEroutersIDNetflowNotFound {
	return &GetDevicesEroutersIDNetflowNotFound{}
}

/* GetDevicesEroutersIDNetflowNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDevicesEroutersIDNetflowNotFound struct {
	Payload *models.Error
}

func (o *GetDevicesEroutersIDNetflowNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/erouters/{id}/netflow][%d] getDevicesEroutersIdNetflowNotFound  %+v", 404, o.Payload)
}
func (o *GetDevicesEroutersIDNetflowNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesEroutersIDNetflowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesEroutersIDNetflowInternalServerError creates a GetDevicesEroutersIDNetflowInternalServerError with default headers values
func NewGetDevicesEroutersIDNetflowInternalServerError() *GetDevicesEroutersIDNetflowInternalServerError {
	return &GetDevicesEroutersIDNetflowInternalServerError{}
}

/* GetDevicesEroutersIDNetflowInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDevicesEroutersIDNetflowInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesEroutersIDNetflowInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/erouters/{id}/netflow][%d] getDevicesEroutersIdNetflowInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDevicesEroutersIDNetflowInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesEroutersIDNetflowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
