// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// PostUserPasswordRequestresetReader is a Reader for the PostUserPasswordRequestreset structure.
type PostUserPasswordRequestresetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUserPasswordRequestresetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUserPasswordRequestresetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUserPasswordRequestresetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUserPasswordRequestresetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostUserPasswordRequestresetOK creates a PostUserPasswordRequestresetOK with default headers values
func NewPostUserPasswordRequestresetOK() *PostUserPasswordRequestresetOK {
	return &PostUserPasswordRequestresetOK{}
}

/* PostUserPasswordRequestresetOK describes a response with status code 200, with default header values.

Successful
*/
type PostUserPasswordRequestresetOK struct {
	Payload *models.Status
}

func (o *PostUserPasswordRequestresetOK) Error() string {
	return fmt.Sprintf("[POST /user/password/requestreset][%d] postUserPasswordRequestresetOK  %+v", 200, o.Payload)
}
func (o *PostUserPasswordRequestresetOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PostUserPasswordRequestresetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserPasswordRequestresetBadRequest creates a PostUserPasswordRequestresetBadRequest with default headers values
func NewPostUserPasswordRequestresetBadRequest() *PostUserPasswordRequestresetBadRequest {
	return &PostUserPasswordRequestresetBadRequest{}
}

/* PostUserPasswordRequestresetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostUserPasswordRequestresetBadRequest struct {
	Payload *models.Error
}

func (o *PostUserPasswordRequestresetBadRequest) Error() string {
	return fmt.Sprintf("[POST /user/password/requestreset][%d] postUserPasswordRequestresetBadRequest  %+v", 400, o.Payload)
}
func (o *PostUserPasswordRequestresetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUserPasswordRequestresetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserPasswordRequestresetInternalServerError creates a PostUserPasswordRequestresetInternalServerError with default headers values
func NewPostUserPasswordRequestresetInternalServerError() *PostUserPasswordRequestresetInternalServerError {
	return &PostUserPasswordRequestresetInternalServerError{}
}

/* PostUserPasswordRequestresetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostUserPasswordRequestresetInternalServerError struct {
	Payload *models.Error
}

func (o *PostUserPasswordRequestresetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /user/password/requestreset][%d] postUserPasswordRequestresetInternalServerError  %+v", 500, o.Payload)
}
func (o *PostUserPasswordRequestresetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUserPasswordRequestresetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}