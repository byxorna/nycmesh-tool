// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// GetDevicesIDSystemUnmsReader is a Reader for the GetDevicesIDSystemUnms structure.
type GetDevicesIDSystemUnmsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesIDSystemUnmsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesIDSystemUnmsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesIDSystemUnmsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesIDSystemUnmsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesIDSystemUnmsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesIDSystemUnmsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesIDSystemUnmsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDevicesIDSystemUnmsOK creates a GetDevicesIDSystemUnmsOK with default headers values
func NewGetDevicesIDSystemUnmsOK() *GetDevicesIDSystemUnmsOK {
	return &GetDevicesIDSystemUnmsOK{}
}

/* GetDevicesIDSystemUnmsOK describes a response with status code 200, with default header values.

Successful
*/
type GetDevicesIDSystemUnmsOK struct {
	Payload *models.UispSetting
}

func (o *GetDevicesIDSystemUnmsOK) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/system/unms][%d] getDevicesIdSystemUnmsOK  %+v", 200, o.Payload)
}
func (o *GetDevicesIDSystemUnmsOK) GetPayload() *models.UispSetting {
	return o.Payload
}

func (o *GetDevicesIDSystemUnmsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UispSetting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesIDSystemUnmsBadRequest creates a GetDevicesIDSystemUnmsBadRequest with default headers values
func NewGetDevicesIDSystemUnmsBadRequest() *GetDevicesIDSystemUnmsBadRequest {
	return &GetDevicesIDSystemUnmsBadRequest{}
}

/* GetDevicesIDSystemUnmsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDevicesIDSystemUnmsBadRequest struct {
	Payload *models.Error
}

func (o *GetDevicesIDSystemUnmsBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/system/unms][%d] getDevicesIdSystemUnmsBadRequest  %+v", 400, o.Payload)
}
func (o *GetDevicesIDSystemUnmsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesIDSystemUnmsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesIDSystemUnmsUnauthorized creates a GetDevicesIDSystemUnmsUnauthorized with default headers values
func NewGetDevicesIDSystemUnmsUnauthorized() *GetDevicesIDSystemUnmsUnauthorized {
	return &GetDevicesIDSystemUnmsUnauthorized{}
}

/* GetDevicesIDSystemUnmsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDevicesIDSystemUnmsUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesIDSystemUnmsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/system/unms][%d] getDevicesIdSystemUnmsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDevicesIDSystemUnmsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesIDSystemUnmsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesIDSystemUnmsForbidden creates a GetDevicesIDSystemUnmsForbidden with default headers values
func NewGetDevicesIDSystemUnmsForbidden() *GetDevicesIDSystemUnmsForbidden {
	return &GetDevicesIDSystemUnmsForbidden{}
}

/* GetDevicesIDSystemUnmsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDevicesIDSystemUnmsForbidden struct {
	Payload *models.Error
}

func (o *GetDevicesIDSystemUnmsForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/system/unms][%d] getDevicesIdSystemUnmsForbidden  %+v", 403, o.Payload)
}
func (o *GetDevicesIDSystemUnmsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesIDSystemUnmsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesIDSystemUnmsNotFound creates a GetDevicesIDSystemUnmsNotFound with default headers values
func NewGetDevicesIDSystemUnmsNotFound() *GetDevicesIDSystemUnmsNotFound {
	return &GetDevicesIDSystemUnmsNotFound{}
}

/* GetDevicesIDSystemUnmsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDevicesIDSystemUnmsNotFound struct {
	Payload *models.Error
}

func (o *GetDevicesIDSystemUnmsNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/system/unms][%d] getDevicesIdSystemUnmsNotFound  %+v", 404, o.Payload)
}
func (o *GetDevicesIDSystemUnmsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesIDSystemUnmsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesIDSystemUnmsInternalServerError creates a GetDevicesIDSystemUnmsInternalServerError with default headers values
func NewGetDevicesIDSystemUnmsInternalServerError() *GetDevicesIDSystemUnmsInternalServerError {
	return &GetDevicesIDSystemUnmsInternalServerError{}
}

/* GetDevicesIDSystemUnmsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDevicesIDSystemUnmsInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesIDSystemUnmsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/system/unms][%d] getDevicesIdSystemUnmsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDevicesIDSystemUnmsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesIDSystemUnmsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}