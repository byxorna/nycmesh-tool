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

// PostDevicesIDVlansReader is a Reader for the PostDevicesIDVlans structure.
type PostDevicesIDVlansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDevicesIDVlansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDevicesIDVlansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDevicesIDVlansBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDevicesIDVlansUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostDevicesIDVlansForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDevicesIDVlansInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDevicesIDVlansOK creates a PostDevicesIDVlansOK with default headers values
func NewPostDevicesIDVlansOK() *PostDevicesIDVlansOK {
	return &PostDevicesIDVlansOK{}
}

/* PostDevicesIDVlansOK describes a response with status code 200, with default header values.

Successful
*/
type PostDevicesIDVlansOK struct {
	Payload *models.Model30
}

func (o *PostDevicesIDVlansOK) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/vlans][%d] postDevicesIdVlansOK  %+v", 200, o.Payload)
}
func (o *PostDevicesIDVlansOK) GetPayload() *models.Model30 {
	return o.Payload
}

func (o *PostDevicesIDVlansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Model30)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDVlansBadRequest creates a PostDevicesIDVlansBadRequest with default headers values
func NewPostDevicesIDVlansBadRequest() *PostDevicesIDVlansBadRequest {
	return &PostDevicesIDVlansBadRequest{}
}

/* PostDevicesIDVlansBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostDevicesIDVlansBadRequest struct {
	Payload *models.Error
}

func (o *PostDevicesIDVlansBadRequest) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/vlans][%d] postDevicesIdVlansBadRequest  %+v", 400, o.Payload)
}
func (o *PostDevicesIDVlansBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDVlansBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDVlansUnauthorized creates a PostDevicesIDVlansUnauthorized with default headers values
func NewPostDevicesIDVlansUnauthorized() *PostDevicesIDVlansUnauthorized {
	return &PostDevicesIDVlansUnauthorized{}
}

/* PostDevicesIDVlansUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostDevicesIDVlansUnauthorized struct {
	Payload *models.Error
}

func (o *PostDevicesIDVlansUnauthorized) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/vlans][%d] postDevicesIdVlansUnauthorized  %+v", 401, o.Payload)
}
func (o *PostDevicesIDVlansUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDVlansUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDVlansForbidden creates a PostDevicesIDVlansForbidden with default headers values
func NewPostDevicesIDVlansForbidden() *PostDevicesIDVlansForbidden {
	return &PostDevicesIDVlansForbidden{}
}

/* PostDevicesIDVlansForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostDevicesIDVlansForbidden struct {
	Payload *models.Error
}

func (o *PostDevicesIDVlansForbidden) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/vlans][%d] postDevicesIdVlansForbidden  %+v", 403, o.Payload)
}
func (o *PostDevicesIDVlansForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDVlansForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDVlansInternalServerError creates a PostDevicesIDVlansInternalServerError with default headers values
func NewPostDevicesIDVlansInternalServerError() *PostDevicesIDVlansInternalServerError {
	return &PostDevicesIDVlansInternalServerError{}
}

/* PostDevicesIDVlansInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostDevicesIDVlansInternalServerError struct {
	Payload *models.Error
}

func (o *PostDevicesIDVlansInternalServerError) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/vlans][%d] postDevicesIdVlansInternalServerError  %+v", 500, o.Payload)
}
func (o *PostDevicesIDVlansInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDVlansInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
