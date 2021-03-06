// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// GetLogsReader is a Reader for the GetLogs structure.
type GetLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLogsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLogsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLogsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLogsOK creates a GetLogsOK with default headers values
func NewGetLogsOK() *GetLogsOK {
	return &GetLogsOK{}
}

/* GetLogsOK describes a response with status code 200, with default header values.

Successful
*/
type GetLogsOK struct {
	Payload *models.Model10
}

func (o *GetLogsOK) Error() string {
	return fmt.Sprintf("[GET /logs][%d] getLogsOK  %+v", 200, o.Payload)
}
func (o *GetLogsOK) GetPayload() *models.Model10 {
	return o.Payload
}

func (o *GetLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Model10)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsBadRequest creates a GetLogsBadRequest with default headers values
func NewGetLogsBadRequest() *GetLogsBadRequest {
	return &GetLogsBadRequest{}
}

/* GetLogsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetLogsBadRequest struct {
	Payload *models.Error
}

func (o *GetLogsBadRequest) Error() string {
	return fmt.Sprintf("[GET /logs][%d] getLogsBadRequest  %+v", 400, o.Payload)
}
func (o *GetLogsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLogsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsUnauthorized creates a GetLogsUnauthorized with default headers values
func NewGetLogsUnauthorized() *GetLogsUnauthorized {
	return &GetLogsUnauthorized{}
}

/* GetLogsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetLogsUnauthorized struct {
	Payload *models.Error
}

func (o *GetLogsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /logs][%d] getLogsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetLogsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsForbidden creates a GetLogsForbidden with default headers values
func NewGetLogsForbidden() *GetLogsForbidden {
	return &GetLogsForbidden{}
}

/* GetLogsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetLogsForbidden struct {
	Payload *models.Error
}

func (o *GetLogsForbidden) Error() string {
	return fmt.Sprintf("[GET /logs][%d] getLogsForbidden  %+v", 403, o.Payload)
}
func (o *GetLogsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLogsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsInternalServerError creates a GetLogsInternalServerError with default headers values
func NewGetLogsInternalServerError() *GetLogsInternalServerError {
	return &GetLogsInternalServerError{}
}

/* GetLogsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetLogsInternalServerError struct {
	Payload *models.Error
}

func (o *GetLogsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /logs][%d] getLogsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetLogsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLogsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
