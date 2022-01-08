// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// PostDevicesIDUnassignReader is a Reader for the PostDevicesIDUnassign structure.
type PostDevicesIDUnassignReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDevicesIDUnassignReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDevicesIDUnassignOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDevicesIDUnassignBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDevicesIDUnassignUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostDevicesIDUnassignForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostDevicesIDUnassignNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDevicesIDUnassignInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDevicesIDUnassignOK creates a PostDevicesIDUnassignOK with default headers values
func NewPostDevicesIDUnassignOK() *PostDevicesIDUnassignOK {
	return &PostDevicesIDUnassignOK{}
}

/* PostDevicesIDUnassignOK describes a response with status code 200, with default header values.

Successful
*/
type PostDevicesIDUnassignOK struct {
	Payload *models.Status
}

func (o *PostDevicesIDUnassignOK) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/unassign][%d] postDevicesIdUnassignOK  %+v", 200, o.Payload)
}
func (o *PostDevicesIDUnassignOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PostDevicesIDUnassignOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDUnassignBadRequest creates a PostDevicesIDUnassignBadRequest with default headers values
func NewPostDevicesIDUnassignBadRequest() *PostDevicesIDUnassignBadRequest {
	return &PostDevicesIDUnassignBadRequest{}
}

/* PostDevicesIDUnassignBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostDevicesIDUnassignBadRequest struct {
	Payload *models.Error
}

func (o *PostDevicesIDUnassignBadRequest) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/unassign][%d] postDevicesIdUnassignBadRequest  %+v", 400, o.Payload)
}
func (o *PostDevicesIDUnassignBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDUnassignBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDUnassignUnauthorized creates a PostDevicesIDUnassignUnauthorized with default headers values
func NewPostDevicesIDUnassignUnauthorized() *PostDevicesIDUnassignUnauthorized {
	return &PostDevicesIDUnassignUnauthorized{}
}

/* PostDevicesIDUnassignUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostDevicesIDUnassignUnauthorized struct {
	Payload *models.Error
}

func (o *PostDevicesIDUnassignUnauthorized) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/unassign][%d] postDevicesIdUnassignUnauthorized  %+v", 401, o.Payload)
}
func (o *PostDevicesIDUnassignUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDUnassignUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDUnassignForbidden creates a PostDevicesIDUnassignForbidden with default headers values
func NewPostDevicesIDUnassignForbidden() *PostDevicesIDUnassignForbidden {
	return &PostDevicesIDUnassignForbidden{}
}

/* PostDevicesIDUnassignForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostDevicesIDUnassignForbidden struct {
	Payload *models.Error
}

func (o *PostDevicesIDUnassignForbidden) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/unassign][%d] postDevicesIdUnassignForbidden  %+v", 403, o.Payload)
}
func (o *PostDevicesIDUnassignForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDUnassignForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDUnassignNotFound creates a PostDevicesIDUnassignNotFound with default headers values
func NewPostDevicesIDUnassignNotFound() *PostDevicesIDUnassignNotFound {
	return &PostDevicesIDUnassignNotFound{}
}

/* PostDevicesIDUnassignNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostDevicesIDUnassignNotFound struct {
	Payload *models.Error
}

func (o *PostDevicesIDUnassignNotFound) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/unassign][%d] postDevicesIdUnassignNotFound  %+v", 404, o.Payload)
}
func (o *PostDevicesIDUnassignNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDUnassignNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDUnassignInternalServerError creates a PostDevicesIDUnassignInternalServerError with default headers values
func NewPostDevicesIDUnassignInternalServerError() *PostDevicesIDUnassignInternalServerError {
	return &PostDevicesIDUnassignInternalServerError{}
}

/* PostDevicesIDUnassignInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostDevicesIDUnassignInternalServerError struct {
	Payload *models.Error
}

func (o *PostDevicesIDUnassignInternalServerError) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/unassign][%d] postDevicesIdUnassignInternalServerError  %+v", 500, o.Payload)
}
func (o *PostDevicesIDUnassignInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDUnassignInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}