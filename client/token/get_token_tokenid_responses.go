// Code generated by go-swagger; DO NOT EDIT.

package token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/models"
)

// GetTokenTokenidReader is a Reader for the GetTokenTokenid structure.
type GetTokenTokenidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTokenTokenidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTokenTokenidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTokenTokenidUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTokenTokenidForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTokenTokenidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetTokenTokenidConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTokenTokenidInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTokenTokenidOK creates a GetTokenTokenidOK with default headers values
func NewGetTokenTokenidOK() *GetTokenTokenidOK {
	return &GetTokenTokenidOK{}
}

/* GetTokenTokenidOK describes a response with status code 200, with default header values.

Information about token.
*/
type GetTokenTokenidOK struct {
	Payload *models.TokenSchema
}

func (o *GetTokenTokenidOK) Error() string {
	return fmt.Sprintf("[GET /token/{tokenId}][%d] getTokenTokenidOK  %+v", 200, o.Payload)
}
func (o *GetTokenTokenidOK) GetPayload() *models.TokenSchema {
	return o.Payload
}

func (o *GetTokenTokenidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TokenSchema)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokenTokenidUnauthorized creates a GetTokenTokenidUnauthorized with default headers values
func NewGetTokenTokenidUnauthorized() *GetTokenTokenidUnauthorized {
	return &GetTokenTokenidUnauthorized{}
}

/* GetTokenTokenidUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetTokenTokenidUnauthorized struct {
	Payload *models.Error
}

func (o *GetTokenTokenidUnauthorized) Error() string {
	return fmt.Sprintf("[GET /token/{tokenId}][%d] getTokenTokenidUnauthorized  %+v", 401, o.Payload)
}
func (o *GetTokenTokenidUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTokenTokenidUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokenTokenidForbidden creates a GetTokenTokenidForbidden with default headers values
func NewGetTokenTokenidForbidden() *GetTokenTokenidForbidden {
	return &GetTokenTokenidForbidden{}
}

/* GetTokenTokenidForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetTokenTokenidForbidden struct {
	Payload *models.Error
}

func (o *GetTokenTokenidForbidden) Error() string {
	return fmt.Sprintf("[GET /token/{tokenId}][%d] getTokenTokenidForbidden  %+v", 403, o.Payload)
}
func (o *GetTokenTokenidForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTokenTokenidForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokenTokenidNotFound creates a GetTokenTokenidNotFound with default headers values
func NewGetTokenTokenidNotFound() *GetTokenTokenidNotFound {
	return &GetTokenTokenidNotFound{}
}

/* GetTokenTokenidNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetTokenTokenidNotFound struct {
	Payload *models.Error
}

func (o *GetTokenTokenidNotFound) Error() string {
	return fmt.Sprintf("[GET /token/{tokenId}][%d] getTokenTokenidNotFound  %+v", 404, o.Payload)
}
func (o *GetTokenTokenidNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTokenTokenidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokenTokenidConflict creates a GetTokenTokenidConflict with default headers values
func NewGetTokenTokenidConflict() *GetTokenTokenidConflict {
	return &GetTokenTokenidConflict{}
}

/* GetTokenTokenidConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetTokenTokenidConflict struct {
	Payload *models.Error
}

func (o *GetTokenTokenidConflict) Error() string {
	return fmt.Sprintf("[GET /token/{tokenId}][%d] getTokenTokenidConflict  %+v", 409, o.Payload)
}
func (o *GetTokenTokenidConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTokenTokenidConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokenTokenidInternalServerError creates a GetTokenTokenidInternalServerError with default headers values
func NewGetTokenTokenidInternalServerError() *GetTokenTokenidInternalServerError {
	return &GetTokenTokenidInternalServerError{}
}

/* GetTokenTokenidInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetTokenTokenidInternalServerError struct {
	Payload *models.Error
}

func (o *GetTokenTokenidInternalServerError) Error() string {
	return fmt.Sprintf("[GET /token/{tokenId}][%d] getTokenTokenidInternalServerError  %+v", 500, o.Payload)
}
func (o *GetTokenTokenidInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTokenTokenidInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
