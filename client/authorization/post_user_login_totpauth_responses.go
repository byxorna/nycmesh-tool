// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/models"
)

// PostUserLoginTotpauthReader is a Reader for the PostUserLoginTotpauth structure.
type PostUserLoginTotpauthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUserLoginTotpauthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUserLoginTotpauthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUserLoginTotpauthBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUserLoginTotpauthInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostUserLoginTotpauthOK creates a PostUserLoginTotpauthOK with default headers values
func NewPostUserLoginTotpauthOK() *PostUserLoginTotpauthOK {
	return &PostUserLoginTotpauthOK{}
}

/* PostUserLoginTotpauthOK describes a response with status code 200, with default header values.

Successful
*/
type PostUserLoginTotpauthOK struct {

	/* User authorization token
	 */
	XAuthToken string

	Payload *models.UserLogin
}

func (o *PostUserLoginTotpauthOK) Error() string {
	return fmt.Sprintf("[POST /user/login/totpauth][%d] postUserLoginTotpauthOK  %+v", 200, o.Payload)
}
func (o *PostUserLoginTotpauthOK) GetPayload() *models.UserLogin {
	return o.Payload
}

func (o *PostUserLoginTotpauthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-auth-token
	hdrXAuthToken := response.GetHeader("x-auth-token")

	if hdrXAuthToken != "" {
		o.XAuthToken = hdrXAuthToken
	}

	o.Payload = new(models.UserLogin)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserLoginTotpauthBadRequest creates a PostUserLoginTotpauthBadRequest with default headers values
func NewPostUserLoginTotpauthBadRequest() *PostUserLoginTotpauthBadRequest {
	return &PostUserLoginTotpauthBadRequest{}
}

/* PostUserLoginTotpauthBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostUserLoginTotpauthBadRequest struct {
	Payload *models.Error
}

func (o *PostUserLoginTotpauthBadRequest) Error() string {
	return fmt.Sprintf("[POST /user/login/totpauth][%d] postUserLoginTotpauthBadRequest  %+v", 400, o.Payload)
}
func (o *PostUserLoginTotpauthBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUserLoginTotpauthBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserLoginTotpauthInternalServerError creates a PostUserLoginTotpauthInternalServerError with default headers values
func NewPostUserLoginTotpauthInternalServerError() *PostUserLoginTotpauthInternalServerError {
	return &PostUserLoginTotpauthInternalServerError{}
}

/* PostUserLoginTotpauthInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostUserLoginTotpauthInternalServerError struct {
	Payload *models.Error
}

func (o *PostUserLoginTotpauthInternalServerError) Error() string {
	return fmt.Sprintf("[POST /user/login/totpauth][%d] postUserLoginTotpauthInternalServerError  %+v", 500, o.Payload)
}
func (o *PostUserLoginTotpauthInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUserLoginTotpauthInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
