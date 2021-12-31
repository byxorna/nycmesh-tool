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

// PutDevicesOltsIDOnuProfilesProfileidReader is a Reader for the PutDevicesOltsIDOnuProfilesProfileid structure.
type PutDevicesOltsIDOnuProfilesProfileidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDevicesOltsIDOnuProfilesProfileidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutDevicesOltsIDOnuProfilesProfileidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutDevicesOltsIDOnuProfilesProfileidBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutDevicesOltsIDOnuProfilesProfileidUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutDevicesOltsIDOnuProfilesProfileidForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutDevicesOltsIDOnuProfilesProfileidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutDevicesOltsIDOnuProfilesProfileidInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutDevicesOltsIDOnuProfilesProfileidOK creates a PutDevicesOltsIDOnuProfilesProfileidOK with default headers values
func NewPutDevicesOltsIDOnuProfilesProfileidOK() *PutDevicesOltsIDOnuProfilesProfileidOK {
	return &PutDevicesOltsIDOnuProfilesProfileidOK{}
}

/* PutDevicesOltsIDOnuProfilesProfileidOK describes a response with status code 200, with default header values.

Successful
*/
type PutDevicesOltsIDOnuProfilesProfileidOK struct {
	Payload *models.OnuProfile
}

func (o *PutDevicesOltsIDOnuProfilesProfileidOK) Error() string {
	return fmt.Sprintf("[PUT /devices/olts/{id}/onu/profiles/{profileId}][%d] putDevicesOltsIdOnuProfilesProfileidOK  %+v", 200, o.Payload)
}
func (o *PutDevicesOltsIDOnuProfilesProfileidOK) GetPayload() *models.OnuProfile {
	return o.Payload
}

func (o *PutDevicesOltsIDOnuProfilesProfileidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OnuProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesOltsIDOnuProfilesProfileidBadRequest creates a PutDevicesOltsIDOnuProfilesProfileidBadRequest with default headers values
func NewPutDevicesOltsIDOnuProfilesProfileidBadRequest() *PutDevicesOltsIDOnuProfilesProfileidBadRequest {
	return &PutDevicesOltsIDOnuProfilesProfileidBadRequest{}
}

/* PutDevicesOltsIDOnuProfilesProfileidBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutDevicesOltsIDOnuProfilesProfileidBadRequest struct {
	Payload *models.Error
}

func (o *PutDevicesOltsIDOnuProfilesProfileidBadRequest) Error() string {
	return fmt.Sprintf("[PUT /devices/olts/{id}/onu/profiles/{profileId}][%d] putDevicesOltsIdOnuProfilesProfileidBadRequest  %+v", 400, o.Payload)
}
func (o *PutDevicesOltsIDOnuProfilesProfileidBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesOltsIDOnuProfilesProfileidBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesOltsIDOnuProfilesProfileidUnauthorized creates a PutDevicesOltsIDOnuProfilesProfileidUnauthorized with default headers values
func NewPutDevicesOltsIDOnuProfilesProfileidUnauthorized() *PutDevicesOltsIDOnuProfilesProfileidUnauthorized {
	return &PutDevicesOltsIDOnuProfilesProfileidUnauthorized{}
}

/* PutDevicesOltsIDOnuProfilesProfileidUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutDevicesOltsIDOnuProfilesProfileidUnauthorized struct {
	Payload *models.Error
}

func (o *PutDevicesOltsIDOnuProfilesProfileidUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /devices/olts/{id}/onu/profiles/{profileId}][%d] putDevicesOltsIdOnuProfilesProfileidUnauthorized  %+v", 401, o.Payload)
}
func (o *PutDevicesOltsIDOnuProfilesProfileidUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesOltsIDOnuProfilesProfileidUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesOltsIDOnuProfilesProfileidForbidden creates a PutDevicesOltsIDOnuProfilesProfileidForbidden with default headers values
func NewPutDevicesOltsIDOnuProfilesProfileidForbidden() *PutDevicesOltsIDOnuProfilesProfileidForbidden {
	return &PutDevicesOltsIDOnuProfilesProfileidForbidden{}
}

/* PutDevicesOltsIDOnuProfilesProfileidForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutDevicesOltsIDOnuProfilesProfileidForbidden struct {
	Payload *models.Error
}

func (o *PutDevicesOltsIDOnuProfilesProfileidForbidden) Error() string {
	return fmt.Sprintf("[PUT /devices/olts/{id}/onu/profiles/{profileId}][%d] putDevicesOltsIdOnuProfilesProfileidForbidden  %+v", 403, o.Payload)
}
func (o *PutDevicesOltsIDOnuProfilesProfileidForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesOltsIDOnuProfilesProfileidForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesOltsIDOnuProfilesProfileidNotFound creates a PutDevicesOltsIDOnuProfilesProfileidNotFound with default headers values
func NewPutDevicesOltsIDOnuProfilesProfileidNotFound() *PutDevicesOltsIDOnuProfilesProfileidNotFound {
	return &PutDevicesOltsIDOnuProfilesProfileidNotFound{}
}

/* PutDevicesOltsIDOnuProfilesProfileidNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutDevicesOltsIDOnuProfilesProfileidNotFound struct {
	Payload *models.Error
}

func (o *PutDevicesOltsIDOnuProfilesProfileidNotFound) Error() string {
	return fmt.Sprintf("[PUT /devices/olts/{id}/onu/profiles/{profileId}][%d] putDevicesOltsIdOnuProfilesProfileidNotFound  %+v", 404, o.Payload)
}
func (o *PutDevicesOltsIDOnuProfilesProfileidNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesOltsIDOnuProfilesProfileidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesOltsIDOnuProfilesProfileidInternalServerError creates a PutDevicesOltsIDOnuProfilesProfileidInternalServerError with default headers values
func NewPutDevicesOltsIDOnuProfilesProfileidInternalServerError() *PutDevicesOltsIDOnuProfilesProfileidInternalServerError {
	return &PutDevicesOltsIDOnuProfilesProfileidInternalServerError{}
}

/* PutDevicesOltsIDOnuProfilesProfileidInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutDevicesOltsIDOnuProfilesProfileidInternalServerError struct {
	Payload *models.Error
}

func (o *PutDevicesOltsIDOnuProfilesProfileidInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /devices/olts/{id}/onu/profiles/{profileId}][%d] putDevicesOltsIdOnuProfilesProfileidInternalServerError  %+v", 500, o.Payload)
}
func (o *PutDevicesOltsIDOnuProfilesProfileidInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesOltsIDOnuProfilesProfileidInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}