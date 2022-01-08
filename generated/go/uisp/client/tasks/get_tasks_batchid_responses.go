// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// GetTasksBatchidReader is a Reader for the GetTasksBatchid structure.
type GetTasksBatchidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTasksBatchidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTasksBatchidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTasksBatchidBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTasksBatchidUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTasksBatchidForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTasksBatchidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTasksBatchidInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTasksBatchidOK creates a GetTasksBatchidOK with default headers values
func NewGetTasksBatchidOK() *GetTasksBatchidOK {
	return &GetTasksBatchidOK{}
}

/* GetTasksBatchidOK describes a response with status code 200, with default header values.

Successful
*/
type GetTasksBatchidOK struct {
	Payload *models.Task
}

func (o *GetTasksBatchidOK) Error() string {
	return fmt.Sprintf("[GET /tasks/{batchId}][%d] getTasksBatchidOK  %+v", 200, o.Payload)
}
func (o *GetTasksBatchidOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *GetTasksBatchidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksBatchidBadRequest creates a GetTasksBatchidBadRequest with default headers values
func NewGetTasksBatchidBadRequest() *GetTasksBatchidBadRequest {
	return &GetTasksBatchidBadRequest{}
}

/* GetTasksBatchidBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetTasksBatchidBadRequest struct {
	Payload *models.Error
}

func (o *GetTasksBatchidBadRequest) Error() string {
	return fmt.Sprintf("[GET /tasks/{batchId}][%d] getTasksBatchidBadRequest  %+v", 400, o.Payload)
}
func (o *GetTasksBatchidBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTasksBatchidBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksBatchidUnauthorized creates a GetTasksBatchidUnauthorized with default headers values
func NewGetTasksBatchidUnauthorized() *GetTasksBatchidUnauthorized {
	return &GetTasksBatchidUnauthorized{}
}

/* GetTasksBatchidUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetTasksBatchidUnauthorized struct {
	Payload *models.Error
}

func (o *GetTasksBatchidUnauthorized) Error() string {
	return fmt.Sprintf("[GET /tasks/{batchId}][%d] getTasksBatchidUnauthorized  %+v", 401, o.Payload)
}
func (o *GetTasksBatchidUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTasksBatchidUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksBatchidForbidden creates a GetTasksBatchidForbidden with default headers values
func NewGetTasksBatchidForbidden() *GetTasksBatchidForbidden {
	return &GetTasksBatchidForbidden{}
}

/* GetTasksBatchidForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetTasksBatchidForbidden struct {
	Payload *models.Error
}

func (o *GetTasksBatchidForbidden) Error() string {
	return fmt.Sprintf("[GET /tasks/{batchId}][%d] getTasksBatchidForbidden  %+v", 403, o.Payload)
}
func (o *GetTasksBatchidForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTasksBatchidForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksBatchidNotFound creates a GetTasksBatchidNotFound with default headers values
func NewGetTasksBatchidNotFound() *GetTasksBatchidNotFound {
	return &GetTasksBatchidNotFound{}
}

/* GetTasksBatchidNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetTasksBatchidNotFound struct {
	Payload *models.Error
}

func (o *GetTasksBatchidNotFound) Error() string {
	return fmt.Sprintf("[GET /tasks/{batchId}][%d] getTasksBatchidNotFound  %+v", 404, o.Payload)
}
func (o *GetTasksBatchidNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTasksBatchidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksBatchidInternalServerError creates a GetTasksBatchidInternalServerError with default headers values
func NewGetTasksBatchidInternalServerError() *GetTasksBatchidInternalServerError {
	return &GetTasksBatchidInternalServerError{}
}

/* GetTasksBatchidInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetTasksBatchidInternalServerError struct {
	Payload *models.Error
}

func (o *GetTasksBatchidInternalServerError) Error() string {
	return fmt.Sprintf("[GET /tasks/{batchId}][%d] getTasksBatchidInternalServerError  %+v", 500, o.Payload)
}
func (o *GetTasksBatchidInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTasksBatchidInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}