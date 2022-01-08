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

// GetNmsVersionReader is a Reader for the GetNmsVersion structure.
type GetNmsVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNmsVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNmsVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetNmsVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNmsVersionOK creates a GetNmsVersionOK with default headers values
func NewGetNmsVersionOK() *GetNmsVersionOK {
	return &GetNmsVersionOK{}
}

/* GetNmsVersionOK describes a response with status code 200, with default header values.

Successful
*/
type GetNmsVersionOK struct {
	Payload *models.Version
}

func (o *GetNmsVersionOK) Error() string {
	return fmt.Sprintf("[GET /nms/version][%d] getNmsVersionOK  %+v", 200, o.Payload)
}
func (o *GetNmsVersionOK) GetPayload() *models.Version {
	return o.Payload
}

func (o *GetNmsVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Version)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNmsVersionInternalServerError creates a GetNmsVersionInternalServerError with default headers values
func NewGetNmsVersionInternalServerError() *GetNmsVersionInternalServerError {
	return &GetNmsVersionInternalServerError{}
}

/* GetNmsVersionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetNmsVersionInternalServerError struct {
	Payload *models.Error
}

func (o *GetNmsVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /nms/version][%d] getNmsVersionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetNmsVersionInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNmsVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}