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

// GetTasksInprogressReader is a Reader for the GetTasksInprogress structure.
type GetTasksInprogressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTasksInprogressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTasksInprogressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTasksInprogressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTasksInprogressForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTasksInprogressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTasksInprogressOK creates a GetTasksInprogressOK with default headers values
func NewGetTasksInprogressOK() *GetTasksInprogressOK {
	return &GetTasksInprogressOK{}
}

/* GetTasksInprogressOK describes a response with status code 200, with default header values.

Successful
*/
type GetTasksInprogressOK struct {
	Payload *models.InProgress
}

func (o *GetTasksInprogressOK) Error() string {
	return fmt.Sprintf("[GET /tasks/in-progress][%d] getTasksInprogressOK  %+v", 200, o.Payload)
}
func (o *GetTasksInprogressOK) GetPayload() *models.InProgress {
	return o.Payload
}

func (o *GetTasksInprogressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InProgress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksInprogressUnauthorized creates a GetTasksInprogressUnauthorized with default headers values
func NewGetTasksInprogressUnauthorized() *GetTasksInprogressUnauthorized {
	return &GetTasksInprogressUnauthorized{}
}

/* GetTasksInprogressUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetTasksInprogressUnauthorized struct {
	Payload *models.Error
}

func (o *GetTasksInprogressUnauthorized) Error() string {
	return fmt.Sprintf("[GET /tasks/in-progress][%d] getTasksInprogressUnauthorized  %+v", 401, o.Payload)
}
func (o *GetTasksInprogressUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTasksInprogressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksInprogressForbidden creates a GetTasksInprogressForbidden with default headers values
func NewGetTasksInprogressForbidden() *GetTasksInprogressForbidden {
	return &GetTasksInprogressForbidden{}
}

/* GetTasksInprogressForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetTasksInprogressForbidden struct {
	Payload *models.Error
}

func (o *GetTasksInprogressForbidden) Error() string {
	return fmt.Sprintf("[GET /tasks/in-progress][%d] getTasksInprogressForbidden  %+v", 403, o.Payload)
}
func (o *GetTasksInprogressForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTasksInprogressForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksInprogressInternalServerError creates a GetTasksInprogressInternalServerError with default headers values
func NewGetTasksInprogressInternalServerError() *GetTasksInprogressInternalServerError {
	return &GetTasksInprogressInternalServerError{}
}

/* GetTasksInprogressInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetTasksInprogressInternalServerError struct {
	Payload *models.Error
}

func (o *GetTasksInprogressInternalServerError) Error() string {
	return fmt.Sprintf("[GET /tasks/in-progress][%d] getTasksInprogressInternalServerError  %+v", 500, o.Payload)
}
func (o *GetTasksInprogressInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTasksInprogressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
