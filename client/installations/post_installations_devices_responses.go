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

// PostInstallationsDevicesReader is a Reader for the PostInstallationsDevices structure.
type PostInstallationsDevicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostInstallationsDevicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostInstallationsDevicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostInstallationsDevicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostInstallationsDevicesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostInstallationsDevicesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostInstallationsDevicesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostInstallationsDevicesOK creates a PostInstallationsDevicesOK with default headers values
func NewPostInstallationsDevicesOK() *PostInstallationsDevicesOK {
	return &PostInstallationsDevicesOK{}
}

/* PostInstallationsDevicesOK describes a response with status code 200, with default header values.

Installation device.
*/
type PostInstallationsDevicesOK struct {
	Payload *models.InstallationDevice
}

func (o *PostInstallationsDevicesOK) Error() string {
	return fmt.Sprintf("[POST /installations/devices][%d] postInstallationsDevicesOK  %+v", 200, o.Payload)
}
func (o *PostInstallationsDevicesOK) GetPayload() *models.InstallationDevice {
	return o.Payload
}

func (o *PostInstallationsDevicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InstallationDevice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInstallationsDevicesUnauthorized creates a PostInstallationsDevicesUnauthorized with default headers values
func NewPostInstallationsDevicesUnauthorized() *PostInstallationsDevicesUnauthorized {
	return &PostInstallationsDevicesUnauthorized{}
}

/* PostInstallationsDevicesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostInstallationsDevicesUnauthorized struct {
	Payload *models.Error
}

func (o *PostInstallationsDevicesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /installations/devices][%d] postInstallationsDevicesUnauthorized  %+v", 401, o.Payload)
}
func (o *PostInstallationsDevicesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostInstallationsDevicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInstallationsDevicesForbidden creates a PostInstallationsDevicesForbidden with default headers values
func NewPostInstallationsDevicesForbidden() *PostInstallationsDevicesForbidden {
	return &PostInstallationsDevicesForbidden{}
}

/* PostInstallationsDevicesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostInstallationsDevicesForbidden struct {
	Payload *models.Error
}

func (o *PostInstallationsDevicesForbidden) Error() string {
	return fmt.Sprintf("[POST /installations/devices][%d] postInstallationsDevicesForbidden  %+v", 403, o.Payload)
}
func (o *PostInstallationsDevicesForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostInstallationsDevicesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInstallationsDevicesInternalServerError creates a PostInstallationsDevicesInternalServerError with default headers values
func NewPostInstallationsDevicesInternalServerError() *PostInstallationsDevicesInternalServerError {
	return &PostInstallationsDevicesInternalServerError{}
}

/* PostInstallationsDevicesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostInstallationsDevicesInternalServerError struct {
	Payload *models.Error
}

func (o *PostInstallationsDevicesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /installations/devices][%d] postInstallationsDevicesInternalServerError  %+v", 500, o.Payload)
}
func (o *PostInstallationsDevicesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostInstallationsDevicesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInstallationsDevicesServiceUnavailable creates a PostInstallationsDevicesServiceUnavailable with default headers values
func NewPostInstallationsDevicesServiceUnavailable() *PostInstallationsDevicesServiceUnavailable {
	return &PostInstallationsDevicesServiceUnavailable{}
}

/* PostInstallationsDevicesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable
*/
type PostInstallationsDevicesServiceUnavailable struct {
	Payload *models.Error
}

func (o *PostInstallationsDevicesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /installations/devices][%d] postInstallationsDevicesServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PostInstallationsDevicesServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostInstallationsDevicesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}