// Code generated by go-swagger; DO NOT EDIT.

package vault

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/models"
)

// PostVaultCredentialsRegenerateReader is a Reader for the PostVaultCredentialsRegenerate structure.
type PostVaultCredentialsRegenerateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVaultCredentialsRegenerateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostVaultCredentialsRegenerateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostVaultCredentialsRegenerateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostVaultCredentialsRegenerateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostVaultCredentialsRegenerateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostVaultCredentialsRegenerateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewPostVaultCredentialsRegeneratePreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostVaultCredentialsRegenerateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostVaultCredentialsRegenerateOK creates a PostVaultCredentialsRegenerateOK with default headers values
func NewPostVaultCredentialsRegenerateOK() *PostVaultCredentialsRegenerateOK {
	return &PostVaultCredentialsRegenerateOK{}
}

/* PostVaultCredentialsRegenerateOK describes a response with status code 200, with default header values.

Successful
*/
type PostVaultCredentialsRegenerateOK struct {
	Payload *models.Status
}

func (o *PostVaultCredentialsRegenerateOK) Error() string {
	return fmt.Sprintf("[POST /vault/credentials/regenerate][%d] postVaultCredentialsRegenerateOK  %+v", 200, o.Payload)
}
func (o *PostVaultCredentialsRegenerateOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PostVaultCredentialsRegenerateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVaultCredentialsRegenerateBadRequest creates a PostVaultCredentialsRegenerateBadRequest with default headers values
func NewPostVaultCredentialsRegenerateBadRequest() *PostVaultCredentialsRegenerateBadRequest {
	return &PostVaultCredentialsRegenerateBadRequest{}
}

/* PostVaultCredentialsRegenerateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostVaultCredentialsRegenerateBadRequest struct {
	Payload *models.Error
}

func (o *PostVaultCredentialsRegenerateBadRequest) Error() string {
	return fmt.Sprintf("[POST /vault/credentials/regenerate][%d] postVaultCredentialsRegenerateBadRequest  %+v", 400, o.Payload)
}
func (o *PostVaultCredentialsRegenerateBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostVaultCredentialsRegenerateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVaultCredentialsRegenerateUnauthorized creates a PostVaultCredentialsRegenerateUnauthorized with default headers values
func NewPostVaultCredentialsRegenerateUnauthorized() *PostVaultCredentialsRegenerateUnauthorized {
	return &PostVaultCredentialsRegenerateUnauthorized{}
}

/* PostVaultCredentialsRegenerateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostVaultCredentialsRegenerateUnauthorized struct {
	Payload *models.Error
}

func (o *PostVaultCredentialsRegenerateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /vault/credentials/regenerate][%d] postVaultCredentialsRegenerateUnauthorized  %+v", 401, o.Payload)
}
func (o *PostVaultCredentialsRegenerateUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostVaultCredentialsRegenerateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVaultCredentialsRegenerateForbidden creates a PostVaultCredentialsRegenerateForbidden with default headers values
func NewPostVaultCredentialsRegenerateForbidden() *PostVaultCredentialsRegenerateForbidden {
	return &PostVaultCredentialsRegenerateForbidden{}
}

/* PostVaultCredentialsRegenerateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostVaultCredentialsRegenerateForbidden struct {
	Payload *models.Error
}

func (o *PostVaultCredentialsRegenerateForbidden) Error() string {
	return fmt.Sprintf("[POST /vault/credentials/regenerate][%d] postVaultCredentialsRegenerateForbidden  %+v", 403, o.Payload)
}
func (o *PostVaultCredentialsRegenerateForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostVaultCredentialsRegenerateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVaultCredentialsRegenerateNotFound creates a PostVaultCredentialsRegenerateNotFound with default headers values
func NewPostVaultCredentialsRegenerateNotFound() *PostVaultCredentialsRegenerateNotFound {
	return &PostVaultCredentialsRegenerateNotFound{}
}

/* PostVaultCredentialsRegenerateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostVaultCredentialsRegenerateNotFound struct {
	Payload *models.Error
}

func (o *PostVaultCredentialsRegenerateNotFound) Error() string {
	return fmt.Sprintf("[POST /vault/credentials/regenerate][%d] postVaultCredentialsRegenerateNotFound  %+v", 404, o.Payload)
}
func (o *PostVaultCredentialsRegenerateNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostVaultCredentialsRegenerateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVaultCredentialsRegeneratePreconditionFailed creates a PostVaultCredentialsRegeneratePreconditionFailed with default headers values
func NewPostVaultCredentialsRegeneratePreconditionFailed() *PostVaultCredentialsRegeneratePreconditionFailed {
	return &PostVaultCredentialsRegeneratePreconditionFailed{}
}

/* PostVaultCredentialsRegeneratePreconditionFailed describes a response with status code 412, with default header values.

Precondition Failed
*/
type PostVaultCredentialsRegeneratePreconditionFailed struct {
	Payload *models.Error
}

func (o *PostVaultCredentialsRegeneratePreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /vault/credentials/regenerate][%d] postVaultCredentialsRegeneratePreconditionFailed  %+v", 412, o.Payload)
}
func (o *PostVaultCredentialsRegeneratePreconditionFailed) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostVaultCredentialsRegeneratePreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVaultCredentialsRegenerateInternalServerError creates a PostVaultCredentialsRegenerateInternalServerError with default headers values
func NewPostVaultCredentialsRegenerateInternalServerError() *PostVaultCredentialsRegenerateInternalServerError {
	return &PostVaultCredentialsRegenerateInternalServerError{}
}

/* PostVaultCredentialsRegenerateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostVaultCredentialsRegenerateInternalServerError struct {
	Payload *models.Error
}

func (o *PostVaultCredentialsRegenerateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /vault/credentials/regenerate][%d] postVaultCredentialsRegenerateInternalServerError  %+v", 500, o.Payload)
}
func (o *PostVaultCredentialsRegenerateInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostVaultCredentialsRegenerateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}