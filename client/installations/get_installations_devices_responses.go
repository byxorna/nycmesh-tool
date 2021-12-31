// Code generated by go-swagger; DO NOT EDIT.

package installations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/models"
)

// GetInstallationsDevicesReader is a Reader for the GetInstallationsDevices structure.
type GetInstallationsDevicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstallationsDevicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInstallationsDevicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetInstallationsDevicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInstallationsDevicesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInstallationsDevicesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetInstallationsDevicesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInstallationsDevicesOK creates a GetInstallationsDevicesOK with default headers values
func NewGetInstallationsDevicesOK() *GetInstallationsDevicesOK {
	return &GetInstallationsDevicesOK{}
}

/* GetInstallationsDevicesOK describes a response with status code 200, with default header values.

Successful
*/
type GetInstallationsDevicesOK struct {
	Payload models.InstallationDeviceList
}

func (o *GetInstallationsDevicesOK) Error() string {
	return fmt.Sprintf("[GET /installations/devices][%d] getInstallationsDevicesOK  %+v", 200, o.Payload)
}
func (o *GetInstallationsDevicesOK) GetPayload() models.InstallationDeviceList {
	return o.Payload
}

func (o *GetInstallationsDevicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallationsDevicesUnauthorized creates a GetInstallationsDevicesUnauthorized with default headers values
func NewGetInstallationsDevicesUnauthorized() *GetInstallationsDevicesUnauthorized {
	return &GetInstallationsDevicesUnauthorized{}
}

/* GetInstallationsDevicesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetInstallationsDevicesUnauthorized struct {
	Payload *models.Error
}

func (o *GetInstallationsDevicesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /installations/devices][%d] getInstallationsDevicesUnauthorized  %+v", 401, o.Payload)
}
func (o *GetInstallationsDevicesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetInstallationsDevicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallationsDevicesForbidden creates a GetInstallationsDevicesForbidden with default headers values
func NewGetInstallationsDevicesForbidden() *GetInstallationsDevicesForbidden {
	return &GetInstallationsDevicesForbidden{}
}

/* GetInstallationsDevicesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetInstallationsDevicesForbidden struct {
	Payload *models.Error
}

func (o *GetInstallationsDevicesForbidden) Error() string {
	return fmt.Sprintf("[GET /installations/devices][%d] getInstallationsDevicesForbidden  %+v", 403, o.Payload)
}
func (o *GetInstallationsDevicesForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetInstallationsDevicesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallationsDevicesInternalServerError creates a GetInstallationsDevicesInternalServerError with default headers values
func NewGetInstallationsDevicesInternalServerError() *GetInstallationsDevicesInternalServerError {
	return &GetInstallationsDevicesInternalServerError{}
}

/* GetInstallationsDevicesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetInstallationsDevicesInternalServerError struct {
	Payload *models.Error
}

func (o *GetInstallationsDevicesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /installations/devices][%d] getInstallationsDevicesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetInstallationsDevicesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetInstallationsDevicesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallationsDevicesServiceUnavailable creates a GetInstallationsDevicesServiceUnavailable with default headers values
func NewGetInstallationsDevicesServiceUnavailable() *GetInstallationsDevicesServiceUnavailable {
	return &GetInstallationsDevicesServiceUnavailable{}
}

/* GetInstallationsDevicesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable
*/
type GetInstallationsDevicesServiceUnavailable struct {
	Payload *models.Error
}

func (o *GetInstallationsDevicesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /installations/devices][%d] getInstallationsDevicesServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetInstallationsDevicesServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetInstallationsDevicesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}