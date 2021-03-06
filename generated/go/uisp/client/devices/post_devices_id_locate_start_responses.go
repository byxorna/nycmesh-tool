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

// PostDevicesIDLocateStartReader is a Reader for the PostDevicesIDLocateStart structure.
type PostDevicesIDLocateStartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDevicesIDLocateStartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDevicesIDLocateStartOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDevicesIDLocateStartBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDevicesIDLocateStartUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostDevicesIDLocateStartForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostDevicesIDLocateStartNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDevicesIDLocateStartInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDevicesIDLocateStartOK creates a PostDevicesIDLocateStartOK with default headers values
func NewPostDevicesIDLocateStartOK() *PostDevicesIDLocateStartOK {
	return &PostDevicesIDLocateStartOK{}
}

/* PostDevicesIDLocateStartOK describes a response with status code 200, with default header values.

Successful
*/
type PostDevicesIDLocateStartOK struct {
	Payload *models.Status
}

func (o *PostDevicesIDLocateStartOK) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/locate/start][%d] postDevicesIdLocateStartOK  %+v", 200, o.Payload)
}
func (o *PostDevicesIDLocateStartOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PostDevicesIDLocateStartOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDLocateStartBadRequest creates a PostDevicesIDLocateStartBadRequest with default headers values
func NewPostDevicesIDLocateStartBadRequest() *PostDevicesIDLocateStartBadRequest {
	return &PostDevicesIDLocateStartBadRequest{}
}

/* PostDevicesIDLocateStartBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostDevicesIDLocateStartBadRequest struct {
	Payload *models.Error
}

func (o *PostDevicesIDLocateStartBadRequest) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/locate/start][%d] postDevicesIdLocateStartBadRequest  %+v", 400, o.Payload)
}
func (o *PostDevicesIDLocateStartBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDLocateStartBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDLocateStartUnauthorized creates a PostDevicesIDLocateStartUnauthorized with default headers values
func NewPostDevicesIDLocateStartUnauthorized() *PostDevicesIDLocateStartUnauthorized {
	return &PostDevicesIDLocateStartUnauthorized{}
}

/* PostDevicesIDLocateStartUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostDevicesIDLocateStartUnauthorized struct {
	Payload *models.Error
}

func (o *PostDevicesIDLocateStartUnauthorized) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/locate/start][%d] postDevicesIdLocateStartUnauthorized  %+v", 401, o.Payload)
}
func (o *PostDevicesIDLocateStartUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDLocateStartUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDLocateStartForbidden creates a PostDevicesIDLocateStartForbidden with default headers values
func NewPostDevicesIDLocateStartForbidden() *PostDevicesIDLocateStartForbidden {
	return &PostDevicesIDLocateStartForbidden{}
}

/* PostDevicesIDLocateStartForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostDevicesIDLocateStartForbidden struct {
	Payload *models.Error
}

func (o *PostDevicesIDLocateStartForbidden) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/locate/start][%d] postDevicesIdLocateStartForbidden  %+v", 403, o.Payload)
}
func (o *PostDevicesIDLocateStartForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDLocateStartForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDLocateStartNotFound creates a PostDevicesIDLocateStartNotFound with default headers values
func NewPostDevicesIDLocateStartNotFound() *PostDevicesIDLocateStartNotFound {
	return &PostDevicesIDLocateStartNotFound{}
}

/* PostDevicesIDLocateStartNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostDevicesIDLocateStartNotFound struct {
	Payload *models.Error
}

func (o *PostDevicesIDLocateStartNotFound) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/locate/start][%d] postDevicesIdLocateStartNotFound  %+v", 404, o.Payload)
}
func (o *PostDevicesIDLocateStartNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDLocateStartNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDLocateStartInternalServerError creates a PostDevicesIDLocateStartInternalServerError with default headers values
func NewPostDevicesIDLocateStartInternalServerError() *PostDevicesIDLocateStartInternalServerError {
	return &PostDevicesIDLocateStartInternalServerError{}
}

/* PostDevicesIDLocateStartInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostDevicesIDLocateStartInternalServerError struct {
	Payload *models.Error
}

func (o *PostDevicesIDLocateStartInternalServerError) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/locate/start][%d] postDevicesIdLocateStartInternalServerError  %+v", 500, o.Payload)
}
func (o *PostDevicesIDLocateStartInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDLocateStartInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
