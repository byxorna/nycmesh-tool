// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// PutNmsMaintenanceBackupReader is a Reader for the PutNmsMaintenanceBackup structure.
type PutNmsMaintenanceBackupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNmsMaintenanceBackupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutNmsMaintenanceBackupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutNmsMaintenanceBackupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutNmsMaintenanceBackupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutNmsMaintenanceBackupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutNmsMaintenanceBackupOK creates a PutNmsMaintenanceBackupOK with default headers values
func NewPutNmsMaintenanceBackupOK() *PutNmsMaintenanceBackupOK {
	return &PutNmsMaintenanceBackupOK{}
}

/* PutNmsMaintenanceBackupOK describes a response with status code 200, with default header values.

Successful
*/
type PutNmsMaintenanceBackupOK struct {
	Payload *models.Status
}

func (o *PutNmsMaintenanceBackupOK) Error() string {
	return fmt.Sprintf("[PUT /nms/maintenance/backup][%d] putNmsMaintenanceBackupOK  %+v", 200, o.Payload)
}
func (o *PutNmsMaintenanceBackupOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PutNmsMaintenanceBackupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNmsMaintenanceBackupUnauthorized creates a PutNmsMaintenanceBackupUnauthorized with default headers values
func NewPutNmsMaintenanceBackupUnauthorized() *PutNmsMaintenanceBackupUnauthorized {
	return &PutNmsMaintenanceBackupUnauthorized{}
}

/* PutNmsMaintenanceBackupUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutNmsMaintenanceBackupUnauthorized struct {
	Payload *models.Error
}

func (o *PutNmsMaintenanceBackupUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /nms/maintenance/backup][%d] putNmsMaintenanceBackupUnauthorized  %+v", 401, o.Payload)
}
func (o *PutNmsMaintenanceBackupUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNmsMaintenanceBackupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNmsMaintenanceBackupForbidden creates a PutNmsMaintenanceBackupForbidden with default headers values
func NewPutNmsMaintenanceBackupForbidden() *PutNmsMaintenanceBackupForbidden {
	return &PutNmsMaintenanceBackupForbidden{}
}

/* PutNmsMaintenanceBackupForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutNmsMaintenanceBackupForbidden struct {
	Payload *models.Error
}

func (o *PutNmsMaintenanceBackupForbidden) Error() string {
	return fmt.Sprintf("[PUT /nms/maintenance/backup][%d] putNmsMaintenanceBackupForbidden  %+v", 403, o.Payload)
}
func (o *PutNmsMaintenanceBackupForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNmsMaintenanceBackupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNmsMaintenanceBackupInternalServerError creates a PutNmsMaintenanceBackupInternalServerError with default headers values
func NewPutNmsMaintenanceBackupInternalServerError() *PutNmsMaintenanceBackupInternalServerError {
	return &PutNmsMaintenanceBackupInternalServerError{}
}

/* PutNmsMaintenanceBackupInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutNmsMaintenanceBackupInternalServerError struct {
	Payload *models.Error
}

func (o *PutNmsMaintenanceBackupInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /nms/maintenance/backup][%d] putNmsMaintenanceBackupInternalServerError  %+v", 500, o.Payload)
}
func (o *PutNmsMaintenanceBackupInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNmsMaintenanceBackupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
