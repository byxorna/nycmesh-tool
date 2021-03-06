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

// PutDevicesDeviceidBackupsReader is a Reader for the PutDevicesDeviceidBackups structure.
type PutDevicesDeviceidBackupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDevicesDeviceidBackupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutDevicesDeviceidBackupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutDevicesDeviceidBackupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutDevicesDeviceidBackupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutDevicesDeviceidBackupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutDevicesDeviceidBackupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutDevicesDeviceidBackupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutDevicesDeviceidBackupsOK creates a PutDevicesDeviceidBackupsOK with default headers values
func NewPutDevicesDeviceidBackupsOK() *PutDevicesDeviceidBackupsOK {
	return &PutDevicesDeviceidBackupsOK{}
}

/* PutDevicesDeviceidBackupsOK describes a response with status code 200, with default header values.

Successful
*/
type PutDevicesDeviceidBackupsOK struct {
	Payload *models.DeviceBackup
}

func (o *PutDevicesDeviceidBackupsOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{deviceId}/backups][%d] putDevicesDeviceidBackupsOK  %+v", 200, o.Payload)
}
func (o *PutDevicesDeviceidBackupsOK) GetPayload() *models.DeviceBackup {
	return o.Payload
}

func (o *PutDevicesDeviceidBackupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBackup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesDeviceidBackupsBadRequest creates a PutDevicesDeviceidBackupsBadRequest with default headers values
func NewPutDevicesDeviceidBackupsBadRequest() *PutDevicesDeviceidBackupsBadRequest {
	return &PutDevicesDeviceidBackupsBadRequest{}
}

/* PutDevicesDeviceidBackupsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutDevicesDeviceidBackupsBadRequest struct {
	Payload *models.Error
}

func (o *PutDevicesDeviceidBackupsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /devices/{deviceId}/backups][%d] putDevicesDeviceidBackupsBadRequest  %+v", 400, o.Payload)
}
func (o *PutDevicesDeviceidBackupsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesDeviceidBackupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesDeviceidBackupsUnauthorized creates a PutDevicesDeviceidBackupsUnauthorized with default headers values
func NewPutDevicesDeviceidBackupsUnauthorized() *PutDevicesDeviceidBackupsUnauthorized {
	return &PutDevicesDeviceidBackupsUnauthorized{}
}

/* PutDevicesDeviceidBackupsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutDevicesDeviceidBackupsUnauthorized struct {
	Payload *models.Error
}

func (o *PutDevicesDeviceidBackupsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /devices/{deviceId}/backups][%d] putDevicesDeviceidBackupsUnauthorized  %+v", 401, o.Payload)
}
func (o *PutDevicesDeviceidBackupsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesDeviceidBackupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesDeviceidBackupsForbidden creates a PutDevicesDeviceidBackupsForbidden with default headers values
func NewPutDevicesDeviceidBackupsForbidden() *PutDevicesDeviceidBackupsForbidden {
	return &PutDevicesDeviceidBackupsForbidden{}
}

/* PutDevicesDeviceidBackupsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutDevicesDeviceidBackupsForbidden struct {
	Payload *models.Error
}

func (o *PutDevicesDeviceidBackupsForbidden) Error() string {
	return fmt.Sprintf("[PUT /devices/{deviceId}/backups][%d] putDevicesDeviceidBackupsForbidden  %+v", 403, o.Payload)
}
func (o *PutDevicesDeviceidBackupsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesDeviceidBackupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesDeviceidBackupsNotFound creates a PutDevicesDeviceidBackupsNotFound with default headers values
func NewPutDevicesDeviceidBackupsNotFound() *PutDevicesDeviceidBackupsNotFound {
	return &PutDevicesDeviceidBackupsNotFound{}
}

/* PutDevicesDeviceidBackupsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutDevicesDeviceidBackupsNotFound struct {
	Payload *models.Error
}

func (o *PutDevicesDeviceidBackupsNotFound) Error() string {
	return fmt.Sprintf("[PUT /devices/{deviceId}/backups][%d] putDevicesDeviceidBackupsNotFound  %+v", 404, o.Payload)
}
func (o *PutDevicesDeviceidBackupsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesDeviceidBackupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDevicesDeviceidBackupsInternalServerError creates a PutDevicesDeviceidBackupsInternalServerError with default headers values
func NewPutDevicesDeviceidBackupsInternalServerError() *PutDevicesDeviceidBackupsInternalServerError {
	return &PutDevicesDeviceidBackupsInternalServerError{}
}

/* PutDevicesDeviceidBackupsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutDevicesDeviceidBackupsInternalServerError struct {
	Payload *models.Error
}

func (o *PutDevicesDeviceidBackupsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /devices/{deviceId}/backups][%d] putDevicesDeviceidBackupsInternalServerError  %+v", 500, o.Payload)
}
func (o *PutDevicesDeviceidBackupsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutDevicesDeviceidBackupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
