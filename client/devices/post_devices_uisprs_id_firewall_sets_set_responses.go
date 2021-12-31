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

// PostDevicesUisprsIDFirewallSetsSetReader is a Reader for the PostDevicesUisprsIDFirewallSetsSet structure.
type PostDevicesUisprsIDFirewallSetsSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDevicesUisprsIDFirewallSetsSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDevicesUisprsIDFirewallSetsSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDevicesUisprsIDFirewallSetsSetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDevicesUisprsIDFirewallSetsSetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostDevicesUisprsIDFirewallSetsSetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDevicesUisprsIDFirewallSetsSetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDevicesUisprsIDFirewallSetsSetOK creates a PostDevicesUisprsIDFirewallSetsSetOK with default headers values
func NewPostDevicesUisprsIDFirewallSetsSetOK() *PostDevicesUisprsIDFirewallSetsSetOK {
	return &PostDevicesUisprsIDFirewallSetsSetOK{}
}

/* PostDevicesUisprsIDFirewallSetsSetOK describes a response with status code 200, with default header values.

Successful
*/
type PostDevicesUisprsIDFirewallSetsSetOK struct {
	Payload *models.Status
}

func (o *PostDevicesUisprsIDFirewallSetsSetOK) Error() string {
	return fmt.Sprintf("[POST /devices/uisprs/{id}/firewall/sets/set][%d] postDevicesUisprsIdFirewallSetsSetOK  %+v", 200, o.Payload)
}
func (o *PostDevicesUisprsIDFirewallSetsSetOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PostDevicesUisprsIDFirewallSetsSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesUisprsIDFirewallSetsSetBadRequest creates a PostDevicesUisprsIDFirewallSetsSetBadRequest with default headers values
func NewPostDevicesUisprsIDFirewallSetsSetBadRequest() *PostDevicesUisprsIDFirewallSetsSetBadRequest {
	return &PostDevicesUisprsIDFirewallSetsSetBadRequest{}
}

/* PostDevicesUisprsIDFirewallSetsSetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostDevicesUisprsIDFirewallSetsSetBadRequest struct {
	Payload *models.Error
}

func (o *PostDevicesUisprsIDFirewallSetsSetBadRequest) Error() string {
	return fmt.Sprintf("[POST /devices/uisprs/{id}/firewall/sets/set][%d] postDevicesUisprsIdFirewallSetsSetBadRequest  %+v", 400, o.Payload)
}
func (o *PostDevicesUisprsIDFirewallSetsSetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesUisprsIDFirewallSetsSetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesUisprsIDFirewallSetsSetUnauthorized creates a PostDevicesUisprsIDFirewallSetsSetUnauthorized with default headers values
func NewPostDevicesUisprsIDFirewallSetsSetUnauthorized() *PostDevicesUisprsIDFirewallSetsSetUnauthorized {
	return &PostDevicesUisprsIDFirewallSetsSetUnauthorized{}
}

/* PostDevicesUisprsIDFirewallSetsSetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostDevicesUisprsIDFirewallSetsSetUnauthorized struct {
	Payload *models.Error
}

func (o *PostDevicesUisprsIDFirewallSetsSetUnauthorized) Error() string {
	return fmt.Sprintf("[POST /devices/uisprs/{id}/firewall/sets/set][%d] postDevicesUisprsIdFirewallSetsSetUnauthorized  %+v", 401, o.Payload)
}
func (o *PostDevicesUisprsIDFirewallSetsSetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesUisprsIDFirewallSetsSetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesUisprsIDFirewallSetsSetNotFound creates a PostDevicesUisprsIDFirewallSetsSetNotFound with default headers values
func NewPostDevicesUisprsIDFirewallSetsSetNotFound() *PostDevicesUisprsIDFirewallSetsSetNotFound {
	return &PostDevicesUisprsIDFirewallSetsSetNotFound{}
}

/* PostDevicesUisprsIDFirewallSetsSetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostDevicesUisprsIDFirewallSetsSetNotFound struct {
	Payload *models.Error
}

func (o *PostDevicesUisprsIDFirewallSetsSetNotFound) Error() string {
	return fmt.Sprintf("[POST /devices/uisprs/{id}/firewall/sets/set][%d] postDevicesUisprsIdFirewallSetsSetNotFound  %+v", 404, o.Payload)
}
func (o *PostDevicesUisprsIDFirewallSetsSetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesUisprsIDFirewallSetsSetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesUisprsIDFirewallSetsSetInternalServerError creates a PostDevicesUisprsIDFirewallSetsSetInternalServerError with default headers values
func NewPostDevicesUisprsIDFirewallSetsSetInternalServerError() *PostDevicesUisprsIDFirewallSetsSetInternalServerError {
	return &PostDevicesUisprsIDFirewallSetsSetInternalServerError{}
}

/* PostDevicesUisprsIDFirewallSetsSetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostDevicesUisprsIDFirewallSetsSetInternalServerError struct {
	Payload *models.Error
}

func (o *PostDevicesUisprsIDFirewallSetsSetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /devices/uisprs/{id}/firewall/sets/set][%d] postDevicesUisprsIdFirewallSetsSetInternalServerError  %+v", 500, o.Payload)
}
func (o *PostDevicesUisprsIDFirewallSetsSetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesUisprsIDFirewallSetsSetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
