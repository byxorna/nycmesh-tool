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

// PostDevicesRefreshReader is a Reader for the PostDevicesRefresh structure.
type PostDevicesRefreshReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDevicesRefreshReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDevicesRefreshOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDevicesRefreshBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDevicesRefreshUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostDevicesRefreshForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDevicesRefreshInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDevicesRefreshOK creates a PostDevicesRefreshOK with default headers values
func NewPostDevicesRefreshOK() *PostDevicesRefreshOK {
	return &PostDevicesRefreshOK{}
}

/* PostDevicesRefreshOK describes a response with status code 200, with default header values.

Successful
*/
type PostDevicesRefreshOK struct {
	Payload *models.Status
}

func (o *PostDevicesRefreshOK) Error() string {
	return fmt.Sprintf("[POST /devices/refresh][%d] postDevicesRefreshOK  %+v", 200, o.Payload)
}
func (o *PostDevicesRefreshOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PostDevicesRefreshOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesRefreshBadRequest creates a PostDevicesRefreshBadRequest with default headers values
func NewPostDevicesRefreshBadRequest() *PostDevicesRefreshBadRequest {
	return &PostDevicesRefreshBadRequest{}
}

/* PostDevicesRefreshBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostDevicesRefreshBadRequest struct {
	Payload *models.Error
}

func (o *PostDevicesRefreshBadRequest) Error() string {
	return fmt.Sprintf("[POST /devices/refresh][%d] postDevicesRefreshBadRequest  %+v", 400, o.Payload)
}
func (o *PostDevicesRefreshBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesRefreshBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesRefreshUnauthorized creates a PostDevicesRefreshUnauthorized with default headers values
func NewPostDevicesRefreshUnauthorized() *PostDevicesRefreshUnauthorized {
	return &PostDevicesRefreshUnauthorized{}
}

/* PostDevicesRefreshUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostDevicesRefreshUnauthorized struct {
	Payload *models.Error
}

func (o *PostDevicesRefreshUnauthorized) Error() string {
	return fmt.Sprintf("[POST /devices/refresh][%d] postDevicesRefreshUnauthorized  %+v", 401, o.Payload)
}
func (o *PostDevicesRefreshUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesRefreshUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesRefreshForbidden creates a PostDevicesRefreshForbidden with default headers values
func NewPostDevicesRefreshForbidden() *PostDevicesRefreshForbidden {
	return &PostDevicesRefreshForbidden{}
}

/* PostDevicesRefreshForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostDevicesRefreshForbidden struct {
	Payload *models.Error
}

func (o *PostDevicesRefreshForbidden) Error() string {
	return fmt.Sprintf("[POST /devices/refresh][%d] postDevicesRefreshForbidden  %+v", 403, o.Payload)
}
func (o *PostDevicesRefreshForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesRefreshForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesRefreshInternalServerError creates a PostDevicesRefreshInternalServerError with default headers values
func NewPostDevicesRefreshInternalServerError() *PostDevicesRefreshInternalServerError {
	return &PostDevicesRefreshInternalServerError{}
}

/* PostDevicesRefreshInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostDevicesRefreshInternalServerError struct {
	Payload *models.Error
}

func (o *PostDevicesRefreshInternalServerError) Error() string {
	return fmt.Sprintf("[POST /devices/refresh][%d] postDevicesRefreshInternalServerError  %+v", 500, o.Payload)
}
func (o *PostDevicesRefreshInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesRefreshInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}