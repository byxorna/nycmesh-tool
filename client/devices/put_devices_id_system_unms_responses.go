// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/models"
)

// PutDevicesIDSystemUnmsReader is a Reader for the PutDevicesIDSystemUnms structure.
type PutDevicesIDSystemUnmsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDevicesIDSystemUnmsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutDevicesIDSystemUnmsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutDevicesIDSystemUnmsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutDevicesIDSystemUnmsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutDevicesIDSystemUnmsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutDevicesIDSystemUnmsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutDevicesIDSystemUnmsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutDevicesIDSystemUnmsOK creates a PutDevicesIDSystemUnmsOK with default headers values
func NewPutDevicesIDSystemUnmsOK() *PutDevicesIDSystemUnmsOK {
	return &PutDevicesIDSystemUnmsOK{}
}

/* PutDevicesIDSystemUnmsOK describes a response with status code 200, with default header values.

Successful
*/
type PutDevicesIDSystemUnmsOK struct {
	Payload *models.UispSetting
}

func (o *PutDevicesIDSystemUnmsOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/system/unms][%d] putDevicesIdSystemUnmsOK  %+v", 200, o.Payload)
}
func (o *PutDevicesIDSystemUnmsOK) GetPayload() *models.UispSetting {
	return o.Payload
}

func (o *PutDevicesIDSystemUnmsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UispSetting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDSystemUnmsBadRequest creates a PutDevicesIDSystemUnmsBadRequest with default headers values
func NewPutDevicesIDSystemUnmsBadRequest() *PutDevicesIDSystemUnmsBadRequest {
	return &PutDevicesIDSystemUnmsBadRequest{}
}

/* PutDevicesIDSystemUnmsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutDevicesIDSystemUnmsBadRequest struct {
	Payload *models.Error
}

func (o *PutDevicesIDSystemUnmsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/system/unms][%d] putDevicesIdSystemUnmsBadRequest  %+v", 400, o.Payload)
}
func (o *PutDevicesIDSystemUnmsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDSystemUnmsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDSystemUnmsUnauthorized creates a PutDevicesIDSystemUnmsUnauthorized with default headers values
func NewPutDevicesIDSystemUnmsUnauthorized() *PutDevicesIDSystemUnmsUnauthorized {
	return &PutDevicesIDSystemUnmsUnauthorized{}
}

/* PutDevicesIDSystemUnmsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutDevicesIDSystemUnmsUnauthorized struct {
	Payload *models.Error
}

func (o *PutDevicesIDSystemUnmsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/system/unms][%d] putDevicesIdSystemUnmsUnauthorized  %+v", 401, o.Payload)
}
func (o *PutDevicesIDSystemUnmsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDSystemUnmsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDSystemUnmsForbidden creates a PutDevicesIDSystemUnmsForbidden with default headers values
func NewPutDevicesIDSystemUnmsForbidden() *PutDevicesIDSystemUnmsForbidden {
	return &PutDevicesIDSystemUnmsForbidden{}
}

/* PutDevicesIDSystemUnmsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutDevicesIDSystemUnmsForbidden struct {
	Payload *models.Error
}

func (o *PutDevicesIDSystemUnmsForbidden) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/system/unms][%d] putDevicesIdSystemUnmsForbidden  %+v", 403, o.Payload)
}
func (o *PutDevicesIDSystemUnmsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDSystemUnmsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDSystemUnmsNotFound creates a PutDevicesIDSystemUnmsNotFound with default headers values
func NewPutDevicesIDSystemUnmsNotFound() *PutDevicesIDSystemUnmsNotFound {
	return &PutDevicesIDSystemUnmsNotFound{}
}

/* PutDevicesIDSystemUnmsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutDevicesIDSystemUnmsNotFound struct {
	Payload *models.Error
}

func (o *PutDevicesIDSystemUnmsNotFound) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/system/unms][%d] putDevicesIdSystemUnmsNotFound  %+v", 404, o.Payload)
}
func (o *PutDevicesIDSystemUnmsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDSystemUnmsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesIDSystemUnmsInternalServerError creates a PutDevicesIDSystemUnmsInternalServerError with default headers values
func NewPutDevicesIDSystemUnmsInternalServerError() *PutDevicesIDSystemUnmsInternalServerError {
	return &PutDevicesIDSystemUnmsInternalServerError{}
}

/* PutDevicesIDSystemUnmsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutDevicesIDSystemUnmsInternalServerError struct {
	Payload *models.Error
}

func (o *PutDevicesIDSystemUnmsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}/system/unms][%d] putDevicesIdSystemUnmsInternalServerError  %+v", 500, o.Payload)
}
func (o *PutDevicesIDSystemUnmsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesIDSystemUnmsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}