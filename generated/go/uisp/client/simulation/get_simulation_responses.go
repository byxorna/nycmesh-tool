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

// GetSimulationReader is a Reader for the GetSimulation structure.
type GetSimulationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSimulationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSimulationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSimulationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSimulationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSimulationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSimulationOK creates a GetSimulationOK with default headers values
func NewGetSimulationOK() *GetSimulationOK {
	return &GetSimulationOK{}
}

/* GetSimulationOK describes a response with status code 200, with default header values.

Successful
*/
type GetSimulationOK struct {
	Payload *models.Simulation
}

func (o *GetSimulationOK) Error() string {
	return fmt.Sprintf("[GET /simulation][%d] getSimulationOK  %+v", 200, o.Payload)
}
func (o *GetSimulationOK) GetPayload() *models.Simulation {
	return o.Payload
}

func (o *GetSimulationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Simulation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSimulationBadRequest creates a GetSimulationBadRequest with default headers values
func NewGetSimulationBadRequest() *GetSimulationBadRequest {
	return &GetSimulationBadRequest{}
}

/* GetSimulationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetSimulationBadRequest struct {
	Payload *models.Error
}

func (o *GetSimulationBadRequest) Error() string {
	return fmt.Sprintf("[GET /simulation][%d] getSimulationBadRequest  %+v", 400, o.Payload)
}
func (o *GetSimulationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSimulationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSimulationUnauthorized creates a GetSimulationUnauthorized with default headers values
func NewGetSimulationUnauthorized() *GetSimulationUnauthorized {
	return &GetSimulationUnauthorized{}
}

/* GetSimulationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetSimulationUnauthorized struct {
	Payload *models.Error
}

func (o *GetSimulationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /simulation][%d] getSimulationUnauthorized  %+v", 401, o.Payload)
}
func (o *GetSimulationUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSimulationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSimulationInternalServerError creates a GetSimulationInternalServerError with default headers values
func NewGetSimulationInternalServerError() *GetSimulationInternalServerError {
	return &GetSimulationInternalServerError{}
}

/* GetSimulationInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetSimulationInternalServerError struct {
	Payload *models.Error
}

func (o *GetSimulationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /simulation][%d] getSimulationInternalServerError  %+v", 500, o.Payload)
}
func (o *GetSimulationInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSimulationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
