// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/models"
)

// GetNmsKeepaliveReader is a Reader for the GetNmsKeepalive structure.
type GetNmsKeepaliveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNmsKeepaliveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNmsKeepaliveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetNmsKeepaliveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNmsKeepaliveOK creates a GetNmsKeepaliveOK with default headers values
func NewGetNmsKeepaliveOK() *GetNmsKeepaliveOK {
	return &GetNmsKeepaliveOK{}
}

/* GetNmsKeepaliveOK describes a response with status code 200, with default header values.

Successful
*/
type GetNmsKeepaliveOK struct {
	Payload *models.Model17
}

func (o *GetNmsKeepaliveOK) Error() string {
	return fmt.Sprintf("[GET /nms/keep-alive][%d] getNmsKeepaliveOK  %+v", 200, o.Payload)
}
func (o *GetNmsKeepaliveOK) GetPayload() *models.Model17 {
	return o.Payload
}

func (o *GetNmsKeepaliveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Model17)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNmsKeepaliveInternalServerError creates a GetNmsKeepaliveInternalServerError with default headers values
func NewGetNmsKeepaliveInternalServerError() *GetNmsKeepaliveInternalServerError {
	return &GetNmsKeepaliveInternalServerError{}
}

/* GetNmsKeepaliveInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetNmsKeepaliveInternalServerError struct {
	Payload *models.Error
}

func (o *GetNmsKeepaliveInternalServerError) Error() string {
	return fmt.Sprintf("[GET /nms/keep-alive][%d] getNmsKeepaliveInternalServerError  %+v", 500, o.Payload)
}
func (o *GetNmsKeepaliveInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNmsKeepaliveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
