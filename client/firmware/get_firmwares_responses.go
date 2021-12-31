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

// GetFirmwaresReader is a Reader for the GetFirmwares structure.
type GetFirmwaresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFirmwaresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFirmwaresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetFirmwaresUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFirmwaresForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFirmwaresInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFirmwaresOK creates a GetFirmwaresOK with default headers values
func NewGetFirmwaresOK() *GetFirmwaresOK {
	return &GetFirmwaresOK{}
}

/* GetFirmwaresOK describes a response with status code 200, with default header values.

Successful
*/
type GetFirmwaresOK struct {
	Payload models.ListOfFirmwares
}

func (o *GetFirmwaresOK) Error() string {
	return fmt.Sprintf("[GET /firmwares][%d] getFirmwaresOK  %+v", 200, o.Payload)
}
func (o *GetFirmwaresOK) GetPayload() models.ListOfFirmwares {
	return o.Payload
}

func (o *GetFirmwaresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFirmwaresUnauthorized creates a GetFirmwaresUnauthorized with default headers values
func NewGetFirmwaresUnauthorized() *GetFirmwaresUnauthorized {
	return &GetFirmwaresUnauthorized{}
}

/* GetFirmwaresUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetFirmwaresUnauthorized struct {
	Payload *models.Error
}

func (o *GetFirmwaresUnauthorized) Error() string {
	return fmt.Sprintf("[GET /firmwares][%d] getFirmwaresUnauthorized  %+v", 401, o.Payload)
}
func (o *GetFirmwaresUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFirmwaresUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFirmwaresForbidden creates a GetFirmwaresForbidden with default headers values
func NewGetFirmwaresForbidden() *GetFirmwaresForbidden {
	return &GetFirmwaresForbidden{}
}

/* GetFirmwaresForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetFirmwaresForbidden struct {
	Payload *models.Error
}

func (o *GetFirmwaresForbidden) Error() string {
	return fmt.Sprintf("[GET /firmwares][%d] getFirmwaresForbidden  %+v", 403, o.Payload)
}
func (o *GetFirmwaresForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFirmwaresForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFirmwaresInternalServerError creates a GetFirmwaresInternalServerError with default headers values
func NewGetFirmwaresInternalServerError() *GetFirmwaresInternalServerError {
	return &GetFirmwaresInternalServerError{}
}

/* GetFirmwaresInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetFirmwaresInternalServerError struct {
	Payload *models.Error
}

func (o *GetFirmwaresInternalServerError) Error() string {
	return fmt.Sprintf("[GET /firmwares][%d] getFirmwaresInternalServerError  %+v", 500, o.Payload)
}
func (o *GetFirmwaresInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFirmwaresInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
