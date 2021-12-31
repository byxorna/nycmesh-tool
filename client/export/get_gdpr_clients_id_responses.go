// Code generated by go-swagger; DO NOT EDIT.

package export

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/models"
)

// GetGdprClientsIDReader is a Reader for the GetGdprClientsID structure.
type GetGdprClientsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGdprClientsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewGetGdprClientsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGdprClientsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGdprClientsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGdprClientsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGdprClientsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGdprClientsIDBadRequest creates a GetGdprClientsIDBadRequest with default headers values
func NewGetGdprClientsIDBadRequest() *GetGdprClientsIDBadRequest {
	return &GetGdprClientsIDBadRequest{}
}

/* GetGdprClientsIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetGdprClientsIDBadRequest struct {
	Payload *models.Error
}

func (o *GetGdprClientsIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /gdpr/clients/{id}][%d] getGdprClientsIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetGdprClientsIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGdprClientsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprClientsIDUnauthorized creates a GetGdprClientsIDUnauthorized with default headers values
func NewGetGdprClientsIDUnauthorized() *GetGdprClientsIDUnauthorized {
	return &GetGdprClientsIDUnauthorized{}
}

/* GetGdprClientsIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetGdprClientsIDUnauthorized struct {
	Payload *models.Error
}

func (o *GetGdprClientsIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /gdpr/clients/{id}][%d] getGdprClientsIdUnauthorized  %+v", 401, o.Payload)
}
func (o *GetGdprClientsIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGdprClientsIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprClientsIDForbidden creates a GetGdprClientsIDForbidden with default headers values
func NewGetGdprClientsIDForbidden() *GetGdprClientsIDForbidden {
	return &GetGdprClientsIDForbidden{}
}

/* GetGdprClientsIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetGdprClientsIDForbidden struct {
	Payload *models.Error
}

func (o *GetGdprClientsIDForbidden) Error() string {
	return fmt.Sprintf("[GET /gdpr/clients/{id}][%d] getGdprClientsIdForbidden  %+v", 403, o.Payload)
}
func (o *GetGdprClientsIDForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGdprClientsIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprClientsIDNotFound creates a GetGdprClientsIDNotFound with default headers values
func NewGetGdprClientsIDNotFound() *GetGdprClientsIDNotFound {
	return &GetGdprClientsIDNotFound{}
}

/* GetGdprClientsIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetGdprClientsIDNotFound struct {
	Payload *models.Error
}

func (o *GetGdprClientsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /gdpr/clients/{id}][%d] getGdprClientsIdNotFound  %+v", 404, o.Payload)
}
func (o *GetGdprClientsIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGdprClientsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprClientsIDInternalServerError creates a GetGdprClientsIDInternalServerError with default headers values
func NewGetGdprClientsIDInternalServerError() *GetGdprClientsIDInternalServerError {
	return &GetGdprClientsIDInternalServerError{}
}

/* GetGdprClientsIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetGdprClientsIDInternalServerError struct {
	Payload *models.Error
}

func (o *GetGdprClientsIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /gdpr/clients/{id}][%d] getGdprClientsIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetGdprClientsIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGdprClientsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
