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

// PutDevicesUisprsIDFirewallNatsRulesReader is a Reader for the PutDevicesUisprsIDFirewallNatsRules structure.
type PutDevicesUisprsIDFirewallNatsRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDevicesUisprsIDFirewallNatsRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutDevicesUisprsIDFirewallNatsRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutDevicesUisprsIDFirewallNatsRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutDevicesUisprsIDFirewallNatsRulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutDevicesUisprsIDFirewallNatsRulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutDevicesUisprsIDFirewallNatsRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutDevicesUisprsIDFirewallNatsRulesOK creates a PutDevicesUisprsIDFirewallNatsRulesOK with default headers values
func NewPutDevicesUisprsIDFirewallNatsRulesOK() *PutDevicesUisprsIDFirewallNatsRulesOK {
	return &PutDevicesUisprsIDFirewallNatsRulesOK{}
}

/* PutDevicesUisprsIDFirewallNatsRulesOK describes a response with status code 200, with default header values.

Successful
*/
type PutDevicesUisprsIDFirewallNatsRulesOK struct {
	Payload *models.Status
}

func (o *PutDevicesUisprsIDFirewallNatsRulesOK) Error() string {
	return fmt.Sprintf("[PUT /devices/uisprs/{id}/firewall/nats/rules][%d] putDevicesUisprsIdFirewallNatsRulesOK  %+v", 200, o.Payload)
}
func (o *PutDevicesUisprsIDFirewallNatsRulesOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PutDevicesUisprsIDFirewallNatsRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesUisprsIDFirewallNatsRulesBadRequest creates a PutDevicesUisprsIDFirewallNatsRulesBadRequest with default headers values
func NewPutDevicesUisprsIDFirewallNatsRulesBadRequest() *PutDevicesUisprsIDFirewallNatsRulesBadRequest {
	return &PutDevicesUisprsIDFirewallNatsRulesBadRequest{}
}

/* PutDevicesUisprsIDFirewallNatsRulesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutDevicesUisprsIDFirewallNatsRulesBadRequest struct {
	Payload *models.Error
}

func (o *PutDevicesUisprsIDFirewallNatsRulesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /devices/uisprs/{id}/firewall/nats/rules][%d] putDevicesUisprsIdFirewallNatsRulesBadRequest  %+v", 400, o.Payload)
}
func (o *PutDevicesUisprsIDFirewallNatsRulesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesUisprsIDFirewallNatsRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesUisprsIDFirewallNatsRulesUnauthorized creates a PutDevicesUisprsIDFirewallNatsRulesUnauthorized with default headers values
func NewPutDevicesUisprsIDFirewallNatsRulesUnauthorized() *PutDevicesUisprsIDFirewallNatsRulesUnauthorized {
	return &PutDevicesUisprsIDFirewallNatsRulesUnauthorized{}
}

/* PutDevicesUisprsIDFirewallNatsRulesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutDevicesUisprsIDFirewallNatsRulesUnauthorized struct {
	Payload *models.Error
}

func (o *PutDevicesUisprsIDFirewallNatsRulesUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /devices/uisprs/{id}/firewall/nats/rules][%d] putDevicesUisprsIdFirewallNatsRulesUnauthorized  %+v", 401, o.Payload)
}
func (o *PutDevicesUisprsIDFirewallNatsRulesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesUisprsIDFirewallNatsRulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesUisprsIDFirewallNatsRulesNotFound creates a PutDevicesUisprsIDFirewallNatsRulesNotFound with default headers values
func NewPutDevicesUisprsIDFirewallNatsRulesNotFound() *PutDevicesUisprsIDFirewallNatsRulesNotFound {
	return &PutDevicesUisprsIDFirewallNatsRulesNotFound{}
}

/* PutDevicesUisprsIDFirewallNatsRulesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutDevicesUisprsIDFirewallNatsRulesNotFound struct {
	Payload *models.Error
}

func (o *PutDevicesUisprsIDFirewallNatsRulesNotFound) Error() string {
	return fmt.Sprintf("[PUT /devices/uisprs/{id}/firewall/nats/rules][%d] putDevicesUisprsIdFirewallNatsRulesNotFound  %+v", 404, o.Payload)
}
func (o *PutDevicesUisprsIDFirewallNatsRulesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesUisprsIDFirewallNatsRulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesUisprsIDFirewallNatsRulesInternalServerError creates a PutDevicesUisprsIDFirewallNatsRulesInternalServerError with default headers values
func NewPutDevicesUisprsIDFirewallNatsRulesInternalServerError() *PutDevicesUisprsIDFirewallNatsRulesInternalServerError {
	return &PutDevicesUisprsIDFirewallNatsRulesInternalServerError{}
}

/* PutDevicesUisprsIDFirewallNatsRulesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutDevicesUisprsIDFirewallNatsRulesInternalServerError struct {
	Payload *models.Error
}

func (o *PutDevicesUisprsIDFirewallNatsRulesInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /devices/uisprs/{id}/firewall/nats/rules][%d] putDevicesUisprsIdFirewallNatsRulesInternalServerError  %+v", 500, o.Payload)
}
func (o *PutDevicesUisprsIDFirewallNatsRulesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesUisprsIDFirewallNatsRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}