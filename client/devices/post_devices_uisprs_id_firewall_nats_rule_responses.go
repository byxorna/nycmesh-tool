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

// PostDevicesUisprsIDFirewallNatsRuleReader is a Reader for the PostDevicesUisprsIDFirewallNatsRule structure.
type PostDevicesUisprsIDFirewallNatsRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDevicesUisprsIDFirewallNatsRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDevicesUisprsIDFirewallNatsRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDevicesUisprsIDFirewallNatsRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDevicesUisprsIDFirewallNatsRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostDevicesUisprsIDFirewallNatsRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDevicesUisprsIDFirewallNatsRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDevicesUisprsIDFirewallNatsRuleOK creates a PostDevicesUisprsIDFirewallNatsRuleOK with default headers values
func NewPostDevicesUisprsIDFirewallNatsRuleOK() *PostDevicesUisprsIDFirewallNatsRuleOK {
	return &PostDevicesUisprsIDFirewallNatsRuleOK{}
}

/* PostDevicesUisprsIDFirewallNatsRuleOK describes a response with status code 200, with default header values.

Successful
*/
type PostDevicesUisprsIDFirewallNatsRuleOK struct {
	Payload *models.Model106
}

func (o *PostDevicesUisprsIDFirewallNatsRuleOK) Error() string {
	return fmt.Sprintf("[POST /devices/uisprs/{id}/firewall/nats/rule][%d] postDevicesUisprsIdFirewallNatsRuleOK  %+v", 200, o.Payload)
}
func (o *PostDevicesUisprsIDFirewallNatsRuleOK) GetPayload() *models.Model106 {
	return o.Payload
}

func (o *PostDevicesUisprsIDFirewallNatsRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Model106)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesUisprsIDFirewallNatsRuleBadRequest creates a PostDevicesUisprsIDFirewallNatsRuleBadRequest with default headers values
func NewPostDevicesUisprsIDFirewallNatsRuleBadRequest() *PostDevicesUisprsIDFirewallNatsRuleBadRequest {
	return &PostDevicesUisprsIDFirewallNatsRuleBadRequest{}
}

/* PostDevicesUisprsIDFirewallNatsRuleBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostDevicesUisprsIDFirewallNatsRuleBadRequest struct {
	Payload *models.Error
}

func (o *PostDevicesUisprsIDFirewallNatsRuleBadRequest) Error() string {
	return fmt.Sprintf("[POST /devices/uisprs/{id}/firewall/nats/rule][%d] postDevicesUisprsIdFirewallNatsRuleBadRequest  %+v", 400, o.Payload)
}
func (o *PostDevicesUisprsIDFirewallNatsRuleBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesUisprsIDFirewallNatsRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesUisprsIDFirewallNatsRuleUnauthorized creates a PostDevicesUisprsIDFirewallNatsRuleUnauthorized with default headers values
func NewPostDevicesUisprsIDFirewallNatsRuleUnauthorized() *PostDevicesUisprsIDFirewallNatsRuleUnauthorized {
	return &PostDevicesUisprsIDFirewallNatsRuleUnauthorized{}
}

/* PostDevicesUisprsIDFirewallNatsRuleUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostDevicesUisprsIDFirewallNatsRuleUnauthorized struct {
	Payload *models.Error
}

func (o *PostDevicesUisprsIDFirewallNatsRuleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /devices/uisprs/{id}/firewall/nats/rule][%d] postDevicesUisprsIdFirewallNatsRuleUnauthorized  %+v", 401, o.Payload)
}
func (o *PostDevicesUisprsIDFirewallNatsRuleUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesUisprsIDFirewallNatsRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesUisprsIDFirewallNatsRuleNotFound creates a PostDevicesUisprsIDFirewallNatsRuleNotFound with default headers values
func NewPostDevicesUisprsIDFirewallNatsRuleNotFound() *PostDevicesUisprsIDFirewallNatsRuleNotFound {
	return &PostDevicesUisprsIDFirewallNatsRuleNotFound{}
}

/* PostDevicesUisprsIDFirewallNatsRuleNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostDevicesUisprsIDFirewallNatsRuleNotFound struct {
	Payload *models.Error
}

func (o *PostDevicesUisprsIDFirewallNatsRuleNotFound) Error() string {
	return fmt.Sprintf("[POST /devices/uisprs/{id}/firewall/nats/rule][%d] postDevicesUisprsIdFirewallNatsRuleNotFound  %+v", 404, o.Payload)
}
func (o *PostDevicesUisprsIDFirewallNatsRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesUisprsIDFirewallNatsRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesUisprsIDFirewallNatsRuleInternalServerError creates a PostDevicesUisprsIDFirewallNatsRuleInternalServerError with default headers values
func NewPostDevicesUisprsIDFirewallNatsRuleInternalServerError() *PostDevicesUisprsIDFirewallNatsRuleInternalServerError {
	return &PostDevicesUisprsIDFirewallNatsRuleInternalServerError{}
}

/* PostDevicesUisprsIDFirewallNatsRuleInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostDevicesUisprsIDFirewallNatsRuleInternalServerError struct {
	Payload *models.Error
}

func (o *PostDevicesUisprsIDFirewallNatsRuleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /devices/uisprs/{id}/firewall/nats/rule][%d] postDevicesUisprsIdFirewallNatsRuleInternalServerError  %+v", 500, o.Payload)
}
func (o *PostDevicesUisprsIDFirewallNatsRuleInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesUisprsIDFirewallNatsRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
