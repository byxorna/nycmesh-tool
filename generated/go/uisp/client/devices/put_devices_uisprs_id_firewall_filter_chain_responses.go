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

// PutDevicesUisprsIDFirewallFilterChainReader is a Reader for the PutDevicesUisprsIDFirewallFilterChain structure.
type PutDevicesUisprsIDFirewallFilterChainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDevicesUisprsIDFirewallFilterChainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutDevicesUisprsIDFirewallFilterChainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutDevicesUisprsIDFirewallFilterChainBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutDevicesUisprsIDFirewallFilterChainUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutDevicesUisprsIDFirewallFilterChainNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutDevicesUisprsIDFirewallFilterChainInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutDevicesUisprsIDFirewallFilterChainOK creates a PutDevicesUisprsIDFirewallFilterChainOK with default headers values
func NewPutDevicesUisprsIDFirewallFilterChainOK() *PutDevicesUisprsIDFirewallFilterChainOK {
	return &PutDevicesUisprsIDFirewallFilterChainOK{}
}

/* PutDevicesUisprsIDFirewallFilterChainOK describes a response with status code 200, with default header values.

Successful
*/
type PutDevicesUisprsIDFirewallFilterChainOK struct {
	Payload *models.Model60
}

func (o *PutDevicesUisprsIDFirewallFilterChainOK) Error() string {
	return fmt.Sprintf("[PUT /devices/uisprs/{id}/firewall/filter/chain][%d] putDevicesUisprsIdFirewallFilterChainOK  %+v", 200, o.Payload)
}
func (o *PutDevicesUisprsIDFirewallFilterChainOK) GetPayload() *models.Model60 {
	return o.Payload
}

func (o *PutDevicesUisprsIDFirewallFilterChainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Model60)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesUisprsIDFirewallFilterChainBadRequest creates a PutDevicesUisprsIDFirewallFilterChainBadRequest with default headers values
func NewPutDevicesUisprsIDFirewallFilterChainBadRequest() *PutDevicesUisprsIDFirewallFilterChainBadRequest {
	return &PutDevicesUisprsIDFirewallFilterChainBadRequest{}
}

/* PutDevicesUisprsIDFirewallFilterChainBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutDevicesUisprsIDFirewallFilterChainBadRequest struct {
	Payload *models.Error
}

func (o *PutDevicesUisprsIDFirewallFilterChainBadRequest) Error() string {
	return fmt.Sprintf("[PUT /devices/uisprs/{id}/firewall/filter/chain][%d] putDevicesUisprsIdFirewallFilterChainBadRequest  %+v", 400, o.Payload)
}
func (o *PutDevicesUisprsIDFirewallFilterChainBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesUisprsIDFirewallFilterChainBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesUisprsIDFirewallFilterChainUnauthorized creates a PutDevicesUisprsIDFirewallFilterChainUnauthorized with default headers values
func NewPutDevicesUisprsIDFirewallFilterChainUnauthorized() *PutDevicesUisprsIDFirewallFilterChainUnauthorized {
	return &PutDevicesUisprsIDFirewallFilterChainUnauthorized{}
}

/* PutDevicesUisprsIDFirewallFilterChainUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutDevicesUisprsIDFirewallFilterChainUnauthorized struct {
	Payload *models.Error
}

func (o *PutDevicesUisprsIDFirewallFilterChainUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /devices/uisprs/{id}/firewall/filter/chain][%d] putDevicesUisprsIdFirewallFilterChainUnauthorized  %+v", 401, o.Payload)
}
func (o *PutDevicesUisprsIDFirewallFilterChainUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesUisprsIDFirewallFilterChainUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesUisprsIDFirewallFilterChainNotFound creates a PutDevicesUisprsIDFirewallFilterChainNotFound with default headers values
func NewPutDevicesUisprsIDFirewallFilterChainNotFound() *PutDevicesUisprsIDFirewallFilterChainNotFound {
	return &PutDevicesUisprsIDFirewallFilterChainNotFound{}
}

/* PutDevicesUisprsIDFirewallFilterChainNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutDevicesUisprsIDFirewallFilterChainNotFound struct {
	Payload *models.Error
}

func (o *PutDevicesUisprsIDFirewallFilterChainNotFound) Error() string {
	return fmt.Sprintf("[PUT /devices/uisprs/{id}/firewall/filter/chain][%d] putDevicesUisprsIdFirewallFilterChainNotFound  %+v", 404, o.Payload)
}
func (o *PutDevicesUisprsIDFirewallFilterChainNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesUisprsIDFirewallFilterChainNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesUisprsIDFirewallFilterChainInternalServerError creates a PutDevicesUisprsIDFirewallFilterChainInternalServerError with default headers values
func NewPutDevicesUisprsIDFirewallFilterChainInternalServerError() *PutDevicesUisprsIDFirewallFilterChainInternalServerError {
	return &PutDevicesUisprsIDFirewallFilterChainInternalServerError{}
}

/* PutDevicesUisprsIDFirewallFilterChainInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutDevicesUisprsIDFirewallFilterChainInternalServerError struct {
	Payload *models.Error
}

func (o *PutDevicesUisprsIDFirewallFilterChainInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /devices/uisprs/{id}/firewall/filter/chain][%d] putDevicesUisprsIdFirewallFilterChainInternalServerError  %+v", 500, o.Payload)
}
func (o *PutDevicesUisprsIDFirewallFilterChainInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesUisprsIDFirewallFilterChainInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}