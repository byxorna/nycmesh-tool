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

// DeleteDevicesUnknownIpaddressReader is a Reader for the DeleteDevicesUnknownIpaddress structure.
type DeleteDevicesUnknownIpaddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDevicesUnknownIpaddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDevicesUnknownIpaddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteDevicesUnknownIpaddressBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteDevicesUnknownIpaddressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteDevicesUnknownIpaddressForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDevicesUnknownIpaddressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDevicesUnknownIpaddressOK creates a DeleteDevicesUnknownIpaddressOK with default headers values
func NewDeleteDevicesUnknownIpaddressOK() *DeleteDevicesUnknownIpaddressOK {
	return &DeleteDevicesUnknownIpaddressOK{}
}

/* DeleteDevicesUnknownIpaddressOK describes a response with status code 200, with default header values.

Successful
*/
type DeleteDevicesUnknownIpaddressOK struct {
	Payload *models.Status
}

func (o *DeleteDevicesUnknownIpaddressOK) Error() string {
	return fmt.Sprintf("[DELETE /devices/unknown/{ipAddress}][%d] deleteDevicesUnknownIpaddressOK  %+v", 200, o.Payload)
}
func (o *DeleteDevicesUnknownIpaddressOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *DeleteDevicesUnknownIpaddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesUnknownIpaddressBadRequest creates a DeleteDevicesUnknownIpaddressBadRequest with default headers values
func NewDeleteDevicesUnknownIpaddressBadRequest() *DeleteDevicesUnknownIpaddressBadRequest {
	return &DeleteDevicesUnknownIpaddressBadRequest{}
}

/* DeleteDevicesUnknownIpaddressBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteDevicesUnknownIpaddressBadRequest struct {
	Payload *models.Error
}

func (o *DeleteDevicesUnknownIpaddressBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /devices/unknown/{ipAddress}][%d] deleteDevicesUnknownIpaddressBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteDevicesUnknownIpaddressBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesUnknownIpaddressBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesUnknownIpaddressUnauthorized creates a DeleteDevicesUnknownIpaddressUnauthorized with default headers values
func NewDeleteDevicesUnknownIpaddressUnauthorized() *DeleteDevicesUnknownIpaddressUnauthorized {
	return &DeleteDevicesUnknownIpaddressUnauthorized{}
}

/* DeleteDevicesUnknownIpaddressUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteDevicesUnknownIpaddressUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteDevicesUnknownIpaddressUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /devices/unknown/{ipAddress}][%d] deleteDevicesUnknownIpaddressUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteDevicesUnknownIpaddressUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesUnknownIpaddressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesUnknownIpaddressForbidden creates a DeleteDevicesUnknownIpaddressForbidden with default headers values
func NewDeleteDevicesUnknownIpaddressForbidden() *DeleteDevicesUnknownIpaddressForbidden {
	return &DeleteDevicesUnknownIpaddressForbidden{}
}

/* DeleteDevicesUnknownIpaddressForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteDevicesUnknownIpaddressForbidden struct {
	Payload *models.Error
}

func (o *DeleteDevicesUnknownIpaddressForbidden) Error() string {
	return fmt.Sprintf("[DELETE /devices/unknown/{ipAddress}][%d] deleteDevicesUnknownIpaddressForbidden  %+v", 403, o.Payload)
}
func (o *DeleteDevicesUnknownIpaddressForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesUnknownIpaddressForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesUnknownIpaddressInternalServerError creates a DeleteDevicesUnknownIpaddressInternalServerError with default headers values
func NewDeleteDevicesUnknownIpaddressInternalServerError() *DeleteDevicesUnknownIpaddressInternalServerError {
	return &DeleteDevicesUnknownIpaddressInternalServerError{}
}

/* DeleteDevicesUnknownIpaddressInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteDevicesUnknownIpaddressInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteDevicesUnknownIpaddressInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /devices/unknown/{ipAddress}][%d] deleteDevicesUnknownIpaddressInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteDevicesUnknownIpaddressInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesUnknownIpaddressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}