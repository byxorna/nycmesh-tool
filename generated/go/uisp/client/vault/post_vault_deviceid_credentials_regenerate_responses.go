// Code generated by go-swagger; DO NOT EDIT.

package vault

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// PostVaultDeviceidCredentialsRegenerateReader is a Reader for the PostVaultDeviceidCredentialsRegenerate structure.
type PostVaultDeviceidCredentialsRegenerateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVaultDeviceidCredentialsRegenerateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostVaultDeviceidCredentialsRegenerateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostVaultDeviceidCredentialsRegenerateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostVaultDeviceidCredentialsRegenerateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostVaultDeviceidCredentialsRegenerateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostVaultDeviceidCredentialsRegenerateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewPostVaultDeviceidCredentialsRegeneratePreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostVaultDeviceidCredentialsRegenerateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostVaultDeviceidCredentialsRegenerateOK creates a PostVaultDeviceidCredentialsRegenerateOK with default headers values
func NewPostVaultDeviceidCredentialsRegenerateOK() *PostVaultDeviceidCredentialsRegenerateOK {
	return &PostVaultDeviceidCredentialsRegenerateOK{}
}

/* PostVaultDeviceidCredentialsRegenerateOK describes a response with status code 200, with default header values.

Successful
*/
type PostVaultDeviceidCredentialsRegenerateOK struct {
	Payload *models.DevicesCredentials
}

func (o *PostVaultDeviceidCredentialsRegenerateOK) Error() string {
	return fmt.Sprintf("[POST /vault/{deviceId}/credentials/regenerate][%d] postVaultDeviceidCredentialsRegenerateOK  %+v", 200, o.Payload)
}
func (o *PostVaultDeviceidCredentialsRegenerateOK) GetPayload() *models.DevicesCredentials {
	return o.Payload
}

func (o *PostVaultDeviceidCredentialsRegenerateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DevicesCredentials)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVaultDeviceidCredentialsRegenerateBadRequest creates a PostVaultDeviceidCredentialsRegenerateBadRequest with default headers values
func NewPostVaultDeviceidCredentialsRegenerateBadRequest() *PostVaultDeviceidCredentialsRegenerateBadRequest {
	return &PostVaultDeviceidCredentialsRegenerateBadRequest{}
}

/* PostVaultDeviceidCredentialsRegenerateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostVaultDeviceidCredentialsRegenerateBadRequest struct {
	Payload *models.Error
}

func (o *PostVaultDeviceidCredentialsRegenerateBadRequest) Error() string {
	return fmt.Sprintf("[POST /vault/{deviceId}/credentials/regenerate][%d] postVaultDeviceidCredentialsRegenerateBadRequest  %+v", 400, o.Payload)
}
func (o *PostVaultDeviceidCredentialsRegenerateBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostVaultDeviceidCredentialsRegenerateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVaultDeviceidCredentialsRegenerateUnauthorized creates a PostVaultDeviceidCredentialsRegenerateUnauthorized with default headers values
func NewPostVaultDeviceidCredentialsRegenerateUnauthorized() *PostVaultDeviceidCredentialsRegenerateUnauthorized {
	return &PostVaultDeviceidCredentialsRegenerateUnauthorized{}
}

/* PostVaultDeviceidCredentialsRegenerateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostVaultDeviceidCredentialsRegenerateUnauthorized struct {
	Payload *models.Error
}

func (o *PostVaultDeviceidCredentialsRegenerateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /vault/{deviceId}/credentials/regenerate][%d] postVaultDeviceidCredentialsRegenerateUnauthorized  %+v", 401, o.Payload)
}
func (o *PostVaultDeviceidCredentialsRegenerateUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostVaultDeviceidCredentialsRegenerateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVaultDeviceidCredentialsRegenerateForbidden creates a PostVaultDeviceidCredentialsRegenerateForbidden with default headers values
func NewPostVaultDeviceidCredentialsRegenerateForbidden() *PostVaultDeviceidCredentialsRegenerateForbidden {
	return &PostVaultDeviceidCredentialsRegenerateForbidden{}
}

/* PostVaultDeviceidCredentialsRegenerateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostVaultDeviceidCredentialsRegenerateForbidden struct {
	Payload *models.Error
}

func (o *PostVaultDeviceidCredentialsRegenerateForbidden) Error() string {
	return fmt.Sprintf("[POST /vault/{deviceId}/credentials/regenerate][%d] postVaultDeviceidCredentialsRegenerateForbidden  %+v", 403, o.Payload)
}
func (o *PostVaultDeviceidCredentialsRegenerateForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostVaultDeviceidCredentialsRegenerateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVaultDeviceidCredentialsRegenerateNotFound creates a PostVaultDeviceidCredentialsRegenerateNotFound with default headers values
func NewPostVaultDeviceidCredentialsRegenerateNotFound() *PostVaultDeviceidCredentialsRegenerateNotFound {
	return &PostVaultDeviceidCredentialsRegenerateNotFound{}
}

/* PostVaultDeviceidCredentialsRegenerateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostVaultDeviceidCredentialsRegenerateNotFound struct {
	Payload *models.Error
}

func (o *PostVaultDeviceidCredentialsRegenerateNotFound) Error() string {
	return fmt.Sprintf("[POST /vault/{deviceId}/credentials/regenerate][%d] postVaultDeviceidCredentialsRegenerateNotFound  %+v", 404, o.Payload)
}
func (o *PostVaultDeviceidCredentialsRegenerateNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostVaultDeviceidCredentialsRegenerateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVaultDeviceidCredentialsRegeneratePreconditionFailed creates a PostVaultDeviceidCredentialsRegeneratePreconditionFailed with default headers values
func NewPostVaultDeviceidCredentialsRegeneratePreconditionFailed() *PostVaultDeviceidCredentialsRegeneratePreconditionFailed {
	return &PostVaultDeviceidCredentialsRegeneratePreconditionFailed{}
}

/* PostVaultDeviceidCredentialsRegeneratePreconditionFailed describes a response with status code 412, with default header values.

Precondition Failed
*/
type PostVaultDeviceidCredentialsRegeneratePreconditionFailed struct {
	Payload *models.Error
}

func (o *PostVaultDeviceidCredentialsRegeneratePreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /vault/{deviceId}/credentials/regenerate][%d] postVaultDeviceidCredentialsRegeneratePreconditionFailed  %+v", 412, o.Payload)
}
func (o *PostVaultDeviceidCredentialsRegeneratePreconditionFailed) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostVaultDeviceidCredentialsRegeneratePreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVaultDeviceidCredentialsRegenerateInternalServerError creates a PostVaultDeviceidCredentialsRegenerateInternalServerError with default headers values
func NewPostVaultDeviceidCredentialsRegenerateInternalServerError() *PostVaultDeviceidCredentialsRegenerateInternalServerError {
	return &PostVaultDeviceidCredentialsRegenerateInternalServerError{}
}

/* PostVaultDeviceidCredentialsRegenerateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostVaultDeviceidCredentialsRegenerateInternalServerError struct {
	Payload *models.Error
}

func (o *PostVaultDeviceidCredentialsRegenerateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /vault/{deviceId}/credentials/regenerate][%d] postVaultDeviceidCredentialsRegenerateInternalServerError  %+v", 500, o.Payload)
}
func (o *PostVaultDeviceidCredentialsRegenerateInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostVaultDeviceidCredentialsRegenerateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
