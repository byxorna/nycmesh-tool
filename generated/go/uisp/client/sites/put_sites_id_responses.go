// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// PutSitesIDReader is a Reader for the PutSitesID structure.
type PutSitesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutSitesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutSitesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutSitesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutSitesIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutSitesIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutSitesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutSitesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutSitesIDOK creates a PutSitesIDOK with default headers values
func NewPutSitesIDOK() *PutSitesIDOK {
	return &PutSitesIDOK{}
}

/* PutSitesIDOK describes a response with status code 200, with default header values.

Successful
*/
type PutSitesIDOK struct {
	Payload *models.Site
}

func (o *PutSitesIDOK) Error() string {
	return fmt.Sprintf("[PUT /sites/{id}][%d] putSitesIdOK  %+v", 200, o.Payload)
}
func (o *PutSitesIDOK) GetPayload() *models.Site {
	return o.Payload
}

func (o *PutSitesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Site)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSitesIDBadRequest creates a PutSitesIDBadRequest with default headers values
func NewPutSitesIDBadRequest() *PutSitesIDBadRequest {
	return &PutSitesIDBadRequest{}
}

/* PutSitesIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutSitesIDBadRequest struct {
	Payload *models.Error
}

func (o *PutSitesIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /sites/{id}][%d] putSitesIdBadRequest  %+v", 400, o.Payload)
}
func (o *PutSitesIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutSitesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSitesIDUnauthorized creates a PutSitesIDUnauthorized with default headers values
func NewPutSitesIDUnauthorized() *PutSitesIDUnauthorized {
	return &PutSitesIDUnauthorized{}
}

/* PutSitesIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutSitesIDUnauthorized struct {
	Payload *models.Error
}

func (o *PutSitesIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /sites/{id}][%d] putSitesIdUnauthorized  %+v", 401, o.Payload)
}
func (o *PutSitesIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutSitesIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSitesIDForbidden creates a PutSitesIDForbidden with default headers values
func NewPutSitesIDForbidden() *PutSitesIDForbidden {
	return &PutSitesIDForbidden{}
}

/* PutSitesIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutSitesIDForbidden struct {
	Payload *models.Error
}

func (o *PutSitesIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /sites/{id}][%d] putSitesIdForbidden  %+v", 403, o.Payload)
}
func (o *PutSitesIDForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutSitesIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSitesIDNotFound creates a PutSitesIDNotFound with default headers values
func NewPutSitesIDNotFound() *PutSitesIDNotFound {
	return &PutSitesIDNotFound{}
}

/* PutSitesIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutSitesIDNotFound struct {
	Payload *models.Error
}

func (o *PutSitesIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /sites/{id}][%d] putSitesIdNotFound  %+v", 404, o.Payload)
}
func (o *PutSitesIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutSitesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSitesIDInternalServerError creates a PutSitesIDInternalServerError with default headers values
func NewPutSitesIDInternalServerError() *PutSitesIDInternalServerError {
	return &PutSitesIDInternalServerError{}
}

/* PutSitesIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutSitesIDInternalServerError struct {
	Payload *models.Error
}

func (o *PutSitesIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /sites/{id}][%d] putSitesIdInternalServerError  %+v", 500, o.Payload)
}
func (o *PutSitesIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutSitesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
