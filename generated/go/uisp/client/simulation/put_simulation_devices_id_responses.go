// Code generated by go-swagger; DO NOT EDIT.

package simulation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// PutSimulationDevicesIDReader is a Reader for the PutSimulationDevicesID structure.
type PutSimulationDevicesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutSimulationDevicesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutSimulationDevicesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutSimulationDevicesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutSimulationDevicesIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutSimulationDevicesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutSimulationDevicesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutSimulationDevicesIDOK creates a PutSimulationDevicesIDOK with default headers values
func NewPutSimulationDevicesIDOK() *PutSimulationDevicesIDOK {
	return &PutSimulationDevicesIDOK{}
}

/* PutSimulationDevicesIDOK describes a response with status code 200, with default header values.

Successful
*/
type PutSimulationDevicesIDOK struct {
	Payload *models.SimulationDevice
}

func (o *PutSimulationDevicesIDOK) Error() string {
	return fmt.Sprintf("[PUT /simulation/devices/{id}][%d] putSimulationDevicesIdOK  %+v", 200, o.Payload)
}
func (o *PutSimulationDevicesIDOK) GetPayload() *models.SimulationDevice {
	return o.Payload
}

func (o *PutSimulationDevicesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimulationDevice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSimulationDevicesIDBadRequest creates a PutSimulationDevicesIDBadRequest with default headers values
func NewPutSimulationDevicesIDBadRequest() *PutSimulationDevicesIDBadRequest {
	return &PutSimulationDevicesIDBadRequest{}
}

/* PutSimulationDevicesIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutSimulationDevicesIDBadRequest struct {
	Payload *models.Error
}

func (o *PutSimulationDevicesIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /simulation/devices/{id}][%d] putSimulationDevicesIdBadRequest  %+v", 400, o.Payload)
}
func (o *PutSimulationDevicesIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutSimulationDevicesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSimulationDevicesIDUnauthorized creates a PutSimulationDevicesIDUnauthorized with default headers values
func NewPutSimulationDevicesIDUnauthorized() *PutSimulationDevicesIDUnauthorized {
	return &PutSimulationDevicesIDUnauthorized{}
}

/* PutSimulationDevicesIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutSimulationDevicesIDUnauthorized struct {
	Payload *models.Error
}

func (o *PutSimulationDevicesIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /simulation/devices/{id}][%d] putSimulationDevicesIdUnauthorized  %+v", 401, o.Payload)
}
func (o *PutSimulationDevicesIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutSimulationDevicesIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSimulationDevicesIDNotFound creates a PutSimulationDevicesIDNotFound with default headers values
func NewPutSimulationDevicesIDNotFound() *PutSimulationDevicesIDNotFound {
	return &PutSimulationDevicesIDNotFound{}
}

/* PutSimulationDevicesIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutSimulationDevicesIDNotFound struct {
	Payload *models.Error
}

func (o *PutSimulationDevicesIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /simulation/devices/{id}][%d] putSimulationDevicesIdNotFound  %+v", 404, o.Payload)
}
func (o *PutSimulationDevicesIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutSimulationDevicesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSimulationDevicesIDInternalServerError creates a PutSimulationDevicesIDInternalServerError with default headers values
func NewPutSimulationDevicesIDInternalServerError() *PutSimulationDevicesIDInternalServerError {
	return &PutSimulationDevicesIDInternalServerError{}
}

/* PutSimulationDevicesIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutSimulationDevicesIDInternalServerError struct {
	Payload *models.Error
}

func (o *PutSimulationDevicesIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /simulation/devices/{id}][%d] putSimulationDevicesIdInternalServerError  %+v", 500, o.Payload)
}
func (o *PutSimulationDevicesIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutSimulationDevicesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
