// Code generated by go-swagger; DO NOT EDIT.

package token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// GetTokenReader is a Reader for the GetToken structure.
type GetTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetTokenConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTokenOK creates a GetTokenOK with default headers values
func NewGetTokenOK() *GetTokenOK {
	return &GetTokenOK{}
}

/* GetTokenOK describes a response with status code 200, with default header values.

List of tokens.
*/
type GetTokenOK struct {
	Payload models.TokenListSchema
}

func (o *GetTokenOK) Error() string {
	return fmt.Sprintf("[GET /token][%d] getTokenOK  %+v", 200, o.Payload)
}
func (o *GetTokenOK) GetPayload() models.TokenListSchema {
	return o.Payload
}

func (o *GetTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokenUnauthorized creates a GetTokenUnauthorized with default headers values
func NewGetTokenUnauthorized() *GetTokenUnauthorized {
	return &GetTokenUnauthorized{}
}

/* GetTokenUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetTokenUnauthorized struct {
	Payload *models.Error
}

func (o *GetTokenUnauthorized) Error() string {
	return fmt.Sprintf("[GET /token][%d] getTokenUnauthorized  %+v", 401, o.Payload)
}
func (o *GetTokenUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokenForbidden creates a GetTokenForbidden with default headers values
func NewGetTokenForbidden() *GetTokenForbidden {
	return &GetTokenForbidden{}
}

/* GetTokenForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetTokenForbidden struct {
	Payload *models.Error
}

func (o *GetTokenForbidden) Error() string {
	return fmt.Sprintf("[GET /token][%d] getTokenForbidden  %+v", 403, o.Payload)
}
func (o *GetTokenForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokenConflict creates a GetTokenConflict with default headers values
func NewGetTokenConflict() *GetTokenConflict {
	return &GetTokenConflict{}
}

/* GetTokenConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetTokenConflict struct {
	Payload *models.Error
}

func (o *GetTokenConflict) Error() string {
	return fmt.Sprintf("[GET /token][%d] getTokenConflict  %+v", 409, o.Payload)
}
func (o *GetTokenConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTokenConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokenInternalServerError creates a GetTokenInternalServerError with default headers values
func NewGetTokenInternalServerError() *GetTokenInternalServerError {
	return &GetTokenInternalServerError{}
}

/* GetTokenInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetTokenInternalServerError struct {
	Payload *models.Error
}

func (o *GetTokenInternalServerError) Error() string {
	return fmt.Sprintf("[GET /token][%d] getTokenInternalServerError  %+v", 500, o.Payload)
}
func (o *GetTokenInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}