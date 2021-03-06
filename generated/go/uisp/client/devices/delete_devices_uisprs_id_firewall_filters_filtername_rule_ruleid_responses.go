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

// DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidReader is a Reader for the DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleid structure.
type DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidOK creates a DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidOK with default headers values
func NewDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidOK() *DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidOK {
	return &DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidOK{}
}

/* DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidOK describes a response with status code 200, with default header values.

Successful
*/
type DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidOK struct {
	Payload *models.Status
}

func (o *DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidOK) Error() string {
	return fmt.Sprintf("[DELETE /devices/uisprs/{id}/firewall/filters/{filterName}/rule/{ruleId}][%d] deleteDevicesUisprsIdFirewallFiltersFilternameRuleRuleidOK  %+v", 200, o.Payload)
}
func (o *DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidBadRequest creates a DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidBadRequest with default headers values
func NewDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidBadRequest() *DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidBadRequest {
	return &DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidBadRequest{}
}

/* DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidBadRequest struct {
	Payload *models.Error
}

func (o *DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /devices/uisprs/{id}/firewall/filters/{filterName}/rule/{ruleId}][%d] deleteDevicesUisprsIdFirewallFiltersFilternameRuleRuleidBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidUnauthorized creates a DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidUnauthorized with default headers values
func NewDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidUnauthorized() *DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidUnauthorized {
	return &DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidUnauthorized{}
}

/* DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /devices/uisprs/{id}/firewall/filters/{filterName}/rule/{ruleId}][%d] deleteDevicesUisprsIdFirewallFiltersFilternameRuleRuleidUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidNotFound creates a DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidNotFound with default headers values
func NewDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidNotFound() *DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidNotFound {
	return &DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidNotFound{}
}

/* DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidNotFound struct {
	Payload *models.Error
}

func (o *DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /devices/uisprs/{id}/firewall/filters/{filterName}/rule/{ruleId}][%d] deleteDevicesUisprsIdFirewallFiltersFilternameRuleRuleidNotFound  %+v", 404, o.Payload)
}
func (o *DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidInternalServerError creates a DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidInternalServerError with default headers values
func NewDeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidInternalServerError() *DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidInternalServerError {
	return &DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidInternalServerError{}
}

/* DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /devices/uisprs/{id}/firewall/filters/{filterName}/rule/{ruleId}][%d] deleteDevicesUisprsIdFirewallFiltersFilternameRuleRuleidInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDevicesUisprsIDFirewallFiltersFilternameRuleRuleidInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
