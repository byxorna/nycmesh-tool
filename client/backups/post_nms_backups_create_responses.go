// Code generated by go-swagger; DO NOT EDIT.

package backups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/models"
)

// PostNmsBackupsCreateReader is a Reader for the PostNmsBackupsCreate structure.
type PostNmsBackupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNmsBackupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostNmsBackupsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostNmsBackupsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostNmsBackupsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostNmsBackupsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostNmsBackupsCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostNmsBackupsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostNmsBackupsCreateOK creates a PostNmsBackupsCreateOK with default headers values
func NewPostNmsBackupsCreateOK() *PostNmsBackupsCreateOK {
	return &PostNmsBackupsCreateOK{}
}

/* PostNmsBackupsCreateOK describes a response with status code 200, with default header values.

Successful
*/
type PostNmsBackupsCreateOK struct {
	Payload *models.UispBackup
}

func (o *PostNmsBackupsCreateOK) Error() string {
	return fmt.Sprintf("[POST /nms/backups/create][%d] postNmsBackupsCreateOK  %+v", 200, o.Payload)
}
func (o *PostNmsBackupsCreateOK) GetPayload() *models.UispBackup {
	return o.Payload
}

func (o *PostNmsBackupsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UispBackup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsBackupsCreateBadRequest creates a PostNmsBackupsCreateBadRequest with default headers values
func NewPostNmsBackupsCreateBadRequest() *PostNmsBackupsCreateBadRequest {
	return &PostNmsBackupsCreateBadRequest{}
}

/* PostNmsBackupsCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostNmsBackupsCreateBadRequest struct {
	Payload *models.Error
}

func (o *PostNmsBackupsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /nms/backups/create][%d] postNmsBackupsCreateBadRequest  %+v", 400, o.Payload)
}
func (o *PostNmsBackupsCreateBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsBackupsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsBackupsCreateUnauthorized creates a PostNmsBackupsCreateUnauthorized with default headers values
func NewPostNmsBackupsCreateUnauthorized() *PostNmsBackupsCreateUnauthorized {
	return &PostNmsBackupsCreateUnauthorized{}
}

/* PostNmsBackupsCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostNmsBackupsCreateUnauthorized struct {
	Payload *models.Error
}

func (o *PostNmsBackupsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /nms/backups/create][%d] postNmsBackupsCreateUnauthorized  %+v", 401, o.Payload)
}
func (o *PostNmsBackupsCreateUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsBackupsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsBackupsCreateForbidden creates a PostNmsBackupsCreateForbidden with default headers values
func NewPostNmsBackupsCreateForbidden() *PostNmsBackupsCreateForbidden {
	return &PostNmsBackupsCreateForbidden{}
}

/* PostNmsBackupsCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostNmsBackupsCreateForbidden struct {
	Payload *models.Error
}

func (o *PostNmsBackupsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /nms/backups/create][%d] postNmsBackupsCreateForbidden  %+v", 403, o.Payload)
}
func (o *PostNmsBackupsCreateForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsBackupsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsBackupsCreateNotFound creates a PostNmsBackupsCreateNotFound with default headers values
func NewPostNmsBackupsCreateNotFound() *PostNmsBackupsCreateNotFound {
	return &PostNmsBackupsCreateNotFound{}
}

/* PostNmsBackupsCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostNmsBackupsCreateNotFound struct {
	Payload *models.Error
}

func (o *PostNmsBackupsCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /nms/backups/create][%d] postNmsBackupsCreateNotFound  %+v", 404, o.Payload)
}
func (o *PostNmsBackupsCreateNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsBackupsCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsBackupsCreateInternalServerError creates a PostNmsBackupsCreateInternalServerError with default headers values
func NewPostNmsBackupsCreateInternalServerError() *PostNmsBackupsCreateInternalServerError {
	return &PostNmsBackupsCreateInternalServerError{}
}

/* PostNmsBackupsCreateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostNmsBackupsCreateInternalServerError struct {
	Payload *models.Error
}

func (o *PostNmsBackupsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /nms/backups/create][%d] postNmsBackupsCreateInternalServerError  %+v", 500, o.Payload)
}
func (o *PostNmsBackupsCreateInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsBackupsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}