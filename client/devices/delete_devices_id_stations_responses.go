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

// DeleteDevicesIDStationsReader is a Reader for the DeleteDevicesIDStations structure.
type DeleteDevicesIDStationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDevicesIDStationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDevicesIDStationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteDevicesIDStationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteDevicesIDStationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteDevicesIDStationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDevicesIDStationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDevicesIDStationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDevicesIDStationsOK creates a DeleteDevicesIDStationsOK with default headers values
func NewDeleteDevicesIDStationsOK() *DeleteDevicesIDStationsOK {
	return &DeleteDevicesIDStationsOK{}
}

/* DeleteDevicesIDStationsOK describes a response with status code 200, with default header values.

Successful
*/
type DeleteDevicesIDStationsOK struct {
	Payload models.Model129
}

func (o *DeleteDevicesIDStationsOK) Error() string {
	return fmt.Sprintf("[DELETE /devices/{id}/stations][%d] deleteDevicesIdStationsOK  %+v", 200, o.Payload)
}
func (o *DeleteDevicesIDStationsOK) GetPayload() models.Model129 {
	return o.Payload
}

func (o *DeleteDevicesIDStationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesIDStationsBadRequest creates a DeleteDevicesIDStationsBadRequest with default headers values
func NewDeleteDevicesIDStationsBadRequest() *DeleteDevicesIDStationsBadRequest {
	return &DeleteDevicesIDStationsBadRequest{}
}

/* DeleteDevicesIDStationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteDevicesIDStationsBadRequest struct {
	Payload *models.Error
}

func (o *DeleteDevicesIDStationsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /devices/{id}/stations][%d] deleteDevicesIdStationsBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteDevicesIDStationsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesIDStationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesIDStationsUnauthorized creates a DeleteDevicesIDStationsUnauthorized with default headers values
func NewDeleteDevicesIDStationsUnauthorized() *DeleteDevicesIDStationsUnauthorized {
	return &DeleteDevicesIDStationsUnauthorized{}
}

/* DeleteDevicesIDStationsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteDevicesIDStationsUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteDevicesIDStationsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /devices/{id}/stations][%d] deleteDevicesIdStationsUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteDevicesIDStationsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesIDStationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesIDStationsForbidden creates a DeleteDevicesIDStationsForbidden with default headers values
func NewDeleteDevicesIDStationsForbidden() *DeleteDevicesIDStationsForbidden {
	return &DeleteDevicesIDStationsForbidden{}
}

/* DeleteDevicesIDStationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteDevicesIDStationsForbidden struct {
	Payload *models.Error
}

func (o *DeleteDevicesIDStationsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /devices/{id}/stations][%d] deleteDevicesIdStationsForbidden  %+v", 403, o.Payload)
}
func (o *DeleteDevicesIDStationsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesIDStationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesIDStationsNotFound creates a DeleteDevicesIDStationsNotFound with default headers values
func NewDeleteDevicesIDStationsNotFound() *DeleteDevicesIDStationsNotFound {
	return &DeleteDevicesIDStationsNotFound{}
}

/* DeleteDevicesIDStationsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteDevicesIDStationsNotFound struct {
	Payload *models.Error
}

func (o *DeleteDevicesIDStationsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /devices/{id}/stations][%d] deleteDevicesIdStationsNotFound  %+v", 404, o.Payload)
}
func (o *DeleteDevicesIDStationsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesIDStationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesIDStationsInternalServerError creates a DeleteDevicesIDStationsInternalServerError with default headers values
func NewDeleteDevicesIDStationsInternalServerError() *DeleteDevicesIDStationsInternalServerError {
	return &DeleteDevicesIDStationsInternalServerError{}
}

/* DeleteDevicesIDStationsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteDevicesIDStationsInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteDevicesIDStationsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /devices/{id}/stations][%d] deleteDevicesIdStationsInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteDevicesIDStationsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesIDStationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
