// Code generated by go-swagger; DO NOT EDIT.

package data_links

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// GetDatalinksDeviceIDReader is a Reader for the GetDatalinksDeviceID structure.
type GetDatalinksDeviceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatalinksDeviceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDatalinksDeviceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDatalinksDeviceIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDatalinksDeviceIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDatalinksDeviceIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDatalinksDeviceIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDatalinksDeviceIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDatalinksDeviceIDOK creates a GetDatalinksDeviceIDOK with default headers values
func NewGetDatalinksDeviceIDOK() *GetDatalinksDeviceIDOK {
	return &GetDatalinksDeviceIDOK{}
}

/* GetDatalinksDeviceIDOK describes a response with status code 200, with default header values.

Successful
*/
type GetDatalinksDeviceIDOK struct {
	Payload *models.DataLink
}

func (o *GetDatalinksDeviceIDOK) Error() string {
	return fmt.Sprintf("[GET /data-links/device/{id}][%d] getDatalinksDeviceIdOK  %+v", 200, o.Payload)
}
func (o *GetDatalinksDeviceIDOK) GetPayload() *models.DataLink {
	return o.Payload
}

func (o *GetDatalinksDeviceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatalinksDeviceIDBadRequest creates a GetDatalinksDeviceIDBadRequest with default headers values
func NewGetDatalinksDeviceIDBadRequest() *GetDatalinksDeviceIDBadRequest {
	return &GetDatalinksDeviceIDBadRequest{}
}

/* GetDatalinksDeviceIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDatalinksDeviceIDBadRequest struct {
	Payload *models.Error
}

func (o *GetDatalinksDeviceIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /data-links/device/{id}][%d] getDatalinksDeviceIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetDatalinksDeviceIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDatalinksDeviceIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatalinksDeviceIDUnauthorized creates a GetDatalinksDeviceIDUnauthorized with default headers values
func NewGetDatalinksDeviceIDUnauthorized() *GetDatalinksDeviceIDUnauthorized {
	return &GetDatalinksDeviceIDUnauthorized{}
}

/* GetDatalinksDeviceIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDatalinksDeviceIDUnauthorized struct {
	Payload *models.Error
}

func (o *GetDatalinksDeviceIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /data-links/device/{id}][%d] getDatalinksDeviceIdUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDatalinksDeviceIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDatalinksDeviceIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatalinksDeviceIDForbidden creates a GetDatalinksDeviceIDForbidden with default headers values
func NewGetDatalinksDeviceIDForbidden() *GetDatalinksDeviceIDForbidden {
	return &GetDatalinksDeviceIDForbidden{}
}

/* GetDatalinksDeviceIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDatalinksDeviceIDForbidden struct {
	Payload *models.Error
}

func (o *GetDatalinksDeviceIDForbidden) Error() string {
	return fmt.Sprintf("[GET /data-links/device/{id}][%d] getDatalinksDeviceIdForbidden  %+v", 403, o.Payload)
}
func (o *GetDatalinksDeviceIDForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDatalinksDeviceIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatalinksDeviceIDNotFound creates a GetDatalinksDeviceIDNotFound with default headers values
func NewGetDatalinksDeviceIDNotFound() *GetDatalinksDeviceIDNotFound {
	return &GetDatalinksDeviceIDNotFound{}
}

/* GetDatalinksDeviceIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDatalinksDeviceIDNotFound struct {
	Payload *models.Error
}

func (o *GetDatalinksDeviceIDNotFound) Error() string {
	return fmt.Sprintf("[GET /data-links/device/{id}][%d] getDatalinksDeviceIdNotFound  %+v", 404, o.Payload)
}
func (o *GetDatalinksDeviceIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDatalinksDeviceIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatalinksDeviceIDInternalServerError creates a GetDatalinksDeviceIDInternalServerError with default headers values
func NewGetDatalinksDeviceIDInternalServerError() *GetDatalinksDeviceIDInternalServerError {
	return &GetDatalinksDeviceIDInternalServerError{}
}

/* GetDatalinksDeviceIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDatalinksDeviceIDInternalServerError struct {
	Payload *models.Error
}

func (o *GetDatalinksDeviceIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /data-links/device/{id}][%d] getDatalinksDeviceIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDatalinksDeviceIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDatalinksDeviceIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}