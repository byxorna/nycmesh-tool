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

// GetDevicesUisprsIDFirewallSetsReader is a Reader for the GetDevicesUisprsIDFirewallSets structure.
type GetDevicesUisprsIDFirewallSetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesUisprsIDFirewallSetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesUisprsIDFirewallSetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesUisprsIDFirewallSetsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesUisprsIDFirewallSetsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesUisprsIDFirewallSetsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesUisprsIDFirewallSetsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDevicesUisprsIDFirewallSetsOK creates a GetDevicesUisprsIDFirewallSetsOK with default headers values
func NewGetDevicesUisprsIDFirewallSetsOK() *GetDevicesUisprsIDFirewallSetsOK {
	return &GetDevicesUisprsIDFirewallSetsOK{}
}

/* GetDevicesUisprsIDFirewallSetsOK describes a response with status code 200, with default header values.

Successful
*/
type GetDevicesUisprsIDFirewallSetsOK struct {
	Payload models.Model66
}

func (o *GetDevicesUisprsIDFirewallSetsOK) Error() string {
	return fmt.Sprintf("[GET /devices/uisprs/{id}/firewall/sets][%d] getDevicesUisprsIdFirewallSetsOK  %+v", 200, o.Payload)
}
func (o *GetDevicesUisprsIDFirewallSetsOK) GetPayload() models.Model66 {
	return o.Payload
}

func (o *GetDevicesUisprsIDFirewallSetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesUisprsIDFirewallSetsBadRequest creates a GetDevicesUisprsIDFirewallSetsBadRequest with default headers values
func NewGetDevicesUisprsIDFirewallSetsBadRequest() *GetDevicesUisprsIDFirewallSetsBadRequest {
	return &GetDevicesUisprsIDFirewallSetsBadRequest{}
}

/* GetDevicesUisprsIDFirewallSetsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDevicesUisprsIDFirewallSetsBadRequest struct {
	Payload *models.Error
}

func (o *GetDevicesUisprsIDFirewallSetsBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/uisprs/{id}/firewall/sets][%d] getDevicesUisprsIdFirewallSetsBadRequest  %+v", 400, o.Payload)
}
func (o *GetDevicesUisprsIDFirewallSetsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesUisprsIDFirewallSetsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesUisprsIDFirewallSetsUnauthorized creates a GetDevicesUisprsIDFirewallSetsUnauthorized with default headers values
func NewGetDevicesUisprsIDFirewallSetsUnauthorized() *GetDevicesUisprsIDFirewallSetsUnauthorized {
	return &GetDevicesUisprsIDFirewallSetsUnauthorized{}
}

/* GetDevicesUisprsIDFirewallSetsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDevicesUisprsIDFirewallSetsUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesUisprsIDFirewallSetsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/uisprs/{id}/firewall/sets][%d] getDevicesUisprsIdFirewallSetsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDevicesUisprsIDFirewallSetsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesUisprsIDFirewallSetsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesUisprsIDFirewallSetsNotFound creates a GetDevicesUisprsIDFirewallSetsNotFound with default headers values
func NewGetDevicesUisprsIDFirewallSetsNotFound() *GetDevicesUisprsIDFirewallSetsNotFound {
	return &GetDevicesUisprsIDFirewallSetsNotFound{}
}

/* GetDevicesUisprsIDFirewallSetsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDevicesUisprsIDFirewallSetsNotFound struct {
	Payload *models.Error
}

func (o *GetDevicesUisprsIDFirewallSetsNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/uisprs/{id}/firewall/sets][%d] getDevicesUisprsIdFirewallSetsNotFound  %+v", 404, o.Payload)
}
func (o *GetDevicesUisprsIDFirewallSetsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesUisprsIDFirewallSetsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesUisprsIDFirewallSetsInternalServerError creates a GetDevicesUisprsIDFirewallSetsInternalServerError with default headers values
func NewGetDevicesUisprsIDFirewallSetsInternalServerError() *GetDevicesUisprsIDFirewallSetsInternalServerError {
	return &GetDevicesUisprsIDFirewallSetsInternalServerError{}
}

/* GetDevicesUisprsIDFirewallSetsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDevicesUisprsIDFirewallSetsInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesUisprsIDFirewallSetsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/uisprs/{id}/firewall/sets][%d] getDevicesUisprsIdFirewallSetsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDevicesUisprsIDFirewallSetsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesUisprsIDFirewallSetsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}