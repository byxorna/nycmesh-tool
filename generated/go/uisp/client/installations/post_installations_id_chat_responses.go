// Code generated by go-swagger; DO NOT EDIT.

package installations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// PostInstallationsIDChatReader is a Reader for the PostInstallationsIDChat structure.
type PostInstallationsIDChatReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostInstallationsIDChatReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostInstallationsIDChatOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostInstallationsIDChatUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostInstallationsIDChatForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostInstallationsIDChatInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostInstallationsIDChatServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostInstallationsIDChatOK creates a PostInstallationsIDChatOK with default headers values
func NewPostInstallationsIDChatOK() *PostInstallationsIDChatOK {
	return &PostInstallationsIDChatOK{}
}

/* PostInstallationsIDChatOK describes a response with status code 200, with default header values.

Installation chat message.
*/
type PostInstallationsIDChatOK struct {
	Payload *models.InstallationChatResponse
}

func (o *PostInstallationsIDChatOK) Error() string {
	return fmt.Sprintf("[POST /installations/{id}/chat][%d] postInstallationsIdChatOK  %+v", 200, o.Payload)
}
func (o *PostInstallationsIDChatOK) GetPayload() *models.InstallationChatResponse {
	return o.Payload
}

func (o *PostInstallationsIDChatOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InstallationChatResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInstallationsIDChatUnauthorized creates a PostInstallationsIDChatUnauthorized with default headers values
func NewPostInstallationsIDChatUnauthorized() *PostInstallationsIDChatUnauthorized {
	return &PostInstallationsIDChatUnauthorized{}
}

/* PostInstallationsIDChatUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostInstallationsIDChatUnauthorized struct {
	Payload *models.Error
}

func (o *PostInstallationsIDChatUnauthorized) Error() string {
	return fmt.Sprintf("[POST /installations/{id}/chat][%d] postInstallationsIdChatUnauthorized  %+v", 401, o.Payload)
}
func (o *PostInstallationsIDChatUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostInstallationsIDChatUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInstallationsIDChatForbidden creates a PostInstallationsIDChatForbidden with default headers values
func NewPostInstallationsIDChatForbidden() *PostInstallationsIDChatForbidden {
	return &PostInstallationsIDChatForbidden{}
}

/* PostInstallationsIDChatForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostInstallationsIDChatForbidden struct {
	Payload *models.Error
}

func (o *PostInstallationsIDChatForbidden) Error() string {
	return fmt.Sprintf("[POST /installations/{id}/chat][%d] postInstallationsIdChatForbidden  %+v", 403, o.Payload)
}
func (o *PostInstallationsIDChatForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostInstallationsIDChatForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInstallationsIDChatInternalServerError creates a PostInstallationsIDChatInternalServerError with default headers values
func NewPostInstallationsIDChatInternalServerError() *PostInstallationsIDChatInternalServerError {
	return &PostInstallationsIDChatInternalServerError{}
}

/* PostInstallationsIDChatInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostInstallationsIDChatInternalServerError struct {
	Payload *models.Error
}

func (o *PostInstallationsIDChatInternalServerError) Error() string {
	return fmt.Sprintf("[POST /installations/{id}/chat][%d] postInstallationsIdChatInternalServerError  %+v", 500, o.Payload)
}
func (o *PostInstallationsIDChatInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostInstallationsIDChatInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInstallationsIDChatServiceUnavailable creates a PostInstallationsIDChatServiceUnavailable with default headers values
func NewPostInstallationsIDChatServiceUnavailable() *PostInstallationsIDChatServiceUnavailable {
	return &PostInstallationsIDChatServiceUnavailable{}
}

/* PostInstallationsIDChatServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable
*/
type PostInstallationsIDChatServiceUnavailable struct {
	Payload *models.Error
}

func (o *PostInstallationsIDChatServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /installations/{id}/chat][%d] postInstallationsIdChatServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PostInstallationsIDChatServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostInstallationsIDChatServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}