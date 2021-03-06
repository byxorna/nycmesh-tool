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

// DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidReader is a Reader for the DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleid structure.
type DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidOK creates a DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidOK with default headers values
func NewDeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidOK() *DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidOK {
	return &DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidOK{}
}

/* DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidOK describes a response with status code 200, with default header values.

Successful
*/
type DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidOK struct {
	Payload *models.Status
}

func (o *DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidOK) Error() string {
	return fmt.Sprintf("[DELETE /devices/uisprs/{id}/firewall/mangles/{mangleName}/rule/{ruleId}][%d] deleteDevicesUisprsIdFirewallManglesManglenameRuleRuleidOK  %+v", 200, o.Payload)
}
func (o *DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidBadRequest creates a DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidBadRequest with default headers values
func NewDeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidBadRequest() *DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidBadRequest {
	return &DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidBadRequest{}
}

/* DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidBadRequest struct {
	Payload *models.Error
}

func (o *DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /devices/uisprs/{id}/firewall/mangles/{mangleName}/rule/{ruleId}][%d] deleteDevicesUisprsIdFirewallManglesManglenameRuleRuleidBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidUnauthorized creates a DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidUnauthorized with default headers values
func NewDeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidUnauthorized() *DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidUnauthorized {
	return &DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidUnauthorized{}
}

/* DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /devices/uisprs/{id}/firewall/mangles/{mangleName}/rule/{ruleId}][%d] deleteDevicesUisprsIdFirewallManglesManglenameRuleRuleidUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidNotFound creates a DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidNotFound with default headers values
func NewDeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidNotFound() *DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidNotFound {
	return &DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidNotFound{}
}

/* DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidNotFound struct {
	Payload *models.Error
}

func (o *DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /devices/uisprs/{id}/firewall/mangles/{mangleName}/rule/{ruleId}][%d] deleteDevicesUisprsIdFirewallManglesManglenameRuleRuleidNotFound  %+v", 404, o.Payload)
}
func (o *DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidInternalServerError creates a DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidInternalServerError with default headers values
func NewDeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidInternalServerError() *DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidInternalServerError {
	return &DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidInternalServerError{}
}

/* DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /devices/uisprs/{id}/firewall/mangles/{mangleName}/rule/{ruleId}][%d] deleteDevicesUisprsIdFirewallManglesManglenameRuleRuleidInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesUisprsIDFirewallManglesManglenameRuleRuleidInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
