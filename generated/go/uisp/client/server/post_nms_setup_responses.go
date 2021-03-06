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

// PostNmsSetupReader is a Reader for the PostNmsSetup structure.
type PostNmsSetupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNmsSetupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostNmsSetupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPostNmsSetupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostNmsSetupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostNmsSetupOK creates a PostNmsSetupOK with default headers values
func NewPostNmsSetupOK() *PostNmsSetupOK {
	return &PostNmsSetupOK{}
}

/* PostNmsSetupOK describes a response with status code 200, with default header values.

Successful
*/
type PostNmsSetupOK struct {
	Payload *models.SetupUispFinish
}

func (o *PostNmsSetupOK) Error() string {
	return fmt.Sprintf("[POST /nms/setup][%d] postNmsSetupOK  %+v", 200, o.Payload)
}
func (o *PostNmsSetupOK) GetPayload() *models.SetupUispFinish {
	return o.Payload
}

func (o *PostNmsSetupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SetupUispFinish)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsSetupForbidden creates a PostNmsSetupForbidden with default headers values
func NewPostNmsSetupForbidden() *PostNmsSetupForbidden {
	return &PostNmsSetupForbidden{}
}

/* PostNmsSetupForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostNmsSetupForbidden struct {
	Payload *models.Error
}

func (o *PostNmsSetupForbidden) Error() string {
	return fmt.Sprintf("[POST /nms/setup][%d] postNmsSetupForbidden  %+v", 403, o.Payload)
}
func (o *PostNmsSetupForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsSetupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsSetupInternalServerError creates a PostNmsSetupInternalServerError with default headers values
func NewPostNmsSetupInternalServerError() *PostNmsSetupInternalServerError {
	return &PostNmsSetupInternalServerError{}
}

/* PostNmsSetupInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostNmsSetupInternalServerError struct {
	Payload *models.Error
}

func (o *PostNmsSetupInternalServerError) Error() string {
	return fmt.Sprintf("[POST /nms/setup][%d] postNmsSetupInternalServerError  %+v", 500, o.Payload)
}
func (o *PostNmsSetupInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsSetupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
