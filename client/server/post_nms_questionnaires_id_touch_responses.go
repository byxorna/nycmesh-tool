// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/models"
)

// PostNmsQuestionnairesIDTouchReader is a Reader for the PostNmsQuestionnairesIDTouch structure.
type PostNmsQuestionnairesIDTouchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNmsQuestionnairesIDTouchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostNmsQuestionnairesIDTouchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostNmsQuestionnairesIDTouchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostNmsQuestionnairesIDTouchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostNmsQuestionnairesIDTouchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostNmsQuestionnairesIDTouchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostNmsQuestionnairesIDTouchOK creates a PostNmsQuestionnairesIDTouchOK with default headers values
func NewPostNmsQuestionnairesIDTouchOK() *PostNmsQuestionnairesIDTouchOK {
	return &PostNmsQuestionnairesIDTouchOK{}
}

/* PostNmsQuestionnairesIDTouchOK describes a response with status code 200, with default header values.

Successful
*/
type PostNmsQuestionnairesIDTouchOK struct {
	Payload *models.Status
}

func (o *PostNmsQuestionnairesIDTouchOK) Error() string {
	return fmt.Sprintf("[POST /nms/questionnaires/{id}/touch][%d] postNmsQuestionnairesIdTouchOK  %+v", 200, o.Payload)
}
func (o *PostNmsQuestionnairesIDTouchOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PostNmsQuestionnairesIDTouchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsQuestionnairesIDTouchBadRequest creates a PostNmsQuestionnairesIDTouchBadRequest with default headers values
func NewPostNmsQuestionnairesIDTouchBadRequest() *PostNmsQuestionnairesIDTouchBadRequest {
	return &PostNmsQuestionnairesIDTouchBadRequest{}
}

/* PostNmsQuestionnairesIDTouchBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostNmsQuestionnairesIDTouchBadRequest struct {
	Payload *models.Error
}

func (o *PostNmsQuestionnairesIDTouchBadRequest) Error() string {
	return fmt.Sprintf("[POST /nms/questionnaires/{id}/touch][%d] postNmsQuestionnairesIdTouchBadRequest  %+v", 400, o.Payload)
}
func (o *PostNmsQuestionnairesIDTouchBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsQuestionnairesIDTouchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsQuestionnairesIDTouchUnauthorized creates a PostNmsQuestionnairesIDTouchUnauthorized with default headers values
func NewPostNmsQuestionnairesIDTouchUnauthorized() *PostNmsQuestionnairesIDTouchUnauthorized {
	return &PostNmsQuestionnairesIDTouchUnauthorized{}
}

/* PostNmsQuestionnairesIDTouchUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostNmsQuestionnairesIDTouchUnauthorized struct {
	Payload *models.Error
}

func (o *PostNmsQuestionnairesIDTouchUnauthorized) Error() string {
	return fmt.Sprintf("[POST /nms/questionnaires/{id}/touch][%d] postNmsQuestionnairesIdTouchUnauthorized  %+v", 401, o.Payload)
}
func (o *PostNmsQuestionnairesIDTouchUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsQuestionnairesIDTouchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsQuestionnairesIDTouchForbidden creates a PostNmsQuestionnairesIDTouchForbidden with default headers values
func NewPostNmsQuestionnairesIDTouchForbidden() *PostNmsQuestionnairesIDTouchForbidden {
	return &PostNmsQuestionnairesIDTouchForbidden{}
}

/* PostNmsQuestionnairesIDTouchForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostNmsQuestionnairesIDTouchForbidden struct {
	Payload *models.Error
}

func (o *PostNmsQuestionnairesIDTouchForbidden) Error() string {
	return fmt.Sprintf("[POST /nms/questionnaires/{id}/touch][%d] postNmsQuestionnairesIdTouchForbidden  %+v", 403, o.Payload)
}
func (o *PostNmsQuestionnairesIDTouchForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsQuestionnairesIDTouchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsQuestionnairesIDTouchInternalServerError creates a PostNmsQuestionnairesIDTouchInternalServerError with default headers values
func NewPostNmsQuestionnairesIDTouchInternalServerError() *PostNmsQuestionnairesIDTouchInternalServerError {
	return &PostNmsQuestionnairesIDTouchInternalServerError{}
}

/* PostNmsQuestionnairesIDTouchInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostNmsQuestionnairesIDTouchInternalServerError struct {
	Payload *models.Error
}

func (o *PostNmsQuestionnairesIDTouchInternalServerError) Error() string {
	return fmt.Sprintf("[POST /nms/questionnaires/{id}/touch][%d] postNmsQuestionnairesIdTouchInternalServerError  %+v", 500, o.Payload)
}
func (o *PostNmsQuestionnairesIDTouchInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsQuestionnairesIDTouchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
