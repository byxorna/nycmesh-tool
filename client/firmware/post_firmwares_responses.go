// Code generated by go-swagger; DO NOT EDIT.

package firmware

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/models"
)

// PostFirmwaresReader is a Reader for the PostFirmwares structure.
type PostFirmwaresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFirmwaresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostFirmwaresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostFirmwaresUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostFirmwaresForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostFirmwaresInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostFirmwaresOK creates a PostFirmwaresOK with default headers values
func NewPostFirmwaresOK() *PostFirmwaresOK {
	return &PostFirmwaresOK{}
}

/* PostFirmwaresOK describes a response with status code 200, with default header values.

Successful
*/
type PostFirmwaresOK struct {
	Payload *models.Firmware
}

func (o *PostFirmwaresOK) Error() string {
	return fmt.Sprintf("[POST /firmwares][%d] postFirmwaresOK  %+v", 200, o.Payload)
}
func (o *PostFirmwaresOK) GetPayload() *models.Firmware {
	return o.Payload
}

func (o *PostFirmwaresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Firmware)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFirmwaresUnauthorized creates a PostFirmwaresUnauthorized with default headers values
func NewPostFirmwaresUnauthorized() *PostFirmwaresUnauthorized {
	return &PostFirmwaresUnauthorized{}
}

/* PostFirmwaresUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostFirmwaresUnauthorized struct {
	Payload *models.Error
}

func (o *PostFirmwaresUnauthorized) Error() string {
	return fmt.Sprintf("[POST /firmwares][%d] postFirmwaresUnauthorized  %+v", 401, o.Payload)
}
func (o *PostFirmwaresUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostFirmwaresUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFirmwaresForbidden creates a PostFirmwaresForbidden with default headers values
func NewPostFirmwaresForbidden() *PostFirmwaresForbidden {
	return &PostFirmwaresForbidden{}
}

/* PostFirmwaresForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostFirmwaresForbidden struct {
	Payload *models.Error
}

func (o *PostFirmwaresForbidden) Error() string {
	return fmt.Sprintf("[POST /firmwares][%d] postFirmwaresForbidden  %+v", 403, o.Payload)
}
func (o *PostFirmwaresForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostFirmwaresForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFirmwaresInternalServerError creates a PostFirmwaresInternalServerError with default headers values
func NewPostFirmwaresInternalServerError() *PostFirmwaresInternalServerError {
	return &PostFirmwaresInternalServerError{}
}

/* PostFirmwaresInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostFirmwaresInternalServerError struct {
	Payload *models.Error
}

func (o *PostFirmwaresInternalServerError) Error() string {
	return fmt.Sprintf("[POST /firmwares][%d] postFirmwaresInternalServerError  %+v", 500, o.Payload)
}
func (o *PostFirmwaresInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostFirmwaresInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
