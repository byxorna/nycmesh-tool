// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/models"
)

// PostSitesIDImagesReader is a Reader for the PostSitesIDImages structure.
type PostSitesIDImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSitesIDImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSitesIDImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSitesIDImagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostSitesIDImagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostSitesIDImagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSitesIDImagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostSitesIDImagesOK creates a PostSitesIDImagesOK with default headers values
func NewPostSitesIDImagesOK() *PostSitesIDImagesOK {
	return &PostSitesIDImagesOK{}
}

/* PostSitesIDImagesOK describes a response with status code 200, with default header values.

Successful
*/
type PostSitesIDImagesOK struct {
	Payload *models.Status
}

func (o *PostSitesIDImagesOK) Error() string {
	return fmt.Sprintf("[POST /sites/{id}/images][%d] postSitesIdImagesOK  %+v", 200, o.Payload)
}
func (o *PostSitesIDImagesOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PostSitesIDImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSitesIDImagesBadRequest creates a PostSitesIDImagesBadRequest with default headers values
func NewPostSitesIDImagesBadRequest() *PostSitesIDImagesBadRequest {
	return &PostSitesIDImagesBadRequest{}
}

/* PostSitesIDImagesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostSitesIDImagesBadRequest struct {
	Payload *models.Error
}

func (o *PostSitesIDImagesBadRequest) Error() string {
	return fmt.Sprintf("[POST /sites/{id}/images][%d] postSitesIdImagesBadRequest  %+v", 400, o.Payload)
}
func (o *PostSitesIDImagesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSitesIDImagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSitesIDImagesUnauthorized creates a PostSitesIDImagesUnauthorized with default headers values
func NewPostSitesIDImagesUnauthorized() *PostSitesIDImagesUnauthorized {
	return &PostSitesIDImagesUnauthorized{}
}

/* PostSitesIDImagesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostSitesIDImagesUnauthorized struct {
	Payload *models.Error
}

func (o *PostSitesIDImagesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /sites/{id}/images][%d] postSitesIdImagesUnauthorized  %+v", 401, o.Payload)
}
func (o *PostSitesIDImagesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSitesIDImagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSitesIDImagesForbidden creates a PostSitesIDImagesForbidden with default headers values
func NewPostSitesIDImagesForbidden() *PostSitesIDImagesForbidden {
	return &PostSitesIDImagesForbidden{}
}

/* PostSitesIDImagesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostSitesIDImagesForbidden struct {
	Payload *models.Error
}

func (o *PostSitesIDImagesForbidden) Error() string {
	return fmt.Sprintf("[POST /sites/{id}/images][%d] postSitesIdImagesForbidden  %+v", 403, o.Payload)
}
func (o *PostSitesIDImagesForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSitesIDImagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSitesIDImagesInternalServerError creates a PostSitesIDImagesInternalServerError with default headers values
func NewPostSitesIDImagesInternalServerError() *PostSitesIDImagesInternalServerError {
	return &PostSitesIDImagesInternalServerError{}
}

/* PostSitesIDImagesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostSitesIDImagesInternalServerError struct {
	Payload *models.Error
}

func (o *PostSitesIDImagesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /sites/{id}/images][%d] postSitesIdImagesInternalServerError  %+v", 500, o.Payload)
}
func (o *PostSitesIDImagesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSitesIDImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
