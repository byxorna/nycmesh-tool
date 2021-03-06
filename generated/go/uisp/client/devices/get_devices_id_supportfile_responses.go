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

// GetDevicesIDSupportfileReader is a Reader for the GetDevicesIDSupportfile structure.
type GetDevicesIDSupportfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesIDSupportfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 401:
		result := NewGetDevicesIDSupportfileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesIDSupportfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesIDSupportfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesIDSupportfileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDevicesIDSupportfileUnauthorized creates a GetDevicesIDSupportfileUnauthorized with default headers values
func NewGetDevicesIDSupportfileUnauthorized() *GetDevicesIDSupportfileUnauthorized {
	return &GetDevicesIDSupportfileUnauthorized{}
}

/* GetDevicesIDSupportfileUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDevicesIDSupportfileUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesIDSupportfileUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/supportfile][%d] getDevicesIdSupportfileUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDevicesIDSupportfileUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesIDSupportfileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesIDSupportfileForbidden creates a GetDevicesIDSupportfileForbidden with default headers values
func NewGetDevicesIDSupportfileForbidden() *GetDevicesIDSupportfileForbidden {
	return &GetDevicesIDSupportfileForbidden{}
}

/* GetDevicesIDSupportfileForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDevicesIDSupportfileForbidden struct {
	Payload *models.Error
}

func (o *GetDevicesIDSupportfileForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/supportfile][%d] getDevicesIdSupportfileForbidden  %+v", 403, o.Payload)
}
func (o *GetDevicesIDSupportfileForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesIDSupportfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesIDSupportfileNotFound creates a GetDevicesIDSupportfileNotFound with default headers values
func NewGetDevicesIDSupportfileNotFound() *GetDevicesIDSupportfileNotFound {
	return &GetDevicesIDSupportfileNotFound{}
}

/* GetDevicesIDSupportfileNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDevicesIDSupportfileNotFound struct {
	Payload *models.Error
}

func (o *GetDevicesIDSupportfileNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/supportfile][%d] getDevicesIdSupportfileNotFound  %+v", 404, o.Payload)
}
func (o *GetDevicesIDSupportfileNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesIDSupportfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesIDSupportfileInternalServerError creates a GetDevicesIDSupportfileInternalServerError with default headers values
func NewGetDevicesIDSupportfileInternalServerError() *GetDevicesIDSupportfileInternalServerError {
	return &GetDevicesIDSupportfileInternalServerError{}
}

/* GetDevicesIDSupportfileInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDevicesIDSupportfileInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesIDSupportfileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/supportfile][%d] getDevicesIdSupportfileInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDevicesIDSupportfileInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesIDSupportfileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
