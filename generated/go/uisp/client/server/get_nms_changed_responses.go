// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// GetNmsChangedReader is a Reader for the GetNmsChanged structure.
type GetNmsChangedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNmsChangedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNmsChangedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetNmsChangedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNmsChangedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNmsChangedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNmsChangedOK creates a GetNmsChangedOK with default headers values
func NewGetNmsChangedOK() *GetNmsChangedOK {
	return &GetNmsChangedOK{}
}

/* GetNmsChangedOK describes a response with status code 200, with default header values.

Successful
*/
type GetNmsChangedOK struct {
	Payload *models.ChangedItems
}

func (o *GetNmsChangedOK) Error() string {
	return fmt.Sprintf("[GET /nms/changed][%d] getNmsChangedOK  %+v", 200, o.Payload)
}
func (o *GetNmsChangedOK) GetPayload() *models.ChangedItems {
	return o.Payload
}

func (o *GetNmsChangedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChangedItems)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNmsChangedUnauthorized creates a GetNmsChangedUnauthorized with default headers values
func NewGetNmsChangedUnauthorized() *GetNmsChangedUnauthorized {
	return &GetNmsChangedUnauthorized{}
}

/* GetNmsChangedUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetNmsChangedUnauthorized struct {
	Payload *models.Error
}

func (o *GetNmsChangedUnauthorized) Error() string {
	return fmt.Sprintf("[GET /nms/changed][%d] getNmsChangedUnauthorized  %+v", 401, o.Payload)
}
func (o *GetNmsChangedUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNmsChangedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNmsChangedForbidden creates a GetNmsChangedForbidden with default headers values
func NewGetNmsChangedForbidden() *GetNmsChangedForbidden {
	return &GetNmsChangedForbidden{}
}

/* GetNmsChangedForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetNmsChangedForbidden struct {
	Payload *models.Error
}

func (o *GetNmsChangedForbidden) Error() string {
	return fmt.Sprintf("[GET /nms/changed][%d] getNmsChangedForbidden  %+v", 403, o.Payload)
}
func (o *GetNmsChangedForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNmsChangedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNmsChangedInternalServerError creates a GetNmsChangedInternalServerError with default headers values
func NewGetNmsChangedInternalServerError() *GetNmsChangedInternalServerError {
	return &GetNmsChangedInternalServerError{}
}

/* GetNmsChangedInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetNmsChangedInternalServerError struct {
	Payload *models.Error
}

func (o *GetNmsChangedInternalServerError) Error() string {
	return fmt.Sprintf("[GET /nms/changed][%d] getNmsChangedInternalServerError  %+v", 500, o.Payload)
}
func (o *GetNmsChangedInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNmsChangedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
