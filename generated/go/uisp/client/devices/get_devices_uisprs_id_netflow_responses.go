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

// GetDevicesUisprsIDNetflowReader is a Reader for the GetDevicesUisprsIDNetflow structure.
type GetDevicesUisprsIDNetflowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesUisprsIDNetflowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesUisprsIDNetflowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesUisprsIDNetflowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesUisprsIDNetflowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesUisprsIDNetflowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesUisprsIDNetflowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDevicesUisprsIDNetflowOK creates a GetDevicesUisprsIDNetflowOK with default headers values
func NewGetDevicesUisprsIDNetflowOK() *GetDevicesUisprsIDNetflowOK {
	return &GetDevicesUisprsIDNetflowOK{}
}

/* GetDevicesUisprsIDNetflowOK describes a response with status code 200, with default header values.

Successful
*/
type GetDevicesUisprsIDNetflowOK struct {
	Payload *models.Model26
}

func (o *GetDevicesUisprsIDNetflowOK) Error() string {
	return fmt.Sprintf("[GET /devices/uisprs/{id}/netflow][%d] getDevicesUisprsIdNetflowOK  %+v", 200, o.Payload)
}
func (o *GetDevicesUisprsIDNetflowOK) GetPayload() *models.Model26 {
	return o.Payload
}

func (o *GetDevicesUisprsIDNetflowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Model26)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesUisprsIDNetflowBadRequest creates a GetDevicesUisprsIDNetflowBadRequest with default headers values
func NewGetDevicesUisprsIDNetflowBadRequest() *GetDevicesUisprsIDNetflowBadRequest {
	return &GetDevicesUisprsIDNetflowBadRequest{}
}

/* GetDevicesUisprsIDNetflowBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDevicesUisprsIDNetflowBadRequest struct {
	Payload *models.Error
}

func (o *GetDevicesUisprsIDNetflowBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/uisprs/{id}/netflow][%d] getDevicesUisprsIdNetflowBadRequest  %+v", 400, o.Payload)
}
func (o *GetDevicesUisprsIDNetflowBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesUisprsIDNetflowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesUisprsIDNetflowUnauthorized creates a GetDevicesUisprsIDNetflowUnauthorized with default headers values
func NewGetDevicesUisprsIDNetflowUnauthorized() *GetDevicesUisprsIDNetflowUnauthorized {
	return &GetDevicesUisprsIDNetflowUnauthorized{}
}

/* GetDevicesUisprsIDNetflowUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDevicesUisprsIDNetflowUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesUisprsIDNetflowUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/uisprs/{id}/netflow][%d] getDevicesUisprsIdNetflowUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDevicesUisprsIDNetflowUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesUisprsIDNetflowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesUisprsIDNetflowNotFound creates a GetDevicesUisprsIDNetflowNotFound with default headers values
func NewGetDevicesUisprsIDNetflowNotFound() *GetDevicesUisprsIDNetflowNotFound {
	return &GetDevicesUisprsIDNetflowNotFound{}
}

/* GetDevicesUisprsIDNetflowNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDevicesUisprsIDNetflowNotFound struct {
	Payload *models.Error
}

func (o *GetDevicesUisprsIDNetflowNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/uisprs/{id}/netflow][%d] getDevicesUisprsIdNetflowNotFound  %+v", 404, o.Payload)
}
func (o *GetDevicesUisprsIDNetflowNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesUisprsIDNetflowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesUisprsIDNetflowInternalServerError creates a GetDevicesUisprsIDNetflowInternalServerError with default headers values
func NewGetDevicesUisprsIDNetflowInternalServerError() *GetDevicesUisprsIDNetflowInternalServerError {
	return &GetDevicesUisprsIDNetflowInternalServerError{}
}

/* GetDevicesUisprsIDNetflowInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDevicesUisprsIDNetflowInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesUisprsIDNetflowInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/uisprs/{id}/netflow][%d] getDevicesUisprsIdNetflowInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDevicesUisprsIDNetflowInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesUisprsIDNetflowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}