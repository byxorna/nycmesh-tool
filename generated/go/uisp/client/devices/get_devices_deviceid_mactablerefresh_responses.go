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

// GetDevicesDeviceidMactablerefreshReader is a Reader for the GetDevicesDeviceidMactablerefresh structure.
type GetDevicesDeviceidMactablerefreshReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesDeviceidMactablerefreshReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesDeviceidMactablerefreshOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesDeviceidMactablerefreshBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesDeviceidMactablerefreshUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesDeviceidMactablerefreshForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesDeviceidMactablerefreshNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesDeviceidMactablerefreshInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDevicesDeviceidMactablerefreshOK creates a GetDevicesDeviceidMactablerefreshOK with default headers values
func NewGetDevicesDeviceidMactablerefreshOK() *GetDevicesDeviceidMactablerefreshOK {
	return &GetDevicesDeviceidMactablerefreshOK{}
}

/* GetDevicesDeviceidMactablerefreshOK describes a response with status code 200, with default header values.

Successful
*/
type GetDevicesDeviceidMactablerefreshOK struct {
	Payload *models.Status
}

func (o *GetDevicesDeviceidMactablerefreshOK) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/mac-table-refresh][%d] getDevicesDeviceidMactablerefreshOK  %+v", 200, o.Payload)
}
func (o *GetDevicesDeviceidMactablerefreshOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *GetDevicesDeviceidMactablerefreshOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesDeviceidMactablerefreshBadRequest creates a GetDevicesDeviceidMactablerefreshBadRequest with default headers values
func NewGetDevicesDeviceidMactablerefreshBadRequest() *GetDevicesDeviceidMactablerefreshBadRequest {
	return &GetDevicesDeviceidMactablerefreshBadRequest{}
}

/* GetDevicesDeviceidMactablerefreshBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDevicesDeviceidMactablerefreshBadRequest struct {
	Payload *models.Error
}

func (o *GetDevicesDeviceidMactablerefreshBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/mac-table-refresh][%d] getDevicesDeviceidMactablerefreshBadRequest  %+v", 400, o.Payload)
}
func (o *GetDevicesDeviceidMactablerefreshBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesDeviceidMactablerefreshBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesDeviceidMactablerefreshUnauthorized creates a GetDevicesDeviceidMactablerefreshUnauthorized with default headers values
func NewGetDevicesDeviceidMactablerefreshUnauthorized() *GetDevicesDeviceidMactablerefreshUnauthorized {
	return &GetDevicesDeviceidMactablerefreshUnauthorized{}
}

/* GetDevicesDeviceidMactablerefreshUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDevicesDeviceidMactablerefreshUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesDeviceidMactablerefreshUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/mac-table-refresh][%d] getDevicesDeviceidMactablerefreshUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDevicesDeviceidMactablerefreshUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesDeviceidMactablerefreshUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesDeviceidMactablerefreshForbidden creates a GetDevicesDeviceidMactablerefreshForbidden with default headers values
func NewGetDevicesDeviceidMactablerefreshForbidden() *GetDevicesDeviceidMactablerefreshForbidden {
	return &GetDevicesDeviceidMactablerefreshForbidden{}
}

/* GetDevicesDeviceidMactablerefreshForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDevicesDeviceidMactablerefreshForbidden struct {
	Payload *models.Error
}

func (o *GetDevicesDeviceidMactablerefreshForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/mac-table-refresh][%d] getDevicesDeviceidMactablerefreshForbidden  %+v", 403, o.Payload)
}
func (o *GetDevicesDeviceidMactablerefreshForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesDeviceidMactablerefreshForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesDeviceidMactablerefreshNotFound creates a GetDevicesDeviceidMactablerefreshNotFound with default headers values
func NewGetDevicesDeviceidMactablerefreshNotFound() *GetDevicesDeviceidMactablerefreshNotFound {
	return &GetDevicesDeviceidMactablerefreshNotFound{}
}

/* GetDevicesDeviceidMactablerefreshNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDevicesDeviceidMactablerefreshNotFound struct {
	Payload *models.Error
}

func (o *GetDevicesDeviceidMactablerefreshNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/mac-table-refresh][%d] getDevicesDeviceidMactablerefreshNotFound  %+v", 404, o.Payload)
}
func (o *GetDevicesDeviceidMactablerefreshNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesDeviceidMactablerefreshNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesDeviceidMactablerefreshInternalServerError creates a GetDevicesDeviceidMactablerefreshInternalServerError with default headers values
func NewGetDevicesDeviceidMactablerefreshInternalServerError() *GetDevicesDeviceidMactablerefreshInternalServerError {
	return &GetDevicesDeviceidMactablerefreshInternalServerError{}
}

/* GetDevicesDeviceidMactablerefreshInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDevicesDeviceidMactablerefreshInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesDeviceidMactablerefreshInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}/mac-table-refresh][%d] getDevicesDeviceidMactablerefreshInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDevicesDeviceidMactablerefreshInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesDeviceidMactablerefreshInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}