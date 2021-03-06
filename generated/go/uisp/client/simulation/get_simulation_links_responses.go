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

// GetSimulationLinksReader is a Reader for the GetSimulationLinks structure.
type GetSimulationLinksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSimulationLinksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSimulationLinksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSimulationLinksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSimulationLinksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSimulationLinksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSimulationLinksOK creates a GetSimulationLinksOK with default headers values
func NewGetSimulationLinksOK() *GetSimulationLinksOK {
	return &GetSimulationLinksOK{}
}

/* GetSimulationLinksOK describes a response with status code 200, with default header values.

Successful
*/
type GetSimulationLinksOK struct {
	Payload models.SimulationLinkList
}

func (o *GetSimulationLinksOK) Error() string {
	return fmt.Sprintf("[GET /simulation/links][%d] getSimulationLinksOK  %+v", 200, o.Payload)
}
func (o *GetSimulationLinksOK) GetPayload() models.SimulationLinkList {
	return o.Payload
}

func (o *GetSimulationLinksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSimulationLinksBadRequest creates a GetSimulationLinksBadRequest with default headers values
func NewGetSimulationLinksBadRequest() *GetSimulationLinksBadRequest {
	return &GetSimulationLinksBadRequest{}
}

/* GetSimulationLinksBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetSimulationLinksBadRequest struct {
	Payload *models.Error
}

func (o *GetSimulationLinksBadRequest) Error() string {
	return fmt.Sprintf("[GET /simulation/links][%d] getSimulationLinksBadRequest  %+v", 400, o.Payload)
}
func (o *GetSimulationLinksBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSimulationLinksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSimulationLinksUnauthorized creates a GetSimulationLinksUnauthorized with default headers values
func NewGetSimulationLinksUnauthorized() *GetSimulationLinksUnauthorized {
	return &GetSimulationLinksUnauthorized{}
}

/* GetSimulationLinksUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetSimulationLinksUnauthorized struct {
	Payload *models.Error
}

func (o *GetSimulationLinksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /simulation/links][%d] getSimulationLinksUnauthorized  %+v", 401, o.Payload)
}
func (o *GetSimulationLinksUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSimulationLinksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSimulationLinksInternalServerError creates a GetSimulationLinksInternalServerError with default headers values
func NewGetSimulationLinksInternalServerError() *GetSimulationLinksInternalServerError {
	return &GetSimulationLinksInternalServerError{}
}

/* GetSimulationLinksInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetSimulationLinksInternalServerError struct {
	Payload *models.Error
}

func (o *GetSimulationLinksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /simulation/links][%d] getSimulationLinksInternalServerError  %+v", 500, o.Payload)
}
func (o *GetSimulationLinksInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSimulationLinksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
