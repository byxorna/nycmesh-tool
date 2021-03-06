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

// DeleteDevicesOltsIDOnuProfilesProfileidReader is a Reader for the DeleteDevicesOltsIDOnuProfilesProfileid structure.
type DeleteDevicesOltsIDOnuProfilesProfileidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDevicesOltsIDOnuProfilesProfileidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDevicesOltsIDOnuProfilesProfileidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteDevicesOltsIDOnuProfilesProfileidBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteDevicesOltsIDOnuProfilesProfileidUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteDevicesOltsIDOnuProfilesProfileidForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDevicesOltsIDOnuProfilesProfileidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDevicesOltsIDOnuProfilesProfileidInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDevicesOltsIDOnuProfilesProfileidOK creates a DeleteDevicesOltsIDOnuProfilesProfileidOK with default headers values
func NewDeleteDevicesOltsIDOnuProfilesProfileidOK() *DeleteDevicesOltsIDOnuProfilesProfileidOK {
	return &DeleteDevicesOltsIDOnuProfilesProfileidOK{}
}

/* DeleteDevicesOltsIDOnuProfilesProfileidOK describes a response with status code 200, with default header values.

Successful
*/
type DeleteDevicesOltsIDOnuProfilesProfileidOK struct {
	Payload *models.Status
}

func (o *DeleteDevicesOltsIDOnuProfilesProfileidOK) Error() string {
	return fmt.Sprintf("[DELETE /devices/olts/{id}/onu/profiles/{profileId}][%d] deleteDevicesOltsIdOnuProfilesProfileidOK  %+v", 200, o.Payload)
}
func (o *DeleteDevicesOltsIDOnuProfilesProfileidOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *DeleteDevicesOltsIDOnuProfilesProfileidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesOltsIDOnuProfilesProfileidBadRequest creates a DeleteDevicesOltsIDOnuProfilesProfileidBadRequest with default headers values
func NewDeleteDevicesOltsIDOnuProfilesProfileidBadRequest() *DeleteDevicesOltsIDOnuProfilesProfileidBadRequest {
	return &DeleteDevicesOltsIDOnuProfilesProfileidBadRequest{}
}

/* DeleteDevicesOltsIDOnuProfilesProfileidBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteDevicesOltsIDOnuProfilesProfileidBadRequest struct {
	Payload *models.Error
}

func (o *DeleteDevicesOltsIDOnuProfilesProfileidBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /devices/olts/{id}/onu/profiles/{profileId}][%d] deleteDevicesOltsIdOnuProfilesProfileidBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteDevicesOltsIDOnuProfilesProfileidBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesOltsIDOnuProfilesProfileidBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesOltsIDOnuProfilesProfileidUnauthorized creates a DeleteDevicesOltsIDOnuProfilesProfileidUnauthorized with default headers values
func NewDeleteDevicesOltsIDOnuProfilesProfileidUnauthorized() *DeleteDevicesOltsIDOnuProfilesProfileidUnauthorized {
	return &DeleteDevicesOltsIDOnuProfilesProfileidUnauthorized{}
}

/* DeleteDevicesOltsIDOnuProfilesProfileidUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteDevicesOltsIDOnuProfilesProfileidUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteDevicesOltsIDOnuProfilesProfileidUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /devices/olts/{id}/onu/profiles/{profileId}][%d] deleteDevicesOltsIdOnuProfilesProfileidUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteDevicesOltsIDOnuProfilesProfileidUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesOltsIDOnuProfilesProfileidUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesOltsIDOnuProfilesProfileidForbidden creates a DeleteDevicesOltsIDOnuProfilesProfileidForbidden with default headers values
func NewDeleteDevicesOltsIDOnuProfilesProfileidForbidden() *DeleteDevicesOltsIDOnuProfilesProfileidForbidden {
	return &DeleteDevicesOltsIDOnuProfilesProfileidForbidden{}
}

/* DeleteDevicesOltsIDOnuProfilesProfileidForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteDevicesOltsIDOnuProfilesProfileidForbidden struct {
	Payload *models.Error
}

func (o *DeleteDevicesOltsIDOnuProfilesProfileidForbidden) Error() string {
	return fmt.Sprintf("[DELETE /devices/olts/{id}/onu/profiles/{profileId}][%d] deleteDevicesOltsIdOnuProfilesProfileidForbidden  %+v", 403, o.Payload)
}
func (o *DeleteDevicesOltsIDOnuProfilesProfileidForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesOltsIDOnuProfilesProfileidForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesOltsIDOnuProfilesProfileidNotFound creates a DeleteDevicesOltsIDOnuProfilesProfileidNotFound with default headers values
func NewDeleteDevicesOltsIDOnuProfilesProfileidNotFound() *DeleteDevicesOltsIDOnuProfilesProfileidNotFound {
	return &DeleteDevicesOltsIDOnuProfilesProfileidNotFound{}
}

/* DeleteDevicesOltsIDOnuProfilesProfileidNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteDevicesOltsIDOnuProfilesProfileidNotFound struct {
	Payload *models.Error
}

func (o *DeleteDevicesOltsIDOnuProfilesProfileidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /devices/olts/{id}/onu/profiles/{profileId}][%d] deleteDevicesOltsIdOnuProfilesProfileidNotFound  %+v", 404, o.Payload)
}
func (o *DeleteDevicesOltsIDOnuProfilesProfileidNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesOltsIDOnuProfilesProfileidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesOltsIDOnuProfilesProfileidInternalServerError creates a DeleteDevicesOltsIDOnuProfilesProfileidInternalServerError with default headers values
func NewDeleteDevicesOltsIDOnuProfilesProfileidInternalServerError() *DeleteDevicesOltsIDOnuProfilesProfileidInternalServerError {
	return &DeleteDevicesOltsIDOnuProfilesProfileidInternalServerError{}
}

/* DeleteDevicesOltsIDOnuProfilesProfileidInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteDevicesOltsIDOnuProfilesProfileidInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteDevicesOltsIDOnuProfilesProfileidInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /devices/olts/{id}/onu/profiles/{profileId}][%d] deleteDevicesOltsIdOnuProfilesProfileidInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteDevicesOltsIDOnuProfilesProfileidInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesOltsIDOnuProfilesProfileidInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
