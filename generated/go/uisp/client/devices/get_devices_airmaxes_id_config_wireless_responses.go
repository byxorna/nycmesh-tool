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

// GetDevicesAirmaxesIDConfigWirelessReader is a Reader for the GetDevicesAirmaxesIDConfigWireless structure.
type GetDevicesAirmaxesIDConfigWirelessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesAirmaxesIDConfigWirelessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesAirmaxesIDConfigWirelessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesAirmaxesIDConfigWirelessBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesAirmaxesIDConfigWirelessUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesAirmaxesIDConfigWirelessForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesAirmaxesIDConfigWirelessNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesAirmaxesIDConfigWirelessInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDevicesAirmaxesIDConfigWirelessOK creates a GetDevicesAirmaxesIDConfigWirelessOK with default headers values
func NewGetDevicesAirmaxesIDConfigWirelessOK() *GetDevicesAirmaxesIDConfigWirelessOK {
	return &GetDevicesAirmaxesIDConfigWirelessOK{}
}

/* GetDevicesAirmaxesIDConfigWirelessOK describes a response with status code 200, with default header values.

Successful
*/
type GetDevicesAirmaxesIDConfigWirelessOK struct {
	Payload *models.AirMaxWifiConfig
}

func (o *GetDevicesAirmaxesIDConfigWirelessOK) Error() string {
	return fmt.Sprintf("[GET /devices/airmaxes/{id}/config/wireless][%d] getDevicesAirmaxesIdConfigWirelessOK  %+v", 200, o.Payload)
}
func (o *GetDevicesAirmaxesIDConfigWirelessOK) GetPayload() *models.AirMaxWifiConfig {
	return o.Payload
}

func (o *GetDevicesAirmaxesIDConfigWirelessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AirMaxWifiConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAirmaxesIDConfigWirelessBadRequest creates a GetDevicesAirmaxesIDConfigWirelessBadRequest with default headers values
func NewGetDevicesAirmaxesIDConfigWirelessBadRequest() *GetDevicesAirmaxesIDConfigWirelessBadRequest {
	return &GetDevicesAirmaxesIDConfigWirelessBadRequest{}
}

/* GetDevicesAirmaxesIDConfigWirelessBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDevicesAirmaxesIDConfigWirelessBadRequest struct {
	Payload *models.Error
}

func (o *GetDevicesAirmaxesIDConfigWirelessBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/airmaxes/{id}/config/wireless][%d] getDevicesAirmaxesIdConfigWirelessBadRequest  %+v", 400, o.Payload)
}
func (o *GetDevicesAirmaxesIDConfigWirelessBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAirmaxesIDConfigWirelessBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAirmaxesIDConfigWirelessUnauthorized creates a GetDevicesAirmaxesIDConfigWirelessUnauthorized with default headers values
func NewGetDevicesAirmaxesIDConfigWirelessUnauthorized() *GetDevicesAirmaxesIDConfigWirelessUnauthorized {
	return &GetDevicesAirmaxesIDConfigWirelessUnauthorized{}
}

/* GetDevicesAirmaxesIDConfigWirelessUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDevicesAirmaxesIDConfigWirelessUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesAirmaxesIDConfigWirelessUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/airmaxes/{id}/config/wireless][%d] getDevicesAirmaxesIdConfigWirelessUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDevicesAirmaxesIDConfigWirelessUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAirmaxesIDConfigWirelessUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAirmaxesIDConfigWirelessForbidden creates a GetDevicesAirmaxesIDConfigWirelessForbidden with default headers values
func NewGetDevicesAirmaxesIDConfigWirelessForbidden() *GetDevicesAirmaxesIDConfigWirelessForbidden {
	return &GetDevicesAirmaxesIDConfigWirelessForbidden{}
}

/* GetDevicesAirmaxesIDConfigWirelessForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDevicesAirmaxesIDConfigWirelessForbidden struct {
	Payload *models.Error
}

func (o *GetDevicesAirmaxesIDConfigWirelessForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/airmaxes/{id}/config/wireless][%d] getDevicesAirmaxesIdConfigWirelessForbidden  %+v", 403, o.Payload)
}
func (o *GetDevicesAirmaxesIDConfigWirelessForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAirmaxesIDConfigWirelessForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAirmaxesIDConfigWirelessNotFound creates a GetDevicesAirmaxesIDConfigWirelessNotFound with default headers values
func NewGetDevicesAirmaxesIDConfigWirelessNotFound() *GetDevicesAirmaxesIDConfigWirelessNotFound {
	return &GetDevicesAirmaxesIDConfigWirelessNotFound{}
}

/* GetDevicesAirmaxesIDConfigWirelessNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDevicesAirmaxesIDConfigWirelessNotFound struct {
	Payload *models.Error
}

func (o *GetDevicesAirmaxesIDConfigWirelessNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/airmaxes/{id}/config/wireless][%d] getDevicesAirmaxesIdConfigWirelessNotFound  %+v", 404, o.Payload)
}
func (o *GetDevicesAirmaxesIDConfigWirelessNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAirmaxesIDConfigWirelessNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAirmaxesIDConfigWirelessInternalServerError creates a GetDevicesAirmaxesIDConfigWirelessInternalServerError with default headers values
func NewGetDevicesAirmaxesIDConfigWirelessInternalServerError() *GetDevicesAirmaxesIDConfigWirelessInternalServerError {
	return &GetDevicesAirmaxesIDConfigWirelessInternalServerError{}
}

/* GetDevicesAirmaxesIDConfigWirelessInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDevicesAirmaxesIDConfigWirelessInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesAirmaxesIDConfigWirelessInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/airmaxes/{id}/config/wireless][%d] getDevicesAirmaxesIdConfigWirelessInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDevicesAirmaxesIDConfigWirelessInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAirmaxesIDConfigWirelessInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
